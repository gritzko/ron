package RDT

import (
	"github.com/gritzko/ron"
	//	"fmt"
)

type RGA struct {
	active_ins  ron.FrameHeap
	waiting_rms map[ron.UUID]ron.UUID
	waiting_ins UUIDFrameMultiMap // TODO   UUID2Map? UUIDIntMultiMap
	loc_ins     ron.UUIDHeap      // just sort 'em
}

var RGA_UUID = ron.NewName("rga")

func MakeRGAReducer () ron.Reducer {
	var rga RGA
	rga.active_ins = ron.MakeFrameHeap(ron.PRIM_EVENT|ron.PRIM_DESC|ron.SEC_LOCATION|ron.SEC_DESC, 2)
	rga.waiting_ins = MakeUUIDFrameMultiMap()
	rga.waiting_rms = make(map[ron.UUID]ron.UUID)
	return rga
}

// [ ] multiframe handling:
// 		[x] Rest(), ranges
//		[ ] FrameHeap EJECT_SUBFRAMES
//		[ ] reinsert subframes
// 		[ ] Split(), another run
//		[ ] violation checks
// [ ] undo/redo

// Reduce RGA frames
// RGA multiframe format:
// ORDER,   WHY
//
func (rga RGA) ReduceAll(inputs []ron.Frame) (result ron.Frame, err ron.UUID) {

	rdtype, object := inputs[0].Type(), inputs[0].Object()

	var version ron.UUID

	for k := 0; k < len(inputs); k++ {
		// TODO check type/object
		raw := !inputs[k].IsHeader()
		at := inputs[k].Ref()
		version = inputs[k].Event()
		if !raw {
			inputs[k].Next()
			// FIXME if over: remove!!!
		}
		if inputs[k].Atoms.Count() == 0 {
			// FIXME what if others are inserts?
			// option: root tree, other trees, then orphan removes
			// break on order violation
			for ii := inputs[k]; !ii.EOF(); ii.Next() {
				pre, ok := rga.waiting_rms[ii.Ref()]
				if !ok || ii.Event().LaterThan(pre) {
					rga.waiting_rms[ii.Ref()] = ii.Event()
				}
			}
		} else { // inserts
			//fmt.Printf("WAIT %s\n", at.String())
			rga.waiting_ins.Put(at, &inputs[k])
			rga.loc_ins.Put(at)
		}
	}
	// multiframe parts must be atomically applied, hence same version id
	header_spec := ron.NewSpec(rdtype, object, version, ron.ZERO_UUID)

	for rga.loc_ins.Len() > 0 {

		loc := rga.loc_ins.PopUnique()
		cu := rga.waiting_ins.Unload(&rga.active_ins, loc)
		if 0 == cu {
			continue
		}
		//fmt.Printf("LOC %s (%d)\n", loc.String(), cu)

		// note any states, if so use ! else ;
		header_spec.SetUUID(ron.SPEC_REF, loc)
		if !loc.IsZero() {
			// ?
		}
		result.AppendStateHeader(header_spec)

		for rga.active_ins.Len() > 0 {
			op := *rga.active_ins.Op()
			event := op.Event()
			atoms := op.Atoms
			if op.IsRaw() {
				header_spec.SetUUID(ron.SPEC_REF, ron.ZERO_UUID)
			} else {
				header_spec.SetUUID(ron.SPEC_REF, op.Ref())
			}
			del, ok := rga.waiting_rms[event]
			if ok && del.LaterThan(op.Ref()) {
				header_spec.SetUUID(ron.SPEC_REF, del)
				delete(rga.waiting_rms, event)
			}
			header_spec.SetUUID(ron.SPEC_EVENT, event)

			result.AppendReduced(header_spec, atoms)
			//fmt.Printf("APPND %c[ %s ]\n", op.Term(), string(op.Atoms.Body))

			rga.active_ins.NextPrim() // idempotency  FIXME rename "prim"
			//fmt.Printf("ACTIVE %d\n", rga.active_ins.Len())

			rga.waiting_ins.Unload(&rga.active_ins, event)
			//fmt.Printf("ADD %s (+%d)\n", event.String(), c)
		}
	}

	if len(rga.waiting_rms) > 0 {
		header_spec.SetUUID(ron.SPEC_REF, ron.NEVER_UUID)
		result.AppendStateHeader(header_spec) // multiframe
		for target, maxEvent := range rga.waiting_rms {
			header_spec.SetUUID(ron.SPEC_EVENT, maxEvent)
			header_spec.SetUUID(ron.SPEC_REF, target)
			result.AppendReduced(header_spec, ron.NO_ATOMS)
			delete(rga.waiting_rms, target)
		}
	}
	// safety: ceil for inserted subtrees - SANITY SCAN!!!

	return
}

func (rga RGA) Reduce(a, b ron.Frame) (res ron.Frame, err ron.UUID) {
	//fmt.Printf("START [ %s ] + [ %s ]\n", a.String(), b.String())
	var frames = [2]ron.Frame{a, b}
	res, err = rga.ReduceAll(frames[0:2])
	return
}

type IMMCell struct {
	p *ron.Frame
	n uint64
}

type UUIDFrameMultiMap struct {
	m map[ron.UUID]IMMCell
	c uint64
}

func MakeUUIDFrameMultiMap() (ret UUIDFrameMultiMap) {
	ret.m = make(map[ron.UUID]IMMCell)
	return
}

func (imm *UUIDFrameMultiMap) Put(key ron.UUID, value *ron.Frame) {
	pre, ok := imm.m[key]
	if ok {
		imm.c++
		synth := ron.NewHashUUID(imm.c, ron.INT60_ERROR)
		imm.m[synth] = pre
		imm.m[key] = IMMCell{value, synth.Value()}
	} else {
		imm.m[key] = IMMCell{value, 0}
	}
}

func (imm UUIDFrameMultiMap) Take(key ron.UUID) (value *ron.Frame, next ron.UUID) {
	pre, ok := imm.m[key]
	if ok {
		delete(imm.m, key)
		value = pre.p
		if pre.n != 0 {
			next = ron.NewHashUUID(pre.n, ron.INT60_ERROR)
		}
	}
	return
}

func (imm UUIDFrameMultiMap) Unload(heap *ron.FrameHeap, key ron.UUID) (count int) {
	for pre, ok := imm.m[key]; ok; pre, ok = imm.m[key] {
		delete(imm.m, key)
		heap.Put(pre.p)
		count++
		if pre.n != 0 {
			key = ron.NewHashUUID(pre.n, ron.INT60_ERROR)
		} else {
			break
		}
	}
	return
}

func init () {
	ron.RDTYPES[RGA_UUID] = MakeRGAReducer
}