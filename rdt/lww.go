package rdt

import (
	"github.com/gritzko/ron"
	"sort"
	"fmt"
)

type LWW struct {

}

// UUID for "lww"
var LWW_UUID = ron.NewName("lww")
var DELTA_UUID = ron.NewName("d")

// LWW arrays and matrices  :)1%)2 :)2   merge is O(N)
func (lww LWW) Reduce (af, bf ron.Frame) (res ron.Frame, err ron.UUID) {

	// just iter, no alloc

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

func (lww LWW) FallbackReduce (inputs []ron.Frame) (res ron.Frame, err ron.UUID) {

	/*object := inputs[0].Object()
	rdtype := inputs[0].Type()
	version := inputs[len(inputs)-1].Event()

	max := make(map[RON.UUID]RON.Op)
	fields := make([]RON.UUID, 0, 128)
	for i:=0; i<len(inputs); i++ {
		f := inputs[i]
		for f.SkipHeader(); !f.EOF(); f.Next() {
			pre, ok := max[f.Ref()]
			if !ok || f.Event().LaterThan(pre.Event()) {
				max[f.Ref()] = f.Op
			}
			if ! ok {
				fields = append(fields, f.Ref())
			}
			fmt.Println()
			for field, op := range max {
				fmt.Println(field.String(),op.String())
			}
		}
	}

	sort.Slice(fields, func(i, j int) bool {
		return fields[j].LaterThan(fields[i])
	})

	res.AppendStateHeader(RON.NewSpec(rdtype, object, version, RON.ZERO_UUID))

	for _, field := range fields {
		res.AppendOp(max[field])
	}*/

	return
}

func MakeLWW () ron.Reducer {
	return LWW{}
}

func init () {
	ron.RDTYPES[LWW_UUID] = MakeLWW
}