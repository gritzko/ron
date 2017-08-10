package RON

import "testing"

func TestIHeap_AddFrame(t *testing.T) {
	frameA := ".lww#test@time1-orig:number=1@(2:string'2'"
	frameB := ".lww#test@time3-orig:number=3@(4:string'4'"
	frameC := ".lww#test@time1-orig:number=1@(2:string'2'@(3:number=3@(4:string'4'"
	heap := IHeap{SortBy: SPEC_EVENT}
	heap.PutFrame(ParseFrame([]byte(frameA)))
	heap.PutFrame(ParseFrame([]byte(frameB)))
	C := heap.Frame()
	if C.String() != frameC {
		t.Fail()
		t.Logf("heap fail: '%s' != '%s'", C.String(), frameC)
	}
}

func TestIHeap_Op(t *testing.T) {
	frameA := ".lww#test@time1-orig:number=1@(2:string'2'"
	frameB := ".lww#test@time3-orig:number=3@(4:string'4'"
	frameC := ".lww#test@time2-orig:number=2@(2:string'2'@(3:number=3@(4:string'4'"
	heap := IHeap{SortBy: SPEC_LOCATION}
	heap.PutFrame(ParseFrame([]byte(frameA)))
	heap.PutFrame(ParseFrame([]byte(frameB)))
	heap.PutFrame(ParseFrame([]byte(frameC)))
	loc := heap.Op().Location()
	count := 0
	for heap.Op().Location()==loc {
		count ++
		heap.Next()
	}
	if count!=3 {
		t.Fail()
	}
}
