package RON

import "fmt"
import "errors"
const trace = false

/*
 * [ ] use parserState inside XParse
 * [ ] move pS to Iter, XP -> Iter.Next
 * [ ] ron.hpp structure
 * [ ] Cursor, separate atom-grammar.rl
 * [ ] Cursor.Integer()... (same as uuid-..., ragel prepares a slice)
 *
 * */

type OpParserPos struct {
    // int60 idx, base64 digit idx
    idx, half, digit uint
}

type OpParserState struct {
    OpParserPos
    // the RON frame (for the streaming mode, probably a bit less or a bit more)
    data []byte
    // parser position
    p int
    // ragel state
    cs int
    // ts, te, act int
    // incomplete uuid/atom data
    incomplete uint128
    // streaming mode switch
    streaming bool
}

/*func (state *OpParserState) atom_slice () (ret []byte) {
    if len(cur_atom) > 0 {
        return append(cur_atom, data[:p])
    } else {
        return data[atom_start:p]
    }
}*/

const (
        PARSED_ERROR = iota
        PARSED_OP
        PARSED_INCOMPLETE
        PARSED_EOF
      )

// Parse consumes one op, unless the buffer ends earlier.
func (it *Iterator) Parse() int {

    if it.IsLast() {
        it.Op = ZERO_OP
        return PARSED_EOF
    }

    %% machine RON;
    %% write data;
    %% access it.state.;

    if it.state.cs==0 {
        it.Reset()
        it.frame = it.state.data;
        if it.term!=TERM_RAW {
            it.term = TERM_REDUCED
        }
    }

	p, pe, eof := it.state.p, len(it.state.data), len(it.state.data)
    n := uint(0)
    done := false
    _ = done
    _ = eof

    if it.state.streaming {
        eof = -1
    }

    i := it.state.incomplete;
    idx := it.state.idx;
    half := it.state.half;
    digit := it.state.digit;

	%%{

        include FRAME "./op-grammar.rl";
        main := FRAME ;

	    write init;
	    write exec;
	}%%

    it.state.incomplete = i;
    it.state.idx = idx;
    it.state.digit = digit;
    it.state.half = half;
    it.state.p = p;

    if it.state.cs == RON_error {
        return PARSED_ERROR
    } else if it.state.cs >= RON_first_final {
        return PARSED_EOF
    } else if p < pe {
        return PARSED_OP
    } else {
        return PARSED_INCOMPLETE
    }
}

var DIGIT_OFFSETS [11]uint8
var PREFIX_MASKS [11]uint64

func init () {
    var one uint64 = 1
    for i:=0; i<11; i++ {
        var bitoff uint8 = uint8(60 - i*6)
        DIGIT_OFFSETS[i] = bitoff - 6
        PREFIX_MASKS[i] = ((one<<60)-1) - ((one<<bitoff)-1)
    }
}


func (ctx_uuid UUID) Parse (data []byte) (UUID, error) {

    %% machine UUID;
    %% write data;

    var i uint128 = ctx_uuid.uint128
    digit := uint(0)
    half := 0

	cs, p, pe, eof := 0, 0, len(data), len(data)
	//var ts, te, act int
    _ = eof
    //_,_,_ = ts,te,act


	%%{ 

        include UUID "./uuid-grammar.rl";
        main := UUID ;

	    write init;
	    write exec;
	}%%

    if cs < UUID_first_final || digit>10 {
        return ERROR_UUID, errors.New(fmt.Sprintf("parse error at pos %d", p))
    } else {
        return UUID{uint128:i}, nil 
    }

}

