%%{

    machine UUID;

    BASE = ( [0-9a-zA-Z~_] @int60_digit )+;
    PREFIX =  [([\{\}\])]  @int60_prefix;
    SIGN = [\-+$%] @uuid_sep;

    VALUE = ( PREFIX | BASE | PREFIX BASE ) %value ;
    ORIGIN = ( ( SIGN | PREFIX | SIGN PREFIX )  BASE? ) %origin ;

    UUID =  VALUE? ORIGIN?
            %uuid
           ;

}%%

