//line dfa.rl:1
package ron

import "fmt"
import "errors"

//line dfa.rl:7
//line dfa.go:12
const RON_start int = 14
const RON_first_final int = 14
const RON_error int = 0

const RON_en_main int = 14

//line dfa.rl:8

//line dfa.rl:9
//line dfa.rl:10
//line dfa.rl:11
// The parser reached end-of-input (in block mode) or
// the closing dot (in streaming mode) successfully.
// The rest of the input is frame.Rest()
const RON_FULL_STOP = -1

// Parse consumes one op from data[], unless the buffer ends earlier.
// Fills atoms[]
func (frame *Frame) Parse() {

	ps := &frame.Parser

	switch ps.state {
	case RON_error:
		if ps.position != 0 {
			return
		}

//line dfa.go:46
		{
			(ps.state) = RON_start
		}

//line dfa.rl:30
		frame.Position = -1
		if len(frame.atoms) < DEFAULT_ATOMS_ALLOC {
			frame.atoms = make([]Atom, 4, DEFAULT_ATOMS_ALLOC)
		}

	case RON_FULL_STOP:
		ps.state = RON_error
		return

	case RON_start:
		ps.offset = ps.position
		frame.atoms = frame.atoms[:4]
		ps.atm, ps.hlf, ps.dgt = 0, 0, 0
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

	atm, hlf, dgt := ps.atm, ps.hlf, ps.dgt
	atoms := frame.atoms
	var e_sgn, e_val, e_frac int
	p := ps.position

//line dfa.go:89
	{
		if p == pe {
			goto _test_eof
		}
		goto _resume

	_again:
		switch ps.state {
		case 14:
			goto st14
		case 0:
			goto st0
		case 1:
			goto st1
		case 15:
			goto st15
		case 16:
			goto st16
		case 17:
			goto st17
		case 18:
			goto st18
		case 19:
			goto st19
		case 2:
			goto st2
		case 3:
			goto st3
		case 20:
			goto st20
		case 21:
			goto st21
		case 4:
			goto st4
		case 5:
			goto st5
		case 22:
			goto st22
		case 6:
			goto st6
		case 23:
			goto st23
		case 24:
			goto st24
		case 7:
			goto st7
		case 8:
			goto st8
		case 9:
			goto st9
		case 10:
			goto st10
		case 25:
			goto st25
		case 11:
			goto st11
		case 12:
			goto st12
		case 26:
			goto st26
		case 27:
			goto st27
		case 28:
			goto st28
		case 13:
			goto st13
		case 29:
			goto st29
		case 30:
			goto st30
		case 31:
			goto st31
		case 32:
			goto st32
		}

		if p++; p == pe {
			goto _test_eof
		}
	_resume:
		switch ps.state {
		case 14:
			goto st_case_14
		case 0:
			goto st_case_0
		case 1:
			goto st_case_1
		case 15:
			goto st_case_15
		case 16:
			goto st_case_16
		case 17:
			goto st_case_17
		case 18:
			goto st_case_18
		case 19:
			goto st_case_19
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		case 20:
			goto st_case_20
		case 21:
			goto st_case_21
		case 4:
			goto st_case_4
		case 5:
			goto st_case_5
		case 22:
			goto st_case_22
		case 6:
			goto st_case_6
		case 23:
			goto st_case_23
		case 24:
			goto st_case_24
		case 7:
			goto st_case_7
		case 8:
			goto st_case_8
		case 9:
			goto st_case_9
		case 10:
			goto st_case_10
		case 25:
			goto st_case_25
		case 11:
			goto st_case_11
		case 12:
			goto st_case_12
		case 26:
			goto st_case_26
		case 27:
			goto st_case_27
		case 28:
			goto st_case_28
		case 13:
			goto st_case_13
		case 29:
			goto st_case_29
		case 30:
			goto st_case_30
		case 31:
			goto st_case_31
		case 32:
			goto st_case_32
		}
		goto st_out
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		switch (frame.Body)[p] {
		case 32:
			goto st1
		case 35:
			goto tr2
		case 42:
			goto tr2
		case 46:
			goto tr25
		case 58:
			goto tr2
		case 64:
			goto tr2
		}
		if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
			goto st1
		}
		goto st0
	st_case_0:
	st0:
		(ps.state) = 0
		goto _out
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		switch (frame.Body)[p] {
		case 32:
			goto st1
		case 35:
			goto tr2
		case 42:
			goto tr2
		case 58:
			goto tr2
		case 64:
			goto tr2
		}
		if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
			goto st1
		}
		goto st0
	tr2:
		(ps.state) = 15
//line ./op-grammar.rl:131
		hlf = 0
		if p > frame.Parser.offset {
			// one op is done, so stop parsing for now
			// make sure the parser restarts with the next op
			p--
			(ps.state) = (RON_start)
			{
				p++
				goto _out
			}
		} else {
			//op_idx++;
		}

//line ./op-grammar.rl:12
		n = (int)(ABC[(frame.Body)[p]])
		hlf = 0
		dgt = 0
		if n < atm {
			// parse #op1#op2#op3 without Ragel state explosion
			(ps.state) = (RON_start)
			frame.Position++
			p--
			{
				p++
				goto _out
			}
		} else {
			// next UUID
			atm = n
		}

		goto _again
	tr26:
//line ./op-grammar.rl:28
		atm++

		goto st15
	tr28:
		(ps.state) = 15
//line ./op-grammar.rl:28
		atm++

//line ./op-grammar.rl:12
		n = (int)(ABC[(frame.Body)[p]])
		hlf = 0
		dgt = 0
		if n < atm {
			// parse #op1#op2#op3 without Ragel state explosion
			(ps.state) = (RON_start)
			frame.Position++
			p--
			{
				p++
				goto _out
			}
		} else {
			// next UUID
			atm = n
		}

		goto _again
	tr39:
		(ps.state) = 15
//line ./op-grammar.rl:144
		frame.Position++

//line ./op-grammar.rl:131
		hlf = 0
		if p > frame.Parser.offset {
			// one op is done, so stop parsing for now
			// make sure the parser restarts with the next op
			p--
			(ps.state) = (RON_start)
			{
				p++
				goto _out
			}
		} else {
			//op_idx++;
		}

//line ./op-grammar.rl:12
		n = (int)(ABC[(frame.Body)[p]])
		hlf = 0
		dgt = 0
		if n < atm {
			// parse #op1#op2#op3 without Ragel state explosion
			(ps.state) = (RON_start)
			frame.Position++
			p--
			{
				p++
				goto _out
			}
		} else {
			// next UUID
			atm = n
		}

		goto _again
	tr46:
		(ps.state) = 15
//line ./op-grammar.rl:12
		n = (int)(ABC[(frame.Body)[p]])
		hlf = 0
		dgt = 0
		if n < atm {
			// parse #op1#op2#op3 without Ragel state explosion
			(ps.state) = (RON_start)
			frame.Position++
			p--
			{
				p++
				goto _out
			}
		} else {
			// next UUID
			atm = n
		}

		goto _again
	tr54:
		(ps.state) = 15
//line ./op-grammar.rl:37

		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:124

//line ./op-grammar.rl:144
		frame.Position++

//line ./op-grammar.rl:131
		hlf = 0
		if p > frame.Parser.offset {
			// one op is done, so stop parsing for now
			// make sure the parser restarts with the next op
			p--
			(ps.state) = (RON_start)
			{
				p++
				goto _out
			}
		} else {
			//op_idx++;
		}

//line ./op-grammar.rl:12
		n = (int)(ABC[(frame.Body)[p]])
		hlf = 0
		dgt = 0
		if n < atm {
			// parse #op1#op2#op3 without Ragel state explosion
			(ps.state) = (RON_start)
			frame.Position++
			p--
			{
				p++
				goto _out
			}
		} else {
			// next UUID
			atm = n
		}

		goto _again
	tr62:
		(ps.state) = 15
//line ./op-grammar.rl:124

//line ./op-grammar.rl:144
		frame.Position++

//line ./op-grammar.rl:131
		hlf = 0
		if p > frame.Parser.offset {
			// one op is done, so stop parsing for now
			// make sure the parser restarts with the next op
			p--
			(ps.state) = (RON_start)
			{
				p++
				goto _out
			}
		} else {
			//op_idx++;
		}

//line ./op-grammar.rl:12
		n = (int)(ABC[(frame.Body)[p]])
		hlf = 0
		dgt = 0
		if n < atm {
			// parse #op1#op2#op3 without Ragel state explosion
			(ps.state) = (RON_start)
			frame.Position++
			p--
			{
				p++
				goto _out
			}
		} else {
			// next UUID
			atm = n
		}

		goto _again
	tr70:
		(ps.state) = 15
//line ./op-grammar.rl:53
		atoms[atm][1] |= ((uint64)(ATOM_INT)) << 62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:124

//line ./op-grammar.rl:144
		frame.Position++

//line ./op-grammar.rl:131
		hlf = 0
		if p > frame.Parser.offset {
			// one op is done, so stop parsing for now
			// make sure the parser restarts with the next op
			p--
			(ps.state) = (RON_start)
			{
				p++
				goto _out
			}
		} else {
			//op_idx++;
		}

//line ./op-grammar.rl:12
		n = (int)(ABC[(frame.Body)[p]])
		hlf = 0
		dgt = 0
		if n < atm {
			// parse #op1#op2#op3 without Ragel state explosion
			(ps.state) = (RON_start)
			frame.Position++
			p--
			{
				p++
				goto _out
			}
		} else {
			// next UUID
			atm = n
		}

		goto _again
	tr78:
		(ps.state) = 15
//line ./op-grammar.rl:114
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:124

//line ./op-grammar.rl:144
		frame.Position++

//line ./op-grammar.rl:131
		hlf = 0
		if p > frame.Parser.offset {
			// one op is done, so stop parsing for now
			// make sure the parser restarts with the next op
			p--
			(ps.state) = (RON_start)
			{
				p++
				goto _out
			}
		} else {
			//op_idx++;
		}

//line ./op-grammar.rl:12
		n = (int)(ABC[(frame.Body)[p]])
		hlf = 0
		dgt = 0
		if n < atm {
			// parse #op1#op2#op3 without Ragel state explosion
			(ps.state) = (RON_start)
			frame.Position++
			p--
			{
				p++
				goto _out
			}
		} else {
			// next UUID
			atm = n
		}

		goto _again
	tr88:
		(ps.state) = 15
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:114
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:124

//line ./op-grammar.rl:144
		frame.Position++

//line ./op-grammar.rl:131
		hlf = 0
		if p > frame.Parser.offset {
			// one op is done, so stop parsing for now
			// make sure the parser restarts with the next op
			p--
			(ps.state) = (RON_start)
			{
				p++
				goto _out
			}
		} else {
			//op_idx++;
		}

//line ./op-grammar.rl:12
		n = (int)(ABC[(frame.Body)[p]])
		hlf = 0
		dgt = 0
		if n < atm {
			// parse #op1#op2#op3 without Ragel state explosion
			(ps.state) = (RON_start)
			frame.Position++
			p--
			{
				p++
				goto _out
			}
		} else {
			// next UUID
			atm = n
		}

		goto _again
	tr97:
		(ps.state) = 15
//line ./op-grammar.rl:85
		if e_sgn == -1 {
			e_val = -e_val - e_frac
		} else {
			e_val = +e_val - e_frac
		}
		if e_val < 0 {
			atoms[atm][1] |= uint64(1) << 33
			e_val = -e_val
		}
		atoms[atm][1] |= uint64(e_val)
		atoms[atm][1] |= ((uint64)(ATOM_FLOAT)) << 62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:124

//line ./op-grammar.rl:144
		frame.Position++

//line ./op-grammar.rl:131
		hlf = 0
		if p > frame.Parser.offset {
			// one op is done, so stop parsing for now
			// make sure the parser restarts with the next op
			p--
			(ps.state) = (RON_start)
			{
				p++
				goto _out
			}
		} else {
			//op_idx++;
		}

//line ./op-grammar.rl:12
		n = (int)(ABC[(frame.Body)[p]])
		hlf = 0
		dgt = 0
		if n < atm {
			// parse #op1#op2#op3 without Ragel state explosion
			(ps.state) = (RON_start)
			frame.Position++
			p--
			{
				p++
				goto _out
			}
		} else {
			// next UUID
			atm = n
		}

		goto _again
	tr106:
		(ps.state) = 15
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:114
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:124

//line ./op-grammar.rl:144
		frame.Position++

//line ./op-grammar.rl:131
		hlf = 0
		if p > frame.Parser.offset {
			// one op is done, so stop parsing for now
			// make sure the parser restarts with the next op
			p--
			(ps.state) = (RON_start)
			{
				p++
				goto _out
			}
		} else {
			//op_idx++;
		}

//line ./op-grammar.rl:12
		n = (int)(ABC[(frame.Body)[p]])
		hlf = 0
		dgt = 0
		if n < atm {
			// parse #op1#op2#op3 without Ragel state explosion
			(ps.state) = (RON_start)
			frame.Position++
			p--
			{
				p++
				goto _out
			}
		} else {
			// next UUID
			atm = n
		}

		goto _again
	tr117:
		(ps.state) = 15
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:114
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:124

//line ./op-grammar.rl:144
		frame.Position++

//line ./op-grammar.rl:131
		hlf = 0
		if p > frame.Parser.offset {
			// one op is done, so stop parsing for now
			// make sure the parser restarts with the next op
			p--
			(ps.state) = (RON_start)
			{
				p++
				goto _out
			}
		} else {
			//op_idx++;
		}

//line ./op-grammar.rl:12
		n = (int)(ABC[(frame.Body)[p]])
		hlf = 0
		dgt = 0
		if n < atm {
			// parse #op1#op2#op3 without Ragel state explosion
			(ps.state) = (RON_start)
			frame.Position++
			p--
			{
				p++
				goto _out
			}
		} else {
			// next UUID
			atm = n
		}

		goto _again
	tr126:
		(ps.state) = 15
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:28
		atm++

//line ./op-grammar.rl:12
		n = (int)(ABC[(frame.Body)[p]])
		hlf = 0
		dgt = 0
		if n < atm {
			// parse #op1#op2#op3 without Ragel state explosion
			(ps.state) = (RON_start)
			frame.Position++
			p--
			{
				p++
				goto _out
			}
		} else {
			// next UUID
			atm = n
		}

		goto _again
	tr135:
		(ps.state) = 15
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:28
		atm++

//line ./op-grammar.rl:12
		n = (int)(ABC[(frame.Body)[p]])
		hlf = 0
		dgt = 0
		if n < atm {
			// parse #op1#op2#op3 without Ragel state explosion
			(ps.state) = (RON_start)
			frame.Position++
			p--
			{
				p++
				goto _out
			}
		} else {
			// next UUID
			atm = n
		}

		goto _again
	tr146:
		(ps.state) = 15
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:28
		atm++

//line ./op-grammar.rl:12
		n = (int)(ABC[(frame.Body)[p]])
		hlf = 0
		dgt = 0
		if n < atm {
			// parse #op1#op2#op3 without Ragel state explosion
			(ps.state) = (RON_start)
			frame.Position++
			p--
			{
				p++
				goto _out
			}
		} else {
			// next UUID
			atm = n
		}

		goto _again
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
//line dfa.go:884
		switch (frame.Body)[p] {
		case 32:
			goto tr26
		case 33:
			goto tr27
		case 35:
			goto tr28
		case 39:
			goto tr30
		case 42:
			goto tr28
		case 44:
			goto tr27
		case 46:
			goto tr32
		case 58:
			goto tr28
		case 59:
			goto tr27
		case 61:
			goto tr34
		case 62:
			goto tr35
		case 63:
			goto tr27
		case 64:
			goto tr28
		case 91:
			goto tr31
		case 93:
			goto tr31
		case 94:
			goto tr36
		case 96:
			goto tr37
		case 123:
			goto tr31
		case 125:
			goto tr31
		case 126:
			goto tr33
		}
		switch {
		case (frame.Body)[p] < 43:
			switch {
			case (frame.Body)[p] < 36:
				if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
					goto tr26
				}
			case (frame.Body)[p] > 37:
				if 40 <= (frame.Body)[p] && (frame.Body)[p] <= 41 {
					goto tr31
				}
			default:
				goto tr29
			}
		case (frame.Body)[p] > 45:
			switch {
			case (frame.Body)[p] < 65:
				if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
					goto tr33
				}
			case (frame.Body)[p] > 90:
				if 95 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr33
				}
			default:
				goto tr33
			}
		default:
			goto tr29
		}
		goto st0
	tr27:
//line ./op-grammar.rl:28
		atm++

//line ./op-grammar.rl:148
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:127
		frame.term = int(ABC[(frame.Body)[p]])

		goto st16
	tr38:
//line ./op-grammar.rl:144
		frame.Position++

		goto st16
	tr45:
//line ./op-grammar.rl:148
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:127
		frame.term = int(ABC[(frame.Body)[p]])

		goto st16
	tr53:
//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:124

//line ./op-grammar.rl:127
		frame.term = int(ABC[(frame.Body)[p]])

		goto st16
	tr61:
//line ./op-grammar.rl:124
//line ./op-grammar.rl:127
		frame.term = int(ABC[(frame.Body)[p]])

		goto st16
	tr69:
//line ./op-grammar.rl:53
		atoms[atm][1] |= ((uint64)(ATOM_INT)) << 62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:124

//line ./op-grammar.rl:127
		frame.term = int(ABC[(frame.Body)[p]])

		goto st16
	tr77:
//line ./op-grammar.rl:114
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:124

//line ./op-grammar.rl:127
		frame.term = int(ABC[(frame.Body)[p]])

		goto st16
	tr87:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:114
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:124

//line ./op-grammar.rl:127
		frame.term = int(ABC[(frame.Body)[p]])

		goto st16
	tr96:
