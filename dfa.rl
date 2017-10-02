package RON

import "fmt"
import "errors"
const trace = false

type parser struct { // TODO
    data []byte
    p, pe, cs int
    ts, te, act int
}

func XParseOp(data []byte, op *Op, context Op) int {
    // TODO phase out pointer-signatures!!!

    %% machine RON;
    %% write data;

    var ctx_uuid UUID = ZERO_UUID
    _ = ctx_uuid
    var uuid *UUID
    var blank UUID
    var i uint64
    var digits uint
    var uuid_value, uuid_origin, uuid_scheme uint64
    var n, old_n int = -1, -1
    var atoms_at, atoms_till int
    var red uint

    op.Count = 0
    op.Body = op.Body[:0]
    if context.Term==TERM_LAST { // default op status
        op.Term = TERM_LAST
    } else {
        op.Term = TERM_INNER
    }

	cs, p, pe, eof := 0, 0, len(data), len(data)
	var ts, te, act int
    _ = eof
    _,_,_ = ts,te,act
    done := false

	%%{

        include OP "./op-grammar.rl";
        main := OP ;

	    write init;
	    write exec;
	}%%

    if done {
        op.Body = data[atoms_at:atoms_till]
        return p-1
    } else {
        return -p
    }
}

var DIGIT_OFFSETS [10]uint8
var PREFIX_MASKS [10]uint64

func init () {
    var one uint64 = 1
    for i:=0; i<10; i++ {
        var bitoff uint8 = uint8(60 - i*6)
        DIGIT_OFFSETS[i] = bitoff - 6
        PREFIX_MASKS[i] = ((one<<60)-1) - ((one<<bitoff)-1)
    }
}


func (ctx_uuid UUID) Parse (data []byte) (ret UUID, err error) {

    %% machine UUID;
    %% write data;

    var i uint64
    var digits uint
    var uuid_value, uuid_origin, uuid_scheme uint64

	cs, p, pe, eof := 0, 0, len(data), len(data)
	var ts, te, act int
    _ = eof
    _,_,_ = ts,te,act


	%%{ 

        include UUID "./uuid-grammar.rl";
        main := UUID ;

	    write init;
	    write exec;
	}%%

    if cs < %%{ write first_final; }%% || digits>10 {
        err = errors.New(fmt.Sprintf("parse error at pos %d", p))
    } else {
        ret = NewUUID(uuid_scheme, uuid_value, uuid_origin)
    }

    return
}

