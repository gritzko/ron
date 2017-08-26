package RON

const (
	FORMAT_FRAME_NL = 1 << iota
	FORMAT_HEADER_SPACE
	FORMAT_OP_SPACE
	FORMAT_OP_TAB
	FORMAT_UUID_SPACE
	FORMAT_OP_NL
	FORMAT_OP_INDENT
	FORMAT_ZIP_REDEF
	FORMAT_ZIP_PREFIX
	FORMAT_SKIP_REPEATS
)

func unzipPrefixSeparator(input []byte) (prefix uint8, length int) {
	var i = ABC[input[0]]
	if i <= -10 {
		prefix = uint8(-i-10) + 4
		length = 1
	}
	return
}

func UnzipBase64(input []byte, number *uint64) int {

	// TODO migrate to Ragel

	var l int = 10
	if l > len(input) {
		l = len(input)
	}
	var i = 0
	var res uint64
	for ; i < l; i++ {
		code := ABC[input[i]]
		if code < 0 {
			break
		}
		res <<= 6
		res |= uint64(code)
	}
	if number != nil {
		*number = res
	}
	return i
}

func (t UUID) Equal(b UUID) bool {
	return t.Value == b.Value && t.Origin == b.Origin
}

func CommonPrefix(value, context uint64) uint {
	// TODO use math.bits
	var xor = value ^ context
	if xor >= 1<<(6*6) {
		return 0
	}
	if xor == 0 {
		return 10
	}
	if xor >= 1<<(3*6) { // 456
		if xor >= 1<<(5*6) {
			return 4
		} else if xor >= 1<<(4*6) {
			return 5
		} else {
			return 6
		}
	} else { // 789
		if xor >= 1<<(2*6) {
			return 7
		} else if xor >= 1<<(1*6) {
			return 8
		} else {
			return 9
		}
	}
}

func ZeroTail(value *uint64) (tail uint) {
	if *value&((1<<30)-1) == 0 {
		tail += 5
		*value >>= 30
	}
	if *value&((1<<18)-1) == 0 {
		tail += 3
		*value >>= 18
	}
	if *value&((1<<12)-1) == 0 {
		tail += 2
		*value >>= 12
	}
	if tail < 10 && *value&((1<<6)-1) == 0 {
		tail += 1
		*value >>= 6
	}
	return
}

const prefix_mask uint64 = 0xffffff << 36

func ZipUUIDString(uuid, context UUID) string {
	var arr [INT60LEN*2+2]byte
	ret := FormatZippedUUID(arr[:0], uuid, context)
	return string(ret)
}

func (uuid UUID) String() (ret string) {
	ret = ZipUUIDString(uuid, ZERO_UUID)
	if len(ret)==0 {
		ret = "0"
	}
	return
}

func FormatTrimmedInt(output []byte, value uint64) (ret []byte) {
	if value == 0 {
		ret = append(output, '0')
		return
	}
	l := 10
	if value&((1<<24)-1) == 0 {
		value >>= 24
		l -= 4
	}
	for value&63 == 0 {
		value >>= 6
		l--
	}
	var vals [10]byte
	add := vals[:l]
	k := l
	for k > 0 {
		k--
		add[k] = base64[value&63]
		value >>= 6
	}
	ret = append(output, add...)
	return
}

func FormatInt(output []byte, value uint64) (ret []byte) {
	ret = FormatTrimmedInt(output, value)
	l := len(ret) - len(output)
	for l < 10 {
		ret = append(ret, '0')
		l++
	}
	return
}

func FormatZippedInt(output []byte, value, context uint64) (ret []byte) {
	var prefix uint = CommonPrefix(value, context)
	ret = output
	if prefix < 4 {
		ret = FormatTrimmedInt(ret, value)
	} else {
		if prefix == 10 {
			return
		}
		ret = append(ret, PREFIX_PUNCT[prefix-4])
		value = (value << (prefix * 6)) & PREFIX10
		if value != 0 {
			ret = FormatTrimmedInt(ret, value)
		}
	}
	return
}

func FormatUUID(output []byte, uuid UUID) []byte {
	ret := FormatTrimmedInt(output, uuid.Value)
	ret = append(ret, UUID_PUNCT[uuid.Sign()])
	ret = FormatTrimmedInt(ret, uuid.Origin)
	return ret
}

func FormatZippedUUID(output []byte, uuid UUID, context UUID) (ret []byte) {

	if uuid == context && uuid != ZERO_UUID { // FIXME options
		return output
	}
	ret = FormatZippedInt(output, uuid.Value, context.Value)
	if uuid.Origin == UUID_NAME_UPPER_BITS {
		return ret
	}
	if uuid.Value == context.Value || uuid.Sign() != context.Sign() ||
		(uuid.Origin&prefix_mask) != (context.Origin&prefix_mask) ||
		(uuid.Replica() == context.Replica() && ABC[ret[len(output)]]>=0) { // FIXME this if
		ret = append(ret, UUID_PUNCT[uuid.Scheme()])
	}
	if uuid.Replica() != context.Replica() {
		ret = FormatZippedInt(ret, uuid.Replica(), context.Replica())
	}
	return
}

