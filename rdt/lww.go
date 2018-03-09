package rdt

import (
	"github.com/gritzko/ron"
)

// LWW is a last-write-wins replicated data type that may host a variety of user-land data types, like:
//
// 		* a dictionary,
// 		* a struct or
// 		* a simple 1D array (no splice, no index shifts),
// 		* a simple 2D array.
//
// This LWW employs client-side logical timestamps to decide which write wins, on a field-by-field basis.
// That is similar to e.g. Cassandra LWW.
//
type LWW struct {
}

var LWW_UUID = ron.NewName("lww")
var DELTA_UUID = ron.NewName("d")

func (lww LWW) Features() int {
	return ACID_FULL
}

func (lww LWW) Reduce(inputs ron.Batch) (res ron.Frame) {
	heap := ron.MakeFrameHeap(ron.RefComparator, ron.EventComparatorDesc, len(inputs))
	spec := inputs[0].Spec()
	spec.SetEvent(inputs[len(inputs)-1].Event())
	if inputs.HasFullState() {
		spec.SetRef(ron.ZERO_UUID)
	} else {
		spec.SetRef(DELTA_UUID)
	}
	for k := 0; k < len(inputs); k++ {
		heap.Put(&inputs[k])
	}
	res.AppendStateHeader(spec)
	for !heap.EOF() {
		res.AppendReduced(*heap.Current())
		heap.NextPrim()
	}

	return
}

func MakeLWW() ron.Reducer {
	return LWW{}
}

func init() {
	ron.RDTYPES[LWW_UUID] = MakeLWW
}
