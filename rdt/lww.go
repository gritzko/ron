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

// LWW arrays and matrices  :)1%)2 :)2   merge is O(N)
func (lww LWW) Reduce (af, bf ron.Frame) (res ron.Frame, err ron.UUID) {
	frames := [2]ron.Frame{af, bf}
	return lww.ReduceAll(frames[0:2])
}

func (lww LWW) ReduceAll(inputs []ron.Frame) (res ron.Frame, err ron.UUID) {
	heap := ron.MakeFrameHeap(ron.PRIM_LOCATION|ron.SEC_EVENT|ron.SEC_DESC, len(inputs))
	var spec ron.Spec
	haveState := false
	for k:=0; k<len(inputs); k++ {
		i := inputs[k]
		spec = i.Spec
		if i.Ref().IsZero() && i.IsHeader() {
			haveState = true
		}
		if i.IsHeader() {
			i.Next()
		}
		heap.Put(&i)
	}
	if !haveState {
		spec.SetUUID(ron.SPEC_REF, DELTA_UUID)
	} else {
		spec.SetUUID(ron.SPEC_REF, ron.ZERO_UUID)
	}
	res.AppendStateHeader(spec)
	for !heap.IsEmpty() {
		atoms := heap.Op().Atoms
		res.AppendReduced(heap.Op().Spec, atoms)
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