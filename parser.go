//line dfa.rl:1
package ron

import "fmt"
import "errors"

//line dfa.rl:7

//line parser.go:12
const RON_start int = 15
const RON_first_final int = 15
const RON_error int = 0

const RON_en_main int = 15

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
		if ps.pos != 0 {
			return
		}

//line parser.go:46
		{
			(ps.state) = RON_start
		}

//line dfa.rl:30
		frame.position = -1
		frame.atoms = frame._atoms[:4]

	case RON_FULL_STOP:
		ps.state = RON_error
		return

	case RON_start:
		ps.off = ps.pos
		frame.atoms = frame._atoms[:4]
		ps.atm, ps.hlf, ps.dgt = 0, VALUE, 0
	}

	if ps.pos >= len(frame.Body) {
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

	atm, hlf, dgt, p, atoms := ps.atm, ps.hlf, ps.dgt, ps.pos, frame.atoms

//line parser.go:84
	{
		if p == pe {
			goto _test_eof
		}
		goto _resume

	_again:
		switch ps.state {
		case 15:
			goto st15
		case 0:
			goto st0
		case 1:
			goto st1
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
		case 22:
			goto st22
		case 4:
			goto st4
		case 5:
			goto st5
		case 23:
			goto st23
		case 6:
			goto st6
		case 24:
			goto st24
		case 25:
			goto st25
		case 7:
			goto st7
		case 8:
			goto st8
		case 9:
			goto st9
		case 10:
			goto st10
		case 26:
			goto st26
		case 11:
			goto st11
		case 12:
			goto st12
		case 27:
			goto st27
		case 13:
			goto st13
		case 28:
			goto st28
		case 29:
			goto st29
		case 30:
			goto st30
		case 31:
			goto st31
		case 14:
			goto st14
		case 32:
			goto st32
		case 33:
			goto st33
		case 34:
			goto st34
		case 35:
			goto st35
		case 36:
			goto st36
		case 37:
			goto st37
		}

		if p++; p == pe {
			goto _test_eof
		}
	_resume:
		switch ps.state {
		case 15:
			goto st_case_15
		case 0:
			goto st_case_0
		case 1:
			goto st_case_1
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
		case 22:
			goto st_case_22
		case 4:
			goto st_case_4
		case 5:
			goto st_case_5
		case 23:
			goto st_case_23
		case 6:
			goto st_case_6
		case 24:
			goto st_case_24
		case 25:
			goto st_case_25
		case 7:
			goto st_case_7
		case 8:
			goto st_case_8
		case 9:
			goto st_case_9
		case 10:
			goto st_case_10
		case 26:
			goto st_case_26
		case 11:
			goto st_case_11
		case 12:
			goto st_case_12
		case 27:
			goto st_case_27
		case 13:
			goto st_case_13
		case 28:
			goto st_case_28
		case 29:
			goto st_case_29
		case 30:
			goto st_case_30
		case 31:
			goto st_case_31
		case 14:
			goto st_case_14
		case 32:
			goto st_case_32
		case 33:
			goto st_case_33
		case 34:
			goto st_case_34
		case 35:
			goto st_case_35
		case 36:
			goto st_case_36
		case 37:
			goto st_case_37
		}
		goto st_out
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
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
			goto tr31
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
	tr160:
//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:128

//line ./op-grammar.rl:124

		frame.position++

		goto st1
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
//line parser.go:315
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
		(ps.state) = 16
//line ./op-grammar.rl:108

		hlf = VALUE
		if p > frame.Parser.off && frame.position != -1 {
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

//line ./op-grammar.rl:104

		frame.term = int(ABC[(frame.Body)[p]])

		goto _again
	tr33:
		(ps.state) = 16
//line ./op-grammar.rl:124

		frame.position++

//line ./op-grammar.rl:108

		hlf = VALUE
		if p > frame.Parser.off && frame.position != -1 {
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

//line ./op-grammar.rl:104

		frame.term = int(ABC[(frame.Body)[p]])

		goto _again
	tr41:
//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:128

//line ./op-grammar.rl:104

		frame.term = int(ABC[(frame.Body)[p]])

		goto st16
	tr56:
//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:101

//line ./op-grammar.rl:104

		frame.term = int(ABC[(frame.Body)[p]])

		goto st16
	tr64:
//line ./op-grammar.rl:101

//line ./op-grammar.rl:104

		frame.term = int(ABC[(frame.Body)[p]])

		goto st16
	tr72:
//line ./op-grammar.rl:60

		atoms[atm][1] |= ((uint64)(ATOM_INT)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:101

//line ./op-grammar.rl:104

		frame.term = int(ABC[(frame.Body)[p]])

		goto st16
	tr80:
//line ./op-grammar.rl:91

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:101

//line ./op-grammar.rl:104

		frame.term = int(ABC[(frame.Body)[p]])

		goto st16
	tr90:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:91

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:101

//line ./op-grammar.rl:104

		frame.term = int(ABC[(frame.Body)[p]])

		goto st16
	tr99:
//line ./op-grammar.rl:69

		atoms[atm].setTill(p)
		atoms[atm].parseValue(frame.Body)

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:101

//line ./op-grammar.rl:104

		frame.term = int(ABC[(frame.Body)[p]])

		goto st16
	tr109:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:91

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:101

//line ./op-grammar.rl:104

		frame.term = int(ABC[(frame.Body)[p]])

		goto st16
	tr120:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:51

		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:91

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:101

//line ./op-grammar.rl:104

		frame.term = int(ABC[(frame.Body)[p]])

		goto st16
	tr130:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:128

//line ./op-grammar.rl:104

		frame.term = int(ABC[(frame.Body)[p]])

		goto st16
	tr140:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:128

//line ./op-grammar.rl:104

		frame.term = int(ABC[(frame.Body)[p]])

		goto st16
	tr151:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:51

		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:128

//line ./op-grammar.rl:104

		frame.term = int(ABC[(frame.Body)[p]])

		goto st16
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
//line parser.go:606
		switch (frame.Body)[p] {
		case 32:
			goto st16
		case 33:
			goto tr33
		case 35:
			goto tr34
		case 39:
			goto tr35
		case 42:
			goto tr34
		case 44:
			goto tr33
		case 46:
			goto tr36
		case 58:
			goto tr34
		case 59:
			goto tr33
		case 61:
			goto tr37
		case 62:
			goto tr38
		case 63:
			goto tr33
		case 64:
			goto tr34
		case 94:
			goto tr39
		}
		if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
			goto st16
		}
		goto st0
	tr3:
		(ps.state) = 17
//line ./op-grammar.rl:108

		hlf = VALUE
		if p > frame.Parser.off && frame.position != -1 {
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

//line ./op-grammar.rl:6

		ps.omitted = 15

//line ./op-grammar.rl:16

		n = (int)(ABC[(frame.Body)[p]])
		hlf, dgt = VALUE, 0
		if n < atm {
			// parse #op1#op2#op3 without Ragel state explosion
			(ps.state) = (RON_start)
			frame.position++
			p--
			{
				p++
				goto _out
			}
		} else {
			// next UUID
			atm = n
			ps.omitted -= 1 << uint(n)
		}

		goto _again
	tr34:
		(ps.state) = 17
//line ./op-grammar.rl:124

		frame.position++

//line ./op-grammar.rl:108

		hlf = VALUE
		if p > frame.Parser.off && frame.position != -1 {
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

//line ./op-grammar.rl:6

		ps.omitted = 15

//line ./op-grammar.rl:16

		n = (int)(ABC[(frame.Body)[p]])
		hlf, dgt = VALUE, 0
		if n < atm {
			// parse #op1#op2#op3 without Ragel state explosion
			(ps.state) = (RON_start)
			frame.position++
			p--
			{
				p++
				goto _out
			}
		} else {
			// next UUID
			atm = n
			ps.omitted -= 1 << uint(n)
		}

		goto _again
	tr42:
		(ps.state) = 17
//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:16

		n = (int)(ABC[(frame.Body)[p]])
		hlf, dgt = VALUE, 0
		if n < atm {
			// parse #op1#op2#op3 without Ragel state explosion
			(ps.state) = (RON_start)
			frame.position++
			p--
			{
				p++
				goto _out
			}
		} else {
			// next UUID
			atm = n
			ps.omitted -= 1 << uint(n)
		}

		goto _again
	tr57:
		(ps.state) = 17
//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:101

//line ./op-grammar.rl:124

		frame.position++

//line ./op-grammar.rl:108

		hlf = VALUE
		if p > frame.Parser.off && frame.position != -1 {
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

//line ./op-grammar.rl:6

		ps.omitted = 15

//line ./op-grammar.rl:16

		n = (int)(ABC[(frame.Body)[p]])
		hlf, dgt = VALUE, 0
		if n < atm {
			// parse #op1#op2#op3 without Ragel state explosion
			(ps.state) = (RON_start)
			frame.position++
			p--
			{
				p++
				goto _out
			}
		} else {
			// next UUID
			atm = n
			ps.omitted -= 1 << uint(n)
		}

		goto _again
	tr65:
		(ps.state) = 17
//line ./op-grammar.rl:101

//line ./op-grammar.rl:124

		frame.position++

//line ./op-grammar.rl:108

		hlf = VALUE
		if p > frame.Parser.off && frame.position != -1 {
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

//line ./op-grammar.rl:6

		ps.omitted = 15

//line ./op-grammar.rl:16

		n = (int)(ABC[(frame.Body)[p]])
		hlf, dgt = VALUE, 0
		if n < atm {
			// parse #op1#op2#op3 without Ragel state explosion
			(ps.state) = (RON_start)
			frame.position++
			p--
			{
				p++
				goto _out
			}
		} else {
			// next UUID
			atm = n
			ps.omitted -= 1 << uint(n)
		}

		goto _again
	tr73:
		(ps.state) = 17
//line ./op-grammar.rl:60

		atoms[atm][1] |= ((uint64)(ATOM_INT)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:101

//line ./op-grammar.rl:124

		frame.position++

//line ./op-grammar.rl:108

		hlf = VALUE
		if p > frame.Parser.off && frame.position != -1 {
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

//line ./op-grammar.rl:6

		ps.omitted = 15

//line ./op-grammar.rl:16

		n = (int)(ABC[(frame.Body)[p]])
		hlf, dgt = VALUE, 0
		if n < atm {
			// parse #op1#op2#op3 without Ragel state explosion
			(ps.state) = (RON_start)
			frame.position++
			p--
			{
				p++
				goto _out
			}
		} else {
			// next UUID
			atm = n
			ps.omitted -= 1 << uint(n)
		}

		goto _again
	tr81:
		(ps.state) = 17
//line ./op-grammar.rl:91

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:101

//line ./op-grammar.rl:124

		frame.position++

//line ./op-grammar.rl:108

		hlf = VALUE
		if p > frame.Parser.off && frame.position != -1 {
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

//line ./op-grammar.rl:6

		ps.omitted = 15

//line ./op-grammar.rl:16

		n = (int)(ABC[(frame.Body)[p]])
		hlf, dgt = VALUE, 0
		if n < atm {
			// parse #op1#op2#op3 without Ragel state explosion
			(ps.state) = (RON_start)
			frame.position++
			p--
			{
				p++
				goto _out
			}
		} else {
			// next UUID
			atm = n
			ps.omitted -= 1 << uint(n)
		}

		goto _again
	tr91:
		(ps.state) = 17
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:91

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:101

//line ./op-grammar.rl:124

		frame.position++

//line ./op-grammar.rl:108

		hlf = VALUE
		if p > frame.Parser.off && frame.position != -1 {
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

//line ./op-grammar.rl:6

		ps.omitted = 15

//line ./op-grammar.rl:16

		n = (int)(ABC[(frame.Body)[p]])
		hlf, dgt = VALUE, 0
		if n < atm {
			// parse #op1#op2#op3 without Ragel state explosion
			(ps.state) = (RON_start)
			frame.position++
			p--
			{
				p++
				goto _out
			}
		} else {
			// next UUID
			atm = n
			ps.omitted -= 1 << uint(n)
		}

		goto _again
	tr100:
		(ps.state) = 17
//line ./op-grammar.rl:69

		atoms[atm].setTill(p)
		atoms[atm].parseValue(frame.Body)

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:101

//line ./op-grammar.rl:124

		frame.position++

//line ./op-grammar.rl:108

		hlf = VALUE
		if p > frame.Parser.off && frame.position != -1 {
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

//line ./op-grammar.rl:6

		ps.omitted = 15

//line ./op-grammar.rl:16

		n = (int)(ABC[(frame.Body)[p]])
		hlf, dgt = VALUE, 0
		if n < atm {
			// parse #op1#op2#op3 without Ragel state explosion
			(ps.state) = (RON_start)
			frame.position++
			p--
			{
				p++
				goto _out
			}
		} else {
			// next UUID
			atm = n
			ps.omitted -= 1 << uint(n)
		}

		goto _again
	tr110:
		(ps.state) = 17
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:91

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:101

//line ./op-grammar.rl:124

		frame.position++

//line ./op-grammar.rl:108

		hlf = VALUE
		if p > frame.Parser.off && frame.position != -1 {
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

//line ./op-grammar.rl:6

		ps.omitted = 15

//line ./op-grammar.rl:16

		n = (int)(ABC[(frame.Body)[p]])
		hlf, dgt = VALUE, 0
		if n < atm {
			// parse #op1#op2#op3 without Ragel state explosion
			(ps.state) = (RON_start)
			frame.position++
			p--
			{
				p++
				goto _out
			}
		} else {
			// next UUID
			atm = n
			ps.omitted -= 1 << uint(n)
		}

		goto _again
	tr121:
		(ps.state) = 17
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:51

		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:91

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:101

//line ./op-grammar.rl:124

		frame.position++

//line ./op-grammar.rl:108

		hlf = VALUE
		if p > frame.Parser.off && frame.position != -1 {
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

//line ./op-grammar.rl:6

		ps.omitted = 15

//line ./op-grammar.rl:16

		n = (int)(ABC[(frame.Body)[p]])
		hlf, dgt = VALUE, 0
		if n < atm {
			// parse #op1#op2#op3 without Ragel state explosion
			(ps.state) = (RON_start)
			frame.position++
			p--
			{
				p++
				goto _out
			}
		} else {
			// next UUID
			atm = n
			ps.omitted -= 1 << uint(n)
		}

		goto _again
	tr131:
		(ps.state) = 17
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:16

		n = (int)(ABC[(frame.Body)[p]])
		hlf, dgt = VALUE, 0
		if n < atm {
			// parse #op1#op2#op3 without Ragel state explosion
			(ps.state) = (RON_start)
			frame.position++
			p--
			{
				p++
				goto _out
			}
		} else {
			// next UUID
			atm = n
			ps.omitted -= 1 << uint(n)
		}

		goto _again
	tr141:
		(ps.state) = 17
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:16

		n = (int)(ABC[(frame.Body)[p]])
		hlf, dgt = VALUE, 0
		if n < atm {
			// parse #op1#op2#op3 without Ragel state explosion
			(ps.state) = (RON_start)
			frame.position++
			p--
			{
				p++
				goto _out
			}
		} else {
			// next UUID
			atm = n
			ps.omitted -= 1 << uint(n)
		}

		goto _again
	tr152:
		(ps.state) = 17
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:51

		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:16

		n = (int)(ABC[(frame.Body)[p]])
		hlf, dgt = VALUE, 0
		if n < atm {
			// parse #op1#op2#op3 without Ragel state explosion
			(ps.state) = (RON_start)
			frame.position++
			p--
			{
				p++
				goto _out
			}
		} else {
			// next UUID
			atm = n
			ps.omitted -= 1 << uint(n)
		}

		goto _again
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
//line parser.go:1267
		switch (frame.Body)[p] {
		case 32:
			goto st17
		case 33:
			goto tr41
		case 35:
			goto tr42
		case 39:
			goto tr44
		case 42:
			goto tr42
		case 44:
			goto tr41
		case 46:
			goto tr46
		case 58:
			goto tr42
		case 59:
			goto tr41
		case 61:
			goto tr48
		case 62:
			goto tr49
		case 63:
			goto tr41
		case 64:
			goto tr42
		case 91:
			goto tr45
		case 93:
			goto tr45
		case 94:
			goto tr50
		case 96:
			goto tr51
		case 123:
			goto tr45
		case 125:
			goto tr45
		case 126:
			goto tr47
		}
		switch {
		case (frame.Body)[p] < 43:
			switch {
			case (frame.Body)[p] < 36:
				if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
					goto st17
				}
			case (frame.Body)[p] > 37:
				if 40 <= (frame.Body)[p] && (frame.Body)[p] <= 41 {
					goto tr45
				}
			default:
				goto tr43
			}
		case (frame.Body)[p] > 45:
			switch {
			case (frame.Body)[p] < 65:
				if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
					goto tr47
				}
			case (frame.Body)[p] > 90:
				if 95 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr47
				}
			default:
				goto tr47
			}
		default:
			goto tr43
		}
		goto st0
	tr43:
//line ././uuid-grammar.rl:5

//line ././uuid-grammar.rl:45

		hlf = 1
		atoms[atm][1] &= INT60_FULL
		atoms[atm][1] |= ((uint64)(ABC[(frame.Body)[p]])) << 60

		goto st18
	tr142:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:45

		hlf = 1
		atoms[atm][1] &= INT60_FULL
		atoms[atm][1] |= ((uint64)(ABC[(frame.Body)[p]])) << 60

		goto st18
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
//line parser.go:1368
		switch (frame.Body)[p] {
		case 32:
			goto st19
		case 33:
			goto tr41
		case 35:
			goto tr42
		case 39:
			goto tr44
		case 42:
			goto tr42
		case 44:
			goto tr41
		case 46:
			goto tr46
		case 58:
			goto tr42
		case 59:
			goto tr41
		case 61:
			goto tr48
		case 62:
			goto tr49
		case 63:
			goto tr41
		case 64:
			goto tr42
		case 91:
			goto tr53
		case 93:
			goto tr53
		case 94:
			goto tr50
		case 95:
			goto tr54
		case 123:
			goto tr53
		case 125:
			goto tr53
		case 126:
			goto tr54
		}
		switch {
		case (frame.Body)[p] < 48:
			switch {
			case (frame.Body)[p] > 13:
				if 40 <= (frame.Body)[p] && (frame.Body)[p] <= 41 {
					goto tr53
				}
			case (frame.Body)[p] >= 9:
				goto st19
			}
		case (frame.Body)[p] > 57:
			switch {
			case (frame.Body)[p] > 90:
				if 97 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr54
				}
			case (frame.Body)[p] >= 65:
				goto tr54
			}
		default:
			goto tr54
		}
		goto st0
	tr129:
//line ././uuid-grammar.rl:37

		goto st19
	tr139:
//line ././uuid-grammar.rl:34

		goto st19
	tr150:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:51

		atoms[atm][1] = UUID_NAME_FLAG

		goto st19
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
//line parser.go:1458
		switch (frame.Body)[p] {
		case 32:
			goto st19
		case 33:
			goto tr41
		case 35:
			goto tr42
		case 39:
			goto tr44
		case 42:
			goto tr42
		case 44:
			goto tr41
		case 46:
			goto tr46
		case 58:
			goto tr42
		case 59:
			goto tr41
		case 61:
			goto tr48
		case 62:
			goto tr49
		case 63:
			goto tr41
		case 64:
			goto tr42
		case 94:
			goto tr50
		}
		if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
			goto st19
		}
		goto st0
	tr4:
		(ps.state) = 2
//line ./op-grammar.rl:108

		hlf = VALUE
		if p > frame.Parser.off && frame.position != -1 {
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

//line ./op-grammar.rl:95

		atm = 4
		hlf = VALUE
		dgt = 0

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto _again
	tr35:
		(ps.state) = 2
//line ./op-grammar.rl:124

		frame.position++

//line ./op-grammar.rl:108

		hlf = VALUE
		if p > frame.Parser.off && frame.position != -1 {
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

//line ./op-grammar.rl:95

		atm = 4
		hlf = VALUE
		dgt = 0

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto _again
	tr44:
//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:128

//line ./op-grammar.rl:95

		atm = 4
		hlf = VALUE
		dgt = 0

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st2
	tr58:
//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st2
	tr66:
//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st2
	tr74:
//line ./op-grammar.rl:60

		atoms[atm][1] |= ((uint64)(ATOM_INT)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st2
	tr82:
//line ./op-grammar.rl:91

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st2
	tr92:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:91

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st2
	tr101:
//line ./op-grammar.rl:69

		atoms[atm].setTill(p)
		atoms[atm].parseValue(frame.Body)

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st2
	tr112:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:91

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st2
	tr122:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:51

		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:91

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st2
	tr132:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:128

//line ./op-grammar.rl:95

		atm = 4
		hlf = VALUE
		dgt = 0

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st2
	tr143:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:128

//line ./op-grammar.rl:95

		atm = 4
		hlf = VALUE
		dgt = 0

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st2
	tr153:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:51

		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:128

//line ./op-grammar.rl:95

		atm = 4
		hlf = VALUE
		dgt = 0

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
//line parser.go:1777
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
//line ./op-grammar.rl:74

		atoms[atm][0] = ((uint64)(p)) << 32

		goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
//line parser.go:1802
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
			goto st14
		}
		goto st3
	tr9:
//line ./op-grammar.rl:74

		atoms[atm][0] = ((uint64)(p)) << 32

//line ./op-grammar.rl:78

		atoms[atm][0] |= uint64(p)
		atoms[atm][1] = ((uint64)(ATOM_STRING)) << 62

		goto st20
	tr12:
//line ./op-grammar.rl:78

		atoms[atm][0] |= uint64(p)
		atoms[atm][1] = ((uint64)(ATOM_STRING)) << 62

		goto st20
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
//line parser.go:1839
		switch (frame.Body)[p] {
		case 32:
			goto tr55
		case 33:
			goto tr56
		case 35:
			goto tr57
		case 39:
			goto tr58
		case 42:
			goto tr57
		case 44:
			goto tr56
		case 46:
			goto tr59
		case 58:
			goto tr57
		case 59:
			goto tr56
		case 61:
			goto tr60
		case 62:
			goto tr61
		case 63:
			goto tr56
		case 64:
			goto tr57
		case 94:
			goto tr62
		}
		if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
			goto tr55
		}
		goto st0
	tr55:
//line ./op-grammar.rl:41

		atm++

		goto st21
	tr71:
//line ./op-grammar.rl:60

		atoms[atm][1] |= ((uint64)(ATOM_INT)) << 62

//line ./op-grammar.rl:41

		atm++

		goto st21
	tr79:
//line ./op-grammar.rl:91

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

		goto st21
	tr89:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:91

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

		goto st21
	tr98:
//line ./op-grammar.rl:69

		atoms[atm].setTill(p)
		atoms[atm].parseValue(frame.Body)

//line ./op-grammar.rl:41

		atm++

		goto st21
	tr108:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:91

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

		goto st21
	tr119:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:51

		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:91

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

		goto st21
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
//line parser.go:1959
		switch (frame.Body)[p] {
		case 32:
			goto st21
		case 33:
			goto tr64
		case 35:
			goto tr65
		case 39:
			goto tr66
		case 42:
			goto tr65
		case 44:
			goto tr64
		case 46:
			goto tr67
		case 58:
			goto tr65
		case 59:
			goto tr64
		case 61:
			goto tr68
		case 62:
			goto tr69
		case 63:
			goto tr64
		case 64:
			goto tr65
		case 94:
			goto tr70
		}
		if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
			goto st21
		}
		goto st0
	tr31:
		(ps.state) = 22
//line ./op-grammar.rl:131

		(ps.state) = (RON_FULL_STOP)
		{
			p++
			goto _out
		}

		goto _again
	tr36:
		(ps.state) = 22
//line ./op-grammar.rl:124

		frame.position++

//line ./op-grammar.rl:131

		(ps.state) = (RON_FULL_STOP)
		{
			p++
			goto _out
		}

		goto _again
	tr46:
		(ps.state) = 22
//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:128

//line ./op-grammar.rl:124

		frame.position++

//line ./op-grammar.rl:131

		(ps.state) = (RON_FULL_STOP)
		{
			p++
			goto _out
		}

		goto _again
	tr59:
		(ps.state) = 22
//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:101

//line ./op-grammar.rl:124

		frame.position++

//line ./op-grammar.rl:131

		(ps.state) = (RON_FULL_STOP)
		{
			p++
			goto _out
		}

		goto _again
	tr67:
		(ps.state) = 22
//line ./op-grammar.rl:101

//line ./op-grammar.rl:124

		frame.position++

//line ./op-grammar.rl:131

		(ps.state) = (RON_FULL_STOP)
		{
			p++
			goto _out
		}

		goto _again
	tr75:
		(ps.state) = 22
//line ./op-grammar.rl:60

		atoms[atm][1] |= ((uint64)(ATOM_INT)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:101

//line ./op-grammar.rl:124

		frame.position++

//line ./op-grammar.rl:131

		(ps.state) = (RON_FULL_STOP)
		{
			p++
			goto _out
		}

		goto _again
	tr84:
		(ps.state) = 22
//line ./op-grammar.rl:91

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:101

//line ./op-grammar.rl:124

		frame.position++

//line ./op-grammar.rl:131

		(ps.state) = (RON_FULL_STOP)
		{
			p++
			goto _out
		}

		goto _again
	tr93:
		(ps.state) = 22
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:91

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:101

//line ./op-grammar.rl:124

		frame.position++

//line ./op-grammar.rl:131

		(ps.state) = (RON_FULL_STOP)
		{
			p++
			goto _out
		}

		goto _again
	tr102:
		(ps.state) = 22
//line ./op-grammar.rl:69

		atoms[atm].setTill(p)
		atoms[atm].parseValue(frame.Body)

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:101

//line ./op-grammar.rl:124

		frame.position++

//line ./op-grammar.rl:131

		(ps.state) = (RON_FULL_STOP)
		{
			p++
			goto _out
		}

		goto _again
	tr114:
		(ps.state) = 22
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:91

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:101

//line ./op-grammar.rl:124

		frame.position++

//line ./op-grammar.rl:131

		(ps.state) = (RON_FULL_STOP)
		{
			p++
			goto _out
		}

		goto _again
	tr123:
		(ps.state) = 22
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:51

		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:91

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:101

//line ./op-grammar.rl:124

		frame.position++

//line ./op-grammar.rl:131

		(ps.state) = (RON_FULL_STOP)
		{
			p++
			goto _out
		}

		goto _again
	tr133:
		(ps.state) = 22
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:128

//line ./op-grammar.rl:124

		frame.position++

//line ./op-grammar.rl:131

		(ps.state) = (RON_FULL_STOP)
		{
			p++
			goto _out
		}

		goto _again
	tr145:
		(ps.state) = 22
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:128

//line ./op-grammar.rl:124

		frame.position++

//line ./op-grammar.rl:131

		(ps.state) = (RON_FULL_STOP)
		{
			p++
			goto _out
		}

		goto _again
	tr154:
		(ps.state) = 22
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:51

		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:128

//line ./op-grammar.rl:124

		frame.position++

//line ./op-grammar.rl:131

		(ps.state) = (RON_FULL_STOP)
		{
			p++
			goto _out
		}

		goto _again
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
//line parser.go:2294
		goto st0
	tr5:
		(ps.state) = 4
//line ./op-grammar.rl:108

		hlf = VALUE
		if p > frame.Parser.off && frame.position != -1 {
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

//line ./op-grammar.rl:95

		atm = 4
		hlf = VALUE
		dgt = 0

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto _again
	tr37:
		(ps.state) = 4
//line ./op-grammar.rl:124

		frame.position++

//line ./op-grammar.rl:108

		hlf = VALUE
		if p > frame.Parser.off && frame.position != -1 {
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

//line ./op-grammar.rl:95

		atm = 4
		hlf = VALUE
		dgt = 0

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto _again
	tr48:
//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:128

//line ./op-grammar.rl:95

		atm = 4
		hlf = VALUE
		dgt = 0

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st4
	tr60:
//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st4
	tr68:
//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st4
	tr76:
//line ./op-grammar.rl:60

		atoms[atm][1] |= ((uint64)(ATOM_INT)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st4
	tr86:
//line ./op-grammar.rl:91

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st4
	tr95:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:91

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st4
	tr103:
//line ./op-grammar.rl:69

		atoms[atm].setTill(p)
		atoms[atm].parseValue(frame.Body)

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st4
	tr116:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:91

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st4
	tr126:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:51

		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:91

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st4
	tr135:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:128

//line ./op-grammar.rl:95

		atm = 4
		hlf = VALUE
		dgt = 0

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st4
	tr147:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:128

//line ./op-grammar.rl:95

		atm = 4
		hlf = VALUE
		dgt = 0

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st4
	tr157:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:51

		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:128

//line ./op-grammar.rl:95

		atm = 4
		hlf = VALUE
		dgt = 0

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
//line parser.go:2580
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
//line ./op-grammar.rl:45

//line ./op-grammar.rl:48

		if (frame.Body)[p] == '-' {
			atoms[atm][1] |= 1
		}

		goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
//line parser.go:2614
		if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
			goto tr17
		}
		goto st0
	tr16:
//line ./op-grammar.rl:45

//line ./op-grammar.rl:54

		atoms[atm][0] *= 10
		atoms[atm][0] += (uint64)((frame.Body)[p] - '0')
		// TODO max size for int/float/string

		goto st23
	tr17:
//line ./op-grammar.rl:54

		atoms[atm][0] *= 10
		atoms[atm][0] += (uint64)((frame.Body)[p] - '0')
		// TODO max size for int/float/string

		goto st23
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
//line parser.go:2643
		switch (frame.Body)[p] {
		case 32:
			goto tr71
		case 33:
			goto tr72
		case 35:
			goto tr73
		case 39:
			goto tr74
		case 42:
			goto tr73
		case 44:
			goto tr72
		case 46:
			goto tr75
		case 58:
			goto tr73
		case 59:
			goto tr72
		case 61:
			goto tr76
		case 62:
			goto tr77
		case 63:
			goto tr72
		case 64:
			goto tr73
		case 94:
			goto tr78
		}
		switch {
		case (frame.Body)[p] > 13:
			if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
				goto tr17
			}
		case (frame.Body)[p] >= 9:
			goto tr71
		}
		goto st0
	tr6:
		(ps.state) = 6
//line ./op-grammar.rl:108

		hlf = VALUE
		if p > frame.Parser.off && frame.position != -1 {
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

//line ./op-grammar.rl:95

		atm = 4
		hlf = VALUE
		dgt = 0

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto _again
	tr38:
		(ps.state) = 6
//line ./op-grammar.rl:124

		frame.position++

//line ./op-grammar.rl:108

		hlf = VALUE
		if p > frame.Parser.off && frame.position != -1 {
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

//line ./op-grammar.rl:95

		atm = 4
		hlf = VALUE
		dgt = 0

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto _again
	tr49:
//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:128

//line ./op-grammar.rl:95

		atm = 4
		hlf = VALUE
		dgt = 0

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st6
	tr61:
//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st6
	tr69:
//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st6
	tr77:
//line ./op-grammar.rl:60

		atoms[atm][1] |= ((uint64)(ATOM_INT)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st6
	tr87:
//line ./op-grammar.rl:91

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st6
	tr96:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:91

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st6
	tr104:
//line ./op-grammar.rl:69

		atoms[atm].setTill(p)
		atoms[atm].parseValue(frame.Body)

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st6
	tr117:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:91

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st6
	tr127:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:51

		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:91

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st6
	tr136:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:128

//line ./op-grammar.rl:95

		atm = 4
		hlf = VALUE
		dgt = 0

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st6
	tr148:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:128

//line ./op-grammar.rl:95

		atm = 4
		hlf = VALUE
		dgt = 0

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st6
	tr158:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:51

		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:128

//line ./op-grammar.rl:95

		atm = 4
		hlf = VALUE
		dgt = 0

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
//line parser.go:2967
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
//line ./op-grammar.rl:83

		if atm == 4 {
			atoms[atm] = atoms[SPEC_OBJECT]
		} else if atoms[atm-1].Type() == ATOM_UUID {
			atoms[atm] = atoms[atm-1]
		}

//line ././uuid-grammar.rl:5

//line ././uuid-grammar.rl:45

		hlf = 1
		atoms[atm][1] &= INT60_FULL
		atoms[atm][1] |= ((uint64)(ABC[(frame.Body)[p]])) << 60

		goto st24
	tr111:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:45

		hlf = 1
		atoms[atm][1] &= INT60_FULL
		atoms[atm][1] |= ((uint64)(ABC[(frame.Body)[p]])) << 60

		goto st24
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
//line parser.go:3050
		switch (frame.Body)[p] {
		case 32:
			goto tr79
		case 33:
			goto tr80
		case 35:
			goto tr81
		case 39:
			goto tr82
		case 42:
			goto tr81
		case 44:
			goto tr80
		case 46:
			goto tr84
		case 58:
			goto tr81
		case 59:
			goto tr80
		case 61:
			goto tr86
		case 62:
			goto tr87
		case 63:
			goto tr80
		case 64:
			goto tr81
		case 91:
			goto tr83
		case 93:
			goto tr83
		case 94:
			goto tr88
		case 95:
			goto tr85
		case 123:
			goto tr83
		case 125:
			goto tr83
		case 126:
			goto tr85
		}
		switch {
		case (frame.Body)[p] < 48:
			switch {
			case (frame.Body)[p] > 13:
				if 40 <= (frame.Body)[p] && (frame.Body)[p] <= 41 {
					goto tr83
				}
			case (frame.Body)[p] >= 9:
				goto tr79
			}
		case (frame.Body)[p] > 57:
			switch {
			case (frame.Body)[p] > 90:
				if 97 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr85
				}
			case (frame.Body)[p] >= 65:
				goto tr85
			}
		default:
			goto tr85
		}
		goto st0
	tr83:
//line ././uuid-grammar.rl:29

		dgt = 0
		hlf = 1

//line ././uuid-grammar.rl:9

		dgt = int(ABC[(frame.Body)[p]]) + 4
		atoms[atm][hlf] &= INT60_FLAGS | PREFIX_MASKS[dgt]

		goto st25
	tr94:
//line ././uuid-grammar.rl:14

		atoms[atm][hlf] |= ((uint64)(ABC[(frame.Body)[p]])) << DIGIT_OFFSETS[dgt]
		dgt++
		if dgt > 10 {
			{
				p++
				(ps.state) = 25
				goto _out
			}
		}

		goto st25
	tr107:
//line ././uuid-grammar.rl:40

		atoms[atm][hlf] <<= 6
		dgt--

		goto st25
	tr113:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:29

		dgt = 0
		hlf = 1

//line ././uuid-grammar.rl:9

		dgt = int(ABC[(frame.Body)[p]]) + 4
		atoms[atm][hlf] &= INT60_FLAGS | PREFIX_MASKS[dgt]

		goto st25
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
//line parser.go:3165
		switch (frame.Body)[p] {
		case 32:
			goto tr89
		case 33:
			goto tr90
		case 35:
			goto tr91
		case 39:
			goto tr92
		case 42:
			goto tr91
		case 44:
			goto tr90
		case 46:
			goto tr93
		case 58:
			goto tr91
		case 59:
			goto tr90
		case 61:
			goto tr95
		case 62:
			goto tr96
		case 63:
			goto tr90
		case 64:
			goto tr91
		case 94:
			goto tr97
		case 95:
			goto tr94
		case 126:
			goto tr94
		}
		switch {
		case (frame.Body)[p] < 48:
			if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
				goto tr89
			}
		case (frame.Body)[p] > 57:
			switch {
			case (frame.Body)[p] > 90:
				if 97 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr94
				}
			case (frame.Body)[p] >= 65:
				goto tr94
			}
		default:
			goto tr94
		}
		goto st0
	tr7:
		(ps.state) = 7
//line ./op-grammar.rl:108

		hlf = VALUE
		if p > frame.Parser.off && frame.position != -1 {
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

//line ./op-grammar.rl:95

		atm = 4
		hlf = VALUE
		dgt = 0

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto _again
	tr39:
		(ps.state) = 7
//line ./op-grammar.rl:124

		frame.position++

//line ./op-grammar.rl:108

		hlf = VALUE
		if p > frame.Parser.off && frame.position != -1 {
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

//line ./op-grammar.rl:95

		atm = 4
		hlf = VALUE
		dgt = 0

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto _again
	tr50:
//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:128

//line ./op-grammar.rl:95

		atm = 4
		hlf = VALUE
		dgt = 0

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st7
	tr62:
//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st7
	tr70:
//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st7
	tr78:
//line ./op-grammar.rl:60

		atoms[atm][1] |= ((uint64)(ATOM_INT)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st7
	tr88:
//line ./op-grammar.rl:91

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st7
	tr97:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:91

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st7
	tr106:
//line ./op-grammar.rl:69

		atoms[atm].setTill(p)
		atoms[atm].parseValue(frame.Body)

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st7
	tr118:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:91

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st7
	tr128:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:51

		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:91

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st7
	tr137:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:128

//line ./op-grammar.rl:95

		atm = 4
		hlf = VALUE
		dgt = 0

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st7
	tr149:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:128

//line ./op-grammar.rl:95

		atm = 4
		hlf = VALUE
		dgt = 0

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st7
	tr159:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:51

		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:128

//line ./op-grammar.rl:95

		atm = 4
		hlf = VALUE
		dgt = 0

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
//line parser.go:3502
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
//line ./op-grammar.rl:64

		atoms[atm].setFloatType()
		atoms[atm].setFrom(p)

		goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
//line parser.go:3532
		if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
			goto st9
		}
		goto st0
	tr24:
//line ./op-grammar.rl:64

		atoms[atm].setFloatType()
		atoms[atm].setFrom(p)

		goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
//line parser.go:3549
		switch (frame.Body)[p] {
		case 46:
			goto st10
		case 69:
			goto st13
		case 101:
			goto st13
		}
		if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
			goto st9
		}
		goto st0
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
			goto st26
		}
		goto st0
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
		switch (frame.Body)[p] {
		case 32:
			goto tr98
		case 33:
			goto tr99
		case 35:
			goto tr100
		case 39:
			goto tr101
		case 42:
			goto tr100
		case 44:
			goto tr99
		case 46:
			goto tr102
		case 58:
			goto tr100
		case 59:
			goto tr99
		case 61:
			goto tr103
		case 62:
			goto tr104
		case 63:
			goto tr99
		case 64:
			goto tr100
		case 69:
			goto st11
		case 94:
			goto tr106
		case 101:
			goto st11
		}
		switch {
		case (frame.Body)[p] > 13:
			if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
				goto st26
			}
		case (frame.Body)[p] >= 9:
			goto tr98
		}
		goto st0
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		switch (frame.Body)[p] {
		case 43:
			goto st12
		case 45:
			goto st12
		}
		if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
			goto st27
		}
		goto st0
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
			goto st27
		}
		goto st0
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
		switch (frame.Body)[p] {
		case 32:
			goto tr98
		case 33:
			goto tr99
		case 35:
			goto tr100
		case 39:
			goto tr101
		case 42:
			goto tr100
		case 44:
			goto tr99
		case 46:
			goto tr102
		case 58:
			goto tr100
		case 59:
			goto tr99
		case 61:
			goto tr103
		case 62:
			goto tr104
		case 63:
			goto tr99
		case 64:
			goto tr100
		case 94:
			goto tr106
		}
		switch {
		case (frame.Body)[p] > 13:
			if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
				goto st27
			}
		case (frame.Body)[p] >= 9:
			goto tr98
		}
		goto st0
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		switch (frame.Body)[p] {
		case 43:
			goto st10
		case 45:
			goto st10
		}
		if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
			goto st26
		}
		goto st0
	tr85:
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
//line parser.go:3726
		switch (frame.Body)[p] {
		case 32:
			goto tr89
		case 33:
			goto tr90
		case 35:
			goto tr91
		case 39:
			goto tr92
		case 42:
			goto tr91
		case 44:
			goto tr90
		case 46:
			goto tr93
		case 47:
			goto tr107
		case 58:
			goto tr91
		case 59:
			goto tr90
		case 61:
			goto tr95
		case 62:
			goto tr96
		case 63:
			goto tr90
		case 64:
			goto tr91
		case 94:
			goto tr97
		case 95:
			goto tr94
		case 126:
			goto tr94
		}
		switch {
		case (frame.Body)[p] < 48:
			if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
				goto tr89
			}
		case (frame.Body)[p] > 57:
			switch {
			case (frame.Body)[p] > 90:
				if 97 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr94
				}
			case (frame.Body)[p] >= 65:
				goto tr94
			}
		default:
			goto tr94
		}
		goto st0
	tr20:
//line ./op-grammar.rl:83

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

		goto st29
	tr115:
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
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
//line parser.go:3817
		switch (frame.Body)[p] {
		case 32:
			goto tr108
		case 33:
			goto tr109
		case 35:
			goto tr110
		case 39:
			goto tr112
		case 42:
			goto tr110
		case 44:
			goto tr109
		case 46:
			goto tr114
		case 58:
			goto tr110
		case 59:
			goto tr109
		case 61:
			goto tr116
		case 62:
			goto tr117
		case 63:
			goto tr109
		case 64:
			goto tr110
		case 91:
			goto tr113
		case 93:
			goto tr113
		case 94:
			goto tr118
		case 95:
			goto tr115
		case 123:
			goto tr113
		case 125:
			goto tr113
		case 126:
			goto tr115
		}
		switch {
		case (frame.Body)[p] < 43:
			switch {
			case (frame.Body)[p] < 36:
				if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
					goto tr108
				}
			case (frame.Body)[p] > 37:
				if 40 <= (frame.Body)[p] && (frame.Body)[p] <= 41 {
					goto tr113
				}
			default:
				goto tr111
			}
		case (frame.Body)[p] > 45:
			switch {
			case (frame.Body)[p] < 65:
				if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
					goto tr115
				}
			case (frame.Body)[p] > 90:
				if 97 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr115
				}
			default:
				goto tr115
			}
		default:
			goto tr111
		}
		goto st0
	tr21:
//line ./op-grammar.rl:83

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
//line parser.go:3924
		switch (frame.Body)[p] {
		case 32:
			goto tr119
		case 33:
			goto tr120
		case 35:
			goto tr121
		case 39:
			goto tr122
		case 42:
			goto tr121
		case 44:
			goto tr120
		case 46:
			goto tr123
		case 47:
			goto tr124
		case 58:
			goto tr121
		case 59:
			goto tr120
		case 61:
			goto tr126
		case 62:
			goto tr127
		case 63:
			goto tr120
		case 64:
			goto tr121
		case 91:
			goto tr113
		case 93:
			goto tr113
		case 94:
			goto tr128
		case 95:
			goto tr125
		case 123:
			goto tr113
		case 125:
			goto tr113
		case 126:
			goto tr125
		}
		switch {
		case (frame.Body)[p] < 43:
			switch {
			case (frame.Body)[p] < 36:
				if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
					goto tr119
				}
			case (frame.Body)[p] > 37:
				if 40 <= (frame.Body)[p] && (frame.Body)[p] <= 41 {
					goto tr113
				}
			default:
				goto tr111
			}
		case (frame.Body)[p] > 45:
			switch {
			case (frame.Body)[p] < 65:
				if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
					goto tr125
				}
			case (frame.Body)[p] > 90:
				if 97 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr125
				}
			default:
				goto tr125
			}
		default:
			goto tr111
		}
		goto st0
	tr125:
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
	tr124:
//line ././uuid-grammar.rl:40

		atoms[atm][hlf] <<= 6
		dgt--

		goto st31
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
//line parser.go:4022
		switch (frame.Body)[p] {
		case 32:
			goto tr119
		case 33:
			goto tr120
		case 35:
			goto tr121
		case 39:
			goto tr122
		case 42:
			goto tr121
		case 44:
			goto tr120
		case 46:
			goto tr123
		case 58:
			goto tr121
		case 59:
			goto tr120
		case 61:
			goto tr126
		case 62:
			goto tr127
		case 63:
			goto tr120
		case 64:
			goto tr121
		case 91:
			goto tr113
		case 93:
			goto tr113
		case 94:
			goto tr128
		case 95:
			goto tr125
		case 123:
			goto tr113
		case 125:
			goto tr113
		case 126:
			goto tr125
		}
		switch {
		case (frame.Body)[p] < 43:
			switch {
			case (frame.Body)[p] < 36:
				if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
					goto tr119
				}
			case (frame.Body)[p] > 37:
				if 40 <= (frame.Body)[p] && (frame.Body)[p] <= 41 {
					goto tr113
				}
			default:
				goto tr111
			}
		case (frame.Body)[p] > 45:
			switch {
			case (frame.Body)[p] < 65:
				if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
					goto tr125
				}
			case (frame.Body)[p] > 90:
				if 97 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr125
				}
			default:
				goto tr125
			}
		default:
			goto tr111
		}
		goto st0
	tr10:
