%%{

    machine UUID;


    action int60_prefix {
        digits = uint(-ABC[fc]-10+4) * 6
        i >>= (60-digits)  // FIXME
        if trace {
            fmt.Printf("PREFIX %c %d %d\n", fc, i, digits)
        }
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
        if trace {
            fmt.Printf("DIGIT %c\n", fc)
        }
    }

    action value {
        if digits>0 {
            uuid.Value = i << (60-digits)
            digits = 0
        }
        i = uuid.Origin
        if trace {
            fmt.Printf("VALUE\n")
        }
    }

    action origin {
        if digits>0 {
            uuid.Origin = i << (60-digits)
        }
        bare = false
        if trace {
            fmt.Printf("ORIGIN\n")
        }
    }

    action uuid_sep {
        uuid.Sign = fc
        i = uuid.Origin
    }

    action uuid {
        length = pe
        if bare && full {
            uuid.Origin = 0
            uuid.Sign = '$'
        }
        if trace {
            fmt.Printf("UUID\n")
        }
    }

    action start_uuid {
        bare, full = true, false
        if trace {
            fmt.Printf("START_UUID\n")
        }
        i = uuid.Value
    }


    BASE = ( [0-9a-zA-Z~_] @int60_digit )+;
    PREFIX =  [([\{\}\])]  @int60_prefix;
    SIGN = [\-+$%] @uuid_sep;

    VALUE = ( PREFIX | BASE | PREFIX BASE ) %value ;
    ORIGIN = ( ( SIGN | PREFIX | SIGN PREFIX )  BASE? ) %origin ;

    UUID =  (VALUE? ORIGIN?)
            >start_uuid %uuid
           ;

# main := UUID;

}%%

