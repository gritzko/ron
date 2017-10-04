package RON

import (
	"math/rand"
	"os"
	"testing"
)

func TestParseUUID(t *testing.T) {
	uuidA, _ := ParseUUID([]byte("1"))
	if uuidA.Value() != (1<<54) || uuidA.Origin() != 0 || uuidA.Scheme() != UUID_NAME {
		t.Fail()
	}
	uuidAB, _ := uuidA.Parse([]byte(")1"))
	if uuidAB.Value() != (1<<54)|1 || uuidAB.Origin() != 0 || uuidAB.Scheme() != UUID_NAME {
		t.Fail()
	}
	hello, _ := ParseUUID([]byte("hello-111"))
	world, _ := hello.Parse([]byte("[world-111"))
	helloworld, _ := ParseUUID([]byte("helloworld-111"))
	if !world.Equal(helloworld) {
		t.Fail()
	}
	err_str := "erro_error$~~~~~~~~~~"
	error_uid, err := ParseUUIDString(err_str)
	if err != nil || error_uid.IsZero() {
		t.Fail()
	}
}

func TestParseFormatUUID(t *testing.T) {
	tests := [][]string{
		{"0", "1", "1"}, // 0
		{"1-x", ")1", "1000000001-x"},
		{"test-1", "-", "test-1"},
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
		context, _ := ParseUUID([]byte(tri[0]))
		uuid, err := context.Parse([]byte(tri[1]))
		if err != nil {
			t.Logf("parse %d fail %s (context: %s)", i, tri[1], tri[0])
			t.Fail()
			continue
		}
		str := uuid.String()
		if str != tri[2] {
			t.Logf("parse %d: %s must be %s", i, str, tri[2])
			t.Fail()
		}
		zip := uuid.ZipString(context)
		if zip != tri[1] {
			t.Logf("format %d: %s must be %s", i, zip, tri[1])
			t.Fail()
		}
	}
}

func TestParseUUIDErrors(t *testing.T) {

}

var test32 = [32][3]string{ // context: 0123456789-abcdefghi
	{"-", "0123456789-abcdefghi"},   // 00000
	{"B", "B"},                      // 00001
	{"(", "0123-abcdefghi"},         // 00010
	{"(B", "0123B-abcdefghi"},       // 00011
	{"+", "0123456789+abcdefghi"},   // 00100
	{"+B", "0123456789+B"},          // 00101
	{"+(", "0123456789+abcd"},       // 00110
	{"+(B", "0123456789+abcdB"},     // 00111
	{"A", "A"},                      // 01000 8
	{"AB", "AB"},                    // 01001
	{"A(", "A-abcd"},                // 01010
	{"A(B", "A-abcdB"},              // 01011
	{"A+", "A+abcdefghi"},           // 01100
	{"A+B", "A+B"},                  // 01101
	{"A+(", "A+abcd"},               // 01110
	{"A+(B", "A+abcdB"},             // 01111
	{")", "012345678-abcdefghi"},    // 10000 16
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
		next, err := def.Parse([]byte(zipped))
		if err != nil && test32[i][1] == "" {
			continue
		}
		if next != unzipped {
			t.Fail()
			t.Logf("uuid parse fail at %d: '%s' should be '%s' context %s len %d", i, next.String(), test32[i][1], defstr, len(zipped))
			break
		}
	}
}

func random_close_int(base uint64, prefix uint) uint64 {
	if prefix == 10 {
		return base
	}
	if prefix == 11 {
		return 0
	}
	var shift uint = (10 - prefix) * 6
	base >>= shift
	base <<= shift
	rnd := rand.Int() & 63
	base |= uint64(rnd << (shift - 6))
	return base
}

func TestParseFrame(t *testing.T) {
	pid := os.Getpid()
	t.Logf("random seed %d", pid)
	rand.Seed(int64(pid))
	defstr := "0123456789-abcdefghi"
	def, _ := ParseUUIDString(defstr)
	var at int
	// 64 random uuids - 8 brackets
	const dim = INT60LEN + 2
	var uuids [dim * dim]UUID
	for bv := 0; bv < dim; bv++ {
		for bo := 0; bo < dim; bo++ {
			v := random_close_int(def.Value(), uint(bv))
			o := random_close_int(def.Origin(), uint(bo))
			uuids[bv*dim+bo] = NewEventUUID(v, o)
		}
	}
	// shuffle to 16 ops
	for i := 0; i < 1000; i++ {
		var f, t int = rand.Int() % len(uuids), rand.Int() % len(uuids)
		uuids[f], uuids[t] = uuids[t], uuids[f]
	}
	// pack into a frame
	frame := MakeFrame(dim*dim*22 + dim*100)
	frame.Format |= FORMAT_OP_LINES
	const ops = 30
	for j := 0; j < ops; j++ {
		at = j << 2
		frame.AppendStateHeader(Spec{uuids: [4]UUID{uuids[at], uuids[at+1], uuids[at+2], uuids[at+3]}})
	}
	t.Logf(frame.String())
	// recover, compare
	iter := frame.Close().Begin()
	for k := 0; k < ops; k++ {
		if iter.IsEmpty() {
			t.Fail()
			t.Logf("Premature end: %d not %d, failed at %d\n", k, ops, iter.state.p)
			break
		}
		at = k << 2
		for u := 0; u < 4; u++ {
			uuid := iter.uuids[u]
			if uuid != uuids[at+u] {
				t.Fail()
				t.Logf("uuid %d decoding failed at %d, '%s' should be '%s' context: '%s' op: '%s'", u, k, iter.uuids[u].String(), uuids[at+u].String(), uuids[at+u-4].String(), string(iter.Op.String()))
			}
		}
		iter.Next()
	}
	if !iter.IsEmpty() {
		t.Fail()
		t.Log("No end")
	}
}

