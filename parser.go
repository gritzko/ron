//line dfa.rl:1
package ron

import "fmt"
import "errors"

//line dfa.rl:7

//line parser.go:12
const RON_start int = 14
const RON_first_final int = 14
const RON_error int = 0

const RON_en_main int = 14

//line dfa.rl:8

//line dfa.rl:9

//line dfa.rl:10

//line dfa.rl:11

const (
	// The parser reached end-of-input (in block mode) or
	// the closing dot (in streaming mode) successfully.
	// The rest of the input is frame.Rest()
	RON_FULL_STOP        = -1
	RON_OPEN      uint64 = 1919905842 // https://play.golang.org/p/vo74Pf-DKh2
	RON_CLOSED    uint64 = 1919905330
	opFlag        uint64 = 3 << 62
)

// Parse consumes one op from data[], unless the buffer ends earlier.
// Fills atoms[]
func (frame *Frame) Parse() {

	ps := &frame.Parser

	switch ps.state {
	case RON_error:
		if ps.pos != 0 {
			return
		}

//line parser.go:51
		{
			(ps.state) = RON_start
		}

//line dfa.rl:35
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

	frame.descriptor[VALUE] = RON_CLOSED // ?
	frame.descriptor[ORIGIN] = 0

	frame.descriptor.setType(opFlag)
	frame.descriptor.setFrom(p)

//line parser.go:95
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
		case 29:
			goto st29
		case 13:
			goto st13
		case 30:
			goto st30
		case 31:
			goto st31
		case 32:
			goto st32
		case 33:
			goto st33
		case 34:
			goto st34
		case 35:
			goto st35
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
		case 29:
			goto st_case_29
		case 13:
			goto st_case_13
		case 30:
			goto st_case_30
		case 31:
			goto st_case_31
		case 32:
			goto st_case_32
		case 33:
			goto st_case_33
		case 34:
			goto st_case_34
		case 35:
			goto st_case_35
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
	tr151:
//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:110

//line ./op-grammar.rl:106

		frame.position++

		goto st1
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
//line parser.go:318
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
//line ./op-grammar.rl:90

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

//line ./op-grammar.rl:86

		frame.term = int(ABC[(frame.Body)[p]])

		goto _again
	tr32:
		(ps.state) = 15
//line ./op-grammar.rl:106

		frame.position++

//line ./op-grammar.rl:90

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

//line ./op-grammar.rl:86

		frame.term = int(ABC[(frame.Body)[p]])

		goto _again
	tr40:
//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:110

//line ./op-grammar.rl:86

		frame.term = int(ABC[(frame.Body)[p]])

		goto st15
	tr55:
//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:83

//line ./op-grammar.rl:86

		frame.term = int(ABC[(frame.Body)[p]])

		goto st15
	tr63:
//line ./op-grammar.rl:83

//line ./op-grammar.rl:86

		frame.term = int(ABC[(frame.Body)[p]])

		goto st15
	tr71:
//line ./op-grammar.rl:60

		atoms[atm].setTill(p)
		atoms[atm].parseValue(frame.Body)

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:83

//line ./op-grammar.rl:86

		frame.term = int(ABC[(frame.Body)[p]])

		goto st15
	tr79:
//line ./op-grammar.rl:73

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:83

//line ./op-grammar.rl:86

		frame.term = int(ABC[(frame.Body)[p]])

		goto st15
	tr89:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:73

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:83

//line ./op-grammar.rl:86

		frame.term = int(ABC[(frame.Body)[p]])

		goto st15
	tr100:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:73

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:83

//line ./op-grammar.rl:86

		frame.term = int(ABC[(frame.Body)[p]])

		goto st15
	tr111:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:51

		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:73

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:83

//line ./op-grammar.rl:86

		frame.term = int(ABC[(frame.Body)[p]])

		goto st15
	tr121:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:110

//line ./op-grammar.rl:86

		frame.term = int(ABC[(frame.Body)[p]])

		goto st15
	tr131:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:110

//line ./op-grammar.rl:86

		frame.term = int(ABC[(frame.Body)[p]])

		goto st15
	tr142:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:51

		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:110

//line ./op-grammar.rl:86

		frame.term = int(ABC[(frame.Body)[p]])

		goto st15
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
//line parser.go:592
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
//line ./op-grammar.rl:90

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
	tr33:
		(ps.state) = 16
//line ./op-grammar.rl:106

		frame.position++

//line ./op-grammar.rl:90

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
	tr41:
		(ps.state) = 16
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
	tr56:
		(ps.state) = 16
//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:83

//line ./op-grammar.rl:106

		frame.position++

//line ./op-grammar.rl:90

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
	tr64:
		(ps.state) = 16
//line ./op-grammar.rl:83

//line ./op-grammar.rl:106

		frame.position++

//line ./op-grammar.rl:90

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
	tr72:
		(ps.state) = 16
//line ./op-grammar.rl:60

		atoms[atm].setTill(p)
		atoms[atm].parseValue(frame.Body)

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:83

//line ./op-grammar.rl:106

		frame.position++

//line ./op-grammar.rl:90

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
	tr80:
		(ps.state) = 16
//line ./op-grammar.rl:73

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:83

//line ./op-grammar.rl:106

		frame.position++

//line ./op-grammar.rl:90

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
	tr90:
		(ps.state) = 16
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:73

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:83

//line ./op-grammar.rl:106

		frame.position++

//line ./op-grammar.rl:90

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
	tr101:
		(ps.state) = 16
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:73

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:83

//line ./op-grammar.rl:106

		frame.position++

//line ./op-grammar.rl:90

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
	tr112:
		(ps.state) = 16
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:51

		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:73

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:83

//line ./op-grammar.rl:106

		frame.position++

//line ./op-grammar.rl:90

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
	tr122:
		(ps.state) = 16
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
	tr132:
		(ps.state) = 16
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
	tr143:
		(ps.state) = 16
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
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
//line parser.go:1199
		switch (frame.Body)[p] {
		case 32:
			goto st16
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
					goto st16
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

//line ././uuid-grammar.rl:45

		hlf = 1
		atoms[atm][1] &= INT60_FULL
		atoms[atm][1] |= ((uint64)(ABC[(frame.Body)[p]])) << 60

		goto st17
	tr133:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:45

		hlf = 1
		atoms[atm][1] &= INT60_FULL
		atoms[atm][1] |= ((uint64)(ABC[(frame.Body)[p]])) << 60

		goto st17
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
//line parser.go:1300
		switch (frame.Body)[p] {
		case 32:
			goto st18
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
				goto st18
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
	tr120:
//line ././uuid-grammar.rl:37

		goto st18
	tr130:
//line ././uuid-grammar.rl:34

		goto st18
	tr141:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:51

		atoms[atm][1] = UUID_NAME_FLAG

		goto st18
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
//line parser.go:1390
		switch (frame.Body)[p] {
		case 32:
			goto st18
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
		case 94:
			goto tr49
		}
		if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
			goto st18
		}
		goto st0
	tr4:
		(ps.state) = 2
//line ./op-grammar.rl:90

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

//line ./op-grammar.rl:77

		atm = 4
		hlf = VALUE
		dgt = 0

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto _again
	tr34:
		(ps.state) = 2
//line ./op-grammar.rl:106

		frame.position++

//line ./op-grammar.rl:90

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

//line ./op-grammar.rl:77

		atm = 4
		hlf = VALUE
		dgt = 0

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto _again
	tr43:
//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:110

//line ./op-grammar.rl:77

		atm = 4
		hlf = VALUE
		dgt = 0

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st2
	tr57:
//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st2
	tr65:
//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st2
	tr73:
//line ./op-grammar.rl:60

		atoms[atm].setTill(p)
		atoms[atm].parseValue(frame.Body)

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st2
	tr81:
//line ./op-grammar.rl:73

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st2
	tr91:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:73

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st2
	tr103:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:73

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st2
	tr113:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:51

		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:73

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st2
	tr123:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:110

//line ./op-grammar.rl:77

		atm = 4
		hlf = VALUE
		dgt = 0

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st2
	tr134:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:110

//line ./op-grammar.rl:77

		atm = 4
		hlf = VALUE
		dgt = 0

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st2
	tr144:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:51

		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:110

//line ./op-grammar.rl:77

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
//line parser.go:1694
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
//line ./op-grammar.rl:55

		atoms[atm].setStringType()
		atoms[atm].setFrom(p)

		goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
//line parser.go:1720
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
//line ./op-grammar.rl:55

		atoms[atm].setStringType()
		atoms[atm].setFrom(p)

//line ./op-grammar.rl:60

		atoms[atm].setTill(p)
		atoms[atm].parseValue(frame.Body)

		goto st19
	tr12:
//line ./op-grammar.rl:60

		atoms[atm].setTill(p)
		atoms[atm].parseValue(frame.Body)

		goto st19
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
//line parser.go:1758
		switch (frame.Body)[p] {
		case 32:
			goto tr54
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
			goto tr54
		}
		goto st0
	tr54:
//line ./op-grammar.rl:41

		atm++

		goto st20
	tr70:
//line ./op-grammar.rl:60

		atoms[atm].setTill(p)
		atoms[atm].parseValue(frame.Body)

//line ./op-grammar.rl:41

		atm++

		goto st20
	tr78:
//line ./op-grammar.rl:73

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

		goto st20
	tr88:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:73

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

		goto st20
	tr99:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:73

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

		goto st20
	tr110:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:51

		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:73

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

		goto st20
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
//line parser.go:1868
		switch (frame.Body)[p] {
		case 32:
			goto st20
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
			goto st20
		}
		goto st0
	tr30:
		(ps.state) = 21
//line ./op-grammar.rl:113

		(ps.state) = (RON_FULL_STOP)
		{
			p++
			goto _out
		}

		goto _again
	tr35:
		(ps.state) = 21
//line ./op-grammar.rl:106

		frame.position++

//line ./op-grammar.rl:113

		(ps.state) = (RON_FULL_STOP)
		{
			p++
			goto _out
		}

		goto _again
	tr45:
		(ps.state) = 21
//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:110

//line ./op-grammar.rl:106

		frame.position++

//line ./op-grammar.rl:113

		(ps.state) = (RON_FULL_STOP)
		{
			p++
			goto _out
		}

		goto _again
	tr58:
		(ps.state) = 21
//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:83

//line ./op-grammar.rl:106

		frame.position++

//line ./op-grammar.rl:113

		(ps.state) = (RON_FULL_STOP)
		{
			p++
			goto _out
		}

		goto _again
	tr66:
		(ps.state) = 21
//line ./op-grammar.rl:83

//line ./op-grammar.rl:106

		frame.position++

//line ./op-grammar.rl:113

		(ps.state) = (RON_FULL_STOP)
		{
			p++
			goto _out
		}

		goto _again
	tr74:
		(ps.state) = 21
//line ./op-grammar.rl:60

		atoms[atm].setTill(p)
		atoms[atm].parseValue(frame.Body)

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:83

//line ./op-grammar.rl:106

		frame.position++

//line ./op-grammar.rl:113

		(ps.state) = (RON_FULL_STOP)
		{
			p++
			goto _out
		}

		goto _again
	tr83:
		(ps.state) = 21
//line ./op-grammar.rl:73

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:83

//line ./op-grammar.rl:106

		frame.position++

//line ./op-grammar.rl:113

		(ps.state) = (RON_FULL_STOP)
		{
			p++
			goto _out
		}

		goto _again
	tr92:
		(ps.state) = 21
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:73

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:83

//line ./op-grammar.rl:106

		frame.position++

//line ./op-grammar.rl:113

		(ps.state) = (RON_FULL_STOP)
		{
			p++
			goto _out
		}

		goto _again
	tr105:
		(ps.state) = 21
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:73

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:83

//line ./op-grammar.rl:106

		frame.position++

//line ./op-grammar.rl:113

		(ps.state) = (RON_FULL_STOP)
		{
			p++
			goto _out
		}

		goto _again
	tr114:
		(ps.state) = 21
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:51

		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:73

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:83

//line ./op-grammar.rl:106

		frame.position++

//line ./op-grammar.rl:113

		(ps.state) = (RON_FULL_STOP)
		{
			p++
			goto _out
		}

		goto _again
	tr124:
		(ps.state) = 21
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:110

//line ./op-grammar.rl:106

		frame.position++

//line ./op-grammar.rl:113

		(ps.state) = (RON_FULL_STOP)
		{
			p++
			goto _out
		}

		goto _again
	tr136:
		(ps.state) = 21
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:110

//line ./op-grammar.rl:106

		frame.position++

//line ./op-grammar.rl:113

		(ps.state) = (RON_FULL_STOP)
		{
			p++
			goto _out
		}

		goto _again
	tr145:
		(ps.state) = 21
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:51

		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:110

//line ./op-grammar.rl:106

		frame.position++

//line ./op-grammar.rl:113

		(ps.state) = (RON_FULL_STOP)
		{
			p++
			goto _out
		}

		goto _again
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
//line parser.go:2180
		goto st0
	tr5:
		(ps.state) = 4
//line ./op-grammar.rl:90

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

//line ./op-grammar.rl:77

		atm = 4
		hlf = VALUE
		dgt = 0

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto _again
	tr36:
		(ps.state) = 4
//line ./op-grammar.rl:106

		frame.position++

//line ./op-grammar.rl:90

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

//line ./op-grammar.rl:77

		atm = 4
		hlf = VALUE
		dgt = 0

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto _again
	tr47:
//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:110

//line ./op-grammar.rl:77

		atm = 4
		hlf = VALUE
		dgt = 0

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st4
	tr59:
//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st4
	tr67:
//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st4
	tr75:
//line ./op-grammar.rl:60

		atoms[atm].setTill(p)
		atoms[atm].parseValue(frame.Body)

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st4
	tr85:
//line ./op-grammar.rl:73

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st4
	tr94:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:73

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st4
	tr107:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:73

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st4
	tr117:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:51

		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:73

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st4
	tr126:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:110

//line ./op-grammar.rl:77

		atm = 4
		hlf = VALUE
		dgt = 0

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st4
	tr138:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:110

//line ./op-grammar.rl:77

		atm = 4
		hlf = VALUE
		dgt = 0

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st4
	tr148:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:51

		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:110

//line ./op-grammar.rl:77

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
//line parser.go:2451
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

		atoms[atm].setIntType()
		atoms[atm].setFrom(p)

		goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
//line parser.go:2481
		if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
			goto st22
		}
		goto st0
	tr16:
//line ./op-grammar.rl:45

		atoms[atm].setIntType()
		atoms[atm].setFrom(p)

		goto st22
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
//line parser.go:2498
		switch (frame.Body)[p] {
		case 32:
			goto tr70
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
		switch {
		case (frame.Body)[p] > 13:
			if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
				goto st22
			}
		case (frame.Body)[p] >= 9:
			goto tr70
		}
		goto st0
	tr6:
		(ps.state) = 6
//line ./op-grammar.rl:90

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

//line ./op-grammar.rl:77

		atm = 4
		hlf = VALUE
		dgt = 0

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto _again
	tr37:
		(ps.state) = 6
//line ./op-grammar.rl:106

		frame.position++

//line ./op-grammar.rl:90

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

//line ./op-grammar.rl:77

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

//line ./op-grammar.rl:110

//line ./op-grammar.rl:77

		atm = 4
		hlf = VALUE
		dgt = 0

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st6
	tr60:
//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st6
	tr68:
//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st6
	tr76:
//line ./op-grammar.rl:60

		atoms[atm].setTill(p)
		atoms[atm].parseValue(frame.Body)

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st6
	tr86:
//line ./op-grammar.rl:73

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st6
	tr95:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:73

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st6
	tr108:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:73

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st6
	tr118:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:51

		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:73

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st6
	tr127:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:110

//line ./op-grammar.rl:77

		atm = 4
		hlf = VALUE
		dgt = 0

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st6
	tr139:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:110

//line ./op-grammar.rl:77

		atm = 4
		hlf = VALUE
		dgt = 0

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st6
	tr149:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:51

		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:110

//line ./op-grammar.rl:77

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
//line parser.go:2807
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
//line ./op-grammar.rl:65

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

		goto st23
	tr102:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:45

		hlf = 1
		atoms[atm][1] &= INT60_FULL
		atoms[atm][1] |= ((uint64)(ABC[(frame.Body)[p]])) << 60

		goto st23
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
//line parser.go:2890
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
			goto tr83
		case 58:
			goto tr80
		case 59:
			goto tr79
		case 61:
			goto tr85
		case 62:
			goto tr86
		case 63:
			goto tr79
		case 64:
			goto tr80
		case 91:
			goto tr82
		case 93:
			goto tr82
		case 94:
			goto tr87
		case 95:
			goto tr84
		case 123:
			goto tr82
		case 125:
			goto tr82
		case 126:
			goto tr84
		}
		switch {
		case (frame.Body)[p] < 48:
			switch {
			case (frame.Body)[p] > 13:
				if 40 <= (frame.Body)[p] && (frame.Body)[p] <= 41 {
					goto tr82
				}
			case (frame.Body)[p] >= 9:
				goto tr78
			}
		case (frame.Body)[p] > 57:
			switch {
			case (frame.Body)[p] > 90:
				if 97 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr84
				}
			case (frame.Body)[p] >= 65:
				goto tr84
			}
		default:
			goto tr84
		}
		goto st0
	tr82:
