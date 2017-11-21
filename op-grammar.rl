%%{

    machine FRAME;
    include UUID "./uuid-grammar.rl";

    action redef_uuid {
        if (atm>0) {
            atoms[atm] = atoms[atm-1];
        }
    }

    action spec_uuid_start {
        n = (int)(ABC[fc]);
        if (n < atm) { 
            // wrong UUID order; must be type-object-event-ref
            fnext *RON_error;
            fbreak;
        } else { 
            // start parsing the UUID
            atm = n;
            hlf = 0;
            dgt = 0;
        }
    }

    action spec_uuid_end {
        // OK, save the UUID
        atm++;
        hlf = 0;
        dgt = 0;
    }

    action atom_start {
        dgt = 0;
        atoms = append(atoms, Atom{})
    }
    action atom_end {
        // TODO max size for int/float/string
        atm++;
        hlf = 0;
        dgt = 0;
    }

    action int_atom_start {
        atoms[atm][1] |= ((uint64)(ATOM_INT))<<62;
    }
    action int_sign {
        if (fc=='-') {
            atoms[atm][1] |= 1;
        }
    }
    action int_digit {
        atoms[atm][0] *= 10;
        atoms[atm][0] += (uint64)(fc-'0');
    }
    action int_atom_end {
    }

    action float_atom_start {
        atoms[atm][1] |= ((uint64)(ATOM_FLOAT))<<62;
    }
    action float_atom_end {
    }

    action string_atom_start {
        atoms[atm][0] = ((uint64)(frame.Parser.position))<<32;
    }
    action string_atom_end {
        atoms[atm][0] |= uint64(frame.Parser.position);
        atoms[atm][1] = ((uint64)(ATOM_STRING))<<62;
    }

    action uuid_atom_start {
        atoms[atm][1] |= ((uint64)(ATOM_UUID))<<62;
    }
    action uuid_atom_end {
    }

    action atoms_start {
        atm = 4;
        hlf = 0;
        dgt = 0;
    }
    action atoms {
    }

    action opterm {
        frame.term = int(ABC[fc]);
    }

    action op_start {
        hlf = 0;
        if (frame.Parser.position>frame.Parser.offset) {
            // one op is done, so stop parsing for now
            // make sure the parser restarts with the next op
            frame.Parser.position--;
            fnext *RON_start;
            fbreak;
        } else {
            //op_idx++;
        }
    }

    action op_end {
    }

    action spec_end {
        if (frame.term!=TERM_RAW) {
            frame.term = TERM_REDUCED;
        }
    }

    action frame_end {
        fnext *RON_EOF;
    }

    # one of op spec UUIDs: type, object, event id or a reference 
    REDEF = "`" @redef_uuid;
    QUANT = [*#@:] @spec_uuid_start ;
    SPEC_UUID = QUANT space* REDEF? UUID? %spec_uuid_end space*;

    # 64-bit signed integer 
    INT_ATOM = ([\-+]? @int_sign ( digit @int_digit )+ ) %int_atom_end >int_atom_start;

    # 64-bit (double) float 
    FLOAT_ATOM = [\-+]? [0-9]+ ("." digit+) ([eE][\-+]?digit+)? >float_atom_start %float_atom_end ;

    # JSON-escaped string 
    UNIESC = "\\u" [0-9a-fA-F]{4};
    ESC = "\\" [^\n\r];
    CHAR = [^'\n\r\\];
    STRING_ATOM = (UNIESC|ESC|CHAR)* %string_atom_end >string_atom_start;

    # an atom (int, float, string or UUID) 
    ATOM = (
            "=" space* INT_ATOM  |
            "^" space* FLOAT_ATOM |
            ['] STRING_ATOM ['] |
            ">" space* UUID
            ) >atom_start %atom_end space*;
    # op value - an atom, an atom tuple, or empty 
    ATOMS = ATOM+ %atoms >atoms_start;

    # an optional op terminator (raw, reduced, header, query)
    OPTERM = [,;!?] @opterm space*;

    # a RON op; types: (0) raw op (1) reduced op (2) frame header (3) query header 
    OP = space* ( SPEC_UUID+ >op_start %spec_end ) ( ATOMS OPTERM? | OPTERM ) %op_end;

    # optional frame terminator; mandatory in the streaming mode 
    DOT = "." @frame_end;

    # RON frame, including multiframes (those have more headers inside) 
    FRAME = OP* DOT? ;

}%%
