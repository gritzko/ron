//line dfa.rl:1
package ron

import "fmt"
import "errors"

//line dfa.rl:7
//line dfa.go:12
const RON_start int = 21
const RON_first_final int = 21
const RON_error int = 0

const RON_en_main int = 21

//line dfa.rl:8

//line dfa.rl:9
//line dfa.rl:10
//line dfa.rl:11
//line dfa.rl:12
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

	if frame.Parser.state == RON_error {
		if frame.Parser.position == 0 {

//line dfa.go:46
			{
				(frame.Parser.state) = RON_start
			}

//line dfa.rl:29
			frame.atoms = make([]Atom, 4, 8)
		} else {
			return
		}
	} else if frame.Parser.state == RON_EOF {
		frame.Parser.state = RON_error
		return
	} else if frame.Parser.state == RON_start {
		frame.Parser.offset = frame.Parser.position
		frame.atoms = frame.atoms[:4]
		frame.Parser.atm, frame.Parser.hlf, frame.Parser.dgt = 0, 0, 0
	}

	pe, eof := len(frame.Body), len(frame.Body)
	n := 0
	_ = eof
	_ = pe // FIXME kill

	if frame.Parser.streaming {
		eof = -1
	}

	atm, hlf, dgt := frame.Parser.atm, frame.Parser.hlf, frame.Parser.dgt
	atoms := frame.atoms