//line ./op-grammar.rl:85
		if e_sgn == -1 {
			e_val = -e_val - e_frac
		} else {
			e_val = +e_val - e_frac
		}
		if e_val < 0 {
			atoms[atm][1] |= uint64(1) << 33
			e_val = -e_val
		}
		atoms[atm][1] |= uint64(e_val)
		atoms[atm][1] |= ((uint64)(ATOM_FLOAT)) << 62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:124

//line ./op-grammar.rl:127
		frame.term = int(ABC[(frame.Body)[p]])

		goto st16
	tr105:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:114
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:124

//line ./op-grammar.rl:127
		frame.term = int(ABC[(frame.Body)[p]])

		goto st16
	tr116:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:114
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:124

//line ./op-grammar.rl:127
		frame.term = int(ABC[(frame.Body)[p]])

		goto st16
	tr125:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:28
		atm++

//line ./op-grammar.rl:148
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:127
		frame.term = int(ABC[(frame.Body)[p]])

		goto st16
	tr134:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:28
		atm++

//line ./op-grammar.rl:148
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:127
		frame.term = int(ABC[(frame.Body)[p]])

		goto st16
	tr145:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:28
		atm++

//line ./op-grammar.rl:148
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:127
		frame.term = int(ABC[(frame.Body)[p]])

		goto st16
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
//line dfa.go:1212
		switch (frame.Body)[p] {
		case 32:
			goto tr38
		case 35:
			goto tr39
		case 42:
			goto tr39
		case 46:
			goto tr40
		case 58:
			goto tr39
		case 64:
			goto tr39
		}
		if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
			goto tr38
		}
		goto st0
	tr25:
		(ps.state) = 17
//line ./op-grammar.rl:154
		(ps.state) = (RON_FULL_STOP)

		goto _again
	tr32:
		(ps.state) = 17
//line ./op-grammar.rl:28
		atm++

//line ./op-grammar.rl:148
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:144
		frame.Position++

//line ./op-grammar.rl:154
		(ps.state) = (RON_FULL_STOP)

		goto _again
	tr40:
		(ps.state) = 17
//line ./op-grammar.rl:144
		frame.Position++

//line ./op-grammar.rl:154
		(ps.state) = (RON_FULL_STOP)

		goto _again
	tr48:
		(ps.state) = 17
//line ./op-grammar.rl:148
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:144
		frame.Position++

//line ./op-grammar.rl:154
		(ps.state) = (RON_FULL_STOP)

		goto _again
	tr56:
		(ps.state) = 17
//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:124

//line ./op-grammar.rl:144
		frame.Position++

//line ./op-grammar.rl:154
		(ps.state) = (RON_FULL_STOP)

		goto _again
	tr64:
		(ps.state) = 17
//line ./op-grammar.rl:124
//line ./op-grammar.rl:144
		frame.Position++

//line ./op-grammar.rl:154
		(ps.state) = (RON_FULL_STOP)

		goto _again
	tr72:
		(ps.state) = 17
//line ./op-grammar.rl:53
		atoms[atm][1] |= ((uint64)(ATOM_INT)) << 62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:124

//line ./op-grammar.rl:144
		frame.Position++

//line ./op-grammar.rl:154
		(ps.state) = (RON_FULL_STOP)

		goto _again
	tr81:
		(ps.state) = 17
//line ./op-grammar.rl:114
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:124

//line ./op-grammar.rl:144
		frame.Position++

//line ./op-grammar.rl:154
		(ps.state) = (RON_FULL_STOP)

		goto _again
	tr90:
		(ps.state) = 17
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:114
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:124

//line ./op-grammar.rl:144
		frame.Position++

//line ./op-grammar.rl:154
		(ps.state) = (RON_FULL_STOP)

		goto _again
	tr99:
		(ps.state) = 17
//line ./op-grammar.rl:85
		if e_sgn == -1 {
			e_val = -e_val - e_frac
		} else {
			e_val = +e_val - e_frac
		}
		if e_val < 0 {
			atoms[atm][1] |= uint64(1) << 33
			e_val = -e_val
		}
		atoms[atm][1] |= uint64(e_val)
		atoms[atm][1] |= ((uint64)(ATOM_FLOAT)) << 62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:124

//line ./op-grammar.rl:144
		frame.Position++

//line ./op-grammar.rl:154
		(ps.state) = (RON_FULL_STOP)

		goto _again
	tr110:
		(ps.state) = 17
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:114
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:124

//line ./op-grammar.rl:144
		frame.Position++

//line ./op-grammar.rl:154
		(ps.state) = (RON_FULL_STOP)

		goto _again
	tr119:
		(ps.state) = 17
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:114
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:124

//line ./op-grammar.rl:144
		frame.Position++

//line ./op-grammar.rl:154
		(ps.state) = (RON_FULL_STOP)

		goto _again
	tr128:
		(ps.state) = 17
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:28
		atm++

//line ./op-grammar.rl:148
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:144
		frame.Position++

//line ./op-grammar.rl:154
		(ps.state) = (RON_FULL_STOP)

		goto _again
	tr139:
		(ps.state) = 17
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:28
		atm++

//line ./op-grammar.rl:148
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:144
		frame.Position++

//line ./op-grammar.rl:154
		(ps.state) = (RON_FULL_STOP)

		goto _again
	tr148:
		(ps.state) = 17
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:28
		atm++

//line ./op-grammar.rl:148
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:144
		frame.Position++

//line ./op-grammar.rl:154
		(ps.state) = (RON_FULL_STOP)

		goto _again
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
//line dfa.go:1562
		goto st0
	tr29:
//line ././uuid-grammar.rl:5

//line ././uuid-grammar.rl:40
		hlf = 1
		atoms[atm][1] &= INT60_FULL
		atoms[atm][1] |= ((uint64)(ABC[(frame.Body)[p]])) << 60

		goto st18
	tr136:
//line ././uuid-grammar.rl:34
//line ././uuid-grammar.rl:40
		hlf = 1
		atoms[atm][1] &= INT60_FULL
		atoms[atm][1] |= ((uint64)(ABC[(frame.Body)[p]])) << 60

		goto st18
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
//line dfa.go:1591
		switch (frame.Body)[p] {
		case 32:
			goto tr41
		case 33:
			goto tr27
		case 35:
			goto tr28
		case 39:
			goto tr30
		case 42:
			goto tr28
		case 44:
			goto tr27
		case 46:
			goto tr32
		case 58:
			goto tr28
		case 59:
			goto tr27
		case 61:
			goto tr34
		case 62:
			goto tr35
		case 63:
			goto tr27
		case 64:
			goto tr28
		case 91:
			goto tr42
		case 93:
			goto tr42
		case 94:
			goto tr36
		case 95:
			goto tr43
		case 123:
			goto tr42
		case 125:
			goto tr42
		case 126:
			goto tr43
		}
		switch {
		case (frame.Body)[p] < 48:
			switch {
			case (frame.Body)[p] > 13:
				if 40 <= (frame.Body)[p] && (frame.Body)[p] <= 41 {
					goto tr42
				}
			case (frame.Body)[p] >= 9:
				goto tr41
			}
		case (frame.Body)[p] > 57:
			switch {
			case (frame.Body)[p] > 90:
				if 97 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr43
				}
			case (frame.Body)[p] >= 65:
				goto tr43
			}
		default:
			goto tr43
		}
		goto st0
	tr41:
//line ./op-grammar.rl:28
		atm++

		goto st19
	tr124:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:28
		atm++

		goto st19
	tr133:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:28
		atm++

		goto st19
	tr144:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:28
		atm++

		goto st19
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
//line dfa.go:1699
		switch (frame.Body)[p] {
		case 32:
			goto st19
		case 33:
			goto tr45
		case 35:
			goto tr46
		case 39:
			goto tr47
		case 42:
			goto tr46
		case 44:
			goto tr45
		case 46:
			goto tr48
		case 58:
			goto tr46
		case 59:
			goto tr45
		case 61:
			goto tr49
		case 62:
			goto tr50
		case 63:
			goto tr45
		case 64:
			goto tr46
		case 94:
			goto tr51
		}
		if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
			goto st19
		}
		goto st0
	tr30:
//line ./op-grammar.rl:28
		atm++

//line ./op-grammar.rl:148
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:119
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st2
	tr47:
//line ./op-grammar.rl:148
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:119
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st2
	tr55:
//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st2
	tr63:
//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st2
	tr71:
//line ./op-grammar.rl:53
		atoms[atm][1] |= ((uint64)(ATOM_INT)) << 62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st2
	tr79:
//line ./op-grammar.rl:114
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st2
	tr89:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:114
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st2
	tr98:
//line ./op-grammar.rl:85
		if e_sgn == -1 {
			e_val = -e_val - e_frac
		} else {
			e_val = +e_val - e_frac
		}
		if e_val < 0 {
			atoms[atm][1] |= uint64(1) << 33
			e_val = -e_val
		}
		atoms[atm][1] |= uint64(e_val)
		atoms[atm][1] |= ((uint64)(ATOM_FLOAT)) << 62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st2
	tr108:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:114
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st2
	tr118:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:114
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st2
	tr127:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:28
		atm++

