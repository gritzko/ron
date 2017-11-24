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

func (lww LWW) Reduce(inputs ron.Batch) (res ron.Frame) {
	heap := ron.MakeFrameHeap(ron.PRIM_LOCATION|ron.SEC_EVENT|ron.SEC_DESC, len(inputs))
	spec := inputs[0].Spec()
	spec.Event = inputs[len(inputs)-1].Event()
	haveState := false
	for k:=0; k<len(inputs); k++ {
		if inputs[k].Ref().IsZero() && inputs[k].IsHeader() {
			haveState = true
		}
		heap.Put(&inputs[k])
	}
	if !haveState {
		spec.Ref = DELTA_UUID
	} else {
		spec.Ref = ron.ZERO_UUID
	}
	res.AppendStateHeader(spec)
	for !heap.IsEmpty() {
		res.AppendReduced(*heap.Current())
		heap.NextPrim()
	}

	return
}

func MakeLWW () ron.Reducer {
	return LWW{}
}

func init () {
	ron.RDTYPES[LWW_UUID] = MakeLWW
}