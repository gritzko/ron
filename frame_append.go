package ron

import (
	"fmt"
	"math/bits"
)

const (
	FORMAT_UNZIP = 1 << iota
	FORMAT_GRID
	FORMAT_SPACE
	FORMAT_HEADER_SPACE
	FORMAT_NOSKIP
	FORMAT_REDEFAULT
	FORMAT_OP_LINES
	FORMAT_FRAME_LINES
	FORMAT_INDENT
)
const FRAME_FORMAT_CARPET = FORMAT_GRID | FORMAT_SPACE | FORMAT_OP_LINES | FORMAT_NOSKIP | FORMAT_UNZIP
const FRAME_FORMAT_TABLE = FORMAT_GRID | FORMAT_SPACE | FORMAT_OP_LINES
const FRAME_FORMAT_LIST = FORMAT_OP_LINES | FORMAT_INDENT
const FRAME_FORMAT_LINE = FORMAT_FRAME_LINES | FORMAT_HEADER_SPACE

//FORMAT_CONDENSED = 1 << iota
//FORMAT_OP_LINES
//FORMAT_FRAMES
//FORMAT_TABLE
const SPACES22 = "                      "
const SPACES88 = SPACES22 + SPACES22 + SPACES22 + SPACES22
const ZEROS10 = "0000000000"

// FormatInt outputs a 60-bit "Base64x64" int into the output slice
func FormatInt(output []byte, value uint64) []byte {
	tail := bits.TrailingZeros64(value)
	if tail > 54 {
		tail = 54
	}
	tail -= tail % 6
	for i := 54; i >= tail; i -= 6 {
		output = append(output, BASE64[(value>>uint(i))&63])
	}
	return output
}

func FormatZipInt(output []byte, value, context uint64) []byte {
	prefix := Int60Prefix(value, context)
	if prefix == 60 {
		return output
	}
	if prefix >= 4*6 {
		prefix -= prefix % 6
		value = (value << uint(prefix)) & INT60_FULL
		pchar := PREFIX_PUNCT[uint(prefix)/6-4]
		output = append(output, pchar)
		if value != 0 {
			output = FormatInt(output, value)
		}
	} else {
		output = FormatInt(output, value)
	}
	return output
}

func Int60Prefix(a, b uint64) int {
	return bits.LeadingZeros64((a^b)&INT60_FULL) - 4
}

func FormatUUID(buf []byte, uuid UUID) []byte {
	buf = FormatInt(buf, uuid.Value())
	if uuid.Origin() != UUID_NAME_UPPER_BITS {
		buf = append(buf, uuid.Sign())
		buf = FormatInt(buf, uuid.Replica())
	}
	return buf
}

func FormatZipUUID(buf []byte, uuid, context UUID) []byte {
	start := len(buf)
	buf = FormatZipInt(buf, uuid.Value(), context.Value())
	if uuid.Origin() == UUID_NAME_UPPER_BITS {
		return buf
	}
	buf = append(buf, uuid.Sign())
	at := len(buf)
	buf = FormatZipInt(buf, uuid.Origin(), context.Origin())
	// sometimes, we may skip UUID separator (+-%$)
	if uuid.Scheme() == context.Scheme() && at > start+1 {
		if len(buf) > at && ABC_KIND[buf[at]] != BASE_KIND {
			copy(buf[at-1:], buf[at:])
			buf = buf[:len(buf)-1]
		} else if len(buf) == at && ABC_KIND[buf[start]] != BASE_KIND {
			buf = buf[:len(buf)-1]
		}
	}
	return buf
}

// TODO FormatInt Float String

func (frame *Frame) appendUUID(uuid UUID, context UUID) {
	if 0 != frame.Format&FORMAT_UNZIP {
		frame.state.data = FormatUUID(frame.state.data, uuid)
	} else if uuid != context {
		frame.state.data = FormatZipUUID(frame.state.data, uuid, context)
	}
}

func (frame *Frame) appendSpec(spec, context Spec) {
	start := len(frame.state.data)
	flags := frame.Format
	for t := 0; t < 4; t++ {
		if 0 != flags&FORMAT_GRID {
			rest := t*22 - (len(frame.state.data) - start)
			frame.state.data = append(frame.state.data, SPACES88[:rest]...)
		} else if 0 != flags&FORMAT_SPACE && t > 0 {
			frame.state.data = append(frame.state.data, ' ')
		}
		if (spec.uuids[t] == context.uuids[t]) && (0 == flags&FORMAT_NOSKIP) {
			continue
		}
		frame.state.data = append(frame.state.data, SPEC_PUNCT[uint(t)])
		if t > 0 && 0 != flags&FORMAT_REDEFAULT {
			ctxAt := 0
			ctxUUID := spec.uuids[t-1]
			ctxPL := spec.uuids[t].prefixWith(ctxUUID)
			for i := 1; i < 4; i++ {
				pl := spec.uuids[t].prefixWith(context.uuids[i])
				if pl > ctxPL {
					ctxPL = pl
					ctxUUID = context.uuids[i]
					ctxAt = i
				}
			}
			if ctxAt != t {
				frame.state.data = append(frame.state.data, REDEF_PUNCT[uint(ctxAt)])
			}
			frame.appendUUID(spec.uuids[t], ctxUUID)
		} else {
			frame.appendUUID(spec.uuids[t], context.uuids[t])
		}
	}
}

