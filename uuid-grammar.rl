%%{

    machine UUID;


    action int60_prefix {
        digits = uint(-ABC[fc]-10+4) * 6
        i >>= (60-digits)  // FIXME
    }

    action int60_digit {
        if digits==0 {
            i = 0
            full = true
        } else if digits>=60 {
            length = -p
            fbreak;
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
        i = uuid.Replica()
    }

    action origin {
        if digits>0 {
            uuid.Origin = i << (60-digits)
        }
        bare = false
    }

    action uuid_sep {
        sign = uuidSep2Bits(fc)
        i = uuid.Replica()
        uuid.Origin &= PREFIX10
        bare = false
    }

    action uuid {
        length = pe
        uuid.Origin |= uint64(sign) << 60
        if bare && full {
            uuid.Origin = 0
        }
    }

    action start_uuid {
        bare, full = true, false
        i = uuid.Value
        sign = uuid.Scheme()
    }


    BASE = ( [0-9a-zA-Z~_] @int60_digit )+;
    PREFIX =  [([\{\}\])]  @int60_prefix;
    SIGN = [\-+\$%] @uuid_sep;
    PBASE = PREFIX BASE?;
    INT = PBASE | BASE;

    VALUE = INT %value ;
    ORIGIN = INT %origin ;

    UUID =  ( VALUE SIGN ORIGIN? | SIGN ORIGIN? | VALUE (PBASE %origin)? )
            >start_uuid %uuid
           ;

# main := UUID;

}%%

