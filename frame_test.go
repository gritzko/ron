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
	id1, id2 := frame.Split2()
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

func TestBatchFrames(t *testing.T) {
	frame1 := "*lww#A@1!:a=1:b=2:c=3"
	frame2 := "*lww#A@2!:d=4"
	var batch Batch
	batch = append(batch, ParseFrameString(frame1))
	batch = append(batch, ParseFrameString(frame2))
	frame12 := batch.Join()
	batchStr := "*lww#A@1!:a=1:b=2:c=3@2:0!:d=4"
	if frame12.String() != batchStr {
		t.Logf("\n%s != \n%s\n", frame12.String(), batchStr)
		t.Fail()
		return
	}
	b2 := frame12.Split()
	if len(b2) != 2 {
		t.Fail()
		t.Log("length", len(b2))
		return
	}
	if b2[0].String() != frame1 {
		t.Fail()
		t.Logf("%s != %s\n", b2[0].String(), frame1)
	}
	if b2[1].String() != frame2 {
		t.Fail()
		t.Logf("%s != %s\n", b2[0].String(), frame1)
	}
}

func TestFrame_SplitMultiframe(t *testing.T) {
	splits := []string{
		"*lww#test!:a=1#best:0!:b=2:c=3:d=4;",
		"*lww#test!:a=1",
		"*lww#best!:b=2:c=3",
		"*lww#best:d=4",
	}
	multi := ParseFrameString(splits[0])
	monos := multi.Split()
	for i := 0; i < len(monos); i++ {
		if monos[i].String() != splits[i+1] {
			t.Fail()
			t.Logf("split fail:\n'%s'\nshould be\n'%s'\n", monos[i].String(), splits[i+1])
		}
	}
}
