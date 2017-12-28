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
		case 2:
			goto st2
		case 3:
			goto st3
		case 19:
			goto st19
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
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		case 19:
			goto st_case_19
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
		case 33:
			goto tr2
		case 35:
			goto tr3
		case 39:
			goto tr4
		case 42:
			goto tr3
		case 44:
			goto tr2
		case 46:
			goto tr30
		case 58:
			goto tr3
		case 59:
			goto tr2
		case 61:
			goto tr5
		case 62:
			goto tr6
		case 63:
			goto tr2
		case 64:
			goto tr3
		case 94:
			goto tr7
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
		case 33:
			goto tr2
		case 35:
			goto tr3
		case 39:
			goto tr4
		case 42:
			goto tr3
		case 44:
			goto tr2
		case 58:
			goto tr3
		case 59:
			goto tr2
		case 61:
			goto tr5
		case 62:
			goto tr6
		case 63:
			goto tr2
		case 64:
			goto tr3
		case 94:
			goto tr7
		}
		if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
			goto st1
		}
		goto st0
	tr2:
		(ps.state) = 15
//line ./op-grammar.rl:132
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
			if frame.term != TERM_RAW {
				frame.term = TERM_REDUCED
			}
		}

//line ./op-grammar.rl:128
		frame.term = int(ABC[(frame.Body)[p]])

		goto _again
	tr32:
		(ps.state) = 15
//line ./op-grammar.rl:148
		frame.Position++

//line ./op-grammar.rl:132
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
			if frame.term != TERM_RAW {
				frame.term = TERM_REDUCED
			}
		}

//line ./op-grammar.rl:128
		frame.term = int(ABC[(frame.Body)[p]])

		goto _again
	tr40:
//line ./op-grammar.rl:27
		atm++

//line ./op-grammar.rl:152

//line ./op-grammar.rl:128
		frame.term = int(ABC[(frame.Body)[p]])

		goto st15
	tr55:
//line ./op-grammar.rl:152

//line ./op-grammar.rl:128
		frame.term = int(ABC[(frame.Body)[p]])

		goto st15
	tr63:
//line ./op-grammar.rl:35
		atm++

//line ./op-grammar.rl:125

//line ./op-grammar.rl:128
		frame.term = int(ABC[(frame.Body)[p]])

		goto st15
	tr71:
//line ./op-grammar.rl:125
//line ./op-grammar.rl:128
		frame.term = int(ABC[(frame.Body)[p]])

		goto st15
	tr79:
//line ./op-grammar.rl:51
		atoms[atm][1] |= ((uint64)(ATOM_INT)) << 62

//line ./op-grammar.rl:35
		atm++

//line ./op-grammar.rl:125

//line ./op-grammar.rl:128
		frame.term = int(ABC[(frame.Body)[p]])

		goto st15
	tr87:
//line ./op-grammar.rl:115
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:35
		atm++

//line ./op-grammar.rl:125

//line ./op-grammar.rl:128
		frame.term = int(ABC[(frame.Body)[p]])

		goto st15
	tr97:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:115
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:35
		atm++

//line ./op-grammar.rl:125

//line ./op-grammar.rl:128
		frame.term = int(ABC[(frame.Body)[p]])

		goto st15
	tr106:
//line ./op-grammar.rl:86
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

//line ./op-grammar.rl:35
		atm++

//line ./op-grammar.rl:125

//line ./op-grammar.rl:128
		frame.term = int(ABC[(frame.Body)[p]])

		goto st15
	tr115:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:115
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:35
		atm++

//line ./op-grammar.rl:125

//line ./op-grammar.rl:128
		frame.term = int(ABC[(frame.Body)[p]])

		goto st15
	tr126:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:115
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:35
		atm++

//line ./op-grammar.rl:125

//line ./op-grammar.rl:128
		frame.term = int(ABC[(frame.Body)[p]])

		goto st15
	tr135:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:27
		atm++

//line ./op-grammar.rl:152

//line ./op-grammar.rl:128
		frame.term = int(ABC[(frame.Body)[p]])

		goto st15
	tr144:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:27
		atm++

//line ./op-grammar.rl:152

//line ./op-grammar.rl:128
		frame.term = int(ABC[(frame.Body)[p]])

		goto st15
	tr155:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:27
		atm++

//line ./op-grammar.rl:152

//line ./op-grammar.rl:128
		frame.term = int(ABC[(frame.Body)[p]])

		goto st15
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
//line dfa.go:595
		switch (frame.Body)[p] {
		case 32:
			goto st15
		case 33:
			goto tr32
		case 35:
			goto tr33
		case 39:
			goto tr34
		case 42:
			goto tr33
		case 44:
			goto tr32
		case 46:
			goto tr35
		case 58:
			goto tr33
		case 59:
			goto tr32
		case 61:
			goto tr36
		case 62:
			goto tr37
		case 63:
			goto tr32
		case 64:
			goto tr33
		case 94:
			goto tr38
		}
		if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
			goto st15
		}
		goto st0
	tr3:
		(ps.state) = 16
//line ./op-grammar.rl:132
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
			if frame.term != TERM_RAW {
				frame.term = TERM_REDUCED
			}
		}

//line ./op-grammar.rl:12
		n = (int)(ABC[(frame.Body)[p]])
		hlf, dgt = 0, 0
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
	tr33:
		(ps.state) = 16
//line ./op-grammar.rl:148
		frame.Position++

//line ./op-grammar.rl:132
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
			if frame.term != TERM_RAW {
				frame.term = TERM_REDUCED
			}
		}

//line ./op-grammar.rl:12
		n = (int)(ABC[(frame.Body)[p]])
		hlf, dgt = 0, 0
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
//line ./op-grammar.rl:27
		atm++

		goto st16
	tr41:
		(ps.state) = 16
//line ./op-grammar.rl:27
		atm++

//line ./op-grammar.rl:12
		n = (int)(ABC[(frame.Body)[p]])
		hlf, dgt = 0, 0
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
	tr56:
		(ps.state) = 16
//line ./op-grammar.rl:12
		n = (int)(ABC[(frame.Body)[p]])
		hlf, dgt = 0, 0
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
	tr64:
		(ps.state) = 16