//line ./op-grammar.rl:148
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:119
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st2
	tr137:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:28
		atm++

//line ./op-grammar.rl:148
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:119
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st2
	tr147:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:28
		atm++

//line ./op-grammar.rl:148
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:119
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
//line dfa.go:2014
		switch (frame.Body)[p] {
		case 10:
			goto st0
		case 13:
			goto st0
		case 34:
			goto st0
		case 39:
			goto tr4
		case 92:
			goto tr5
		}
		goto tr3
	tr3:
//line ./op-grammar.rl:99
		atoms[atm][0] = ((uint64)(p)) << 32

		goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
//line dfa.go:2039
		switch (frame.Body)[p] {
		case 10:
			goto st0
		case 13:
			goto st0
		case 34:
			goto st0
		case 39:
			goto tr7
		case 92:
			goto st13
		}
		goto st3
	tr4:
//line ./op-grammar.rl:99
		atoms[atm][0] = ((uint64)(p)) << 32

//line ./op-grammar.rl:102
		atoms[atm][0] |= uint64(p)
		atoms[atm][1] = ((uint64)(ATOM_STRING)) << 62

		goto st20
	tr7:
//line ./op-grammar.rl:102
		atoms[atm][0] |= uint64(p)
		atoms[atm][1] = ((uint64)(ATOM_STRING)) << 62

		goto st20
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
//line dfa.go:2076
		switch (frame.Body)[p] {
		case 32:
			goto tr52
		case 33:
			goto tr53
		case 35:
			goto tr54
		case 39:
			goto tr55
		case 42:
			goto tr54
		case 44:
			goto tr53
		case 46:
			goto tr56
		case 58:
			goto tr54
		case 59:
			goto tr53
		case 61:
			goto tr57
		case 62:
			goto tr58
		case 63:
			goto tr53
		case 64:
			goto tr54
		case 94:
			goto tr59
		}
		if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
			goto tr52
		}
		goto st0
	tr52:
//line ./op-grammar.rl:37

		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:124

//line ./op-grammar.rl:144
		frame.Position++

		goto st21
	tr60:
//line ./op-grammar.rl:124
//line ./op-grammar.rl:144
		frame.Position++

		goto st21
	tr68:
//line ./op-grammar.rl:53
		atoms[atm][1] |= ((uint64)(ATOM_INT)) << 62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:124

//line ./op-grammar.rl:144
		frame.Position++

		goto st21
	tr76:
//line ./op-grammar.rl:114
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:124

//line ./op-grammar.rl:144
		frame.Position++

		goto st21
	tr86:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:114
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:124

//line ./op-grammar.rl:144
		frame.Position++

		goto st21
	tr95:
//line ./op-grammar.rl:85
		if e_sgn == -1 {
			e_val = -e_val - e_frac
		} else {
			e_val = +e_val - e_frac
		}
		if e_val < 0 {
			atoms[atm][1] |= uint64(1) << 33
			e_val = -e_val
		}
		atoms[atm][1] |= uint64(e_val)
		atoms[atm][1] |= ((uint64)(ATOM_FLOAT)) << 62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:124

//line ./op-grammar.rl:144
		frame.Position++

		goto st21
	tr104:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:114
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:124

//line ./op-grammar.rl:144
		frame.Position++

		goto st21
	tr115:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:114
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:124

//line ./op-grammar.rl:144
		frame.Position++

		goto st21
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
//line dfa.go:2270
		switch (frame.Body)[p] {
		case 32:
			goto tr60
		case 33:
			goto tr61
		case 35:
			goto tr62
		case 39:
			goto tr63
		case 42:
			goto tr62
		case 44:
			goto tr61
		case 46:
			goto tr64
		case 58:
			goto tr62
		case 59:
			goto tr61
		case 61:
			goto tr65
		case 62:
			goto tr66
		case 63:
			goto tr61
		case 64:
			goto tr62
		case 94:
			goto tr67
		}
		if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
			goto tr60
		}
		goto st0
	tr34:
//line ./op-grammar.rl:28
		atm++

//line ./op-grammar.rl:148
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:119
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st4
	tr49:
//line ./op-grammar.rl:148
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:119
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st4
	tr57:
//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st4
	tr65:
//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st4
	tr73:
//line ./op-grammar.rl:53
		atoms[atm][1] |= ((uint64)(ATOM_INT)) << 62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st4
	tr83:
//line ./op-grammar.rl:114
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st4
	tr92:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:114
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st4
	tr100:
//line ./op-grammar.rl:85
		if e_sgn == -1 {
			e_val = -e_val - e_frac
		} else {
			e_val = +e_val - e_frac
		}
		if e_val < 0 {
			atoms[atm][1] |= uint64(1) << 33
			e_val = -e_val
		}
		atoms[atm][1] |= uint64(e_val)
		atoms[atm][1] |= ((uint64)(ATOM_FLOAT)) << 62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st4
	tr112:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:114
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st4
	tr121:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:114
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st4
	tr130:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:28
		atm++

//line ./op-grammar.rl:148
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:119
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st4
	tr141:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:28
		atm++

//line ./op-grammar.rl:148
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:119
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st4
	tr150:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:28
		atm++

//line ./op-grammar.rl:148
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:119
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
//line dfa.go:2585
		switch (frame.Body)[p] {
		case 32:
			goto st4
		case 43:
			goto tr10
		case 45:
			goto tr10
		}
		switch {
		case (frame.Body)[p] > 13:
			if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
				goto tr11
			}
		case (frame.Body)[p] >= 9:
			goto st4
		}
		goto st0
	tr10:
//line ./op-grammar.rl:42

//line ./op-grammar.rl:44
		if (frame.Body)[p] == '-' {
			atoms[atm][1] |= 1
		}

		goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
//line dfa.go:2619
		if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
			goto tr12
		}
		goto st0
	tr11:
//line ./op-grammar.rl:42

//line ./op-grammar.rl:49
		atoms[atm][0] *= 10
		atoms[atm][0] += (uint64)((frame.Body)[p] - '0')

		goto st22
	tr12:
//line ./op-grammar.rl:49
		atoms[atm][0] *= 10
		atoms[atm][0] += (uint64)((frame.Body)[p] - '0')

		goto st22
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
//line dfa.go:2646
		switch (frame.Body)[p] {
		case 32:
			goto tr68
		case 33:
			goto tr69
		case 35:
			goto tr70
		case 39:
			goto tr71
		case 42:
			goto tr70
		case 44:
			goto tr69
		case 46:
			goto tr72
		case 58:
			goto tr70
		case 59:
			goto tr69
		case 61:
			goto tr73
		case 62:
			goto tr74
		case 63:
			goto tr69
		case 64:
			goto tr70
		case 94:
			goto tr75
		}
		switch {
		case (frame.Body)[p] > 13:
			if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
				goto tr12
			}
		case (frame.Body)[p] >= 9:
			goto tr68
		}
		goto st0
	tr35:
//line ./op-grammar.rl:28
		atm++

//line ./op-grammar.rl:148
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:119
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st6
	tr50:
//line ./op-grammar.rl:148
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:119
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st6
	tr58:
//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st6
	tr66:
//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st6
	tr74:
//line ./op-grammar.rl:53
		atoms[atm][1] |= ((uint64)(ATOM_INT)) << 62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st6
	tr84:
//line ./op-grammar.rl:114
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st6
	tr93:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:114
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st6
	tr101:
//line ./op-grammar.rl:85
		if e_sgn == -1 {
			e_val = -e_val - e_frac
		} else {
			e_val = +e_val - e_frac
		}
		if e_val < 0 {
			atoms[atm][1] |= uint64(1) << 33
			e_val = -e_val
		}
		atoms[atm][1] |= uint64(e_val)
		atoms[atm][1] |= ((uint64)(ATOM_FLOAT)) << 62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st6
	tr113:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:114
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st6
	tr122:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:114
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st6
	tr131:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:28
		atm++

//line ./op-grammar.rl:148
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:119
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st6
	tr142:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:28
		atm++

//line ./op-grammar.rl:148
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:119
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st6
	tr151:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:28
		atm++

//line ./op-grammar.rl:148
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:119
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
//line dfa.go:2966
		switch (frame.Body)[p] {
		case 32:
			goto st6
		case 43:
			goto tr14
		case 45:
			goto tr14
		case 91:
			goto tr15
		case 93:
			goto tr15
		case 95:
			goto tr16
		case 123:
			goto tr15
		case 125:
			goto tr15
		case 126:
			goto tr16
		}
		switch {
		case (frame.Body)[p] < 40:
			switch {
			case (frame.Body)[p] > 13:
				if 36 <= (frame.Body)[p] && (frame.Body)[p] <= 37 {
					goto tr14
				}
			case (frame.Body)[p] >= 9:
				goto st6
			}
		case (frame.Body)[p] > 41:
			switch {
			case (frame.Body)[p] < 65:
				if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
					goto tr16
				}
			case (frame.Body)[p] > 90:
				if 97 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr16
				}
			default:
				goto tr16
			}
		default:
			goto tr15
		}
		goto st0
	tr14:
