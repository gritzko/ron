package ron

import "testing"

func TestOp_Event(t *testing.T) {
	var a = [...]byte{1, 2, 3}
	var b = [...]byte{1, 2, 3}
	if a != b {
		t.Fail()
	}
}

func TestParseOp(t *testing.T) {
	t.Log("Parser")
	var frame = "*lww#test-author@(time-origin:loc=1''>test"
	iter := ParseFrameString(frame)
	if iter.Type().String() != "lww" {
		t.Logf("'%s' %v != '%s'\n", iter.Type().String(), []byte(iter.Type().String()), "lww")
		t.Fail()
	}
	if iter.Object().String() != "test-author" {
		t.Logf("'%s' %v != '%s'\n", iter.Type().String(), []byte(iter.Object().String()), "test-author")
		t.Fail()
	}
	t.Log(iter.OpString())
	i := iter.Integer(0)
	if i != 1 {
		t.Logf("int parse fails: %s", iter.String())
		t.Fail()
	}
}

func BenchmarkParseOp(b *testing.B) {
	//var frame= "*lww#test-author@(time-origin:loc=1''>test\n"
	var framestr = "@(time-origin:loc=1"
	var frames []byte = make([]byte, 0, len(framestr)*b.N+10)
	for i := 0; i < b.N; i++ {
		frames = append(frames, []byte(framestr)...)
	}
	b.Logf("a frame of %d bytes\n", len(frames))
	origin, _ := ParseUUID([]byte("1-origin"))
	b.ResetTimer()
	frame := ParseFrame(frames)
	var off int
	for i := 0; i < b.N; i++ {
		if frame.Event().Origin() != origin.Origin() || frame.Count() != 1 || frame.Atom(0).Type() != ATOM_INT {
			b.Logf("parse fail: off %d len %d '%s'", off, len(framestr), frame.OpString())
			b.Fail()
			break
		}
		frame.Next()
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
