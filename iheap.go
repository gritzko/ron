package RON

import "container/heap"

//import "github.com/gritzko/RON"

type IHeap struct {
	// most of the time, this is a h of 1-2 elements, optimize for that
	// sometimes, it can get millions of elements, ensure that is O(NlogN)
	iters []*Iterator
	SortBy int
}

func (h IHeap) Less(i, j int) bool {
	switch 	h.SortBy {
		case SPEC_EVENT: return h.iters[i].Spec[SPEC_EVENT].EarlierThan(h.iters[j].Spec[SPEC_EVENT])
		case -SPEC_EVENT: return h.iters[j].Spec[SPEC_EVENT].LaterThan(h.iters[i].Spec[SPEC_EVENT])
		case SPEC_LOCATION: return h.iters[i].Spec[SPEC_LOCATION].EarlierThan(h.iters[j].Spec[SPEC_LOCATION])
		case -SPEC_LOCATION: return h.iters[j].Spec[SPEC_LOCATION].LaterThan(h.iters[i].Spec[SPEC_LOCATION])
		default : panic("unsupported sort mode")
	}
}

func (h IHeap) Len() int { return len(h.iters) }

func (h IHeap) Swap(i, j int) {
	h.iters[i], h.iters[j] = h.iters[j], h.iters[i]
}

func (h *IHeap) Push(x interface{}) {
	item := x.(*Iterator)
	h.iters = append(h.iters, item)
}

func (h *IHeap) Pop() interface{} {
	n := len(h.iters)
	item := h.iters[n-1]
	h.iters = h.iters[0 : n-1]
	return item
}

func (h *IHeap) Op() (op Op) {
	return h.iters[0].Op
}

func (h *IHeap) Next() (op Op) {
	h.iters[0].Next()
	if h.iters[0].IsEmpty() {
		heap.Pop(h)
	} else {
		heap.Fix(h, 0)
	}
	if len(h.iters)>0 {
		op = h.iters[0].Op
	}
	return
}

func (h *IHeap) AddFrame(frame Frame) {
	b := frame.Begin()
	heap.Push(h, &b)
}

func (h *IHeap) AddIterator(i *Iterator) {
	heap.Push(h, i)
}

func (h *IHeap) IsEmpty() bool {
	return h.Len() == 0
}

func (h *IHeap) Frame () (ret Frame) {
	for ! h.IsEmpty() {
		ret.AppendOp(h.Op())
		h.Next()
	}
	return
}

/*

func (h *IHeap) sink(i int) {
	u := h.SortBy
	to := i
	j := i<<1
	if j<len(h.iters) && h.iters[j].Spec[u].LaterThan(h.iters[to].Spec[u]) {
		to = j
	}
	j++
	if j<len(h.iters) && h.iters[j].Spec[u].LaterThan(h.iters[to].Spec[u]) {
		to = j
	}
	if to != i {
		h.iters[to], h.iters[i] = h.iters[i], h.iters[to]
		h.sink(to)
	}
}

func (h *IHeap) pop(i int) {
	j := i>>1
	u := h.SortBy
	if h.iters[i].Spec[u].LaterThan(h.iters[j].Spec[u]) {
		h.iters[j], h.iters[i] = h.iters[i], h.iters[j]
	}
	if j>1 {
		h.pop(j)
	}
}
 */