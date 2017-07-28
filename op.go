package RON

import (
	"github.com/pkg/errors"
)

func Compare(a, b UUID) int {
	diff := a.Value - b.Value
	if diff == 0 {
		diff = uint64(a.Sign) - uint64(b.Sign)
		if diff == 0 {
			diff = a.Origin - b.Origin
		}
	}
	if diff < 0 {
		return -1
	} else if diff > 0 {
		return 1
	} else {
		return 0
	}
}

func (op *Op) Empty() bool {
	return op.AtomCount == 0
}

func (a *Op) Same(b *Op) bool {
	return a.uuids == b.uuids
}

func (op *Op) IsHeader() bool {
	return op.AtomTypes[0] == '!'
}

// not good - op is detached from a frame here
func CreateOp(rdtype, object, event, location UUID, value string) (ret Op, err error) {
	l := XParseOp([]byte(value), &ret, ZERO_OP)
	if l <= 0 {
		err = errors.New("invalid atom string")
		return
	}
	ret.uuids = [4]UUID{rdtype, object, event, location}
	return
}

func CreateFrame(rdtype, object, event, location, value string) Frame {
	return Frame{}
}

func (i *Iterator) Next() bool {

	if i.AtEnd() {
		return false
	} else if i.IsLast() {
		i.Op = ZERO_OP
		i.Op.AtomCount = 0
		return false
	}
	var prev Op = i.Op
	l := XParseOp(i.frame.Body[i.offset:], &i.Op, prev)
	i.offset += l

	return i.AtEnd()
}

func (i *Iterator) Rest () [] byte {
	return []byte{}
}

func (frame *Frame) Begin() (i Iterator) {
	i.frame = frame
	i.Op = ZERO_OP // TODO  ZERO_OP is exactly Op{}
	i.Next()
	return
}

func (frame *Frame) End() (i Iterator) {
	i.frame = frame
	i.offset = len(frame.Body)
	return
}

// A frame's end position is an op having a value of !!! and UUIDs from
// the last valid op (zeroes for an empty frame).
// The end op may be explicit, i.e. actually exist in the frame.
// An explicit end op can not be abbreviated.
func (i Iterator) AtEnd() bool {
	return i.offset>0 && i.AtomCount==0
	//i.AtomTypes[0] == '!' && i.AtomTypes[1] == '!' && i.AtomTypes[2] == '!'
}

func (i Iterator) IsLast () bool {
	return i.offset >= len(i.frame.Body)
}

func MakeFrame(prealloc_bytes int) Frame {
	var buf = make([]byte, 0, prealloc_bytes)
	return Frame{buf, ZERO_OP}
}

func (op *Op) GetUUIDp (i int) *UUID {
	return & op.uuids[i]
}

func (op *Op) GetUUID (i int) UUID {
	return op.uuids[i]
}

func (uuid *UUID) IsZero() bool {
	return uuid.Value==0 && uuid.Origin==0
}

func (spec *Op) isZero () bool {
	for t:=0; t<4; t++ {
		if !spec.uuids[t].IsZero() {
			return false
		}
	}
	return true
}