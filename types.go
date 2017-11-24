package ron

const INT60LEN = 10
const MAX_ATOMS_PER_OP = 1 << 20

type Atom [2]uint64

type UUID Atom

type Spec struct {
	RDType, Object, Event, Ref UUID
}

type ParserState struct {
	// position in the atom array, in the atom, in the half-atom
	atm, hlf, dgt int
	// ragel parser state
	state int
	// byte offset of the current op
	offset int
	// parsing byte offset
	position int
	// whether the frame might get more data
	streaming bool
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
	// The current position in the frame (op idx).
	Position int
	// ints hosts the current op: 4 pairs for spec uuid entries, the rest is values (also pairs).
	// General convention: hte first int is hte value, the second is flags and other stuff.
	atoms []Atom
	// Op terminator (see OP_TERM)
	term int
	// Frame body, raw bytes.
	Body []byte
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
// [x] clean-up: uuid-grammar.rl (take from Java)
// [x] iterator - parse error
// [x] value parsing NEW DEAL - [2]uint64
//		[ ] (all types - tables, safe ranges, length limits)
//		[x] int
//		[x] float
//		[x] string
//		[x] uuid (maybe same extension as hashes, ranges?)
//			?;,.!:
//			UUID EXTENSION hash%start<hash%end, on-demand parse
//			to UUIDVector, remember offsets, defaults to `
//			Extension bit!!!
//		[x] int-in-uuid, float-in-uuid, string-in-uuid
//			float is two ints: value and E? (read IEEE)
//		[-] Spec -> QuadUUID, 4 values max
//		[-] optionally, atom vectorization =1<2<3<4 OR arb number of atoms
//		[-] Cursor, aka atom iterator
// [x] lww - idempotency
// [x] parse: imply . if no frame header seen previously (is_frame_open)
// [x] parse: get rid of the "NEXT" hack, check ragel docs
// [x] sorter: pre-detect errors, split multiframes, etc
// [ ] parser: proper UTF-8 CHAR pattern
// [ ] AppendRange, Iterator.offset, IsEmpty()
// [x] Frame.Header() parsed header fast access
// [x] Parsers: err = SMOE_ERROR
// [x] FRAME IS A SLICE
//      [x] no *Frame
//      [x] fr = fr.Append(...)
//      [x] first, last *Op
// [x] MakeNameUUID("name")
//
// [x] continuation test *a!*b=1*c=1!*d,*e., clean cs states, fhold
// [x] frame.State() OP PART ERROR
//     for frame:=ParseFrame(); !frame.IsEmpty(); frame.Next() {}
// [ ] trailing space test, Rest(), multiframes
//
// [ ] RGA reducer (fn, errors)
//		[x] Reduce()
//		[x] tab tests
//		[ ] benchmark: 1mln ops
//
// [ ] fuzzer go-fuzz (need samples)
// [ ] defensive atom parsing
// [ ] LWW: out-of-order entries - restart the algo (with alloc)
// [ ] iheap: seek the loop - reimpl (see UUIDHeap), bench
// [ ] LWW: 1000x1000 array test
//
// ## NEW ORDER ##
// [x] @~! explicit frame terminator - or ;  frame.Close() frame.Join()
// [-] parser-private adaptor fns  _set_digit()
// [ ] unified grammar files: Java, C++, Go
// [-] Op: 4 UUIDs, []byte atoms
// [x] Iterator, ret code, error/incomplete input
// [-] separate atom parser
// [x] reader.Next() reader.ReadInt()...
// [-] ron.Writer
// [x] Frame, Reader, Writer inherit Op (see C++)
// [x] type Batch []Frame, type Flow chan Batch
// [ ] auto-gen ABC! (base64: take from the file)
// [x] Cursor API:  SetObject(uuid), AddInteger(int), Append()
//                  AppendFrame(), AppendAll(), AppendRange()
// [ ] Nice sigs, frame.read.stream, frame.write.format
// [ ] No Rewind(), just Clone()
//
// [ ] Minimize copying in Frame.Parse()
// [ ] clonable Frames (by value)
// [ ] Atom parser/iterator
// [ ] frame splitting ("op" and "rest", no reset?, more tests)
//
// [ ] AppendXXX(t,o,e,r) - Spec... spread sign
//
// [x] reducer registry
// [x] reducer flags (at least, formatting)
// [x] nice base64 constant definitions (ron ... // "comment")
// [-] error header   @~~~~~~~~~~:reference "error message" (to reduce)
// [-] copy generic reduction errors
// [x] struct Reducer - mimic Rocks, (a,b) or (a,b,c,d,...)
// [x] prereduce - optional, may fail (RGA subtrees)
// [x] Frame.Split() multiframe parsing  ;,,.,,!,. etc
// [x] multiframe Sorter
// [x] consider ?!,; instead of !.,; and ?
// [x]   insert ; or , depending on the prev op
// [ ] test redefs!
// [x] test op term defaulting (Append, op before frame, etc)
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
//		[ ] tablist formatting!!!
// [x] kill 2 impl of zip UUID
// [x] test formatting
// [ ] test redefaults - BACKTICK ONLY (replaces the quant)

// [ ] strings: either escaped byte buffer or an unescaped string!!!!!!

// Reducer is essentially a replicated data type.
// A reduction of the object's full op log produces its RON state.
// A reduction of a log segment produces a patch.
// A reduced frame has the same object id and, in most cases, type.
// Event id is the one of the last input frame.
// Complexity guarantees: max O(log N)
// (could be made to reduce 1mln single-op frames)
// Associative, commutative*, idempotent.
type Reducer interface {
	Reduce(batch Batch) Frame
}

type ReducerMaker func () Reducer

var RDTYPES map[UUID]ReducerMaker

type Batch []Frame

type RawUUID []byte

type Environment map[uint64]UUID

var BASE64 = string(BASE_PUNCT)

const INT60_FULL uint64 = 1<<60 - 1
const INT60_ERROR = INT60_FULL
const INT60_INFINITY = 63 << (6 * 9)
const INT60_FLAGS = uint64(15) << 60

const UUID_NAME_UPPER_BITS = uint64(UUID_NAME) << 60

var ZERO_UUID = NewNameUUID(0, 0)

var NEVER_UUID = NewNameUUID(INT60_INFINITY, 0)

var ERROR_UUID = NewNameUUID(INT60_ERROR, 0)

func init() {

	RDTYPES = make(map[UUID]ReducerMaker, 10)

}
