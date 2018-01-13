package ron

import (
//	"fmt"
)

func (frame Frame) IsSameOp(b Frame) bool {
	return frame.Event() == b.Event() && frame.Object() == b.Object() && frame.Ref() == b.Ref() && frame.Type() == b.Type()
}

func (frame Frame) Type() UUID {
	return frame.UUID(SPEC_TYPE)
}

func (frame Frame) Object() UUID {
	return frame.UUID(SPEC_OBJECT)
}

func (frame Frame) Event() UUID {
	return frame.UUID(SPEC_EVENT)
}

func (frame Frame) Ref() UUID {
	return frame.UUID(SPEC_REF)
}

func (frame Frame) Term() int {
	return frame.term
}

func (frame Frame) IsQuery() bool {
	return frame.Term() == TERM_QUERY
}

func (frame Frame) IsHeader() bool {
	return frame.Term() == TERM_HEADER
}

func (frame Frame) IsFullState() bool {
	return frame.IsHeader() && frame.Ref().IsZero()
}

func (frame Frame) IsFramed() bool {
	return frame.Term() == TERM_REDUCED
}

func (frame Frame) IsRaw() bool {
	return frame.Term() == TERM_RAW
}

func (frame Frame) IsOn() bool {
	return frame.IsQuery() && frame.Ref() != NEVER_UUID
}

func (frame Frame) IsOff() bool {
	return frame.IsQuery() && frame.Ref() == NEVER_UUID
}

func CreateFrame(rdtype, object, event, location, value string) Frame {
	return Frame{}
}

func (frame Frame) Stamp(clock Clock) Frame {
	cur := MakeFrame(frame.Len() + 20)
	stamps := map[uint64]UUID{}
	for !frame.EOF() {
		for t := 0; t < 4; t++ {
			uuid := frame.UUID(t)
			if uuid.IsTemplate() {
				stamp, ok := stamps[uuid.Origin()]
				if !ok {
					stamp = clock.Time()
					stamps[uuid.Origin()] = stamp
				}
				frame.atoms[t] = Atom(stamp)
			}
		}
		cur.Append(frame)
		frame.Next()
	}
	return cur.Close()
}

func (frame Frame) isZero() bool {
	for t := 0; t < 4; t++ {
		if !frame.UUID(t).IsZero() {
			return false
		}
	}
	return true
}

func (frame Frame) OpString() string {
	cur := MakeFrame(128)
	cur.Append(frame)
	return string(cur.Body)
}

func (frame Frame) Origin() uint64 {
	return frame.atoms[SPEC_EVENT][1]
}

func (frame Frame) Time() uint64 {
	return frame.atoms[SPEC_EVENT][0]
}
