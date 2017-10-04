package RON

const INT60LEN = 10
const MAX_ATOMS = 8

type uint128 [2]uint64

type UUID struct {
	uint128
}

type Spec struct {
	uuids [4]UUID
}

type Atoms struct {
	_atoms [2]uint128
	atoms  []uint128
	frame  []byte
}

// OP is an immutable atomic operation object - no write access
type Op struct { // ~128 bytes
	Spec
	Atoms
	term uint
}

// Immutable RON op Frame; the first op is pre-parsed
type Frame struct {
	Op
	state  OpParserState
	Format uint
}

// Checker performs sanity checks on incoming data. Note that a Checker
// may accumulate data, e.g. keep a max timestamp seen.
type Checker interface {
	Check(frame Frame) error
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
// [x] multiframe (still atomic)   Frame.Split(iterator)
//
// cli FIXME
// [ ] clean-up: uuid-grammar.rl (take from Java)
// [x] iterator - parse error
// [ ] value parsing NEW DEAL - [2]uint64
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
//		[-] Spec -> QuadUUID, 4 values max
//		[-] optionally, atom vectorization =1<2<3<4 OR arb number of atoms
//		[ ] Cursor, aka atom iterator
// [x] lww - idempotency
// [x] parse: imply . if no frame header seen previously (is_frame_open)
// [ ] parse: get rid of the "NEXT" hack, check ragel docs
// [x] sorter: pre-detect errors, split multiframes, etc
// [ ] parser: proper UTF-8 CHAR pattern
// [ ] AppendRange, Iterator.offset, IsEmpty()
// [ ] Frame.Header() parsed header fast access
// [ ] Parsers: err = SMOE_ERROR
// [ ] FRAME IS A SLICE
//      [ ] no *Frame
//      [ ] fr = fr.Append(...)
//      [ ] first, last *Op
// [x] MakeNameUUID("name")
//
//
// [ ] RGA reducer (fn, errors)
//		[x] Reduce()
//		[x] tab tests
//		[ ] benchmark: 1mln ops
//
// [ ] fuzzer go-fuzz (need samples)
// [ ] defensive atom parsing
// [ ] LWW: out-of-order entries - restart the algo (with alloc)
// [ ] iheap: seek the loop - reimpl (see UHeap), bench
// [ ] LWW: 1000x1000 array test
//
// ## NEW ORDER ##
// [x] @~! explicit frame terminator - or ;  frame.Close() frame.Join()
// [ ] parser-private adaptor fns  _set_digit()
// [ ] unified grammar files: Java, C++, Go
// [ ] Op: 4 UUIDs, []byte atoms
// [ ] Iterator, ret code, error/incomplete input
// [ ] separate atom parser
// [ ] reader.Next() reader.ReadInt()...
// [ ] ron.Writer
// [ ] Frame, Reader, Writer inherit Op (see C++)
// [ ] type Batch []Frame, type Flow chan Batch
// [ ] auto-gen ABC! (base64: take from the file)
// [ ] Cursor API:  SetObject(uuid), AddInteger(int), Append()
//                  AppendFrame(), AppendAll(), AppendRange()
//
// [ ] reducer registry
// [ ] reducer flags (at least, formatting)
// [x] nice base64 constant definitions (ron ... // "comment")
// [ ] error header   @~~~~~~~~~~:reference "error message" (to reduce)
// [-] copy generic reduction errors
// [x] struct Reducer - mimic Rocks, (a,b) or (a,b,c,d,...)
// [x] prereduce - optional, may fail (RGA subtrees)
// [x] Frame.Split() multiframe parsing  ;,,.,,!,. etc
// [x] multiframe Sorter
// [x] consider ?!,; instead of !.,; and ?
// [x]   insert ; or , depending on the prev op
// [ ] test redefs!
// [ ] test op term defaulting (Append, op before frame, etc)
// [ ] ron.go --> cmd_reduce.go
// [x] go fmt hook
// [ ] test/benchmark hook
// [x] reducers to ignore empty frames
// [ ] Frame.Realloc() // put valuues on a new slab, release old slices
// [x] clock.Authority, clock.See() bool
// [x] ParseUUID sig
// [-] far future: 64 bit uuid, 2bit type, 2bit 1..4 bytes of origin
//
// [x] formatting options
// 		[x] indenting
// 		[x] newlines
// 		[x] trimming/zipping
// 		[ ] redefs (bench - fast prefix - bit ops)
// [x] kill 2 impl of zip UUID
// [x] test formatting
// [ ] test redefaults - BACKTICK ONLY (replaces the quant)

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

type RawUUID []byte

type Environment map[uint64]UUID

const BASE64 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz~"

var ABC [256]int8

// FIXME bracket order to match the numeric order!!!!   }{][)(

const INT60_FULL uint64 = 1<<60 - 1
const INT60_ERROR = INT60_FULL
const INT60_INFINITY = 63 << (6 * 9)
const INT60_FLAGS = uint64(15) << 60

const UUID_NAME_UPPER_BITS = uint64(UUID_NAME) << 60

var ZERO_UUID = NewNameUUID(0, 0)

var NEVER_UUID = NewNameUUID(INT60_INFINITY, 0)

var ERROR_UUID = NewNameUUID(INT60_ERROR, 0)

var ZERO_OP = Op{}

var NO_ATOMS = Atoms{}

func init() {
	// TODO move to bitsep.go
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
	for i := 0; i < len(TERM_PUNCT); i++ {
		ABC[TERM_PUNCT[i]] = -5
	}
}
