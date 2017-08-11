%%{

    machine OP;
    include UUID "./uuid-grammar.rl";

    action redef_uuid {
        switch fc {
        case '`':  *uuid = *prev_uuid
        case '\\': *uuid = context.Spec[1]
        case '|':  *uuid = context.Spec[2]
        case '/':  *uuid = context.Spec[3]
        }
    }

    action toel_start {
        if trace {
            fmt.Printf("TOEL %c\n", fc);
        }
        i = 0
        digits = 0
        old_n = n
        n = -int(ABC[fc]) -30
        uuid = &op.Spec[n]
        *uuid = context.Spec[n]
        if n <= old_n {
            if trace {
                fmt.Printf("MISORDERED UUIDs %c %d %d\n", fc, n, old_n);
            }
            fbreak;
        }
    }

    action toel_uuid {
        prev_uuid = uuid
    }

    action atom_start {
        if op.Count >= MAX_ATOMS {
            fbreak;
        }
        op.Types[op.Count] = fc
        op.Offsets[op.Count] = p - atoms_at
        op.Count++
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
        op.Body = data[atoms_at:p]
    }

    action int_atom {
    }
    action float_atom {
    }
    action string_atom1 {
    }
    action string_atom2 {
    }
    action atoms_start {
        uuid = &blank
        atoms_at = p
    }
    action opterm {
        op.Atoms.Types[MAX_ATOMS] = fc
    }

    action next {
        if trace {
            fmt.Printf("NEXT at %d\n", p)
        }
        p-=1
        done = true
        fbreak;
    }
    action over {
        if trace {
            fmt.Printf("OVER at %d\n", p)
        }
        p-=1
        done = true
        fbreak;
    }

    UNIESC = /\\u[0-9a-fA-F]{4}/;

    INT_ATOM = [\-+]? [0-9]+ %int_atom ;
    FLOAT_ATOM = [\-+]? [0-9]+ "." digit+ ([eE][\-+]?digit+)? %float_atom ;
    STRING_ATOM1 = (UNIESC|"\\" [^\n\r]|[^'])* %string_atom1;
    STRING_ATOM2 = (UNIESC|"\\" [^\n\r]|[^"])* %string_atom2;

    ATOM = space* (
            "=" space* INT_ATOM  |
            "^" space* FLOAT_ATOM |
            "'" STRING_ATOM1 "'" |
            '"' STRING_ATOM2 '"' |
            ">" space* UUID
            ) >atom_start %atom;

    REDEF = ([`\\|/]|"") @redef_uuid;

    OPTERM = space* ("?" [,\.;!]? | [,\.;!] ) @opterm ;

    OP = (
            ( space* [*#@:] @toel_start space* REDEF UUID %toel_uuid )*
            ( (OPTERM | ATOM+ %atoms OPTERM?) >atoms_start ) space*
            ( [*#@:] @next )? %/over
         ) ;

    # main := OP;

}%%
// TODO exact value syntax and case-cov tests
