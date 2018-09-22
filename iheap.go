package ron

import "sort"

//import "github.com/gritzko/RON"

// FrameHeap is an iterator heap - gives the minimum available element
// at every step. Useful for merge sort like algorithms.
type FrameHeap struct {
	// Most of the time, a heap has 2 elements, optimize for that.
	// Sometimes, it can get millions of elements, ensure that is O(NlogN)
	iters              []*Frame
	primary, secondary Comparator
}

func MakeFrameHeap(primary, secondary Comparator, size int) (ret FrameHeap) {
	ret.iters = make([]*Frame, 1, size+1)
	ret.primary = primary
	ret.secondary = secondary
	return
}

func (heap FrameHeap) less(i, j int) bool {
	c := heap.primary(heap.iters[i], heap.iters[j])
	if c == 0 {
		if heap.secondary != nil {
			c = heap.secondary(heap.iters[i], heap.iters[j])
		} else {
			c = int64(j) - int64(i)
		}
	}
	//fmt.Printf("CMP %s %s GOT %d\n", heap.iters[i].OpString(), heap.iters[j].OpString(), c)
	return c < 0
}

func (heap *FrameHeap) sink(i int) {
	to := i
	j := i << 1
	if j < len(heap.iters) && heap.less(j, i) {
		to = j
	}
	j++
	if j < len(heap.iters) && heap.less(j, to) {
		to = j
	}
	if to != i {
		heap.swap(i, to)
		heap.sink(to)
	}
}

func (heap *FrameHeap) raise(i int) {
	j := i >> 1
	if j > 0 && heap.less(i, j) {
		heap.swap(i, j)
		if j > 1 {
			heap.raise(j)
		}
	}
}

func (heap FrameHeap) Len() int { return len(heap.iters) - 1 }

func (heap FrameHeap) swap(i, j int) {
	//fmt.Printf("SWAP %d %d\n", i, j)
	heap.iters[i], heap.iters[j] = heap.iters[j], heap.iters[i]
}

func (heap *FrameHeap) PutAll(b Batch) {
	for i := 0; i < len(b); i++ {
		heap.Put(&b[i])
	}
}

func (heap *FrameHeap) Put(i *Frame) {
	for {
		if !i.EOF() && (i.IsHeader() || i.IsQuery()) {
			i.Next()
		} else {
			break
		}
	}
	if !i.EOF() && !i.IsHeader() {
		at := len(heap.iters)
		heap.iters = append(heap.iters, i)
		heap.raise(at)
	}
}

func (heap *FrameHeap) Current() (frame *Frame) {
	if len(heap.iters) > 1 {
		return heap.iters[1]
	} else {
		return nil
	}
}

func (heap *FrameHeap) remove(i int) {
	heap.iters[i] = heap.iters[len(heap.iters)-1]
	heap.iters = heap.iters[:len(heap.iters)-1]
	heap.sink(i)
}

func (heap *FrameHeap) next(i int) {
	heap.iters[i].Next()
	if heap.iters[i].EOF() || heap.iters[i].IsHeader() {
		heap.remove(i)
	} else {
		heap.sink(i)
	}
}

func (heap *FrameHeap) Next() (frame *Frame) {
	heap.next(1)
	for len(heap.iters) > 1 && heap.iters[1].Type() == COMMENT_UUID { // skip comments
		heap.next(1)
	}
	return heap.Current()
}

func (heap FrameHeap) listEqs(at int, eqs *[]int) {
	*eqs = append(*eqs, at)
	l := at << 1
	if l < len(heap.iters) {
		if 0 == heap.primary(heap.iters[1], heap.iters[l]) {
			heap.listEqs(l, eqs)
		}
		r := l | 1
		if r < len(heap.iters) {
			if 0 == heap.primary(heap.iters[1], heap.iters[r]) {
				heap.listEqs(r, eqs)
			}
		}
	}
}

func (heap *FrameHeap) NextPrim() (frame *Frame) {
	var _eqs [16]int
	eqs := _eqs[0:0:16]
	heap.listEqs(1, &eqs)
	if len(eqs) > 1 {
		sort.Ints(eqs)
	}
	for i := len(eqs) - 1; i >= 0; i-- {
		heap.next(eqs[i])
		heap.sink(eqs[i])
	}
	return heap.Current()
}

func (heap *FrameHeap) PutFrame(frame Frame) {
	heap.Put(&frame)
}

func (heap *FrameHeap) EOF() bool {
	return len(heap.iters) == 1
}

func (heap *FrameHeap) Frame() Frame {
	cur := MakeFrame(128)
	for !heap.EOF() {
		cur.Append(*heap.Current())
		heap.Next()
	}
	return cur.Close()
}

func (heap *FrameHeap) Clear() {
	heap.iters = heap.iters[:1]
}

func EventComparator(a, b *Frame) int64 {
	return a.Event().Compare(b.Event())
}

func EventComparatorDesc(a, b *Frame) int64 {
	return b.Event().Compare(a.Event())
}

func RefComparator(a, b *Frame) int64 {
	return a.Ref().Compare(b.Ref())
}

func RefComparatorDesc(a, b *Frame) int64 {
	return b.Ref().Compare(a.Ref())
}