//line ./op-grammar.rl:74

		atoms[atm][0] = ((uint64)(p)) << 32

		goto st14
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
//line parser.go:4107
		switch (frame.Body)[p] {
		case 10:
			goto st0
		case 13:
			goto st0
		}
		goto st3
	tr53:
//line ././uuid-grammar.rl:29

		dgt = 0
		hlf = 1

//line ././uuid-grammar.rl:9

		dgt = int(ABC[(frame.Body)[p]]) + 4
		atoms[atm][hlf] &= INT60_FLAGS | PREFIX_MASKS[dgt]

		goto st32
	tr134:
//line ././uuid-grammar.rl:14

		atoms[atm][hlf] |= ((uint64)(ABC[(frame.Body)[p]])) << DIGIT_OFFSETS[dgt]
		dgt++
		if dgt > 10 {
			{
				p++
				(ps.state) = 32
				goto _out
			}
		}

		goto st32
	tr138:
//line ././uuid-grammar.rl:40

		atoms[atm][hlf] <<= 6
		dgt--

		goto st32
	tr144:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:29

		dgt = 0
		hlf = 1

//line ././uuid-grammar.rl:9

		dgt = int(ABC[(frame.Body)[p]]) + 4
		atoms[atm][hlf] &= INT60_FLAGS | PREFIX_MASKS[dgt]

		goto st32
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
//line parser.go:4164
		switch (frame.Body)[p] {
		case 32:
			goto tr129
		case 33:
			goto tr130
		case 35:
			goto tr131
		case 39:
			goto tr132
		case 42:
			goto tr131
		case 44:
			goto tr130
		case 46:
			goto tr133
		case 58:
			goto tr131
		case 59:
			goto tr130
		case 61:
			goto tr135
		case 62:
			goto tr136
		case 63:
			goto tr130
		case 64:
			goto tr131
		case 94:
			goto tr137
		case 95:
			goto tr134
		case 126:
			goto tr134
		}
		switch {
		case (frame.Body)[p] < 48:
			if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
				goto tr129
			}
		case (frame.Body)[p] > 57:
			switch {
			case (frame.Body)[p] > 90:
				if 97 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr134
				}
			case (frame.Body)[p] >= 65:
				goto tr134
			}
		default:
			goto tr134
		}
		goto st0
	tr54:
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
				(ps.state) = 33
				goto _out
			}
		}

		goto st33
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
//line parser.go:4241
		switch (frame.Body)[p] {
		case 32:
			goto tr129
		case 33:
			goto tr130
		case 35:
			goto tr131
		case 39:
			goto tr132
		case 42:
			goto tr131
		case 44:
			goto tr130
		case 46:
			goto tr133
		case 47:
			goto tr138
		case 58:
			goto tr131
		case 59:
			goto tr130
		case 61:
			goto tr135
		case 62:
			goto tr136
		case 63:
			goto tr130
		case 64:
			goto tr131
		case 94:
			goto tr137
		case 95:
			goto tr134
		case 126:
			goto tr134
		}
		switch {
		case (frame.Body)[p] < 48:
			if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
				goto tr129
			}
		case (frame.Body)[p] > 57:
			switch {
			case (frame.Body)[p] > 90:
				if 97 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr134
				}
			case (frame.Body)[p] >= 65:
				goto tr134
			}
		default:
			goto tr134
		}
		goto st0
	tr45:
