package rdt

import "github.com/gritzko/ron"

type VV struct {
	heap ron.FrameHeap
}

var VV_UUID = ron.NewName("vv")

func VVComparator(a, b *ron.Frame) int64 {
	return int64(a.Origin()) - int64(b.Origin())
}

func MakeVVReducer() ron.Reducer {
	return VV{
		heap: ron.MakeFrameHeap(VVComparator, ron.EventComparatorDesc, 2),
	}
}

func (vv VV) Features() int {
	return ron.ACID_FULL
}

func (vv VV) Reduce(batch ron.Batch) ron.Frame {
	spec := ron.NewSpec(
		VV_UUID,
		batch[0].Object(),
		batch[len(batch)-1].Event(),
		ron.ZERO_UUID,
	)
	vv.heap.PutAll(batch)
	re := ron.NewFrame()
	re.AppendStateHeader(spec)
	for !vv.heap.EOF() {
		spec[ron.SPEC_EVENT] = vv.heap.Current().Atoms()[ron.SPEC_EVENT]
		re.AppendEmpty(spec, ron.TERM_REDUCED)
		vv.heap.NextPrim()
	}
	return re
}

func init() {
	ron.RDTYPES[VV_UUID] = MakeVVReducer
}
