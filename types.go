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

const INT60LEN = 10

type UUID struct {
	Value  uint64
	Sign   byte // TODO maybe fit into 16 bytes
	Origin uint64
}

// OP is an immutable atomic operation object - no write access
type Op struct { // ~128 bytes
	uuids		[4]UUID
	AtomCount   int
	AtomTypes   [8]byte
	AtomOffsets [8]int
	Body        []byte
}

// Frame... mutable, but append-only
type Frame struct {
	Body []byte
	last Op
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
// [ ] sign = 0 1 2 3   UUID{} ~ ZERO_UUID, upper bits
//			Origin() vs Replica(), 128 bits, google memory layouts
// [x] end -- test
// [x] Op fields/array/GetUUID(i) [4]UUID  -- GetUUID(i), ABC
// [x] Format - nil context
// [x] open/closed frame => static error strings "=400'parsing error'"
//	   cause the end op can be displaced!!!
// cli FIXME
// [ ] iterator - parse error
// [ ] value parsing (all types - tables, safe ranges, length limits)
//		[ ] int
//		[ ] float
//		[ ] string
//		[ ] uuid (save?)
// [x] lww - idempotency
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

type Reducer func(a, b Iterator, to *Frame) UUID

type zip int8

const ZIP_SKIP zip = -1
const ZIP_KEEP zip = 1
const ZIP_WAIT zip = 0

type Zipper func(a Iterator, b Iterator) (left, right zip)

type RawUUID []byte

const BASE64 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz~"

var base64 = []byte(BASE64)
var ABC [256]int8

func (op *Op) Type () UUID {
	return op.uuids[0]
}

func (op *Op) Object () UUID {
	return op.uuids[1]
}

func (op *Op) Event () UUID {
	return op.uuids[2]
}

func (op *Op) Location () UUID {
	return op.uuids[3]
}

const SPEC_PUNCT = ".#@:"
const (
	SPEC_TYPE_SEP     = '.'
	SPEC_OBJECT_SEP   = '#'
	SPEC_EVENT_SEP    = '@'
	SPEC_LOCATION_SEP = ':'
)

var REDEF_PUNCT = "`\\|/"

// FIXME bracket order to match the numeric order!!!!   }{][)(
var PREFIX_PUNCT = "([{}])"

const UUID_PUNCT = "-+$%"
const NAME_UUID_SEP = byte('$')
const EVENT_UUID_SEP = byte('+')
const REF_UUID_SEP = byte('-')
const HASH_UUID_SEP = byte('%')

const INT60_ERROR uint64 = 1<<60 - 1

var ZERO_UUID = UUID{0, '$', 0}

var ERROR_UUID = UUID{INT60_ERROR, '$', 0}

var ZERO_OP = Op{uuids:[4]UUID{ZERO_UUID,ZERO_UUID,ZERO_UUID,ZERO_UUID}}

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
	for i:=0; i<len(SPEC_PUNCT); i++ {
		ABC[SPEC_PUNCT[i]] = -30 -int8(i)
	}
	TYPE_MISMATCH_ERROR_UUID, _ = ParseUUIDString("type_msmch$~~~~~~~~~~")
	UNKNOWN_TYPE_ERROR_UUID, _ = ParseUUIDString("type_unknw$~~~~~~~~~~")
}
