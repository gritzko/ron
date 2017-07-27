package RON

var Reducers map[UUID]Reducer

// Reduce picks a reducer function, performs all the sanity checks,
// creates the header, invokes the reducer, returns the result
func Reduce (a Iterator, b Iterator) Frame {
	// parse key
	// pick fn
	return Frame{}
}


var LWW_UUID, _ = ParseUUIDString("lww")

func ReduceLWW (a Iterator, b Iterator) Frame {
	if a.IsHeader() {
		a.Next()
	}
	if b.IsHeader() {
		b.Next()
	}
	var ret Frame
	for !a.AtEnd() && !b.AtEnd() {
		loc_cmp := Compare(a.Location(), b.Location())
		if loc_cmp == 0 {
			ev_cmp := Compare(a.Event(), b.Event())
			if ev_cmp < 0 {
				ret.AppendOp(&b.Op)
			} else {
				ret.AppendOp(&a.Op)
			}
			a.Next()
			b.Next()
		} else if loc_cmp < 0 {
			ret.AppendOp(&a.Op)
			a.Next()
		} else {
			ret.AppendOp(&b.Op)
			b.Next()
		}
	}
	if !a.AtEnd() {
		ret.AppendAll(a)
	}
	if !b.AtEnd() {
		ret.AppendAll(b)
	}
	return ret
}
