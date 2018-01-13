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
	if uuid.IsTranscendentName() {
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

func sharedPrefix(uuid, context UUID) (ret int) {
	vp := bits.LeadingZeros64(uuid.Value() ^ context.Value())
	vp -= vp % 6
	op := bits.LeadingZeros64((uuid.Origin() ^ context.Origin()) & INT60_FULL)
	op -= op % 6
	ret = vp + op
	if uuid.Scheme() != context.Scheme() {
		ret--
	}
	return
}

func (frame *Frame) appendUUID(uuid UUID, context UUID) {
	if 0 != frame.Serializer.Format&FORMAT_UNZIP {
		frame.Body = FormatUUID(frame.Body, uuid)
	} else if uuid != context {
		frame.Body = FormatZipUUID(frame.Body, uuid, context)
	}
}

var zeros []byte = []byte("0000000000")

func (frame *Frame) appendFloat(a Atom) {
	if a[0] == 0 {
		frame.Body = append(frame.Body, "0.0"...)
		return
	}
	intStr := fmt.Sprintf("%d", a[0])
	e := int(a[1] & INT32_FULL)
	if a[1]&BIT32 != 0 {
		frame.Body = append(frame.Body, '-')
	}
	if a[1]&BIT33 != 0 { // neg e
		ip := len(intStr) - e
		if ip > 0 { // integer part
			frame.Body = append(frame.Body, intStr[:ip]...)
			frame.Body = append(frame.Body, '.')
			tail := intStr[ip:]
			for len(tail) > 1 && tail[len(tail)-1] == '0' {
				tail = tail[:len(tail)-1]
			}
			frame.Body = append(frame.Body, tail...)
		} else if ip == 0 {
			frame.Body = append(frame.Body, '0', '.')
			frame.Body = append(frame.Body, intStr...)
		} else {
			de := 1 - ip
			frame.Body = append(frame.Body, intStr[:1]...)
			frame.Body = append(frame.Body, '.')
			if len(intStr) > 1 {
				frame.Body = append(frame.Body, intStr[1:]...)
			} else {
				frame.Body = append(frame.Body, '0')
			}
			frame.Body = append(frame.Body, 'e', '-')
			exp := fmt.Sprintf("%d", de)
			frame.Body = append(frame.Body, exp...)
		}
	} else {
		if e+len(intStr) <= 10 {
			frame.Body = append(frame.Body, intStr...)
			frame.Body = append(frame.Body, zeros[:e]...)
			frame.Body = append(frame.Body, ".0"...)
		} else {
			exp := fmt.Sprintf("%d", e+len(intStr)-1)
			frame.Body = append(frame.Body, intStr[0])
			frame.Body = append(frame.Body, '.')
			frame.Body = append(frame.Body, intStr[1:]...)
			frame.Body = append(frame.Body, "e+"...)
			frame.Body = append(frame.Body, exp...)
		}
	}
}

func (frame *Frame) appendSpec(other Frame) {

	start := len(frame.Body)
	flags := frame.Serializer.Format
	skips := 0
	spec := other.atoms[:4]
	context := frame.atoms[:4]

	do_grid := flags&FORMAT_GRID != 0
	do_space := flags&FORMAT_SPACE != 0
	do_noskip := flags&FORMAT_NOSKIP != 0
	do_redef := flags&FORMAT_REDEFAULT != 0

	k := 4
	if spec[SPEC_TYPE] == Atom(COMMENT_UUID) {
		k = 1
	}

	for t := 0; t < k; t++ {
		if do_grid {
			rest := t*22 - (len(frame.Body) - start)
			frame.Body = append(frame.Body, SPACES88[:rest]...)
		} else if do_space && t > 0 {
			frame.Body = append(frame.Body, ' ')
		}
		if !do_noskip && spec[t] == context[t] && (other.term == TERM_REDUCED || spec[t] == ZERO_UUID_ATOM) {
			skips++
			continue
		}
		frame.Body = append(frame.Body, SPEC_PUNCT[uint(t)])
		if t > 0 && do_redef {
			ctxAt := 0
			ctxUUID := UUID(spec[t-1])
			ctxPL := sharedPrefix(UUID(spec[t]), ctxUUID)
			for i := 1; i < 4; i++ {
				pl := sharedPrefix(UUID(spec[t]), UUID(context[i]))
				if pl > ctxPL {
					ctxPL = pl
					ctxUUID = UUID(context[i])
					ctxAt = i
				}
			}
			if ctxAt != t {
				frame.Body = append(frame.Body, REDEF_PUNCT[uint(ctxAt)])
			}
			frame.appendUUID(UUID(spec[t]), ctxUUID)
		} else {
			frame.appendUUID(UUID(spec[t]), UUID(context[t]))
		}
	}
	if skips == 4 {
		frame.Body = append(frame.Body, '@')
	}
}

func (frame *Frame) appendAtoms(other Frame) {
	for i := 4; i < len(other.atoms); i++ {
		a := other.atoms[i]
		switch a.Type() {
		case ATOM_INT:
			{
				frame.Body = append(frame.Body, ATOM_INT_SEP)
				s := fmt.Sprint(a.Integer())
				frame.Body = append(frame.Body, []byte(s)...)
			}
		case ATOM_STRING:
			{
				frame.Body = append(frame.Body, ATOM_STRING_SEP)
				frame.Body = append(frame.Body, other.EscString(i-4)...)
				frame.Body = append(frame.Body, ATOM_STRING_SEP)
			}
		case ATOM_FLOAT:
			{
				frame.Body = append(frame.Body, ATOM_FLOAT_SEP)
				frame.appendFloat(a)
			}
		case ATOM_UUID:
			{
				frame.Body = append(frame.Body, ATOM_UUID_SEP)
				frame.appendUUID(a.UUID(), ZERO_UUID) // TODO defaults
			}
		}
	}
}

func (frame *Frame) Append(other Frame) {

	flags := frame.Serializer.Format
	start := len(frame.Body)
	if len(frame.Body) > 0 && (0 != flags&FORMAT_OP_LINES || (0 != flags&FORMAT_FRAME_LINES && !other.IsFramed())) {
		frame.Body = append(frame.Body, '\n')
		if 0 != flags&FORMAT_INDENT && !other.IsHeader() {
			frame.Body = append(frame.Body, "    "...)
		}
	} else if 0 != flags&FORMAT_HEADER_SPACE && frame.IsHeader() {
		frame.Body = append(frame.Body, ' ')
	}

	if len(frame.atoms) == 0 {
		frame.atoms = frame._atoms[:4]
	}
	frame.appendSpec(other)

	if 0 != flags&FORMAT_GRID {
		rest := 4*22 - (len(frame.Body) - start)
		frame.Body = append(frame.Body, SPACES88[:rest]...)
	}

	frame.appendAtoms(other)

	defaultTerm := TERM_REDUCED
	if frame.term == TERM_RAW {
		defaultTerm = TERM_RAW
	}

	if other.term != defaultTerm || other.Count() == 0 {
		frame.Body = append(frame.Body, TERM_PUNCT[other.term])
	}

	if len(other.atoms) > len(frame.atoms) {
		copy(frame.atoms, other.atoms[:len(frame.atoms)])
		frame.atoms = append(frame.atoms, other.atoms[len(frame.atoms)])
	} else {
		copy(frame.atoms, other.atoms)
		frame.atoms = frame.atoms[:len(other.atoms)]
	}
	frame.term = other.term
	frame.Position++

}

func (frame *Frame) AppendReducedRef(ref UUID, other Frame) {
	tmpRef, tmpTerm := other.atoms[SPEC_REF], other.term
	other.atoms[SPEC_REF] = Atom(ref)
	other.term = TERM_REDUCED
	frame.Append(other)
	other.atoms[SPEC_REF], other.term = tmpRef, tmpTerm
}

func (frame *Frame) AppendReduced(other Frame) {
	tmpTerm := other.term
	other.term = TERM_REDUCED
	frame.Append(other)
	other.term = tmpTerm
}

func (frame *Frame) AppendEmpty(spec Spec, term int) {
	atoms := make([]Atom, 0, 6)
	atoms = append(atoms, spec[0:4]...)
	tmp := Frame{atoms: atoms, term: term}
	frame.Append(tmp)
}

func (frame *Frame) AppendEmptyReducedOp(spec Spec) {
	frame.AppendEmpty(spec, TERM_REDUCED)
}

func NewOpFrame(t, o, e, r UUID, term int) Frame {
	frame := NewFrame()
	frame.AppendEmpty(NewSpec(t, o, e, r), term)
	return frame.Rewind()
}

func NewRawOp(t, o, e, r UUID) Frame {
	return NewOpFrame(t, o, e, r, TERM_RAW)
}

func NewStateHeader(t, o, e, r UUID) Frame {
	return NewOpFrame(t, o, e, r, TERM_HEADER)
}

func (frame *Frame) AppendAmended(spec Spec, values Frame, term int) {
	tmp := values.Clone()
	tmp.term = term
	copy(tmp.atoms, spec)
	frame.Append(tmp)
}

func (frame *Frame) AppendSpecValT(spec Spec, value Atom, term int) {
	var atoms [5]Atom
	copy(atoms[:], spec)
	atoms[4] = value
	tmp := Frame{atoms: atoms[:], term: term}
	frame.Append(tmp)
}

func (frame *Frame) AppendReducedOpInt(spec Spec, value int64) {
	frame.AppendSpecValT(spec, NewIntegerAtom(value), TERM_REDUCED)
}

func (frame *Frame) AppendReducedOpUUID(spec Spec, value UUID) {
	frame.AppendSpecValT(spec, NewUUIDAtom(value), TERM_REDUCED)
}

func (frame *Frame) AppendStateHeader(spec Spec) {
	frame.AppendEmpty(spec, TERM_HEADER)
}

func (frame *Frame) AppendQueryHeader(spec Spec) {
	frame.AppendEmpty(spec, TERM_QUERY)
}

//func (frame *Frame) AppendRaw (spec Spec, value Atom) {
//	frame.AppendSpecValT(spec, value, TERM_RAW)
//}

func (frame Frame) Atoms() []Atom {
	return frame.atoms[:]
}

func (frame *Frame) AppendAll(i Frame) {
	if i.IsEmpty() {
		return
	}
	for !i.EOF() {
		frame.Append(i)
		i.Next()
	}
}

func (frame *Frame) AppendFrame(second Frame) {
	frame.AppendAll(second)
}

func (frame *Frame) Close() Frame {
	return ParseFrame(frame.Body)
}

func (frame *Frame) AddReducedOp(event, ref UUID, atoms ...Atom) {
	tmp := Frame{term: TERM_REDUCED}
	a := tmp._atoms[:0]
	a = append(a, frame.atoms[:2]...)
	a = append(a, Atom(event), Atom(ref))
	a = append(a, atoms...)
	tmp.atoms = a
	frame.Append(tmp)
}

func MakeQueryFrame(headerSpec Spec) Frame {
	cur := MakeFrame(128)
	cur.AppendQueryHeader(headerSpec)
	return cur.Close()
}

// overall, serialized batches are used in rare cases
// (delivery fails, cross-key transactions)
// hence, we don't care about performance that much
// still, may consider explicit-length formats at some point
func (frame Frame) Split() Batch {
	ret := Batch{}
	for !frame.EOF() {
		next := NewFrame()
		next.Append(frame)
		frame.Next()
		for !frame.EOF() && frame.Term() == TERM_REDUCED {
			next.Append(frame)
			frame.Next()
		}
		ret = append(ret, next.Rewind())
	}
	return ret
}

func (batch Batch) Join() Frame {
	size := 0
	for _, f := range batch {
		size += len(f.Body)
	}
	ret := MakeFrame(size)
	for _, f := range batch {
		ret.AppendFrame(f)
	}

	return ret.Rewind()
}
