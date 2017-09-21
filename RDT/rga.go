package RDT

import (
	"github.com/gritzko/RON"
//	"fmt"
)

type IPChain struct { // 64x2, sweet
	p *RON.Iterator
	n *IPChain
}

type RGA struct {
	active_ins, removes RON.IHeap
	waiting_rms         UUIDMaxMap
	waiting_ins         IMultiMap
	loc_ins             RON.UHeap
}

// UUID for "rga"
var RGA_UUID = RON.UUID{985043671231496192, RON.UUID_NAME_UPPER_BITS}

func (rga RGA) ReduceAll(inputs []RON.Frame) (result RON.Frame, err RON.UUID) {

	rga.active_ins = RON.MakeIHeap(RON.PRIM_EVENT|RON.PRIM_DESC|RON.SEC_LOCATION|RON.SEC_DESC, len(inputs))
	rga.removes = RON.MakeIHeap(RON.PRIM_EVENT|RON.SEC_LOCATION, len(inputs))
	rga.waiting_ins = MakeMultiMap()
	rga.waiting_rms = MakeUUIDMaxMap()
	iii := make([]RON.Iterator, len(inputs))

	var version RON.UUID
	header_spec := RON.Spec{}

	for k := 0; k < len(inputs); k++ {
		iii[k] = inputs[k].Begin()
		header_spec[RON.SPEC_TYPE] = iii[k].Type() //FIXME
		header_spec[RON.SPEC_OBJECT] = iii[k].Object()
		raw := !iii[k].IsHeader()
		at := iii[k].Ref()
		version = iii[k].Event()
		if !raw {
			iii[k].Next()
		}
		if iii[k].Count == 0 {
			rga.removes.Put(&iii[k])
			for ii := iii[k]; !ii.IsEmpty(); ii.Next() {
				rga.waiting_rms.Put(ii.Ref(), ii.Event())
			}
		} else { // inserts
			//fmt.Printf("WAIT %s\n", at.String())
			rga.waiting_ins.Put(at, &iii[k])
			rga.loc_ins.Put(at)
		}
	}
	// multiframe parts must be atomically applied, hence same version id
	header_spec[RON.SPEC_EVENT] = version

	for rga.loc_ins.Len() > 0 {

		loc := rga.loc_ins.TakeUUID()
		cu := rga.waiting_ins.Unload(&rga.active_ins, loc)
		if 0 == cu {
			continue
		}
		//fmt.Printf("LOC %s (%d)\n", loc.String(), cu)

		// note any states, if so use ! else ;
		header_spec[RON.SPEC_REF] = loc
		if !loc.IsZero() {
			// ?
		}
		result.AppendStateHeader(header_spec)

		for rga.active_ins.Len() > 0 {
			op := *rga.active_ins.Op()
			event := op.Event()

			spec := op.Spec
			atoms := op.Atoms
			if op.IsRaw() {
				spec[RON.SPEC_REF] = RON.ZERO_UUID
			}
			del := rga.waiting_rms.Take(event)
			if del.LaterThan(op.Ref()) {
				spec[RON.SPEC_REF] = del
			}
			result.AppendReduced(spec,atoms)
			//fmt.Printf("APPND %c[ %s ]\n", op.Term(), string(op.Atoms.Body))

			rga.active_ins.NextPrim() // idempotency
			//fmt.Printf("ACTIVE %d\n", rga.active_ins.Len())

			rga.waiting_ins.Unload(&rga.active_ins, event)
			//fmt.Printf("ADD %s (+%d)\n", event.String(), c)
		}
	}

	if rga.waiting_rms.Len() > 0 {
		header_spec[RON.SPEC_REF] = RON.NEVER_UUID
		result.AppendStateHeader(header_spec)

		for !rga.removes.IsEmpty() {
			still := rga.waiting_rms.Take(rga.removes.Op().Ref())
			if !still.IsZero() {
				result.AppendOp(*rga.removes.Op())
			}
			rga.removes.Next()
		}
	}
	// safety: ceil for inserted subtrees - SANITY SCAN!!!
	// TODO ensure undo/redo ordering

	return
}

func (rga RGA) Reduce(a, b RON.Frame) (res RON.Frame, err RON.UUID) {
	//fmt.Printf("START [ %s ] + [ %s ]\n", a.String(), b.String())
	var frames = [2]RON.Frame{a, b}
	res, err = rga.ReduceAll(frames[0:2])
	return
}

type UUIDMaxMap struct {
	m map[RON.UUID]RON.UUID
}

func MakeUUIDMaxMap() (ret UUIDMaxMap) {
	ret.m = make(map[RON.UUID]RON.UUID)
	return
}

func (umm *UUIDMaxMap) Put(key, value RON.UUID) {
	pre, ok := umm.m[key]
	if !ok || value.LaterThan(pre) {
		umm.m[key] = value
	}
}

func (umm UUIDMaxMap) Take(key RON.UUID) RON.UUID {
	uuid, ok := umm.m[key]
	if ok {
		delete(umm.m, key)
	}
	return uuid
}

func (umm UUIDMaxMap) Len() int {
	return len(umm.m)
}

type IMMCell struct {
	p *RON.Iterator
	n uint64
}

type IMultiMap struct {
	m map[RON.UUID]IMMCell
	c uint64
}

func MakeMultiMap() (ret IMultiMap) {
	ret.m = make(map[RON.UUID]IMMCell)
	return
}

func (imm *IMultiMap) Put(key RON.UUID, value *RON.Iterator) {
	pre, ok := imm.m[key]
	if ok {
		imm.c++
		synth := RON.UUID{imm.c, RON.INT60_ERROR}
		imm.m[synth] = pre
		imm.m[key] = IMMCell{value, synth.Value}
	} else {
		imm.m[key] = IMMCell{value, 0}
	}
}

func (imm IMultiMap) Take(key RON.UUID) (value *RON.Iterator, next RON.UUID) {
	pre, ok := imm.m[key]
	if ok {
		delete(imm.m, key)
		value = pre.p
		if pre.n != 0 {
			next = RON.UUID{pre.n, RON.INT60_ERROR}
		}
	}
	return
}

func (imm IMultiMap) Unload(heap *RON.IHeap, key RON.UUID) (count int) {
	for pre, ok := imm.m[key]; ok; pre, ok = imm.m[key] {
		delete(imm.m, key)
		heap.Put(pre.p)
		count++
		if pre.n != 0 {
			key = RON.UUID{pre.n, RON.INT60_ERROR}
		} else {
			break
		}
	}
	return
}
