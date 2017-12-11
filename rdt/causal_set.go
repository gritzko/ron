package rdt

import "github.com/gritzko/ron"

// Causal set, assumes causally consistent op delivery.
// Hence, no tombstones.
// You can either add or remove an atom/tuple.
// Equal elements possible.
type CausalSet struct {
	heap ron.FrameHeap
}

var CAUSAL_SET_UUID = ron.NewName("cas")

func SetComparator(af, bf *ron.Frame) int64 {
	a, b := af.Event(), bf.Event()
	if !af.Ref().IsZero() {
		a = af.Ref()
	}
	if !bf.Ref().IsZero() {
		b = bf.Ref()
	}
	return -a.Compare(b)
}

func MakeCausalSetReducer() ron.Reducer {
	ret := CausalSet{
		heap: ron.MakeFrameHeap(SetComparator, ron.RefComparatorDesc, 16),
		}
	return ret
}

func (cs CausalSet) Reduce(batch ron.Batch) ron.Frame {
	ret := ron.MakeFrame(batch.Len())
	ref := DELTA_UUID
	if batch.HasFullState() {
		ref = ron.ZERO_UUID
	}
	ret.AppendStateHeader(ron.NewSpec(
		CAUSAL_SET_UUID,
		batch[0].Object(),
		batch[len(batch)-1].Event(),
		ref,
	))
	cs.heap.PutAll(batch)
	for !cs.heap.EOF() {
		if cs.heap.Current().Ref().IsZero() {
			ret.AppendReduced(*cs.heap.Current())
		}
		cs.heap.NextPrim()
	}
	return ret
}