//line ././uuid-grammar.rl:29

		dgt = 0
		hlf = 1

//line ././uuid-grammar.rl:9

		dgt = int(ABC[(frame.Body)[p]]) + 4
		atoms[atm][hlf] &= INT60_FLAGS | PREFIX_MASKS[dgt]

		goto st24
	tr93:
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
	tr98:
//line ././uuid-grammar.rl:40

		atoms[atm][hlf] <<= 6
		dgt--

		goto st24
	tr104:
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
//line parser.go:3005
		switch (frame.Body)[p] {
		case 32:
			goto tr88
		case 33:
			goto tr89
		case 35:
			goto tr90
		case 39:
			goto tr91
		case 42:
			goto tr90
		case 44:
			goto tr89
		case 46:
			goto tr92
		case 58:
			goto tr90
		case 59:
			goto tr89
		case 61:
			goto tr94
		case 62:
			goto tr95
		case 63:
			goto tr89
		case 64:
			goto tr90
		case 94:
			goto tr96
		case 95:
			goto tr93
		case 126:
			goto tr93
		}
		switch {
		case (frame.Body)[p] < 48:
			if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
				goto tr88
			}
		case (frame.Body)[p] > 57:
			switch {
			case (frame.Body)[p] > 90:
				if 97 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr93
				}
			case (frame.Body)[p] >= 65:
				goto tr93
			}
		default:
			goto tr93
		}
		goto st0
	tr7:
		(ps.state) = 7
