package RON

import "testing"

func TestUHeap_TakeUUID(t *testing.T) {
	var h UHeap
	h.Put(ZERO_UUID)
	h.Put(ZERO_UUID)
	h.Put(NEVER_UUID)
	h.Put(NEVER_UUID)
	h.Put(NEVER_UUID)
	if h.Len()!=5 {
		t.Fail()
	}
	if h.Take()!=ZERO_UUID {
		t.Fail()
	}
	if h.TakeUUID()!=ZERO_UUID {
		t.Fail()
	}
	if h.Len()!=3 {
		t.Fail()
	}
	if h.TakeUUID()!=NEVER_UUID {
		t.Fail()
	}
	if h.Len()!=0 {
		t.Fail()
	}
}
