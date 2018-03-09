package ron

type EmptyReducer struct {
}

func (r EmptyReducer) Reduce(inputs Batch) (frame Frame) {
	ai := &inputs[0]
	loc := ai.Ref()
	if !ai.IsHeader() {
		loc = ai.Event()
	}
	cur := MakeFrame(128)
	spec := ai.Spec()
	spec.SetEvent(inputs[len(inputs)-1].Event())
	spec.SetRef(loc)
	cur.AppendStateHeader(spec)
	return cur.Close()
}

type OmniReducer struct {
	empty EmptyReducer
	Types map[UUID]Reducer
}

var REDUCER = OmniReducer{}

func NewOmniReducer() (ret OmniReducer) {
	ret.Types = make(map[UUID]Reducer)
	for rdt, fct := range RDTYPES {
		ret.Types[rdt] = fct()
	}
	return
}

func (omni OmniReducer) Features () int {
	return 0 // should it be a reducer?
}

func (omni EmptyReducer) Features () int {
	return 0
}

func (omni OmniReducer) AddType(id UUID, r Reducer) {
	omni.Types[id] = r
}

func (omni OmniReducer) pickReducer(t UUID) Reducer {
	r := omni.Types[t]
	if r == nil {
		r = omni.empty
	}
	return r
}

// Reduce picks a reducer function, performs all the sanity checks,
// creates the header, invokes the reducer, returns the result
func (omni OmniReducer) Reduce(ins Batch) Frame {
	r := omni.pickReducer(ins[0].Type())
	// TODO sanity checks?
	return r.Reduce(ins)
}
