package RON

import "testing"

func TestOp_Event(t *testing.T) {
	var a = [...]byte{1, 2, 3}
	var b = [...]byte{1, 2, 3}
	if a != b {
		t.Fail()
	}
}

// Op parser TODO
// [ ] pointer based sigs
// [ ] term (next op, eof exit)  +length
// [ ] error handling  -length
// [x] pointer shifting (.#@:)
// [x] values/atoms
// [ ] strconv -- value parsing methods

func TestParseOp (t *testing.T) {
    t.Log("Parser")
	var frame = ".lww#test-author@(time-origin:loc=1''>test\n"
	var op Op
    if len(frame)-1 != XParseOp ( []byte(frame), &op, &ZERO_OP ) {
		t.Fail()
	}
	if op.Type.String() != "lww" {
		t.Logf("'%s' %v != '%s'\n", op.Type.String(), []byte(op.Type.String()), "lww")
		t.Fail()
	}
	if op.Object.String() != "test-author" {
		t.Logf("'%s' %v != '%s'\n", op.Type.String(), []byte(op.Object.String()), "test-author")
		t.Fail()
	}
}

func BenchmarkParseOp(b *testing.B) {
	//var frame= ".lww#test-author@(time-origin:loc=1''>test\n"
	var frame= "@(time-origin:loc=1\n"
	var frames []byte = make([]byte, 0, len(frame)*b.N+10)
	for i := 0; i < b.N; i++ {
		frames = append(frames, []byte(frame)...)
	}
	var off int
	var op Op
	for i := 0; i < b.N; i++ {
		l := XParseOp(frames[off:], &op, &ZERO_OP)
		if l!=len(frame)-1 {
			b.Logf("off %d len %d", off, l)
			b.Fail()
		}
		off+=l+1
	}
}

/*
func BenchmarkIterator_Next(b *testing.B) {
	var clock = Clock{}
	var buf = make([]byte, 100*b.N)
	var frame = Frame{buf}
	var ops = make([]Op, b.N)
	for i := 0; i < b.N; i++ {
		ops[i] = CreateOp("lww", "test", clock.Time(), "field", "=1")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		frame.Append(ops[i])
	}
	iter := frame.First()
	for i := 0; i < b.N; i++ {
		if !iter.Op.Same(ops[i]) {
			b.Fail()
		}
	}
}
*/
