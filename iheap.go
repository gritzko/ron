package RON

import "container/heap"

//import "github.com/gritzko/RON"

type IHeap struct {
	// most of the time, this is a h of 1-2 elements, optimize for that
	// sometimes, it can get millions of elements, ensure that is O(NlogN)
	iters []*Iterator
	SortBy int8
}

const LOC_ASC_EVENT_DESC int8 = SPEC_LOCATION*4-SPEC_EVENT // lww
const EVENT_ASC_LOC_ASC int8 = SPEC_EVENT*4+SPEC_LOCATION // rga rms
const EVENT_DESC_LOC_DESC int8 = -SPEC_EVENT*4-SPEC_LOCATION // rga concs
const EVENT_ASC int8 = SPEC_EVENT
const EVENT_DESC int8 = -SPEC_EVENT
const LOC_ASC int8 = SPEC_LOCATION
const LOC_DESC int = -SPEC_LOCATION

func (h IHeap) Less(i, j int) bool {
	switch 	h.SortBy {
		case SPEC_EVENT: return h.iters[i].Spec[SPEC_EVENT].EarlierThan(h.iters[j].Spec[SPEC_EVENT])
		case -SPEC_EVENT: return h.iters[j].Spec[SPEC_EVENT].LaterThan(h.iters[i].Spec[SPEC_EVENT])
		case SPEC_LOCATION: return h.iters[i].Spec[SPEC_LOCATION].EarlierThan(h.iters[j].Spec[SPEC_LOCATION])
		case -SPEC_LOCATION: return h.iters[j].Spec[SPEC_LOCATION].LaterThan(h.iters[i].Spec[SPEC_LOCATION])
		case LOC_ASC_EVENT_DESC: {
			i := Compare(h.iters[i].Spec[SPEC_LOCATION], h.iters[j].Spec[SPEC_LOCATION])
			if i==0 {
				i = Compare(h.iters[j].Spec[SPEC_EVENT], h.iters[i].Spec[SPEC_EVENT])
			}
			return i < 0
		}
		case EVENT_ASC_LOC_ASC: {
			i := Compare(h.iters[i].Spec[SPEC_EVENT], h.iters[j].Spec[SPEC_EVENT])
			if i==0 {
				i = Compare(h.iters[i].Spec[SPEC_LOCATION], h.iters[j].Spec[SPEC_LOCATION])
			}
			return i < 0
		}
		case EVENT_DESC_LOC_DESC: {
			i := Compare(h.iters[i].Spec[SPEC_EVENT], h.iters[j].Spec[SPEC_EVENT])
			if i==0 {
				i = Compare(h.iters[i].Spec[SPEC_LOCATION], h.iters[j].Spec[SPEC_LOCATION])
			}
			return i > 0
		}
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

func (h *IHeap) PutFrame(frame Frame) {
	b := frame.Begin()
	heap.Push(h, &b)
}

func (h *IHeap) PutIterator(i *Iterator) {
	if !i.IsEmpty() {
		heap.Push(h, i)
	}
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

func (h *IHeap) Clear () {
	h.iters = h.iters[:0]
}

