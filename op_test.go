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

func TestParseOp(t *testing.T) {
	t.Log("Parser")
	var frame = "*lww#test-author@(time-origin:loc=1''>test"
	iter := NewStringIterator(frame)
	if iter.Spec.Type().String() != "lww" {
		t.Logf("'%s' %v != '%s'\n", iter.Type().String(), []byte(iter.Type().String()), "lww")
		t.Fail()
	}
	if iter.Spec.Object().String() != "test-author" {
		t.Logf("'%s' %v != '%s'\n", iter.Type().String(), []byte(iter.Object().String()), "test-author")
		t.Fail()
	}
	i := iter.Integer(0)
	if i != 1 {
		t.Logf("int parse fails: %d", iter)
		t.Fail()
	}
}

func BenchmarkParseOp(b *testing.B) {
	//var frame= "*lww#test-author@(time-origin:loc=1''>test\n"
	var frame = "@(time-origin:loc=1"
	var frames []byte = make([]byte, 0, len(frame)*b.N+10)
	for i := 0; i < b.N; i++ {
		frames = append(frames, []byte(frame)...)
	}
	origin, _ := ParseUUID([]byte("1-origin"))
	iter := NewBufferIterator(frames)
	var off int
	var op Op
	for i := 0; i < b.N; i++ {
		if op.Event().Origin() != origin.Origin() || op.Atoms.Count() != 1 || op.Atoms.AType(0) != ATOM_INT {
			b.Logf("parse fail: off %d len %d '%s'", off, len(frame), string(frames[off:]))
			b.Fail()
			break
		}
		iter.Next()
	}
}

/*
func BenchmarkIterator_Next(b *testing.B) {
	var clock = Clock{}
	var frame = MakeFrame(100 * b.N)
	var times = make([]UUID, b.N)
	var test_uuid, _ = ParseUUIDString("test")
	var field_uuid, _ = ParseUUIDString("field")
	var LWW_UUID = NewName("lww")
	atoms1 := ParseAtoms([]byte("=1"))
	for i := 0; i < b.N; i++ {
		time := clock.Time()
		times[i] = time
		frame.AppendReduced(Spec{uuids:[4]UUID{LWW_UUID, test_uuid, time, field_uuid}}, atoms1)
	}
	b.Logf("'%s'", string(frame.body))
	iter := frame.Begin()
	for i := 0; i < b.N; i++ {
		if !iter.Event().Equal(times[i]) {
			b.Logf("parse fail at %d, %s != %s", i, iter.Event().String(), times[i].String())
			b.Fail()
		}
	}
}
*/
