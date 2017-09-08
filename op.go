package RON

import (
	"github.com/pkg/errors"
//	"fmt"
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

func (a UUID) LaterThan(b UUID) bool {
	if a.Value == b.Value {
		return a.Origin > b.Origin
	} else {
		return a.Value > b.Value
	}
}

func (a UUID) EarlierThan(b UUID) bool {
	// FIXME define through Compare
	if a.Value == b.Value {
		return a.Origin < b.Origin
	} else {
		return a.Value < b.Value
	}
}

func (a UUID) Scheme() uint {
	return uint(a.Origin >> 60)
}

func (a UUID) Sign() byte {
	return uuidBits2Sep(a.Scheme())
}

func (a UUID) Replica() uint64 {
	return a.Origin & PREFIX10
}

func (a UUID) SameAs(b UUID) bool {
	if a.Value != b.Value {
		return false
	} else if a.Origin == b.Origin {
		return true
	} else if (a.Origin^b.Origin)&PREFIX10 != 0 {
		return false
	} else {
		return a.Origin&UUID_UPPER_BITS == b.Origin&UUID_UPPER_BITS
	}
}

func (a Op) Same(b *Op) bool {
	return a.Spec == b.Spec
}

func (op Op) Term() byte {
	return opBits2Sep(op.Class())
}

func (op Op) Class () uint {
	return 3 & op.Flags
}

func (op Op) IsQuery () bool {
	return op.Flags & OP_QUERY_BIT != 0
}

func (op Op) IsHeader() bool {
	return op.Class()&OP_STATE_BIT != 0
}

func (op Op) IsFramed() bool {
	return op.Class()==OP_REDUCED
}

func (op Op) IsState() bool {
	return op.Class() == OP_STATE
}

func (op Op) IsRaw() bool {
	return op.Class() == OP_RAW
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

	if i.offset!=0 && i.IsEmpty() { // FIXME test corner cases more
		return false
	} else if i.IsLast() {
		i.Op = ZERO_OP
		return false
	}
	var prev Op = i.Op
	l := XParseOp(i.frame.Body[i.offset:], &i.Op, prev)

	if l > 0 {
		//fmt.Printf("PARSED [ %s ] REST [ %s ]\n", string(i.frame.Body[i.offset:i.offset+l]), string(i.frame.Body[i.offset+l:]))
		i.offset += l
		return true
	} else {
		i.Op = ZERO_OP
		i.offset = l - i.offset
		return false
	}
}

func (uuid UUID) IsTemplate() bool {
	return uuid.Sign()==UUID_NAME && uuid.Value == 0 && uuid.Origin != 0
}

func (frame Frame) Stamp(clock Clock) (ret Frame) {
	stamps := map[uint64]UUID{}
	i := frame.Begin()
	for !i.IsEmpty() {
		op := i.Op
		for t := 0; t < 4; t++ {
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
	return op.Spec[0]==ZERO_UUID && op.Spec[1]==ZERO_UUID && op.Spec[2]==ZERO_UUID && op.Spec[3]==ZERO_UUID
}

func (frame *Frame) Begin() (i Iterator) {
	i.frame = frame
	i.offset = 0
	if frame.first.IsEmpty() {
		i.Op = ZERO_OP
		i.Next()
	} else {
		i.Op = frame.first
	}
	return
}

func (frame *Frame) End() Iterator {
	return Iterator{frame: frame, offset: len(frame.Body)}
}

// A frame's end position is an op having a value of !!! and UUIDs from
// the last valid op (zeroes for an empty frame).
// The end op may be explicit, i.e. actually exist in the frame.
// An explicit end op can not be abbreviated.
//func (i Iterator) AtEnd() bool {
//	return i.offset>0 && i.Count==0
//	//i.AtomTypes[0] == '!' && i.AtomTypes[1] == '!' && i.AtomTypes[2] == '!'
//}

func (i Iterator) IsLast() bool {
	return i.offset >= len(i.frame.Body)
}

func MakeFrame(prealloc_bytes int) Frame {
	var buf = make([]byte, 0, prealloc_bytes)
	return Frame{Body: buf}
}

func (op *Spec) GetUUIDp(i int) *UUID {
	return &op[i]
}

func (op *Spec) GetUUID(i int) UUID {
	return op[i]
}

func (uuid *UUID) IsZero() bool {
	return uuid.Value == 0 && uuid.Origin == 0
}

func (spec *Op) isZero() bool {
	for t := 0; t < 4; t++ {
		if !spec.Spec[t].IsZero() {
			return false
		}
	}
	return true
}

func (i Iterator) Offset() int {
	return i.offset
}

func (a Atoms) Type (i uint) uint {
	return (a.Types >> (i<<1)) & 3
}

func NewEventUUID (time, origin uint64) UUID {
	return UUID{Value:time, Origin:(origin&INT60_ERROR)|UUID_EVENT_UPPER_BITS}
}