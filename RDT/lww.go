package RDT

import (
	"github.com/gritzko/RON"
)

type LWW struct {

}

// UUID for "lww"
var LWW_UUID = RON.NewName("lww")
var DELTA_UUID = RON.NewName("d")

// LWW arrays and matrices  :)1%)2 :)2   merge is O(N)
func (lww LWW) Reduce (af, bf RON.Frame) (res RON.Frame, err RON.UUID) {
	frames := [2]RON.Frame{af, bf}
	return lww.ReduceAll(frames[0:2])
}

func (lww LWW) ReduceAll(inputs []RON.Frame) (res RON.Frame, err RON.UUID) {
	heap := RON.MakeFrameHeap(RON.PRIM_LOCATION|RON.SEC_EVENT|RON.SEC_DESC, len(inputs))
	var spec RON.Spec
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
		spec.SetUUID(RON.SPEC_REF, DELTA_UUID)
	} else {
		spec.SetUUID(RON.SPEC_REF, RON.ZERO_UUID)
	}
	res.AppendStateHeader(spec)
	for !heap.IsEmpty() {
		atoms := heap.Op().Atoms
		res.AppendReduced(heap.Op().Spec, atoms)
		heap.NextPrim()
	}

	return
}

func MakeLWW () RON.Reducer {
	return LWW{}
}

func init () {
	RON.RDTYPES[LWW_UUID] = MakeLWW
}