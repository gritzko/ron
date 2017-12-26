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
	ret.atoms = make([]Atom, 4, DEFAULT_ATOMS_ALLOC)
	ret.Serializer.Format = format
	return
}

func MakeFrame(prealloc_bytes int) (ret Frame) {
	ret.Body = make([]byte, 0, prealloc_bytes)
	ret.atoms = make([]Atom, 4, DEFAULT_ATOMS_ALLOC)
	return
}

func ParseStream (buf []byte) Frame {
	ret := MakeFrame(1000+len(buf))
	ret.AppendBytes(buf)
	ret.Parser.streaming = true
	ret.Next()
	return ret
}

func MakeStream(prealloc_bytes int) (ret Frame) {
	ret.Body = make([]byte, 0, prealloc_bytes)
	ret.atoms = make([]Atom, 4, DEFAULT_ATOMS_ALLOC)
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
    return (frame.Parser.state == RON_start && frame.Position>=0) || frame.Parser.state == RON_FULL_STOP
}

func (ps ParserState) State () int {
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
    return len(batch)>0
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

func NewQuery(t,o,e,r UUID) Frame {
	ret := NewFrame()
	ret.AppendQueryHeader(NewSpec(t,o,e,r))
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

func NewSpec (t,o,e,l UUID) Spec {
	return Spec{Atom(t), Atom(o), Atom(e), Atom(l)}
}

func (frame Frame) Values() []Atom {
    return frame.atoms[4:]
}

func (spec Spec) Type () UUID {
	return UUID(spec[SPEC_TYPE])
}
func (spec Spec) Object () UUID {
	return UUID(spec[SPEC_OBJECT])
}
func (spec Spec) Event () UUID {
	return UUID(spec[SPEC_EVENT])
}
func (spec Spec) Ref () UUID {
	return UUID(spec[SPEC_REF])
}

func (spec Spec) SetType (uuid UUID) {
	spec[SPEC_TYPE] = Atom(uuid)
}
func (spec Spec) SetObject (uuid UUID) {
	spec[SPEC_OBJECT] = Atom(uuid)
}
func (spec Spec) SetEvent (uuid UUID) {
	spec[SPEC_EVENT] = Atom(uuid)
}
func (spec Spec) SetRef (uuid UUID) {
	spec[SPEC_REF] = Atom(uuid)
}

// Verify the syntax, return the offset where error was found. -1 means OK.
func (frame Frame) Verify () int {
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

// When we copy a frame by value, we keep a reference to the slice
// of atom values. Hence, we can't iterate without messing up the
// original frame (fixme).
// Hence, Unshare()
func (frame *Frame) Unshare () {
	newAtoms := make([]Atom, len(frame.atoms))
	copy(newAtoms, frame.atoms)
	frame.atoms = newAtoms
}

func (frame Frame) Equal (other Frame) bool {
	ret := true
	frame.Unshare()
	other.Unshare()
	for ret && !frame.EOF() && !other.EOF() {
		ret = ret && frame.Term()== other.Term()
		for i:=0; i<4 && ret; i++ {
			ret = ret && frame.atoms[i]== other.atoms[i]
		}
		ret = ret && frame.Count()== other.Count()
		// TODO atoms
		frame.Next()
		other.Next()
	}
	ret = ret && frame.EOF()
	ret = ret && other.EOF()
	return ret
}

func (batch Batch) Equal (other Batch) bool {
	if len(batch)!=len(other) {
		return false
	}
	for i:=0; i<len(batch); i++ {
		if !batch[i].Equal(other[i]) {
			return false
		}
	}
	return true
}