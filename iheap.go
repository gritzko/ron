package ron

//import "github.com/gritzko/RON"

// FrameHeap is an iterator heap - gives the minimum available element
// at every step. Useful for merge sort like algorithms.
type FrameHeap struct {
	// Most of the time, a heap has 2 elements, optimize for that.
	// Sometimes, it can get millions of elements, ensure that is O(NlogN)
	iters               []*Frame
	primary, secondary  int
	prim_desc, sec_desc bool
}

// sort modes, e.g. PRIM_EVENT|PRIM_DESC|SEC_LOCATION
const (
	PRIM_DESC     = 4
	PRIM_TYPE     = SPEC_TYPE
	PRIM_OBJECT   = SPEC_OBJECT
	PRIM_EVENT    = SPEC_EVENT
	PRIM_LOCATION = SPEC_REF
	SEC_DESC      = 32
	SEC_TYPE      = SPEC_TYPE << 3
	SEC_OBJECT    = SPEC_OBJECT << 3
	SEC_EVENT     = SPEC_EVENT << 3
	SEC_LOCATION  = SPEC_REF << 3
)

func MakeFrameHeap(mode, size int) (ret FrameHeap) {
	ret.iters = make([]*Frame, 1, size+1)
	ret.prim_desc = (mode & PRIM_DESC) != 0
	ret.sec_desc = (mode & SEC_DESC) != 0
	ret.primary = mode & 3
	ret.secondary = (mode >> 3) & 3
	return
}

func (heap FrameHeap) less(i, j int) bool {
	c := Compare(UUID(heap.iters[i].atoms[heap.primary]), UUID(heap.iters[j].atoms[heap.primary]))
	if c == 0 {
		c = Compare(UUID(heap.iters[i].atoms[heap.secondary]), UUID(heap.iters[j].atoms[heap.secondary]))
		if heap.sec_desc {
			c = -c
		}
	} else if heap.prim_desc {
		c = -c
	}
	//fmt.Printf("CMP %s %s %d\n", h.iters[i].String(), h.iters[j].String(), c)
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

func (heap *FrameHeap) Put(i *Frame) {
	if !i.EOF() && i.IsHeader() {
		i.Next()
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
	return heap.Current()
}

func (heap *FrameHeap) nexteq(i int, uuid UUID) {
	if heap.iters[i].UUID(heap.primary) == uuid {
		j := i << 1
		if j < len(heap.iters) {
			if j+1 < len(heap.iters) { // FIXME rightmost first!
				heap.nexteq(j+1, uuid)
			}
			heap.nexteq(j, uuid)
		}
		heap.next(i)
		for i < len(heap.iters) && heap.iters[i].UUID(heap.primary) == uuid {
			heap.next(i) // FIXME this fix (recheck after removal)
		}
	}
}

func (heap *FrameHeap) NextPrim() (frame *Frame) {
	if !heap.IsEmpty() {
		event := heap.iters[1].UUID(heap.primary)
		heap.nexteq(1, event)
	}
	return heap.Current()
}

func (heap *FrameHeap) PutFrame(frame Frame) {
	heap.Put(&frame)
}

func (heap *FrameHeap) IsEmpty() bool {
	return len(heap.iters) == 1
}

func (heap *FrameHeap) Frame() Frame {
	cur := MakeFrame(128)
	for !heap.IsEmpty() {
		cur.Append(*heap.Current())
		heap.Next()
	}
	return cur.Close()
}

func (heap *FrameHeap) Clear() {
	heap.iters = heap.iters[:1]
}
