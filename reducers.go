package RON

var HEADER_ATOMS []byte = []byte("!")
var ERROR_ATOMS []byte = []byte("!!!")

// Reduce picks a reducer function, performs all the sanity checks,
// creates the header, invokes the reducer, returns the result
func Reduce (a, b Frame) Frame {
	var length int = len(a.Body) + len(b.Body)
	var ret Frame = Frame{Body:make([]byte, 0, length), last:ZERO_OP} //FIXME last
	// FIXME: context A$B value A-B must produce "-"
	i, j := a.Begin(), b.Begin()
	var error_uuid UUID = ZERO_UUID
	var fn Reducer = Reducers[i.Type()]
	if fn==nil {
		error_uuid = UNKNOWN_TYPE_ERROR_UUID
	} else if i.Type()!=j.Type() {
		error_uuid = TYPE_MISMATCH_ERROR_UUID
	}
	// plant header
	if error_uuid==ZERO_UUID {
		var loc UUID = i.Location()
		if !i.IsHeader() {
			loc = i.Event()
		}
		ret.Append(i.Type(), i.Object(), j.Event(), loc, HEADER_ATOMS)
		error_uuid = fn(i, j, &ret)
	}
	if error_uuid!=ZERO_UUID {
		ret = Frame{Body:make([]byte,100)}
		ret.Append(i.Type(), i.Object(), ERROR_UUID, error_uuid, ERROR_ATOMS)
	}
	return ret
}

func ReduceAll (frames []Frame) (ret Frame) {
	ret = frames[0]
	for i:=1; i<len(frames); i++ {
		ret = Reduce(ret, frames[i])
	}
	return
}


var LWW_UUID = UUID{881557636825219072, '$', 0}

func ReduceLWW (a Iterator, b Iterator, ret *Frame) UUID {
	if a.IsHeader() {
		a.Next()
	}
	if b.IsHeader() {
		b.Next()
	}
	for !a.AtEnd() && !b.AtEnd() {
		loc_cmp := Compare(a.Location(), b.Location())
		if loc_cmp == 0 {
			ev_cmp := Compare(a.Event(), b.Event())
			if ev_cmp > 0 {
				ret.AppendOp(&b.Op)
			} else {
				ret.AppendOp(&a.Op)
			}
			a.Next()
			b.Next()
		} else if loc_cmp > 0 {
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
	return ZERO_UUID
}

var Reducers = map[UUID]Reducer{}

func init () {
	Reducers[LWW_UUID] = ReduceLWW
}