//line ./op-grammar.rl:35
		atm++

//line ./op-grammar.rl:125

//line ./op-grammar.rl:148
		frame.Position++

//line ./op-grammar.rl:132
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
			if frame.term != TERM_RAW {
				frame.term = TERM_REDUCED
			}
		}

//line ./op-grammar.rl:12
		n = (int)(ABC[(frame.Body)[p]])
		hlf, dgt = 0, 0
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
	tr72:
		(ps.state) = 16
//line ./op-grammar.rl:125

//line ./op-grammar.rl:148
		frame.Position++

//line ./op-grammar.rl:132
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
			if frame.term != TERM_RAW {
				frame.term = TERM_REDUCED
			}
		}

//line ./op-grammar.rl:12
		n = (int)(ABC[(frame.Body)[p]])
		hlf, dgt = 0, 0
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
	tr80:
		(ps.state) = 16
//line ./op-grammar.rl:51
		atoms[atm][1] |= ((uint64)(ATOM_INT)) << 62

//line ./op-grammar.rl:35
		atm++

//line ./op-grammar.rl:125

//line ./op-grammar.rl:148
		frame.Position++

//line ./op-grammar.rl:132
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
			if frame.term != TERM_RAW {
				frame.term = TERM_REDUCED
			}
		}

//line ./op-grammar.rl:12
		n = (int)(ABC[(frame.Body)[p]])
		hlf, dgt = 0, 0
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
		(ps.state) = 16
//line ./op-grammar.rl:115
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:35
		atm++

//line ./op-grammar.rl:125

//line ./op-grammar.rl:148
		frame.Position++

//line ./op-grammar.rl:132
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
			if frame.term != TERM_RAW {
				frame.term = TERM_REDUCED
			}
		}

//line ./op-grammar.rl:12
		n = (int)(ABC[(frame.Body)[p]])
		hlf, dgt = 0, 0
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
	tr98:
		(ps.state) = 16
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:115
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:35
		atm++

//line ./op-grammar.rl:125

//line ./op-grammar.rl:148
		frame.Position++

//line ./op-grammar.rl:132
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
			if frame.term != TERM_RAW {
				frame.term = TERM_REDUCED
			}
		}

//line ./op-grammar.rl:12
		n = (int)(ABC[(frame.Body)[p]])
		hlf, dgt = 0, 0
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
	tr107:
		(ps.state) = 16
//line ./op-grammar.rl:86
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

//line ./op-grammar.rl:35
		atm++

//line ./op-grammar.rl:125

//line ./op-grammar.rl:148
		frame.Position++

//line ./op-grammar.rl:132
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
			if frame.term != TERM_RAW {
				frame.term = TERM_REDUCED
			}
		}

//line ./op-grammar.rl:12
		n = (int)(ABC[(frame.Body)[p]])
		hlf, dgt = 0, 0
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
	tr116:
		(ps.state) = 16
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:115
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:35
		atm++

//line ./op-grammar.rl:125

//line ./op-grammar.rl:148
		frame.Position++

//line ./op-grammar.rl:132
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
			if frame.term != TERM_RAW {
				frame.term = TERM_REDUCED
			}
		}

//line ./op-grammar.rl:12
		n = (int)(ABC[(frame.Body)[p]])
		hlf, dgt = 0, 0
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
	tr127:
		(ps.state) = 16
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:115
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:35
		atm++

//line ./op-grammar.rl:125

//line ./op-grammar.rl:148
		frame.Position++

//line ./op-grammar.rl:132
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
			if frame.term != TERM_RAW {
				frame.term = TERM_REDUCED
			}
		}

//line ./op-grammar.rl:12
		n = (int)(ABC[(frame.Body)[p]])
		hlf, dgt = 0, 0
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
	tr136:
		(ps.state) = 16
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:27
		atm++

//line ./op-grammar.rl:12
		n = (int)(ABC[(frame.Body)[p]])
		hlf, dgt = 0, 0
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
	tr145:
		(ps.state) = 16
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:27
		atm++

//line ./op-grammar.rl:12
		n = (int)(ABC[(frame.Body)[p]])
		hlf, dgt = 0, 0
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
	tr156:
		(ps.state) = 16
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:27
		atm++

//line ./op-grammar.rl:12
		n = (int)(ABC[(frame.Body)[p]])
		hlf, dgt = 0, 0
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
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
//line dfa.go:1235
		switch (frame.Body)[p] {
		case 32:
			goto tr39
		case 33:
			goto tr40
		case 35:
			goto tr41
		case 39:
			goto tr43
		case 42:
			goto tr41
		case 44:
			goto tr40
		case 46:
			goto tr45
		case 58:
			goto tr41
		case 59:
			goto tr40
		case 61:
			goto tr47
		case 62:
			goto tr48
		case 63:
			goto tr40
		case 64:
			goto tr41
		case 91:
			goto tr44
		case 93:
			goto tr44
		case 94:
			goto tr49
		case 96:
			goto tr50
		case 123:
			goto tr44
		case 125:
			goto tr44
		case 126:
			goto tr46
		}
		switch {
		case (frame.Body)[p] < 43:
			switch {
			case (frame.Body)[p] < 36:
				if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
					goto tr39
				}
			case (frame.Body)[p] > 37:
				if 40 <= (frame.Body)[p] && (frame.Body)[p] <= 41 {
					goto tr44
				}
			default:
				goto tr42
			}
		case (frame.Body)[p] > 45:
			switch {
			case (frame.Body)[p] < 65:
				if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
					goto tr46
				}
			case (frame.Body)[p] > 90:
				if 95 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr46
				}
			default:
				goto tr46
			}
		default:
			goto tr42
		}
		goto st0
	tr42:
//line ././uuid-grammar.rl:5

//line ././uuid-grammar.rl:40
		hlf = 1
		atoms[atm][1] &= INT60_FULL
		atoms[atm][1] |= ((uint64)(ABC[(frame.Body)[p]])) << 60

		goto st17
	tr146:
