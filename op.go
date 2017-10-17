package RON

import (
//	"fmt"
)

func NewSpec (rdtype, object, event, ref UUID) Spec {
	return Spec{uuids:[4]UUID{rdtype, object, event, ref}}
}

func (spec *Spec) SetUUID (idx int, uuid UUID) {
	spec.uuids[idx] = uuid
}

func (a Spec) IsSame(b Spec) bool {
	return a.Event() == b.Event() && a.Object() == b.Object() && a.Ref() == b.Ref() && a.Type() == b.Type()
}

func (spec Spec) UUID(i uint) UUID {
	return spec.uuids[i]
}

func (spec Spec) Type() UUID {
	return spec.uuids[SPEC_TYPE]
}

func (spec Spec) Object() UUID {
	return spec.uuids[SPEC_OBJECT]
}

func (spec Spec) Event() UUID {
	return spec.uuids[SPEC_EVENT]
}

func (spec Spec) Ref() UUID {
	return spec.uuids[SPEC_REF]
}

func (op Op) Term() uint {
	return op.term
}

func (op Op) IsQuery() bool {
	return op.Term() == TERM_QUERY
}

func (op Op) IsHeader() bool {
	return op.Term() == TERM_HEADER
}

func (op Op) IsFramed() bool {
	return op.Term() == TERM_REDUCED
}

func (op Op) IsRaw() bool {
	return op.Term() == TERM_RAW
}

func (op Op) IsOn() bool {
	return op.IsQuery() && op.Ref() != NEVER_UUID
}

func (op Op) IsOff() bool {
	return op.IsQuery() && op.Ref() == NEVER_UUID
}

func CreateFrame(rdtype, object, event, location, value string) Frame {
	return Frame{}
}

func (frame Frame) Stamp(clock Clock) Frame {
	cur := MakeFrame(frame.Len() + 20)
	stamps := map[uint64]UUID{}
	for !frame.IsEmpty() {
		op := frame.Op
		for t := uint(0); t < 4; t++ {
			uuid := op.Spec.UUID(t)
			if uuid.IsTemplate() {
				stamp, ok := stamps[uuid.Origin()]
				if !ok {
					stamp = clock.Time()
					stamps[uuid.Origin()] = stamp
				}
				op.uuids[t] = stamp
			}
		}
		cur.AppendOp(op)
		frame.Next()
	}
	return cur.Close()
}

func (op Op) IsEmpty() bool {
	return op.IsSame(ZERO_OP.Spec)
}

func (spec *Op) isZero() bool {
	for t := 0; t < 4; t++ {
		if !spec.uuids[t].IsZero() {
			return false
		}
	}
	return true
}

func (op Op) String() string {
	cur := MakeFrame(128)
	cur.AppendOp(op)
	return string(cur.Body())
}
