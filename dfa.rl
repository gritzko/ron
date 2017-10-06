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

var SYNTAX_ERROR = NewError("BadSyntax")
var LIMIT_ERROR = NewError("SyntxLimit")
var EOF_ERROR = NewError("EOF")
var INCOMPLETE_ERROR = NewError("Incomplete")

%% machine RON;
%% write data;
%% access it.state.;

// Parse consumes one op, unless the buffer ends earlier.
func (it *Frame) Parse() error {

    //fmt.Println("GO");

    if it.IsLast() {
        it.Op = ZERO_OP
        return EOF_ERROR
    }

    if it.state.cs==0 {
        //fmt.Println("INIT");
	    %% write init;
    } else if it.state.cs>=RON_first_final {
        it.state.cs = RON_start
    }

	p, pe, eof := it.state.p, len(it.state.data), len(it.state.data)
    n := uint(0)
    _ = eof
    _ = pe // FIXME kill

    if it.state.streaming {
        eof = -1
    }

    i := it.state.incomplete;
    idx := it.state.idx;
    half := it.state.half;
    digit := it.state.digit;
    //fmt.Println("GO!", it.state.cs, "at", p, "with", it.state.data[p]);

	%%{

        include FRAME "./op-grammar.rl";
        main := FRAME ;

	    write exec;
	}%%
    //fmt.Println("DONE", it.state.cs, "at", p);

    it.state.incomplete = i;
    it.state.idx = idx;
    it.state.digit = digit;
    it.state.half = half;
    it.state.p = p;

    if p>=pe && !it.state.streaming && it.state.cs!=RON_start && it.state.cs<RON_first_final {
        it.state.cs = RON_error
        //fmt.Println("BAD", p, pe, it.state.cs)
    }

    if it.state.cs == RON_start || it.state.cs>=RON_first_final || p==pe {
        return nil
    } else if it.state.cs == RON_error {
        //fmt.Println("DONE1", p);
        it.Op = ZERO_OP;
        return SYNTAX_ERROR
    } else  {
        //fmt.Println("DONE2", p);
        return INCOMPLETE_ERROR
    }
}

func (frame Frame) EOF () bool {
    return frame.state.cs == RON_error
}

func (frame Frame) Offset () int {
    return frame.state.p
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