//line ./op-grammar.rl:107
		if atm == 4 {
			atoms[atm] = atoms[SPEC_OBJECT]
		} else if atoms[atm-1].Type() == ATOM_UUID {
			atoms[atm] = atoms[atm-1]
		}

//line ././uuid-grammar.rl:5

//line ././uuid-grammar.rl:40
		hlf = 1
		atoms[atm][1] &= INT60_FULL
		atoms[atm][1] |= ((uint64)(ABC[(frame.Body)[p]])) << 60

		goto st23
	tr107:
//line ././uuid-grammar.rl:34
//line ././uuid-grammar.rl:40
		hlf = 1
		atoms[atm][1] &= INT60_FULL
		atoms[atm][1] |= ((uint64)(ABC[(frame.Body)[p]])) << 60

		goto st23
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
//line dfa.go:3049
		switch (frame.Body)[p] {
		case 32:
			goto tr76
		case 33:
			goto tr77
		case 35:
			goto tr78
		case 39:
			goto tr79
		case 42:
			goto tr78
		case 44:
			goto tr77
		case 46:
			goto tr81
		case 58:
			goto tr78
		case 59:
			goto tr77
		case 61:
			goto tr83
		case 62:
			goto tr84
		case 63:
			goto tr77
		case 64:
			goto tr78
		case 91:
			goto tr80
		case 93:
			goto tr80
		case 94:
			goto tr85
		case 95:
			goto tr82
		case 123:
			goto tr80
		case 125:
			goto tr80
		case 126:
			goto tr82
		}
		switch {
		case (frame.Body)[p] < 48:
			switch {
			case (frame.Body)[p] > 13:
				if 40 <= (frame.Body)[p] && (frame.Body)[p] <= 41 {
					goto tr80
				}
			case (frame.Body)[p] >= 9:
				goto tr76
			}
		case (frame.Body)[p] > 57:
			switch {
			case (frame.Body)[p] > 90:
				if 97 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr82
				}
			case (frame.Body)[p] >= 65:
				goto tr82
			}
		default:
			goto tr82
		}
		goto st0
	tr80:
//line ././uuid-grammar.rl:29
		dgt = 0
		hlf = 1

//line ././uuid-grammar.rl:9
		dgt = int(ABC[(frame.Body)[p]]) + 4
		atoms[atm][hlf] &= INT60_FLAGS | PREFIX_MASKS[dgt]

		goto st24
	tr82:
//line ././uuid-grammar.rl:29
		dgt = 0
		hlf = 1

//line ././uuid-grammar.rl:22
		atoms[atm][hlf] &= INT60_FLAGS

//line ././uuid-grammar.rl:14
		atoms[atm][hlf] |= ((uint64)(ABC[(frame.Body)[p]])) << DIGIT_OFFSETS[dgt]
		dgt++
		if dgt > 10 {
			{
				p++
				(ps.state) = 24
				goto _out
			}
		}

		goto st24
	tr91:
//line ././uuid-grammar.rl:14
		atoms[atm][hlf] |= ((uint64)(ABC[(frame.Body)[p]])) << DIGIT_OFFSETS[dgt]
		dgt++
		if dgt > 10 {
			{
				p++
				(ps.state) = 24
				goto _out
			}
		}

		goto st24
	tr109:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:29
		dgt = 0
		hlf = 1

//line ././uuid-grammar.rl:9
		dgt = int(ABC[(frame.Body)[p]]) + 4
		atoms[atm][hlf] &= INT60_FLAGS | PREFIX_MASKS[dgt]

		goto st24
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
//line dfa.go:3176
		switch (frame.Body)[p] {
		case 32:
			goto tr86
		case 33:
			goto tr87
		case 35:
			goto tr88
		case 39:
			goto tr89
		case 42:
			goto tr88
		case 44:
			goto tr87
		case 46:
			goto tr90
		case 58:
			goto tr88
		case 59:
			goto tr87
		case 61:
			goto tr92
		case 62:
			goto tr93
		case 63:
			goto tr87
		case 64:
			goto tr88
		case 94:
			goto tr94
		case 95:
			goto tr91
		case 126:
			goto tr91
		}
		switch {
		case (frame.Body)[p] < 48:
			if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
				goto tr86
			}
		case (frame.Body)[p] > 57:
			switch {
			case (frame.Body)[p] > 90:
				if 97 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr91
				}
			case (frame.Body)[p] >= 65:
				goto tr91
			}
		default:
			goto tr91
		}
		goto st0
	tr36:
//line ./op-grammar.rl:28
		atm++

//line ./op-grammar.rl:148
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:119
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st7
	tr51:
//line ./op-grammar.rl:148
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:119
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st7
	tr59:
//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st7
	tr67:
//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st7
	tr75:
//line ./op-grammar.rl:53
		atoms[atm][1] |= ((uint64)(ATOM_INT)) << 62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st7
	tr85:
//line ./op-grammar.rl:114
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st7
	tr94:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:114
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st7
	tr103:
//line ./op-grammar.rl:85
		if e_sgn == -1 {
			e_val = -e_val - e_frac
		} else {
			e_val = +e_val - e_frac
		}
		if e_val < 0 {
			atoms[atm][1] |= uint64(1) << 33
			e_val = -e_val
		}
		atoms[atm][1] |= uint64(e_val)
		atoms[atm][1] |= ((uint64)(ATOM_FLOAT)) << 62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st7
	tr114:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:114
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st7
	tr123:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:114
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++

//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st7
	tr132:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:28
		atm++

//line ./op-grammar.rl:148
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:119
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st7
	tr143:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:28
		atm++

//line ./op-grammar.rl:148
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:119
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st7
	tr152:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:28
		atm++

//line ./op-grammar.rl:148
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:119
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:32
		hlf = 0
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
//line dfa.go:3509
		switch (frame.Body)[p] {
		case 32:
			goto st7
		case 43:
			goto tr18
		case 45:
			goto tr18
		}
		switch {
		case (frame.Body)[p] > 13:
			if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
				goto tr19
			}
		case (frame.Body)[p] >= 9:
			goto st7
		}
		goto st0
	tr18:
//line ./op-grammar.rl:57
		e_sgn = 0
		e_val = 0
		e_frac = 0

//line ./op-grammar.rl:66
		if (frame.Body)[p] == '-' {
			atoms[atm][1] |= uint64(1) << 32
		}

		goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
//line dfa.go:3546
		if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
			goto tr20
		}
		goto st0
	tr19:
//line ./op-grammar.rl:57
		e_sgn = 0
		e_val = 0
		e_frac = 0

//line ./op-grammar.rl:62
		atoms[atm][0] *= 10
		atoms[atm][0] += (uint64)((frame.Body)[p] - '0')

		goto st9
	tr20:
//line ./op-grammar.rl:62
		atoms[atm][0] *= 10
		atoms[atm][0] += (uint64)((frame.Body)[p] - '0')

		goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
//line dfa.go:3576
		if (frame.Body)[p] == 46 {
			goto st10
		}
		if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
			goto tr20
		}
		goto st0
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
			goto tr22
		}
		goto st0
	tr22:
//line ./op-grammar.rl:71
		atoms[atm][0] *= 10
		atoms[atm][0] += (uint64)((frame.Body)[p] - '0')
		e_frac++

		goto st25
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
//line dfa.go:3606
		switch (frame.Body)[p] {
		case 32:
			goto tr95
		case 33:
			goto tr96
		case 35:
			goto tr97
		case 39:
			goto tr98
		case 42:
			goto tr97
		case 44:
			goto tr96
		case 46:
			goto tr99
		case 58:
			goto tr97
		case 59:
			goto tr96
		case 61:
			goto tr100
		case 62:
			goto tr101
		case 63:
			goto tr96
		case 64:
			goto tr97
		case 69:
			goto st11
		case 94:
			goto tr103
		case 101:
			goto st11
		}
		switch {
		case (frame.Body)[p] > 13:
			if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
				goto tr22
			}
		case (frame.Body)[p] >= 9:
			goto tr95
		}
		goto st0
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		switch (frame.Body)[p] {
		case 43:
			goto tr23
		case 45:
			goto tr23
		}
		if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
			goto tr24
		}
		goto st0
	tr23:
//line ./op-grammar.rl:76
		if (frame.Body)[p] == '-' {
			e_sgn = -1
		}

		goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
//line dfa.go:3678
		if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
			goto tr24
		}
		goto st0
	tr24:
//line ./op-grammar.rl:81
		e_val *= 10
		e_val += int((frame.Body)[p] - '0')

		goto st26
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
//line dfa.go:3695
		switch (frame.Body)[p] {
		case 32:
			goto tr95
		case 33:
			goto tr96
		case 35:
			goto tr97
		case 39:
			goto tr98
		case 42:
			goto tr97
		case 44:
			goto tr96
		case 46:
			goto tr99
		case 58:
			goto tr97
		case 59:
			goto tr96
		case 61:
			goto tr100
		case 62:
			goto tr101
		case 63:
			goto tr96
		case 64:
			goto tr97
		case 94:
			goto tr103
		}
		switch {
		case (frame.Body)[p] > 13:
			if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
				goto tr24
			}
		case (frame.Body)[p] >= 9:
			goto tr95
		}
		goto st0
	tr15:
//line ./op-grammar.rl:107
		if atm == 4 {
			atoms[atm] = atoms[SPEC_OBJECT]
		} else if atoms[atm-1].Type() == ATOM_UUID {
			atoms[atm] = atoms[atm-1]
		}

//line ././uuid-grammar.rl:5

//line ././uuid-grammar.rl:26

//line ././uuid-grammar.rl:9
		dgt = int(ABC[(frame.Body)[p]]) + 4
		atoms[atm][hlf] &= INT60_FLAGS | PREFIX_MASKS[dgt]

		goto st27
	tr111:
//line ././uuid-grammar.rl:14
		atoms[atm][hlf] |= ((uint64)(ABC[(frame.Body)[p]])) << DIGIT_OFFSETS[dgt]
		dgt++
		if dgt > 10 {
			{
				p++
				(ps.state) = 27
				goto _out
			}
		}

		goto st27
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
//line dfa.go:3771
		switch (frame.Body)[p] {
		case 32:
			goto tr104
		case 33:
			goto tr105
		case 35:
			goto tr106
		case 39:
			goto tr108
		case 42:
			goto tr106
		case 44:
			goto tr105
		case 46:
			goto tr110
		case 58:
			goto tr106
		case 59:
			goto tr105
		case 61:
			goto tr112
		case 62:
			goto tr113
		case 63:
			goto tr105
		case 64:
			goto tr106
		case 91:
			goto tr109
		case 93:
			goto tr109
		case 94:
			goto tr114
		case 95:
			goto tr111
		case 123:
			goto tr109
		case 125:
			goto tr109
		case 126:
			goto tr111
		}
		switch {
		case (frame.Body)[p] < 43:
			switch {
			case (frame.Body)[p] < 36:
				if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
					goto tr104
				}
			case (frame.Body)[p] > 37:
				if 40 <= (frame.Body)[p] && (frame.Body)[p] <= 41 {
					goto tr109
				}
			default:
				goto tr107
			}
		case (frame.Body)[p] > 45:
			switch {
			case (frame.Body)[p] < 65:
				if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
					goto tr111
				}
			case (frame.Body)[p] > 90:
				if 97 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr111
				}
			default:
				goto tr111
			}
		default:
			goto tr107
		}
		goto st0
	tr16:
//line ./op-grammar.rl:107
		if atm == 4 {
			atoms[atm] = atoms[SPEC_OBJECT]
		} else if atoms[atm-1].Type() == ATOM_UUID {
			atoms[atm] = atoms[atm-1]
		}

//line ././uuid-grammar.rl:5

//line ././uuid-grammar.rl:26

//line ././uuid-grammar.rl:22
		atoms[atm][hlf] &= INT60_FLAGS

//line ././uuid-grammar.rl:14
		atoms[atm][hlf] |= ((uint64)(ABC[(frame.Body)[p]])) << DIGIT_OFFSETS[dgt]
		dgt++
		if dgt > 10 {
			{
				p++
				(ps.state) = 28
				goto _out
			}
		}

		goto st28
	tr120:
//line ././uuid-grammar.rl:14
		atoms[atm][hlf] |= ((uint64)(ABC[(frame.Body)[p]])) << DIGIT_OFFSETS[dgt]
		dgt++
		if dgt > 10 {
			{
				p++
				(ps.state) = 28
				goto _out
			}
		}

		goto st28
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
//line dfa.go:3888
		switch (frame.Body)[p] {
		case 32:
			goto tr115
		case 33:
			goto tr116
		case 35:
			goto tr117
		case 39:
			goto tr118
		case 42:
			goto tr117
		case 44:
			goto tr116
		case 46:
			goto tr119
		case 58:
			goto tr117
		case 59:
			goto tr116
		case 61:
			goto tr121
		case 62:
			goto tr122
		case 63:
			goto tr116
		case 64:
			goto tr117
		case 91:
			goto tr109
		case 93:
			goto tr109
		case 94:
			goto tr123
		case 95:
			goto tr120
		case 123:
			goto tr109
		case 125:
			goto tr109
		case 126:
			goto tr120
		}
		switch {
		case (frame.Body)[p] < 43:
			switch {
			case (frame.Body)[p] < 36:
				if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
					goto tr115
				}
			case (frame.Body)[p] > 37:
				if 40 <= (frame.Body)[p] && (frame.Body)[p] <= 41 {
					goto tr109
				}
			default:
				goto tr107
			}
		case (frame.Body)[p] > 45:
			switch {
			case (frame.Body)[p] < 65:
				if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
					goto tr120
				}
			case (frame.Body)[p] > 90:
				if 97 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr120
				}
			default:
				goto tr120
			}
		default:
			goto tr107
		}
		goto st0
	tr5:
//line ./op-grammar.rl:99
		atoms[atm][0] = ((uint64)(p)) << 32

		goto st13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
//line dfa.go:3973
		switch (frame.Body)[p] {
		case 10:
			goto st0
		case 13:
			goto st0
		}
		goto st3
	tr42:
//line ././uuid-grammar.rl:29
		dgt = 0
		hlf = 1

//line ././uuid-grammar.rl:9
		dgt = int(ABC[(frame.Body)[p]]) + 4
		atoms[atm][hlf] &= INT60_FLAGS | PREFIX_MASKS[dgt]

		goto st29
	tr43:
//line ././uuid-grammar.rl:29
		dgt = 0
		hlf = 1

//line ././uuid-grammar.rl:22
		atoms[atm][hlf] &= INT60_FLAGS

//line ././uuid-grammar.rl:14
		atoms[atm][hlf] |= ((uint64)(ABC[(frame.Body)[p]])) << DIGIT_OFFSETS[dgt]
		dgt++
		if dgt > 10 {
			{
				p++
				(ps.state) = 29
				goto _out
			}
		}

		goto st29
	tr129:
//line ././uuid-grammar.rl:14
		atoms[atm][hlf] |= ((uint64)(ABC[(frame.Body)[p]])) << DIGIT_OFFSETS[dgt]
		dgt++
		if dgt > 10 {
			{
				p++
				(ps.state) = 29
				goto _out
			}
		}

		goto st29
	tr138:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:29
		dgt = 0
		hlf = 1

//line ././uuid-grammar.rl:9
		dgt = int(ABC[(frame.Body)[p]]) + 4
		atoms[atm][hlf] &= INT60_FLAGS | PREFIX_MASKS[dgt]

		goto st29
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
//line dfa.go:4042
		switch (frame.Body)[p] {
		case 32:
			goto tr124
		case 33:
			goto tr125
		case 35:
			goto tr126
		case 39:
			goto tr127
		case 42:
			goto tr126
		case 44:
			goto tr125
		case 46:
			goto tr128
		case 58:
			goto tr126
		case 59:
			goto tr125
		case 61:
			goto tr130
		case 62:
			goto tr131
		case 63:
			goto tr125
		case 64:
			goto tr126
		case 94:
			goto tr132
		case 95:
			goto tr129
		case 126:
			goto tr129
		}
		switch {
		case (frame.Body)[p] < 48:
			if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
				goto tr124
			}
		case (frame.Body)[p] > 57:
			switch {
			case (frame.Body)[p] > 90:
				if 97 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr129
				}
			case (frame.Body)[p] >= 65:
				goto tr129
			}
		default:
			goto tr129
		}
		goto st0
	tr31:
//line ././uuid-grammar.rl:5

//line ././uuid-grammar.rl:26

//line ././uuid-grammar.rl:9
		dgt = int(ABC[(frame.Body)[p]]) + 4
		atoms[atm][hlf] &= INT60_FLAGS | PREFIX_MASKS[dgt]

		goto st30
	tr140:
//line ././uuid-grammar.rl:14
		atoms[atm][hlf] |= ((uint64)(ABC[(frame.Body)[p]])) << DIGIT_OFFSETS[dgt]
		dgt++
		if dgt > 10 {
			{
				p++
				(ps.state) = 30
				goto _out
			}
		}

		goto st30
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
//line dfa.go:4123
		switch (frame.Body)[p] {
		case 32:
			goto tr133
		case 33:
			goto tr134
		case 35:
			goto tr135
		case 39:
			goto tr137
		case 42:
			goto tr135
		case 44:
			goto tr134
		case 46:
			goto tr139
		case 58:
			goto tr135
		case 59:
			goto tr134
		case 61:
			goto tr141
		case 62:
			goto tr142
		case 63:
			goto tr134
		case 64:
			goto tr135
		case 91:
			goto tr138
		case 93:
			goto tr138
		case 94:
			goto tr143
		case 95:
			goto tr140
		case 123:
			goto tr138
		case 125:
			goto tr138
		case 126:
			goto tr140
		}
		switch {
		case (frame.Body)[p] < 43:
			switch {
			case (frame.Body)[p] < 36:
				if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
					goto tr133
				}
			case (frame.Body)[p] > 37:
				if 40 <= (frame.Body)[p] && (frame.Body)[p] <= 41 {
					goto tr138
				}
			default:
				goto tr136
			}
		case (frame.Body)[p] > 45:
			switch {
			case (frame.Body)[p] < 65:
				if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
					goto tr140
				}
			case (frame.Body)[p] > 90:
				if 97 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr140
				}
			default:
				goto tr140
			}
		default:
			goto tr136
		}
		goto st0
	tr33:
