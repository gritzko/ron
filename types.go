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
const MAX_ATOMS = 8

type UUID struct {
	Value uint64
	Origin uint64
}

type Spec [4]UUID

type Atoms struct {
	Count   uint
	Types   uint
	Offsets [MAX_ATOMS]uint
	Body    []byte
}

// OP is an immutable atomic operation object - no write access
type Op struct { // ~128 bytes
	Spec
	Atoms
	Flags uint
}

// Frame... mutable, but append-only
type Frame struct {
	Body        []byte
	first, last Op
	Format		int
	Source      int
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

// [ ] Compare tests!!! [1,2,3] - mix,sort,cmp (all types, derived is +1)
// [x] void atom , -- sweet  "op1, op2, op3" is perfectly OK
// [x] op.Atoms && tests
// [x] typedef Spec [4]UUID,
// [x] typedef Atoms, Atoms.Count()
// [x] Location -> Reference
// [x] ?!,; term/mark/kind/status/headerness
// [x] AppendOp/Query/Patch/State - Spec/Atoms
// [ ] multiframe (still atomic)   Frame.Split(iterator)
//
// cli FIXME
// [ ] clean-up: uuid-grammar.rl
// [x] iterator - parse error
// [ ] value parsing NEW DEAL
//		[ ] (all types - tables, safe ranges, length limits)
//		[x] int
//		[x] float
//		[ ] string
//		[ ] uuid (maybe same extension as hashes, ranges?)
//			?;,.!:
//			UUID EXTENSION hash%start<hash%end, on-demand parse
//			to UUIDVector, remember offsets, defaults to `
//			Extension bit!!!
//		[ ] int-in-uuid, float-in-uuid, string-in-uuid
//			float is two ints: value and E? (read IEEE)
//		[ ] Spec -> QuadUUID, 4 values max
//		[ ] optionally, atom vectorization =1<2<3<4
// [x] lww - idempotency
// [x] parse: imply . if no frame header seen previously (is_frame_open)
// [ ] parse: get rid of the "NEXT" hack, check ragel docs
// [ ] sorter: pre-detect errors, split multiframes, etc
// [ ] parser: proper UTF-8 CHAR
//
// [ ] RGA reducer (fn, errors)
//		[x] Reduce()
//		[x] tab tests
//		[ ] benchmark: 1mln ops
//
// [ ] fuzzer go-fuzz (need samples)
// [ ] defensive atom parsing
//
// [ ] reducer registry
// [ ] reducer flags (at least, formatting)
// [x] nice base64 constant definitions (ron ... // "comment")
// [ ] error header   @~~~~~~~~~~:reference "error message" (to reduce)
// [-] copy generic reduction errors
// [x] struct Reducer - mimic Rocks, (a,b) or (a,b,c,d,...)
// [x] prereduce - optional, may fail (RGA subtrees)
// [-] Frame.Split() multiframe parsing  ;,,.,,!,. etc
// [ ] multiframe Sorter
//
// [x] formatting options
// 		[ ] indenting
// 		[x] newlines
// 		[ ] trimming/zipping
// 		[ ] redefs (bench - fast prefix - bit ops)

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

var NO_ATOMS = Atoms{}

type RawUUID []byte

const BASE64 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz~"

var base64 = []byte(BASE64)
var ABC [256]int8

func (op Op) Type() UUID {
	return op.Spec[SPEC_TYPE]
}

func (op Op) Object() UUID {
	return op.Spec[SPEC_OBJECT]
}

func (op Op) Event() UUID {
	return op.Spec[SPEC_EVENT]
}

func (op Op) Reference() UUID {
	return op.Spec[SPEC_REF]
}

const OP_UPDATE_BIT = 1
const OP_STATE_BIT = 2
const OP_QUERY_BIT = 4

// FIXME bracket order to match the numeric order!!!!   }{][)(

const INT60_FULL uint64 = 1<<60 - 1
const INT60_ERROR = INT60_FULL
const INT60_NEVER = 63 << (6 * 9)
const UUID_NAME_UPPER_BITS uint64 = UUID_NAME<<60
const UUID_EVENT_UPPER_BITS uint64 = UUID_EVENT<<60
const UUID_DERIVED_UPPER_BITS uint64 = UUID_DERIVED<<60
const UUID_HASH_UPPER_BITS uint64 = UUID_HASH<<60
const UUID_UPPER_BITS uint64 = 3<<60

var ZERO_UUID = UUID{0, UUID_NAME_UPPER_BITS}

var NEVER_UUID = UUID{INT60_NEVER, UUID_NAME_UPPER_BITS}

var ERROR_UUID = UUID{INT60_ERROR, UUID_NAME_UPPER_BITS}

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
