%%{

    machine UUID;

    action start_uuid {
        //fmt.Println("START");
        half = 0;
    }


    action int60_prefix {
        digit = prefixSep2Bits(fc)+4;
        i[half] &= INT60_FLAGS | PREFIX_MASKS[digit];
    }

    action int60_digit {
        i[half] |= uint64(ABC[fc]) << DIGIT_OFFSETS[digit]
        digit++
        if (digit>10) {
            fbreak;
        }
    }

    action start_full_int {
        i[half] &= INT60_FLAGS;
    }

    action start_value {
    }

    action start_origin {
        digit = 0;
        half |= 1;
    }

    action end_value {
    }

    action end_origin {
    }

    action uuid_sep {
        half |= 1;
        i[half] &= INT60_FULL;
        i[half] |= uint64(uuidSep2Bits(fc))<<60;
    }

    action end_name {
        i[1] = UUID_NAME_UPPER_BITS;
    }

    # Base64 value
    BASE = ( [0-9a-zA-Z~_] @int60_digit )+;
    # prefix compression 
    PREFIX =  [([\{\}\])]  @int60_prefix;
    # UUID type: name, hash, event or derived event 
    SIGN = [\$\%\+\-] @uuid_sep;
    # prefix-compressed int (half of UUID) 
    PBASE = PREFIX BASE?;
    # full int 
    FBASE = BASE >start_full_int;
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

