package ron

import (
	"math"
	"strconv"
)

const ATOM_INT_62 = uint64(ATOM_INT) << 62
const ATOM_FLOAT_62 = uint64(ATOM_FLOAT) << 62
const ATOM_STRING_62 = uint64(ATOM_STRING) << 62
const ATOM_UUID_62 = uint64(ATOM_UUID) << 62

func NewIntegerAtom(i int64) Atom {
	var a Atom
	a[ORIGIN] = ATOM_INT_62
	if i > 0 {
		a[VALUE] = uint64(i)
	} else {
		a[VALUE] = uint64(-i)
		a[ORIGIN] |= 1
	}
	return a
}

func (frame Frame) Count() int {
	return len(frame.atoms) - 4
}

func (a Atom) Type() uint {
	return uint(a[ORIGIN] >> 62)
}

func (a Atom) Integer() int64 {
	neg := a[ORIGIN] & (1 << 60)
	ret := int64(a[VALUE])
	if neg == 0 {
		return ret
	} else {
		return -ret
	}
}

func (a Atom) IsUUID() bool {
	return a.Type() == ATOM_UUID
}

func (a Atom) UUID() UUID {
	return UUID(a)
}

var BIT60 = uint64(1) << 60
var BIT61 = uint64(1) << 61

// We can't rely on standard floats cause they MUTATE THE VALUE.
// If 3.141592 is parsed then serialized, it becomes 3.141591(9)
// or something, that is entirely platform-dependent.
// Overall, floats are NOT commutative. Any floating arithmetic
// is highly discouraged inside CRDT type implementations.
func (a Atom) Float() float64 {
	return math.Float64frombits(a[VALUE])
}

func (a *Atom) setType(t uint64) {
	a[ORIGIN] = ((a[ORIGIN] << 4) >> 4) | t
}

func (a *Atom) setFloatType() {
	a.setType(ATOM_FLOAT_62)
}

func (a *Atom) setFrom(from int) {
	a[ORIGIN] |= uint64(from) << 30
}

func (a *Atom) setTill(till int) {
	a[ORIGIN] |= uint64(till)
}

func (a *Atom) parseValue(b []byte) {
	if a.Type() == ATOM_FLOAT {
		from := (a[ORIGIN] >> 30) & INT30_FULL
		till := a[ORIGIN] & INT30_FULL
		f, err := strconv.ParseFloat(string(b[from:till]), 64)
		if err == nil {
			a[VALUE] = math.Float64bits(f)
		}
	} else {
		// TODO: implement in nominal RON format
		panic("parsing is not implemented for this type")
	}
}

// add JSON escapes
func esc(str []byte) []byte {
	return str
}

// remove JSON escapes
func unesc(str []byte) []byte {
	// TODO
	return str
}

func (a Atom) RawString(body []byte) string {
	from := a[0] >> 32
	till := a[0] & INT32_FULL
	// FIXME check if binary
	return string(unesc(body[from:till]))
}

var INT32_FULL uint64 = (1 << 32) - 1
var INT30_FULL uint64 = (1 << 30) - 1

func (a Atom) EscString(body []byte) []byte {
	from := a[0] >> 32
	till := a[0] & INT32_FULL
	// FIXME check if binary
	return body[from:till]
}

func (frame Frame) RawString(idx int) string {
	atom := frame.atoms[idx+4]
	if atom.Type() != ATOM_STRING {
		return ""
	}
	return atom.RawString(frame.Body)
}

func (frame Frame) EscString(idx int) []byte {
	return frame.atoms[idx+4].EscString(frame.Body)
}

func (a *Atom) init64(half Half, flags uint8) {
	a[half] = uint64(flags) << 60
}

func (a *Atom) set1(half Half, idx uint) {
	a[half] |= uint64(1) << idx
}

func (a *Atom) set2(half Half, idx uint, value uint64) {
	a[half] |= value << (idx << 1)
}

func (a *Atom) set4(half Half, idx uint, value uint64) {
	a[half] |= value << (idx << 2)
}

func (a *Atom) reset4(half Half, idx uint, value uint8) {
	a[half] &^= 15 << (idx << 2)
	a[half] |= uint64(value) << (idx << 2)
}

func (a *Atom) set6(half Half, dgt int, value uint8) {
	a[half] |= uint64(value) << DIGIT_OFFSETS[dgt] // FIXME reverse numbering
}

func (a Atom) get6(half Half, dgt int) uint8 {
	return uint8((a[half] >> DIGIT_OFFSETS[dgt]) & 63)
}

func (a *Atom) trim6(half Half, dgt int) {
	a[half] &= INT60_FLAGS | PREFIX_MASKS[dgt]
}

func (a *Atom) set32(half Half, idx uint, value int) {
	a[half] |= uint64(value) << (idx << 5)
}

const INT16_FULL = (1 << 16) - 1

func (a *Atom) inc16(half Half, idx uint) {
	shift := uint(idx << 4)
	i := a[half] >> shift
	i++
	a[half] &^= INT16_FULL << shift
	a[half] |= (i & INT16_FULL) << shift
}

func (a *Atom) arab16(half Half, value byte) {
	i := a[half] & INT16_FULL
	i *= 10
	i += uint64(value)
	a[half] &^= INT16_FULL
	a[half] |= i & INT16_FULL
}

func (a *Atom) set64(half Half, value uint64) {
	a[half] = value
}

func (a *Atom) arab64(idx Half, value byte) {
	a[idx] *= 10
	a[idx] += uint64(value)
}
