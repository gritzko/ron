package ron

/// tweak abc2go
/// re /(\w+)\s+(\S+)\s+(.*)/_PUNCT $1 $2\n_ENUM $1 $3\n_SEPS $1 $2 $3\n/
/// fn /_PUNCT (\w+) (.*)/ (s,enm,seps) => { return "var "+enm+"_PUNCT = []byte(\""+seps.replace(/\\/,"\\\\")+'")' }
/// fn /_ENUM (\w+) (.*)/ (s,enm,vals)=>{ return "const (\n" + vals.split(/\s+/).map(name=>'\t'+enm+"_"+name+" = iota").join('\n') + "\n)" }
/// fn /_SEPS (\w+) (\S+) (.*)/ (s,enm,sepstr,names) => { seps=sepstr.match(/./g).reverse(); return names.split(/\s+/g).map(name=>"const "+enm+"_"+name+"_SEP = \'"+seps.pop().replace(/([\\'])/,"\\$1")+"'").join('\n') }
/// end

/// paste ABC [0755330a]
/// use abc2go [59a4c136]
var SPEC_PUNCT = []byte("*#@:")

const (
	SPEC_TYPE   = iota
	SPEC_OBJECT = iota
	SPEC_EVENT  = iota
	SPEC_REF    = iota
)
const SPEC_TYPE_SEP = '*'
const SPEC_OBJECT_SEP = '#'
const SPEC_EVENT_SEP = '@'
const SPEC_REF_SEP = ':'

var UUID_PUNCT = []byte("$%+-")

const (
	UUID_NAME    = iota
	UUID_HASH    = iota
	UUID_EVENT   = iota
	UUID_DERIVED = iota
)
const UUID_NAME_SEP = '$'
const UUID_HASH_SEP = '%'
const UUID_EVENT_SEP = '+'
const UUID_DERIVED_SEP = '-'

var ATOM_PUNCT = []byte(">='^")

const (
	ATOM_UUID   = iota
	ATOM_INT    = iota
	ATOM_STRING = iota
	ATOM_FLOAT  = iota
)
const ATOM_UUID_SEP = '>'
const ATOM_INT_SEP = '='
const ATOM_STRING_SEP = '\''
const ATOM_FLOAT_SEP = '^'

var TERM_PUNCT = []byte(";,!?")

const (
	TERM_RAW     = iota
	TERM_REDUCED = iota
	TERM_HEADER  = iota
	TERM_QUERY   = iota
)
const TERM_RAW_SEP = ';'
const TERM_REDUCED_SEP = ','
const TERM_HEADER_SEP = '!'
const TERM_QUERY_SEP = '?'

var REDEF_PUNCT = []byte("`\\|/")

const (
	REDEF_PREV   = iota
	REDEF_OBJECT = iota
	REDEF_EVENT  = iota
	REDEF_REF    = iota
)
const REDEF_PREV_SEP = '`'
const REDEF_OBJECT_SEP = '\\'
const REDEF_EVENT_SEP = '|'
const REDEF_REF_SEP = '/'

var PREFIX_PUNCT = []byte("([{}])")

const (
	PREFIX_PRE4 = iota
	PREFIX_PRE5 = iota
	PREFIX_PRE6 = iota
	PREFIX_PRE7 = iota
	PREFIX_PRE8 = iota
	PREFIX_PRE9 = iota
)
const PREFIX_PRE4_SEP = '('
const PREFIX_PRE5_SEP = '['
const PREFIX_PRE6_SEP = '{'
const PREFIX_PRE7_SEP = '}'
const PREFIX_PRE8_SEP = ']'
const PREFIX_PRE9_SEP = ')'

var BASE_PUNCT = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz~")

