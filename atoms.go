package ron

import (
	"math"
	"strconv"
)

const (
	int_atom_flags    = uint64(4|ATOM_INT) << 60
	float_atom_flags  = uint64(4|ATOM_FLOAT) << 60
	string_atom_flags = uint64(4|ATOM_STRING) << 60
	int30Full         = uint64(1<<30) - 1
)

func NewIntegerAtom(i int64) Atom {
	a := Atom{0, int_atom_flags}
	if i < 0 {
		a[VALUE] = uint64(-i) | 1<<63
	} else {
		a[VALUE] = uint64(i)
	}
	return a
}

func NewFloatAtom(f float64) Atom {
	return Atom{math.Float64bits(f), float_atom_flags}
}

func (frame Frame) Count() int {
	// TODO: move to frame implementation file
	return len(frame.atoms) - 4
}

func (a Atom) Type() uint {
	// Note:
	//   00xx - UUID type
	//   01xx - scalar types
	if a[ORIGIN]>>62 == 0 {
		return ATOM_UUID
	} else {
		return uint((a[ORIGIN] << 2) >> 62)
	}
}

func (a Atom) Integer() int64 {
	if a[VALUE]&(1<<63) != 0 {
		return -int64(a[VALUE])
	} else {
		return int64(a[VALUE])
	}
}

func (a Atom) IsUUID() bool {
	return a.Type() == ATOM_UUID
}

func (a Atom) UUID() UUID {
	return UUID(a)
}

// We can't rely on standard floats cause they MUTATE THE VALUE.
// If 3.141592 is parsed then serialized, it becomes 3.141591(9)
// or something, that is entirely platform-dependent.
// Overall, floats are NOT commutative. Any floating arithmetic
// is highly discouraged inside CRDT type implementations.
func (a Atom) Float() float64 {
	return math.Float64frombits(a[VALUE])
}

func (a *Atom) setType(t uint64) {
	// setType resets first 4 bits, so it
	// assumed only int|string|float types
	a[ORIGIN] = ((a[ORIGIN] << 4) >> 4) | ((4 | t) << 60)
}

func (a *Atom) setFrom(from int) {
	a[ORIGIN] |= uint64(from) << 30
}

func (a *Atom) setTill(till int) {
	a[ORIGIN] |= uint64(till)
}

func (a *Atom) parseValue(b []byte) {
	// TODO: handle parsing error
	switch t := a.Type(); true {
	case t == ATOM_FLOAT:
		f, err := strconv.ParseFloat(a.getSource(b), 64)
		if err == nil {
			a[VALUE] = math.Float64bits(f)
		}
	case t == ATOM_INT:
		i, err := strconv.ParseInt(a.getSource(b), 10, 64)
		if err == nil {
			if i < 0 {
				a[VALUE] = uint64(-i) | 1<<63
			} else {
				a[VALUE] = uint64(i)
			}
		}
	case t == ATOM_STRING:
		// TODO: save short strings in a VALUE slot for advanced optimizations
		break
	default:
		panic("parsing is not implemented for this type")
	}
}

func (a Atom) getFrom() uint64 {
	return (a[ORIGIN] >> 30) & int30Full
}

func (a Atom) getTill() uint64 {
	return a[ORIGIN] & int30Full
}

func (a Atom) getSource(b []byte) string {
	return string(b[a.getFrom():a.getTill()])
}

// remove JSON escapes
func unesc(str []byte) []byte {
	// TODO
	return str
}

func (a Atom) RawString(b []byte) string {
	// FIXME check if binary
	return string(unesc(b[a.getFrom():a.getTill()]))
}

func (a Atom) EscString(b []byte) []byte {
	// FIXME check if binary
	return b[a.getFrom():a.getTill()]
}

func (frame Frame) RawString(idx int) string {
	// TODO: move to frame implementation file
	atom := frame.atoms[idx+4]
	if atom.Type() != ATOM_STRING {
		return ""
	}
	return atom.RawString(frame.Body)
}

func (frame Frame) EscString(idx int) []byte {
	return frame.atoms[idx+4].EscString(frame.Body)
}
