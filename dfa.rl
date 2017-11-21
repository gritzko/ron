package ron

import "fmt"
import "errors"

%% machine RON;
%% write data;
%% access frame.Parser.;
%% variable p frame.Parser.position;
%% variable data frame.Body;
%% variable cs frame.Parser.state;
const RON_EOF = -1


// Parse consumes one op from data[], unless the buffer ends earlier.
// Fills atoms[], returns op term (TERM_RAW etc) or TERM_ERROR
func (frame *Frame) Parse() {

    if frame.Parser.position >= len(frame.Body) {
        if !frame.Parser.streaming {
            frame.Parser.state = RON_error
        }
        return 
    }

    if frame.Parser.state==RON_error {
        if frame.Parser.position==0 {
	        %% write init;
            frame.atoms = make([]Atom, 4, 8)
        } else {
            return
        }
    } else if frame.Parser.state==RON_EOF {
        frame.Parser.state = RON_error
        return
    } else if frame.Parser.state==RON_start {
        frame.Parser.offset = frame.Parser.position;
        frame.atoms = frame.atoms[:4];
        frame.Parser.atm, frame.Parser.hlf, frame.Parser.dgt = 0, 0, 0;
    }

	pe, eof := len(frame.Body), len(frame.Body)
    n := 0
    _ = eof
    _ = pe // FIXME kill

    if frame.Parser.streaming {
        eof = -1
    }

    atm, hlf, dgt := frame.Parser.atm, frame.Parser.hlf, frame.Parser.dgt;
    atoms := frame.atoms;

	%%{

        include FRAME "./op-grammar.rl";
        main := FRAME ;

	    write exec;
	}%%

    frame.Parser.atm, frame.Parser.hlf, frame.Parser.dgt = atm, hlf, dgt;
    frame.atoms = atoms;

    if !frame.Parser.streaming && frame.Parser.state<RON_first_final && frame.Parser.state>0 {
        frame.Parser.state = RON_error
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