/*
func TestXParseOpWhitespace(t *testing.T) {
	str := []byte(" #test ;\n#next?")
	var op Op
	l := XParseOp(str, &op, ZERO_OP)
	if l != bytes.IndexByte(str, '\n')+1 {
		t.Fail()
	}
	l2 := XParseOp(str[l:], &op, ZERO_OP)
	if l+l2 != len(str) {
		t.Fail()
	}
}

func TestXParseMalformedOp(t *testing.T) {
	var tests = []string{
		"#novalue",
		"# broken - uuid?",
		"#too-many@values!!??=5=6^7.0^8.0'extra'",
		"#invalid-float ^31.41.5",
		"",
		"'unescaped ' quote'",
		">badreference",
		"#no_uuid-sep@$$",
		"#trailing garbage",
		"#reordered .uuids =1",
		"#repeat #uuids =1",
	}
	for i, str := range tests {
		var op Op
		l := XParseOp([]byte(str), &op, ZERO_OP)
		if l > 0 {
			t.Logf("parsed %d but invalid: '%s' (%d)", i, str, l)
			t.Fail()
		}
	}
}

func TestOp_ParseFloat(t *testing.T) {
	var tests = []string{
		"^3.1415",
		"^1.0e6",
		"^1.2345e6",
		"^0",
	}
	var correct = []float64{
		3.1415,
		1e6,
		1.2345e6,
		0,
	}
	for i, str := range tests {
		var op Op
		l := XParseOp([]byte(str), &op, ZERO_OP)
		if l != len(str) {
			t.Logf("not parsed %d: '%s' (%d)", i, str, l)
			t.Fail()
			break
		}
		val, _ := op.ParseFloat(0)
		if val != correct[i] {
			t.Logf("misparsed %d: '%e' (%e)", i, val, correct[i])
			t.Fail()
		}
	}
}

func TestOp_ParseAtoms(t *testing.T) {
	var tests = [5][2]string{
		{">0>1>2>3", ">>>>"},
		{">0>0,#next>0>0", ">>"},
		{",", ""},
		{"=1^2.0", "=^"},
		{"'str''quoted \\'mid\\' str'", "''"},
	}
	for i := 0; i < len(tests); i++ {
		var op Op
		str := tests[i][0]
		l := XParseOp([]byte(str), &op, ZERO_OP)
		if l <= 0 {
			t.Logf("not parsed %d: '%s' (%d)", i, str, l)
			t.Fail()
			break
		}
		types := ""
		for a := uint(0); a < op.Atoms.Count; a++ {
			types += string(atomBits2Sep(op.Atoms.AType(a)))
		}
		if types != tests[i][1] {
			t.Logf("misparsed %d: '%s' (%s)", i, types, tests[i][1])
			t.Fail()
		}
	}
}

func TestParse_Errors(t *testing.T) {
	frames := []string{
		"#test>linkмусор",
		"#string'unfinishe",
		"#id#id#id",
		"#bad@term??",
		"#no-term",
		"#notfloat^a",
		"#notesc'\\'",
		"*type=1NT",
		"*ty&",
	}
	for k, f := range frames {
		buf := []byte(f)
		var op Op
		l := XParseOp(buf, &op, ZERO_OP)
		if l > 0 {
			t.Fail()
			t.Logf("mistakenly parsed %d [ %s ] %d\n", k, f, l)
		}
	}
}

func TestFrame_SplitMultiframe(t *testing.T) {
	splits := []string{
		"*lww#test!:a=1#best:0!:b=2:c=3:d=4;",
		"*lww#test!:a=1",
		"*lww#best!:b=2:c=3",
		"*lww#best:d=4;",
	}
	multi := ParseFrameString(splits[0])
	monos, err := multi.SplitMultiframe(nil)
	if err != nil {
		t.Fail()
		t.Log(err)
	}
	for i := 0; i < len(monos); i++ {
		if monos[i].String() != splits[i+1] {
			t.Fail()
			t.Logf("split fail:\n'%s'\nshould be\n'%s'\n", monos[i].String(), splits[i+1])
		}
	}
}
*/