//line ././uuid-grammar.rl:34
//line ././uuid-grammar.rl:40
		hlf = 1
		atoms[atm][1] &= INT60_FULL
		atoms[atm][1] |= ((uint64)(ABC[(frame.Body)[p]])) << 60

		goto st17
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
//line dfa.go:1336
		switch (frame.Body)[p] {
		case 32:
			goto tr51
		case 33:
			goto tr40
		case 35:
			goto tr41
		case 39:
			goto tr43
		case 42:
			goto tr41
		case 44:
			goto tr40
		case 46:
			goto tr45
		case 58:
			goto tr41
		case 59:
			goto tr40
		case 61:
			goto tr47
		case 62:
			goto tr48
		case 63:
			goto tr40
		case 64:
			goto tr41
		case 91:
			goto tr52
		case 93:
			goto tr52
		case 94:
			goto tr49
		case 95:
			goto tr53
		case 123:
			goto tr52
		case 125:
			goto tr52
		case 126:
			goto tr53
		}
		switch {
		case (frame.Body)[p] < 48:
			switch {
			case (frame.Body)[p] > 13:
				if 40 <= (frame.Body)[p] && (frame.Body)[p] <= 41 {
					goto tr52
				}
			case (frame.Body)[p] >= 9:
				goto tr51
			}
		case (frame.Body)[p] > 57:
			switch {
			case (frame.Body)[p] > 90:
				if 97 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr53
				}
			case (frame.Body)[p] >= 65:
				goto tr53
			}
		default:
			goto tr53
		}
		goto st0
	tr51:
//line ./op-grammar.rl:27
		atm++

		goto st18
	tr134:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:27
		atm++

		goto st18
	tr143:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:27
		atm++

		goto st18
	tr154:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:27
		atm++

		goto st18
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
//line dfa.go:1444
		switch (frame.Body)[p] {
		case 32:
			goto st18
		case 33:
			goto tr55
		case 35:
			goto tr56
		case 39:
			goto tr57
		case 42:
			goto tr56
		case 44:
			goto tr55
		case 46:
			goto tr58
		case 58:
			goto tr56
		case 59:
			goto tr55
		case 61:
			goto tr59
		case 62:
			goto tr60
		case 63:
			goto tr55
		case 64:
			goto tr56
		case 94:
			goto tr61
		}
		if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
			goto st18
		}
		goto st0
	tr4:
		(ps.state) = 2
//line ./op-grammar.rl:132
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
			if frame.term != TERM_RAW {
				frame.term = TERM_REDUCED
			}
		}

//line ./op-grammar.rl:120
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto _again
	tr34:
		(ps.state) = 2
//line ./op-grammar.rl:148
		frame.Position++

//line ./op-grammar.rl:132
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
			if frame.term != TERM_RAW {
				frame.term = TERM_REDUCED
			}
		}

//line ./op-grammar.rl:120
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto _again
	tr43:
//line ./op-grammar.rl:27
		atm++

//line ./op-grammar.rl:152

//line ./op-grammar.rl:120
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st2
	tr57:
//line ./op-grammar.rl:152

//line ./op-grammar.rl:120
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st2
	tr65:
//line ./op-grammar.rl:35
		atm++

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st2
	tr73:
//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st2
	tr81:
//line ./op-grammar.rl:51
		atoms[atm][1] |= ((uint64)(ATOM_INT)) << 62

//line ./op-grammar.rl:35
		atm++

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st2
	tr89:
//line ./op-grammar.rl:115
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:35
		atm++

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st2
	tr99:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:115
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:35
		atm++

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st2
	tr108:
//line ./op-grammar.rl:86
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

//line ./op-grammar.rl:35
		atm++

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st2
	tr118:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:115
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:35
		atm++

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st2
	tr128:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:115
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:35
		atm++

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st2
	tr137:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:27
		atm++

//line ./op-grammar.rl:152

//line ./op-grammar.rl:120
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st2
	tr147:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:27
		atm++

//line ./op-grammar.rl:152

//line ./op-grammar.rl:120
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st2
	tr157:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:27
		atm++

//line ./op-grammar.rl:152

//line ./op-grammar.rl:120
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
//line dfa.go:1788
		switch (frame.Body)[p] {
		case 10:
			goto st0
		case 13:
			goto st0
		case 34:
			goto st0
		case 39:
			goto tr9
		case 92:
			goto tr10
		}
		goto tr8
	tr8:
//line ./op-grammar.rl:100
		atoms[atm][0] = ((uint64)(p)) << 32

		goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
//line dfa.go:1813
		switch (frame.Body)[p] {
		case 10:
			goto st0
		case 13:
			goto st0
		case 34:
			goto st0
		case 39:
			goto tr12
		case 92:
			goto st13
		}
		goto st3
	tr9:
//line ./op-grammar.rl:100
		atoms[atm][0] = ((uint64)(p)) << 32

//line ./op-grammar.rl:103
		atoms[atm][0] |= uint64(p)
		atoms[atm][1] = ((uint64)(ATOM_STRING)) << 62

		goto st19
	tr12:
//line ./op-grammar.rl:103
		atoms[atm][0] |= uint64(p)
		atoms[atm][1] = ((uint64)(ATOM_STRING)) << 62

		goto st19
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
//line dfa.go:1850
		switch (frame.Body)[p] {
		case 32:
			goto tr62
		case 33:
			goto tr63
		case 35:
			goto tr64
		case 39:
			goto tr65
		case 42:
			goto tr64
		case 44:
			goto tr63
		case 46:
			goto tr66
		case 58:
			goto tr64
		case 59:
			goto tr63
		case 61:
			goto tr67
		case 62:
			goto tr68
		case 63:
			goto tr63
		case 64:
			goto tr64
		case 94:
			goto tr69
		}
		if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
			goto tr62
		}
		goto st0
	tr62:
//line ./op-grammar.rl:35
		atm++

		goto st20
	tr78:
//line ./op-grammar.rl:51
		atoms[atm][1] |= ((uint64)(ATOM_INT)) << 62

//line ./op-grammar.rl:35
		atm++

		goto st20
	tr86:
//line ./op-grammar.rl:115
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:35
		atm++

		goto st20
	tr96:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:115
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:35
		atm++

		goto st20
	tr105:
//line ./op-grammar.rl:86
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

