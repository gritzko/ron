package ron

import "fmt"
import "errors"

%% machine RON;
%% write data;
%% access ps.;
%% variable data frame.Body;
%% variable cs ps.state;

// The parser reached end-of-input (in block mode) or
// the closing dot (in streaming mode) successfully.
// The rest of the input is frame.Rest()
const RON_FULL_STOP = -1;


// Parse consumes one op from data[], unless the buffer ends earlier.
// Fills atoms[]
func (frame *Frame) Parse() {

    ps := &frame.Parser

    switch ps.state {
        case RON_error:
            if ps.position!=0 {
                return
            }
	        %% write init;
            frame.Position = -1
            if len(frame.atoms)<DEFAULT_ATOMS_ALLOC {
                frame.atoms = make([]Atom, 4, DEFAULT_ATOMS_ALLOC)
            }

        case RON_FULL_STOP:
            ps.state = RON_error
            return

        case RON_start:
            ps.offset = ps.position;
            frame.atoms = frame.atoms[:4];
            ps.atm, ps.hlf, ps.dgt = 0, 0, 0;
    }

    if ps.position >= len(frame.Body) {
        if !ps.streaming {
            ps.state = RON_error
        }
        return 
    }

	pe, eof := len(frame.Body), len(frame.Body)
    n := 0
    _ = eof
    _ = pe // FIXME kill

    if ps.streaming {
        eof = -1
    }

    atm, hlf, dgt := ps.atm, ps.hlf, ps.dgt;
    atoms := frame.atoms;
    var e_sgn, e_val, e_frac int
    p := ps.position;

	%%{

        include FRAME "./op-grammar.rl";
        main := FRAME ;

	    write exec;
	}%%

    ps.atm, ps.hlf, ps.dgt = atm, hlf, dgt;
    ps.position = p;
    frame.atoms = atoms;

    switch {
        case ps.state==RON_error:
            frame.Position = -1
        case ps.state>=RON_first_final: // one of end states
            if !ps.streaming && p>=eof {
                // in the block mode, the final dot is optional/implied
                ps.state = RON_FULL_STOP
            }
        case ps.state==RON_FULL_STOP:
        case ps.state==RON_start:
        default:
            if !ps.streaming {
                ps.state = RON_error
                frame.Position = -1
            }
    }

}


func (ctx_uuid UUID) Parse (data []byte) (UUID, error) {

    %% machine UUID;
    %% write data;

	cs, p, pe, eof := 0, 0, len(data), len(data)
    _ = eof

    atm, hlf, dgt := 0, 0, 0

    atoms := [1]Atom{Atom(ctx_uuid)}

	%%{ 

        include UUID "./uuid-grammar.rl";
        main := UUID ;

	    write init;
	    write exec;
	}%%

    if cs < UUID_first_final || dgt>10 {
        return ERROR_UUID, errors.New(fmt.Sprintf("parse error at pos %d", p))
    } else {
        return UUID(atoms[0]), nil 
    }

}

