package ron

import (
	"io"
)

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

func MakeStreamedFrame(prealloc_bytes int) (ret Frame) {
	ret.state.data = make([]byte, 0, prealloc_bytes)
	ret.state.streaming = true
	return
}

func StrippedFrame(header Op, body []byte) Frame {
	frame := Frame{}
	frame.state.stripped = true
	frame.Op = header
	frame.state.cs = RON_start
	frame.state.data = body
	return frame
}

func (frame Frame) Reallocate(new_size int) (ret Frame) {
	if new_size < frame.Len() {
		panic("realloc to a smaller size")
	}
	ret = frame
	ret.state.data = make([]byte, frame.Len(), new_size)
	ret.state.streaming = frame.state.streaming
	copy(ret.state.data, frame.state.data)
	return
}

func (frame Frame) Cap() int {
	return cap(frame.state.data)
}

func (frame Frame) Head() Op {
	return frame.Op
}

func (frame Frame) Body() []byte {
	if frame.state.stripped {
		ret := Frame{}
		ret.Format = frame.Format
		ret.AppendOp(frame.Op)
		ret.Append(frame.state.data)
		return ret.Body()
	} else {
		return frame.state.data
	}
}

func (frame *Frame) Read(reader io.Reader) (length int, err error) {
	len, cap := frame.Len(), frame.Cap()
	length, err = reader.Read(frame.state.data[len:cap])
	if length > 0 {
		frame.state.data = frame.state.data[:len+length]
	}
	return
}

func (frame Frame) Fill(clock Clock, env Environment) Frame {
	ret := MakeFrame(frame.Len() << 1)
	now := clock.Time()
	for !frame.EOF() {
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
	for !frame.EOF() {
		ret.AppendOp(frame.Op)
		frame.Next()
	}
	return ret.Close()
}

func (frame Frame) Clone() (clone Frame) {
	clone = frame
	if frame.Atoms.Count()<=2 { // FIXME
		copy(clone._atoms[:frame.Atoms.Count()], frame.atoms)
		clone.atoms = clone._atoms[:frame.Atoms.Count()]
	} else {
		clone.atoms = make([]uint128, frame.Atoms.Count())
		copy(clone.atoms, frame.atoms)
	}
	return
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

func (frame *Frame) Next() bool {
	frame.Parse()
	if frame.state.cs == RON_error {
		return false
	}
	if frame.state.streaming && (frame.state.cs!=RON_start && frame.state.cs!=RON_EOF) {
		return false
	}
	return true
}

func (frame Frame) ParserState() int { // FIXME frame.Parser.State
	return frame.state.cs
}

func (frame Frame) Restart() Frame {
	if frame.state.stripped {
		panic("can't restart") // TODO
	}
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
	s, p int
	//
	position int
	// ragel state
	cs int
	// incomplete uuid/atom data
	incomplete uint128
	// streaming mode switch
	streaming bool
	// whether the head(er) op is present in the body
	stripped bool
}

func (frame Frame) Position() int {
	return frame.state.position // FIXME
}

func (frame Frame) EOF() bool {
	return frame.state.cs == RON_error
}

func (frame *Frame) SkipHeader() {
	if frame.IsHeader() {
		frame.Next()
	}
}

func (frame Frame) Offset() int {
	return frame.state.p
}

// [ ] needs a formal state machine
func (frame Frame) IsComplete() bool {
	return frame.state.cs == RON_start ||
		(!frame.state.streaming && frame.state.p >= len(frame.state.data))
}

// Write a frame to a stream (non-trivial because of event mark rewrites)
func (frame Frame) Write(w io.Writer) error {
	_, err := w.Write(frame.state.data)
	w.Write(FRAME_TERM_ARR[:])
	return err
}

// Write a batch as a multi-frame
func (batch Batch) WriteAll(w io.Writer) (err error) {
	for i := 0; i < len(batch) && err == nil; i++ {
		err = batch[i].Write(w)
	}
	//if err == nil {
	//	w.Write(FRAME_TERM_ARR[:])
	//}
	return
}

func (batch Batch) String() (ret string) {
	for _, frame := range batch {
		ret += frame.String()
	}
	return
}

func (frame *Frame) Append(data []byte) {
	frame.state.data = append(frame.state.data, data...)
}

// Split returns two frames: one from the start to the current position (exclusive),
// another from the current pos (incl) to the end. The right one is "stripped".
func (frame Frame) Split() (left, right Frame) {
	left = ParseFrame(frame.state.data[0:frame.state.s])
	right = StrippedFrame(frame.Op, frame.state.data[frame.state.p:])
	return
}

func (frame Frame) SplitInclusive() Frame {
	if frame.state.stripped {
		panic("oops") // FIXME
	}
	at := frame.state.p
	if at>0 && frame.state.data[at-1]==FRAME_TERM_SEP {
		at -- // strip the frame terminator
	}
	return ParseFrame(frame.state.data[0:at])
}

func (frame Frame) Rest() (rest Frame) {
	rest.state.data = frame.state.data[frame.state.p:]
	rest.state.streaming = true
	return
}

var FRAME_TERM_ARR = [2]byte{FRAME_TERM_SEP, '\n'}
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
