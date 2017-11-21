package ron

import (
	"testing"
)

func TestIHeap_AddFrame(t *testing.T) {
	frameA := "*lww#test@time1-orig:number=1@(2:string'2'"
	frameB := "*lww#test@time3-orig:number=3@(4:string'4'"
	frameC := "*lww#test@time1-orig:number=1@(2:string'2'@(3:number=3@(4:string'4'"
	heap := MakeFrameHeap(PRIM_EVENT, 2)
	heap.PutFrame(ParseFrame([]byte(frameA)))
	heap.PutFrame(ParseFrame([]byte(frameB)))
	C := heap.Frame()
	if C.String() != frameC {
		t.Fail()
		t.Logf("heap fail: \n'%s' must be \n'%s'", C.String(), frameC)
	}
}

func TestIHeap_Op(t *testing.T) {
	frameA := "*lww#test@time1-orig:number=1@(2:string'2'"
	frameB := "*lww#test@time3-orig:number=3@(4:string'4'"
	frameC := "*lww#test@time2-orig:number=2@(2:string'2'@(3:number=3@(4:string'4'"
	heap := MakeFrameHeap(PRIM_LOCATION, 2)
	heap.PutFrame(ParseFrame([]byte(frameA)))
	heap.PutFrame(ParseFrame([]byte(frameB)))
	heap.PutFrame(ParseFrame([]byte(frameC)))
	loc := heap.Current().Ref()
	count := 0
	for heap.Current().Ref() == loc {
		count++
		heap.Next()
	}
	if count != 3 {
		t.Fail()
	}
}

func TestIHeap_Merge(t *testing.T) {
	frameA := "*rga#test@1:0'A'@2'B'" //  D E A C B
	frameB := "*rga#test@1:0'A'@3'C'"
	frameC := "*rga#test@4:0'D'@5'E'"
	frameR := "*rga#test@4'D'@5'E'@1'A'@3'C'@2'B'"
	heap := MakeFrameHeap(PRIM_EVENT|PRIM_DESC|SEC_LOCATION, 4)
	heap.PutFrame(ParseFrame([]byte(frameA)))
	heap.PutFrame(ParseFrame([]byte(frameB)))
	heap.PutFrame(ParseFrame([]byte(frameC)))
	res := MakeFrame(128)
	for !heap.IsEmpty() {
		res.Append(*heap.Current())
		heap.NextPrim()
	}
	if res.String() != frameR {
		t.Fail()
		t.Logf("merge fail: \n'%s' must be \n'%s'", res.String(), frameR)
	}
}
