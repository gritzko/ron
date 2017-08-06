package RON

const (
	PREFIX10 = (1<<60 - 1) - (1<<(iota*6) - 1)
	PREFIX9
	PREFIX8
	PREFIX7
	PREFIX6
	PREFIX5
	PREFIX4
	PREFIX3
	PREFIX2
	PREFIX1
)

const NAME_SIGN uint64 = 0
const EVENT_SIGN uint64 = 2
const DERIVED_SIGN uint64 = 1
const HASH_SIGN uint64  = DERIVED_SIGN
const DERIVED_EVENT_SIGN = EVENT_SIGN|DERIVED_SIGN

const NAME_SIGN_BITS uint64 = NAME_SIGN<<60
const EVENT_SIGN_BIT uint64 = EVENT_SIGN<<60
const DERIVED_SIGN_BIT = DERIVED_SIGN<<60
const HASH_SIGN_BIT = HASH_SIGN<<60
const DERIVED_EVENT_SIGN_BITS = DERIVED_EVENT_SIGN<<60

const INT60LEN = 10

type UUID struct {
	Value  uint64
	//Sign   byte // TODO maybe fit into 16 bytes
	Origin uint64
}

// OP is an immutable atomic operation object - no write access
type Op struct { // ~128 bytes
	uuids       [4]UUID
	AtomCount   int
	AtomTypes   [8]byte
	AtomOffsets [8]int
	Body        []byte
}

// Frame... mutable, but append-only
type Frame struct {
	Body        []byte
	first, last Op
}

// Iterator is a mutable iterator over a frame; each position is an op.
type Iterator struct {
	Op
	frame  *Frame
	offset int
}

// Frame Open Q
// [x] ERRORS   !!!=code"text"
// [x]  including length limits!!!
// [x] ron CLI
// [x] whitespace
// [x] sign = 0 1 2 3  50% footprint red!!! UUID{} ~ ZERO_UUID, upper bits
//			Origin() vs Replica(), 128 bits, google memory layouts
// [x] end -- test
// [x] Op fields/array/GetUUID(i) [4]UUID  -- GetUUID(i), ABC
// [x] Format - nil context
// [x] open/closed frame => static error strings "=400'parsing error'"
//	   cause the end op can be displaced!!!

// [ ] Compare tests!!! (all types)
// [x] void atom , -- sweet  "op1, op2, op3" is perfectly OK
// [ ] op.Atoms && tests

// cli FIXME
// [ ] clean-up: uuid-grammar.rl
// [ ] iterator - parse error
// [ ] value parsing (all types - tables, safe ranges, length limits)
//		[ ] int
//		[ ] float
//		[ ] string
//		[ ] uuid (save?)
// [x] lww - idempotency
//
// [ ] RGA reducer (fn, errors)
//		[ ] Reduce()
//		[ ] tab tests
//
// [ ] fuzzer go-fuzz (need samples)
//
// [ ] reducer features
// [ ] error header   @~~~~~~~~~~:reference "error message" (to reduce)
// [ ] copy generic reduction errors
// [ ] struct Reducer - mimic Rocks, (a,b) or (a,b,c,d,...)
// [ ] prereduce - optional, may fail (RGA subtrees)
//
// [ ] formatting options
// [ ] indenting
// [ ] newlines
// [ ] trimming/zipping
// [ ] redefs (bench - fast prefix - bit ops)

type Reducer interface {
	Reduce(a, b Frame) (result Frame, err UUID)
	ReduceAll(inputs []Frame) (result Frame, err UUID)
}

var HEADER_ATOMS []byte = []byte("!")

type RawUUID []byte

const BASE64 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz~"

var base64 = []byte(BASE64)
var ABC [256]int8

func (op *Op) Type() UUID {
	return op.uuids[0]
}

func (op *Op) Object() UUID {
	return op.uuids[1]
}

func (op *Op) Event() UUID {
	return op.uuids[2]
}

func (op *Op) Location() UUID {
	return op.uuids[3]
}

const SPEC_PUNCT = ".#@:"

const (
	SPEC_TYPE_SEP     = "."
	SPEC_OBJECT_SEP   = '#'
	SPEC_EVENT_SEP    = '@'
	SPEC_LOCATION_SEP = ':'
)

var REDEF_PUNCT = "`\\|/"

// FIXME bracket order to match the numeric order!!!!   }{][)(
var PREFIX_PUNCT = "([{}])"

const UUID_PUNCT = "$%+-"
const NAME_UUID_SEP = byte('$')
const EVENT_UUID_SEP = byte('+')
const DERIVED_EVENT_SEP = byte('-')
const HASH_UUID_SEP = byte('%')

const INT60_ERROR uint64 = 1<<60 - 1
const INT60_NEVER = 63<<(6*9)

var ZERO_UUID = UUID{0, NAME_SIGN_BITS}

var NEVER_UUID = UUID{INT60_NEVER, NAME_SIGN_BITS}

var ERROR_UUID = UUID{INT60_ERROR, NAME_SIGN_BITS}

var ZERO_OP = Op{uuids: [4]UUID{ZERO_UUID, ZERO_UUID, ZERO_UUID, ZERO_UUID}}

var EMPTY_FRAME Frame

var TYPE_MISMATCH_ERROR_UUID UUID
var UNKNOWN_TYPE_ERROR_UUID UUID

func init() {
	for i := 0; i < len(ABC); i++ {
		ABC[i] = -1
	}
	for i := 0; i < len(REDEF_PUNCT); i++ {
		ABC[REDEF_PUNCT[i]] = -2
	}
	for i := 0; i < len(PREFIX_PUNCT); i++ {
		ABC[PREFIX_PUNCT[i]] = -10 - int8(i)
	}
	for i := 0; i < len(UUID_PUNCT); i++ {
		ABC[UUID_PUNCT[i]] = -4
	}
	for i := 0; i < len(BASE64); i++ {
		ABC[BASE64[i]] = int8(i)
	}
	for i := 0; i < len(SPEC_PUNCT); i++ {
		ABC[SPEC_PUNCT[i]] = -30 - int8(i)
	}
	TYPE_MISMATCH_ERROR_UUID, _ = ParseUUIDString("type_msmch$~~~~~~~~~~")
	UNKNOWN_TYPE_ERROR_UUID, _ = ParseUUIDString("type_unknw$~~~~~~~~~~")
}