const (
	BASE_0  = iota
	BASE_1  = iota
	BASE_2  = iota
	BASE_3  = iota
	BASE_4  = iota
	BASE_5  = iota
	BASE_6  = iota
	BASE_7  = iota
	BASE_8  = iota
	BASE_9  = iota
	BASE_10 = iota
	BASE_11 = iota
	BASE_12 = iota
	BASE_13 = iota
	BASE_14 = iota
	BASE_15 = iota
	BASE_16 = iota
	BASE_17 = iota
	BASE_18 = iota
	BASE_19 = iota
	BASE_20 = iota
	BASE_21 = iota
	BASE_22 = iota
	BASE_23 = iota
	BASE_24 = iota
	BASE_25 = iota
	BASE_26 = iota
	BASE_27 = iota
	BASE_28 = iota
	BASE_29 = iota
	BASE_30 = iota
	BASE_31 = iota
	BASE_32 = iota
	BASE_33 = iota
	BASE_34 = iota
	BASE_35 = iota
	BASE_36 = iota
	BASE_37 = iota
	BASE_38 = iota
	BASE_39 = iota
	BASE_40 = iota
	BASE_41 = iota
	BASE_42 = iota
	BASE_43 = iota
	BASE_44 = iota
	BASE_45 = iota
	BASE_46 = iota
	BASE_47 = iota
	BASE_48 = iota
	BASE_49 = iota
	BASE_50 = iota
	BASE_51 = iota
	BASE_52 = iota
	BASE_53 = iota
	BASE_54 = iota
	BASE_55 = iota
	BASE_56 = iota
	BASE_57 = iota
	BASE_58 = iota
	BASE_59 = iota
	BASE_60 = iota
	BASE_61 = iota
	BASE_62 = iota
	BASE_63 = iota
)
const BASE_0_SEP = '0'
const BASE_1_SEP = '1'
const BASE_2_SEP = '2'
const BASE_3_SEP = '3'
const BASE_4_SEP = '4'
const BASE_5_SEP = '5'
const BASE_6_SEP = '6'
const BASE_7_SEP = '7'
const BASE_8_SEP = '8'
const BASE_9_SEP = '9'
const BASE_10_SEP = 'A'
const BASE_11_SEP = 'B'
const BASE_12_SEP = 'C'
const BASE_13_SEP = 'D'
const BASE_14_SEP = 'E'
const BASE_15_SEP = 'F'
const BASE_16_SEP = 'G'
const BASE_17_SEP = 'H'
const BASE_18_SEP = 'I'
const BASE_19_SEP = 'J'
const BASE_20_SEP = 'K'
const BASE_21_SEP = 'L'
const BASE_22_SEP = 'M'
const BASE_23_SEP = 'N'
const BASE_24_SEP = 'O'
const BASE_25_SEP = 'P'
const BASE_26_SEP = 'Q'
const BASE_27_SEP = 'R'
const BASE_28_SEP = 'S'
const BASE_29_SEP = 'T'
const BASE_30_SEP = 'U'
const BASE_31_SEP = 'V'
const BASE_32_SEP = 'W'
const BASE_33_SEP = 'X'
const BASE_34_SEP = 'Y'
const BASE_35_SEP = 'Z'
const BASE_36_SEP = '_'
const BASE_37_SEP = 'a'
const BASE_38_SEP = 'b'
const BASE_39_SEP = 'c'
const BASE_40_SEP = 'd'
const BASE_41_SEP = 'e'
const BASE_42_SEP = 'f'
const BASE_43_SEP = 'g'
const BASE_44_SEP = 'h'
const BASE_45_SEP = 'i'
const BASE_46_SEP = 'j'
const BASE_47_SEP = 'k'
const BASE_48_SEP = 'l'
const BASE_49_SEP = 'm'
const BASE_50_SEP = 'n'
const BASE_51_SEP = 'o'
const BASE_52_SEP = 'p'
const BASE_53_SEP = 'q'
const BASE_54_SEP = 'r'
const BASE_55_SEP = 's'
const BASE_56_SEP = 't'
const BASE_57_SEP = 'u'
const BASE_58_SEP = 'v'
const BASE_59_SEP = 'w'
const BASE_60_SEP = 'x'
const BASE_61_SEP = 'y'
const BASE_62_SEP = 'z'
const BASE_63_SEP = '~'

var FRAME_PUNCT = []byte(".")

const (
	FRAME_TERM = iota
)
const FRAME_TERM_SEP = '.'

/// end

/// tweak var2go
/// fn /(\S+)\s+(.*)/ (s,key,seq) => { i=0; return seq.split(/\s+/g).map(vrt=>{nm=key+"_"+vrt; return "const "+nm+" = "+i+"\nvar "+nm+"_SEP = BASE_PUNCT["+(i++)+"]"}).join("\n") }
/// end

/// paste RON_UUID [368fd2fd]
/// use var2go [191338df]
const UUID_NAME_TRANSCENDENT = 0

var UUID_NAME_TRANSCENDENT_SEP = BASE_PUNCT[0]

const UUID_NAME_ISBN = 1

var UUID_NAME_ISBN_SEP = BASE_PUNCT[1]

const UUID_EVENT_CALENDAR_RECORD = 0

var UUID_EVENT_CALENDAR_RECORD_SEP = BASE_PUNCT[0]

const UUID_EVENT_LOGICAL_RECORD = 1

var UUID_EVENT_LOGICAL_RECORD_SEP = BASE_PUNCT[1]

const UUID_EVENT_EPOCH_RECORD = 2

var UUID_EVENT_EPOCH_RECORD_SEP = BASE_PUNCT[2]

/// end

var ABC [128]uint8
var IS_BASE [4]uint64

func init() {
	for i := 0; i < 128; i++ {
		ABC[i] = 255
	}
	for i, l := range BASE_PUNCT {
		li := uint(l)
		ABC[li] = uint8(i)
		IS_BASE[li>>6] |= uint64(1) << (li & 63)
	}
	for i, l := range PREFIX_PUNCT {
		ABC[l] = uint8(i)
	}
	for i, l := range TERM_PUNCT {
		ABC[l] = uint8(i)
	}
	for i, l := range UUID_PUNCT {
		ABC[l] = uint8(i)
	}
	for i, l := range SPEC_PUNCT {
		ABC[l] = uint8(i)
	}
}
