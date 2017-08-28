package RON

import (
	"testing"
	//"sort"
	"math"
	"math/rand"
	//"bytes"
)

func TestUUIDPrimitives(t *testing.T) {
	var p, r uint8
	var l int
	p, l = unzipPrefixSeparator([]byte("[abc"))
	if p != 5 || l != 1 {
		t.Fail()
	}
	r, l = unzipPrefixSeparator([]byte("abc"))
	if r != 0 || l != 0 {
		t.Fail()
	}
	var n uint64
	l = UnzipBase64([]byte("10,"), &n)
	if n != 64 || l != 2 {
		t.Fail()
	}
}

func TestUUID_String(t *testing.T) {
	tests := [][]string{
		{"0", "1", "1"},
		{"0", "0000000001-orig", ")1-orig"},
		{"1time01-src", "1time02+src", "{2+"},
		{"0$author", "name$author2", "name{2"},
		{"hash%here", "hash%there", "%there"},
		{"1", ")1", "0000000001"}, //5
		{"0", "name$0", "name"},
		{"time+orig", "time1+orig2", "(1(2"},
		{"time-orig", "time1+orig2", "(1+(2"},
		{"[1s9L3-[Wj8oO", "[1s9L3-(2Biejq", "-(2Biejq"},
		{"}DcR-}L8w", "}IYI-", "}IYI}"}, //10
		{"}DcR-L8w", "}IYI-", "}IYI-0"},
		// FIXME: context A$B value A-B must produce "-"
	}
	for i, tri := range tests {
		context, _ := ParseUUID([]byte(tri[0]), ZERO_UUID)
		uuid, _ := ParseUUID([]byte(tri[1]), ZERO_UUID)
		zip := ZipUUIDString(uuid, context)
		if zip != tri[2] {
			t.Logf("case %d: %s must be %s", i, zip, tri[2])
			t.Fail()
		}
	}
}

func TestCommonPrefix(t *testing.T) {
	var a, b uint64 = 0, 1
	for i := 9; i > 0; i-- {
		pre := CommonPrefix(a, b)
		if i > 3 && int(pre) != i {
			t.Logf("prefix %d!=%d", pre, i)
			t.Fail()
		}
		a <<= 6
		b <<= 6
	}
	if CommonPrefix(100, 100) != 10 {
		t.Fail()
	}
	if CommonPrefix(0, math.MaxUint64) != 0 {
		t.Fail()
	}
}

func TestZeroTail(t *testing.T) {
	var a uint64 = 1
	for i := 0; i < 10; i++ {
		b := a
		tail := ZeroTail(&b)
		if int(tail) != i {
			t.Logf("tail %d!=%d", tail, i)
			t.Fail()
		}
		a <<= 6
	}
}

func RandUint() uint64 {
	var mask uint64 = (1 << 60) - 1
	var num = rand.Uint64() & mask
	length := rand.Uint32() % 60
	num >>= length
	if length > 6 {
		shift := rand.Uint32() % length
		num <<= shift
	}
	return num
}

func BenchmarkUnzip(b *testing.B) {
	uuids := make([]UUID, b.N)
	const m32 = 0xffffffff
	for i := 0; i < b.N; i++ {
		uuids[i] = UUID{RandUint(), UUID_EVENT_UPPER_BITS | RandUint()}
		//uuids[i] = UUID{uint64(i), '-', 100}
		// FIXME optimize close ids - bench CT/RGA
	}
	//sort.Slice(uuids, func(i, j int) bool { return uuids[i].LessThan(uuids[j]) })
	zipped := make([]byte, 0, b.N*22+22)
	lens := make([]int, b.N*2)

	zipped = FormatZippedUUID(zipped, uuids[0], ZERO_UUID)
	lens[0] = len(zipped)
	zipped = append(zipped, ' ')
	for i := 1; i < b.N; i++ {
		zipped = FormatZippedUUID(zipped, uuids[i], uuids[i-1])
		lens[i] = len(zipped)-lens[i-1]
		zipped = append(zipped, ' ')
	}

	b.ResetTimer()

	context := ZERO_UUID
	var ro int = 0
	for j := 0; j < b.N; j++ {
		ind := lens[j] //bytes.IndexByte(zipped[ro:], ' ')
		//unzip, l := ParseUUID(zipped[ro:ro+ind], context)
		unzip := context
		l := XParseUUID(zipped[ro:ro+ind], &unzip)
		if l < 0 {
			b.Logf("parse fail at %d: %s should be %s context %s text '%s'",
				j, unzip.String(), uuids[j].String(), context.String(), string(zipped[ro:]))
			b.Fail()
			break
		}
		if unzip != uuids[j] {
			b.Logf("incorrect unzip at %d: %s should be %s context %s len %d segm %v\n",
				j, unzip.String(), uuids[j].String(),
				context.String(), l, string(zipped[ro:ro+l]))
			b.Fail()
			break
		}
		ro += l
		//fmt.Println(unzip.String())
		ro += 1
		context = unzip
	}
	//fmt.Println("compressed")
	//b.Logf("%d bytes parsed\n", ro)
}

func TestOp_String(t *testing.T) {
	// FIXME EMPTY_OP.String() is ".0#0..." !!!
	str := "*lww#object@time-origin:loc=1"
	op, l := ParseOp([]byte(str), ZERO_OP)
	if l!= len(str) {
		t.Fail()
		t.Logf("misparsed %s", str)
		return
	}
	context := op
	op.Spec[2].Value++
	op.Spec[3].Value++
	buf := make([]byte, 0, 100)
	buf = FormatOp(buf, op, context)
	if len(buf) <= 0 {
		t.Fail()
		return
	}
	opstr := string(buf)
	correct := "@)1:)1=1"
	if opstr != correct {
		t.Logf("incorrect: '%s' != '%s'", opstr, correct)
		t.Fail()
	}
}

func BenchmarkFormatOp(b *testing.B) {
	str := "*lww#object@time-origin:loc=1"
	op, _ := ParseOp([]byte(str), ZERO_OP)
	var context Op = op
	buf := make([]byte, 0, b.N*len(str)*2+100)
	buf = FormatOp(buf, op, ZERO_OP)
	for i := 0; i < b.N; i++ {
		context = op
		op.Spec[2].Value++
		op.Spec[3].Value++
		buf = FormatOp(buf, op, context)
	}
}

type sample struct {
	flags int
	correct string
}

func TestFormatOptions(t *testing.T) {
	framestr := "*lww#test;@1:key'value'@2:number=1*rga#text@3'T'.*rga#text@6:3;@4'e'@5'x'@6't'"
	formats := []sample{
		{
			FORMAT_FRAME_NL | FORMAT_HEADER_SPACE,
			"*lww#test; @1:key'value'@2:number=1\n*rga#text@3'T'.\n@6:3; @4'e'@5'x'@6't'",
		},
	}
	frame := ParseFrame([]byte(framestr))
	for k, f := range formats {
		var formatted = Frame{Format:f.flags}
		i := frame.Begin()
		for !i.IsEmpty() {
			formatted.AppendOp(i.Op)
			i.Next()
		}
		if formatted.String() != f.correct {
			t.Fail()
			t.Logf("incorrect format at %d\n---\n%s\n---should be---\n%s\n", k, formatted.String(), f.correct)
		}
	}
}