//line ./op-grammar.rl:90

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

//line ./op-grammar.rl:77

		atm = 4
		hlf = VALUE
		dgt = 0

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto _again
	tr38:
		(ps.state) = 7
//line ./op-grammar.rl:106

		frame.position++

//line ./op-grammar.rl:90

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

//line ./op-grammar.rl:77

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

//line ./op-grammar.rl:110

//line ./op-grammar.rl:77

		atm = 4
		hlf = VALUE
		dgt = 0

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st7
	tr61:
//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st7
	tr69:
//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st7
	tr77:
//line ./op-grammar.rl:60

		atoms[atm].setTill(p)
		atoms[atm].parseValue(frame.Body)

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st7
	tr87:
//line ./op-grammar.rl:73

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st7
	tr96:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:73

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st7
	tr109:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:73

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st7
	tr119:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:51

		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:73

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

		atm++

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st7
	tr128:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:110

//line ./op-grammar.rl:77

		atm = 4
		hlf = VALUE
		dgt = 0

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st7
	tr140:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:110

//line ./op-grammar.rl:77

		atm = 4
		hlf = VALUE
		dgt = 0

//line ./op-grammar.rl:36

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st7
	tr150:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:51

		atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:32

		atm++

//line ./op-grammar.rl:110

//line ./op-grammar.rl:77

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
//line parser.go:3327
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
//line ./op-grammar.rl:50

		atoms[atm].setFloatType()
		atoms[atm].setFrom(p)

		goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
