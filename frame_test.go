package ron

import "testing"

func TestFrame_Split(t *testing.T) {
	frame := ParseFrameString("*lww#id1!:val=1*#id2:0!:val=2")
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
	batchStr := "*lww#A@1!:a=1:b=2:c=3*#@2:0!:d=4"
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
	multiStr := "*lww#test!:a=1*#best:0!:b=2:c=3*#:d=4;"
	splits := []string{
		"*lww#test!:a=1",
		"*lww#best!:b=2:c=3",
		"*lww#best:d=4",
	}
	multi := ParseFrameString(multiStr)

	monos := multi.Split()
	for i := 0; i < len(monos); i++ {
		if monos[i].String() != splits[i] {
			t.Fail()
			t.Logf("split fail at %d:\n'%s'\nshould be\n'%s'\n", i, monos[i].String(), splits[i])
		}
	}
}

func TestBatch_Equal(t *testing.T) {
	b1 := ParseStringBatch([]string{"*one", "*two"})
	b2 := ParseStringBatch([]string{"*one*two"})
	if !b1.Equal(b2) {
		t.Fail()
	}
	b2 = append(b2, ParseFrameString("*three"))
	if b1.Equal(b2) {
		t.Fail()
	}
}

func TestFrame_Copy(t *testing.T) {
	a := ParseFrameString("*~'comment' *lww#obj!")
	b := a
	if b.Type() != COMMENT_UUID {
		t.Log("improper copy")
		t.Fail()
	}
	b.Next()
	if a.Type() != COMMENT_UUID {
		t.Log("the copy is still linked")
		t.Fail()
	}
}

func TestFrame_Split2(t *testing.T) {
	frame := ParseFrameString("*rga#test@4!@1'A'@2'B'*#@4:rm!:3,")
	split := frame.Split()
	if !split.Equal(Batch{frame}) {
		t.Fail()
		t.Logf("split fail, \n%s\nbecame\n%s", frame.String(), split.String())
	}
}
