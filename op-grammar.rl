%%{

    machine OP;
    include UUID "./uuid-grammar.rl";

# TODO  REDEF_UUID

    INT_ATOM = [\-+]? [0-9]+ $dgt;

    STRING_ATOM = /[^']*/;

    ATOM = ("?"|"!"|"="INT_ATOM|"'"STRING_ATOM"'"|">"UUID) >atom_start %atom;

    OP = (
            ( "." UUID )?   %type >type_start
            ( "#" UUID )?   %object >object_start
            ( "@" UUID )?   %event >event_start
            ( ":" UUID )?   %location >location_start
            (ATOM+)     %value
         ) "\n";

}%%

