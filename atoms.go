package RON

import "math"

func NewAtoms() (ret Atoms) {
	ret.atoms = ret._atoms[0:0:len(ret._atoms)]
	return
}

const ATOM_INT_62 = uint64(ATOM_INT) << 62
const ATOM_FLOAT_62 = uint64(ATOM_FLOAT) << 62
const ATOM_STRING_62 = uint64(ATOM_STRING) << 62
const ATOM_UUID_62 = uint64(ATOM_UUID) << 62

func (a *Atoms) AddAtom(atom uint128) {
	a.atoms = append(a.atoms, atom)
}

func (a *Atoms) AddInteger(i int64) {
	second := ATOM_INT_62
	var first uint64
	if i >= 0 {
		first = uint64(i)
	} else {
		first = uint64(-i)
		second |= 1
	}
	a.AddAtom(uint128{first, second})
}

func (a *Atoms) AddStringRange(from, till uint64) {
	a.AddAtom(uint128{from, till | ATOM_STRING_62})
}

func (a *Atoms) AddString(s string) {
	// TODO
}

func (a *Atoms) AddFloat(i float64) {
	a.AddAtom(uint128{uint64(i), ATOM_FLOAT_62}) // TODO
}

func (a *Atoms) AddUUID(uuid UUID) {
	a.AddAtom(uint128{uuid.uint128[0], uuid.uint128[1] | ATOM_UUID_62})
}

func (a *Atoms) Reset() {
	a.atoms = a.atoms[:0]
}

func (a Atoms) Count() int {
	return len(a.atoms)
}

func (a *Atoms) AType(i int) uint {
	return uint(a.atoms[i][1] >> 62)
}

func (a *Atoms) Integer(i int) int64 {
	neg := a.atoms[i][1] & 1
	ret := int64(a.atoms[i][0])
	if neg == 0 {
		return ret
	} else {
		return -ret
	}
}

func (a *Atoms) Float(i int) float64 {
	num := a.atoms[i][0]
	exp := a.atoms[i][1]
	return math.Pow10(int(exp)) * float64(num)
}

func (a *Atoms) String(i int) string {
	from := a.atoms[i][0] & INT60_FULL
	till := a.atoms[i][1] & INT60_FULL
	return string(a.frame[from:till])
}