//line ././uuid-grammar.rl:5

//line ././uuid-grammar.rl:26

//line ././uuid-grammar.rl:9

		dgt = int(ABC[(frame.Body)[p]]) + 4
		atoms[atm][hlf] &= INT60_FLAGS | PREFIX_MASKS[dgt]

		goto st34
	tr146:
//line ././uuid-grammar.rl:14

		atoms[atm][hlf] |= ((uint64)(ABC[(frame.Body)[p]])) << DIGIT_OFFSETS[dgt]
		dgt++
		if dgt > 10 {
			{
				p++
				(ps.state) = 34
				goto _out
			}
		}

		goto st34
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
//line parser.go:4324
		switch (frame.Body)[p] {
		case 32:
			goto tr139
		case 33:
			goto tr140
		case 35:
			goto tr141
		case 39:
			goto tr143
		case 42:
			goto tr141
		case 44:
			goto tr140
		case 46:
			goto tr145
		case 58:
			goto tr141
		case 59:
			goto tr140
		case 61:
			goto tr147
		case 62:
			goto tr148
		case 63:
			goto tr140
		case 64:
			goto tr141
		case 91:
			goto tr144
		case 93:
			goto tr144
		case 94:
			goto tr149
		case 95:
			goto tr146
		case 123:
			goto tr144
		case 125:
			goto tr144
		case 126:
			goto tr146
		}
		switch {
		case (frame.Body)[p] < 43:
			switch {
			case (frame.Body)[p] < 36:
				if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
					goto tr139
				}
			case (frame.Body)[p] > 37:
				if 40 <= (frame.Body)[p] && (frame.Body)[p] <= 41 {
					goto tr144
				}
			default:
				goto tr142
			}
		case (frame.Body)[p] > 45:
			switch {
			case (frame.Body)[p] < 65:
				if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
					goto tr146
				}
			case (frame.Body)[p] > 90:
				if 97 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr146
				}
			default:
				goto tr146
			}
		default:
			goto tr142
		}
		goto st0
	tr47:
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
				(ps.state) = 35
				goto _out
			}
		}

		goto st35
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
//line parser.go:4423
		switch (frame.Body)[p] {
		case 32:
			goto tr150
		case 33:
			goto tr151
		case 35:
			goto tr152
		case 39:
			goto tr153
		case 42:
			goto tr152
		case 44:
			goto tr151
		case 46:
			goto tr154
		case 47:
			goto tr155
		case 58:
			goto tr152
		case 59:
			goto tr151
		case 61:
			goto tr157
		case 62:
			goto tr158
		case 63:
			goto tr151
		case 64:
			goto tr152
		case 91:
			goto tr144
		case 93:
			goto tr144
		case 94:
			goto tr159
		case 95:
			goto tr156
		case 123:
			goto tr144
		case 125:
			goto tr144
		case 126:
			goto tr156
		}
		switch {
		case (frame.Body)[p] < 43:
			switch {
			case (frame.Body)[p] < 36:
				if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
					goto tr150
				}
			case (frame.Body)[p] > 37:
				if 40 <= (frame.Body)[p] && (frame.Body)[p] <= 41 {
					goto tr144
				}
			default:
				goto tr142
			}
		case (frame.Body)[p] > 45:
			switch {
			case (frame.Body)[p] < 65:
				if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
					goto tr156
				}
			case (frame.Body)[p] > 90:
				if 97 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr156
				}
			default:
				goto tr156
			}
		default:
			goto tr142
		}
		goto st0
	tr156:
