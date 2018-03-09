package ron

import (
	"io"
)

func OpenFrame(data []byte) Frame {
	frame := Frame{}
	frame.Body = data
	frame.Parse()
	return frame
}

func MakeFormattedFrame(format uint, prealloc_bytes int) (ret Frame) {
	ret.Body = make([]byte, 0, prealloc_bytes)
	ret.Serializer.Format = format
	return
}

func MakeFrame(prealloc_bytes int) (ret Frame) {
	ret.Body = make([]byte, 0, prealloc_bytes)
	return
}

func ParseStream(buf []byte) Frame {
	ret := MakeFrame(1000 + len(buf))
	ret.AppendBytes(buf)
	ret.Parser.streaming = true
	ret.Next()
	return ret
}

func MakeStream(prealloc_bytes int) (ret Frame) {
	ret.Body = make([]byte, 0, prealloc_bytes)
	ret.Parser.streaming = true
	//ret.Parser.state = RON_start
	return
}

func (frame Frame) Cap() int {
	return cap(frame.Body)
}

func (frame Frame) Spec() Spec {
	var ret [4]Atom
	copy(ret[:], frame.atoms[:4])
	return ret[:]
}

func (frame *Frame) Read(reader io.Reader) (length int, err error) {
	len, cap := frame.Len(), frame.Cap()
	length, err = reader.Read(frame.Body[len:cap])
	if length > 0 {
		frame.Body = frame.Body[:len+length]
	}
	return
}

func (frame Frame) IsEmpty() bool {
	return len(frame.Body) == 0
}

func (frame Frame) IsComment() bool {
	return frame.Type() == COMMENT_UUID
}

func (frame Frame) UUID(idx int) UUID {
	return UUID(frame.atoms[idx])
}

func (frame Frame) Fill(clock Clock, env Environment) Frame {
	ret := MakeFrame(frame.Len() << 1)
	// FIXME header
	for !frame.EOF() {
		spec := frame.Spec()
		if spec.Event().IsZero() {
			spec.SetEvent(clock.Time())
		}
		ret.AppendAmended(spec, frame, frame.Term())
		frame.Next()
	}
	return ret
}

func (frame Frame) Reformat(format uint) Frame {
	ret := MakeFrame(frame.Len())
	ret.Serializer.Format = format
	clone := frame.Clone()
	for !clone.EOF() {
		ret.Append(clone)
		clone.Next()
	}
	return ret.Rewind()
}

func (frame Frame) Clone() (clone Frame) {
	clone = frame
	clone.atoms = make([]Atom, len(frame.atoms))
	copy(clone.atoms, frame.atoms)
	l := len(frame.Body)
	// prevent from appending to the same buffer
	clone.Body = frame.Body[0:l:l]
	return
}

func (frame Frame) String() string {
	return string(frame.Body)
}

func NewBufferFrame(data []byte) (i Frame) {
	i.Body = data
	i.Parse()
	return
}

func NewStringFrame(data string) (i Frame) {
	return NewBufferFrame([]byte(data))
}

func (frame Frame) IsLast() bool {
	return frame.Parser.position >= len(frame.Body)
}

func (frame *Frame) Next() bool {
	frame.Parse()
	if frame.Parser.state == RON_error {
		return false
	}
	if frame.Parser.streaming {
		return frame.IsComplete()
	}
	return true
}

func (frame Frame) Rewind() Frame {
	return ParseFrame(frame.Body)
}

func (frame Frame) Len() int {
	return len(frame.Body)
}

// True if we are past the last op
func (frame Frame) EOF() bool {
	return frame.Parser.state == RON_error
}

func (frame *Frame) SkipHeader() {
	if frame.IsHeader() {
		frame.Next()
	}
}

func (frame Frame) Offset() int {
	return frame.Parser.position
}

// Whether op parsing is complete (not always the case for the streaming mode)
func (frame Frame) IsComplete() bool {
	return (frame.Parser.state == RON_start && frame.Position >= 0) || frame.Parser.state == RON_FULL_STOP
}

func (ps ParserState) State() int {
	return ps.state
}

// Write a frame to a stream (non-trivial because of event mark rewrites)
func (frame Frame) Write(w io.Writer) error {
	_, err := w.Write(frame.Body)
	//fmt.Printf("WROTE: '%s'\n", string(frame.Body))
	return err
}

// Write a batch as a multi-frame
// FIXME merge into a solid frame
func (batch Batch) WriteAll(w io.Writer) (err error) {
	for i := 0; i < len(batch) && err == nil; i++ {
		err = batch[i].Write(w)
	}
	return
}

func (batch Batch) String() (ret string) {
	for _, frame := range batch {
		ret += frame.String() + "\n"
	}
	return
}

func (batch Batch) Len() int {
	ret := 0
	for _, f := range batch {
		ret += len(f.Body)
	}
	return ret
}

func (batch Batch) IsEmpty() bool {
	return len(batch) > 0
}

func (state ParserState) IsFail() bool {
	return state.position > 0 && state.state == RON_error
}

func (batch Batch) HasFullState() bool {
	for _, f := range batch {
		if f.IsFullState() {
			return true
		}
	}
	return false
}

func (frame *Frame) AppendBytes(data []byte) {
	frame.Body = append(frame.Body, data...)
}

func NewFrame() Frame {
	return NewBufferFrame(make([]byte, 0, 1024))
}

func NewQuery(t, o, e, r UUID) Frame {
	ret := NewFrame()
	ret.AppendQueryHeader(NewSpec(t, o, e, r))
	return ret
}

func NewFormattedFrame(format uint) (ret Frame) {
	ret = NewFrame()
	ret.Serializer.Format = format
	return
}

func (frame Frame) Rest() []byte {
	return frame.Body[frame.Parser.position:]
}

