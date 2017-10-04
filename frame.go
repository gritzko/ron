package RON

func (frame Frame) Begin() (i Iterator) {
	i.Op = frame.Op
	i.state.data = frame.body
	i.state.p = frame.offset
	return
}

func (frame Frame) Head() Op {
	return frame.Op
}

func (frame *Frame) Fill(clock Clock, env Environment) Frame {
	ret := MakeFrame(len(frame.body) << 1)
	now := clock.Time()
	i := frame.Begin()
	for !i.IsEmpty() {
		spec := i.Spec
		if spec.uuids[SPEC_EVENT] == ZERO_UUID {
			spec.uuids[SPEC_EVENT] = now
		}
		// TODO implement env fill
		ret.AppendSpecAtomsFlags(spec, i.Atoms, i.term)
		i.Next()
	}
	return ret.Close()
}

func (frame Frame) Reformat(format uint) Frame {
	ret := MakeFrame(len(frame.body))
	ret.Format = format
	i := frame.Begin()
	for !i.IsEmpty() {
		ret.AppendOp(i.Op)
		i.Next()
	}
	return ret.Close()
}

func (frame Frame) Clone() Frame {
	body := make([]byte, 0, len(frame.body))
	copy(body, frame.body)
	return Frame{body: body}
}

func (frame *Frame) String() string {
	return string(frame.body)
}