//line dfa.go:78
	{
		if (frame.Parser.position) == pe {
			goto _test_eof
		}
		goto _resume

	_again:
		switch frame.Parser.state {
		case 21:
			goto st21
		case 0:
			goto st0
		case 1:
			goto st1
		case 2:
			goto st2
		case 22:
			goto st22
		case 23:
			goto st23
		case 3:
			goto st3
		case 4:
			goto st4
		case 5:
			goto st5
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
		case 26:
			goto st26
		case 9:
			goto st9
		case 27:
			goto st27
		case 28:
			goto st28
		case 10:
			goto st10
		case 11:
			goto st11
		case 12:
			goto st12
		case 13:
			goto st13
		case 29:
			goto st29
		case 14:
			goto st14
		case 15:
			goto st15
		case 30:
			goto st30
		case 31:
			goto st31
		case 32:
			goto st32
		case 16:
			goto st16
		case 17:
			goto st17
		case 18:
			goto st18
		case 19:
			goto st19
		case 20:
			goto st20
		}

		if (frame.Parser.position)++; (frame.Parser.position) == pe {
			goto _test_eof
		}
	_resume:
		switch frame.Parser.state {
		case 21:
			goto st_case_21
		case 0:
			goto st_case_0
		case 1:
			goto st_case_1
		case 2:
			goto st_case_2
		case 22:
			goto st_case_22
		case 23:
			goto st_case_23
		case 3:
			goto st_case_3
		case 4:
			goto st_case_4
		case 5:
			goto st_case_5
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
		case 26:
			goto st_case_26
		case 9:
			goto st_case_9
		case 27:
			goto st_case_27
		case 28:
			goto st_case_28
		case 10:
			goto st_case_10
		case 11:
			goto st_case_11
		case 12:
			goto st_case_12
		case 13:
			goto st_case_13
		case 29:
			goto st_case_29
		case 14:
			goto st_case_14
		case 15:
			goto st_case_15
		case 30:
			goto st_case_30
		case 31:
			goto st_case_31
		case 32:
			goto st_case_32
		case 16:
			goto st_case_16
		case 17:
			goto st_case_17
		case 18:
			goto st_case_18
		case 19:
			goto st_case_19
		case 20:
			goto st_case_20
		}
		goto st_out
	st21:
		if (frame.Parser.position)++; (frame.Parser.position) == pe {
			goto _test_eof21
		}
	st_case_21:
		switch (frame.Body)[(frame.Parser.position)] {
		case 32:
			goto st1
		case 35:
			goto tr2
		case 42:
			goto tr2
		case 46:
			goto tr71
		case 58:
			goto tr2
		case 64:
			goto tr2
		}
		if 9 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 13 {
			goto st1
		}
		goto st0
	st_case_0:
	st0:
		(frame.Parser.state) = 0
		goto _out
	st1:
		if (frame.Parser.position)++; (frame.Parser.position) == pe {
			goto _test_eof1
		}
	st_case_1:
		switch (frame.Body)[(frame.Parser.position)] {
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
		if 9 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 13 {
			goto st1
		}
		goto st0
	tr2:
		(frame.Parser.state) = 2
//line ./op-grammar.rl:91
		hlf = 0
		if frame.Parser.position > frame.Parser.offset {
			// one op is done, so stop parsing for now
			// make sure the parser restarts with the next op
			frame.Parser.position--
			(frame.Parser.state) = (RON_start)
			{
				(frame.Parser.position)++
				goto _out
			}
		} else {
			//op_idx++;
		}

//line ./op-grammar.rl:12
		n = (int)(ABC[(frame.Body)[(frame.Parser.position)]])
		if n < atm {
			// wrong UUID order; must be type-object-event-ref
			(frame.Parser.state) = (RON_error)
			{
				(frame.Parser.position)++
				goto _out
			}
		} else {
			// start parsing the UUID
			atm = n
			hlf = 0
			dgt = 0
		}

		goto _again
	tr3:
//line ./op-grammar.rl:26
		// OK, save the UUID
		atm++
		hlf = 0
		dgt = 0

		goto st2
	tr5:
		(frame.Parser.state) = 2
//line ./op-grammar.rl:26
		// OK, save the UUID
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:12
		n = (int)(ABC[(frame.Body)[(frame.Parser.position)]])
		if n < atm {
			// wrong UUID order; must be type-object-event-ref
			(frame.Parser.state) = (RON_error)
			{
				(frame.Parser.position)++
				goto _out
			}
		} else {
			// start parsing the UUID
			atm = n
			hlf = 0
			dgt = 0
		}

		goto _again
	tr19:
		(frame.Parser.state) = 2
//line ./op-grammar.rl:12
		n = (int)(ABC[(frame.Body)[(frame.Parser.position)]])
		if n < atm {
			// wrong UUID order; must be type-object-event-ref
			(frame.Parser.state) = (RON_error)
			{
				(frame.Parser.position)++
				goto _out
			}
		} else {
			// start parsing the UUID
			atm = n
			hlf = 0
			dgt = 0
		}

		goto _again
	tr47:
		(frame.Parser.state) = 2
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:26

		// OK, save the UUID
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:12
		n = (int)(ABC[(frame.Body)[(frame.Parser.position)]])
		if n < atm {
			// wrong UUID order; must be type-object-event-ref
			(frame.Parser.state) = (RON_error)
			{
				(frame.Parser.position)++
				goto _out
			}
		} else {
			// start parsing the UUID
			atm = n
			hlf = 0
			dgt = 0
		}

		goto _again
	tr55:
		(frame.Parser.state) = 2
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:26

		// OK, save the UUID
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:12
		n = (int)(ABC[(frame.Body)[(frame.Parser.position)]])
		if n < atm {
			// wrong UUID order; must be type-object-event-ref
			(frame.Parser.state) = (RON_error)
			{
				(frame.Parser.position)++
				goto _out
			}
		} else {
			// start parsing the UUID
			atm = n
			hlf = 0
			dgt = 0
		}

		goto _again
	tr65:
		(frame.Parser.state) = 2
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
		atoms[atm][1] = ((uint64)(UUID_NAME)) << 60

//line ./op-grammar.rl:26

		// OK, save the UUID
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:12
		n = (int)(ABC[(frame.Body)[(frame.Parser.position)]])
		if n < atm {
			// wrong UUID order; must be type-object-event-ref
			(frame.Parser.state) = (RON_error)
			{
				(frame.Parser.position)++
				goto _out
			}
		} else {
			// start parsing the UUID
			atm = n
			hlf = 0
			dgt = 0
		}

		goto _again
	tr73:
		(frame.Parser.state) = 2
//line ./op-grammar.rl:104

//line ./op-grammar.rl:91
		hlf = 0
		if frame.Parser.position > frame.Parser.offset {
			// one op is done, so stop parsing for now
			// make sure the parser restarts with the next op
			frame.Parser.position--
			(frame.Parser.state) = (RON_start)
			{
				(frame.Parser.position)++
				goto _out
			}
		} else {
			//op_idx++;
		}

//line ./op-grammar.rl:12
		n = (int)(ABC[(frame.Body)[(frame.Parser.position)]])
		if n < atm {
			// wrong UUID order; must be type-object-event-ref
			(frame.Parser.state) = (RON_error)
			{
				(frame.Parser.position)++
				goto _out
			}
		} else {
			// start parsing the UUID
			atm = n
			hlf = 0
			dgt = 0
		}

		goto _again
	tr77:
		(frame.Parser.state) = 2
//line ./op-grammar.rl:37

		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:84

//line ./op-grammar.rl:104

//line ./op-grammar.rl:91
		hlf = 0
		if frame.Parser.position > frame.Parser.offset {
			// one op is done, so stop parsing for now
			// make sure the parser restarts with the next op
			frame.Parser.position--
			(frame.Parser.state) = (RON_start)
			{
				(frame.Parser.position)++
				goto _out
			}
		} else {
			//op_idx++;
		}

//line ./op-grammar.rl:12
		n = (int)(ABC[(frame.Body)[(frame.Parser.position)]])
		if n < atm {
			// wrong UUID order; must be type-object-event-ref
			(frame.Parser.state) = (RON_error)
			{
				(frame.Parser.position)++
				goto _out
			}
		} else {
			// start parsing the UUID
			atm = n
			hlf = 0
			dgt = 0
		}

		goto _again
	tr85:
		(frame.Parser.state) = 2
//line ./op-grammar.rl:84

//line ./op-grammar.rl:104

//line ./op-grammar.rl:91
		hlf = 0
		if frame.Parser.position > frame.Parser.offset {
			// one op is done, so stop parsing for now
			// make sure the parser restarts with the next op
			frame.Parser.position--
			(frame.Parser.state) = (RON_start)
			{
				(frame.Parser.position)++
				goto _out
			}
		} else {
			//op_idx++;
		}

//line ./op-grammar.rl:12
		n = (int)(ABC[(frame.Body)[(frame.Parser.position)]])
		if n < atm {
			// wrong UUID order; must be type-object-event-ref
			(frame.Parser.state) = (RON_error)
			{
				(frame.Parser.position)++
				goto _out
			}
		} else {
			// start parsing the UUID
			atm = n
			hlf = 0
			dgt = 0
		}

		goto _again
	tr93:
		(frame.Parser.state) = 2
//line ./op-grammar.rl:56

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:84

//line ./op-grammar.rl:104

//line ./op-grammar.rl:91
		hlf = 0
		if frame.Parser.position > frame.Parser.offset {
			// one op is done, so stop parsing for now
			// make sure the parser restarts with the next op
			frame.Parser.position--
			(frame.Parser.state) = (RON_start)
			{
				(frame.Parser.position)++
				goto _out
			}
		} else {
			//op_idx++;
		}

//line ./op-grammar.rl:12
		n = (int)(ABC[(frame.Body)[(frame.Parser.position)]])
		if n < atm {
			// wrong UUID order; must be type-object-event-ref
			(frame.Parser.state) = (RON_error)
			{
				(frame.Parser.position)++
				goto _out
			}
		} else {
			// start parsing the UUID
			atm = n
			hlf = 0
			dgt = 0
		}

		goto _again
	tr103:
		(frame.Parser.state) = 2
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:37

		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:84

//line ./op-grammar.rl:104

//line ./op-grammar.rl:91
		hlf = 0
		if frame.Parser.position > frame.Parser.offset {
			// one op is done, so stop parsing for now
			// make sure the parser restarts with the next op
			frame.Parser.position--
			(frame.Parser.state) = (RON_start)
			{
				(frame.Parser.position)++
				goto _out
			}
		} else {
			//op_idx++;
		}

//line ./op-grammar.rl:12
		n = (int)(ABC[(frame.Body)[(frame.Parser.position)]])
		if n < atm {
			// wrong UUID order; must be type-object-event-ref
			(frame.Parser.state) = (RON_error)
			{
				(frame.Parser.position)++
				goto _out
			}
		} else {
			// start parsing the UUID
			atm = n
			hlf = 0
			dgt = 0
		}

		goto _again
	tr112:
		(frame.Parser.state) = 2
//line ./op-grammar.rl:59
		atoms[atm][1] |= ((uint64)(ATOM_FLOAT)) << 62

//line ./op-grammar.rl:62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:84

//line ./op-grammar.rl:104

//line ./op-grammar.rl:91
		hlf = 0
		if frame.Parser.position > frame.Parser.offset {
			// one op is done, so stop parsing for now
			// make sure the parser restarts with the next op
			frame.Parser.position--
			(frame.Parser.state) = (RON_start)
			{
				(frame.Parser.position)++
				goto _out
			}
		} else {
			//op_idx++;
		}

//line ./op-grammar.rl:12
		n = (int)(ABC[(frame.Body)[(frame.Parser.position)]])
		if n < atm {
			// wrong UUID order; must be type-object-event-ref
			(frame.Parser.state) = (RON_error)
			{
				(frame.Parser.position)++
				goto _out
			}
		} else {
			// start parsing the UUID
			atm = n
			hlf = 0
			dgt = 0
		}

		goto _again
	tr121:
		(frame.Parser.state) = 2
//line ./op-grammar.rl:62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:84

//line ./op-grammar.rl:104

//line ./op-grammar.rl:91
		hlf = 0
		if frame.Parser.position > frame.Parser.offset {
			// one op is done, so stop parsing for now
			// make sure the parser restarts with the next op
			frame.Parser.position--
			(frame.Parser.state) = (RON_start)
			{
				(frame.Parser.position)++
				goto _out
			}
		} else {
			//op_idx++;
		}

//line ./op-grammar.rl:12
		n = (int)(ABC[(frame.Body)[(frame.Parser.position)]])
		if n < atm {
			// wrong UUID order; must be type-object-event-ref
			(frame.Parser.state) = (RON_error)
			{
				(frame.Parser.position)++
				goto _out
			}
		} else {
			// start parsing the UUID
			atm = n
			hlf = 0
			dgt = 0
		}

		goto _again
	tr129:
		(frame.Parser.state) = 2
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:37

		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:84

//line ./op-grammar.rl:104

//line ./op-grammar.rl:91
		hlf = 0
		if frame.Parser.position > frame.Parser.offset {
			// one op is done, so stop parsing for now
			// make sure the parser restarts with the next op
			frame.Parser.position--
			(frame.Parser.state) = (RON_start)
			{
				(frame.Parser.position)++
				goto _out
			}
		} else {
			//op_idx++;
		}

//line ./op-grammar.rl:12
		n = (int)(ABC[(frame.Body)[(frame.Parser.position)]])
		if n < atm {
			// wrong UUID order; must be type-object-event-ref
			(frame.Parser.state) = (RON_error)
			{
				(frame.Parser.position)++
				goto _out
			}
		} else {
			// start parsing the UUID
			atm = n
			hlf = 0
			dgt = 0
		}

		goto _again
	tr140:
		(frame.Parser.state) = 2
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
		atoms[atm][1] = ((uint64)(UUID_NAME)) << 60

//line ./op-grammar.rl:37

		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:84

//line ./op-grammar.rl:104

//line ./op-grammar.rl:91
		hlf = 0
		if frame.Parser.position > frame.Parser.offset {
			// one op is done, so stop parsing for now
			// make sure the parser restarts with the next op
			frame.Parser.position--
			(frame.Parser.state) = (RON_start)
			{
				(frame.Parser.position)++
				goto _out
			}
		} else {
			//op_idx++;
		}

//line ./op-grammar.rl:12
		n = (int)(ABC[(frame.Body)[(frame.Parser.position)]])
		if n < atm {
			// wrong UUID order; must be type-object-event-ref
			(frame.Parser.state) = (RON_error)
			{
				(frame.Parser.position)++
				goto _out
			}
		} else {
			// start parsing the UUID
			atm = n
			hlf = 0
			dgt = 0
		}

		goto _again
	st2:
		if (frame.Parser.position)++; (frame.Parser.position) == pe {
			goto _test_eof2
		}
	st_case_2:
//line dfa.go:842
		switch (frame.Body)[(frame.Parser.position)] {
		case 32:
			goto tr3
		case 33:
			goto tr4
		case 35:
			goto tr5
		case 39:
			goto tr7
		case 42:
			goto tr5
		case 44:
			goto tr4
		case 58:
			goto tr5
		case 59:
			goto tr4
		case 61:
			goto tr10
		case 62:
			goto tr11
		case 63:
			goto tr4
		case 64:
			goto tr5
		case 91:
			goto tr8
		case 93:
			goto tr8
		case 94:
			goto tr12
		case 96:
			goto tr13
		case 123:
			goto tr8
		case 125:
			goto tr8
		case 126:
			goto tr9
		}
		switch {
		case (frame.Body)[(frame.Parser.position)] < 43:
			switch {
			case (frame.Body)[(frame.Parser.position)] < 36:
				if 9 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 13 {
					goto tr3
				}
			case (frame.Body)[(frame.Parser.position)] > 37:
				if 40 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 41 {
					goto tr8
				}
			default:
				goto tr6
			}
		case (frame.Body)[(frame.Parser.position)] > 45:
			switch {
			case (frame.Body)[(frame.Parser.position)] < 65:
				if 48 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 57 {
					goto tr9
				}
			case (frame.Body)[(frame.Parser.position)] > 90:
				if 95 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 122 {
					goto tr9
				}
			default:
				goto tr9
			}
		default:
			goto tr6
		}
		goto st0
	tr4:
//line ./op-grammar.rl:26

		// OK, save the UUID
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:107
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:87
		frame.term = int(ABC[(frame.Body)[(frame.Parser.position)]])

		goto st22
	tr18:
//line ./op-grammar.rl:107
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:87
		frame.term = int(ABC[(frame.Body)[(frame.Parser.position)]])

		goto st22
	tr46:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:26

		// OK, save the UUID
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:107
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:87
		frame.term = int(ABC[(frame.Body)[(frame.Parser.position)]])

		goto st22
	tr54:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:26

		// OK, save the UUID
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:107
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:87
		frame.term = int(ABC[(frame.Body)[(frame.Parser.position)]])

		goto st22
	tr64:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
		atoms[atm][1] = ((uint64)(UUID_NAME)) << 60

//line ./op-grammar.rl:26

		// OK, save the UUID
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:107
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:87
		frame.term = int(ABC[(frame.Body)[(frame.Parser.position)]])

		goto st22
	tr72:
//line ./op-grammar.rl:104

		goto st22
	tr76:
//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:84

//line ./op-grammar.rl:87
		frame.term = int(ABC[(frame.Body)[(frame.Parser.position)]])

		goto st22
	tr84:
//line ./op-grammar.rl:84
//line ./op-grammar.rl:87
		frame.term = int(ABC[(frame.Body)[(frame.Parser.position)]])

		goto st22
	tr92:
//line ./op-grammar.rl:56
//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:84

//line ./op-grammar.rl:87
		frame.term = int(ABC[(frame.Body)[(frame.Parser.position)]])

		goto st22
	tr102:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:37

		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:84

//line ./op-grammar.rl:87
		frame.term = int(ABC[(frame.Body)[(frame.Parser.position)]])

		goto st22
	tr111:
//line ./op-grammar.rl:59
		atoms[atm][1] |= ((uint64)(ATOM_FLOAT)) << 62

//line ./op-grammar.rl:62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:84

//line ./op-grammar.rl:87
		frame.term = int(ABC[(frame.Body)[(frame.Parser.position)]])

		goto st22
	tr120:
//line ./op-grammar.rl:62
//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:84

//line ./op-grammar.rl:87
		frame.term = int(ABC[(frame.Body)[(frame.Parser.position)]])

		goto st22
	tr128:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:37

		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:84

//line ./op-grammar.rl:87
		frame.term = int(ABC[(frame.Body)[(frame.Parser.position)]])

		goto st22
	tr139:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
		atoms[atm][1] = ((uint64)(UUID_NAME)) << 60

//line ./op-grammar.rl:37

		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:84

//line ./op-grammar.rl:87
		frame.term = int(ABC[(frame.Body)[(frame.Parser.position)]])

		goto st22
	st22:
		if (frame.Parser.position)++; (frame.Parser.position) == pe {
			goto _test_eof22
		}
	st_case_22:
//line dfa.go:1172
		switch (frame.Body)[(frame.Parser.position)] {
		case 32:
			goto tr72
		case 35:
			goto tr73
		case 42:
			goto tr73
		case 46:
			goto tr74
		case 58:
			goto tr73
		case 64:
			goto tr73
		}
		if 9 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 13 {
			goto tr72
		}
		goto st0
	tr71:
		(frame.Parser.state) = 23
//line ./op-grammar.rl:113
		(frame.Parser.state) = (RON_EOF)

		goto _again
	tr74:
		(frame.Parser.state) = 23
//line ./op-grammar.rl:104
//line ./op-grammar.rl:113
		(frame.Parser.state) = (RON_EOF)

		goto _again
	tr79:
		(frame.Parser.state) = 23
//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:84

//line ./op-grammar.rl:104

//line ./op-grammar.rl:113
		(frame.Parser.state) = (RON_EOF)

		goto _again
	tr87:
		(frame.Parser.state) = 23
//line ./op-grammar.rl:84
//line ./op-grammar.rl:104

//line ./op-grammar.rl:113
		(frame.Parser.state) = (RON_EOF)

		goto _again
	tr95:
		(frame.Parser.state) = 23
//line ./op-grammar.rl:56
//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:84

//line ./op-grammar.rl:104

//line ./op-grammar.rl:113
		(frame.Parser.state) = (RON_EOF)

		goto _again
	tr105:
		(frame.Parser.state) = 23
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:37

		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:84

//line ./op-grammar.rl:104

//line ./op-grammar.rl:113
		(frame.Parser.state) = (RON_EOF)

		goto _again
	tr114:
		(frame.Parser.state) = 23
//line ./op-grammar.rl:59
		atoms[atm][1] |= ((uint64)(ATOM_FLOAT)) << 62

//line ./op-grammar.rl:62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:84

//line ./op-grammar.rl:104

//line ./op-grammar.rl:113
		(frame.Parser.state) = (RON_EOF)

		goto _again
	tr123:
		(frame.Parser.state) = 23
//line ./op-grammar.rl:62
//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:84

//line ./op-grammar.rl:104

//line ./op-grammar.rl:113
		(frame.Parser.state) = (RON_EOF)

		goto _again
	tr133:
		(frame.Parser.state) = 23
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:37

		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:84

//line ./op-grammar.rl:104

//line ./op-grammar.rl:113
		(frame.Parser.state) = (RON_EOF)

		goto _again
	tr142:
		(frame.Parser.state) = 23
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
		atoms[atm][1] = ((uint64)(UUID_NAME)) << 60

//line ./op-grammar.rl:37

		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:84

//line ./op-grammar.rl:104

//line ./op-grammar.rl:113
		(frame.Parser.state) = (RON_EOF)

		goto _again
	st23:
		if (frame.Parser.position)++; (frame.Parser.position) == pe {
			goto _test_eof23
		}
	st_case_23:
//line dfa.go:1392
		goto st0
	tr6:
//line ././uuid-grammar.rl:5

//line ././uuid-grammar.rl:40
		hlf = 1
		atoms[atm][hlf] &= INT60_FULL
		atoms[atm][hlf] |= ((uint64)(ABC[(frame.Body)[(frame.Parser.position)]])) << 60

		goto st3
	tr56:
//line ././uuid-grammar.rl:34
//line ././uuid-grammar.rl:40
		hlf = 1
		atoms[atm][hlf] &= INT60_FULL
		atoms[atm][hlf] |= ((uint64)(ABC[(frame.Body)[(frame.Parser.position)]])) << 60

		goto st3
	st3:
		if (frame.Parser.position)++; (frame.Parser.position) == pe {
			goto _test_eof3
		}
	st_case_3:
//line dfa.go:1421
		switch (frame.Body)[(frame.Parser.position)] {
		case 32:
			goto tr14
		case 33:
			goto tr4
		case 35:
			goto tr5
		case 39:
			goto tr7
		case 42:
			goto tr5
		case 44:
			goto tr4
		case 58:
			goto tr5
		case 59:
			goto tr4
		case 61:
			goto tr10
		case 62:
			goto tr11
		case 63:
			goto tr4
		case 64:
			goto tr5
		case 91:
			goto tr15
		case 93:
			goto tr15
		case 94:
			goto tr12
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
		case (frame.Body)[(frame.Parser.position)] < 48:
			switch {
			case (frame.Body)[(frame.Parser.position)] > 13:
				if 40 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 41 {
					goto tr15
				}
			case (frame.Body)[(frame.Parser.position)] >= 9:
				goto tr14
			}
		case (frame.Body)[(frame.Parser.position)] > 57:
			switch {
			case (frame.Body)[(frame.Parser.position)] > 90:
				if 97 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 122 {
					goto tr16
				}
			case (frame.Body)[(frame.Parser.position)] >= 65:
				goto tr16
			}
		default:
			goto tr16
		}
		goto st0
	tr14:
//line ./op-grammar.rl:26

		// OK, save the UUID
		atm++
		hlf = 0
		dgt = 0

		goto st4
	tr45:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:26

		// OK, save the UUID
		atm++
		hlf = 0
		dgt = 0

		goto st4
	tr53:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:26

		// OK, save the UUID
		atm++
		hlf = 0
		dgt = 0

		goto st4
	tr63:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
		atoms[atm][1] = ((uint64)(UUID_NAME)) << 60

//line ./op-grammar.rl:26

		// OK, save the UUID
		atm++
		hlf = 0
		dgt = 0

		goto st4
	st4:
		if (frame.Parser.position)++; (frame.Parser.position) == pe {
			goto _test_eof4
		}
	st_case_4:
//line dfa.go:1539
		switch (frame.Body)[(frame.Parser.position)] {
		case 32:
			goto st4
		case 33:
			goto tr18
		case 35:
			goto tr19
		case 39:
			goto tr20
		case 42:
			goto tr19
		case 44:
			goto tr18
		case 58:
			goto tr19
		case 59:
			goto tr18
		case 61:
			goto tr21
		case 62:
			goto tr22
		case 63:
			goto tr18
		case 64:
			goto tr19
		case 94:
			goto tr23
		}
		if 9 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 13 {
			goto st4
		}
		goto st0
	tr7:
//line ./op-grammar.rl:26

		// OK, save the UUID
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:107
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:79
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st5
	tr20:
//line ./op-grammar.rl:107
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:79
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st5
	tr48:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:26

		// OK, save the UUID
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:107
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:79
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st5
	tr57:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:26

		// OK, save the UUID
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:107
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:79
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st5
	tr66:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
		atoms[atm][1] = ((uint64)(UUID_NAME)) << 60

//line ./op-grammar.rl:26

		// OK, save the UUID
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:107
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:79
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st5
	tr78:
//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st5
	tr86:
//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st5
	tr94:
//line ./op-grammar.rl:56

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st5
	tr104:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:37

		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st5
	tr113:
//line ./op-grammar.rl:59
		atoms[atm][1] |= ((uint64)(ATOM_FLOAT)) << 62

//line ./op-grammar.rl:62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st5
	tr122:
//line ./op-grammar.rl:62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st5
	tr131:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:37

		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st5
	tr141:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
		atoms[atm][1] = ((uint64)(UUID_NAME)) << 60

//line ./op-grammar.rl:37

		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st5
	st5:
		if (frame.Parser.position)++; (frame.Parser.position) == pe {
			goto _test_eof5
		}
	st_case_5:
//line dfa.go:1844
		switch (frame.Body)[(frame.Parser.position)] {
		case 10:
			goto st0
		case 13:
			goto st0
		case 39:
			goto tr25
		case 92:
			goto tr26
		}
		goto tr24
	tr24:
//line ./op-grammar.rl:65
		atoms[atm][0] = ((uint64)(frame.Parser.position)) << 32

		goto st6
	st6:
		if (frame.Parser.position)++; (frame.Parser.position) == pe {
			goto _test_eof6
		}
	st_case_6:
//line dfa.go:1867
		switch (frame.Body)[(frame.Parser.position)] {
		case 10:
			goto st0
		case 13:
			goto st0
		case 39:
			goto tr28
		case 92:
			goto st16
		}
		goto st6
	tr25:
//line ./op-grammar.rl:65
		atoms[atm][0] = ((uint64)(frame.Parser.position)) << 32

//line ./op-grammar.rl:68
		atoms[atm][0] |= uint64(frame.Parser.position)
		atoms[atm][1] = ((uint64)(ATOM_STRING)) << 62

		goto st24
	tr28:
//line ./op-grammar.rl:68
		atoms[atm][0] |= uint64(frame.Parser.position)
		atoms[atm][1] = ((uint64)(ATOM_STRING)) << 62

		goto st24
	st24:
		if (frame.Parser.position)++; (frame.Parser.position) == pe {
			goto _test_eof24
		}
	st_case_24:
//line dfa.go:1902
		switch (frame.Body)[(frame.Parser.position)] {
		case 32:
			goto tr75
		case 33:
			goto tr76
		case 35:
			goto tr77
		case 39:
			goto tr78
		case 42:
			goto tr77
		case 44:
			goto tr76
		case 46:
			goto tr79
		case 58:
			goto tr77
		case 59:
			goto tr76
		case 61:
			goto tr80
		case 62:
			goto tr81
		case 63:
			goto tr76
		case 64:
			goto tr77
		case 94:
			goto tr82
		}
		if 9 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 13 {
			goto tr75
		}
		goto st0
	tr75:
//line ./op-grammar.rl:37

		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:84

//line ./op-grammar.rl:104

		goto st25
	tr83:
//line ./op-grammar.rl:84
//line ./op-grammar.rl:104

		goto st25
	tr91:
//line ./op-grammar.rl:56
//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:84

//line ./op-grammar.rl:104

		goto st25
	tr101:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:37

		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:84

//line ./op-grammar.rl:104

		goto st25
	tr110:
//line ./op-grammar.rl:59
		atoms[atm][1] |= ((uint64)(ATOM_FLOAT)) << 62

//line ./op-grammar.rl:62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:84

//line ./op-grammar.rl:104

		goto st25
	tr119:
//line ./op-grammar.rl:62
//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:84

//line ./op-grammar.rl:104

		goto st25
	tr127:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:37

		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:84

//line ./op-grammar.rl:104

		goto st25
	tr138:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
		atoms[atm][1] = ((uint64)(UUID_NAME)) << 60

//line ./op-grammar.rl:37

		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:84

//line ./op-grammar.rl:104

		goto st25
	st25:
		if (frame.Parser.position)++; (frame.Parser.position) == pe {
			goto _test_eof25
		}
	st_case_25:
//line dfa.go:2081
		switch (frame.Body)[(frame.Parser.position)] {
		case 32:
			goto tr83
		case 33:
			goto tr84
		case 35:
			goto tr85
		case 39:
			goto tr86
		case 42:
			goto tr85
		case 44:
			goto tr84
		case 46:
			goto tr87
		case 58:
			goto tr85
		case 59:
			goto tr84
		case 61:
			goto tr88
		case 62:
			goto tr89
		case 63:
			goto tr84
		case 64:
			goto tr85
		case 94:
			goto tr90
		}
		if 9 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 13 {
			goto tr83
		}
		goto st0
	tr10:
//line ./op-grammar.rl:26

		// OK, save the UUID
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:107
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:79
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st7
	tr21:
//line ./op-grammar.rl:107
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:79
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st7
	tr50:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:26

		// OK, save the UUID
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:107
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:79
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st7
	tr60:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:26

		// OK, save the UUID
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:107
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:79
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st7
	tr68:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
		atoms[atm][1] = ((uint64)(UUID_NAME)) << 60

//line ./op-grammar.rl:26

		// OK, save the UUID
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:107
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:79
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st7
	tr80:
//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st7
	tr88:
//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st7
	tr96:
//line ./op-grammar.rl:56

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st7
	tr107:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:37

		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st7
	tr115:
//line ./op-grammar.rl:59
		atoms[atm][1] |= ((uint64)(ATOM_FLOAT)) << 62

//line ./op-grammar.rl:62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st7
	tr124:
//line ./op-grammar.rl:62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st7
	tr135:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:37

		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st7
	tr144:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
		atoms[atm][1] = ((uint64)(UUID_NAME)) << 60

//line ./op-grammar.rl:37

		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st7
	st7:
		if (frame.Parser.position)++; (frame.Parser.position) == pe {
			goto _test_eof7
		}
	st_case_7:
//line dfa.go:2388
		switch (frame.Body)[(frame.Parser.position)] {
		case 32:
			goto st7
		case 43:
			goto tr31
		case 45:
			goto tr31
		}
		switch {
		case (frame.Body)[(frame.Parser.position)] > 13:
			if 48 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 57 {
				goto tr32
			}
		case (frame.Body)[(frame.Parser.position)] >= 9:
			goto st7
		}
		goto st0
	tr31:
//line ./op-grammar.rl:44
		atoms[atm][1] |= ((uint64)(ATOM_INT)) << 62

//line ./op-grammar.rl:47
		if (frame.Body)[(frame.Parser.position)] == '-' {
			atoms[atm][1] |= 1
		}

		goto st8
	st8:
		if (frame.Parser.position)++; (frame.Parser.position) == pe {
			goto _test_eof8
		}
	st_case_8:
//line dfa.go:2423
		if 48 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 57 {
			goto tr33
		}
		goto st0
	tr32:
//line ./op-grammar.rl:44
		atoms[atm][1] |= ((uint64)(ATOM_INT)) << 62

//line ./op-grammar.rl:52
		atoms[atm][0] *= 10
		atoms[atm][0] += (uint64)((frame.Body)[(frame.Parser.position)] - '0')

		goto st26
	tr33:
//line ./op-grammar.rl:52
		atoms[atm][0] *= 10
		atoms[atm][0] += (uint64)((frame.Body)[(frame.Parser.position)] - '0')

		goto st26
	st26:
		if (frame.Parser.position)++; (frame.Parser.position) == pe {
			goto _test_eof26
		}
	st_case_26:
//line dfa.go:2451
		switch (frame.Body)[(frame.Parser.position)] {
		case 32:
			goto tr91
		case 33:
			goto tr92
		case 35:
			goto tr93
		case 39:
			goto tr94
		case 42:
			goto tr93
		case 44:
			goto tr92
		case 46:
			goto tr95
		case 58:
			goto tr93
		case 59:
			goto tr92
		case 61:
			goto tr96
		case 62:
			goto tr97
		case 63:
			goto tr92
		case 64:
			goto tr93
		case 94:
			goto tr98
		}
		switch {
		case (frame.Body)[(frame.Parser.position)] > 13:
			if 48 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 57 {
				goto tr33
			}
		case (frame.Body)[(frame.Parser.position)] >= 9:
			goto tr91
		}
		goto st0
	tr11:
//line ./op-grammar.rl:26

		// OK, save the UUID
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:107
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:79
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st9
	tr22:
//line ./op-grammar.rl:107
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:79
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st9
	tr51:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:26

		// OK, save the UUID
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:107
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:79
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st9
	tr61:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:26

		// OK, save the UUID
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:107
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:79
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st9
	tr69:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
		atoms[atm][1] = ((uint64)(UUID_NAME)) << 60

//line ./op-grammar.rl:26

		// OK, save the UUID
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:107
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:79
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st9
	tr81:
//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st9
	tr89:
//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st9
	tr97:
//line ./op-grammar.rl:56

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st9
	tr108:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:37

		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st9
	tr116:
//line ./op-grammar.rl:59
		atoms[atm][1] |= ((uint64)(ATOM_FLOAT)) << 62

//line ./op-grammar.rl:62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st9
	tr125:
//line ./op-grammar.rl:62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st9
	tr136:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:37

		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st9
	tr145:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
		atoms[atm][1] = ((uint64)(UUID_NAME)) << 60

//line ./op-grammar.rl:37

		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st9
	st9:
		if (frame.Parser.position)++; (frame.Parser.position) == pe {
			goto _test_eof9
		}
	st_case_9:
//line dfa.go:2763
		switch (frame.Body)[(frame.Parser.position)] {
		case 32:
			goto st9
		case 43:
			goto tr35
		case 45:
			goto tr35
		case 91:
			goto tr36
		case 93:
			goto tr36
		case 95:
			goto tr37
		case 123:
			goto tr36
		case 125:
			goto tr36
		case 126:
			goto tr37
		}
		switch {
		case (frame.Body)[(frame.Parser.position)] < 40:
			switch {
			case (frame.Body)[(frame.Parser.position)] > 13:
				if 36 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 37 {
					goto tr35
				}
			case (frame.Body)[(frame.Parser.position)] >= 9:
				goto st9
			}
		case (frame.Body)[(frame.Parser.position)] > 41:
			switch {
			case (frame.Body)[(frame.Parser.position)] < 65:
				if 48 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 57 {
					goto tr37
				}
			case (frame.Body)[(frame.Parser.position)] > 90:
				if 97 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 122 {
					goto tr37
				}
			default:
				goto tr37
			}
		default:
			goto tr36
		}
		goto st0
	tr35:
//line ././uuid-grammar.rl:5

//line ././uuid-grammar.rl:40
		hlf = 1
		atoms[atm][hlf] &= INT60_FULL
		atoms[atm][hlf] |= ((uint64)(ABC[(frame.Body)[(frame.Parser.position)]])) << 60

		goto st27
	tr130:
//line ././uuid-grammar.rl:34
//line ././uuid-grammar.rl:40
		hlf = 1
		atoms[atm][hlf] &= INT60_FULL
		atoms[atm][hlf] |= ((uint64)(ABC[(frame.Body)[(frame.Parser.position)]])) << 60

		goto st27
	st27:
		if (frame.Parser.position)++; (frame.Parser.position) == pe {
			goto _test_eof27
		}
	st_case_27:
//line dfa.go:2838
		switch (frame.Body)[(frame.Parser.position)] {
		case 32:
			goto tr75
		case 33:
			goto tr76
		case 35:
			goto tr77
		case 39:
			goto tr78
		case 42:
			goto tr77
		case 44:
			goto tr76
		case 46:
			goto tr79
		case 58:
			goto tr77
		case 59:
			goto tr76
		case 61:
			goto tr80
		case 62:
			goto tr81
		case 63:
			goto tr76
		case 64:
			goto tr77
		case 91:
			goto tr99
		case 93:
			goto tr99
		case 94:
			goto tr82
		case 95:
			goto tr100
		case 123:
			goto tr99
		case 125:
			goto tr99
		case 126:
			goto tr100
		}
		switch {
		case (frame.Body)[(frame.Parser.position)] < 48:
			switch {
			case (frame.Body)[(frame.Parser.position)] > 13:
				if 40 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 41 {
					goto tr99
				}
			case (frame.Body)[(frame.Parser.position)] >= 9:
				goto tr75
			}
		case (frame.Body)[(frame.Parser.position)] > 57:
			switch {
			case (frame.Body)[(frame.Parser.position)] > 90:
				if 97 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 122 {
					goto tr100
				}
			case (frame.Body)[(frame.Parser.position)] >= 65:
				goto tr100
			}
		default:
			goto tr100
		}
		goto st0
	tr99:
//line ././uuid-grammar.rl:29
		dgt = 0
		hlf = 1

//line ././uuid-grammar.rl:9
		dgt = int(ABC[(frame.Body)[(frame.Parser.position)]]) + 4
		atoms[atm][hlf] &= INT60_FLAGS | PREFIX_MASKS[dgt]

		goto st28
	tr100:
//line ././uuid-grammar.rl:29
		dgt = 0
		hlf = 1

//line ././uuid-grammar.rl:22
		atoms[atm][hlf] &= INT60_FLAGS

//line ././uuid-grammar.rl:14
		atoms[atm][hlf] |= ((uint64)(ABC[(frame.Body)[(frame.Parser.position)]])) << DIGIT_OFFSETS[dgt]
		dgt++
		if dgt > 10 {
			{
				(frame.Parser.position)++
				(frame.Parser.state) = 28
				goto _out
			}
		}

		goto st28
	tr106:
//line ././uuid-grammar.rl:14
		atoms[atm][hlf] |= ((uint64)(ABC[(frame.Body)[(frame.Parser.position)]])) << DIGIT_OFFSETS[dgt]
		dgt++
		if dgt > 10 {
			{
				(frame.Parser.position)++
				(frame.Parser.state) = 28
				goto _out
			}
		}

		goto st28
	tr132:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:29
		dgt = 0
		hlf = 1

//line ././uuid-grammar.rl:9
		dgt = int(ABC[(frame.Body)[(frame.Parser.position)]]) + 4
		atoms[atm][hlf] &= INT60_FLAGS | PREFIX_MASKS[dgt]

		goto st28
	st28:
		if (frame.Parser.position)++; (frame.Parser.position) == pe {
			goto _test_eof28
		}
	st_case_28:
//line dfa.go:2965
		switch (frame.Body)[(frame.Parser.position)] {
		case 32:
			goto tr101
		case 33:
			goto tr102
		case 35:
			goto tr103
		case 39:
			goto tr104
		case 42:
			goto tr103
		case 44:
			goto tr102
		case 46:
			goto tr105
		case 58:
			goto tr103
		case 59:
			goto tr102
		case 61:
			goto tr107
		case 62:
			goto tr108
		case 63:
			goto tr102
		case 64:
			goto tr103
		case 94:
			goto tr109
		case 95:
			goto tr106
		case 126:
			goto tr106
		}
		switch {
		case (frame.Body)[(frame.Parser.position)] < 48:
			if 9 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 13 {
				goto tr101
			}
		case (frame.Body)[(frame.Parser.position)] > 57:
			switch {
			case (frame.Body)[(frame.Parser.position)] > 90:
				if 97 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 122 {
					goto tr106
				}
			case (frame.Body)[(frame.Parser.position)] >= 65:
				goto tr106
			}
		default:
			goto tr106
		}
		goto st0
	tr12:
//line ./op-grammar.rl:26

		// OK, save the UUID
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:107
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:79
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st10
	tr23:
//line ./op-grammar.rl:107
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:79
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st10
	tr52:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:26

		// OK, save the UUID
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:107
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:79
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st10
	tr62:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:26

		// OK, save the UUID
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:107
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:79
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st10
	tr70:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
		atoms[atm][1] = ((uint64)(UUID_NAME)) << 60

//line ./op-grammar.rl:26

		// OK, save the UUID
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:107
		if frame.term != TERM_RAW {
			frame.term = TERM_REDUCED
		}

//line ./op-grammar.rl:79
		atm = 4
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st10
	tr82:
//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st10
	tr90:
//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st10
	tr98:
//line ./op-grammar.rl:56

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st10
	tr109:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:37

		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st10
	tr118:
//line ./op-grammar.rl:59
		atoms[atm][1] |= ((uint64)(ATOM_FLOAT)) << 62

//line ./op-grammar.rl:62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st10
	tr126:
//line ./op-grammar.rl:62

//line ./op-grammar.rl:37
		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st10
	tr137:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:37

		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st10
	tr146:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
		atoms[atm][1] = ((uint64)(UUID_NAME)) << 60

//line ./op-grammar.rl:37

		// TODO max size for int/float/string
		atm++
		hlf = 0
		dgt = 0

//line ./op-grammar.rl:33
		dgt = 0
		atoms = append(atoms, Atom{})

		goto st10
	st10:
		if (frame.Parser.position)++; (frame.Parser.position) == pe {
			goto _test_eof10
		}
	st_case_10:
//line dfa.go:3290
		switch (frame.Body)[(frame.Parser.position)] {
		case 32:
			goto st10
		case 43:
			goto st11
		case 45:
			goto st11
		}
		switch {
		case (frame.Body)[(frame.Parser.position)] > 13:
			if 48 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 57 {
				goto st12
			}
		case (frame.Body)[(frame.Parser.position)] >= 9:
			goto st10
		}
		goto st0
	st11:
		if (frame.Parser.position)++; (frame.Parser.position) == pe {
			goto _test_eof11
		}
	st_case_11:
		if 48 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 57 {
			goto st12
		}
		goto st0
	st12:
		if (frame.Parser.position)++; (frame.Parser.position) == pe {
			goto _test_eof12
		}
	st_case_12:
		if (frame.Body)[(frame.Parser.position)] == 46 {
			goto st13
		}
		if 48 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 57 {
			goto st12
		}
		goto st0
	st13:
		if (frame.Parser.position)++; (frame.Parser.position) == pe {
			goto _test_eof13
		}
	st_case_13:
		if 48 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 57 {
			goto st29
		}
		goto st0
	st29:
		if (frame.Parser.position)++; (frame.Parser.position) == pe {
			goto _test_eof29
		}
	st_case_29:
		switch (frame.Body)[(frame.Parser.position)] {
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
			goto tr115
		case 62:
			goto tr116
		case 63:
			goto tr111
		case 64:
			goto tr112
		case 69:
			goto tr117
		case 94:
			goto tr118
		case 101:
			goto tr117
		}
		switch {
		case (frame.Body)[(frame.Parser.position)] > 13:
			if 48 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 57 {
				goto st29
			}
		case (frame.Body)[(frame.Parser.position)] >= 9:
			goto tr110
		}
		goto st0
	tr117:
//line ./op-grammar.rl:59
		atoms[atm][1] |= ((uint64)(ATOM_FLOAT)) << 62

		goto st14
	st14:
		if (frame.Parser.position)++; (frame.Parser.position) == pe {
			goto _test_eof14
		}
	st_case_14:
//line dfa.go:3397
		switch (frame.Body)[(frame.Parser.position)] {
		case 43:
			goto st15
		case 45:
			goto st15
		}
		if 48 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 57 {
			goto st30
		}
		goto st0
	st15:
		if (frame.Parser.position)++; (frame.Parser.position) == pe {
			goto _test_eof15
		}
	st_case_15:
		if 48 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 57 {
			goto st30
		}
		goto st0
	st30:
		if (frame.Parser.position)++; (frame.Parser.position) == pe {
			goto _test_eof30
		}
	st_case_30:
		switch (frame.Body)[(frame.Parser.position)] {
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
			goto tr124
		case 62:
			goto tr125
		case 63:
			goto tr120
		case 64:
			goto tr121
		case 94:
			goto tr126
		}
		switch {
		case (frame.Body)[(frame.Parser.position)] > 13:
			if 48 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 57 {
				goto st30
			}
		case (frame.Body)[(frame.Parser.position)] >= 9:
			goto tr119
		}
		goto st0
	tr36:
//line ././uuid-grammar.rl:5

//line ././uuid-grammar.rl:26

//line ././uuid-grammar.rl:9
		dgt = int(ABC[(frame.Body)[(frame.Parser.position)]]) + 4
		atoms[atm][hlf] &= INT60_FLAGS | PREFIX_MASKS[dgt]

		goto st31
	tr134:
//line ././uuid-grammar.rl:14
		atoms[atm][hlf] |= ((uint64)(ABC[(frame.Body)[(frame.Parser.position)]])) << DIGIT_OFFSETS[dgt]
		dgt++
		if dgt > 10 {
			{
				(frame.Parser.position)++
				(frame.Parser.state) = 31
				goto _out
			}
		}

		goto st31
	st31:
		if (frame.Parser.position)++; (frame.Parser.position) == pe {
			goto _test_eof31
		}
	st_case_31:
//line dfa.go:3489
		switch (frame.Body)[(frame.Parser.position)] {
		case 32:
			goto tr127
		case 33:
			goto tr128
		case 35:
			goto tr129
		case 39:
			goto tr131
		case 42:
			goto tr129
		case 44:
			goto tr128
		case 46:
			goto tr133
		case 58:
			goto tr129
		case 59:
			goto tr128
		case 61:
			goto tr135
		case 62:
			goto tr136
		case 63:
			goto tr128
		case 64:
			goto tr129
		case 91:
			goto tr132
		case 93:
			goto tr132
		case 94:
			goto tr137
		case 95:
			goto tr134
		case 123:
			goto tr132
		case 125:
			goto tr132
		case 126:
			goto tr134
		}
		switch {
		case (frame.Body)[(frame.Parser.position)] < 43:
			switch {
			case (frame.Body)[(frame.Parser.position)] < 36:
				if 9 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 13 {
					goto tr127
				}
			case (frame.Body)[(frame.Parser.position)] > 37:
				if 40 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 41 {
					goto tr132
				}
			default:
				goto tr130
			}
		case (frame.Body)[(frame.Parser.position)] > 45:
			switch {
			case (frame.Body)[(frame.Parser.position)] < 65:
				if 48 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 57 {
					goto tr134
				}
			case (frame.Body)[(frame.Parser.position)] > 90:
				if 97 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 122 {
					goto tr134
				}
			default:
				goto tr134
			}
		default:
			goto tr130
		}
		goto st0
	tr37:
//line ././uuid-grammar.rl:5

//line ././uuid-grammar.rl:26

//line ././uuid-grammar.rl:22
		atoms[atm][hlf] &= INT60_FLAGS

//line ././uuid-grammar.rl:14
		atoms[atm][hlf] |= ((uint64)(ABC[(frame.Body)[(frame.Parser.position)]])) << DIGIT_OFFSETS[dgt]
		dgt++
		if dgt > 10 {
			{
				(frame.Parser.position)++
				(frame.Parser.state) = 32
				goto _out
			}
		}

		goto st32
	tr143:
//line ././uuid-grammar.rl:14
		atoms[atm][hlf] |= ((uint64)(ABC[(frame.Body)[(frame.Parser.position)]])) << DIGIT_OFFSETS[dgt]
		dgt++
		if dgt > 10 {
			{
				(frame.Parser.position)++
				(frame.Parser.state) = 32
				goto _out
			}
		}

		goto st32
	st32:
		if (frame.Parser.position)++; (frame.Parser.position) == pe {
			goto _test_eof32
		}
	st_case_32:
//line dfa.go:3598
		switch (frame.Body)[(frame.Parser.position)] {
		case 32:
			goto tr138
		case 33:
			goto tr139
		case 35:
			goto tr140
		case 39:
			goto tr141
		case 42:
			goto tr140
		case 44:
			goto tr139
		case 46:
			goto tr142
		case 58:
			goto tr140
		case 59:
			goto tr139
		case 61:
			goto tr144
		case 62:
			goto tr145
		case 63:
			goto tr139
		case 64:
			goto tr140
		case 91:
			goto tr132
		case 93:
			goto tr132
		case 94:
			goto tr146
		case 95:
			goto tr143
		case 123:
			goto tr132
		case 125:
			goto tr132
		case 126:
			goto tr143
		}
		switch {
		case (frame.Body)[(frame.Parser.position)] < 43:
			switch {
			case (frame.Body)[(frame.Parser.position)] < 36:
				if 9 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 13 {
					goto tr138
				}
			case (frame.Body)[(frame.Parser.position)] > 37:
				if 40 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 41 {
					goto tr132
				}
			default:
				goto tr130
			}
		case (frame.Body)[(frame.Parser.position)] > 45:
			switch {
			case (frame.Body)[(frame.Parser.position)] < 65:
				if 48 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 57 {
					goto tr143
				}
			case (frame.Body)[(frame.Parser.position)] > 90:
				if 97 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 122 {
					goto tr143
				}
			default:
				goto tr143
			}
		default:
			goto tr130
		}
		goto st0
	tr26:
//line ./op-grammar.rl:65
		atoms[atm][0] = ((uint64)(frame.Parser.position)) << 32

		goto st16
	st16:
		if (frame.Parser.position)++; (frame.Parser.position) == pe {
			goto _test_eof16
		}
	st_case_16:
//line dfa.go:3683
		switch (frame.Body)[(frame.Parser.position)] {
		case 10:
			goto st0
		case 13:
			goto st0
		}
		goto st6
	tr15:
//line ././uuid-grammar.rl:29
		dgt = 0
		hlf = 1

//line ././uuid-grammar.rl:9
		dgt = int(ABC[(frame.Body)[(frame.Parser.position)]]) + 4
		atoms[atm][hlf] &= INT60_FLAGS | PREFIX_MASKS[dgt]

		goto st17
	tr16:
//line ././uuid-grammar.rl:29
		dgt = 0
		hlf = 1

//line ././uuid-grammar.rl:22
		atoms[atm][hlf] &= INT60_FLAGS

//line ././uuid-grammar.rl:14
		atoms[atm][hlf] |= ((uint64)(ABC[(frame.Body)[(frame.Parser.position)]])) << DIGIT_OFFSETS[dgt]
		dgt++
		if dgt > 10 {
			{
				(frame.Parser.position)++
				(frame.Parser.state) = 17
				goto _out
			}
		}

		goto st17
	tr49:
//line ././uuid-grammar.rl:14
		atoms[atm][hlf] |= ((uint64)(ABC[(frame.Body)[(frame.Parser.position)]])) << DIGIT_OFFSETS[dgt]
		dgt++
		if dgt > 10 {
			{
				(frame.Parser.position)++
				(frame.Parser.state) = 17
				goto _out
			}
		}

		goto st17
	tr58:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:29
		dgt = 0
		hlf = 1

//line ././uuid-grammar.rl:9
		dgt = int(ABC[(frame.Body)[(frame.Parser.position)]]) + 4
		atoms[atm][hlf] &= INT60_FLAGS | PREFIX_MASKS[dgt]

		goto st17
	st17:
		if (frame.Parser.position)++; (frame.Parser.position) == pe {
			goto _test_eof17
		}
	st_case_17:
//line dfa.go:3752
		switch (frame.Body)[(frame.Parser.position)] {
		case 32:
			goto tr45
		case 33:
			goto tr46
		case 35:
			goto tr47
		case 39:
			goto tr48
		case 42:
			goto tr47
		case 44:
			goto tr46
		case 58:
			goto tr47
		case 59:
			goto tr46
		case 61:
			goto tr50
		case 62:
			goto tr51
		case 63:
			goto tr46
		case 64:
			goto tr47
		case 94:
			goto tr52
		case 95:
			goto tr49
		case 126:
			goto tr49
		}
		switch {
		case (frame.Body)[(frame.Parser.position)] < 48:
			if 9 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 13 {
				goto tr45
			}
		case (frame.Body)[(frame.Parser.position)] > 57:
			switch {
			case (frame.Body)[(frame.Parser.position)] > 90:
				if 97 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 122 {
					goto tr49
				}
			case (frame.Body)[(frame.Parser.position)] >= 65:
				goto tr49
			}
		default:
			goto tr49
		}
		goto st0
	tr8:
//line ././uuid-grammar.rl:5

//line ././uuid-grammar.rl:26

//line ././uuid-grammar.rl:9
		dgt = int(ABC[(frame.Body)[(frame.Parser.position)]]) + 4
		atoms[atm][hlf] &= INT60_FLAGS | PREFIX_MASKS[dgt]

		goto st18
	tr59:
//line ././uuid-grammar.rl:14
		atoms[atm][hlf] |= ((uint64)(ABC[(frame.Body)[(frame.Parser.position)]])) << DIGIT_OFFSETS[dgt]
		dgt++
		if dgt > 10 {
			{
				(frame.Parser.position)++
				(frame.Parser.state) = 18
				goto _out
			}
		}

		goto st18
	st18:
		if (frame.Parser.position)++; (frame.Parser.position) == pe {
			goto _test_eof18
		}
	st_case_18:
//line dfa.go:3831
		switch (frame.Body)[(frame.Parser.position)] {
		case 32:
			goto tr53
		case 33:
			goto tr54
		case 35:
			goto tr55
		case 39:
			goto tr57
		case 42:
			goto tr55
		case 44:
			goto tr54
		case 58:
			goto tr55
		case 59:
			goto tr54
		case 61:
			goto tr60
		case 62:
			goto tr61
		case 63:
			goto tr54
		case 64:
			goto tr55
		case 91:
			goto tr58
		case 93:
			goto tr58
		case 94:
			goto tr62
		case 95:
			goto tr59
		case 123:
			goto tr58
		case 125:
			goto tr58
		case 126:
			goto tr59
		}
		switch {
		case (frame.Body)[(frame.Parser.position)] < 43:
			switch {
			case (frame.Body)[(frame.Parser.position)] < 36:
				if 9 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 13 {
					goto tr53
				}
			case (frame.Body)[(frame.Parser.position)] > 37:
				if 40 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 41 {
					goto tr58
				}
			default:
				goto tr56
			}
		case (frame.Body)[(frame.Parser.position)] > 45:
			switch {
			case (frame.Body)[(frame.Parser.position)] < 65:
				if 48 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 57 {
					goto tr59
				}
			case (frame.Body)[(frame.Parser.position)] > 90:
				if 97 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 122 {
					goto tr59
				}
			default:
				goto tr59
			}
		default:
			goto tr56
		}
		goto st0
	tr9:
//line ././uuid-grammar.rl:5

//line ././uuid-grammar.rl:26

//line ././uuid-grammar.rl:22
		atoms[atm][hlf] &= INT60_FLAGS

//line ././uuid-grammar.rl:14
		atoms[atm][hlf] |= ((uint64)(ABC[(frame.Body)[(frame.Parser.position)]])) << DIGIT_OFFSETS[dgt]
		dgt++
		if dgt > 10 {
			{
				(frame.Parser.position)++
				(frame.Parser.state) = 19
				goto _out
			}
		}

		goto st19
	tr67:
//line ././uuid-grammar.rl:14
		atoms[atm][hlf] |= ((uint64)(ABC[(frame.Body)[(frame.Parser.position)]])) << DIGIT_OFFSETS[dgt]
		dgt++
		if dgt > 10 {
			{
				(frame.Parser.position)++
				(frame.Parser.state) = 19
				goto _out
			}
		}

		goto st19
	st19:
		if (frame.Parser.position)++; (frame.Parser.position) == pe {
			goto _test_eof19
		}
	st_case_19:
//line dfa.go:3938
		switch (frame.Body)[(frame.Parser.position)] {
		case 32:
			goto tr63
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
		case 91:
			goto tr58
		case 93:
			goto tr58
		case 94:
			goto tr70
		case 95:
			goto tr67
		case 123:
			goto tr58
		case 125:
			goto tr58
		case 126:
			goto tr67
		}
		switch {
		case (frame.Body)[(frame.Parser.position)] < 43:
			switch {
			case (frame.Body)[(frame.Parser.position)] < 36:
				if 9 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 13 {
					goto tr63
				}
			case (frame.Body)[(frame.Parser.position)] > 37:
				if 40 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 41 {
					goto tr58
				}
			default:
				goto tr56
			}
		case (frame.Body)[(frame.Parser.position)] > 45:
			switch {
			case (frame.Body)[(frame.Parser.position)] < 65:
				if 48 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 57 {
					goto tr67
				}
			case (frame.Body)[(frame.Parser.position)] > 90:
				if 97 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 122 {
					goto tr67
				}
			default:
				goto tr67
			}
		default:
			goto tr56
		}
		goto st0
	tr13:
//line ./op-grammar.rl:6
		if atm > 0 {
			atoms[atm] = atoms[atm-1]
		}

		goto st20
	st20:
		if (frame.Parser.position)++; (frame.Parser.position) == pe {
			goto _test_eof20
		}
	st_case_20:
//line dfa.go:4023
		switch (frame.Body)[(frame.Parser.position)] {
		case 32:
			goto tr14
		case 33:
			goto tr4
		case 35:
			goto tr5
		case 39:
			goto tr7
		case 42:
			goto tr5
		case 44:
			goto tr4
		case 58:
			goto tr5
		case 59:
			goto tr4
		case 61:
			goto tr10
		case 62:
			goto tr11
		case 63:
			goto tr4
		case 64:
			goto tr5
		case 91:
			goto tr8
		case 93:
			goto tr8
		case 94:
			goto tr12
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
		case (frame.Body)[(frame.Parser.position)] < 43:
			switch {
			case (frame.Body)[(frame.Parser.position)] < 36:
				if 9 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 13 {
					goto tr14
				}
			case (frame.Body)[(frame.Parser.position)] > 37:
				if 40 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 41 {
					goto tr8
				}
			default:
				goto tr6
			}
		case (frame.Body)[(frame.Parser.position)] > 45:
			switch {
			case (frame.Body)[(frame.Parser.position)] < 65:
				if 48 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 57 {
					goto tr9
				}
			case (frame.Body)[(frame.Parser.position)] > 90:
				if 97 <= (frame.Body)[(frame.Parser.position)] && (frame.Body)[(frame.Parser.position)] <= 122 {
					goto tr9
				}
			default:
				goto tr9
			}
		default:
			goto tr6
		}
		goto st0
	st_out:
	_test_eof21:
		(frame.Parser.state) = 21
		goto _test_eof
	_test_eof1:
		(frame.Parser.state) = 1
		goto _test_eof
	_test_eof2:
		(frame.Parser.state) = 2
		goto _test_eof
	_test_eof22:
		(frame.Parser.state) = 22
		goto _test_eof
	_test_eof23:
		(frame.Parser.state) = 23
		goto _test_eof
	_test_eof3:
		(frame.Parser.state) = 3
		goto _test_eof
	_test_eof4:
		(frame.Parser.state) = 4
		goto _test_eof
	_test_eof5:
		(frame.Parser.state) = 5
		goto _test_eof
	_test_eof6:
		(frame.Parser.state) = 6
		goto _test_eof
	_test_eof24:
		(frame.Parser.state) = 24
		goto _test_eof
	_test_eof25:
		(frame.Parser.state) = 25
		goto _test_eof
	_test_eof7:
		(frame.Parser.state) = 7
		goto _test_eof
	_test_eof8:
		(frame.Parser.state) = 8
		goto _test_eof
	_test_eof26:
		(frame.Parser.state) = 26
		goto _test_eof
	_test_eof9:
		(frame.Parser.state) = 9
		goto _test_eof
	_test_eof27:
		(frame.Parser.state) = 27
		goto _test_eof
	_test_eof28:
		(frame.Parser.state) = 28
		goto _test_eof
	_test_eof10:
		(frame.Parser.state) = 10
		goto _test_eof
	_test_eof11:
		(frame.Parser.state) = 11
		goto _test_eof
	_test_eof12:
		(frame.Parser.state) = 12
		goto _test_eof
	_test_eof13:
		(frame.Parser.state) = 13
		goto _test_eof
	_test_eof29:
		(frame.Parser.state) = 29
		goto _test_eof
	_test_eof14:
		(frame.Parser.state) = 14
		goto _test_eof
	_test_eof15:
		(frame.Parser.state) = 15
		goto _test_eof
	_test_eof30:
		(frame.Parser.state) = 30
		goto _test_eof
	_test_eof31:
		(frame.Parser.state) = 31
		goto _test_eof
	_test_eof32:
		(frame.Parser.state) = 32
		goto _test_eof
	_test_eof16:
		(frame.Parser.state) = 16
		goto _test_eof
	_test_eof17:
		(frame.Parser.state) = 17
		goto _test_eof
	_test_eof18:
		(frame.Parser.state) = 18
		goto _test_eof
	_test_eof19:
		(frame.Parser.state) = 19
		goto _test_eof
	_test_eof20:
		(frame.Parser.state) = 20
		goto _test_eof

	_test_eof:
		{
		}
		if (frame.Parser.position) == eof {
			switch frame.Parser.state {
			case 22:
//line ./op-grammar.rl:104

			case 25:
//line ./op-grammar.rl:84
//line ./op-grammar.rl:104

			case 24, 27:
//line ./op-grammar.rl:37
				// TODO max size for int/float/string
				atm++
				hlf = 0
				dgt = 0

//line ./op-grammar.rl:84

//line ./op-grammar.rl:104

			case 31:
//line ././uuid-grammar.rl:34

//line ./op-grammar.rl:37

				// TODO max size for int/float/string
				atm++
				hlf = 0
				dgt = 0

//line ./op-grammar.rl:84

//line ./op-grammar.rl:104

			case 28:
//line ././uuid-grammar.rl:37

//line ./op-grammar.rl:37

				// TODO max size for int/float/string
				atm++
				hlf = 0
				dgt = 0

//line ./op-grammar.rl:84

//line ./op-grammar.rl:104

			case 26:
//line ./op-grammar.rl:56
//line ./op-grammar.rl:37
				// TODO max size for int/float/string
				atm++
				hlf = 0
				dgt = 0

//line ./op-grammar.rl:84

//line ./op-grammar.rl:104

			case 30:
//line ./op-grammar.rl:62
//line ./op-grammar.rl:37
				// TODO max size for int/float/string
				atm++
				hlf = 0
				dgt = 0

//line ./op-grammar.rl:84

//line ./op-grammar.rl:104

			case 32:
//line ././uuid-grammar.rl:34

//line ././uuid-grammar.rl:46
				atoms[atm][1] = ((uint64)(UUID_NAME)) << 60

//line ./op-grammar.rl:37

				// TODO max size for int/float/string
				atm++
				hlf = 0
				dgt = 0

//line ./op-grammar.rl:84

//line ./op-grammar.rl:104

			case 29:
//line ./op-grammar.rl:59
				atoms[atm][1] |= ((uint64)(ATOM_FLOAT)) << 62

//line ./op-grammar.rl:62

//line ./op-grammar.rl:37
				// TODO max size for int/float/string
				atm++
				hlf = 0
				dgt = 0

//line ./op-grammar.rl:84

//line ./op-grammar.rl:104

//line dfa.go:4267
			}
		}

	_out:
		{
		}
	}

//line dfa.rl:60

	frame.Parser.atm, frame.Parser.hlf, frame.Parser.dgt = atm, hlf, dgt
	frame.atoms = atoms

	if !frame.Parser.streaming && frame.Parser.state < RON_first_final && frame.Parser.state > 0 {
		frame.Parser.state = RON_error
	}

}

