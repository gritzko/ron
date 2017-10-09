package RON

import (
	"math/rand"
	"testing"
)

func TestUUID_String(t *testing.T) {
	tests := [][]string{
		{"}DcR-L8w", "}IYI-", "}IYI-0"},
		{"0", "1", "1"},
		{"0", "123-0", "123-"},
		{"0", "0000000001-orig", ")1-orig"},
		{"1time01-src", "1time02+src", "{2+"},
		{"0$author", "name$author2", "name{2"},
		{"hash%here", "hash%there", "%there"},
		{"1", ")1", "0000000001"}, //7
		{"0", "name$0", "name"},
		{"time+orig", "time1+orig2", "(1(2"},
		{"time-orig", "time1+orig2", "(1+(2"},
		{"[1s9L3-[Wj8oO", "[1s9L3-(2Biejq", "-(2Biejq"},
		{"}DcR-}L8w", "}IYI-", "}IYI}"}, //12
		{"A$B", "A-B", "-"},
	}
	for i, tri := range tests {
		context, e1 := ParseUUID([]byte(tri[0]))
		uuid, e2 := ParseUUID([]byte(tri[1]))
		if e1 != nil || e2 != nil {
			t.Fail()
			t.Log(e1, e2)
			break
		}
		zip := uuid.ZipString(context)
		if zip != tri[2] {
			t.Logf("case %d: %s must be %s (%s, %s)", i, zip, tri[2], uuid.String(), context.String())
			t.Fail()
		}
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
		uuids[i] = NewEventUUID(RandUint(), RandUint())
		//uuids[i] = UUID{uint64(i), '-', 100}
	}
	//sort.Slice(uuids, func(i, j int) bool { return uuids[i].LessThan(uuids[j]) })
	zipped := make([]byte, 0, b.N*22+22)
	lens := make([]int, b.N*2)

	zipped = append(zipped, uuids[0].ZipString(ZERO_UUID)...)
	lens[0] = len(zipped)
	zipped = append(zipped, ' ')
	for i := 1; i < b.N; i++ {
		zipped = append(zipped, uuids[i].ZipString(uuids[i-1])...)
		lens[i] = len(zipped) - lens[i-1]
		zipped = append(zipped, ' ')
	}

	b.ResetTimer()

	context := ZERO_UUID
	var ro int = 0
	for j := 0; j < b.N; j++ {
		ind := lens[j] //bytes.IndexByte(zipped[ro:], ' ')
		//unzip, l := ParseUUID(zipped[ro:ro+ind], context)
		unzip := context
		unzip, err := context.Parse(zipped[ro : ro+ind])
		if err != nil {
			b.Logf("parse fail at %d: %s should be %s context %s text '%s' err %s",
				j, unzip.String(), uuids[j].String(), context.String(), string(zipped[ro:]), err.Error())
			b.Fail()
			break
		}
		if unzip != uuids[j] {
			b.Logf("incorrect unzip at %d: %s should be %s context %s len %d segm %v\n",
				j, unzip.String(), uuids[j].String(),
				context.String(), lens[j], string(zipped[ro:]))
			b.Fail()
			break
		}
		ro += lens[j]
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
	op := ParseOp([]byte(str))
	if op.Atoms.Count() != 1 {
		t.Fail()
		t.Logf("misparsed %s", str)
		return
	}
	context := op
	op.uuids[2].uint128[0]++
	op.uuids[3].uint128[0]++
	cur := MakeFrame(1024)
	cur.AppendOp(context)
	le := cur.Len()
	cur.AppendOp(op)
	opstr := string(cur.Body()[le:])
	correct := "@)1:)1=1"
	if opstr != correct {
		t.Logf("incorrect: '%s' != '%s'", opstr, correct)
		t.Fail()
	}
}

func BenchmarkFormatOp(b *testing.B) {
	str := "*lww#object@time-origin:loc=1"
	op := ParseOp([]byte(str))
	frame := MakeFrame(b.N*len(str)*2 + 100)
	frame.AppendOp(op)
	for i := 0; i < b.N; i++ {
		op.uuids[2].uint128[0]++
		op.uuids[3].uint128[0]++
		frame.AppendOp(op)
	}
}

type sample struct {
	flags   uint
	correct string
}

func TestFormatOptions(t *testing.T) {
	framestr := "*lww#test!@1:key'value'@2:number=1*rga#text@3'T'!*rga#text@6:3,@4'e'@5'x'@6't'*lww#more:a=1;"
	formats := []sample{
		{
			FORMAT_FRAME_LINES | FORMAT_HEADER_SPACE,
			"*lww#test! @1:key'value'@2:number=1\n*rga#text@3'T'! @6:3,@4'e'@5'x'@6't'\n*lww#more:a=1;",
		},
	}
	frame := OpenFrame([]byte(framestr))
	if frame.EOF() {
		t.Fail()
		return
	}
	for k, f := range formats {
		formatted := MakeFormattedFrame(f.flags, 1024)
		for !frame.EOF() {
			formatted.AppendOp(frame.Op)
			frame.Next()
		}
		if formatted.String() != f.correct {
			t.Fail()
			t.Logf("incorrect format at %d\n---\n%s\n---should be---\n%s\n", k, formatted.String(), f.correct)
		}
	}
}
