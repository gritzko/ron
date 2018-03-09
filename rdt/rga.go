package rdt

import (
	"github.com/gritzko/ron"
	//	"fmt"
	"sort"
)

// RGA is a Replicated Growable Array data type, an ordered collection of anything
// (numbers, strings, objects). This algorithm is actually closer to Causal Trees,
// albeit that name may turn confusing (RG Arrays are actually trees, Causal Trees
// are actually arrays, but nevermind).
//
// Algorithmically, this differs from Operational Transforms by bruteforcing the
// problem: all the elements of an array have unique ids, so concurrent changes
// can't introduce confusion.
type RGA struct {
	active ron.FrameHeap         // active subtrees, a frame heap
	rms    map[ron.UUID]ron.UUID // removes
	ins    []*ron.Frame          // subtrees-to-insert, ordered by ref
	traps  map[ron.UUID]int      // points to an offset at ins
}

var RGA_UUID = ron.NewName("rga")
var RM_UUID = ron.NewName("rm")
var NO_ATOMS []ron.Atom

func MakeRGAReducer() ron.Reducer {
	var rga RGA
	rga.active = ron.MakeFrameHeap(ron.EventComparatorDesc, ron.RefComparatorDesc, 2)
	rga.rms = make(map[ron.UUID]ron.UUID)
	rga.ins = make([]*ron.Frame, 32)
	rga.traps = make(map[ron.UUID]int)
	return rga
}

func (rga RGA) Features() int {
	return ACID_FULL
}

// [x] multiframe handling: the O(N) multiframe merge
// [ ] undo/redo

func AddMax(rmmap map[ron.UUID]ron.UUID, event, target ron.UUID) {
	rm, ok := rmmap[target]
	if !ok || event.LaterThan(rm) {
		rmmap[target] = event
	}
}

type RefOrderedBatch []*ron.Frame

func (b RefOrderedBatch) Len() int           { return len(b) }
func (b RefOrderedBatch) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b RefOrderedBatch) Less(i, j int) bool { return b[j].Ref().LaterThan(b[i].Ref()) }

type RevOrderedUUIDSlice []ron.UUID

func (b RevOrderedUUIDSlice) Len() int           { return len(b) }
func (b RevOrderedUUIDSlice) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b RevOrderedUUIDSlice) Less(i, j int) bool { return b[i].LaterThan(b[j]) }

// Reduce RGA frames
func (rga RGA) Reduce(batch ron.Batch) ron.Frame {

	rdtype, object := batch[0].Type(), batch[0].Object()
	event := batch[len(batch)-1].Event()
	// multiframe parts must be atomically applied, hence same version id
	spec := ron.NewSpec(rdtype, object, event, ron.ZERO_UUID)
	_produce := [4]ron.Frame{}
	produce := ron.Batch(_produce[:0])
	pending := rga.ins[:0]

	for k := 0; k < len(batch); k++ {
		b := &batch[k]
		if !b.IsHeader() {
			if b.Count() == 0 {
				AddMax(rga.rms, b.Event(), b.Ref())
			} else {
				pending = append(pending, b)
			}
		} else {
			if b.Ref() == RM_UUID { // rm batch, must be the last
				b.Next()
				for !b.EOF() && !b.IsHeader() {
					AddMax(rga.rms, b.Event(), b.Ref())
					b.Next()
				}
			} else {
				pending = append(pending, b)
			}
		}
	}

	sort.Sort(RefOrderedBatch(pending))
	for i := len(pending) - 1; i >= 0; i-- {
		rga.traps[pending[i].Ref()] = i
	}

	for i := 0; i < len(pending); {

		result := ron.MakeFrame(1024)

		at := pending[i].Ref()
		for ; i < len(pending) && !pending[i].EOF() && pending[i].Ref() == at; i++ {
			rga.active.Put(pending[i])
		}
		delete(rga.traps, at)

		spec.SetRef(at)
		spec.SetEvent(event)
		result.AppendStateHeader(spec)

		for !rga.active.EOF() {
			op := rga.active.Current()
			ev := op.Event()
			spec.SetEvent(ev)
			ref := op.Ref()
			if op.IsRaw() {
				ref = ron.ZERO_UUID
			}
			rm, ok := rga.rms[ev]
			if ok {
				if rm.LaterThan(ref) {
					ref = rm
				}
				delete(rga.rms, ev)
			}

			result.AppendReducedRef(ref, *op)
			rga.active.NextPrim()

			for t, ok := rga.traps[ev]; ok && t < len(pending); t++ {
				if !pending[t].EOF() && pending[t].Ref() == ev {
					rga.active.Put(pending[t])
				} else {
					break
				}
			}

		}

		produce = append(produce, result.Rewind())

		for i < len(pending) && pending[i].EOF() {
			i++
		}
	}

	// a separate frame for all the removes we don't have a target for
	if len(rga.rms) > 0 {
		result := ron.MakeFrame(1024)
		spec.SetEvent(event)
		spec.SetRef(RM_UUID)
		result.AppendStateHeader(spec)
		// take removed event ids
		refs := make([]ron.UUID, 0, len(rga.rms))
		for ref := range rga.rms {
			refs = append(refs, ref)
		}
		sort.Sort(RevOrderedUUIDSlice(refs))
		// scan, append
		for _, key := range refs {
			spec.SetRef(key)
			spec.SetEvent(rga.rms[key])
			result.AppendEmptyReducedOp(spec)
			delete(rga.rms, key)
		}
		produce = append(produce, result.Rewind())

	}

	rga.ins = pending[:0] // reuse memory
	for x := range rga.traps {
		delete(rga.traps, x)
	}

	l := len(produce)
	for i := 0; i < len(pending); i++ {
		if !pending[i].EOF() {
			produce = append(produce, pending[i].Split()...)
		}
	}

	if len(produce) == 1 {
		return produce[0]
	} else if l == len(produce) {
		return produce.Join()
	} else {
		//for i:=0; i<len(produce); i++ {
		//	fmt.Printf("    %d %s\n", i, produce[i].String())
		//}
		produce[0], produce[len(produce)-1] = produce[len(produce)-1], produce[0]
		return rga.Reduce(produce)
	}
	// [ ] TODO safety: ceil for inserted subtrees - unified sanity checker
}

type TxtMapper struct {
}

func (txt TxtMapper) Map(batch ron.Batch) string {
	if len(batch) == 0 {
		return ""
	}
	rga := batch[0]
	if rga.Type() != RGA_UUID || !rga.IsHeader() {
		return ""
	}
	ret := []byte{}
	for rga.Next(); !rga.EOF() && !rga.IsHeader(); rga.Next() {
		if rga.Ref().IsZero() {
			ret = append(ret, rga.RawString(0)...)
		}
	}
	return string(ret)
}

func init() {
	ron.RDTYPES[RGA_UUID] = MakeRGAReducer
}
