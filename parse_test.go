package ron

import (
	"math"
	"math/rand"
	"os"
	"strings"
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
	frame.Serializer.Format |= FORMAT_OP_LINES
	const ops = 30
	for j := 0; j < ops; j++ {
		at = j << 2
		frame.AppendStateHeader(NewSpec(uuids[at], uuids[at+1], uuids[at+2], uuids[at+3]))
	}
	t.Logf(frame.String())
	// recover, compare
	iter := frame.Rewind()
	for k := 0; k < ops; k++ {
		if iter.EOF() {
			t.Fail()
			t.Logf("Premature end: %d not %d, failed at %d\n", k, ops, iter.Parser.pos)
			break
		}
		at = k << 2
		for u := 0; u < 4; u++ {
			uuid := iter.UUID(u)
			if uuid != uuids[at+u] {
				t.Fail()
				t.Logf("uuid %d decoding failed in op#%d, '%s' should be '%s'", u, k, iter.UUID(u).String(), uuids[at+u].String())
			}
		}
		iter.Next()
	}
	if !iter.EOF() {
		t.Fail()
		t.Log("No end")
	}
}

func TestFrame_Next(t *testing.T) {
	ops := []string{"*a!", "*b=1", "*c=1!", "*d,", "*e,"}
	// "*a!*b=1*c=1!*d,*e,"
	frameStr := strings.Join(ops, "") + "."
	//t.Log(frameStr)
	frame := ParseFrame([]byte(frameStr))
	names := ""
	i, l := 0, 0
	for !frame.EOF() {
		l += len(ops[i])
		if i == len(ops)-1 {
			l++ //? ragel
		}
		if frame.Offset() != l {
			t.Fail()
			t.Logf("bad off: %d not %d '%s'", frame.Offset(), l, frameStr)
		} else {
			//t.Logf("OK %d %s", i, frame.Type().String())
		}
		i++
		names += frame.Type().String()
		frame.Parse()
	}
	if i != len(ops) || names != "abcde" {
		t.Logf("bad end: %d not %d, at %d, '%s' should be 'abcde'", i, len(ops), frame.Offset(), names)
		t.Fail()
	}
}

func TestFrame_EOF2(t *testing.T) {
	multi := []byte("*a#A:1!:2=2:3=3.*b#B:1?:2=2.")
	states := []int{RON_start, RON_start, RON_FULL_STOP, RON_start, RON_FULL_STOP}
	o := 0
	frame := MakeStream(128)
	for i := 0; i < len(multi); i++ {
		frame.AppendBytes(multi[i : i+1])
		if frame.Next() {
			if states[o] != frame.Parser.state {
				t.Fail()
				t.Logf("state %d at pos %d op %d, expected %d", frame.Parser.state, frame.Parser.pos, o, states[o])
				break
			} else {
				//t.Logf("OK state %d at pos %d op %d", frame.Parser.state, frame.Parser.pos, o)
			}
			if frame.Parser.state == RON_FULL_STOP {
				frame = MakeStream(1024)
			}
			o++

		}
		t.Log(i, frame.Parser.State())
	}
	if o != len(states) {
		t.Fail()
		t.Logf("%d ops, needed %d", o, len(states))
	}
}

func TestFrame_EOF(t *testing.T) {
	var streams = []string{
		"#id . #one! #two! #three!. ",
		"...",
		"#first#incomplete",
	}
	var states = [][]int{
		{RON_FULL_STOP, RON_start, RON_start, RON_FULL_STOP},
		{RON_FULL_STOP, RON_FULL_STOP, RON_FULL_STOP},
		{RON_start},
	}
	for k, stream := range streams {
		frame := ParseStream([]byte{})
		// feed by 1 char
		// EOF -> Rest()
		s := 0
		for i := 0; i < len(stream); i++ {
			frame.AppendBytes([]byte(stream[i : i+1]))
			frame.Next()
			//t.Log(k, i, stream[i:i+1], frame.pos, frame.Parser.State(), frame.IsComplete())
			if frame.IsComplete() {
				if s > len(states[k]) {
					t.Fail()
					t.Logf("stream %d off %d got %d need nothing", k, i, frame.Parser.State())
					break
				}
				if frame.Parser.State() != states[k][s] {
					t.Logf("stream %d off %d got %d need %d", k, i, frame.Parser.State(), states[k][s])
					t.Fail()
					break
				}
				s++
				if frame.Parser.State() == RON_FULL_STOP {
					frame = ParseStream(frame.Rest())
				}
			}
		}
		if s != len(states[k]) {
			t.Logf("need %d complete states, got %d", len(states[k]), s)
			t.Fail()
		}
	}
}

