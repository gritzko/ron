package RON

const UUID_SEPS = ".#@:"
const (
	TYPE_UUID_SEP     = '.'
	OBJECT_UUID_SEP   = '#'
	EVENT_UUID_SEP    = '@'
	LOCATION_UUID_SEP = ':'
)

const (
	PREFIX10 = (1<<60 - 1) - (1<<(iota*6) - 1)
	PREFIX9
	PREFIX8
	PREFIX7
	PREFIX6
	PREFIX5
	PREFIX4
	PREFIX3
	PREFIX2
	PREFIX1
)

type UUID struct {
	Value  uint64 // FIXME rename (confl op values)
	Sign   byte   // TODO maybe fit into 16 bytes
	Origin uint64
}

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

type Op struct {
	Type, Object, Event, Location UUID
	AtomCount                     int
	AtomTypes                     [8]byte
	Body                          []byte
	AtomOffsets                   [8]int
}

type Frame struct {
	Body []byte
	last Iterator
}

type Iterator struct {
	Op
	frame *Frame
	pos   int
}

type Reducer func(a Iterator, b Iterator) Frame

type RawUUID []byte

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

func CreateOp(rdtype, object, event, location, value string) Op {
	// no compression - easy
	// add zeros
	return ZERO_OP
}

func CreateFrame(rdtype, object, event, location, value string) Frame {
	return Frame{}
}

var ZERO_OP Op

func init () {
    ZERO_OP.Type = ZERO_UUID
    ZERO_OP.Object = ZERO_UUID
    ZERO_OP.Event = ZERO_UUID
    ZERO_OP.Location = ZERO_UUID
}

func readUIUD(input []byte, uuid []byte) int {
	return 0
}

func readOp(input []byte, op *Op) int {
	// regex - would solve whitespaces, uniformity, nuances
	// redefault swaps
	//if uuids[3] {
	//	readUUID(uuids[3], op.Type())
	//} perf :(
	// pick values

	//var off int
	//if input[off]==TYPE_UUID_SEP {
	//	off++
	//	off += readUUID(input[off:], op.Type())
	//}
	//if input[off]==OBJECT_UUID_SEP {
	//	off++
	//	off += readUUID(input[off:], op.Object())
	//}
	return 0
}

func (i *Iterator) Next() (op *Op, ok bool) {
	input := i.frame.Body[i.pos:]
	readOp(input, &i.Op)
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

var EMPTY_FRAME Frame

func (frame *Frame) Append(second Frame) Frame {
	// bigger frames: skip compression
	// if frame is small || have last => parse, peek
	// else append()
	// if second is small => remember last
	// logairthmic event => repack
	return EMPTY_FRAME
}

func (frame *Frame) AppendOp(i Iterator) Frame {
	// last==0 => either parse or skip abbrev
	// future TODO
	// end-op  .lww#id@ev:loc!!! - optional
	//         either implicit or explicit (retain explicit)
	return Frame{}
}

func (frame *Frame) AppendAll(i Iterator) Frame {
	frame.AppendOp(i)
	// add the rest as a chunk, last=0
	return Frame{}
}

func (i *Iterator) End() bool {
	return true
}
