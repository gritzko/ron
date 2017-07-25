%%{

    machine OP;
    include UUID "./uuid-grammar.rl";

    action redef_uuid {
        switch fc {
        case '`':  *uuid = *prev_uuid
        case '\\': *uuid = context.Object
        case '|':  *uuid = context.Event
        case '/':  *uuid = context.Location
        }
    }

    action toel_start {
        if trace {
            fmt.Printf("TOEL %c\n", fc);
        }
        i = 0
        digits = 0
        switch fc {
            case '.': {
                          n = 0
                          uuid = & (op.Type)
                          *uuid = context.Type
                      }
            case '#': {
                          n = 1
                          uuid = & op.Object
                          *uuid = context.Object
                      }
            case '@': {
                          n = 2
                          uuid = & op.Event
                          *uuid = context.Event
                      }
            case ':': {
                          n = 3
                          uuid = & op.Location
                          *uuid = context.Location
                      }
        }
        if n < prev_uuid_ind {
            // error
        }
    }

    action toel_uuid {
        prev_uuid = uuid
    }

    action atom_start {
        op.AtomTypes[op.AtomCount] = fc
        op.AtomOffsets[op.AtomCount] = p
        op.AtomCount++
        if trace {
            fmt.Printf("ATOM_START %c at %d\n", fc, p);
        }
    }
    action atom {
        if trace {
            fmt.Printf("ATOM at %d\n", p);
        }
    }

    action atoms {
        if trace {
            fmt.Printf("ATOMS at %d\n", p)
        }
        ret = p
    }

    action int_atom {
    }
    action float_atom {
    }
    action string_atom1 {
    }
    action string_atom2 {
    }

    action next {
        if trace {
            fmt.Printf("NEXT at %d\n", p)
        }
        ret = p
        fbreak;
    }


    INT_ATOM = [\-+]? [0-9]+ %int_atom ;
    FLOAT_ATOM = [\-+]? [0-9]+ "." digit+ ([eE][\-+]?digit+)? %float_atom ;
    STRING_ATOM1 = /[^']*/ %string_atom1;
    STRING_ATOM2 = /[^"]*/ %string_atom2;

    ATOM = (
            "?" |
            "!" |
            "=" INT_ATOM  |
            "^" FLOAT_ATOM |
            "'" STRING_ATOM1 "'" |
            '"' STRING_ATOM2 '"' |
            ">" UUID
            ) >atom_start %atom;

    REDEF = ([`\\|/]|"") @redef_uuid;

    OP = (
            ( [\.#@:] @toel_start REDEF UUID %toel_uuid )*
            (ATOM+ %atoms ) 
            ( [\.#@:] @next ) ?
         ) ;

    # main := OP;

}%%
// TODO exact value syntax and case-cov tests
