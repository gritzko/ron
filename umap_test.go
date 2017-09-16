package RON

import "testing"

func TestUUID2Map_List(t *testing.T) {
	um := MakeUUID2Map()
	um.Add(NEVER_UUID, ZERO_UUID)
	um.Add(NEVER_UUID, NEVER_UUID)
	for i := 0; i < 100; i++ {
		um.Add(ZERO_UUID, UUID{uint64(i), UUID_NAME_UPPER_BITS})
	}
	nv := um.List(NEVER_UUID)
	if len(nv) != 2 || nv[0] != ZERO_UUID || nv[1] != NEVER_UUID {
		t.Fail()
	}
	hu := um.List(ZERO_UUID)
	if len(hu) != 100 {
		t.Fail()
	}
	for i := 0; i < 100; i++ {
		if hu[i].Value != uint64(i) {
			t.Fail()
		}
	}
	for i := 0; i < 100; i+=2 {
		um.Remove(ZERO_UUID, UUID{uint64(i), UUID_NAME_UPPER_BITS})
	}
	odd := um.List(ZERO_UUID)
	for i := 1; i < 100; i+=2 {
		if odd[i].Value != uint64(i) {
			t.Fail()
		}
	}
}
