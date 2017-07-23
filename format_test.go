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


func TestZip(t *testing.T) {
	tests := [][]string{
		{"0", "1", "1"},
		{"0", "0000000001-orig", ")1-orig"},
		{"1time01-src", "1time02+src", "{2+"},
		{"0$author", "name$author2", "name{2"},
		{"hash%here", "hash%there", "%there"},
		{"1", ")1", "0000000001"},
		{"0", "name$0", "name"},
		{"time+orig", "time1+orig2", "(1(2"},
		{"time-orig", "time1+orig2", "(1+(2"},
		{"[1s9L3-[Wj8oO", "[1s9L3-(2Biejq", "-(2Biejq"},
		{"}DcR-}L8w", "}IYI-", "}IYI}"},
		{"}DcR-L8w", "}IYI-", "}IYI-0"},
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
		uuids[i] = UUID{RandUint(), '-', RandUint()}
		//uuids[i] = UUID{uint64(i), '-', 100}
		// FIXME optimize close ids - bench CT/RGA
	}
	//sort.Slice(uuids, func(i, j int) bool { return uuids[i].LessThan(uuids[j]) })
	zipped := make([]byte, b.N*22+22)
	lens := make([]int, b.N*2)

	off := FormatUUID(zipped, uuids[0], ZERO_UUID)
	lens[0] = off
	zipped[off] = ' '
	off++
	for i := 1; i < b.N; i++ {
		l := FormatUUID(zipped[off:], uuids[i], uuids[i-1])
		off += l
		lens[i] = l
		zipped[off] = ' '
		off++
	}

	b.ResetTimer()

	context := ZERO_UUID
	var ro int = 0
	for j := 0; j < b.N; j++ {
		ind := lens[j] //bytes.IndexByte(zipped[ro:], ' ')
		//unzip, l := ParseUUID(zipped[ro:ro+ind], context)
		unzip := context
		l := XParseUUID(zipped[ro:ro+ind], &unzip)
		if l<0 {
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

// TODO
// [x] bench
// [x] prettify
// [x] coverage
// [x] table tests
// [x] hashes
// [x] names
// [x] order
// [ ] clock