// A RON-text file must start with '*'
func TestXParseOpWhitespace(t *testing.T) {
	str := "*lww \t #test ;\n#next?"
	frame := ParseFrameString(str)
	if str[frame.Offset()-1] != '\n' {
		t.Fail()
	}
	frame.Next()
	if frame.Offset() != len(str) {
		t.Fail()
	}
}

func TestXParseMalformedOp(t *testing.T) {
	var tests = []string{
		"novalue",
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
		frame := ParseFrameString(str)
		if !frame.EOF() && frame.Offset() >= len(str) {
			t.Logf("parsed %d but invalid: '%s' (%d)", i, str, frame.Offset())
			t.Fail()
			break
		}
	}
}

func TestParseComment(t *testing.T) {
	tests := [][]string{
		{
			"*lww#object@time!:field'value' *~'comment'! *rga!.",
			"*rga!",
		},
		{
			"*lww#object@time! :field'value', *~'comment', *lww:another'value'.",
			"*lww #object @time :another'value',",
		},
	}
	for k, test := range tests {
		frame := ParseFrameString(test[0])
		correct := ParseFrameString(test[1])
		for frame.Parser.State() != RON_FULL_STOP {
			frame.Next()
		}
		eq := frame.Equal(correct)
		if !eq {
			t.Fail()
			t.Logf("%d need \n'%s'\n got \n'%s'\n", k, correct.OpString(), frame.OpString())
		}
	}
}

func TestFormatComment(t *testing.T) {
	frameStr := "*lww#obj@time!:key1'value1' *~'comment'! *ack@time!"
	frame := ParseFrameString(frameStr)
	clone := frame.Reformat(0)
	for !frame.EOF() && !clone.EOF() {
		eq := frame.Equal(clone)
		if !eq {
			t.Fail()
			t.Logf("%s != %s", frame.OpString(), clone.OpString())
			break
		}
		frame.Next()
		clone.Next()
	}
	q := ParseFrameString("*~?")
	if !q.IsQuery() {
		t.Fail()
	}
}

func TestParseTermDuplet(t *testing.T) {
	frameStr := "*lww#object@time+orig?! :keyA 'А' :keyB 'Б'"
	frame := ParseFrameString(frameStr)
	if !frame.IsQuery() {
		t.Log("no query parsed")
		t.Fail()
	}
	obj := frame.Object()
	frame.Next()
	if frame.EOF() || !frame.IsHeader() || frame.Object() != obj {
		t.Log("state header not parsed")
		t.Fail()
	}
	frame.Next()
	if frame.EOF() || frame.Term() != TERM_REDUCED || frame.Object() != obj {
		t.Log("inner op not parsed")
		t.Fail()
	}
}

/*
func TestOp_ParseFloat(t *testing.T) {
	var tests = []string{
		"*a^3.1415",
		"*a^1.0e6",
		"*a^1.2345e6",
		"*a^0",
	}
	var correct = []float64{
		3.1415,
		1e6,
		1.2345e6,
		0,
	}
	for i, str := range tests {
		frame := ParseFrameString(str)
		if frame.Offset() != len(str) {
			t.Logf("not parsed %d: '%s' (%d)", i, str, l)
			t.Fail()
			break
		}
		val := frame.Float(0)
		if val != correct[i] {
			t.Logf("misparsed %d: '%e' (%e)", i, val, correct[i])
			t.Fail()
		}
	}
}*/

func TestOp_ParseAtoms(t *testing.T) {
	var tests = [5][2]string{
		{"*a>0>1>2>3", ">>>>"},
		{"*a>0>0,#next>0>0", ">>"},
		{"*a,", ""},
		{"*a=1^2.0", "=^"},
		{"*a'str''quoted \\'mid\\' str'", "''"},
	}
	for i := 0; i < len(tests); i++ {
		str := tests[i][0]
		frame := ParseFrameString(str)
		if frame.EOF() {
			t.Logf("not parsed %d: '%s' (%d)", i, str, frame.Offset())
			t.Fail()
			break
		}
		types := ""
		for a := 0; a < frame.Count(); a++ {
			types += string(ATOM_PUNCT[frame.Atom(a).Type()])
		}
		if types != tests[i][1] {
			t.Logf("misparsed %d: '%s' (%s)", i, types, tests[i][1])
			t.Fail()
		}
	}
}

