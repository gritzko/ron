%%{

    machine UUID;

    action start_uuid {
    }


    action int60_prefix {
        dgt = int(ABC[fc])+4;
        atoms[atm][hlf] &= INT60_FLAGS | PREFIX_MASKS[dgt];
    }

    action int60_dgt {
        atoms[atm][hlf] |= ((uint64)(ABC[fc])) << DIGIT_OFFSETS[dgt];
        dgt++;
        if (dgt>10) {
            fbreak;
        }
    }

    action start_full_int {
        atoms[atm][hlf] &= INT60_FLAGS;
    }

    action start_value {
    }

    action start_origin {
        dgt = 0;
        hlf = 1;
    }

    action end_value {
    }

    action end_origin {
    }

    action variety {
        atoms[atm][hlf] <<= 6;
        dgt--;
    }

    action uuid_sep {
        hlf = 1;
        atoms[atm][1] &= INT60_FULL;
        atoms[atm][1] |= ((uint64)(ABC[fc]))<<60;
    }

    action end_name {
        atoms[atm][1] = UUID_NAME_FLAG;
    }

    # Base64 value
    DGT = [0-9a-zA-Z~_] @int60_dgt;
    BASE = DGT+;
    # prefix compression 
    PREFIX =  [([\{\}\])]  @int60_prefix;
    # UUID type: name, hash, event or derived event 
    SIGN = [\$\%\+\-] @uuid_sep;
    # prefix-compressed int (half of UUID) 
    PBASE = PREFIX BASE?;
    # full int 
    FBASE = (DGT ([\/] @variety)? DGT*) >start_full_int;
    # int, either compressed or not 
    INT = PBASE | FBASE;

    # first half of an UUID 
    VALUE = INT >start_value %end_value ;
    # second half of an UUID 
    ORIGIN = INT >start_origin %end_origin ;
    # prefix-compressed 2nd half 
    PORIGIN = PBASE >start_origin %end_origin ;
    # global name UUID, e.g. "lww" (aka transcendent constant) 
    NAME = FBASE >start_value %end_value %end_name;

    # RON 128 bit UUID 
    UUID =  ( NAME | VALUE ( SIGN ORIGIN? | PORIGIN? ) | SIGN ORIGIN? )
            >start_uuid
           ;

    # main := UUID;

}%%

