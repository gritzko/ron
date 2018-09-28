%%{

    machine FRAME;
    include UUID "./uuid-grammar.rl";

    action spec_start {
        ps.omitted = 15;
    }

    action redef_uuid {
        if (atm>0) {
            atoms[atm] = atoms[atm-1];
        }
    }

    action spec_uuid_start {
        n = (int)(ABC[fc]);
        hlf, dgt = VALUE, 0;
        if (n < atm) { 
            // parse #op1#op2#op3 without Ragel state explosion
            fnext *RON_start;
            frame.position ++
            p--;
            fbreak;
        } else { 
            // next UUID
            atm = n;
            ps.omitted -= 1<<uint(n);
        }
    }

    action spec_uuid_end {
        atm++;
    }

    action atom_start {
        hlf, dgt = VALUE, 0;
        atoms = append(atoms, Atom{})
    }

    action atom_end {
        atm++;
    }

    action int_atom_start {
        atoms[atm].setIntType()
        atoms[atm].setFrom(p)
    }

    action float_atom_start {
        atoms[atm].setFloatType()
        atoms[atm].setFrom(p)
    }
    
    action string_atom_start {
        atoms[atm].setStringType()
        atoms[atm].setFrom(p)
    }

    action scalar_atom_end {
        atoms[atm].setTill(p)
        atoms[atm].parseValue(frame.Body)
    }

    action uuid_atom_start {
        if (atm==4) {
            atoms[atm] = atoms[SPEC_OBJECT];
        } else if (atoms[atm-1].Type()==ATOM_UUID) {
            atoms[atm] = atoms[atm-1];
        }
    }

    action uuid_atom_end {
        atoms[atm][1] |= ((uint64)(ATOM_UUID))<<62;
    }

    action atoms_start {
        atm = 4;
        hlf = VALUE;
        dgt = 0;
    }

    action atoms {
    }

    action opterm {
        frame.term = int(ABC[fc]);
    }

    action op_start {
        hlf = VALUE;
        if (p>frame.Parser.off && frame.position!=-1) {
            // one op is done, so stop parsing for now
            // make sure the parser restarts with the next op
            p--;
            fnext *RON_start;
            fbreak;
        } else {
            //op_idx++;
            if (frame.term!=TERM_RAW) {
                frame.term = TERM_REDUCED;
            }
        }
    }

    action op_end {
        frame.position ++
    }

    action spec_end {
    }

    action frame_end {
        fnext *RON_FULL_STOP;
        fbreak;
    }

    # one of op spec UUIDs: type, object, event id or a reference 
    REDEF = "`" @redef_uuid;
    QUANT = [*#@:] @spec_uuid_start ;
    SPEC_UUID = QUANT space* REDEF? (UUID space*)? %spec_uuid_end ;

    # 64-bit signed integer 
    INT_ATOM = ([\-+]? digit+ ) >int_atom_start %scalar_atom_end;

    # 64-bit (double) float 
    FLOAT_ATOM = ( [\-+]? [0-9]+ ("." | ([eE] [\-+]?)) [0-9]+ ([eE] [\-+]? digit+ )? ) >float_atom_start %scalar_atom_end;

    UUID_ATOM = UUID >uuid_atom_start %uuid_atom_end;

    # JSON-escaped string 
    UNIESC = "\\u" [0-9a-fA-F]{4};
    ESC = "\\" [^\n\r];
    CHAR = [^"'\n\r\\];
    STRING_ATOM = (UNIESC|ESC|CHAR)* %scalar_atom_end >string_atom_start;

    # an atom (int, float, string or UUID) 
    ATOM = (
            "=" space* INT_ATOM  |
            "^" space* FLOAT_ATOM |
            ['] STRING_ATOM ['] |
            ">" space* UUID_ATOM
            ) >atom_start %atom_end space*;
    # op value - an atom, an atom tuple, or empty 
    ATOMS = ATOM+ %atoms >atoms_start;

    # an optional op terminator (raw, reduced, header, query)
    OPTERM = [,;!?] @opterm space*;

    # a RON op "specifier" (four UUIDs for its type, object, event, and ref)
    SPEC = SPEC_UUID+ >spec_start %spec_end ;

    # a RON op
    # op types: (0) raw op (1) reduced op (2) frame header (3) query header 
    OP = space* ( SPEC ATOMS? OPTERM? | ATOMS OPTERM? | OPTERM) $2 %1 >op_start %op_end;

    # optional frame terminator; mandatory in the streaming mode 
    DOT = "." @frame_end;

    # RON frame, including multiframes (those have more headers inside) 
    FRAME = OP* DOT? ;

}%%