// Split returns two frames: one from the start to the current position (exclusive),
// another from the current pos (incl) to the end. The right one is "stripped".
func (frame Frame) Split2() (left, right Frame) {
	// TODO text vs binary
	left = ParseFrame(frame.Body[0:frame.Parser.offset])
	right = NewBufferFrame(make([]byte, 0, 128+frame.Len()-frame.Parser.position))
	right.Append(frame)
	right.AppendBytes(frame.Rest())
	return
}

func (frame Frame) SplitInclusive() Frame {
	at := frame.Parser.position
	if at > 0 && frame.Body[at-1] == FRAME_TERM_SEP {
		at-- // strip the frame terminator
	}
	return ParseFrame(frame.Body[0:at])
}

var FRAME_TERM_ARR = [2]byte{FRAME_TERM_SEP, '\n'}
var DIGIT_OFFSETS = [11]uint8{54, 48, 42, 36, 30, 24, 18, 12, 6, 0, 255}
var PREFIX_MASKS = [11]uint64{0, 1134907106097364992, 1152640029630136320, 1152917106560335872, 1152921435887370240, 1152921503533105152, 1152921504590069760, 1152921504606584832, 1152921504606842880, 1152921504606846912, 1152921504606846975}

func (frame Frame) Integer(i int) int64 {
	return frame.atoms[i+4].Integer()
}

func (frame Frame) Atom(i int) Atom {
	return frame.atoms[i+4]
}

func NewSpec(t, o, e, l UUID) Spec {
	return Spec{Atom(t), Atom(o), Atom(e), Atom(l)}
}

func (frame Frame) Values() []Atom {
	return frame.atoms[4:]
}

func (spec Spec) Type() UUID {
	return UUID(spec[SPEC_TYPE])
}
func (spec Spec) Object() UUID {
	return UUID(spec[SPEC_OBJECT])
}
func (spec Spec) Event() UUID {
	return UUID(spec[SPEC_EVENT])
}
func (spec Spec) Ref() UUID {
	return UUID(spec[SPEC_REF])
}

func (spec Spec) SetType(uuid UUID) {
	spec[SPEC_TYPE] = Atom(uuid)
}
func (spec Spec) SetObject(uuid UUID) {
	spec[SPEC_OBJECT] = Atom(uuid)
}
func (spec Spec) SetEvent(uuid UUID) {
	spec[SPEC_EVENT] = Atom(uuid)
}
func (spec Spec) SetRef(uuid UUID) {
	spec[SPEC_REF] = Atom(uuid)
}

// Verify the syntax, return the offset where error was found. -1 means OK.
func (frame Frame) Verify() int {
	ve := frame.Rewind()
	for !ve.EOF() {
		ve.Next()
	}
	if ve.Offset() != len(ve.Body) {
		return ve.Offset()
	} else {
		return -1
	}
}

// IsEqual checks for single-op equality
func (frame Frame) Compare(other Frame) (eq bool, at int) {
	if frame.EOF() || other.EOF() {
		return frame.EOF() && other.EOF(), 0
	}
	if frame.Term() != other.Term() {
		return false, -1
	}
	for i := 0; i < 4; i++ { // FIXME strings are difficult
		if frame.atoms[i] != other.atoms[i] {
			return false, i
		}
	}
	if frame.Count() != other.Count() {
		return false, -2
	}
	return true, 0
}

func (frame Frame) CompareAll(other Frame) (eq bool, op, at int) {
	for !frame.EOF() && !other.EOF() {
		eq, at = frame.Compare(other)
		if !eq {
			return
		}
		op++
		frame.Next()
		other.Next()
	}
	if !frame.EOF() || !other.EOF() {
		eq = false
		return
	}
	return
}

// Equal checks two batches for op-by-op equality (irrespectively of frame borders)
func (batch Batch) Compare(other Batch) (eq bool, op, at int) {
	bi, oi := 0, 0
	bf := Frame{}
	of := Frame{}
	for (!bf.EOF() || bi < len(batch)) && (!of.EOF() || oi < len(other)) {
		for bf.EOF() && bi < len(batch) {
			bf = batch[bi]
			bi++
		}
		for of.EOF() && oi < len(other) {
			of = other[oi]
			oi++
		}
		eq, at = bf.Compare(of)
		if !eq {
			return
		}
		op++
		bf.Next()
		of.Next()
	}
	if bi != len(batch) || oi != len(other) {
		eq = false
	}
	return
}

func (frame Frame) Equal(other Frame) bool {
	eq, _, _ := frame.CompareAll(other)
	return eq
}

func (batch Batch) Equal(other Batch) bool {
	eq, _, _ := batch.Compare(other)
	return eq
}

func (batch *Batch) AppendFrame(f Frame) {
	*batch = append(*batch, f)
}

func (frame Frame) HasUUIDAt(i int) bool {
	return frame.Count() > i && frame.Atom(i).Type() == ATOM_UUID
}

func (frame Frame) HasIntAt(i int) bool {
	return frame.Count() > i && frame.Atom(i).Type() == ATOM_INT
}

func (frame Frame) GetUUID(i int) UUID {
	if frame.Count() <= i {
		return ZERO_UUID
	}
	atom := frame.Atom(i)
	if atom.Type() != ATOM_UUID {
		return ZERO_UUID
	}
	return atom.UUID()
}

func (frame Frame) GetString(i int) string {
	return frame.RawString(i)
}

func (frame Frame) GetInteger(i int) int64 {
	if frame.Count() <= i {
		return 0
	}
	atom := frame.Atom(i)
	if atom.Type() != ATOM_INT {
		return 0
	}
	return atom.Integer()
}

func (frame Frame) GetInt(i int) int {
	return int(frame.GetInteger(i))
}
