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
const HASH_SIGN uint64 = DERIVED_SIGN
const DERIVED_EVENT_SIGN = EVENT_SIGN | DERIVED_SIGN

const NAME_SIGN_BITS uint64 = NAME_SIGN << 60
const EVENT_SIGN_BIT uint64 = EVENT_SIGN << 60
const DERIVED_SIGN_BIT = DERIVED_SIGN << 60
const HASH_SIGN_BIT = HASH_SIGN << 60
const DERIVED_EVENT_SIGN_BITS = DERIVED_EVENT_SIGN << 60

const INT60LEN = 10
const MAX_ATOMS = 7

type UUID struct {
	Value uint64
	//Sign   byte // TODO maybe fit into 16 bytes
	Origin uint64
}

type Spec [4]UUID

type Atoms struct {
	Count   int
	Types   [MAX_ATOMS + 1]byte
	Offsets [MAX_ATOMS]int
	Body    []byte
}

// OP is an immutable atomic operation object - no write access
type Op struct { // ~128 bytes
	Spec
	Atoms
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

// [ ] Compare tests!!! (all types, derived is +1?)
// [x] void atom , -- sweet  "op1, op2, op3" is perfectly OK
// [x] op.Atoms && tests
// [x] typedef Spec [4]UUID,
// [x] typedef Atoms, Atoms.Count()
// [ ] Location -> Reference
// [x] ?!,; term/mark/kind/status/headerness
// [ ] multiframe (still atomic)   Frame.Next() etc
// [x] AppendOp/Query/Patch/State - Spec/Atoms
//
// cli FIXME
// [ ] clean-up: uuid-grammar.rl
// [x] iterator - parse error
// [ ] value parsing (all types - tables, safe ranges, length limits)
//		[x] int
//		[x] float
//		[ ] string
//		[ ] uuid (save?)
// [x] lww - idempotency
//
// [ ] RGA reducer (fn, errors)
//		[x] Reduce()
//		[ ] tab tests
//		[ ] benchmark: 1mln ops
//
// [ ] fuzzer go-fuzz (need samples)
//
// [ ] reducer flags
// [ ] nice base64 constant definitions (ron ... // "comment")
// [ ] error header   @~~~~~~~~~~:reference "error message" (to reduce)
// [ ] copy generic reduction errors
// [x] struct Reducer - mimic Rocks, (a,b) or (a,b,c,d,...)
// [x] prereduce - optional, may fail (RGA subtrees)
// [ ] multiframe parsing  ;,,.,,!,. etc
//
// [ ] formatting options
// [ ] indenting
// [ ] newlines
// [ ] trimming/zipping
// [ ] redefs (bench - fast prefix - bit ops)

// Reducer is essentially a replicated data type.
// It provides two reducing functions: total and incremental.
// A reduction of the object's full op log produces its RON state.
// A reduction of a log segment produces a patch.
// A reduced frame has same type, object id; event id is the one
// of the last input frame.
type Reducer interface {
	// Reduce is a non-reordering incremental reducer.
	// It turns two adjacent frames into a single reduced frame,
	// if that is possible (quite often, two ops can not
	// be meaningfully combined without having the full state).
	// For a full op log, chained Reduce() must produce exactly
	// the same end result as ReduceAll()
	// Associative, commutative*, idempotent.
	Reduce(a, b Frame) (result Frame, err UUID)
	// ReduceAll is a reordering batch reducer. It turns a sequence
	// of frames into a reduced multiframe. In case the input is
	// the full log, the result must match that of chained Reduce().
	// Complexity guarantees: max O(log N)
	// (could be made to reduce 1mln single-op frames)
	// Associative, commutative*, idempotent.
	ReduceAll(inputs []Frame) (result Frame, err UUID)
}

var STATE_HEADER_ATOMS = ParseAtoms([]byte("!"))
var PATCH_HEADER_ATOMS = ParseAtoms([]byte(";"))
var RAW_OP_ATOMS = ParseAtoms([]byte("."))
var OP_ATOMS = ParseAtoms([]byte(","))

type RawUUID []byte

const BASE64 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz~"

var base64 = []byte(BASE64)
var ABC [256]int8

const (
	SPEC_TYPE = iota
	SPEC_OBJECT
	SPEC_EVENT
	SPEC_LOCATION
)

func (op Op) Type() UUID {
	return op.Spec[SPEC_TYPE]
}

func (op Op) Object() UUID {
	return op.Spec[SPEC_OBJECT]
}

func (op Op) Event() UUID {
	return op.Spec[SPEC_EVENT]
}

func (op Op) Location() UUID {
	return op.Spec[SPEC_LOCATION]
}

const SPEC_PUNCT = "*#@:"

const (
	SPEC_TYPE_SEP     = "*"
	SPEC_OBJECT_SEP   = '#'
	SPEC_EVENT_SEP    = '@'
	SPEC_LOCATION_SEP = ':'
)

const OP_PUNCT = ",.;!"
const (
	OP_SEP           = byte(',')
	RAW_OP_SEP       = byte('.')
	PATCH_HEADER_SEP = byte(';')
	STATE_HEADER_SEP = byte('!')
	QUERY_HEADER_SEP = byte('?')
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
const INT60_NEVER = 63 << (6 * 9)

var ZERO_UUID = UUID{0, NAME_SIGN_BITS}

var NEVER_UUID = UUID{INT60_NEVER, NAME_SIGN_BITS}

var ERROR_UUID = UUID{INT60_ERROR, NAME_SIGN_BITS}

var ZERO_OP = Op{}

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
	for i := 0; i < len(OP_PUNCT); i++ {
		ABC[OP_PUNCT[i]] = -5
	}
	TYPE_MISMATCH_ERROR_UUID, _ = ParseUUIDString("type_msmch$~~~~~~~~~~")
	UNKNOWN_TYPE_ERROR_UUID, _ = ParseUUIDString("type_unknw$~~~~~~~~~~")
}
