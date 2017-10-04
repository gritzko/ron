%%{

    machine FRAME;
    include UUID "./uuid-grammar.rl";

    action redef_uuid {
        if idx!=0 {
            it.uuids[idx] = it.uuids[idx-1];
        }
    }

    action toel_start {
        n = specSep2Bits(fc) << 1
        if n <= idx {
            fbreak;
        }
        idx = n
    }

    action toel_uuid {
    }

    action atom_start {
        digit = 0;
        i = uint128{0,0};
    }
    action atom_end {
        // TODO max size for int/float/string
        it.AddAtom(i);
    }

    action int_atom_start {
    }
    action int_sign {
        if fc=='-' {
            i[1] |= 1;
        }
    }
    action int_digit {
        i[0] *= 10;
        i[0] += uint64(int(fc) - int('0'));
    }
    action int_atom_end {
        i[1] |= ATOM_INT_62;
    }

    action float_atom_start {
    }
    action float_atom_end {
        i[1] |= ATOM_FLOAT_62;
    }

    action string_atom_start {
        i[0] = uint64(p);
    }
    action string_atom_end {
        i[1] = uint64(p) | ATOM_STRING_62;
    }

    action uuid_atom_start {
    }
    action uuid_atom_end {
        i[1] |= ATOM_UUID_62;
    }

    action atoms_start {
    }
    action atoms {
    }

    action opterm {
        it.term = termSep2Bits(fc)
    }

    action op_end {
        done = true
        fbreak;
    }

    # one of op spec UUIDs: type, object, event id or a reference 
    REDEF = "`" @redef_uuid;
    SPEC_UUID = [*#@:] @toel_start space* REDEF? UUID %toel_uuid space*;

    # 64-bit signed integer 
    INT_ATOM = ([\-+]? @int_sign [0-9]+ @int_digit) %int_atom_end >int_atom_start;

    # 64-bit (double) float 
    FLOAT_ATOM = [\-+]? [0-9]+ ("." digit+)? ([eE][\-+]?digit+)? >float_atom_start %float_atom_end ;

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

    # an optional op terminator, see op types 
    OPTERM = [,;!?] @opterm space*;

    # a RON op; types: (0) raw op (1) reduced op (2) frame header (3) query header 
    OP = space* SPEC_UUID* ( OPTERM | ATOMS OPTERM? ) %op_end;

    # optional frame terminator; mandatory in the streaming mode 
    DOT = "." space*;

    # RON frame, including multiframes (those have more headers inside) 
    FRAME = OP+ DOT*;

}%%