//line parser.go:3357
		if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
			goto st9
		}
		goto st0
	tr24:
//line ./op-grammar.rl:50

		atoms[atm].setFloatType()
		atoms[atm].setFrom(p)

		goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
//line parser.go:3374
		switch (frame.Body)[p] {
		case 46:
			goto st10
		case 69:
			goto st12
		case 101:
			goto st12
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
			goto st25
		}
		goto st0
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
		switch (frame.Body)[p] {
		case 32:
			goto tr70
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
		case 69:
			goto st11
		case 94:
			goto tr77
		case 101:
			goto st11
		}
		switch {
		case (frame.Body)[p] > 13:
			if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
				goto st25
			}
		case (frame.Body)[p] >= 9:
			goto tr70
		}
		goto st0
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		switch (frame.Body)[p] {
		case 43:
			goto st5
		case 45:
			goto st5
		}
		if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
			goto st22
		}
		goto st0
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		switch (frame.Body)[p] {
		case 43:
			goto st10
		case 45:
			goto st10
		}
		if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
			goto st25
		}
		goto st0
	tr84:
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
				(ps.state) = 26
				goto _out
			}
		}

		goto st26
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
//line parser.go:3498
		switch (frame.Body)[p] {
		case 32:
			goto tr88
		case 33:
			goto tr89
		case 35:
			goto tr90
		case 39:
			goto tr91
		case 42:
			goto tr90
		case 44:
			goto tr89
		case 46:
			goto tr92
		case 47:
			goto tr98
		case 58:
			goto tr90
		case 59:
			goto tr89
		case 61:
			goto tr94
		case 62:
			goto tr95
		case 63:
			goto tr89
		case 64:
			goto tr90
		case 94:
			goto tr96
		case 95:
			goto tr93
		case 126:
			goto tr93
		}
		switch {
		case (frame.Body)[p] < 48:
			if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
				goto tr88
			}
		case (frame.Body)[p] > 57:
			switch {
			case (frame.Body)[p] > 90:
				if 97 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr93
				}
			case (frame.Body)[p] >= 65:
				goto tr93
			}
		default:
			goto tr93
		}
		goto st0
	tr20:
//line ./op-grammar.rl:65

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
	tr106:
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
//line parser.go:3589
		switch (frame.Body)[p] {
		case 32:
			goto tr99
		case 33:
			goto tr100
		case 35:
			goto tr101
		case 39:
			goto tr103
		case 42:
			goto tr101
		case 44:
			goto tr100
		case 46:
			goto tr105
		case 58:
			goto tr101
		case 59:
			goto tr100
		case 61:
			goto tr107
		case 62:
			goto tr108
		case 63:
			goto tr100
		case 64:
			goto tr101
		case 91:
			goto tr104
		case 93:
			goto tr104
		case 94:
			goto tr109
		case 95:
			goto tr106
		case 123:
			goto tr104
		case 125:
			goto tr104
		case 126:
			goto tr106
		}
		switch {
		case (frame.Body)[p] < 43:
			switch {
			case (frame.Body)[p] < 36:
				if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
					goto tr99
				}
			case (frame.Body)[p] > 37:
				if 40 <= (frame.Body)[p] && (frame.Body)[p] <= 41 {
					goto tr104
				}
			default:
				goto tr102
			}
		case (frame.Body)[p] > 45:
			switch {
			case (frame.Body)[p] < 65:
				if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
					goto tr106
				}
			case (frame.Body)[p] > 90:
				if 97 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr106
				}
			default:
				goto tr106
			}
		default:
			goto tr102
		}
		goto st0
	tr21:
//line ./op-grammar.rl:65

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
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
//line parser.go:3696
		switch (frame.Body)[p] {
		case 32:
			goto tr110
		case 33:
			goto tr111
		case 35:
			goto tr112
		case 39:
			goto tr113
		case 42:
			goto tr112
		case 44:
			goto tr111
		case 46:
			goto tr114
		case 47:
			goto tr115
		case 58:
			goto tr112
		case 59:
			goto tr111
		case 61:
			goto tr117
		case 62:
			goto tr118
		case 63:
			goto tr111
		case 64:
			goto tr112
		case 91:
			goto tr104
		case 93:
			goto tr104
		case 94:
			goto tr119
		case 95:
			goto tr116
		case 123:
			goto tr104
		case 125:
			goto tr104
		case 126:
			goto tr116
		}
		switch {
		case (frame.Body)[p] < 43:
			switch {
			case (frame.Body)[p] < 36:
				if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
					goto tr110
				}
			case (frame.Body)[p] > 37:
				if 40 <= (frame.Body)[p] && (frame.Body)[p] <= 41 {
					goto tr104
				}
			default:
				goto tr102
			}
		case (frame.Body)[p] > 45:
			switch {
			case (frame.Body)[p] < 65:
				if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
					goto tr116
				}
			case (frame.Body)[p] > 90:
				if 97 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr116
				}
			default:
				goto tr116
			}
		default:
			goto tr102
		}
		goto st0
	tr116:
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
	tr115:
//line ././uuid-grammar.rl:40

		atoms[atm][hlf] <<= 6
		dgt--

		goto st29
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
//line parser.go:3794
		switch (frame.Body)[p] {
		case 32:
			goto tr110
		case 33:
			goto tr111
		case 35:
			goto tr112
		case 39:
			goto tr113
		case 42:
			goto tr112
		case 44:
			goto tr111
		case 46:
			goto tr114
		case 58:
			goto tr112
		case 59:
			goto tr111
		case 61:
			goto tr117
		case 62:
			goto tr118
		case 63:
			goto tr111
		case 64:
			goto tr112
		case 91:
			goto tr104
		case 93:
			goto tr104
		case 94:
			goto tr119
		case 95:
			goto tr116
		case 123:
			goto tr104
		case 125:
			goto tr104
		case 126:
			goto tr116
		}
		switch {
		case (frame.Body)[p] < 43:
			switch {
			case (frame.Body)[p] < 36:
				if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
					goto tr110
				}
			case (frame.Body)[p] > 37:
				if 40 <= (frame.Body)[p] && (frame.Body)[p] <= 41 {
					goto tr104
				}
			default:
				goto tr102
			}
		case (frame.Body)[p] > 45:
			switch {
			case (frame.Body)[p] < 65:
				if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
					goto tr116
				}
			case (frame.Body)[p] > 90:
				if 97 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr116
				}
			default:
				goto tr116
			}
		default:
			goto tr102
		}
		goto st0
	tr10:
//line ./op-grammar.rl:55

		atoms[atm].setStringType()
		atoms[atm].setFrom(p)

		goto st13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
//line parser.go:3880
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

		goto st30
	tr125:
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
	tr129:
//line ././uuid-grammar.rl:40

		atoms[atm][hlf] <<= 6
		dgt--

		goto st30
	tr135:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:29

		dgt = 0
		hlf = 1

