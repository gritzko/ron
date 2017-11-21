package ron

import ()

func ParseUUIDString(uuid string) (ret UUID, err error) {
	return ZERO_UUID.Parse([]byte(uuid))
}

func ParseUUID(data []byte) (uuid UUID, err error) {
	return ZERO_UUID.Parse(data)
}

func ParseFrame(data []byte) Frame { // TODO swap with OpenFrame
	return OpenFrame(data)
}

func ParseFrameString(frame string) Frame {
	return ParseFrame([]byte(frame))
}

// SplitMultiframe scans a frame detecting any headers; all resulting
// frames are returned in a slice. In case the frame is a monoframe,
// the return slice is empty. The sanity checker is invoked on every
// op, on error the function aborts (all the completed frames still
// in the slice).
//func (frame Frame) SplitMultiframe(sanity Checker) (ret []Frame, err error) {
//
//
//	// FIXME format and copy!!!
//
//
//	for !frame.IsEmpty() {
//		if sanity != nil {
//			err = sanity.Check(frame)
//			if err != nil {
//				return
//			}
//		}
//		if frame.IsHeader() {
//			ret = append(ret, Frame{})
//		}
//		ret[len(ret)-1].Append(frame)
//		frame.Next()
//	}
//	return
//	// TODO make slice frames (head op not in the body), avoid copy
//}
