package RON

type EmptyReducer struct {
}

func (r EmptyReducer) Reduce(a, b Frame) (result Frame, err UUID) {
	ai, bi := a.Begin(), b.Begin()
	loc := ai.Location()
	if !ai.IsHeader() {
		loc = ai.Event()
	}
	result.AppendSpecAtoms(Spec{ai.Type(), ai.Object(), bi.Event(), loc}, STATE_HEADER_ATOMS)
	return
}

func (r EmptyReducer) ReduceAll(inputs []Frame) (result Frame, err UUID) {
	return r.Reduce(inputs[0], inputs[len(inputs)-1])
}

type OmniReducer struct {
	Types map[uint64]Reducer
}

var REDUCER = OmniReducer{}

// Reduce picks a reducer function, performs all the sanity checks,
// creates the header, invokes the reducer, returns the result
func (omni OmniReducer) Reduce(a, b Frame) (result Frame, err UUID) {
	var length int = len(a.Body) + len(b.Body)
	ret := Frame{Body: make([]byte, 0, length)} //FIXME last
	i, j := a.Begin(), b.Begin()
	error_uuid := ZERO_UUID
	var fn Reducer = omni.Types[i.Type().Value]
	if fn == nil {
		error_uuid = UNKNOWN_TYPE_ERROR_UUID
	} else if i.Type() != j.Type() {
		error_uuid = TYPE_MISMATCH_ERROR_UUID
	}
	// plant header
	if error_uuid == ZERO_UUID {
		ret, error_uuid = fn.Reduce(a, b)
	}
	if error_uuid != ZERO_UUID {
		ret = Frame{Body: make([]byte, 100)}
		ret.AppendSpecAtoms(Spec{i.Type(), i.Object(), ERROR_UUID, error_uuid}, STATE_HEADER_ATOMS)
	}
	return
}

func (omni OmniReducer) ReduceAll(inputs []Frame) (result Frame, err UUID) {
	//ret = frames[0]
	//for i:=1; i<len(frames); i++ {
	//	ret = Reduce(ret, frames[i])
	//}
	return
}
