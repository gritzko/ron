package RON

import (
	"github.com/pkg/errors"
)


func (a *UUID) Compare(b UUID) int {
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
	return a.Type == b.Type && a.Object == b.Object &&
		a.Event == b.Event && a.Location == b.Location
}

func (op *Op) IsHeader() bool {
	return op.AtomTypes[0] == '!'
}

// not good - op is detached from a frame here
func CreateOp(rdtype, object, event, location UUID, value string) (ret Op, err error) {
	l := XParseOp([]byte(value), &ret, &ZERO_OP)
	if l<=0 {
		err = errors.New("invalid atom string")
		return
	}
	ret.Type = rdtype
	ret.Object = object
	ret.Event = event
	ret.Location = location
	return
}

func CreateFrame(rdtype, object, event, location, value string) Frame {
	return Frame{}
}


func (i *Iterator) Next() (op *Op, ok bool) {
	//input := i.frame.Body[i.pos:]
	//readOp(input, &i.Op)
	return
}

func (i *Iterator) Clone() Iterator {
	return Iterator{}
}

func (frame *Frame) First() Iterator {
	return Iterator{}
}

func (frame *Frame) Last() Iterator {
	return Iterator{}
}
func (i *Iterator) End() bool {
	return true
}

func MakeFrame (prealloc_bytes int) Frame {
	var buf = make([]byte, prealloc_bytes)
	return Frame{buf, Iterator{}}
}