func (frame *Frame) appendAtoms(a Atoms) {
	for i := 0; i < a.Count(); i++ {
		switch a.AType(i) {
		case ATOM_INT:
			{
				frame.state.data = append(frame.state.data, ATOM_INT_SEP)
				s := fmt.Sprint(a.atoms[i][0])
				frame.state.data = append(frame.state.data, []byte(s)...)
			}
		case ATOM_STRING:
			{
				frame.state.data = append(frame.state.data, ATOM_STRING_SEP)
				ft := a.atoms[i]
				frame.state.data = append(frame.state.data, a.frame[ft[0]&INT60_FULL:ft[1]&INT60_FULL]...)
				frame.state.data = append(frame.state.data, ATOM_STRING_SEP)
			}
		case ATOM_FLOAT:
			{
				frame.state.data = append(frame.state.data, ATOM_FLOAT_SEP)
				s := fmt.Sprint(a.atoms[i][0])
				frame.state.data = append(frame.state.data, []byte(s)...)
				frame.state.data = append(frame.state.data, '.')
			}
		case ATOM_UUID:
			{
				frame.state.data = append(frame.state.data, ATOM_UUID_SEP)
				frame.appendUUID(UUID{a.atoms[i]}, ZERO_UUID)
			}
		}
	}
}

func (frame *Frame) AppendOp(op Op) {

	flags := frame.Format
	start := len(frame.state.data)
	if len(frame.state.data) > 0 && (0 != flags&FORMAT_OP_LINES || (0 != flags&FORMAT_FRAME_LINES && !op.IsFramed())) {
		frame.state.data = append(frame.state.data, '\n')
		if 0 != flags&FORMAT_INDENT && !op.IsHeader() {
			frame.state.data = append(frame.state.data, "    "...)
		}
	} else if 0 != flags&FORMAT_HEADER_SPACE && frame.Op.IsHeader() {
		frame.state.data = append(frame.state.data, ' ')
	}

	frame.appendSpec(op.Spec, frame.Op.Spec)

	if 0 != flags&FORMAT_GRID {
		rest := 4*22 - (len(frame.state.data) - start)
		frame.state.data = append(frame.state.data, SPACES88[:rest]...)
	}

	frame.appendAtoms(op.Atoms)

	if op.IsHeader() || (op.IsRaw() && !frame.Op.IsRaw()) || op.Atoms.Count() == 0 {
		frame.state.data = append(frame.state.data, TERM_PUNCT[op.term])
	}

	frame.Op = op
}

func (frame *Frame) AppendSpecAtomsFlags(spec Spec, atoms Atoms, flags uint) {
	frame.AppendOp(Op{spec, atoms, flags})
}

func (frame *Frame) AppendReduced(spec Spec, atoms Atoms) {
	frame.AppendOp(Op{spec, atoms, TERM_REDUCED})
}

func (frame *Frame) AppendRaw(spec Spec, atoms Atoms) {
	frame.AppendOp(Op{spec, atoms, TERM_RAW})
}

func (frame *Frame) AppendStateHeader(spec Spec) {
	frame.AppendSpecAtomsFlags(spec, NO_ATOMS, TERM_HEADER)
}

func (frame *Frame) AppendQueryHeader(spec Spec) {
	frame.AppendSpecAtomsFlags(spec, NO_ATOMS, TERM_QUERY)
}

func (frame *Frame) AppendSpecInt(spec Spec, i int64) {
	a := NewAtoms()
	a.AddInteger(i)
	frame.AppendSpecAtomsFlags(spec, a, TERM_REDUCED)
}

func (frame *Frame) AppendSpecUUID(spec Spec, uuid UUID) {
	a := NewAtoms()
	a.AddUUID(uuid)
	frame.AppendSpecAtomsFlags(spec, a, TERM_REDUCED)
}

//
//func (frame *Frame) AppendRange(i, j Frame) {
//	if !i.IsEmpty() && !j.IsEmpty() && i.offset >= j.offset {
//		return
//	}
//	if i.IsEmpty() {
//		return
//	}
//	if i.frame != j.frame {
//		panic("mismatching iterators")
//	}
//	// FIXME: last op exclusive!!!
//	frame.AppendOp(i.Op)
//	from := i.offset //+ len(i.Body)
//	till := j.offset
//	if till > from { // more than 1 op
//		frame.Body = append(frame.Body, i.frame.Body[from:till]...)
//	}
//}

func (frame *Frame) AppendAll(i Frame) {
	if i.IsEmpty() {
		return
	}
	for ; !i.IsEmpty(); i.Next() {
		frame.AppendOp(i.Op)
	}
}

func (frame *Frame) AppendRange(i, j Frame) {
	for ; !i.IsEmpty() && !i.IsSame(j.Spec); i.Next() {
		frame.AppendOp(i.Op)
	}
}

func (frame *Frame) AppendFrame(second Frame) {
	frame.AppendAll(second)
}

func (frame *Frame) Close() Frame {
	return ParseFrame(frame.state.data)
}

func MakeQueryFrame(headerSpec Spec) Frame {
	cur := MakeFrame(128)
	cur.AppendQueryHeader(headerSpec)
	return cur.Close()
}
