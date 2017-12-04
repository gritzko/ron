package ron

import "testing"

func TestVVector_Add(t *testing.T) {
	vec := make(VVector)
	vec.AddString("1+A")
	vec.AddString("2+B")
	vec.AddString("3-B")
	vec.AddString("4+A")
	A, _ := ParseUUIDString("+A")
	if vec.GetUUID(A).String() != "4+A" {
		t.Fail()
	}
	pB, _ := ParseUUIDString("+B")
	if vec.GetUUID(pB).String() != "2+B" {
		t.Fail()
	}
	mB, _ := ParseUUIDString("-B")
	if vec.GetUUID(mB).String() != "3-B" {
		t.Fail()
	}
}