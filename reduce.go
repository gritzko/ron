package RON

type EmptyReducer struct {
}

func (r EmptyReducer) Reduce(a, b Frame) (frame Frame, err UUID) {
	ai, bi := a.Begin(), b.Begin()
	loc := ai.Ref()
	if !ai.IsHeader() {
		loc = ai.Event()
	}
	cur := MakeFrame(128)
	spec := ai.Spec
	spec.uuids[SPEC_EVENT] = bi.Event()
	spec.uuids[SPEC_REF] = loc
	cur.AppendStateHeader(spec)
	return cur.Close(), ZERO_UUID
}

func (r EmptyReducer) ReduceAll(inputs []Frame) (result Frame, err UUID) {
	return r.Reduce(inputs[0], inputs[len(inputs)-1])
}

type OmniReducer struct {
	empty EmptyReducer
	Types map[uint64]Reducer
}

var REDUCER = OmniReducer{}

func NewOmniReducer() (ret OmniReducer) {
	ret.Types = make(map[uint64]Reducer)
	return
}

func (omni OmniReducer) AddType(id UUID, r Reducer) {
	omni.Types[id.Value()] = r
}

func (omni OmniReducer) pickReducer(t UUID) Reducer {
	r := omni.Types[t.Value()]
	if r == nil {
		r = omni.empty
	}
	return r
}

// Reduce picks a reducer function, performs all the sanity checks,
// creates the header, invokes the reducer, returns the result
func (omni OmniReducer) Reduce(a, b Frame) (result Frame, err UUID) {
	r := omni.pickReducer(a.Begin().Spec.Type())
	return r.Reduce(a, b)
}

func (omni OmniReducer) ReduceAll(inputs []Frame) (result Frame, err UUID) {
	r := omni.pickReducer(inputs[0].Begin().Spec.Type())
	// TODO sanity checks
	return r.ReduceAll(inputs)
}
