package RON

import (
	"testing"
	"math/rand"
)

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
		{"0", "1", "1"}, // 0
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
		{"0123G-abcdb", "(4566(efF", "01234566-abcdefF"},
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
		l := FormatZippedUUID(fmt[:], uuid, context)
		zip := string(fmt[:l])
		if zip != tri[1] {
			t.Logf("format %d: %s must be %s", i, zip, tri[1])
			t.Fail()
		}
	}
}

func TestParseUUIDErrors(t *testing.T) {

}

var test32 = [32][3]string{ // context: 0123456789-abcdefghi
	{"", "0123456789-abcdefghi"},    // 00000
	{"B", "B"},                      // 00001
	{"(", "0123-abcdefghi"},         // 00010
	{"(B", "0123B-abcdefghi"},       // 00011
	{"+", "0123456789+abcdefghi"},   // 00100
	{"+B", "0123456789+B"},          // 00101
	{"+(", "0123456789+abcd"},       // 00110
	{"+(B", "0123456789+abcdB"},     // 00111
	{"A", "A"},                      // 01000
	{"AB", "AB"},                    // 01001
	{"A(", "A-abcd"},                // 01010
	{"A(B", "A-abcdB"},              // 01011
	{"A+", "A+abcdefghi"},           // 01100
	{"A+B", "A+B"},                  // 01101
	{"A+(", "A+abcd"},               // 01110
	{"A+(B", "A+abcdB"},             // 01111
	{")", "012345678-abcdefghi"},    // 10000
	{")B", "012345678B-abcdefghi"},  // 10001
	{")(", "012345678-abcd"},        // 10010
	{")(B", "012345678-abcdB"},      // 10011
	{")+", "012345678+abcdefghi"},   // 10100
	{")+B", "012345678+B"},          // 10101
	{")+(", "012345678+abcd"},       // 10110
	{")+(B", "012345678+abcdB"},     // 10111
	{")A", "012345678A-abcdefghi"},  // 11000
	{")AB", ""},                     // 11001 error - length
	{")A(", "012345678A-abcd"},      // 11010
	{")A(B", "012345678A-abcdB"},    // 11011
	{")A+", "012345678A+abcdefghi"}, // 11100
	{")A+B", "012345678A+B"},        // 11101
	{")A+(", "012345678A+abcd"},     // 11110
	{")A+(B", "012345678A+abcdB"},   // 11111
}

func TestParseUUID2(t *testing.T) {
	defstr := "0123456789-abcdefghi"
	def, _ := ParseUUIDString(defstr)
	for i := 0; i < len(test32); i++ {
		zipped := test32[i][0]
		unzipped, _ := ParseUUIDString(test32[i][1])
		next, l := ParseUUID([]byte(zipped), def)
		if l<0 && test32[i][1]=="" {
			continue
		}
		if l != len(zipped) || next != unzipped {
			t.Fail()
			t.Logf("uuid parse fail at %d: '%s' should be '%s' context %s", i, next.String(), test32[i][1], defstr)
		}
	}
}

func random_close_int (base uint64, prefix uint) uint64 {
	if prefix == 10 {
		return base
	}
	var shift uint = (10-prefix)*6
	base >>= shift
	base <<= shift
	rnd := rand.Int()&63
	base |= uint64(rnd<<(shift-6))
	return base
}

func TestParseFrame(t *testing.T) {
	defstr := "0123456789-abcdefghi"
	def, _ := ParseUUIDString(defstr)
	var at int
	// 64 random uuids - 8 brackets
	var uuids [64]UUID
	for bv:=0; bv<8; bv++ {
		for bo:=0; bo<8; bo++ {
			v := random_close_int(def.Value, uint(bv))
			o := random_close_int(def.Origin, uint(bo))
			uuids[bv<<3+bo] = UUID{v, '-', o}
		}
	}
	// shuffle to 16 ops
	for i:=0; i<1000; i++ {
		var f,t int = rand.Int()&63, rand.Int()&63
		uuids[f], uuids[t] = uuids[t], uuids[f]
	}
	// pack into a frame
	frame := MakeFrame(64*4*22+640)
	for j:=0; j<16; j++ {
		at = j<<2
		frame.Append(uuids[at], uuids[at+1], uuids[at+2], uuids[at+3], []byte("!"))
	}
	// recover, compare
	iter := frame.Begin()
	for k :=0; k <16; k++ {
		if iter.AtEnd() {
			t.Fail()
			t.Log("Premature end")
			break
		}
		at = k << 2
		if iter.Type!=uuids[at] {
			t.Fail()
			t.Logf("type decoding failed at %d, '%s' should be '%s' op: '%s'", k, iter.Type.String(), uuids[at].String(), string(iter.Op.Body))
			break
		}
		if iter.Object!=uuids[at+1] {
			t.Fail()
			t.Logf("object decoding failed at %d, '%s' should be '%s'", k, iter.Object.String(), uuids[at+1].String())
		}
		if iter.Event!=uuids[at+2] {
			t.Fail()
			t.Logf("event decoding failed at %d, '%s' should be '%s'", k, iter.Event.String(), uuids[at+2].String())
		}
		if iter.Location!=uuids[at+3] {
			t.Fail()
			t.Logf("location decoding failed at %d, '%s' should be '%s'", k, iter.Location.String(), uuids[at+3].String())
		}
		iter.Next()
	}
	if !iter.AtEnd() {
		t.Fail()
		t.Log("No end")
	}
}