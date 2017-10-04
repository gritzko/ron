package RON

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
		pchar := prefixBits2Sep(uint(prefix)/6 - 4)
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
		if len(buf) > at && ABC[buf[at]] < 0 {
			copy(buf[at-1:], buf[at:])
			buf = buf[:len(buf)-1]
		} else if len(buf) == at && ABC[buf[start]] < 0 {
			buf = buf[:len(buf)-1]
		}
	}
	return buf
}

// TODO FormatInt Float String

func (cur *Cursor) appendUUID(uuid UUID, context UUID) {
	if 0 != cur.Format&FORMAT_UNZIP {
		cur.body = FormatUUID(cur.body, uuid)
	} else if uuid != context {
		cur.body = FormatZipUUID(cur.body, uuid, context)
	}
}

func (cur *Cursor) appendSpec(spec, context Spec) {
	buf := cur.body
	start := len(buf)
	flags := cur.Format
	for t := 0; t < 4; t++ {
		if 0 != flags&FORMAT_GRID {
			rest := t*22 - (len(buf) - start)
			buf = append(buf, SPACES88[:rest]...)
		} else if 0 != flags&FORMAT_SPACE && t > 0 {
			buf = append(buf, ' ')
		}
		if (spec.uuids[t] == context.uuids[t]) && (0 == flags&FORMAT_NOSKIP) {
			continue
		}
		buf = append(buf, specBits2Sep(uint(t)))
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
				buf = append(buf, redefBits2Sep(uint(ctxAt)))
			}
			cur.appendUUID(spec.uuids[t], ctxUUID)
		} else {
			cur.appendUUID(spec.uuids[t], context.uuids[t])
		}
	}
	cur.body = buf
}

func (cur *Cursor) appendAtoms(a Atoms) {
	for i := 0; i < a.Count(); i++ {
		switch a.AType(i) {
		case ATOM_INT:
			{
				cur.body = append(cur.body, ATOM_INT_SEP)
				s := fmt.Sprint(a.atoms[i][0])
				cur.body = append(cur.body, []byte(s)...)
			}
		case ATOM_STRING:
			{
				cur.body = append(cur.body, ATOM_STRING_SEP)
				ft := a.atoms[i]
				cur.body = append(cur.body, a.frame[ft[0]&INT60_FULL:ft[1]&INT60_FULL]...)
				cur.body = append(cur.body, ATOM_STRING_SEP)
			}
		case ATOM_FLOAT:
			{
				cur.body = append(cur.body, ATOM_FLOAT_SEP)
				s := fmt.Sprint(a.atoms[i][0])
				cur.body = append(cur.body, []byte(s)...)
				cur.body = append(cur.body, '.')
			}
		case ATOM_UUID:
			{
				cur.body = append(cur.body, ATOM_UUID_SEP)
				cur.appendUUID(UUID{a.atoms[i]}, ZERO_UUID)
			}
		}
	}
}

func (cur *Cursor) AppendOp(op Op) {

	if cur.seq == 0 {
		cur.first = op
	}
	flags := cur.Format
	start := len(cur.body)
	if len(cur.body) > 0 && (0 != flags&FORMAT_OP_LINES || (0 != flags&FORMAT_FRAME_LINES && !op.IsFramed())) {
		cur.body = append(cur.body, '\n')
		if 0 != flags&FORMAT_INDENT && !op.IsHeader() {
			cur.body = append(cur.body, "    "...)
		}
	} else if 0 != flags&FORMAT_HEADER_SPACE && cur.last.IsHeader() {
		cur.body = append(cur.body, ' ')
	}

	cur.appendSpec(op.Spec, cur.last.Spec)

	if 0 != flags&FORMAT_GRID {
		rest := 4*22 - (len(cur.body) - start)
		cur.body = append(cur.body, SPACES88[:rest]...)
	}

	cur.appendAtoms(op.Atoms)

	if op.IsHeader() || (op.IsRaw() && !cur.last.IsRaw()) || op.Atoms.Count() == 0 {
		cur.body = append(cur.body, termBits2Sep(op.term))
	}

	cur.seq++
	cur.prep = false
	cur.last = op
}

func (cur *Cursor) AppendSpecAtomsFlags(spec Spec, atoms Atoms, flags uint) {
	cur.AppendOp(Op{spec, atoms, flags})
}

func (cur *Cursor) AppendReduced(spec Spec, atoms Atoms) {
	cur.AppendOp(Op{spec, atoms, TERM_REDUCED})
}

func (cur *Cursor) AppendRaw(spec Spec, atoms Atoms) {
	cur.AppendOp(Op{spec, atoms, TERM_RAW})
}

func (cur *Cursor) AppendStateHeader(spec Spec) {
	cur.AppendSpecAtomsFlags(spec, NO_ATOMS, TERM_HEADER)
}

func (cur *Cursor) AppendQueryHeader(spec Spec) {
	cur.AppendSpecAtomsFlags(spec, NO_ATOMS, TERM_QUERY)
}

func (cur *Cursor) AppendSpecInt(spec Spec, i int64) {
	a := NewAtoms()
	a.AddInteger(i)
	cur.AppendSpecAtomsFlags(spec, a, TERM_REDUCED)
}

func (cur *Cursor) AppendSpecUUID(spec Spec, uuid UUID) {
	a := NewAtoms()
	a.AddUUID(uuid)
	cur.AppendSpecAtomsFlags(spec, a, TERM_REDUCED)
}

//
//func (frame *Cursor) AppendRange(i, j Iterator) {
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

func (cur *Cursor) Append() {
	if !cur.prep {
		panic("no prepping")
	}
	cur.AppendOp(cur.new)
	cur.new.Reset()
}

func (cur *Cursor) AppendAll(i Iterator) {
	if i.IsEmpty() {
		return
	}
	for ; !i.IsEmpty(); i.Next() {
		cur.AppendOp(i.Op)
	}
}

func (cur *Cursor) AppendRange(i, j Iterator) {
	for ; !i.IsEmpty() && !i.IsSame(j.Spec); i.Next() {
		cur.AppendOp(i.Op)
	}
}

func (cur *Cursor) AppendFrame(second Frame) {
	cur.AppendAll(second.Begin())
}

func (cur *Cursor) Close() Frame {
	return Frame{Op: cur.first, body: cur.body}
}

func MakeQueryFrame(headerSpec Spec) Frame {
	cur := MakeFrame(128)
	cur.AppendQueryHeader(headerSpec)
	return cur.Close()
}

func (cur *Cursor) Prepare() {
	if !cur.prep {
		cur.prep = true
		cur.new = cur.last
	}
}

func (cur *Cursor) AddInteger(i int64) {
	cur.Prepare()
	cur.new.AddInteger(i)
}

func (cur *Cursor) AddString(s string) {
	cur.Prepare()
	// TODO JSON escape
	cur.new.AddString(s)
}

func MakeFormattedFrame(format uint, prealloc_bytes int) Cursor {
	return Cursor{body: make([]byte, format, prealloc_bytes)}
}

func MakeFrame(prealloc_bytes int) Cursor {
	return Cursor{body: make([]byte, 0, prealloc_bytes)}
}

func (cur Cursor) String() string {
	return string(cur.body)
}
