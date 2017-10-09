package RON

import "container/heap"

type UUIDHeap struct {
	uuids []UUID
	desc  bool
}

func MakeUHeap(desc bool, size int) (ret UUIDHeap) {
	ret.uuids = make([]UUID, 0, size)
	ret.desc = desc
	return
}

func (h *UUIDHeap) Less(i, j int) bool {
	c := Compare(h.uuids[i], h.uuids[j])
	if h.desc {
		c = -c
	}
	return c < 0
}

func (h UUIDHeap) Len() int { return len(h.uuids) }

func (h UUIDHeap) Swap(i, j int) {
	h.uuids[i], h.uuids[j] = h.uuids[j], h.uuids[i]
}

func (h *UUIDHeap) Push(x interface{}) {
	item := x.(UUID)
	h.uuids = append(h.uuids, item)
}

func (h *UUIDHeap) Pop() interface{} {
	n := len(h.uuids)
	item := h.uuids[n-1]
	h.uuids = h.uuids[0 : n-1]
	return item
}

func (h *UUIDHeap) Put(u UUID) {
	heap.Push(h, u)
}

func (h *UUIDHeap) Take() UUID {
	return heap.Pop(h).(UUID)
}

func (h *UUIDHeap) PopUnique() (ret UUID) {
	if len(h.uuids) == 0 {
		return ZERO_UUID
	}
	ret = heap.Pop(h).(UUID)
	for len(h.uuids) > 0 && h.uuids[0] == ret {
		heap.Pop(h)
	}
	return
}
