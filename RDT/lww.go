package RDT

import (
	"github.com/gritzko/RON"
)

type LWW struct {

}

var LWW_UUID = RON.UUID{881557636825219072, RON.NAME_SIGN_BITS}

// LWW arrays and matrices  :)1%)2 :)2   merge is O(N)
func (lww LWW) Reduce (af, bf RON.Frame) (res RON.Frame, err RON.UUID) {
	frames := [2]RON.Frame{af, bf}
	return lww.ReduceAll(frames[0:2])
}

func (lww LWW) ReduceAll(inputs []RON.Frame) (res RON.Frame, err RON.UUID) {
	heap := RON.MakeIHeap(RON.PRIM_LOCATION|RON.SEC_EVENT|RON.SEC_DESC, len(inputs))
	var spec RON.Spec
	for k:=0; k<len(inputs); k++ {
		i := inputs[k].Begin()
		spec = i.Spec
		if i.IsHeader() {
			i.Next()
		}
		heap.Put(&i)
	}
	spec[RON.SPEC_LOCATION] = RON.ZERO_UUID
	res.AppendSpecAtoms(spec, RON.STATE_HEADER_ATOMS)
	for !heap.IsEmpty() {
		atoms := heap.Op().Atoms
		atoms.Types[RON.MAX_ATOMS] = RON.OP_SEP
		res.AppendSpecAtoms(heap.Op().Spec, atoms)
		heap.NextPrim()
	}

	return
}

