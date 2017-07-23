package RON

import "fmt"
const trace = true

type parser struct {
    data []byte
    p, pe, cs int
    ts, te, act int
}

func ParseOp(data []byte, context Op) (op Op, length int) {
    op = context
    length = XParseOp(data, &op, &context)
    return
}

func XParseOp(data []byte, op *Op, context *Op) int {

    if context==nil {
        context = &ZERO_OP
    }

    %% machine RON;
    %% write data;

    var prev_uuid *UUID = &ZERO_UUID
    var ret int
    var uuid *UUID
    var i uint64
    var digits uint
    var n int
    var prev_uuid_ind int
    var length = -1
    _ = length

	cs, p, pe, eof := 0, 0, len(data), len(data)
	var ts, te, act int
    _ = eof
    _,_,_ = ts,te,act

	%%{

        include OP "./op-grammar.rl";
        main := OP ;

	    write init;
	    write exec;
	}%%

    return ret
}

func ParseUUID(data []byte, context UUID) (uuid UUID, length int) {
    uuid = context
    length = XParseUUID(data, &uuid)
    return
}

// BIG FIXME  ERROR HANDLING, TESTS
func XParseUUID(data []byte, uuid* UUID) (length int) {

    %% machine UUID;
    %% write data;

    var i uint64 = uuid.Value
    var digits uint
    length = -1

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

    // FIXME checkk all input is parsed

    return
}

