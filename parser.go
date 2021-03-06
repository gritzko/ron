package ron

import "fmt"
import "errors"

const RON_start int = 13
const RON_first_final int = 13
const RON_error int = 0

const RON_en_main int = 13

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

		{
			(ps.state) = RON_start
		}

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

	{
		if p == pe {
			goto _test_eof
		}
		goto _resume

	_again:
		switch ps.state {
		case 13:
			goto st13
		case 0:
			goto st0
		case 1:
			goto st1
		case 14:
			goto st14
		case 15:
			goto st15
		case 16:
			goto st16
		case 17:
			goto st17
		case 2:
			goto st2
		case 3:
			goto st3
		case 18:
			goto st18
		case 19:
			goto st19
		case 20:
			goto st20
		case 4:
			goto st4
		case 5:
			goto st5
		case 21:
			goto st21
		case 6:
			goto st6
		case 22:
			goto st22
		case 23:
			goto st23
		case 7:
			goto st7
		case 8:
			goto st8
		case 9:
			goto st9
		case 10:
			goto st10
		case 24:
			goto st24
		case 11:
			goto st11
		case 25:
			goto st25
		case 26:
			goto st26
		case 27:
			goto st27
		case 28:
			goto st28
		case 12:
			goto st12
		case 29:
			goto st29
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
		}

		if p++; p == pe {
			goto _test_eof
		}
	_resume:
		switch ps.state {
		case 13:
			goto st_case_13
		case 0:
			goto st_case_0
		case 1:
			goto st_case_1
		case 14:
			goto st_case_14
		case 15:
			goto st_case_15
		case 16:
			goto st_case_16
		case 17:
			goto st_case_17
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		case 18:
			goto st_case_18
		case 19:
			goto st_case_19
		case 20:
			goto st_case_20
		case 4:
			goto st_case_4
		case 5:
			goto st_case_5
		case 21:
			goto st_case_21
		case 6:
			goto st_case_6
		case 22:
			goto st_case_22
		case 23:
			goto st_case_23
		case 7:
			goto st_case_7
		case 8:
			goto st_case_8
		case 9:
			goto st_case_9
		case 10:
			goto st_case_10
		case 24:
			goto st_case_24
		case 11:
			goto st_case_11
		case 25:
			goto st_case_25
		case 26:
			goto st_case_26
		case 27:
			goto st_case_27
		case 28:
			goto st_case_28
		case 12:
			goto st_case_12
		case 29:
			goto st_case_29
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
		}
		goto st_out
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
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
	tr150:

		atm++

		frame.position++

		goto st1
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
		(ps.state) = 14

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

		frame.term = int(ABC[(frame.Body)[p]])

		goto _again
	tr32:
		(ps.state) = 14

		frame.position++

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

		frame.term = int(ABC[(frame.Body)[p]])

		goto _again
	tr40:

		atm++

		frame.term = int(ABC[(frame.Body)[p]])

		goto st14
	tr55:

		atm++

		frame.term = int(ABC[(frame.Body)[p]])

		goto st14
	tr63:

		frame.term = int(ABC[(frame.Body)[p]])

		goto st14
	tr71:

		atoms[atm].setTill(p)
		atoms[atm].parseValue(frame.Body)

		atm++

		frame.term = int(ABC[(frame.Body)[p]])

		goto st14
	tr79:

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

		atm++

		frame.term = int(ABC[(frame.Body)[p]])

		goto st14
	tr89:

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

		atm++

		frame.term = int(ABC[(frame.Body)[p]])

		goto st14
	tr99:

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

		atm++

		frame.term = int(ABC[(frame.Body)[p]])

		goto st14
	tr110:

		atoms[atm][1] = UUID_NAME_FLAG

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

		atm++

		frame.term = int(ABC[(frame.Body)[p]])

		goto st14
	tr120:

		atm++

		frame.term = int(ABC[(frame.Body)[p]])

		goto st14
	tr130:

		atm++

		frame.term = int(ABC[(frame.Body)[p]])

		goto st14
	tr141:

		atoms[atm][1] = UUID_NAME_FLAG

		atm++

		frame.term = int(ABC[(frame.Body)[p]])

		goto st14
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		switch (frame.Body)[p] {
		case 32:
			goto st14
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
			goto st14
		}
		goto st0
	tr3:
		(ps.state) = 15

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

		ps.omitted = 15

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
		(ps.state) = 15

		frame.position++

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

		ps.omitted = 15

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
		(ps.state) = 15

		atm++

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
		(ps.state) = 15

		atm++

		frame.position++

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

		ps.omitted = 15

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
		(ps.state) = 15

		frame.position++

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

		ps.omitted = 15

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
		(ps.state) = 15

		atoms[atm].setTill(p)
		atoms[atm].parseValue(frame.Body)

		atm++

		frame.position++

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

		ps.omitted = 15

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
		(ps.state) = 15

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

		atm++

		frame.position++

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

		ps.omitted = 15

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
		(ps.state) = 15

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

		atm++

		frame.position++

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

		ps.omitted = 15

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
		(ps.state) = 15

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

		atm++

		frame.position++

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

		ps.omitted = 15

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
	tr111:
		(ps.state) = 15

		atoms[atm][1] = UUID_NAME_FLAG

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

		atm++

		frame.position++

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

		ps.omitted = 15

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
		(ps.state) = 15

		atm++

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
		(ps.state) = 15

		atm++

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
	tr142:
		(ps.state) = 15

		atoms[atm][1] = UUID_NAME_FLAG

		atm++

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
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
		switch (frame.Body)[p] {
		case 32:
			goto st15
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
					goto st15
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

		hlf = 1
		atoms[atm][1] &= INT60_FULL
		atoms[atm][1] |= ((uint64)(ABC[(frame.Body)[p]])) << 60

		goto st16
	tr132:

		hlf = 1
		atoms[atm][1] &= INT60_FULL
		atoms[atm][1] |= ((uint64)(ABC[(frame.Body)[p]])) << 60

		goto st16
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		switch (frame.Body)[p] {
		case 32:
			goto st17
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
				goto st17
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
	tr119:

		goto st17
	tr129:

		goto st17
	tr140:

		atoms[atm][1] = UUID_NAME_FLAG

		goto st17
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
		switch (frame.Body)[p] {
		case 32:
			goto st17
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
			goto st17
		}
		goto st0
	tr4:
		(ps.state) = 2

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

		atm = 4
		hlf = VALUE
		dgt = 0

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto _again
	tr34:
		(ps.state) = 2

		frame.position++

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

		atm = 4
		hlf = VALUE
		dgt = 0

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto _again
	tr43:

		atm++

		atm = 4
		hlf = VALUE
		dgt = 0

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st2
	tr57:

		atm++

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st2
	tr65:

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st2
	tr73:

		atoms[atm].setTill(p)
		atoms[atm].parseValue(frame.Body)

		atm++

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st2
	tr81:

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

		atm++

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st2
	tr91:

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

		atm++

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st2
	tr102:

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

		atm++

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st2
	tr112:

		atoms[atm][1] = UUID_NAME_FLAG

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

		atm++

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st2
	tr122:

		atm++

		atm = 4
		hlf = VALUE
		dgt = 0

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st2
	tr133:

		atm++

		atm = 4
		hlf = VALUE
		dgt = 0

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st2
	tr143:

		atoms[atm][1] = UUID_NAME_FLAG

		atm++

		atm = 4
		hlf = VALUE
		dgt = 0

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
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

		atoms[atm].setType(ATOM_STRING)
		atoms[atm].setFrom(p)

		goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
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
			goto st12
		}
		goto st3
	tr9:

		atoms[atm].setType(ATOM_STRING)
		atoms[atm].setFrom(p)

		atoms[atm].setTill(p)
		atoms[atm].parseValue(frame.Body)

		goto st18
	tr12:

		atoms[atm].setTill(p)
		atoms[atm].parseValue(frame.Body)

		goto st18
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
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

		atm++

		goto st19
	tr70:

		atoms[atm].setTill(p)
		atoms[atm].parseValue(frame.Body)

		atm++

		goto st19
	tr78:

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

		atm++

		goto st19
	tr88:

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

		atm++

		goto st19
	tr98:

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

		atm++

		goto st19
	tr109:

		atoms[atm][1] = UUID_NAME_FLAG

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

		atm++

		goto st19
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
		switch (frame.Body)[p] {
		case 32:
			goto st19
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
			goto st19
		}
		goto st0
	tr30:
		(ps.state) = 20

		(ps.state) = (RON_FULL_STOP)
		{
			p++
			goto _out
		}

		goto _again
	tr35:
		(ps.state) = 20

		frame.position++

		(ps.state) = (RON_FULL_STOP)
		{
			p++
			goto _out
		}

		goto _again
	tr45:
		(ps.state) = 20

		atm++

		frame.position++

		(ps.state) = (RON_FULL_STOP)
		{
			p++
			goto _out
		}

		goto _again
	tr58:
		(ps.state) = 20

		atm++

		frame.position++

		(ps.state) = (RON_FULL_STOP)
		{
			p++
			goto _out
		}

		goto _again
	tr66:
		(ps.state) = 20

		frame.position++

		(ps.state) = (RON_FULL_STOP)
		{
			p++
			goto _out
		}

		goto _again
	tr74:
		(ps.state) = 20

		atoms[atm].setTill(p)
		atoms[atm].parseValue(frame.Body)

		atm++

		frame.position++

		(ps.state) = (RON_FULL_STOP)
		{
			p++
			goto _out
		}

		goto _again
	tr83:
		(ps.state) = 20

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

		atm++

		frame.position++

		(ps.state) = (RON_FULL_STOP)
		{
			p++
			goto _out
		}

		goto _again
	tr92:
		(ps.state) = 20

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

		atm++

		frame.position++

		(ps.state) = (RON_FULL_STOP)
		{
			p++
			goto _out
		}

		goto _again
	tr104:
		(ps.state) = 20

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

		atm++

		frame.position++

		(ps.state) = (RON_FULL_STOP)
		{
			p++
			goto _out
		}

		goto _again
	tr113:
		(ps.state) = 20

		atoms[atm][1] = UUID_NAME_FLAG

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

		atm++

		frame.position++

		(ps.state) = (RON_FULL_STOP)
		{
			p++
			goto _out
		}

		goto _again
	tr123:
		(ps.state) = 20

		atm++

		frame.position++

		(ps.state) = (RON_FULL_STOP)
		{
			p++
			goto _out
		}

		goto _again
	tr135:
		(ps.state) = 20

		atm++

		frame.position++

		(ps.state) = (RON_FULL_STOP)
		{
			p++
			goto _out
		}

		goto _again
	tr144:
		(ps.state) = 20

		atoms[atm][1] = UUID_NAME_FLAG

		atm++

		frame.position++

		(ps.state) = (RON_FULL_STOP)
		{
			p++
			goto _out
		}

		goto _again
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
		goto st0
	tr5:
		(ps.state) = 4

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

		atm = 4
		hlf = VALUE
		dgt = 0

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto _again
	tr36:
		(ps.state) = 4

		frame.position++

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

		atm = 4
		hlf = VALUE
		dgt = 0

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto _again
	tr47:

		atm++

		atm = 4
		hlf = VALUE
		dgt = 0

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st4
	tr59:

		atm++

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st4
	tr67:

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st4
	tr75:

		atoms[atm].setTill(p)
		atoms[atm].parseValue(frame.Body)

		atm++

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st4
	tr85:

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

		atm++

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st4
	tr94:

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

		atm++

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st4
	tr106:

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

		atm++

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st4
	tr116:

		atoms[atm][1] = UUID_NAME_FLAG

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

		atm++

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st4
	tr125:

		atm++

		atm = 4
		hlf = VALUE
		dgt = 0

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st4
	tr137:

		atm++

		atm = 4
		hlf = VALUE
		dgt = 0

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st4
	tr147:

		atoms[atm][1] = UUID_NAME_FLAG

		atm++

		atm = 4
		hlf = VALUE
		dgt = 0

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
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

		atoms[atm].setType(ATOM_INT)
		atoms[atm].setFrom(p)

		goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
			goto st21
		}
		goto st0
	tr16:

		atoms[atm].setType(ATOM_INT)
		atoms[atm].setFrom(p)

		goto st21
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
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
				goto st21
			}
		case (frame.Body)[p] >= 9:
			goto tr70
		}
		goto st0
	tr6:
		(ps.state) = 6

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

		atm = 4
		hlf = VALUE
		dgt = 0

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto _again
	tr37:
		(ps.state) = 6

		frame.position++

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

		atm = 4
		hlf = VALUE
		dgt = 0

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto _again
	tr48:

		atm++

		atm = 4
		hlf = VALUE
		dgt = 0

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st6
	tr60:

		atm++

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st6
	tr68:

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st6
	tr76:

		atoms[atm].setTill(p)
		atoms[atm].parseValue(frame.Body)

		atm++

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st6
	tr86:

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

		atm++

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st6
	tr95:

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

		atm++

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st6
	tr107:

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

		atm++

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st6
	tr117:

		atoms[atm][1] = UUID_NAME_FLAG

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

		atm++

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st6
	tr126:

		atm++

		atm = 4
		hlf = VALUE
		dgt = 0

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st6
	tr138:

		atm++

		atm = 4
		hlf = VALUE
		dgt = 0

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st6
	tr148:

		atoms[atm][1] = UUID_NAME_FLAG

		atm++

		atm = 4
		hlf = VALUE
		dgt = 0

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
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

		if atm == 4 {
			atoms[atm] = atoms[SPEC_OBJECT]
		} else if atoms[atm-1].Type() == ATOM_UUID {
			atoms[atm] = atoms[atm-1]
		}

		hlf = 1
		atoms[atm][1] &= INT60_FULL
		atoms[atm][1] |= ((uint64)(ABC[(frame.Body)[p]])) << 60

		goto st22
	tr101:

		hlf = 1
		atoms[atm][1] &= INT60_FULL
		atoms[atm][1] |= ((uint64)(ABC[(frame.Body)[p]])) << 60

		goto st22
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
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

		dgt = 0
		hlf = 1

		dgt = int(ABC[(frame.Body)[p]]) + 4
		atoms[atm][hlf] &= INT60_FLAGS | PREFIX_MASKS[dgt]

		goto st23
	tr93:

		atoms[atm][hlf] |= ((uint64)(ABC[(frame.Body)[p]])) << DIGIT_OFFSETS[dgt]
		dgt++
		if dgt > 10 {
			{
				p++
				(ps.state) = 23
				goto _out
			}
		}

		goto st23
	tr97:

		atoms[atm][hlf] <<= 6
		dgt--

		goto st23
	tr103:

		dgt = 0
		hlf = 1

		dgt = int(ABC[(frame.Body)[p]]) + 4
		atoms[atm][hlf] &= INT60_FLAGS | PREFIX_MASKS[dgt]

		goto st23
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
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

		atm = 4
		hlf = VALUE
		dgt = 0

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto _again
	tr38:
		(ps.state) = 7

		frame.position++

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

		atm = 4
		hlf = VALUE
		dgt = 0

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto _again
	tr49:

		atm++

		atm = 4
		hlf = VALUE
		dgt = 0

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st7
	tr61:

		atm++

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st7
	tr69:

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st7
	tr77:

		atoms[atm].setTill(p)
		atoms[atm].parseValue(frame.Body)

		atm++

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st7
	tr87:

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

		atm++

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st7
	tr96:

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

		atm++

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st7
	tr108:

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

		atm++

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st7
	tr118:

		atoms[atm][1] = UUID_NAME_FLAG

		atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

		atm++

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st7
	tr127:

		atm++

		atm = 4
		hlf = VALUE
		dgt = 0

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st7
	tr139:

		atm++

		atm = 4
		hlf = VALUE
		dgt = 0

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st7
	tr149:

		atoms[atm][1] = UUID_NAME_FLAG

		atm++

		atm = 4
		hlf = VALUE
		dgt = 0

		hlf, dgt = VALUE, 0
		atoms = append(atoms, Atom{})

		goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
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

		atoms[atm].setType(ATOM_FLOAT)
		atoms[atm].setFrom(p)

		goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
			goto st9
		}
		goto st0
	tr24:

		atoms[atm].setType(ATOM_FLOAT)
		atoms[atm].setFrom(p)

		goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		switch (frame.Body)[p] {
		case 46:
			goto st10
		case 69:
			goto st11
		case 101:
			goto st11
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
			goto st24
		}
		goto st0
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
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
				goto st24
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
			goto st21
		}
		goto st0
	tr84:

		dgt = 0
		hlf = 1

		atoms[atm][hlf] &= INT60_FLAGS

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
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
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
			goto tr97
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

		if atm == 4 {
			atoms[atm] = atoms[SPEC_OBJECT]
		} else if atoms[atm-1].Type() == ATOM_UUID {
			atoms[atm] = atoms[atm-1]
		}

		dgt = int(ABC[(frame.Body)[p]]) + 4
		atoms[atm][hlf] &= INT60_FLAGS | PREFIX_MASKS[dgt]

		goto st26
	tr105:

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
		switch (frame.Body)[p] {
		case 32:
			goto tr98
		case 33:
			goto tr99
		case 35:
			goto tr100
		case 39:
			goto tr102
		case 42:
			goto tr100
		case 44:
			goto tr99
		case 46:
			goto tr104
		case 58:
			goto tr100
		case 59:
			goto tr99
		case 61:
			goto tr106
		case 62:
			goto tr107
		case 63:
			goto tr99
		case 64:
			goto tr100
		case 91:
			goto tr103
		case 93:
			goto tr103
		case 94:
			goto tr108
		case 95:
			goto tr105
		case 123:
			goto tr103
		case 125:
			goto tr103
		case 126:
			goto tr105
		}
		switch {
		case (frame.Body)[p] < 43:
			switch {
			case (frame.Body)[p] < 36:
				if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
					goto tr98
				}
			case (frame.Body)[p] > 37:
				if 40 <= (frame.Body)[p] && (frame.Body)[p] <= 41 {
					goto tr103
				}
			default:
				goto tr101
			}
		case (frame.Body)[p] > 45:
			switch {
			case (frame.Body)[p] < 65:
				if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
					goto tr105
				}
			case (frame.Body)[p] > 90:
				if 97 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr105
				}
			default:
				goto tr105
			}
		default:
			goto tr101
		}
		goto st0
	tr21:

		if atm == 4 {
			atoms[atm] = atoms[SPEC_OBJECT]
		} else if atoms[atm-1].Type() == ATOM_UUID {
			atoms[atm] = atoms[atm-1]
		}

		atoms[atm][hlf] &= INT60_FLAGS

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
		switch (frame.Body)[p] {
		case 32:
			goto tr109
		case 33:
			goto tr110
		case 35:
			goto tr111
		case 39:
			goto tr112
		case 42:
			goto tr111
		case 44:
			goto tr110
		case 46:
			goto tr113
		case 47:
			goto tr114
		case 58:
			goto tr111
		case 59:
			goto tr110
		case 61:
			goto tr116
		case 62:
			goto tr117
		case 63:
			goto tr110
		case 64:
			goto tr111
		case 91:
			goto tr103
		case 93:
			goto tr103
		case 94:
			goto tr118
		case 95:
			goto tr115
		case 123:
			goto tr103
		case 125:
			goto tr103
		case 126:
			goto tr115
		}
		switch {
		case (frame.Body)[p] < 43:
			switch {
			case (frame.Body)[p] < 36:
				if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
					goto tr109
				}
			case (frame.Body)[p] > 37:
				if 40 <= (frame.Body)[p] && (frame.Body)[p] <= 41 {
					goto tr103
				}
			default:
				goto tr101
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
			goto tr101
		}
		goto st0
	tr115:

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
	tr114:

		atoms[atm][hlf] <<= 6
		dgt--

		goto st28
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
		switch (frame.Body)[p] {
		case 32:
			goto tr109
		case 33:
			goto tr110
		case 35:
			goto tr111
		case 39:
			goto tr112
		case 42:
			goto tr111
		case 44:
			goto tr110
		case 46:
			goto tr113
		case 58:
			goto tr111
		case 59:
			goto tr110
		case 61:
			goto tr116
		case 62:
			goto tr117
		case 63:
			goto tr110
		case 64:
			goto tr111
		case 91:
			goto tr103
		case 93:
			goto tr103
		case 94:
			goto tr118
		case 95:
			goto tr115
		case 123:
			goto tr103
		case 125:
			goto tr103
		case 126:
			goto tr115
		}
		switch {
		case (frame.Body)[p] < 43:
			switch {
			case (frame.Body)[p] < 36:
				if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
					goto tr109
				}
			case (frame.Body)[p] > 37:
				if 40 <= (frame.Body)[p] && (frame.Body)[p] <= 41 {
					goto tr103
				}
			default:
				goto tr101
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
			goto tr101
		}
		goto st0
	tr10:

		atoms[atm].setType(ATOM_STRING)
		atoms[atm].setFrom(p)

		goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		switch (frame.Body)[p] {
		case 10:
			goto st0
		case 13:
			goto st0
		}
		goto st3
	tr52:

		dgt = 0
		hlf = 1

		dgt = int(ABC[(frame.Body)[p]]) + 4
		atoms[atm][hlf] &= INT60_FLAGS | PREFIX_MASKS[dgt]

		goto st29
	tr124:

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
	tr128:

		atoms[atm][hlf] <<= 6
		dgt--

		goto st29
	tr134:

		dgt = 0
		hlf = 1

		dgt = int(ABC[(frame.Body)[p]]) + 4
		atoms[atm][hlf] &= INT60_FLAGS | PREFIX_MASKS[dgt]

		goto st29
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
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
			goto tr125
		case 62:
			goto tr126
		case 63:
			goto tr120
		case 64:
			goto tr121
		case 94:
			goto tr127
		case 95:
			goto tr124
		case 126:
			goto tr124
		}
		switch {
		case (frame.Body)[p] < 48:
			if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
				goto tr119
			}
		case (frame.Body)[p] > 57:
			switch {
			case (frame.Body)[p] > 90:
				if 97 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr124
				}
			case (frame.Body)[p] >= 65:
				goto tr124
			}
		default:
			goto tr124
		}
		goto st0
	tr53:

		dgt = 0
		hlf = 1

		atoms[atm][hlf] &= INT60_FLAGS

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
			goto tr128
		case 58:
			goto tr121
		case 59:
			goto tr120
		case 61:
			goto tr125
		case 62:
			goto tr126
		case 63:
			goto tr120
		case 64:
			goto tr121
		case 94:
			goto tr127
		case 95:
			goto tr124
		case 126:
			goto tr124
		}
		switch {
		case (frame.Body)[p] < 48:
			if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
				goto tr119
			}
		case (frame.Body)[p] > 57:
			switch {
			case (frame.Body)[p] > 90:
				if 97 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr124
				}
			case (frame.Body)[p] >= 65:
				goto tr124
			}
		default:
			goto tr124
		}
		goto st0
	tr44:

		dgt = int(ABC[(frame.Body)[p]]) + 4
		atoms[atm][hlf] &= INT60_FLAGS | PREFIX_MASKS[dgt]

		goto st31
	tr136:

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
		switch (frame.Body)[p] {
		case 32:
			goto tr129
		case 33:
			goto tr130
		case 35:
			goto tr131
		case 39:
			goto tr133
		case 42:
			goto tr131
		case 44:
			goto tr130
		case 46:
			goto tr135
		case 58:
			goto tr131
		case 59:
			goto tr130
		case 61:
			goto tr137
		case 62:
			goto tr138
		case 63:
			goto tr130
		case 64:
			goto tr131
		case 91:
			goto tr134
		case 93:
			goto tr134
		case 94:
			goto tr139
		case 95:
			goto tr136
		case 123:
			goto tr134
		case 125:
			goto tr134
		case 126:
			goto tr136
		}
		switch {
		case (frame.Body)[p] < 43:
			switch {
			case (frame.Body)[p] < 36:
				if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
					goto tr129
				}
			case (frame.Body)[p] > 37:
				if 40 <= (frame.Body)[p] && (frame.Body)[p] <= 41 {
					goto tr134
				}
			default:
				goto tr132
			}
		case (frame.Body)[p] > 45:
			switch {
			case (frame.Body)[p] < 65:
				if 48 <= (frame.Body)[p] && (frame.Body)[p] <= 57 {
					goto tr136
				}
			case (frame.Body)[p] > 90:
				if 97 <= (frame.Body)[p] && (frame.Body)[p] <= 122 {
					goto tr136
				}
			default:
				goto tr136
			}
		default:
			goto tr132
		}
		goto st0
	tr46:

		atoms[atm][hlf] &= INT60_FLAGS

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
		switch (frame.Body)[p] {
		case 32:
			goto tr140
		case 33:
			goto tr141
		case 35:
			goto tr142
		case 39:
			goto tr143
		case 42:
			goto tr142
		case 44:
			goto tr141
		case 46:
			goto tr144
		case 47:
			goto tr145
		case 58:
			goto tr142
		case 59:
			goto tr141
		case 61:
			goto tr147
		case 62:
			goto tr148
		case 63:
			goto tr141
		case 64:
			goto tr142
		case 91:
			goto tr134
		case 93:
			goto tr134
		case 94:
			goto tr149
		case 95:
			goto tr146
		case 123:
			goto tr134
		case 125:
			goto tr134
		case 126:
			goto tr146
		}
		switch {
		case (frame.Body)[p] < 43:
			switch {
			case (frame.Body)[p] < 36:
				if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
					goto tr140
				}
			case (frame.Body)[p] > 37:
				if 40 <= (frame.Body)[p] && (frame.Body)[p] <= 41 {
					goto tr134
				}
			default:
				goto tr132
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
			goto tr132
		}
		goto st0
	tr146:

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
	tr145:

		atoms[atm][hlf] <<= 6
		dgt--

		goto st33
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
		switch (frame.Body)[p] {
		case 32:
			goto tr140
		case 33:
			goto tr141
		case 35:
			goto tr142
		case 39:
			goto tr143
		case 42:
			goto tr142
		case 44:
			goto tr141
		case 46:
			goto tr144
		case 58:
			goto tr142
		case 59:
			goto tr141
		case 61:
			goto tr147
		case 62:
			goto tr148
		case 63:
			goto tr141
		case 64:
			goto tr142
		case 91:
			goto tr134
		case 93:
			goto tr134
		case 94:
			goto tr149
		case 95:
			goto tr146
		case 123:
			goto tr134
		case 125:
			goto tr134
		case 126:
			goto tr146
		}
		switch {
		case (frame.Body)[p] < 43:
			switch {
			case (frame.Body)[p] < 36:
				if 9 <= (frame.Body)[p] && (frame.Body)[p] <= 13 {
					goto tr140
				}
			case (frame.Body)[p] > 37:
				if 40 <= (frame.Body)[p] && (frame.Body)[p] <= 41 {
					goto tr134
				}
			default:
				goto tr132
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
			goto tr132
		}
		goto st0
	tr50:

		if atm > 0 {
			atoms[atm] = atoms[atm-1]
		}

		goto st34
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
		switch (frame.Body)[p] {
		case 32:
			goto tr150
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
					goto tr150
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
	_test_eof13:
		(ps.state) = 13
		goto _test_eof
	_test_eof1:
		(ps.state) = 1
		goto _test_eof
	_test_eof14:
		(ps.state) = 14
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
	_test_eof2:
		(ps.state) = 2
		goto _test_eof
	_test_eof3:
		(ps.state) = 3
		goto _test_eof
	_test_eof18:
		(ps.state) = 18
		goto _test_eof
	_test_eof19:
		(ps.state) = 19
		goto _test_eof
	_test_eof20:
		(ps.state) = 20
		goto _test_eof
	_test_eof4:
		(ps.state) = 4
		goto _test_eof
	_test_eof5:
		(ps.state) = 5
		goto _test_eof
	_test_eof21:
		(ps.state) = 21
		goto _test_eof
	_test_eof6:
		(ps.state) = 6
		goto _test_eof
	_test_eof22:
		(ps.state) = 22
		goto _test_eof
	_test_eof23:
		(ps.state) = 23
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
	_test_eof24:
		(ps.state) = 24
		goto _test_eof
	_test_eof11:
		(ps.state) = 11
		goto _test_eof
	_test_eof25:
		(ps.state) = 25
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
	_test_eof12:
		(ps.state) = 12
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
	_test_eof33:
		(ps.state) = 33
		goto _test_eof
	_test_eof34:
		(ps.state) = 34
		goto _test_eof

	_test_eof:
		{
		}
		if p == eof {
			switch ps.state {
			case 14:

				frame.position++

			case 19:

				frame.position++

			case 15, 16, 17, 34:

				atm++

				frame.position++

			case 18:

				atm++

				frame.position++

			case 31:

				atm++

				frame.position++

			case 29, 30:

				atm++

				frame.position++

			case 21, 24:

				atoms[atm].setTill(p)
				atoms[atm].parseValue(frame.Body)

				atm++

				frame.position++

			case 22:

				atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

				atm++

				frame.position++

			case 32, 33:

				atoms[atm][1] = UUID_NAME_FLAG

				atm++

				frame.position++

			case 26:

				atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

				atm++

				frame.position++

			case 23, 25:

				atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

				atm++

				frame.position++

			case 27, 28:

				atoms[atm][1] = UUID_NAME_FLAG

				atoms[atm][1] |= ((uint64)(ATOM_UUID)) << 62

				atm++

				frame.position++

			}
		}

	_out:
		{
		}
	}

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

	const UUID_start int = 1
	const UUID_first_final int = 2
	const UUID_error int = 0

	const UUID_en_main int = 1

	cs, p, pe, eof := 0, 0, len(data), len(data)
	_ = eof

	atm, hlf, dgt := 0, 0, 0

	atoms := [1]Atom{Atom(ctx_uuid)}

	{
		cs = UUID_start
	}

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

		hlf = 1
		atoms[atm][1] &= INT60_FULL
		atoms[atm][1] |= ((uint64)(ABC[data[p]])) << 60

		goto st2
	tr8:

		hlf = 1
		atoms[atm][1] &= INT60_FULL
		atoms[atm][1] |= ((uint64)(ABC[data[p]])) << 60

		goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
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

		dgt = 0
		hlf = 1

		dgt = int(ABC[data[p]]) + 4
		atoms[atm][hlf] &= INT60_FLAGS | PREFIX_MASKS[dgt]

		goto st3
	tr6:

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

		atoms[atm][hlf] <<= 6
		dgt--

		goto st3
	tr9:

		dgt = 0
		hlf = 1

		dgt = int(ABC[data[p]]) + 4
		atoms[atm][hlf] &= INT60_FLAGS | PREFIX_MASKS[dgt]

		goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
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

		dgt = 0
		hlf = 1

		atoms[atm][hlf] &= INT60_FLAGS

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

		dgt = int(ABC[data[p]]) + 4
		atoms[atm][hlf] &= INT60_FLAGS | PREFIX_MASKS[dgt]

		goto st5
	tr10:

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

		atoms[atm][hlf] &= INT60_FLAGS

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

		atoms[atm][hlf] <<= 6
		dgt--

		goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
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

			case 3, 4:

			case 6, 7:

				atoms[atm][1] = UUID_NAME_FLAG

			}
		}

	_out:
		{
		}
	}

	if cs < UUID_first_final || dgt > 10 {
		return ERROR_UUID, errors.New(fmt.Sprintf("parse error at pos %d", p))
	} else {
		return UUID(atoms[0]), nil
	}

}