//line ././uuid-grammar.rl:5

//line ././uuid-grammar.rl:26

//line ././uuid-grammar.rl:22
		atoms[atm][hlf] &= INT60_FLAGS

//line ././uuid-grammar.rl:14
		atoms[atm][hlf] |= ((uint64)(ABC[(frame.Body)[p]])) << DIGIT_OFFSETS[dgt]
		dgt++
		if dgt > 10 {
			{
				p++
				(ps.state) = 31
				goto _out
			}
		}

		goto st31
	tr149:
//line ././uuid-grammar.rl:14
		atoms[atm][hlf] |= ((uint64)(ABC[(frame.Body)[p]])) << DIGIT_OFFSETS[dgt]
		dgt++
		if dgt > 10 {
			{
				p++
				(ps.state) = 31
				goto _out
			}
		}

		goto st31
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
//line dfa.go:4232
		switch (frame.Body)[p] {
		case 32:
			goto tr144
		case 33:
			goto tr145
		case 35:
			goto tr146
		case 39:
			goto tr147
		case 42:
			goto tr146
		case 44:
			goto tr145
		case 46:
			goto tr148
		case 58:
			goto tr146
		case 59:
			goto tr145
		case 61:
			goto tr150
		case 62:
			goto tr151
		case 63:
			goto tr145
		case 64:
			goto tr146
		case 91:
			goto tr138
		case 93:
			goto tr138
		case 94:
			goto tr152
		case 95:
			goto tr149
		case 123:
			goto tr138
		case 125:
			goto tr138
		case 126:
			goto tr149
		}
		switch {
		case (frame.Body)[p] < 43:
			switch {
			case (frame.Body)[p] < 36:
				if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
					goto tr144
				}
			case (frame.Body)[p] > 37:
				if 40 <= (frame.Body)[p] && (frame.Body)[p] <= 41 {
					goto tr138
				}
			default:
				goto tr136
			}
		case (frame.Body)[p] > 45:
			switch {
			case (frame.Body)[p] < 65:
				if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
					goto tr149
				}
			case (frame.Body)[p] > 90:
				if 97 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr149
				}
			default:
				goto tr149
			}
		default:
			goto tr136
		}
		goto st0
	tr37:
//line ./op-grammar.rl:6
		if atm > 0 {
			atoms[atm] = atoms[atm-1]
		}

		goto st32
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
//line dfa.go:4319
		switch (frame.Body)[p] {
		case 32:
			goto tr41
		case 33:
			goto tr27
		case 35:
			goto tr28
		case 39:
			goto tr30
		case 42:
			goto tr28
		case 44:
			goto tr27
		case 46:
			goto tr32
		case 58:
			goto tr28
		case 59:
			goto tr27
		case 61:
			goto tr34
		case 62:
			goto tr35
		case 63:
			goto tr27
		case 64:
			goto tr28
		case 91:
			goto tr31
		case 93:
			goto tr31
		case 94:
			goto tr36
		case 95:
			goto tr33
		case 123:
			goto tr31
		case 125:
			goto tr31
		case 126:
			goto tr33
		}
		switch {
		case (frame.Body)[p] < 43:
			switch {
			case (frame.Body)[p] < 36:
				if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
					goto tr41
				}
			case (frame.Body)[p] > 37:
				if 40 <= (frame.Body)[p] && (frame.Body)[p] <= 41 {
					goto tr31
				}
			default:
				goto tr29
			}
		case (frame.Body)[p] > 45:
			switch {
			case (frame.Body)[p] < 65:
				if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
					goto tr33
				}
			case (frame.Body)[p] > 90:
				if 97 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr33
				}
			default:
				goto tr33
			}
		default:
			goto tr29
		}
		goto st0
	st_out:
	_test_eof14:
		(ps.state) = 14
		goto _test_eof
	_test_eof1:
		(ps.state) = 1
		goto _test_eof
	_test_eof15:
		(ps.state) = 15
		goto _test_eof
	_test_eof16:
		(ps.state) = 16
		goto _test_eof
	_test_eof17:
		(ps.state) = 17
		goto _test_eof
	_test_eof18:
		(ps.state) = 18
		goto _test_eof
	_test_eof19:
		(ps.state) = 19
		goto _test_eof
	_test_eof2:
		(ps.state) = 2
		goto _test_eof
	_test_eof3:
		(ps.state) = 3
		goto _test_eof
	_test_eof20:
		(ps.state) = 20
		goto _test_eof
	_test_eof21:
		(ps.state) = 21
		goto _test_eof
	_test_eof4:
		(ps.state) = 4
		goto _test_eof
	_test_eof5:
		(ps.state) = 5
		goto _test_eof
	_test_eof22:
		(ps.state) = 22
		goto _test_eof
	_test_eof6:
		(ps.state) = 6
		goto _test_eof
	_test_eof23:
		(ps.state) = 23
		goto _test_eof
	_test_eof24:
		(ps.state) = 24
		goto _test_eof
	_test_eof7:
		(ps.state) = 7
		goto _test_eof
	_test_eof8:
		(ps.state) = 8
		goto _test_eof
	_test_eof9:
		(ps.state) = 9
		goto _test_eof
	_test_eof10:
		(ps.state) = 10
		goto _test_eof
	_test_eof25:
		(ps.state) = 25
		goto _test_eof
	_test_eof11:
		(ps.state) = 11
		goto _test_eof
	_test_eof12:
		(ps.state) = 12
		goto _test_eof
	_test_eof26:
		(ps.state) = 26
		goto _test_eof
	_test_eof27:
		(ps.state) = 27
		goto _test_eof
	_test_eof28:
		(ps.state) = 28
		goto _test_eof
	_test_eof13:
		(ps.state) = 13
		goto _test_eof
	_test_eof29:
		(ps.state) = 29
		goto _test_eof
	_test_eof30:
		(ps.state) = 30
		goto _test_eof
	_test_eof31:
		(ps.state) = 31
		goto _test_eof
	_test_eof32:
		(ps.state) = 32
		goto _test_eof

	_test_eof:
		{
		}
		if p == eof {
			switch ps.state {
			case 16:
//line ./op-grammar.rl:144
				frame.Position++

			case 21:
//line ./op-grammar.rl:124
//line ./op-grammar.rl:144
				frame.Position++

			case 19:
//line ./op-grammar.rl:148
				if frame.term != TERM_RAW {
					frame.term = TERM_REDUCED
				}

//line ./op-grammar.rl:144
				frame.Position++

			case 15, 18, 32:
//line ./op-grammar.rl:28
				atm++

//line ./op-grammar.rl:148
				if frame.term != TERM_RAW {
					frame.term = TERM_REDUCED
				}

//line ./op-grammar.rl:144
				frame.Position++

			case 20:
//line ./op-grammar.rl:37
				// TODO max size for int/float/string
				atm++

//line ./op-grammar.rl:124

//line ./op-grammar.rl:144
				frame.Position++

			case 30:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:28
				atm++

//line ./op-grammar.rl:148
				if frame.term != TERM_RAW {
					frame.term = TERM_REDUCED
				}

//line ./op-grammar.rl:144
				frame.Position++

			case 29:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:28
				atm++

//line ./op-grammar.rl:148
				if frame.term != TERM_RAW {
					frame.term = TERM_REDUCED
				}

//line ./op-grammar.rl:144
				frame.Position++

			case 22:
//line ./op-grammar.rl:53
				atoms[atm][1] |= ((uint64)(ATOM_INT)) << 62

//line ./op-grammar.rl:37
				// TODO max size for int/float/string
				atm++

//line ./op-grammar.rl:124

//line ./op-grammar.rl:144
				frame.Position++

			case 25, 26:
//line ./op-grammar.rl:85
				if e_sgn == -1 {
					e_val = -e_val - e_frac
				} else {
					e_val = +e_val - e_frac
				}
				if e_val < 0 {
					atoms[atm][1] |= uint64(1) << 33
					e_val = -e_val
				}
				atoms[atm][1] |= uint64(e_val)
				atoms[atm][1] |= ((uint64)(ATOM_FLOAT)) << 62

//line ./op-grammar.rl:37
				// TODO max size for int/float/string
				atm++

//line ./op-grammar.rl:124

//line ./op-grammar.rl:144
				frame.Position++

			case 23:
//line ./op-grammar.rl:114
				atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:37
				// TODO max size for int/float/string
				atm++

//line ./op-grammar.rl:124

//line ./op-grammar.rl:144
				frame.Position++

			case 31:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
				atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:28
				atm++

//line ./op-grammar.rl:148
				if frame.term != TERM_RAW {
					frame.term = TERM_REDUCED
				}

//line ./op-grammar.rl:144
				frame.Position++

			case 27:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:114
				atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:37
				// TODO max size for int/float/string
				atm++

//line ./op-grammar.rl:124

//line ./op-grammar.rl:144
				frame.Position++

			case 24:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:114
				atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:37
				// TODO max size for int/float/string
				atm++

//line ./op-grammar.rl:124

//line ./op-grammar.rl:144
				frame.Position++

			case 28:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
				atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:114
				atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:37
				// TODO max size for int/float/string
				atm++

//line ./op-grammar.rl:124

//line ./op-grammar.rl:144
				frame.Position++

//line dfa.go:4665
			}
		}

	_out:
		{
		}
	}

