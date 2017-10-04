package RON

import ()

func ParseUUIDString(uuid string) (ret UUID, err error) {
	return ZERO_UUID.Parse([]byte(uuid))
}

func ParseUUID(data []byte) (uuid UUID, err error) {
	return ZERO_UUID.Parse(data)
}

func ParseOp(data []byte, context Op) (op Op, length int) {
	op = context
	// FIXME length = XParseOp(data, &op, context)
	return
}

func ParseFrame(data []byte) (ret Frame) {
	ret.body = data
	return
}

func ParseFrameString(frame string) Frame {
	return ParseFrame([]byte(frame))
}

func Parse(str string) (Frame, error) {
	ret := Frame{body: []byte(str)}
	_ = ret.Begin() // FIXME iterator - errors
	return ret, nil
}

// SplitMultiframe scans a frame detecting any headers; all resulting
// frames are returned in a slice. In case the frame is a monoframe,
// the return slice is empty. The sanity checker is invoked on every
// op, on error the function aborts (all the completed frames still
// in the slice).
func (frame Frame) SplitMultiframe(sanity Checker) (ret []Frame, err error) {
	from := frame.Begin()
	till := from
	for !till.IsEmpty() {
		if sanity != nil {
			err = sanity.Check(till)
			if err != nil {
				return
			}
		}
		prev := till // FIXME!!!
		till.Next()
		if !till.IsEmpty() && till.IsHeader() {
			next := MakeFrame(128)
			next.AppendRange(from, prev)
			//			ret = append(ret, next)
			from = till
		}
	}
	return
	// TODO make slice frames (head op not in the body), avoid copy
}
