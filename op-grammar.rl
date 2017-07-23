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
        op.AtomOffsets[op.AtomCount] = ts
        op.AtomCount++
        ////fmt.Printf("ATOM_START %c\n", fc);
    }
    action atom {
        ////fmt.Printf("ATOM\n");
    }

    action atoms {
        if trace {
            fmt.Print("ATOMS\n")
        }
        ret = p
    }


    INT_ATOM = [\-+]? [0-9]+ ;
    FLOAT_ATOM = [\-+]? [0-9]+ ;
    STRING_ATOM1 = /[^']*/;
    STRING_ATOM2 = /[^"]*/;

    ATOM = (
            "?" |
            "!" |
            "=" INT_ATOM |
            "^" FLOAT_ATOM |
            "'" STRING_ATOM1 "'" |
            '"' STRING_ATOM2 '"' |
            ">" UUID
            ) >atom_start %atom;

    REDEF = [`\\|/] @redef_uuid;

    OP = (
            ( [\.#@:] @toel_start REDEF? UUID %toel_uuid )+
            (ATOM+ %atoms ) 
         ) "\n";

    # main := OP;

}%%