//line dfa.rl:72

	if atoms[SPEC_TYPE] == Atom(COMMENT_UUID) && frame.IsComplete() {
		atoms[SPEC_OBJECT] = Atom(ZERO_UUID)
		atoms[SPEC_EVENT] = Atom(ZERO_UUID)
		atoms[SPEC_REF] = Atom(ZERO_UUID)
	}

	ps.atm, ps.hlf, ps.dgt = atm, hlf, dgt
	ps.position = p
	frame.atoms = atoms

	switch {
	case ps.state == RON_error:
		frame.Position = -1
	case ps.state >= RON_first_final: // one of end states
		if !ps.streaming && p >= eof {
			// in the block mode, the final dot is optional/implied
			ps.state = RON_FULL_STOP
		}
	case ps.state == RON_FULL_STOP:
	case ps.state == RON_start:
	default:
		if !ps.streaming {
			ps.state = RON_error
			frame.Position = -1
		}
	}

}

func (ctx_uuid UUID) Parse(data []byte) (UUID, error) {

//line dfa.rl:107
//line dfa.go:4710
	const UUID_start int = 1
	const UUID_first_final int = 2
	const UUID_error int = 0

	const UUID_en_main int = 1

//line dfa.rl:108
	cs, p, pe, eof := 0, 0, len(data), len(data)
	_ = eof

	atm, hlf, dgt := 0, 0, 0

	atoms := [1]Atom{Atom(ctx_uuid)}

//line dfa.go:4728
	{
		cs = UUID_start
	}

//line dfa.go:4733
	{
		if p == pe {
			goto _test_eof
		}
		switch cs {
		case 1:
			goto st_case_1
		case 0:
			goto st_case_0
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		case 4:
			goto st_case_4
		case 5:
			goto st_case_5
		}
		goto st_out
	st_case_1:
		switch data[p] {
		case 43:
			goto tr0
		case 45:
			goto tr0
		case 91:
			goto tr2
		case 93:
			goto tr2
		case 95:
			goto tr3
		case 123:
			goto tr2
		case 125:
			goto tr2
		case 126:
			goto tr3
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 37:
				if 40 <= data[p] && data[p] <= 41 {
					goto tr2
				}
			case data[p] >= 36:
				goto tr0
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr3
				}
			case data[p] >= 65:
				goto tr3
			}
		default:
			goto tr3
		}
		goto st0
	st_case_0:
	st0:
		cs = 0
		goto _out
	tr0:
//line ./uuid-grammar.rl:5

//line ./uuid-grammar.rl:40
		hlf = 1
		atoms[atm][1] &= INT60_FULL
		atoms[atm][1] |= ((uint64)(ABC[data[p]])) << 60

		goto st2
	tr7:
//line ./uuid-grammar.rl:34
//line ./uuid-grammar.rl:40
		hlf = 1
		atoms[atm][1] &= INT60_FULL
		atoms[atm][1] |= ((uint64)(ABC[data[p]])) << 60

		goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
//line dfa.go:4826
		switch data[p] {
		case 91:
			goto tr4
		case 93:
			goto tr4
		case 95:
			goto tr5
		case 123:
			goto tr4
		case 125:
			goto tr4
		case 126:
			goto tr5
		}
		switch {
		case data[p] < 48:
			if 40 <= data[p] && data[p] <= 41 {
				goto tr4
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr5
				}
			case data[p] >= 65:
				goto tr5
			}
		default:
			goto tr5
		}
		goto st0
	tr4:
//line ./uuid-grammar.rl:29
		dgt = 0
		hlf = 1

//line ./uuid-grammar.rl:9
		dgt = int(ABC[data[p]]) + 4
		atoms[atm][hlf] &= INT60_FLAGS | PREFIX_MASKS[dgt]

		goto st3
	tr5:
//line ./uuid-grammar.rl:29
		dgt = 0
		hlf = 1

//line ./uuid-grammar.rl:22
		atoms[atm][hlf] &= INT60_FLAGS

//line ./uuid-grammar.rl:14
		atoms[atm][hlf] |= ((uint64)(ABC[data[p]])) << DIGIT_OFFSETS[dgt]
		dgt++
		if dgt > 10 {
			{
				p++
				cs = 3
				goto _out
			}
		}

		goto st3
	tr6:
//line ./uuid-grammar.rl:14
		atoms[atm][hlf] |= ((uint64)(ABC[data[p]])) << DIGIT_OFFSETS[dgt]
		dgt++
		if dgt > 10 {
			{
				p++
				cs = 3
				goto _out
			}
		}

		goto st3
	tr8:
//line ./uuid-grammar.rl:34

//line ./uuid-grammar.rl:29
		dgt = 0
		hlf = 1

//line ./uuid-grammar.rl:9
		dgt = int(ABC[data[p]]) + 4
		atoms[atm][hlf] &= INT60_FLAGS | PREFIX_MASKS[dgt]

		goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
//line dfa.go:4920
		switch data[p] {
		case 95:
			goto tr6
		case 126:
			goto tr6
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr6
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr6
			}
		default:
			goto tr6
		}
		goto st0
	tr2:
//line ./uuid-grammar.rl:5

//line ./uuid-grammar.rl:26

//line ./uuid-grammar.rl:9
		dgt = int(ABC[data[p]]) + 4
		atoms[atm][hlf] &= INT60_FLAGS | PREFIX_MASKS[dgt]

		goto st4
	tr9:
//line ./uuid-grammar.rl:14
		atoms[atm][hlf] |= ((uint64)(ABC[data[p]])) << DIGIT_OFFSETS[dgt]
		dgt++
		if dgt > 10 {
			{
				p++
				cs = 4
				goto _out
			}
		}

		goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
//line dfa.go:4968
		switch data[p] {
		case 43:
			goto tr7
		case 45:
			goto tr7
		case 91:
			goto tr8
		case 93:
			goto tr8
		case 95:
			goto tr9
		case 123:
			goto tr8
		case 125:
			goto tr8
		case 126:
			goto tr9
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 37:
				if 40 <= data[p] && data[p] <= 41 {
					goto tr8
				}
			case data[p] >= 36:
				goto tr7
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr9
				}
			case data[p] >= 65:
				goto tr9
			}
		default:
			goto tr9
		}
		goto st0
	tr3:
//line ./uuid-grammar.rl:5

//line ./uuid-grammar.rl:26

//line ./uuid-grammar.rl:22
		atoms[atm][hlf] &= INT60_FLAGS

//line ./uuid-grammar.rl:14
		atoms[atm][hlf] |= ((uint64)(ABC[data[p]])) << DIGIT_OFFSETS[dgt]
		dgt++
		if dgt > 10 {
			{
				p++
				cs = 5
				goto _out
			}
		}

		goto st5
	tr10:
//line ./uuid-grammar.rl:14
		atoms[atm][hlf] |= ((uint64)(ABC[data[p]])) << DIGIT_OFFSETS[dgt]
		dgt++
		if dgt > 10 {
			{
				p++
				cs = 5
				goto _out
			}
		}

		goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
//line dfa.go:5045
		switch data[p] {
		case 43:
			goto tr7
		case 45:
			goto tr7
		case 91:
			goto tr8
		case 93:
			goto tr8
		case 95:
			goto tr10
		case 123:
			goto tr8
		case 125:
			goto tr8
		case 126:
			goto tr10
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 37:
				if 40 <= data[p] && data[p] <= 41 {
					goto tr8
				}
			case data[p] >= 36:
				goto tr7
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr10
				}
			case data[p] >= 65:
				goto tr10
			}
		default:
			goto tr10
		}
		goto st0
	st_out:
	_test_eof2:
		cs = 2
		goto _test_eof
	_test_eof3:
		cs = 3
		goto _test_eof
	_test_eof4:
		cs = 4
		goto _test_eof
	_test_eof5:
		cs = 5
		goto _test_eof

	_test_eof:
		{
		}
		if p == eof {
			switch cs {
			case 4:
//line ./uuid-grammar.rl:34

			case 3:
//line ./uuid-grammar.rl:37

			case 5:
//line ./uuid-grammar.rl:34
//line ./uuid-grammar.rl:46
				atoms[atm][1] = UUID_NAME_FLAG

//line dfa.go:5112
			}
		}

	_out:
		{
		}
	}

//line dfa.rl:123

	if cs < UUID_first_final || dgt > 10 {
		return ERROR_UUID, errors.New(fmt.Sprintf("parse error at pos %d", p))
	} else {
		return UUID(atoms[0]), nil
	}

}
