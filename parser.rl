package RON


type parser struct {
    data []byte
    p, pe, cs int
    ts, te, act int
}

//func ParseOp(data []byte, op *Op) int {
//    return -1
//}

func ParseOp(data []byte, op *Op) int {

    %% machine RON;
    %% write data;

    var uuid UUID
    var ret int
    var context UUID
    uuid = context
    var i uint64 = context.Value
    var digits uint
    var length = -1
    _ = length

	cs, p, pe, eof := 0, 0, len(data), len(data)
	var ts, te, act int
    _ = eof
    _,_,_ = ts,te,act

	%%{ 

        action uuid_redef {
        }

        action type_start {
            //uuid = zero
        }
        action object_start {
            //uuid = zero
        }
        action event_start {
            //uuid = zero
        }
        action location_start {
            //uuid = zero
        }
        action type {
        }
        action object {
        }
        action event {
        }
        action location {
        }

        action atom_start {
            ////fmt.Printf("ATOM_START %c\n", fc);
        }
        action atom {
            ////fmt.Printf("ATOM\n");
        }

        action atoms {
            ////fmt.Printf("VALUE\n");
            ret = p
        }

        include OP "./op-grammar.rl";
        main := OP ;

	    write init;
	    write exec;
	}%%

        return ret
}



// BIG FIXME  ERROR HANDLING, TESTS
func ParseUUID(data []byte, context UUID) (uuid UUID, length int) {

    %% machine UUID;
    %% write data;

    uuid = context
    var i uint64 = context.Value
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

