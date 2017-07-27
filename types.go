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

type UUID struct {
	Value  uint64
	Sign   byte // TODO maybe fit into 16 bytes
	Origin uint64
}

// treat as an immutable object? - allocations
type Op struct { // ~128 bytes
	Type, Object, Event, Location UUID
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
// [ ] ERRORS   !!!=code"text"
// [ ] sign = 0 1 2 3   UUID{} ~ ZERO_UUID
// [ ] end -- test
// [ ] Op fields/array/GetUUID(i) [4]UUID  -- GetUUID(i), ABC
// [ ] Format - nil context
// [ ] open/closed frame => static error strings "=400'parsing error'"
//	   cause the end op can be displaced!!!

type Reducer func(a Iterator, b Iterator) Frame

type zip int8

const ZIP_SKIP zip = -1
const ZIP_KEEP zip = 1
const ZIP_WAIT zip = 0

type Zipper func(a Iterator, b Iterator) (left, right zip)

type RawUUID []byte

const BASE64 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz~"

var base64 = []byte(BASE64)
var ABC [256]int8

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

var ZERO_OP = Op{Type: ZERO_UUID, Object: ZERO_UUID, Event: ZERO_UUID, Location: ZERO_UUID, AtomCount: 1, AtomTypes: [8]byte{byte('?')}, AtomOffsets: [8]int{0}, Body: []byte("?")}

var EMPTY_FRAME Frame

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
}