//line ./op-grammar.rl:35
		atm++

		goto st20
	tr114:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:115
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:35
		atm++

		goto st20
	tr125:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:115
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:35
		atm++

		goto st20
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
//line dfa.go:1979
		switch (frame.Body)[p] {
		case 32:
			goto st20
		case 33:
			goto tr71
		case 35:
			goto tr72
		case 39:
			goto tr73
		case 42:
			goto tr72
		case 44:
			goto tr71
		case 46:
			goto tr74
		case 58:
			goto tr72
		case 59:
			goto tr71
		case 61:
			goto tr75
		case 62:
			goto tr76
		case 63:
			goto tr71
		case 64:
			goto tr72
		case 94:
			goto tr77
		}
		if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
			goto st20
		}
		goto st0
	tr30:
		(ps.state) = 21
//line ./op-grammar.rl:155
		(ps.state) = (RON_FULL_STOP)

		goto _again
	tr35:
		(ps.state) = 21
//line ./op-grammar.rl:148
		frame.Position++

//line ./op-grammar.rl:155
		(ps.state) = (RON_FULL_STOP)

		goto _again
	tr45:
		(ps.state) = 21
//line ./op-grammar.rl:27
		atm++

//line ./op-grammar.rl:152

//line ./op-grammar.rl:148
		frame.Position++

//line ./op-grammar.rl:155
		(ps.state) = (RON_FULL_STOP)

		goto _again
	tr58:
		(ps.state) = 21
//line ./op-grammar.rl:152
//line ./op-grammar.rl:148
		frame.Position++

//line ./op-grammar.rl:155
		(ps.state) = (RON_FULL_STOP)

		goto _again
	tr66:
		(ps.state) = 21
//line ./op-grammar.rl:35
		atm++

//line ./op-grammar.rl:125

//line ./op-grammar.rl:148
		frame.Position++

//line ./op-grammar.rl:155
		(ps.state) = (RON_FULL_STOP)

		goto _again
	tr74:
		(ps.state) = 21
//line ./op-grammar.rl:125
//line ./op-grammar.rl:148
		frame.Position++

//line ./op-grammar.rl:155
		(ps.state) = (RON_FULL_STOP)

		goto _again
	tr82:
		(ps.state) = 21
//line ./op-grammar.rl:51
		atoms[atm][1] |= ((uint64)(ATOM_INT)) << 62

//line ./op-grammar.rl:35
		atm++

//line ./op-grammar.rl:125

//line ./op-grammar.rl:148
		frame.Position++

//line ./op-grammar.rl:155
		(ps.state) = (RON_FULL_STOP)

		goto _again
	tr91:
		(ps.state) = 21
//line ./op-grammar.rl:115
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:35
		atm++

//line ./op-grammar.rl:125

//line ./op-grammar.rl:148
		frame.Position++

//line ./op-grammar.rl:155
		(ps.state) = (RON_FULL_STOP)

		goto _again
	tr100:
		(ps.state) = 21
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:115
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:35
		atm++

//line ./op-grammar.rl:125

//line ./op-grammar.rl:148
		frame.Position++

//line ./op-grammar.rl:155
		(ps.state) = (RON_FULL_STOP)

		goto _again
	tr109:
		(ps.state) = 21
//line ./op-grammar.rl:86
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

//line ./op-grammar.rl:35
		atm++

//line ./op-grammar.rl:125

//line ./op-grammar.rl:148
		frame.Position++

//line ./op-grammar.rl:155
		(ps.state) = (RON_FULL_STOP)

		goto _again
	tr120:
		(ps.state) = 21
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:115
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:35
		atm++

//line ./op-grammar.rl:125

//line ./op-grammar.rl:148
		frame.Position++

//line ./op-grammar.rl:155
		(ps.state) = (RON_FULL_STOP)

		goto _again
	tr129:
		(ps.state) = 21
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:115
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:35
		atm++

//line ./op-grammar.rl:125

//line ./op-grammar.rl:148
		frame.Position++

//line ./op-grammar.rl:155
		(ps.state) = (RON_FULL_STOP)

		goto _again
	tr138:
		(ps.state) = 21
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:27
		atm++

//line ./op-grammar.rl:152

//line ./op-grammar.rl:148
		frame.Position++

//line ./op-grammar.rl:155
		(ps.state) = (RON_FULL_STOP)

		goto _again
	tr149:
		(ps.state) = 21
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:27
		atm++

//line ./op-grammar.rl:152

//line ./op-grammar.rl:148
		frame.Position++

//line ./op-grammar.rl:155
		(ps.state) = (RON_FULL_STOP)

		goto _again
	tr158:
		(ps.state) = 21
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:27
		atm++

//line ./op-grammar.rl:152

//line ./op-grammar.rl:148
		frame.Position++

//line ./op-grammar.rl:155
		(ps.state) = (RON_FULL_STOP)

		goto _again
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
//line dfa.go:2323
		goto st0
	tr5:
		(ps.state) = 4
//line ./op-grammar.rl:132
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
			if frame.term != TERM_RAW {
				frame.term = TERM_REDUCED
			}
		}

//line ./op-grammar.rl:120
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto _again
	tr36:
		(ps.state) = 4
//line ./op-grammar.rl:148
		frame.Position++

//line ./op-grammar.rl:132
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
			if frame.term != TERM_RAW {
				frame.term = TERM_REDUCED
			}
		}

//line ./op-grammar.rl:120
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto _again
	tr47:
//line ./op-grammar.rl:27
		atm++

//line ./op-grammar.rl:152

//line ./op-grammar.rl:120
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st4
	tr59:
//line ./op-grammar.rl:152

//line ./op-grammar.rl:120
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st4
	tr67:
//line ./op-grammar.rl:35
		atm++

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st4
	tr75:
//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st4
	tr83:
//line ./op-grammar.rl:51
		atoms[atm][1] |= ((uint64)(ATOM_INT)) << 62

//line ./op-grammar.rl:35
		atm++

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st4
	tr93:
//line ./op-grammar.rl:115
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:35
		atm++

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st4
	tr102:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:115
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:35
		atm++

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st4
	tr110:
//line ./op-grammar.rl:86
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

