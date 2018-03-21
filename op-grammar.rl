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
        hlf, dgt = 0, 0;
        if (n < atm) { 
            // parse #op1#op2#op3 without Ragel state explosion
            fnext *RON_start;
            frame.Position++
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
        hlf, dgt = 0, 0;
        atoms = append(atoms, Atom{})
    }
    action atom_end {
        atm++;
    }

    action int_atom_start {
    }
    action int_sign {
        if (fc=='-') {
            atoms[atm][1] |= 1;
        }
    }
    action int_digit {
        atoms[atm][0] *= 10;
        atoms[atm][0] += (uint64)(fc-'0');
        // TODO max size for int/float/string
    }
    action int_atom_end {
        atoms[atm][1] |= ((uint64)(ATOM_INT))<<62;
    }

    action float_atom_start {
        e_sgn = 0;
        e_val = 0;
        e_frac = 0;
    }
    action float_dgt {
        atoms[atm][0] *= 10;
        atoms[atm][0] += (uint64)(fc-'0');
        // TODO max size for int/float/string
    }
    action float_sgn {
        if (fc=='-') {
            atoms[atm][1] |= uint64(1)<<32;
        }
    }
    action float_frac_dgt {
        atoms[atm][0] *= 10;
        atoms[atm][0] += (uint64)(fc-'0');
        e_frac++;
        // TODO max size for int/float/string
    }
    action float_e_sgn {
        if (fc=='-') {
            e_sgn = -1;
        }
    }
    action float_e_dgt {
        e_val *= 10;
        e_val += int(fc-'0');
        // TODO max size for int/float/string
    }
    action float_atom_end {
        if (e_sgn==-1) {
            e_val = -e_val -e_frac;
        } else {
            e_val = +e_val -e_frac;
        }
        if (e_val<0) {
            atoms[atm][1] |= uint64(1)<<33;
            e_val = -e_val;
        }
        atoms[atm][1] |= uint64(e_val)
        atoms[atm][1] |= ((uint64)(ATOM_FLOAT))<<62;
    }

    action string_atom_start {
        atoms[atm][0] = ((uint64)(p))<<32;
    }
    action string_atom_end {
        atoms[atm][0] |= uint64(p);
        atoms[atm][1] = ((uint64)(ATOM_STRING))<<62;
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
        if (p>frame.Parser.offset && frame.Position!=-1) {
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
        frame.Position++
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
    INT_ATOM = ([\-+]? @int_sign ( digit @int_digit )+ ) %int_atom_end >int_atom_start;

    # 64-bit (double) float 
    FLOAT_ATOM = ( [\-+]? @float_sgn [0-9]+ @float_dgt "." [0-9]+ @float_frac_dgt ([eE] [\-+]? @float_e_sgn digit+ @float_e_dgt)? ) >float_atom_start %float_atom_end ;

    UUID_ATOM = UUID >uuid_atom_start %uuid_atom_end;

    # JSON-escaped string 
    UNIESC = "\\u" [0-9a-fA-F]{4};
    ESC = "\\" [^\n\r];
    CHAR = [^"'\n\r\\];
    STRING_ATOM = (UNIESC|ESC|CHAR)* %string_atom_end >string_atom_start;

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
