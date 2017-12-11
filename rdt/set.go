package rdt

import "github.com/gritzko/ron"

// Set, fully commutative, with tombstones.
// You can either add or remove an atom/tuple.
// Equal elements possible.
type Set struct {
	heap ron.FrameHeap
}

var SET_UUID = ron.NewName("set")

func MakeSetReducer() ron.Reducer {
	ret := Set{
		heap: ron.MakeFrameHeap(SetComparator, ron.RefComparatorDesc, 16),
	}
	return ret
}

func (cs Set) Reduce(batch ron.Batch) ron.Frame {
	ret := ron.MakeFrame(batch.Len())
	ref := DELTA_UUID
	if batch.HasFullState() {
		ref = ron.ZERO_UUID
	}
	ret.AppendStateHeader(ron.NewSpec(
		SET_UUID,
		batch[0].Object(),
		batch[len(batch)-1].Event(),
		ref,
	))
	cs.heap.PutAll(batch)
	for !cs.heap.EOF() {
		ret.AppendReduced(*cs.heap.Current())
		cs.heap.NextPrim()
	}
	return ret
}

func init () {
	ron.RDTYPES[SET_UUID] = MakeSetReducer
}