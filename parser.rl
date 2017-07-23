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
    var _, int60_off int
    var ret int

	cs, p, pe, eof := 0, 0, len(data), len(data)
	var ts, te, act int
    _ = eof
    _,_,_ = ts,te,act

	%%{ 

        action int60_prefix      { 
            ////fmt.Printf("INT60_PREFIX: %c\n", fc);
            int60_off = 4
        }
        action int60_digit {
            ////fmt.Printf("INT60_DIGIT: %c\n", fc);
            //uuid[_+int60_off] = fc
            uuid.Value = (uuid.Value<<6) | uint64(ABC[fc])
            int60_off++
        }
        action int60 {
            ////fmt.Printf("INT60: %c\n", fc);
            // FIXME eats cycles
            for ; int60_off<10; int60_off++ {
               // uuid[_+int60_off] = '0'
               uuid.Value <<= 6
            }
            int60_off=0
            //PIN |= uuid.Value
            // FIXME length limit checks
        }

        action origin {

        }
        action uuid_redef {
        }
        action uuid_value {
            ////fmt.Printf("UUID_VALUE\n");
        }
        action uuid_sep {
            uuid.Sign = fc;
            _ = 11;
        }
        action uuid_origin {
        }
        action uuid     {
            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        }

        action spec_uuid {
            ////fmt.Printf("SPEC_UUID %c\n", fc);
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

        action value {
            ////fmt.Printf("VALUE\n");
            ret = p
        }

        action dgt {
            ////fmt.Printf("DIGIT %c\n", fc);
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

        action int60_prefix {
            digits = uint(-ABC[fc]-10+4) * 6
            i >>= (60-digits)  // FIXME
        }

        action int60_digit {
            if digits==0 {
                i = 0
            }
            digits+=6
            i <<= 6
            i |= uint64(ABC[fc])
        }

        action value {
            if digits>0 {
                uuid.Value = i << (60-digits)
                digits = 0
            }
            i = context.Origin
        }

        action origin {
            if digits>0 {
                uuid.Origin = i << (60-digits)
            }
        }

        action uuid_sep {
            uuid.Sign = fc
            i = context.Origin
        }

        action uuid {
            length = pe
        }

        include UUID "./uuid-grammar.rl";
        main := UUID ;

	    write init;
	    write exec;
	}%%

    // FIXME checkk all input is parsed

    return
}

