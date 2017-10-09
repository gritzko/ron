package RON

func OpenFrame(data []byte) Frame {
	frame := Frame{}
	frame.state.data = data
	frame.Parse()
	return frame
}

func MakeFormattedFrame(format uint, prealloc_bytes int) (ret Frame) {
	ret.state.data = make([]byte, 0, prealloc_bytes)
	ret.Format = format
	return
}

func MakeFrame(prealloc_bytes int) (ret Frame) {
	ret.state.data = make([]byte, 0, prealloc_bytes)
	return
}

func (frame Frame) Head() Op {
	return frame.Op
}

func (frame Frame) Body() []byte {
	return frame.state.data
}

func (frame Frame) Fill(clock Clock, env Environment) Frame {
	ret := MakeFrame(frame.Len() << 1)
	now := clock.Time()
	for !frame.IsEmpty() {
		spec := frame.Spec
		if spec.uuids[SPEC_EVENT] == ZERO_UUID {
			spec.uuids[SPEC_EVENT] = now
		}
		// TODO implement env fill
		ret.AppendSpecAtomsFlags(spec, frame.Atoms, frame.term)
		frame.Next()
	}
	return ret.Close()
}

func (frame Frame) Reformat(format uint) Frame {
	ret := MakeFrame(frame.Len())
	ret.Format = format
	for !frame.IsEmpty() {
		ret.AppendOp(frame.Op)
		frame.Next()
	}
	return ret.Close()
}

func (frame Frame) Clone() Frame {
	return frame
}

func (frame Frame) String() string {
	return string(frame.Body())
}

func NewBufferFrame(data []byte) (i Frame) {
	i.state.data = data
	i.Parse()
	return
}

func NewStringFrame(data string) (i Frame) {
	return NewBufferFrame([]byte(data))
}

func (frame Frame) IsLast() bool {
	return frame.state.p >= len(frame.state.data)
}

func (frame *Frame) Next() Op {
	frame.Parse()
	return frame.Op
}

func (frame Frame) Restart() Frame {
	return ParseFrame(frame.Body())
}

func (frame Frame) Len() int {
	return len(frame.state.data)
}

type OpParserPos struct {
	// int60 idx, base64 digit idx
	idx, half, digit uint
}

type OpParserState struct {
	OpParserPos
	// the RON frame (for the streaming mode, probably a bit less or a bit more)
	data []byte
	// parser position
	p int
	// ragel state
	cs int
	// incomplete uuid/atom data
	incomplete uint128
	// streaming mode switch
	streaming bool
}

func (frame Frame) EOF() bool {
	return frame.state.cs == RON_error
}

func (frame Frame) Offset() int {
	return frame.state.p
}

func (frame Frame) IsComplete() bool {
	return frame.state.cs == RON_start ||
		(!frame.state.streaming && frame.state.p >= len(frame.state.data))
}

var DIGIT_OFFSETS [11]uint8
var PREFIX_MASKS [11]uint64

func init() {
	var one uint64 = 1
	for i := 0; i < 11; i++ {
		var bitoff uint8 = uint8(60 - i*6)
		DIGIT_OFFSETS[i] = bitoff - 6
		PREFIX_MASKS[i] = ((one << 60) - 1) - ((one << bitoff) - 1)
	}
}
