%%{

    machine UUID;

    action start_uuid {
        uuid_value = ctx_uuid.Value;
        uuid_scheme = ctx_uuid.Scheme()
        uuid_origin = ctx_uuid.Replica();
    }


    action int60_prefix {
        digits = prefixSep2Bits(fc)+4;
        i &= PREFIX_MASKS[digits];
    }

    action int60_digit {
        if digits==0 {
            i = 0;
        } else if digits>9 {
            digits++
            fbreak;
        }
        i |= uint64(baseSep2Bits(fc)) << DIGIT_OFFSETS[digits]
        digits++
    }

    action start_value {
        i = ctx_uuid.Value
        digits = 0
    }

    action start_origin {
        i = ctx_uuid.Replica()
        digits = 0
    }

    action end_value {
        uuid_value = i
    }

    action end_origin {
        uuid_origin = i
    }

    action uuid_sep {
        uuid_scheme = uint64(uuidSep2Bits(fc))
    }

    action end_name {
        uuid_origin = 0;
        uuid_scheme = UUID_NAME;
    }

    BASE = ( [0-9a-zA-Z~_] @int60_digit )+;
    PREFIX =  [([\{\}\])]  @int60_prefix;
    SIGN = [\-+\$%] @uuid_sep;
    PBASE = PREFIX BASE?;
    INT = PBASE | BASE;

    VALUE = INT >start_value %end_value ;
    ORIGIN = INT >start_origin %end_origin ;
    PORIGIN = PBASE >start_origin %end_origin ;
    NAME = BASE >start_value %end_value %end_name;

    UUID =  ( NAME | VALUE ( SIGN ORIGIN? | PORIGIN? ) | SIGN ORIGIN? )
            >start_uuid
           ;

# main := UUID;

}%%

