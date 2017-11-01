package ron

import "testing"

func TestFrame_Split(t *testing.T) {
	frame := ParseFrameString("*lww#id1!:val=1#id2:0!:val=2")
	h1 := "*lww#id1!:val=1"
	h2 := "*lww#id2!:val=2"
	frame.Next()
	frame.Next()
	if frame.Term() != TERM_HEADER {
		t.Fail()
		return
	}
	id1, id2 := frame.Split()
	if id1.String() != h1 {
		t.Fail()
		t.Logf("\nneed: %s\nhave: %s\n", h1, id1)
	}
	if id2.String() != h2 {
		t.Fail()
		t.Logf("\nneed: %s\nhave: %s\n", h2, id2)
	}
	if id2.Type() != NewName("lww") {
		t.Fail()
	}
}
