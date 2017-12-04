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

func ParseStringBatch(strFrames []string) Batch {
	ret := Batch{}
	for _, s := range strFrames {
		ret = append(ret, ParseFrameString(s))
	}
	return ret
}