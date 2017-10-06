package RON

import "fmt"
import "errors"

%% machine RON;
%% write data;
%% access it.state.;


// Parse consumes one op, unless the buffer ends earlier.
func (it *Frame) Parse() {

    //fmt.Println("GO");

    if it.state.p>=len(it.state.data) {
        if !it.state.streaming {
            it.Op = ZERO_OP
            it.state.cs = RON_error
        }
        return
    }

    if it.state.cs==0 && it.state.p==0 {
	    %% write init;
    } 

    had_end := false 
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

    if !it.state.streaming && it.state.cs<RON_first_final {
        it.state.cs = RON_error
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

