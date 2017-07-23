%%{

    machine OP;

    BASE = [0-9a-zA-Z~_];

    INT60U = BASE{1,10}   $int60_digit ;

    INT60Z = (
                [([\{\}\])]          $int60_prefix
                BASE{,10}   $int60_digit
            );

    INT60 = ( INT60U | INT60Z ) %int60 ;

    UUID = (
            INT60
            ( [\-+$%] @uuid_sep INT60 | INT60Z %int60 )?
            ) %uuid
        ;

    #        ( [`\\|/]        %uuid_redef )?  

    #number = (
    #    ( [0-9]+ $dgt %done ) ( '.' @dec [0-9]+ $dgt )?
    #    ( [eE] ( [+\-] $exp_sign )? [0-9]+ $exp )?
    #) %number;

    INT_ATOM = [\-+]? [0-9]{1,17} $dgt;

    STRING_ATOM = [^']*;

    ATOM = ("?"|"!"|"="INT_ATOM|"'"STRING_ATOM"'"|">"UUID) >atom_start %atom;

    OP = (
            ( "." UUID )?   %type >type_start
            ( "#" UUID )?   %object >object_start
            ( "@" UUID )?   %event >event_start
            ( ":" UUID )?   %location >location_start
            (ATOM{1,8})     %value
         ) "\n";

}%%

