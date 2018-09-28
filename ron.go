package ron

const INT60LEN = 10
const MAX_ATOMS_PER_OP = 1 << 20

// An atom is a constant of a RON type: int, float, string or UUID
type Atom [2]uint64

// half of an atom
type Half uint8

const (
	VALUE  Half = 0
	ORIGIN Half = 1
)

type UUID Atom

type Atoms []Atom

type Spec []Atom

type ParserState struct {
	// pos in the atom array, in the atom, in the half-atom
	atm, dgt int
	hlf      Half
	// ragel parser state
	state int
	// byte off of the current op
	off int
	// parsing byte off
	pos int
	// parser mode: streaming (might get more bytes) / block (complete frame)
	streaming bool
	// which spec uuids are omitted/defaults in the current op
	omitted uint8
}

type SerializerState struct {
	Format uint
}

// RON Frame is a vector of immutable RON ops.
// A frame is always positioned on some op (initially, the first one).
// In a sense, Frame is its own iterator: frame.Next(), returns true is the
// frame is re-positioned to the next op, false on error (EOF is an error too).
// That is made to minimize boilerplate as Frames are forwarded based on the
// frame header (the first op).
// Frame is not thread-safe; the underlying buffer is append-only, thus thread-safe.
type Frame struct {
	Parser     ParserState
	Serializer SerializerState
	// RON coding: binary/text
	binary bool
	// The current pos in the frame (op idx).
	position int
	// ints hosts the current op: 4 pairs for spec uuid entries, the rest is values (also pairs).
	// General convention: hte first int is hte value, the second is flags and other stuff.
	_atoms [DEFAULT_ATOMS_ALLOC]Atom
	atoms  []Atom
	// Op terminator (see OP_TERM)
	term       int
	descriptor Atom
	// Frame body, raw bytes.
	Body []byte
}

// [ ] type Op { term int  atoms Atoms }
// [ ] the big divorce: separate Frame/Append() and Cursor/Next()

// Checker performs sanity checks on incoming data. Note that a Checker
// may accumulate data, e.g. keep a max timestamp seen.
type Checker interface {
	Check(frame Frame) (err UUID)
}

// [ ] parser: proper UTF-8 CHAR pattern
//
// [ ] tough counter (T-Counter)
// [ ] op-based counter
//
// [ ] consistent header.Event(): max, ++ ?
//
// [x] sweeten the grammar (@1@2@3), last_at
// [ ] *0 fast-append
// [ ] slice-append: remember prevFrame, prevPos
//
// [x] header to zero defaults. less smartness!
//
// [ ] multilevel-logical clock modeAA
//
// [ ] lww json mapper (strings only)
// [x] vim syn file: string json escaping
//
// [x] non-idiomatic Frame behavior (copy->shared atoms array)
//     fix:  parser uses cur *Atom, _atoms [6]Atom atoms []Atom
//
// [x] make rewinds *very* explicit, test (query/header differs?)
//
// [ ] fuzzer go-fuzz (need samples)
// [ ] defensive atom parsing
// [ ] LWW: out-of-order entries - restart the algo (with alloc)
// [ ] iheap: seek the loop - reimpl (see UUIDHeap), bench
// [ ] LWW: 1000x1000 array test
// [ ] ACI property tests for everything
//
// [ ] ron.go --> cmd_reduce.go
// [ ] strings: either escaped byte buffer or an unescaped string!!!!!!
//
// [ ] explicit monoframe lengths for trusted sources
//		*~ :blen =12345!  *~ :flen =1234!
// [ ] iheap strip comments
//

// Reducer is essentially a replicated data type.
// A reduction of the object's full op log produces its RON state.
// A reduction of a log segment produces a patch.
// A reduced frame has the same object id and, in most cases, type.
// Event id is the one of the last input frame.
// Complexity guarantees: max O(log N)
// (could be made to reduce 1mln single-op frames)
// Associative, commutative*, idempotent.
type Reducer interface {
	Features() int // see ACID
	Reduce(batch Batch) Frame
}

type Mapper interface {
	Map(batch Batch) interface{}
}

// Compare two ops; >1 for a>b, <1 for a<b, 0 for equal.
type Comparator func(a, b *Frame) int64

type StringMapper interface {
	Map(batch Batch) string
}

type ReducerMaker func() Reducer // FIXME params map[UUID]Atom - no free-form strings

var RDTYPES map[UUID]ReducerMaker

type Batch []Frame

type RawUUID []byte

type VVector map[uint64]uint64

type Environment map[uint64]UUID

var BASE64 = string(BASE_PUNCT)

const INT60_FULL uint64 = 1<<60 - 1
const INT60_ERROR = INT60_FULL
const INT60_INFINITY = 63 << (6 * 9)
const INT60_FLAGS = uint64(15) << 60

const UUID_NAME_UPPER_BITS = uint64(UUID_NAME) << 60

var ZERO_UUID = NewNameUUID(0, 0)
var ZERO_UUID_ATOM = Atom(ZERO_UUID)

var NEVER_UUID = NewNameUUID(INT60_INFINITY, 0)

var COMMENT_UUID = NEVER_UUID

var ERROR_UUID = NewNameUUID(INT60_ERROR, 0)

const DEFAULT_ATOMS_ALLOC = 6

func init() {

	RDTYPES = make(map[UUID]ReducerMaker, 10)

}
