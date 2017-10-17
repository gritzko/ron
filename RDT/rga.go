package RDT

import (
	"github.com/gritzko/RON"
	//	"fmt"
)

type RGA struct {
	active_ins  RON.FrameHeap
	waiting_rms map[RON.UUID]RON.UUID
	waiting_ins UUIDFrameMultiMap // TODO   UUID2Map? UUIDIntMultiMap
	loc_ins     RON.UUIDHeap   // just sort 'em
}

var RGA_UUID = RON.NewName("rga")

func MakeRGAReducer () RON.Reducer {
	var rga RGA
	rga.active_ins = RON.MakeFrameHeap(RON.PRIM_EVENT|RON.PRIM_DESC|RON.SEC_LOCATION|RON.SEC_DESC, 2)
	rga.waiting_ins = MakeUUIDFrameMultiMap()
	rga.waiting_rms = make(map[RON.UUID]RON.UUID)
	return rga
}

// [ ] multiframe handling:
// 		[ ] Rest(), ranges
//		[ ] FrameHeap EJECT_SUBFRAMES
//		[ ] reinsert subframes
// 		[ ] Split(), another run
//		[ ] violation checks
// [ ] undo/redo

// Reduce RGA frames
// RGA multiframe format:
// ORDER,   WHY
//
func (rga RGA) ReduceAll(inputs []RON.Frame) (result RON.Frame, err RON.UUID) {

	rdtype, object := inputs[0].Type(), inputs[0].Object()

	var version RON.UUID

	for k := 0; k < len(inputs); k++ {
		// TODO check type/object
		raw := !inputs[k].IsHeader()
		at := inputs[k].Ref()
		version = inputs[k].Event()
		if !raw {
			inputs[k].Next()
			// FIXME if over: remove!!!
		}
		if inputs[k].Atoms.Count() == 0 {
			// FIXME what if others are inserts?
			// option: root tree, other trees, then orphan removes
			// break on order violation
			for ii := inputs[k]; !ii.EOF(); ii.Next() {
				pre, ok := rga.waiting_rms[ii.Ref()]
				if !ok || ii.Event().LaterThan(pre) {
					rga.waiting_rms[ii.Ref()] = ii.Event()
				}
			}
		} else { // inserts
			//fmt.Printf("WAIT %s\n", at.String())
			rga.waiting_ins.Put(at, &inputs[k])
			rga.loc_ins.Put(at)
		}
	}
	// multiframe parts must be atomically applied, hence same version id
	header_spec := RON.NewSpec(rdtype, object, version, RON.ZERO_UUID)

	for rga.loc_ins.Len() > 0 {

		loc := rga.loc_ins.PopUnique()
		cu := rga.waiting_ins.Unload(&rga.active_ins, loc)
		if 0 == cu {
			continue
		}
		//fmt.Printf("LOC %s (%d)\n", loc.String(), cu)

		// note any states, if so use ! else ;
		header_spec.SetUUID(RON.SPEC_REF, loc)
		if !loc.IsZero() {
			// ?
		}
		result.AppendStateHeader(header_spec)

		for rga.active_ins.Len() > 0 {
			op := *rga.active_ins.Op()
			event := op.Event()
			atoms := op.Atoms
			if op.IsRaw() {
				header_spec.SetUUID(RON.SPEC_REF, RON.ZERO_UUID)
			} else {
				header_spec.SetUUID(RON.SPEC_REF, op.Ref())
			}
			del, ok := rga.waiting_rms[event]
			if ok && del.LaterThan(op.Ref()) {
				header_spec.SetUUID(RON.SPEC_REF, del)
				delete(rga.waiting_rms, event)
			}
			header_spec.SetUUID(RON.SPEC_EVENT, event)

			result.AppendReduced(header_spec, atoms)
			//fmt.Printf("APPND %c[ %s ]\n", op.Term(), string(op.Atoms.Body))

			rga.active_ins.NextPrim() // idempotency  FIXME rename "prim"
			//fmt.Printf("ACTIVE %d\n", rga.active_ins.Len())

			rga.waiting_ins.Unload(&rga.active_ins, event)
			//fmt.Printf("ADD %s (+%d)\n", event.String(), c)
		}
	}

	if len(rga.waiting_rms) > 0 {
		header_spec.SetUUID(RON.SPEC_REF, RON.NEVER_UUID)
		result.AppendStateHeader(header_spec) // multiframe
		for target, maxEvent := range rga.waiting_rms {
			header_spec.SetUUID(RON.SPEC_EVENT, maxEvent)
			header_spec.SetUUID(RON.SPEC_REF, target)
			result.AppendReduced(header_spec, RON.NO_ATOMS)
			delete(rga.waiting_rms, target)
		}
	}
	// safety: ceil for inserted subtrees - SANITY SCAN!!!

	return
}

func (rga RGA) Reduce(a, b RON.Frame) (res RON.Frame, err RON.UUID) {
	//fmt.Printf("START [ %s ] + [ %s ]\n", a.String(), b.String())
	var frames = [2]RON.Frame{a, b}
	res, err = rga.ReduceAll(frames[0:2])
	return
}

type IMMCell struct {
	p *RON.Frame
	n uint64
}

type UUIDFrameMultiMap struct {
	m map[RON.UUID]IMMCell
	c uint64
}

func MakeUUIDFrameMultiMap() (ret UUIDFrameMultiMap) {
	ret.m = make(map[RON.UUID]IMMCell)
	return
}

func (imm *UUIDFrameMultiMap) Put(key RON.UUID, value *RON.Frame) {
	pre, ok := imm.m[key]
	if ok {
		imm.c++
		synth := RON.NewHashUUID(imm.c, RON.INT60_ERROR)
		imm.m[synth] = pre
		imm.m[key] = IMMCell{value, synth.Value()}
	} else {
		imm.m[key] = IMMCell{value, 0}
	}
}

func (imm UUIDFrameMultiMap) Take(key RON.UUID) (value *RON.Frame, next RON.UUID) {
	pre, ok := imm.m[key]
	if ok {
		delete(imm.m, key)
		value = pre.p
		if pre.n != 0 {
			next = RON.NewHashUUID(pre.n, RON.INT60_ERROR)
		}
	}
	return
}

func (imm UUIDFrameMultiMap) Unload(heap *RON.FrameHeap, key RON.UUID) (count int) {
	for pre, ok := imm.m[key]; ok; pre, ok = imm.m[key] {
		delete(imm.m, key)
		heap.Put(pre.p)
		count++
		if pre.n != 0 {
			key = RON.NewHashUUID(pre.n, RON.INT60_ERROR)
		} else {
			break
		}
	}
	return
}

func init () {
	RON.RDTYPES[RGA_UUID] = MakeRGAReducer
}