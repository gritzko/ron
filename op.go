package RON

import (
	"github.com/pkg/errors"
)

func Compare(a, b UUID) int {
	diff := int64(a.Value) - int64(b.Value)
	if diff == 0 {
		diff = int64(a.Origin) - int64(b.Origin)
	}
	if diff < 0 {
		return -1
	} else if diff > 0 {
		return 1
	} else {
		return 0
	}
}

func (a UUID) LaterThan (b UUID) bool {
	if a.Value==b.Value {
		return a.Origin > b.Origin
	} else {
		return a.Value > b.Value
	}
}

func (a UUID) EarlierThan (b UUID) bool {
	if a.Value==b.Value {
		return a.Origin < b.Origin
	} else {
		return a.Value < b.Value
	}
}

func (a UUID) Sign () uint64 {
	return a.Origin >> 60
}

func (a UUID) Replica () uint64 {
	return a.Origin & PREFIX10
}

func (a UUID) SameAs (b UUID) bool {
	if a.Value!=b.Value {
		return false
	} else if a.Origin==b.Origin {
		return true
	} else if (a.Origin^b.Origin) & PREFIX10 != 0 {
		return false
	} else {
		return a.Origin&EVENT_SIGN_BIT==1 && b.Origin&EVENT_SIGN_BIT==1
	}
}

func (op *Op) Empty() bool {
	return op.Count == 0
}

func (a *Op) Same(b *Op) bool {
	return a.Spec == b.Spec
}

func (op *Op) IsHeader() bool {
	return op.Types[0] == '!'
}

// not good - op is detached from a frame here
func CreateOp(rdtype, object, event, location UUID, value string) (ret Op, err error) {
	l := XParseOp([]byte(value), &ret, ZERO_OP)
	if l <= 0 {
		err = errors.New("invalid atom string")
		return
	}
	ret.Spec = Spec{rdtype, object, event, location}
	return
}

func CreateFrame(rdtype, object, event, location, value string) Frame {
	return Frame{}
}

func (i *Iterator) Next() bool {

	if i.IsEmpty() {
		return false
	} else if i.IsLast() {
		i.Op = ZERO_OP
		return false
	}
	var prev Op = i.Op
	l := XParseOp(i.frame.Body[i.offset:], &i.Op, prev)

	if l>0 {
		i.offset += l
		return true
	} else {
		i.Op = ZERO_OP
		return false
	}
}

func (uuid UUID) IsTemplate () bool {
	return uuid.Sign()==NAME_SIGN && uuid.Value==0 && uuid.Origin!=0
}

func (frame Frame) Stamp (clock Clock) (ret Frame) {
	stamps := map[uint64]UUID{}
	i := frame.Begin()
	for !i.IsEmpty() {
		op := i.Op
		for t:=0; t<4; t++ {
			uuid := op.Spec[t]
			if uuid.IsTemplate() {
				stamp, ok := stamps[uuid.Origin]
				if !ok {
					stamp = clock.Time()
					stamps[uuid.Origin] = stamp
				}
				op.Spec[t] = stamp
			}
		}
		ret.AppendOp(op)
		i.Next()
	}
	return
}

func (op Op) IsEmpty() bool {
	return op.Count==0
}

func (frame *Frame) Begin() (i Iterator) {
	i.frame = frame
	i.offset = 0
	if frame.first.IsEmpty() {
		i.Op = ZERO_OP
		i.Count=1 // HACK
		i.Next()
	} else {
		i.Op = frame.first
	}
	return
}

func (frame *Frame) End() Iterator {
	return Iterator{frame:frame, offset:len(frame.Body)}
}

// A frame's end position is an op having a value of !!! and UUIDs from
// the last valid op (zeroes for an empty frame).
// The end op may be explicit, i.e. actually exist in the frame.
// An explicit end op can not be abbreviated.
func (i Iterator) AtEnd() bool {
	return i.offset>0 && i.Count==0
	//i.AtomTypes[0] == '!' && i.AtomTypes[1] == '!' && i.AtomTypes[2] == '!'
}

func (i Iterator) IsLast () bool {
	return i.offset >= len(i.frame.Body)
}

func MakeFrame(prealloc_bytes int) Frame {
	var buf = make([]byte, 0, prealloc_bytes)
	return Frame{Body: buf}
}

func (op *Op) GetUUIDp (i int) *UUID {
	return & op.Spec[i]
}

func (op *Op) GetUUID (i int) UUID {
	return op.Spec[i]
}

func (uuid *UUID) IsZero() bool {
	return uuid.Value==0 && uuid.Origin==0
}

func (spec *Op) isZero () bool {
	for t:=0; t<4; t++ {
		if !spec.Spec[t].IsZero() {
			return false
		}
	}
	return true
}
