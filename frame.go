package RON

func OpenFrame(data []byte) (Frame, error) {
	frame := Frame{}
	frame.state.data = data
	err := frame.Parse()
	return frame, err
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

func (frame Frame) Restart_() Frame { /// FIXME ret new
	return ParseFrame(frame.Body())
}

func (frame Frame) Len() int {
	return len(frame.state.data)
}
