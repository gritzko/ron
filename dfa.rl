package RON

import "fmt"
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

    var prev_uuid *UUID = &ZERO_UUID
    _ = prev_uuid
    var ret int
    var uuid *UUID
    var blank UUID
    var i uint64
    var digits uint
    n, old_n := -1, -1
    var length = -1
    _ = length
    var atoms_at int

    op.Count = 0
    op.Body = op.Body[:0]
    op.Flags = 0

	cs, p, pe, eof := 0, 0, len(data), len(data)
	var ts, te, act int
    _ = eof
    _,_,_ = ts,te,act
    var bare, full bool
    var sign uint = 0
    done := false

	%%{

        include OP "./op-grammar.rl";
        main := OP ;

	    write init;
	    write exec;
	}%%

    if ret>0 {
        if trace {
            fmt.Printf("ATOMS: %d..%d\n", op.Offsets[0], ret);
        }
    }

    if done {
        return p
    } else {
        return -p
    }
}


// BIG FIXME  ERROR HANDLING, TESTS
// FIXME context UUID
func XParseUUID(data []byte, uuid *UUID) (length int) {

    %% machine UUID;
    %% write data;

    var i uint64
    var digits uint
    length = -1

	cs, p, pe, eof := 0, 0, len(data), len(data)
	var ts, te, act int
    _ = eof
    _,_,_ = ts,te,act
    var bare, full bool
    var sign uint = 0


	%%{ 

        include UUID "./uuid-grammar.rl";
        main := UUID ;

	    write init;
	    write exec;
	}%%

    // FIXME checkk all input is parsed

    return
}

