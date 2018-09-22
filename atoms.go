package ron

import (
	"math"
)

const ATOM_INT_62 = uint64(ATOM_INT) << 62
const ATOM_FLOAT_62 = uint64(ATOM_FLOAT) << 62
const ATOM_STRING_62 = uint64(ATOM_STRING) << 62
const ATOM_UUID_62 = uint64(ATOM_UUID) << 62

func NewIntegerAtom(i int64) (a Atom) {
	a[1] = ATOM_INT_62
	if i > 0 {
		a[0] = uint64(i)
	} else {
		a[0] = uint64(-i)
		a[1] |= 1
	}
	return
}

func NewIntAtom(i int) Atom {
	return NewIntegerAtom(int64(i))
}

func NewStringRangeAtom(from, till int) Atom {
	return Atom{uint64((from << 32) | till), ATOM_STRING_62}
}

func NewFloatAtom(f float64) Atom {
	return Atom{math.Float64bits(f), ATOM_FLOAT_62}
}

func NewUUIDAtom(uuid UUID) Atom {
	return Atom(uuid)
}

func (frame Frame) Count() int {
	return len(frame.atoms) - 4
}

func (a Atom) Type() uint {
	return uint(a[1] >> 62)
}

func (a Atom) Integer() int64 {
	neg := a[1] & (1 << 60)
	ret := int64(a[0])
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
// Hence, we work that around by storing a 64-bit integer 3141592 and
// a 32-bit exponent.
// Overall, floats are NOT commutative. Any floating arithmetic
// is highly discouraged inside CRDT type implementations.
func (a Atom) Float() float64 {
	pow := a.pow()
	ret := float64(a[VALUE]) * math.Pow10(pow)
	if a[ORIGIN]&BIT60 != 0 {
		ret = -ret
	}
	return ret
}

func (a Atom) pow() int {
	pow := int(a[ORIGIN] & INT16_FULL)
	if a[ORIGIN]&BIT61 != 0 {
		pow = -pow
	}
	pow -= int((a[ORIGIN] >> 16) & INT16_FULL)
	return pow
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