//line ././uuid-grammar.rl:14

		atoms[atm][hlf] |= ((uint64)(ABC[(frame.Body)[p]])) << DIGIT_OFFSETS[dgt]
		dgt++
		if dgt > 10 {
			{
				p++
				(ps.state) = 36
				goto _out
			}
		}

		goto st36
	tr155:
//line ././uuid-grammar.rl:40

		atoms[atm][hlf] <<= 6
		dgt--

		goto st36
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
//line parser.go:4521
		switch (frame.Body)[p] {
		case 32:
			goto tr150
		case 33:
			goto tr151
		case 35:
			goto tr152
		case 39:
			goto tr153
		case 42:
			goto tr152
		case 44:
			goto tr151
		case 46:
			goto tr154
		case 58:
			goto tr152
		case 59:
			goto tr151
		case 61:
			goto tr157
		case 62:
			goto tr158
		case 63:
			goto tr151
		case 64:
			goto tr152
		case 91:
			goto tr144
		case 93:
			goto tr144
		case 94:
			goto tr159
		case 95:
			goto tr156
		case 123:
			goto tr144
		case 125:
			goto tr144
		case 126:
			goto tr156
		}
		switch {
		case (frame.Body)[p] < 43:
			switch {
			case (frame.Body)[p] < 36:
				if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
					goto tr150
				}
			case (frame.Body)[p] > 37:
				if 40 <= (frame.Body)[p] && (frame.Body)[p] <= 41 {
					goto tr144
				}
			default:
				goto tr142
			}
		case (frame.Body)[p] > 45:
			switch {
			case (frame.Body)[p] < 65:
				if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
					goto tr156
				}
			case (frame.Body)[p] > 90:
				if 97 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr156
				}
			default:
				goto tr156
			}
		default:
			goto tr142
		}
		goto st0
	tr51:
//line ./op-grammar.rl:10

		if atm > 0 {
			atoms[atm] = atoms[atm-1]
		}

		goto st37
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
//line parser.go:4608
		switch (frame.Body)[p] {
		case 32:
			goto tr160
		case 33:
			goto tr41
		case 35:
			goto tr42
		case 39:
			goto tr44
		case 42:
			goto tr42
		case 44:
			goto tr41
		case 46:
			goto tr46
		case 58:
			goto tr42
		case 59:
			goto tr41
		case 61:
			goto tr48
		case 62:
			goto tr49
		case 63:
			goto tr41
		case 64:
			goto tr42
		case 91:
			goto tr45
		case 93:
			goto tr45
		case 94:
			goto tr50
		case 95:
			goto tr47
		case 123:
			goto tr45
		case 125:
			goto tr45
		case 126:
			goto tr47
		}
		switch {
		case (frame.Body)[p] < 43:
			switch {
			case (frame.Body)[p] < 36:
				if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
					goto tr160
				}
			case (frame.Body)[p] > 37:
				if 40 <= (frame.Body)[p] && (frame.Body)[p] <= 41 {
					goto tr45
				}
			default:
				goto tr43
			}
		case (frame.Body)[p] > 45:
			switch {
			case (frame.Body)[p] < 65:
				if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
					goto tr47
				}
			case (frame.Body)[p] > 90:
				if 97 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr47
				}
			default:
				goto tr47
			}
		default:
			goto tr43
		}
		goto st0
	st_out:
	_test_eof15:
		(ps.state) = 15
		goto _test_eof
	_test_eof1:
		(ps.state) = 1
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
	_test_eof22:
		(ps.state) = 22
		goto _test_eof
	_test_eof4:
		(ps.state) = 4
		goto _test_eof
	_test_eof5:
		(ps.state) = 5
		goto _test_eof
	_test_eof23:
		(ps.state) = 23
		goto _test_eof
	_test_eof6:
		(ps.state) = 6
		goto _test_eof
	_test_eof24:
		(ps.state) = 24
		goto _test_eof
	_test_eof25:
		(ps.state) = 25
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
	_test_eof26:
		(ps.state) = 26
		goto _test_eof
	_test_eof11:
		(ps.state) = 11
		goto _test_eof
	_test_eof12:
		(ps.state) = 12
		goto _test_eof
	_test_eof27:
		(ps.state) = 27
		goto _test_eof
	_test_eof13:
		(ps.state) = 13
		goto _test_eof
	_test_eof28:
		(ps.state) = 28
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
	_test_eof14:
		(ps.state) = 14
		goto _test_eof
	_test_eof32:
		(ps.state) = 32
		goto _test_eof
	_test_eof33:
		(ps.state) = 33
		goto _test_eof
	_test_eof34:
		(ps.state) = 34
		goto _test_eof
	_test_eof35:
		(ps.state) = 35
		goto _test_eof
	_test_eof36:
		(ps.state) = 36
		goto _test_eof
	_test_eof37:
		(ps.state) = 37
		goto _test_eof

	_test_eof:
		{
		}
		if p == eof {
			switch ps.state {
			case 16:
//line ./op-grammar.rl:124

				frame.position++

			case 21:
//line ./op-grammar.rl:101

//line ./op-grammar.rl:124

				frame.position++

			case 17, 18, 19, 37:
//line ./op-grammar.rl:32

				atm++

//line ./op-grammar.rl:128

//line ./op-grammar.rl:124

				frame.position++

			case 20:
//line ./op-grammar.rl:41

				atm++

//line ./op-grammar.rl:101

//line ./op-grammar.rl:124

				frame.position++

			case 34:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:32

				atm++

//line ./op-grammar.rl:128

//line ./op-grammar.rl:124

				frame.position++

			case 32, 33:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:32

				atm++

//line ./op-grammar.rl:128

//line ./op-grammar.rl:124

				frame.position++

			case 23:
//line ./op-grammar.rl:60

				atoms[atm][1] |= ((uint64)(ATOM_INT)) << 62

//line ./op-grammar.rl:41

				atm++

//line ./op-grammar.rl:101

//line ./op-grammar.rl:124

				frame.position++

			case 26, 27:
//line ./op-grammar.rl:69

				atoms[atm].setTill(p)
				atoms[atm].parseValue(frame.Body)

//line ./op-grammar.rl:41

				atm++

//line ./op-grammar.rl:101

//line ./op-grammar.rl:124

				frame.position++

			case 24:
//line ./op-grammar.rl:91

				atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

				atm++

//line ./op-grammar.rl:101

//line ./op-grammar.rl:124

				frame.position++

			case 35, 36:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:51

				atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:32

				atm++

//line ./op-grammar.rl:128

//line ./op-grammar.rl:124

				frame.position++

			case 29:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:91

				atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

				atm++

//line ./op-grammar.rl:101

//line ./op-grammar.rl:124

				frame.position++

			case 25, 28:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:91

				atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

				atm++

//line ./op-grammar.rl:101

//line ./op-grammar.rl:124

				frame.position++

			case 30, 31:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:51

				atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:91

				atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

				atm++

//line ./op-grammar.rl:101

//line ./op-grammar.rl:124

				frame.position++

//line parser.go:4920
			}
		}

	_out:
		{
		}
	}