//line ./op-grammar.rl:35
		atm++

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st4
	tr122:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:115
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:35
		atm++

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st4
	tr131:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:115
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:35
		atm++

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st4
	tr140:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:27
		atm++

//line ./op-grammar.rl:152

//line ./op-grammar.rl:120
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st4
	tr151:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:27
		atm++

//line ./op-grammar.rl:152

//line ./op-grammar.rl:120
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st4
	tr160:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:27
		atm++

//line ./op-grammar.rl:152

//line ./op-grammar.rl:120
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
//line dfa.go:2634
		switch (frame.Body)[p] {
		case 32:
			goto st4
		case 43:
			goto tr15
		case 45:
			goto tr15
		}
		switch {
		case (frame.Body)[p] > 13:
			if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
				goto tr16
			}
		case (frame.Body)[p] >= 9:
			goto st4
		}
		goto st0
	tr15:
//line ./op-grammar.rl:39

//line ./op-grammar.rl:41
		if (frame.Body)[p] == '-' {
			atoms[atm][1] |= 1
		}

		goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
//line dfa.go:2668
		if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
			goto tr17
		}
		goto st0
	tr16:
//line ./op-grammar.rl:39

//line ./op-grammar.rl:46
		atoms[atm][0] *= 10
		atoms[atm][0] += (uint64)((frame.Body)[p] - '0')
		// TODO max size for int/float/string

		goto st22
	tr17:
//line ./op-grammar.rl:46
		atoms[atm][0] *= 10
		atoms[atm][0] += (uint64)((frame.Body)[p] - '0')
		// TODO max size for int/float/string

		goto st22
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
//line dfa.go:2697
		switch (frame.Body)[p] {
		case 32:
			goto tr78
		case 33:
			goto tr79
		case 35:
			goto tr80
		case 39:
			goto tr81
		case 42:
			goto tr80
		case 44:
			goto tr79
		case 46:
			goto tr82
		case 58:
			goto tr80
		case 59:
			goto tr79
		case 61:
			goto tr83
		case 62:
			goto tr84
		case 63:
			goto tr79
		case 64:
			goto tr80
		case 94:
			goto tr85
		}
		switch {
		case (frame.Body)[p] > 13:
			if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
				goto tr17
			}
		case (frame.Body)[p] >= 9:
			goto tr78
		}
		goto st0
	tr6:
		(ps.state) = 6
//line ./op-grammar.rl:132
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
			if frame.term != TERM_RAW {
				frame.term = TERM_REDUCED
			}
		}

//line ./op-grammar.rl:120
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto _again
	tr37:
		(ps.state) = 6
//line ./op-grammar.rl:148
		frame.Position++

//line ./op-grammar.rl:132
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
			if frame.term != TERM_RAW {
				frame.term = TERM_REDUCED
			}
		}

//line ./op-grammar.rl:120
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto _again
	tr48:
//line ./op-grammar.rl:27
		atm++

//line ./op-grammar.rl:152

//line ./op-grammar.rl:120
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st6
	tr60:
//line ./op-grammar.rl:152

//line ./op-grammar.rl:120
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st6
	tr68:
//line ./op-grammar.rl:35
		atm++

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st6
	tr76:
//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st6
	tr84:
//line ./op-grammar.rl:51
		atoms[atm][1] |= ((uint64)(ATOM_INT)) << 62

//line ./op-grammar.rl:35
		atm++

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st6
	tr94:
//line ./op-grammar.rl:115
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:35
		atm++

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st6
	tr103:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:115
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:35
		atm++

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st6
	tr111:
//line ./op-grammar.rl:86
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

//line ./op-grammar.rl:35
		atm++

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st6
	tr123:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:115
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:35
		atm++

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st6
	tr132:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:115
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:35
		atm++

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st6
	tr141:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:27
		atm++

//line ./op-grammar.rl:152

//line ./op-grammar.rl:120
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st6
	tr152:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:27
		atm++

//line ./op-grammar.rl:152

//line ./op-grammar.rl:120
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st6
	tr161:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:27
		atm++

//line ./op-grammar.rl:152

//line ./op-grammar.rl:120
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
//line dfa.go:3046
		switch (frame.Body)[p] {
		case 32:
			goto st6
		case 43:
			goto tr19
		case 45:
			goto tr19
		case 91:
			goto tr20
		case 93:
			goto tr20
		case 95:
			goto tr21
		case 123:
			goto tr20
		case 125:
			goto tr20
		case 126:
			goto tr21
		}
		switch {
		case (frame.Body)[p] < 40:
			switch {
			case (frame.Body)[p] > 13:
				if 36 <= (frame.Body)[p] && (frame.Body)[p] <= 37 {
					goto tr19
				}
			case (frame.Body)[p] >= 9:
				goto st6
			}
		case (frame.Body)[p] > 41:
			switch {
			case (frame.Body)[p] < 65:
				if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
					goto tr21
				}
			case (frame.Body)[p] > 90:
				if 97 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr21
				}
			default:
				goto tr21
			}
		default:
			goto tr20
		}
		goto st0
	tr19:
//line ./op-grammar.rl:108
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
	tr117:
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
//line dfa.go:3129
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
			goto tr91
		case 58:
			goto tr88
		case 59:
			goto tr87
		case 61:
			goto tr93
		case 62:
			goto tr94
		case 63:
			goto tr87
		case 64:
			goto tr88
		case 91:
			goto tr90
		case 93:
			goto tr90
		case 94:
			goto tr95
		case 95:
			goto tr92
		case 123:
			goto tr90
		case 125:
			goto tr90
		case 126:
			goto tr92
		}
		switch {
		case (frame.Body)[p] < 48:
			switch {
			case (frame.Body)[p] > 13:
				if 40 <= (frame.Body)[p] && (frame.Body)[p] <= 41 {
					goto tr90
				}
			case (frame.Body)[p] >= 9:
				goto tr86
			}
		case (frame.Body)[p] > 57:
			switch {
			case (frame.Body)[p] > 90:
				if 97 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr92
				}
			case (frame.Body)[p] >= 65:
				goto tr92
			}
		default:
			goto tr92
		}
		goto st0
	tr90:
//line ././uuid-grammar.rl:29
		dgt = 0
		hlf = 1