//line ././uuid-grammar.rl:9

		dgt = int(ABC[(frame.Body)[p]]) + 4
		atoms[atm][hlf] &= INT60_FLAGS | PREFIX_MASKS[dgt]

		goto st30
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
//line parser.go:3937
		switch (frame.Body)[p] {
		case 32:
			goto tr120
		case 33:
			goto tr121
		case 35:
			goto tr122
		case 39:
			goto tr123
		case 42:
			goto tr122
		case 44:
			goto tr121
		case 46:
			goto tr124
		case 58:
			goto tr122
		case 59:
			goto tr121
		case 61:
			goto tr126
		case 62:
			goto tr127
		case 63:
			goto tr121
		case 64:
			goto tr122
		case 94:
			goto tr128
		case 95:
			goto tr125
		case 126:
			goto tr125
		}
		switch {
		case (frame.Body)[p] < 48:
			if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
				goto tr120
			}
		case (frame.Body)[p] > 57:
			switch {
			case (frame.Body)[p] > 90:
				if 97 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr125
				}
			case (frame.Body)[p] >= 65:
				goto tr125
			}
		default:
			goto tr125
		}
		goto st0
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
//line parser.go:4014
		switch (frame.Body)[p] {
		case 32:
			goto tr120
		case 33:
			goto tr121
		case 35:
			goto tr122
		case 39:
			goto tr123
		case 42:
			goto tr122
		case 44:
			goto tr121
		case 46:
			goto tr124
		case 47:
			goto tr129
		case 58:
			goto tr122
		case 59:
			goto tr121
		case 61:
			goto tr126
		case 62:
			goto tr127
		case 63:
			goto tr121
		case 64:
			goto tr122
		case 94:
			goto tr128
		case 95:
			goto tr125
		case 126:
			goto tr125
		}
		switch {
		case (frame.Body)[p] < 48:
			if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
				goto tr120
			}
		case (frame.Body)[p] > 57:
			switch {
			case (frame.Body)[p] > 90:
				if 97 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr125
				}
			case (frame.Body)[p] >= 65:
				goto tr125
			}
		default:
			goto tr125
		}
		goto st0
	tr44:
//line ././uuid-grammar.rl:5

//line ././uuid-grammar.rl:26

//line ././uuid-grammar.rl:9

		dgt = int(ABC[(frame.Body)[p]]) + 4
		atoms[atm][hlf] &= INT60_FLAGS | PREFIX_MASKS[dgt]

		goto st32
	tr137:
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
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
//line parser.go:4097
		switch (frame.Body)[p] {
		case 32:
			goto tr130
		case 33:
			goto tr131
		case 35:
			goto tr132
		case 39:
			goto tr134
		case 42:
			goto tr132
		case 44:
			goto tr131
		case 46:
			goto tr136
		case 58:
			goto tr132
		case 59:
			goto tr131
		case 61:
			goto tr138
		case 62:
			goto tr139
		case 63:
			goto tr131
		case 64:
			goto tr132
		case 91:
			goto tr135
		case 93:
			goto tr135
		case 94:
			goto tr140
		case 95:
			goto tr137
		case 123:
			goto tr135
		case 125:
			goto tr135
		case 126:
			goto tr137
		}
		switch {
		case (frame.Body)[p] < 43:
			switch {
			case (frame.Body)[p] < 36:
				if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
					goto tr130
				}
			case (frame.Body)[p] > 37:
				if 40 <= (frame.Body)[p] && (frame.Body)[p] <= 41 {
					goto tr135
				}
			default:
				goto tr133
			}
		case (frame.Body)[p] > 45:
			switch {
			case (frame.Body)[p] < 65:
				if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
					goto tr137
				}
			case (frame.Body)[p] > 90:
				if 97 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr137
				}
			default:
				goto tr137
			}
		default:
			goto tr133
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
//line parser.go:4196
		switch (frame.Body)[p] {
		case 32:
			goto tr141
		case 33:
			goto tr142
		case 35:
			goto tr143
		case 39:
			goto tr144
		case 42:
			goto tr143
		case 44:
			goto tr142
		case 46:
			goto tr145
		case 47:
			goto tr146
		case 58:
			goto tr143
		case 59:
			goto tr142
		case 61:
			goto tr148
		case 62:
			goto tr149
		case 63:
			goto tr142
		case 64:
			goto tr143
		case 91:
			goto tr135
		case 93:
			goto tr135
		case 94:
			goto tr150
		case 95:
			goto tr147
		case 123:
			goto tr135
		case 125:
			goto tr135
		case 126:
			goto tr147
		}
		switch {
		case (frame.Body)[p] < 43:
			switch {
			case (frame.Body)[p] < 36:
				if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
					goto tr141
				}
			case (frame.Body)[p] > 37:
				if 40 <= (frame.Body)[p] && (frame.Body)[p] <= 41 {
					goto tr135
				}
			default:
				goto tr133
			}
		case (frame.Body)[p] > 45:
			switch {
			case (frame.Body)[p] < 65:
				if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
					goto tr147
				}
			case (frame.Body)[p] > 90:
				if 97 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr147
				}
			default:
				goto tr147
			}
		default:
			goto tr133
		}
		goto st0
	tr147:
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
	tr146:
//line ././uuid-grammar.rl:40

		atoms[atm][hlf] <<= 6
		dgt--

		goto st34
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
//line parser.go:4294
		switch (frame.Body)[p] {
		case 32:
			goto tr141
		case 33:
			goto tr142
		case 35:
			goto tr143
		case 39:
			goto tr144
		case 42:
			goto tr143
		case 44:
			goto tr142
		case 46:
			goto tr145
		case 58:
			goto tr143
		case 59:
			goto tr142
		case 61:
			goto tr148
		case 62:
			goto tr149
		case 63:
			goto tr142
		case 64:
			goto tr143
		case 91:
			goto tr135
		case 93:
			goto tr135
		case 94:
			goto tr150
		case 95:
			goto tr147
		case 123:
			goto tr135
		case 125:
			goto tr135
		case 126:
			goto tr147
		}
		switch {
		case (frame.Body)[p] < 43:
			switch {
			case (frame.Body)[p] < 36:
				if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
					goto tr141
				}
			case (frame.Body)[p] > 37:
				if 40 <= (frame.Body)[p] && (frame.Body)[p] <= 41 {
					goto tr135
				}
			default:
				goto tr133
			}
		case (frame.Body)[p] > 45:
			switch {
			case (frame.Body)[p] < 65:
				if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
					goto tr147
				}
			case (frame.Body)[p] > 90:
				if 97 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr147
				}
			default:
				goto tr147
			}
		default:
			goto tr133
		}
		goto st0
	tr50:
//line ./op-grammar.rl:10

		if atm > 0 {
			atoms[atm] = atoms[atm-1]
		}

		goto st35
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
//line parser.go:4381
		switch (frame.Body)[p] {
		case 32:
			goto tr151
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
					goto tr151
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
	_test_eof29:
		(ps.state) = 29
		goto _test_eof
	_test_eof13:
		(ps.state) = 13
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
	_test_eof33:
		(ps.state) = 33
		goto _test_eof
	_test_eof34:
		(ps.state) = 34
		goto _test_eof
	_test_eof35:
		(ps.state) = 35
		goto _test_eof

	_test_eof:
		{
		}
		if p == eof {
			switch ps.state {
			case 15:
//line ./op-grammar.rl:106

				frame.position++

			case 20:
//line ./op-grammar.rl:83

//line ./op-grammar.rl:106

				frame.position++

			case 16, 17, 18, 35:
//line ./op-grammar.rl:32

				atm++

//line ./op-grammar.rl:110

//line ./op-grammar.rl:106

				frame.position++

			case 19:
//line ./op-grammar.rl:41

				atm++

//line ./op-grammar.rl:83

//line ./op-grammar.rl:106

				frame.position++

			case 32:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:32

				atm++

//line ./op-grammar.rl:110

//line ./op-grammar.rl:106

				frame.position++

			case 30, 31:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:32

				atm++

//line ./op-grammar.rl:110

//line ./op-grammar.rl:106

				frame.position++

			case 22, 25:
//line ./op-grammar.rl:60

				atoms[atm].setTill(p)
				atoms[atm].parseValue(frame.Body)

//line ./op-grammar.rl:41

				atm++

//line ./op-grammar.rl:83

//line ./op-grammar.rl:106

				frame.position++

			case 23:
//line ./op-grammar.rl:73

				atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

				atm++

//line ./op-grammar.rl:83

//line ./op-grammar.rl:106

				frame.position++

			case 33, 34:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:51

				atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:32

				atm++

//line ./op-grammar.rl:110

//line ./op-grammar.rl:106

				frame.position++

			case 27:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:73

				atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

				atm++

//line ./op-grammar.rl:83

//line ./op-grammar.rl:106

				frame.position++

			case 24, 26:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:73

				atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

				atm++

//line ./op-grammar.rl:83

//line ./op-grammar.rl:106

				frame.position++

			case 28, 29:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:51

				atoms[atm][1] = UUID_NAME_FLAG

//line ./op-grammar.rl:73

				atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

//line ./op-grammar.rl:41

				atm++

//line ./op-grammar.rl:83

//line ./op-grammar.rl:106

				frame.position++

//line parser.go:4675
			}
		}

	_out:
		{
		}
	}

//line dfa.rl:78

	frame.descriptor.setTill(p)
	frame.descriptor[VALUE] |= uint64(len(atoms)) << 32

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

//line dfa.rl:116

//line parser.go:4723
	const UUID_start int = 1
	const UUID_first_final int = 2
	const UUID_error int = 0

	const UUID_en_main int = 1

//line dfa.rl:117

	cs, p, pe, eof := 0, 0, len(data), len(data)
	_ = eof

	atm, hlf, dgt := 0, 0, 0

	atoms := [1]Atom{Atom(ctx_uuid)}

//line parser.go:4741
	{
		cs = UUID_start
	}

//line parser.go:4746
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
//line parser.go:4843
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
//line parser.go:4925
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
//line parser.go:4969
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
//line parser.go:5019
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
//line parser.go:5086
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
//line parser.go:5152
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

//line parser.go:5221
			}
		}

	_out:
		{
		}
	}

//line dfa.rl:132

	if cs < UUID_first_final || dgt > 10 {
		return ERROR_UUID, errors.New(fmt.Sprintf("parse error at pos %d", p))
	} else {
		return UUID(atoms[0]), nil
	}

}