func FormatSpec(output []byte, op Op) []byte {
	// expand to 88+values
	for t := 0; t < 4; t++ {
		output = append(output, SPEC_PUNCT[t])
		output = FormatUUID(output, op.GetUUID(t))
	}
	return output
}

func FormatZippedSpec(output []byte, op Op, context Op) []byte {
	// expand to 88+values
	for t := 0; t < 4; t++ {
		if op.GetUUID(t) == context.GetUUID(t) {
			continue
		}
		output = append(output, SPEC_PUNCT[t])
		output = FormatZippedUUID(output, op.GetUUID(t), context.GetUUID(t))
	}
	return output
}

// optimize for close values
// context==nil is valid?
func FormatOp(output []byte, op Op, context Op) []byte {
	output = FormatZippedSpec(output, op, context)
	from := op.Offsets[0]
	//copy(output, op.Body[from:])
	output = append(output, op.Body[from:]...)
	//off += len(op.Body) - from
	if op.Class()!=OP_REDUCED || op.Count==0 {
		output = append(output, op.Term())
	}
	return output
}

func (op Op) String() string {
	var arr[INT60LEN*8+8+40]byte
	buf := FormatOp(arr[:0], op, ZERO_OP)
	return string(buf)
}

func (frame *Frame) String() string {
	return string(frame.Body)
}


func (frame *Frame) AppendOp(op Op) {
	if (0!=frame.Format&FORMAT_FRAME_NL) && len(frame.Body)>0 && !op.IsFramed() {
		frame.Body = append(frame.Body, '\n')
	}
	var uuid_arr [11 * 2 * 4]byte
	uuids := uuid_arr[:0]
	if !frame.last.isZero() || len(frame.Body) == 0 {
		uuids = FormatZippedSpec(uuids, op, frame.last)
	} else {
		uuids = FormatSpec(uuids, op)
	}
	if (0!=frame.Format&FORMAT_HEADER_SPACE) && frame.last.IsHeader() && op.IsFramed() {
		frame.Body = append(frame.Body, ' ')
	}
	frame.Body = append(frame.Body, uuids...)
	frame.Body = append(frame.Body, op.Body[op.Offsets[0]:]...)
	if op.Term()!=',' || op.Count==0 {
		frame.Body = append(frame.Body, op.Term())
	}
	frame.last = op
}

func (frame *Frame) AppendSpecAtomsFlags(spec Spec, atoms Atoms, flags uint) {
	frame.AppendOp(Op{spec,atoms, flags})
}

func (frame *Frame) AppendReduced(spec Spec, atoms Atoms) {
	frame.AppendOp(Op{spec,atoms, OP_REDUCED})
}

func (frame *Frame) AppendRaw(spec Spec, atoms Atoms) {
	frame.AppendOp(Op{spec,atoms, OP_RAW})
}

func (frame *Frame) AppendSpecBody(toel Spec, body []byte, flags uint) {
	frame.AppendSpecAtomsFlags(toel, ParseAtoms(body), flags)
}

func (frame *Frame) AppendPatchHeader (spec Spec) {
	frame.AppendSpecAtomsFlags(spec, NO_ATOMS, OP_PATCH)
}

func (frame *Frame) AppendStateHeader (spec Spec) {
	frame.AppendSpecAtomsFlags(spec, NO_ATOMS, OP_STATE)
}

func (frame *Frame) AppendRange(i, j Iterator) {
	if !i.IsEmpty() && !j.IsEmpty() && i.offset >= j.offset {
		return
	}
	if i.IsEmpty() {
		return
	}
	if i.frame != j.frame {
		panic("mismatching iterators")
	}
	frame.AppendOp(i.Op)
	from := i.offset //+ len(i.Body)
	till := j.offset
	if till > from { // more than 1 op
		frame.Body = append(frame.Body, i.frame.Body[from:till]...)
	}
}

func (frame *Frame) AppendAll(i Iterator) {
	if i.IsEmpty() {
		return
	}
	frame.AppendOp(i.Op)
	from := i.offset // + len(i.Body)
	frame.Body = append(frame.Body, i.frame.Body[from:]...)
	frame.last = i.frame.last
}

func (frame *Frame) AppendFrame(second Frame) {
	frame.AppendRange(second.Begin(), second.End())
}

//func (frame *Frame) AppendEnd() {
// no explicit end marker for now
//}

func (frame Frame) Clone() Frame {
	body := make([]byte, 0, len(frame.Body))
	copy(body, frame.Body)
	return Frame{Body: body, last: frame.last}
}