//line ././uuid-grammar.rl:9
		dgt = int(ABC[(frame.Body)[p]]) + 4
		atoms[atm][hlf] &= INT60_FLAGS | PREFIX_MASKS[dgt]

		goto st24
	tr92:
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
	tr101:
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
	tr119:
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
//line dfa.go:3256
		switch (frame.Body)[p] {
		case 32:
			goto tr96
		case 33:
			goto tr97
		case 35:
			goto tr98
		case 39:
			goto tr99
		case 42:
			goto tr98
		case 44:
			goto tr97
		case 46:
			goto tr100
		case 58:
			goto tr98
		case 59:
			goto tr97
		case 61:
			goto tr102
		case 62:
			goto tr103
		case 63:
			goto tr97
		case 64:
			goto tr98
		case 94:
			goto tr104
		case 95:
			goto tr101
		case 126:
			goto tr101
		}
		switch {
		case (frame.Body)[p] < 48:
			if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
				goto tr96
			}
		case (frame.Body)[p] > 57:
			switch {
			case (frame.Body)[p] > 90:
				if 97 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr101
				}
			case (frame.Body)[p] >= 65:
				goto tr101
			}
		default:
			goto tr101
		}
		goto st0
	tr7:
		(ps.state) = 7
//line ./op-grammar.rl:132
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
			if frame.term != TERM_RAW {
				frame.term = TERM_REDUCED
			}
		}

//line ./op-grammar.rl:120
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto _again
	tr38:
		(ps.state) = 7
//line ./op-grammar.rl:148
		frame.Position++

//line ./op-grammar.rl:132
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
			if frame.term != TERM_RAW {
				frame.term = TERM_REDUCED
			}
		}

//line ./op-grammar.rl:120
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto _again
	tr49:
//line ./op-grammar.rl:27
		atm++

//line ./op-grammar.rl:152

//line ./op-grammar.rl:120
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st7
	tr61:
//line ./op-grammar.rl:152

//line ./op-grammar.rl:120
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st7
	tr69:
//line ./op-grammar.rl:35
		atm++

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st7
	tr77:
//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st7
	tr85:
//line ./op-grammar.rl:51
		atoms[atm][1] |= ((uint64)(ATOM_INT)) << 62

//line ./op-grammar.rl:35
		atm++

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st7
	tr95:
//line ./op-grammar.rl:115
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:35
		atm++

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st7
	tr104:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:115
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:35
		atm++

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st7
	tr113:
//line ./op-grammar.rl:86
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

//line ./op-grammar.rl:35
		atm++

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st7
	tr124:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:115
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:35
		atm++

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st7
	tr133:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:115
		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:35
		atm++

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st7
	tr142:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:27
		atm++

//line ./op-grammar.rl:152

//line ./op-grammar.rl:120
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st7
	tr153:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:27
		atm++

//line ./op-grammar.rl:152

//line ./op-grammar.rl:120
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st7
	tr162:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:27
		atm++

//line ./op-grammar.rl:152

//line ./op-grammar.rl:120
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:31
		hlf, dgt = 0, 0
		atoms = append(atoms, Atom{})

		goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
//line dfa.go:3618
		switch (frame.Body)[p] {
		case 32:
			goto st7
		case 43:
			goto tr23
		case 45:
			goto tr23
		}
		switch {
		case (frame.Body)[p] > 13:
			if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
				goto tr24
			}
		case (frame.Body)[p] >= 9:
			goto st7
		}
		goto st0
	tr23:
//line ./op-grammar.rl:55
		e_sgn = 0
		e_val = 0
		e_frac = 0

//line ./op-grammar.rl:65
		if (frame.Body)[p] == '-' {
			atoms[atm][1] |= uint64(1) << 32
		}

		goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
//line dfa.go:3655
		if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
			goto tr25
		}
		goto st0
	tr24:
//line ./op-grammar.rl:55
		e_sgn = 0
		e_val = 0
		e_frac = 0

//line ./op-grammar.rl:60
		atoms[atm][0] *= 10
		atoms[atm][0] += (uint64)((frame.Body)[p] - '0')
		// TODO max size for int/float/string

		goto st9
	tr25:
//line ./op-grammar.rl:60
		atoms[atm][0] *= 10
		atoms[atm][0] += (uint64)((frame.Body)[p] - '0')
		// TODO max size for int/float/string

		goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
//line dfa.go:3687
		if (frame.Body)[p] == 46 {
			goto st10
		}
		if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
			goto tr25
		}
		goto st0
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
			goto tr27
		}
		goto st0
	tr27:
//line ./op-grammar.rl:70
		atoms[atm][0] *= 10
		atoms[atm][0] += (uint64)((frame.Body)[p] - '0')
		e_frac++
		// TODO max size for int/float/string

		goto st25
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
//line dfa.go:3718
		switch (frame.Body)[p] {
		case 32:
			goto tr105
		case 33:
			goto tr106
		case 35:
			goto tr107
		case 39:
			goto tr108
		case 42:
			goto tr107
		case 44:
			goto tr106
		case 46:
			goto tr109
		case 58:
			goto tr107
		case 59:
			goto tr106
		case 61:
			goto tr110
		case 62:
			goto tr111
		case 63:
			goto tr106
		case 64:
			goto tr107
		case 69:
			goto st11
		case 94:
			goto tr113
		case 101:
			goto st11
		}
		switch {
		case (frame.Body)[p] > 13:
			if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
				goto tr27
			}
		case (frame.Body)[p] >= 9:
			goto tr105
		}
		goto st0
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		switch (frame.Body)[p] {
		case 43:
			goto tr28
		case 45:
			goto tr28
		}
		if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
			goto tr29
		}
		goto st0
	tr28:
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
//line dfa.go:3790
		if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
			goto tr29
		}
		goto st0
	tr29:
//line ./op-grammar.rl:81
		e_val *= 10
		e_val += int((frame.Body)[p] - '0')
		// TODO max size for int/float/string

		goto st26
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
//line dfa.go:3808
		switch (frame.Body)[p] {
		case 32:
			goto tr105
		case 33:
			goto tr106
		case 35:
			goto tr107
		case 39:
			goto tr108
		case 42:
			goto tr107
		case 44:
			goto tr106
		case 46:
			goto tr109
		case 58:
			goto tr107
		case 59:
			goto tr106
		case 61:
			goto tr110
		case 62:
			goto tr111
		case 63:
			goto tr106
		case 64:
			goto tr107
		case 94:
			goto tr113
		}
		switch {
		case (frame.Body)[p] > 13:
			if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
				goto tr29
			}
		case (frame.Body)[p] >= 9:
			goto tr105
		}
		goto st0
	tr20:
//line ./op-grammar.rl:108
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
	tr121:
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
//line dfa.go:3884
		switch (frame.Body)[p] {
		case 32:
			goto tr114
		case 33:
			goto tr115
		case 35:
			goto tr116
		case 39:
			goto tr118
		case 42:
			goto tr116
		case 44:
			goto tr115
		case 46:
			goto tr120
		case 58:
			goto tr116
		case 59:
			goto tr115
		case 61:
			goto tr122
		case 62:
			goto tr123
		case 63:
			goto tr115
		case 64:
			goto tr116
		case 91:
			goto tr119
		case 93:
			goto tr119
		case 94:
			goto tr124
		case 95:
			goto tr121
		case 123:
			goto tr119
		case 125:
			goto tr119
		case 126:
			goto tr121
		}
		switch {
		case (frame.Body)[p] < 43:
			switch {
			case (frame.Body)[p] < 36:
				if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
					goto tr114
				}
			case (frame.Body)[p] > 37:
				if 40 <= (frame.Body)[p] && (frame.Body)[p] <= 41 {
					goto tr119
				}
			default:
				goto tr117
			}
		case (frame.Body)[p] > 45:
			switch {
			case (frame.Body)[p] < 65:
				if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
					goto tr121
				}
			case (frame.Body)[p] > 90:
				if 97 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr121
				}
			default:
				goto tr121
			}
		default:
			goto tr117
		}
		goto st0
	tr21:
//line ./op-grammar.rl:108
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
	tr130:
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
//line dfa.go:4001
		switch (frame.Body)[p] {
		case 32:
			goto tr125
		case 33:
			goto tr126
		case 35:
			goto tr127
		case 39:
			goto tr128
		case 42:
			goto tr127
		case 44:
			goto tr126
		case 46:
			goto tr129
		case 58:
			goto tr127
		case 59:
			goto tr126
		case 61:
			goto tr131
		case 62:
			goto tr132
		case 63:
			goto tr126
		case 64:
			goto tr127
		case 91:
			goto tr119
		case 93:
			goto tr119
		case 94:
			goto tr133
		case 95:
			goto tr130
		case 123:
			goto tr119
		case 125:
			goto tr119
		case 126:
			goto tr130
		}
		switch {
		case (frame.Body)[p] < 43:
			switch {
			case (frame.Body)[p] < 36:
				if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
					goto tr125
				}
			case (frame.Body)[p] > 37:
				if 40 <= (frame.Body)[p] && (frame.Body)[p] <= 41 {
					goto tr119
				}
			default:
				goto tr117
			}
		case (frame.Body)[p] > 45:
			switch {
			case (frame.Body)[p] < 65:
				if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
					goto tr130
				}
			case (frame.Body)[p] > 90:
				if 97 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr130
				}
			default:
				goto tr130
			}
		default:
			goto tr117
		}
		goto st0
	tr10:
//line ./op-grammar.rl:100
		atoms[atm][0] = ((uint64)(p)) << 32

		goto st13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
//line dfa.go:4086
		switch (frame.Body)[p] {
		case 10:
			goto st0
		case 13:
			goto st0
		}
		goto st3
	tr52:
//line ././uuid-grammar.rl:29
		dgt = 0
		hlf = 1

//line ././uuid-grammar.rl:9
		dgt = int(ABC[(frame.Body)[p]]) + 4
		atoms[atm][hlf] &= INT60_FLAGS | PREFIX_MASKS[dgt]

		goto st29
	tr53:
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
	tr139:
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
	tr148:
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
//line dfa.go:4155
		switch (frame.Body)[p] {
		case 32:
			goto tr134
		case 33:
			goto tr135
		case 35:
			goto tr136
		case 39:
			goto tr137
		case 42:
			goto tr136
		case 44:
			goto tr135
		case 46:
			goto tr138
		case 58:
			goto tr136
		case 59:
			goto tr135
		case 61:
			goto tr140
		case 62:
			goto tr141
		case 63:
			goto tr135
		case 64:
			goto tr136
		case 94:
			goto tr142
		case 95:
			goto tr139
		case 126:
			goto tr139
		}
		switch {
		case (frame.Body)[p] < 48:
			if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
				goto tr134
			}
		case (frame.Body)[p] > 57:
			switch {
			case (frame.Body)[p] > 90:
				if 97 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr139
				}
			case (frame.Body)[p] >= 65:
				goto tr139
			}
		default:
			goto tr139
		}
		goto st0
	tr44:
//line ././uuid-grammar.rl:5

//line ././uuid-grammar.rl:26

