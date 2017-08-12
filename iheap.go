package RON

//import "github.com/gritzko/RON"

// IHeap is an iterator heap - gives the minimum available element
// at every step. Useful for merge sort like algorithms.
type IHeap struct {
	// Most of the time, a heap has 2 elements, optimize for that.
	// Sometimes, it can get millions of elements, ensure that is O(NlogN)
	iters               []*Iterator
	primary, secondary  int
	prim_desc, sec_desc bool
}

// sort modes, e.g. PRIM_EVENT|PRIM_DESC|SEC_LOCATION
const (
	PRIM_DESC     = 4
	PRIM_TYPE     = SPEC_TYPE
	PRIM_OBJECT   = SPEC_OBJECT
	PRIM_EVENT    = SPEC_EVENT
	PRIM_LOCATION = SPEC_LOCATION
	SEC_DESC      = 32
	SEC_TYPE      = SPEC_TYPE << 3
	SEC_OBJECT    = SPEC_OBJECT << 3
	SEC_EVENT     = SPEC_EVENT << 3
	SEC_LOCATION  = SPEC_LOCATION << 3
)

func MakeIHeap(mode, size int) (ret IHeap) {
	ret.iters = make([]*Iterator, 1, size+1)
	ret.prim_desc = (mode & PRIM_DESC) != 0
	ret.sec_desc = (mode & SEC_DESC) != 0
	ret.primary = mode & 3
	ret.secondary = (mode >> 3) & 3
	return
}

func (h IHeap) less(i, j int) bool {
	c := Compare(h.iters[i].Spec[h.primary], h.iters[j].Spec[h.primary])
	if c == 0 {
		c = Compare(h.iters[i].Spec[h.secondary], h.iters[j].Spec[h.secondary])
		if h.sec_desc {
			c = -c
		}
	} else if h.prim_desc {
		c = -c
	}
	//fmt.Printf("CMP %s %s %d\n", h.iters[i].String(), h.iters[j].String(), c)
	return c < 0
}

func (h *IHeap) sink(i int) {
	to := i
	j := i << 1
	if j < len(h.iters) && h.less(j, i) {
		to = j
	}
	j++
	if j < len(h.iters) && h.less(j, to) {
		to = j
	}
	if to != i {
		h.swap(i, to)
		h.sink(to)
	}
}

func (h *IHeap) raise(i int) {
	j := i >> 1
	if j > 0 && h.less(i, j) {
		h.swap(i, j)
		if j > 1 {
			h.raise(j)
		}
	}
}

func (h IHeap) Len() int { return len(h.iters) - 1 }

func (h IHeap) swap(i, j int) {
	//fmt.Printf("SWAP %d %d\n", i, j)
	h.iters[i], h.iters[j] = h.iters[j], h.iters[i]
}

func (h *IHeap) Put(i *Iterator) {
	if !i.IsEmpty() {
		at := len(h.iters)
		h.iters = append(h.iters, i)
		h.raise(at)
	}
}

func (h *IHeap) Op() (op *Op) {
	if len(h.iters) > 1 {
		op = &h.iters[1].Op
	}
	return
}

func (h *IHeap) remove(i int) {
	h.iters[i] = h.iters[len(h.iters)-1]
	h.iters = h.iters[:len(h.iters)-1]
	h.sink(i)
}

func (h *IHeap) next(i int) {
	h.iters[i].Next()
	if h.iters[i].IsEmpty() {
		h.remove(i)
	} else {
		h.sink(i)
	}
}

func (h *IHeap) Next() (op *Op) {
	h.next(1)
	return h.Op()
}

func (h *IHeap) nexteq(i int, uuid UUID) {
	if h.iters[i].Spec[h.primary] == uuid {
		j := i << 1
		if j < len(h.iters) {
			if j+1 < len(h.iters) { // rightmost first!
				h.nexteq(j+1, uuid)
			}
			h.nexteq(j, uuid)
		}
		h.next(i)
	}
}

func (h *IHeap) NextPrim() (op *Op) {
	if !h.IsEmpty() {
		event := h.iters[1].Spec[h.primary]
		h.nexteq(1, event)
	}
	return h.Op()
}

func (h *IHeap) PutFrame(frame Frame) {
	b := frame.Begin()
	h.Put(&b)
}

func (h *IHeap) IsEmpty() bool {
	return len(h.iters) == 1
}

func (h *IHeap) Frame() (ret Frame) {
	for !h.IsEmpty() {
		ret.AppendOp(*h.Op())
		h.Next()
	}
	return
}

func (h *IHeap) Clear() {
	h.iters = h.iters[:1]
}
