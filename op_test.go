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
// [x] pointer based sigs
// [ ] term (next op, eof exit)  +length
// [ ] error handling  -length
// [x] pointer shifting (.#@:)
// [x] values/atoms
// [ ] strconv -- value parsing methods
// [ ] Iterator!!!

func TestParseOp (t *testing.T) {
    t.Log("Parser")
	var frame = ".lww#test-author@(time-origin:loc=1''>test"
	var op Op
	pl := XParseOp ( []byte(frame), &op, ZERO_OP )
    if len(frame) != pl {
		t.Fail()
	}
	if op.Type().String() != "lww" {
		t.Logf("'%s' %v != '%s'\n", op.Type().String(), []byte(op.Type().String()), "lww")
		t.Fail()
	}
	if op.Object().String() != "test-author" {
		t.Logf("'%s' %v != '%s'\n", op.Type().String(), []byte(op.Object().String()), "test-author")
		t.Fail()
	}
	i, e := op.ParseInt(0)
	if e!=nil || i!=1 {
		t.Logf("int parse fails: %d", i)
		t.Fail()
	}
}

func BenchmarkParseOp(b *testing.B) {
	//var frame= ".lww#test-author@(time-origin:loc=1''>test\n"
	var frame= "@(time-origin:loc=1"
	var frames []byte = make([]byte, 0, len(frame)*b.N+10)
	for i := 0; i < b.N; i++ {
		frames = append(frames, []byte(frame)...)
	}
	origin, _ := ParseUUID([]byte("1-origin"), ZERO_UUID)
	var off int
	var op Op
	for i := 0; i < b.N; i++ {
		l := XParseOp(frames[off:], &op, ZERO_OP)
		if l != len(frame) || op.Event().Origin != origin.Origin || op.Count !=1 || op.Types[0]!='=' {
			b.Logf("parse fail: off %d len %d(%d) '%s'", off, l, len(frame), string(frames[off:]))
			b.Fail()
			break
		}
		off+=l
	}
}


func BenchmarkIterator_Next(b *testing.B) {
	var clock = Clock{}
	var frame = MakeFrame(100*b.N)
	var times = make([]UUID, b.N)
	var test_uuid, _ = ParseUUIDString("test")
	var field_uuid, _ = ParseUUIDString("field")
	var LWW_UUID = UUID{881557636825219072, NAME_SIGN_BITS}
	for i := 0; i < b.N; i++ {
		time := clock.Time()
		times[i] = time
		frame.Append(Spec{LWW_UUID, test_uuid, time, field_uuid}, []byte("=1"))
	}
	b.Logf("'%s'", string(frame.Body))
	iter := frame.Begin()
	for i := 0; i < b.N; i++ {
		if !iter.Event().Equal(times[i]) {
			b.Logf("parse fail at %d, %s != %s", i, iter.Event().String(), times[i].String())
			b.Fail()
		}
	}
}

