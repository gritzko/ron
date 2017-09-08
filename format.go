package RON

import "math/bits"

const (
	FORMAT_ZIP = 1<<iota
	FORMAT_GRID
	FORMAT_SPACE
	FORMAT_HEADER_SPACE
	FORMAT_SKIP_EQ
	FORMAT_REDEFAULT
	FORMAT_OP_LINES
	FORMAT_FRAME_LINES
)
//FORMAT_CONDENSED = 1 << iota
//FORMAT_OP_LINES
//FORMAT_FRAMES
//FORMAT_TABLE
const SPACES22 = "                      "
const SPACES88 = SPACES22+SPACES22+SPACES22+SPACES22
const ZEROS10 = "0000000000"

func unzipPrefixSeparator(input []byte) (prefix uint8, length int) {
	var i = ABC[input[0]]
	if i <= -10 {
		prefix = uint8(-i-10) + 4
		length = 1
	}
	return
}

func (t UUID) Equal(b UUID) bool {
	return t.Value == b.Value && t.Origin == b.Origin
}

// FormatInt outputs a 60-bit "Base64x64" int into the output slice
func FormatInt(output []byte, value uint64) []byte {
	tail := bits.TrailingZeros64(value)
	if tail>54 {
		tail = 54
	}
	tail -= tail%6
	for i:=54; i>=tail; i-=6 {
		output = append( output, BASE64[(value>>uint(i))&63] )
	}
	return output
}

func FormatZipInt(output []byte, value, context uint64) []byte {
	prefix := Int60Prefix(value,context)
	if prefix==60 {
		return output
	}
	if prefix >= 4*6 {
		prefix -= prefix%6
		value = (value << uint(prefix)) & INT60_FULL
		pchar := prefixBits2Sep ( uint(prefix)/6 - 4 )
		output = append(output, pchar)
		if value!=0 {
			output = FormatInt(output, value)
		}
	} else {
		output = FormatInt(output, value)
	}
	return output
}

func Int60Prefix (a, b uint64) int {
	return bits.LeadingZeros64((a^b)&INT60_FULL)-4
}

func (uuid UUID) ZipString(context UUID) string {
	var arr [INT60LEN*2+2]byte
	ret := FormatZipInt(arr[:0], uuid.Value, context.Value)
	if uuid.Origin != UUID_NAME_UPPER_BITS {
		ret = append(ret, uuid.Sign())
		at := len(ret)
		ret = FormatZipInt(ret, uuid.Replica(), context.Replica())
		if uuid.Scheme()==context.Scheme() && at>1 {
			if len(ret)>at && ABC[ret[at]]<0 {
				copy(ret[at-1:], ret[at:])
				ret = ret[:len(ret)-1]
			} else if len(ret)==at {
				ret = ret[:len(ret)-1]
			}
		}
	}
	return string(ret)
}

func (uuid UUID) String() (ret string) {
	ret = uuid.ZipString(ZERO_UUID)
	if len(ret)==0 {
		ret = "0"
	}
	return
}

func (frame *Frame) appendUUID(buf []byte, uuid UUID, context UUID) []byte {
	if uuid == context /*&& uuid != ZERO_UUID*/ {
		return buf
	}
	start := len(buf)
	buf = FormatZipInt(buf, uuid.Value, context.Value)
	if uuid.Origin == UUID_NAME_UPPER_BITS {
		return buf
	}
	buf = append(buf, uuid.Sign())
	at := len(buf)
	buf = FormatZipInt(buf, uuid.Origin, context.Origin)
	// sometimes, we may skip UUID separator (+-%$)
	if uuid.Scheme()==context.Scheme() && at>start+1 {
		if len(buf)>at && ABC[buf[at]]<0 {
			copy(buf[at-1:], buf[at:])
			buf = buf[:len(buf)-1]
		} else if len(buf)==at && ABC[buf[start]]<0 {
			buf = buf[:len(buf)-1]
		}
	}
	return buf
}

func (uuid UUID) prefixWith (context UUID) (ret int) {
	vp := bits.LeadingZeros64(uuid.Value^context.Value)
	vp -= vp%6
	op := bits.LeadingZeros64((uuid.Origin^context.Origin)&INT60_FULL)
	op -= op%6
	ret = vp + op
	if uuid.Scheme()!=context.Scheme() {
		ret--
	}
	return
}

func (frame *Frame) appendSpec(spec, context Spec) {
	buf := frame.Body
	start := len(buf)
	flags := frame.Format
	for t := 0; t < 4; t++ {
		if 0!=flags&FORMAT_GRID {
			rest := t*22 - (len(buf)-start)
			buf = append(buf, SPACES88[:rest]...)
		} else if 0!=flags&FORMAT_SPACE && t>0 {
			buf = append(buf, ' ')
		}
		if spec[t] == context[t] /*&& 0!=flags&FORMAT_SKIP_EQ*/ {
			continue
		}
		buf = append(buf, specBits2Sep(uint(t)))
		if t>0 && 0!=flags&FORMAT_REDEFAULT {
			ctxAt := 0
			ctxUUID := spec[t-1]
			ctxPL := spec[t].prefixWith(ctxUUID)
			for i:=1; i<4; i++ {
				pl := spec[t].prefixWith(context[i])
				if pl > ctxPL {
					ctxPL = pl
					ctxUUID = context[i]
					ctxAt = i
				}
			}
			if ctxAt != t {
				buf = append(buf, redefBits2Sep(uint(ctxAt)))
			}
			buf = frame.appendUUID(buf, spec[t], ctxUUID)
		} else {
			buf = frame.appendUUID(buf, spec[t], context[t])
		}
	}
	frame.Body = buf
}

func (op Op) String() string {
	var frame Frame
	frame.AppendOp(op)
	return string(frame.Body)
}

func (frame *Frame) String() string {
	return string(frame.Body)
}

func (frame *Frame) AppendOp(op Op) {

	flags := frame.Format
	if len(frame.Body)>0 && ( 0!=flags&FORMAT_OP_LINES || (0!=flags&FORMAT_FRAME_LINES && op.IsHeader()) ) {
		frame.Body = append(frame.Body, '\n')
	} else if 0!=flags&FORMAT_HEADER_SPACE && frame.last.IsHeader() {
		frame.Body = append(frame.Body, ' ')
	}


	frame.appendSpec(op.Spec, frame.last.Spec)

	frame.Body = append(frame.Body, op.Body[op.Offsets[0]:]...)

	if op.IsHeader() || (op.Class()==OP_RAW && frame.last.Class()!=OP_RAW) || op.Count==0 {
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