//line dfa.rl:67

	ps.atm, ps.hlf, ps.dgt, ps.pos, frame.atoms = atm, hlf, dgt, p, atoms

	switch {
	case ps.state == RON_error:
		frame.position = -1
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
			frame.position = -1
		}
	}

	//fmt.Println("omits", frame.IsComplete(), frame.term!=TERM_REDUCED,  ps.omitted, frame.Parser.state, ps.pos)
	if frame.IsComplete() && frame.term != TERM_REDUCED && ps.omitted != 0 {
		for u := uint(0); u < 4; u++ {
			if ps.omitted&(1<<u) != 0 {
				frame.atoms[u] = Atom(ZERO_UUID)
			}
		}
	}
}

func (ctx_uuid UUID) Parse(data []byte) (UUID, error) {

//line dfa.rl:102

//line parser.go:4965
	const UUID_start int = 1
	const UUID_first_final int = 2
	const UUID_error int = 0

	const UUID_en_main int = 1

//line dfa.rl:103

	cs, p, pe, eof := 0, 0, len(data), len(data)
	_ = eof

	atm, hlf, dgt := 0, 0, 0

	atoms := [1]Atom{Atom(ctx_uuid)}

//line parser.go:4983
	{
		cs = UUID_start
	}

//line parser.go:4988
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
		case 6:
			goto st_case_6
		case 7:
			goto st_case_7
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

//line ./uuid-grammar.rl:45

		hlf = 1
		atoms[atm][1] &= INT60_FULL
		atoms[atm][1] |= ((uint64)(ABC[data[p]])) << 60

		goto st2
	tr8:
//line ./uuid-grammar.rl:34

//line ./uuid-grammar.rl:45

		hlf = 1
		atoms[atm][1] &= INT60_FULL
		atoms[atm][1] |= ((uint64)(ABC[data[p]])) << 60

		goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
//line parser.go:5085
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
	tr7:
//line ./uuid-grammar.rl:40

		atoms[atm][hlf] <<= 6
		dgt--

		goto st3
	tr9:
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
//line parser.go:5167
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
//line parser.go:5211
		switch data[p] {
		case 47:
			goto tr7
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
//line parser.go:5261
		switch data[p] {
		case 43:
			goto tr8
		case 45:
			goto tr8
		case 91:
			goto tr9
		case 93:
			goto tr9
		case 95:
			goto tr10
		case 123:
			goto tr9
		case 125:
			goto tr9
		case 126:
			goto tr10
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 37:
				if 40 <= data[p] && data[p] <= 41 {
					goto tr9
				}
			case data[p] >= 36:
				goto tr8
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
				cs = 6
				goto _out
			}
		}

		goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
//line parser.go:5328
		switch data[p] {
		case 43:
			goto tr8
		case 45:
			goto tr8
		case 47:
			goto tr11
		case 91:
			goto tr9
		case 93:
			goto tr9
		case 95:
			goto tr12
		case 123:
			goto tr9
		case 125:
			goto tr9
		case 126:
			goto tr12
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 37:
				if 40 <= data[p] && data[p] <= 41 {
					goto tr9
				}
			case data[p] >= 36:
				goto tr8
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr12
				}
			case data[p] >= 65:
				goto tr12
			}
		default:
			goto tr12
		}
		goto st0
	tr12:
//line ./uuid-grammar.rl:14

		atoms[atm][hlf] |= ((uint64)(ABC[data[p]])) << DIGIT_OFFSETS[dgt]
		dgt++
		if dgt > 10 {
			{
				p++
				cs = 7
				goto _out
			}
		}

		goto st7
	tr11:
//line ./uuid-grammar.rl:40

		atoms[atm][hlf] <<= 6
		dgt--

		goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
//line parser.go:5394
		switch data[p] {
		case 43:
			goto tr8
		case 45:
			goto tr8
		case 91:
			goto tr9
		case 93:
			goto tr9
		case 95:
			goto tr12
		case 123:
			goto tr9
		case 125:
			goto tr9
		case 126:
			goto tr12
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 37:
				if 40 <= data[p] && data[p] <= 41 {
					goto tr9
				}
			case data[p] >= 36:
				goto tr8
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr12
				}
			case data[p] >= 65:
				goto tr12
			}
		default:
			goto tr12
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
	_test_eof6:
		cs = 6
		goto _test_eof
	_test_eof7:
		cs = 7
		goto _test_eof

	_test_eof:
		{
		}
		if p == eof {
			switch cs {
			case 5:
//line ./uuid-grammar.rl:34

			case 3, 4:
//line ./uuid-grammar.rl:37

			case 6, 7:
//line ./uuid-grammar.rl:34

//line ./uuid-grammar.rl:51

				atoms[atm][1] = UUID_NAME_FLAG

//line parser.go:5463
			}
		}

	_out:
		{
		}
	}

//line dfa.rl:118

	if cs < UUID_first_final || dgt > 10 {
		return ERROR_UUID, errors.New(fmt.Sprintf("parse error at pos %d", p))
	} else {
		return UUID(atoms[0]), nil
	}

}