func (ctx_uuid UUID) Parse(data []byte) (UUID, error) {

//line dfa.rl:75
//line dfa.go:4292
	const UUID_start int = 1
	const UUID_first_final int = 2
	const UUID_error int = 0

	const UUID_en_main int = 1

//line dfa.rl:76
	cs, p, pe, eof := 0, 0, len(data), len(data)
	_ = eof

	atm, hlf, dgt := 0, 0, 0

	atoms := [1]Atom{Atom(ctx_uuid)}

//line dfa.go:4310
	{
		cs = UUID_start
	}

//line dfa.go:4315
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
		atoms[atm][hlf] &= INT60_FULL
		atoms[atm][hlf] |= ((uint64)(ABC[data[p]])) << 60

		goto st2
	tr7:
//line ./uuid-grammar.rl:34
//line ./uuid-grammar.rl:40
		hlf = 1
		atoms[atm][hlf] &= INT60_FULL
		atoms[atm][hlf] |= ((uint64)(ABC[data[p]])) << 60

		goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
//line dfa.go:4408
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
//line dfa.go:4502
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
//line dfa.go:4550
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
//line dfa.go:4627
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
				atoms[atm][1] = ((uint64)(UUID_NAME)) << 60

//line dfa.go:4694
			}
		}

	_out:
		{
		}
	}

//line dfa.rl:91

	if cs < UUID_first_final || dgt > 10 {
		return ERROR_UUID, errors.New(fmt.Sprintf("parse error at pos %d", p))
	} else {
		return UUID(atoms[0]), nil
	}

}
