package RON

import (
	"bytes"
)



// serializer file
// FormatXXX
// String()

var REDEFS = []byte("`\\|/")

func unzipRedefaultSeparator(input []byte) (redef int8, length int) {
	if ABC[input[0]] == -2 {
		redef = int8(bytes.IndexByte(REDEFS, input[0]))
		length = 1
	}
	return
}

var PREFIXES = []byte("([{}])")

func unzipPrefixSeparator(input []byte) (prefix uint8, length int) {
	var i = ABC[input[0]]
	if i <= -10 {
		prefix = uint8(-i-10) + 4
		length = 1
	}
	return
}

const BASE64 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz~"

var base64 = []byte(BASE64)
var ABC [256]int8

func UnzipBase64(input []byte, number *uint64) int {

	// TODO migrate to Ragel

	var l int = 10
	if l > len(input) {
		l = len(input)
	}
	var i = 0
	var res uint64
	for ; i < l; i++ {
		code := ABC[input[i]]
		if code < 0 {
			break
		}
		res <<= 6
		res |= uint64(code)
	}
	if number != nil {
		*number = res
	}
	return i
}

const SIGNS = "-+$%"
const NAME_UUID = byte('$')
const EVENT_UUID = byte('+')
const REF_UUID = byte('-')
const HASH_UUID = byte('%')


func (t *UUID) Equal(b UUID) bool {
	return t.Value == b.Value && t.Origin == b.Origin
}

func CommonPrefix(value, context uint64) uint {
	// TODO use math.bits
	var xor = value ^ context
	if xor >= 1<<(6*6) {
		return 0
	}
	if xor == 0 {
		return 10
	}
	if xor >= 1<<(3*6) { // 456
		if xor >= 1<<(5*6) {
			return 4
		} else if xor >= 1<<(4*6) {
			return 5
		} else {
			return 6
		}
	} else { // 789
		if xor >= 1<<(2*6) {
			return 7
		} else if xor >= 1<<(1*6) {
			return 8
		} else {
			return 9
		}
	}
}

func ZeroTail(value *uint64) (tail uint) {
	if *value&((1<<30)-1) == 0 {
		tail += 5
		*value >>= 30
	}
	if *value&((1<<18)-1) == 0 {
		tail += 3
		*value >>= 18
	}
	if *value&((1<<12)-1) == 0 {
		tail += 2
		*value >>= 12
	}
	if tail < 10 && *value&((1<<6)-1) == 0 {
		tail += 1
		*value >>= 6
	}
	return
}

const prefix_mask uint64 = 0xffffff << 36

func ZipUUIDString(uuid, context UUID) string {
	var ret = make([]byte, 21, 21)
	len := FormatUUID(ret, uuid, context)
	return string(ret[0:len])
}

func (uuid *UUID) String() string {
	return ZipUUIDString(*uuid, ZERO_UUID)
}

func (a *UUID) LessThan(b UUID) bool {
	if a.Value == b.Value {
		if a.Sign == b.Sign {
			return a.Origin < b.Origin
		} else {
			return a.Sign < b.Sign
		}
	} else {
		return a.Value < b.Value
	}
}

const INT60_ERROR uint64 = 1<<60 - 1

func FormatInt(slice []byte, value, context uint64) (off int) {

	var prefix uint = CommonPrefix(value, context)
	if context==INT60_ERROR {
		prefix = 0 // FIXME
	}
	var shift uint = 60
	var mask uint64 = (1 << 60) - 1
	if prefix != 0 {
		if prefix == 10 {
			return 0
		}
		slice[0] = PREFIXES[prefix-4]
		off++
		shift -= prefix * 6
		mask = (1 << shift) - 1
	}
	for 0 != value&mask && shift < 64 {
		shift -= 6
		slice[off] = base64[(value>>shift)&63]
		mask >>= 6
		off++
	}
	if off == 0 {
		slice[0] = '0'
		off++
	}
	return
}

func FormatUUID(output []byte, uuid UUID, context UUID) int {

	if uuid==context && uuid!=ZERO_UUID { // FIXME options
		return 0
	}
	off := FormatInt(output, uuid.Value, context.Value)
	if uuid.Sign == NAME_UUID && uuid.Origin == 0 {
		return off
	}
	if uuid.Value == context.Value || uuid.Sign != context.Sign ||
		(uuid.Origin&prefix_mask) != (context.Origin&prefix_mask) {
		output[off] = uuid.Sign
		off++
	}
	l := FormatInt(output[off:], uuid.Origin, context.Origin)
	return off + l
}

func FormatOp(output []byte, op *Op, context *Op) int {
	var off int
	if context == nil {
		context = &ZERO_OP
	}
	// expand to 88+values
	if op.Type != context.Type {
		output[off] = TYPE_UUID_SEP
		off++
		off += FormatUUID(output[off:], op.Type, context.Type)
	}
	if op.Object != context.Object {
		output[off] = OBJECT_UUID_SEP
		off++
		off += FormatUUID(output[off:], op.Object, context.Object)
	}
	if op.Event != context.Event {
		output[off] = EVENT_UUID_SEP
		off++
		off += FormatUUID(output[off:], op.Event, context.Event)
	}
	if op.Location != context.Location {
		output[off] = LOCATION_UUID_SEP
		off++
		off += FormatUUID(output[off:], op.Location, context.Location)
	}
	copy(output[off:], op.Atoms)
	off += len(op.Atoms)
	return off
}

func (op *Op) String () string {
	buf := make([]byte, op.AtomOffsets[op.AtomCount-1]+100) // FIXME!!!
	l := FormatOp(buf, op, &ZERO_OP)
	return string(buf[:l])
}

var ZERO_UUID = UUID{0, '$', 0}

func init() {
	for i := 0; i < len(ABC); i++ {
		ABC[i] = -1
	}
	// optimize for close values
	// TODO -2 -3 -4
	for i := 0; i < len(REDEFS); i++ {
		ABC[REDEFS[i]] = -2
	}
	for i := 0; i < len(PREFIXES); i++ {
		ABC[PREFIXES[i]] = -10 - int8(i)
	}
	for i := 0; i < len(SIGNS); i++ {
		ABC[SIGNS[i]] = -4
	}
	for i := 0; i < len(BASE64); i++ {
		ABC[BASE64[i]] = int8(i)
	}
}