func TestOp_ParseFloat(t *testing.T) {
	frames := [][2]string{
		{"*lww#id^3.141592", "*lww#id^3.141592e+00"},
		{"*lww#id^-0.25", "*lww#id^-2.5e-01"},
		{"*lww#id^-25e-01", "*lww#id^-2.5e+00"},
		{"*lww#id^0.000001", "*lww#id^1e-06"},
		{"*lww#id^-0.00000e+02", "*lww#id^-0e+00"},
		{"*lww#id^-1.00000e+09", "*lww#id^-1e+09"},
		{"*lww#id^1000000000.0e-1", "*lww#id^1e+08"},
		{"*lww#id^12345.6789e+16", "*lww#id^1.23456789e+20"},
	}
	vals := []float64{
		3.141592,
		-0.25,
		-2.5,
		0.000001,
		0,
		-1e+9,
		1e+8,
		1.23456789e+20,
	}
	for i := 0; i < len(frames); i++ {
		frame := ParseFrameString(frames[i][0])
		if frame.Count() != 1 || frame.Atom(0).Type() != ATOM_FLOAT {
			t.Fail()
			t.Log("misparsed a float")
		}
		atom := frame.Atom(0)
		val := atom.Float()
		if math.Abs(val-vals[i]) > 0.001 {
			t.Fail()
			t.Logf("%d float value unparsed %e!=%e", i, val, vals[i])
		}
		back := NewFrame()
		back.Append(frame)
		if back.String() != frames[i][1] {
			t.Fail()
			t.Logf("float serialize fail (got/want):\n%s\n%s\n", back.String(), frames[i][1])
		}
	}
}

func TestParse_SpecOnly(t *testing.T) {
	str := "#test:)1#test:)2#test:)3"
	frame := ParseFrameString(str)
	c := 0
	for !frame.EOF() {
		c++
		if frame.Ref().Value() != uint64(c) {
			t.Fail()
		}
		frame.Next()
	}
	if c != 3 {
		t.Fail()
	}
}

func TestParse_Errors(t *testing.T) {
	frames := []string{
		"#test>linkмусор",
		"#string'unfinishe",
		"#id<",
		"#bad@term=?",
		"#no-term?-",
		"#notfloat^a",
		"#notesc'\\'",
		"*type=1NT",
		"*ty&",
	}
	for k, f := range frames {
		buf := []byte(f)
		frame := ParseFrame(buf)
		if frame.Parser.State() != RON_error {
			t.Fail()
			t.Logf("mistakenly parsed %d [ %s ] %d\n", k, f, frame.Offset())
		}
	}
}

func TestFrame_ParseStream(t *testing.T) {
	str := "*op1=123*op2!*op3!."
	frame := MakeStream(1024)
	count := 0
	for i := 0; i < len(str); i++ {
		frame.Body = append(frame.Body, str[i])
		frame.Parse()
		//fmt.Println(frame.Parser.state, "AT", frame.Offset(), string(frame.Body[:frame.Offset()]))
		if frame.IsComplete() {
			//fmt.Println("TADAAM", frame.OpString(), frame.Count(), "\n")
			count++
		}
	}
	if count != 3 {
		t.Logf("count %d!=3", count)
		t.Fail()
	}
}

func TestAtom_UUID(t *testing.T) {
	str := "*lww#1TUAQ+gritzko@`:bar=1 #(R@`:foo > (Q"
	frame := ParseFrameString(str)
	uuid1 := frame.Object()
	uuid2 := frame.Event()
	if uuid1 != uuid2 {
		t.Fail()
	}
	frame.Next()
	uuid3 := frame.Atom(0).UUID()
	if uuid1 != uuid3 {
		t.Fail()
	}
}

func TestParserState_Omitted(t *testing.T) {
	frame := ParseFrameString("*type#id!@ev@ev:")
	if frame.Parser.omitted != 4|8 {
		t.Logf("omitted is %d for %s", frame.Parser.omitted, frame.OpString())
		t.Fail()
	}
	frame.Next()
	if frame.Parser.omitted != 1|2|8 {
		t.Fail()
	}
	frame.Next()
	if frame.Parser.omitted != 1|2 {
		t.Fail()
	}
}