//line ././uuid-grammar.rl:9
		dgt = int(ABC[(frame.Body)[p]]) + 4
		atoms[atm][hlf] &= INT60_FLAGS | PREFIX_MASKS[dgt]

		goto st30
	tr150:
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
//line dfa.go:4236
		switch (frame.Body)[p] {
		case 32:
			goto tr143
		case 33:
			goto tr144
		case 35:
			goto tr145
		case 39:
			goto tr147
		case 42:
			goto tr145
		case 44:
			goto tr144
		case 46:
			goto tr149
		case 58:
			goto tr145
		case 59:
			goto tr144
		case 61:
			goto tr151
		case 62:
			goto tr152
		case 63:
			goto tr144
		case 64:
			goto tr145
		case 91:
			goto tr148
		case 93:
			goto tr148
		case 94:
			goto tr153
		case 95:
			goto tr150
		case 123:
			goto tr148
		case 125:
			goto tr148
		case 126:
			goto tr150
		}
		switch {
		case (frame.Body)[p] < 43:
			switch {
			case (frame.Body)[p] < 36:
				if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
					goto tr143
				}
			case (frame.Body)[p] > 37:
				if 40 <= (frame.Body)[p] && (frame.Body)[p] <= 41 {
					goto tr148
				}
			default:
				goto tr146
			}
		case (frame.Body)[p] > 45:
			switch {
			case (frame.Body)[p] < 65:
				if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
					goto tr150
				}
			case (frame.Body)[p] > 90:
				if 97 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr150
				}
			default:
				goto tr150
			}
		default:
			goto tr146
		}
		goto st0
	tr46:
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
	tr159:
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
//line dfa.go:4345
		switch (frame.Body)[p] {
		case 32:
			goto tr154
		case 33:
			goto tr155
		case 35:
			goto tr156
		case 39:
			goto tr157
		case 42:
			goto tr156
		case 44:
			goto tr155
		case 46:
			goto tr158
		case 58:
			goto tr156
		case 59:
			goto tr155
		case 61:
			goto tr160
		case 62:
			goto tr161
		case 63:
			goto tr155
		case 64:
			goto tr156
		case 91:
			goto tr148
		case 93:
			goto tr148
		case 94:
			goto tr162
		case 95:
			goto tr159
		case 123:
			goto tr148
		case 125:
			goto tr148
		case 126:
			goto tr159
		}
		switch {
		case (frame.Body)[p] < 43:
			switch {
			case (frame.Body)[p] < 36:
				if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
					goto tr154
				}
			case (frame.Body)[p] > 37:
				if 40 <= (frame.Body)[p] && (frame.Body)[p] <= 41 {
					goto tr148
				}
			default:
				goto tr146
			}
		case (frame.Body)[p] > 45:
			switch {
			case (frame.Body)[p] < 65:
				if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
					goto tr159
				}
			case (frame.Body)[p] > 90:
				if 97 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr159
				}
			default:
				goto tr159
			}
		default:
			goto tr146
		}
		goto st0
	tr50:
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
//line dfa.go:4432
		switch (frame.Body)[p] {
		case 32:
			goto tr51
		case 33:
			goto tr40
		case 35:
			goto tr41
		case 39:
			goto tr43
		case 42:
			goto tr41
		case 44:
			goto tr40
		case 46:
			goto tr45
		case 58:
			goto tr41
		case 59:
			goto tr40
		case 61:
			goto tr47
		case 62:
			goto tr48
		case 63:
			goto tr40
		case 64:
			goto tr41
		case 91:
			goto tr44
		case 93:
			goto tr44
		case 94:
			goto tr49
		case 95:
			goto tr46
		case 123:
			goto tr44
		case 125:
			goto tr44
		case 126:
			goto tr46
		}
		switch {
		case (frame.Body)[p] < 43:
			switch {
			case (frame.Body)[p] < 36:
				if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
					goto tr51
				}
			case (frame.Body)[p] > 37:
				if 40 <= (frame.Body)[p] && (frame.Body)[p] <= 41 {
					goto tr44
				}
			default:
				goto tr42
			}
		case (frame.Body)[p] > 45:
			switch {
			case (frame.Body)[p] < 65:
				if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
					goto tr46
				}
			case (frame.Body)[p] > 90:
				if 97 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr46
				}
			default:
				goto tr46
			}
		default:
			goto tr42
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
	_test_eof2:
		(ps.state) = 2
		goto _test_eof
	_test_eof3:
		(ps.state) = 3
		goto _test_eof
	_test_eof19:
		(ps.state) = 19
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
			case 15:
//line ./op-grammar.rl:148
				frame.Position++

			case 20:
//line ./op-grammar.rl:125
//line ./op-grammar.rl:148
				frame.Position++

			case 18:
//line ./op-grammar.rl:152

//line ./op-grammar.rl:148
				frame.Position++

			case 16, 17, 32:
//line ./op-grammar.rl:27
				atm++

//line ./op-grammar.rl:152

//line ./op-grammar.rl:148
				frame.Position++

			case 19:
//line ./op-grammar.rl:35
				atm++

//line ./op-grammar.rl:125

//line ./op-grammar.rl:148
				frame.Position++

			case 30:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:27
				atm++

//line ./op-grammar.rl:152

//line ./op-grammar.rl:148
				frame.Position++

			case 29:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:27
				atm++

//line ./op-grammar.rl:152

//line ./op-grammar.rl:148
				frame.Position++

			case 22:
//line ./op-grammar.rl:51
				atoms[atm][1] |= ((uint64)(ATOM_INT)) << 62

//line ./op-grammar.rl:35
				atm++

//line ./op-grammar.rl:125

//line ./op-grammar.rl:148
				frame.Position++

			case 25, 26:
//line ./op-grammar.rl:86
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

//line ./op-grammar.rl:35
				atm++

//line ./op-grammar.rl:125

//line ./op-grammar.rl:148
				frame.Position++

			case 23:
//line ./op-grammar.rl:115
				atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:35
				atm++

//line ./op-grammar.rl:125

//line ./op-grammar.rl:148
				frame.Position++

			case 31:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
				atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:27
				atm++

//line ./op-grammar.rl:152

//line ./op-grammar.rl:148
				frame.Position++

			case 27:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:115
				atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:35
				atm++

//line ./op-grammar.rl:125

//line ./op-grammar.rl:148
				frame.Position++

			case 24:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:115
				atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:35
				atm++

//line ./op-grammar.rl:125

//line ./op-grammar.rl:148
				frame.Position++

			case 28:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
				atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:115
				atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:35
				atm++

//line ./op-grammar.rl:125

//line ./op-grammar.rl:148
				frame.Position++

//line dfa.go:4756
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
//line dfa.go:4801
	const UUID_start int = 1
	const UUID_first_final int = 2
	const UUID_error int = 0

	const UUID_en_main int = 1

//line dfa.rl:108
	cs, p, pe, eof := 0, 0, len(data), len(data)
	_ = eof

	atm, hlf, dgt := 0, 0, 0

	atoms := [1]Atom{Atom(ctx_uuid)}

//line dfa.go:4819
	{
		cs = UUID_start
	}

//line dfa.go:4824
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
//line dfa.go:4917
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
//line dfa.go:5011
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
//line dfa.go:5059
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
//line dfa.go:5136
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

//line dfa.go:5203
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
