package RON

import "testing"

func TestParseUUID(t *testing.T) {
	uuidA, _ := ParseUUID([]byte("1"), ZERO_UUID)
	if uuidA.Value != (1<<54) || uuidA.Origin != 0 || uuidA.Sign != '$' {
		t.Fail()
	}
	uuidAB, _ := ParseUUID([]byte(")1"), uuidA)
	if uuidAB.Value != (1<<54)|1 || uuidAB.Origin != 0 || uuidAB.Sign != '$' {
		t.Fail()
	}
	hello, _ := ParseUUID([]byte("hello-111"), ZERO_UUID)
	world, _ := ParseUUID([]byte("[world-111"), hello)
	helloworld, _ := ParseUUID([]byte("helloworld-111"), ZERO_UUID)
	if !world.Equal(helloworld) {
		t.Fail()
	}
}

func TestParseFormatUUID(t *testing.T) {
	tests := [][]string{
		{"0", "1", "1"},  // 0
		{"1-x", ")1", "1000000001-x"},
		{"test-1", "", "test-1"},
		{"hello-111", "[world", "helloworld-111"},
		{"helloworld-111", "[", "hello-111"},
		{"100001-orig", "[", "1-orig"}, // 5
		{"1+orig", "(2-", "10002-orig"},
		{"time+orig", "(1(2", "time1+orig2"},
		// TODO		{"name$user", "$scoped", "scoped$user"},
		{"any-thing", "hash%here", "hash%here"},
		{"[1s9L3-[Wj8oO", "-(2Biejq", "[1s9L3-(2Biejq"}, // 9
		{"0123456789-abcdefghij", ")~)~", "012345678~-abcdefghi~"},
		{"(2-[1jHH~", "-[00yAl", "(2-}yAl"},
	}
	for i, tri := range tests {
		context, _ := ParseUUID([]byte(tri[0]), ZERO_UUID)
		uuid, length := ParseUUID([]byte(tri[1]), context)
		if length < 0 {
			t.Logf("parse %d fail %s (context: %s)", i, tri[1], tri[0])
			t.Fail()
			continue
		}
		str := uuid.String()
		if str != tri[2] {
			t.Logf("parse %d: %s must be %s", i, str, tri[2])
			t.Fail()
		}
		var fmt [21]byte
		l := FormatUUID(fmt[:], uuid, context)
		zip := string(fmt[:l])
		if zip != tri[1] {
			t.Logf("format %d: %s must be %s", i, zip, tri[1])
			t.Fail()
		}
	}
}

func TestParseUUIDErrors(t *testing.T) {

}