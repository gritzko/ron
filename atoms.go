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
	neg := a[1] & 1
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

var BIT32 = uint64(1)<<32
var BIT33 = uint64(1)<<33

// We can't rely on standard floats cause they MUTATE THE VALUE.
// If 3.141592 is parsed then serialized, it becomes 3.141591(9)
// or something, that is entirely platform-dependent.
// Hence, we work that around by storing a 64-bit integer 3141592 and
// a 32-bit exponent.
// Overall, floats are NOT commutative. Any floating arithmetic
// is highly discouraged inside CRDT type implementations.
func (a Atom) Float() float64 {
	pow := int(a[1]&INT32_FULL)
	if a[1]&BIT33 != 0 {
		pow = -pow
	}
	ret := float64(a[0]) * math.Pow10(pow)
	if a[1]&BIT32 != 0 {
		ret = -ret
	}
	return ret
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
	if atom.Type()!=ATOM_STRING {
		return ""
	}
	return atom.RawString(frame.Body)
}

func (frame Frame) EscString(idx int) []byte {
	return frame.atoms[idx+4].EscString(frame.Body)
}
