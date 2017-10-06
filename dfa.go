//line dfa.rl:1
package RON

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

// Parse consumes one op, unless the buffer ends earlier.
func (it *Frame) Parse() {

	if it.state.p >= len(it.state.data) {
		if !it.state.streaming {
			it.Op = ZERO_OP
			it.state.cs = RON_error
		}
		return
	}

	if it.state.cs == 0 && it.state.p == 0 {

//line dfa.go:38
		{
			it.state.cs = RON_start
		}

//line dfa.rl:24
	}

	had_end := false
	p, pe, eof := it.state.p, len(it.state.data), len(it.state.data)
	n := uint(0)
	_ = eof
	_ = pe // FIXME kill

	if it.state.streaming {
		eof = -1
	}

	i := it.state.incomplete
	idx := it.state.idx
	half := it.state.half
	digit := it.state.digit

//line dfa.go:62
	{
		if p == pe {
			goto _test_eof
		}
		goto _resume

	_again:
		switch it.state.cs {
		case 21:
			goto st21
		case 0:
			goto st0
		case 1:
			goto st1
		case 2:
			goto st2
		case 3:
			goto st3
		case 4:
			goto st4
		case 22:
			goto st22
		case 23:
			goto st23
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

		if p++; p == pe {
			goto _test_eof
		}
	_resume:
		switch it.state.cs {
		case 21:
			goto st_case_21
		case 0:
			goto st_case_0
		case 1:
			goto st_case_1
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		case 4:
			goto st_case_4
		case 22:
			goto st_case_22
		case 23:
			goto st_case_23
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
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
		switch it.state.data[p] {
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
		if 9 <= it.state.data[p] && it.state.data[p] <= 13 {
			goto st1
		}
		goto st0
	st_case_0:
	st0:
		it.state.cs = 0
		goto _out
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		switch it.state.data[p] {
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
		if 9 <= it.state.data[p] && it.state.data[p] <= 13 {
			goto st1
		}
		goto st0
	tr2:
		it.state.cs = 2
//line ./op-grammar.rl:85
		idx = 0
		if had_end {
			// one op is done, so stop parsing for now
			// make sure the parser restarts with the next op
			p--
			it.state.cs = (RON_start)
			{
				p++
				goto _out
			}
		}

//line ./op-grammar.rl:12
		n = specSep2Bits(it.state.data[p])
		if n < idx {
			// wrong UUID order; must be type-object-event-ref
			it.state.cs = (RON_error)
			{
				p++
				goto _out
			}
		} else {
			// start parsing the UUID
			idx = n
			i = it.uuids[idx].uint128
			digit = 0
		}

		goto _again
	tr10:
		it.state.cs = 2
//line ./op-grammar.rl:26
		// OK, save the UUID
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:12
		n = specSep2Bits(it.state.data[p])
		if n < idx {
			// wrong UUID order; must be type-object-event-ref
			it.state.cs = (RON_error)
			{
				p++
				goto _out
			}
		} else {
			// start parsing the UUID
			idx = n
			i = it.uuids[idx].uint128
			digit = 0
		}

		goto _again
	tr19:
		it.state.cs = 2
//line ./op-grammar.rl:12
		n = specSep2Bits(it.state.data[p])
		if n < idx {
			// wrong UUID order; must be type-object-event-ref
			it.state.cs = (RON_error)
			{
				p++
				goto _out
			}
		} else {
			// start parsing the UUID
			idx = n
			i = it.uuids[idx].uint128
			digit = 0
		}

		goto _again
	tr47:
		it.state.cs = 2
//line ././uuid-grammar.rl:38

//line ./op-grammar.rl:26

		// OK, save the UUID
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:12
		n = specSep2Bits(it.state.data[p])
		if n < idx {
			// wrong UUID order; must be type-object-event-ref
			it.state.cs = (RON_error)
			{
				p++
				goto _out
			}
		} else {
			// start parsing the UUID
			idx = n
			i = it.uuids[idx].uint128
			digit = 0
		}

		goto _again
	tr55:
		it.state.cs = 2
//line ././uuid-grammar.rl:35

//line ./op-grammar.rl:26

		// OK, save the UUID
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:12
		n = specSep2Bits(it.state.data[p])
		if n < idx {
			// wrong UUID order; must be type-object-event-ref
			it.state.cs = (RON_error)
			{
				p++
				goto _out
			}
		} else {
			// start parsing the UUID
			idx = n
			i = it.uuids[idx].uint128
			digit = 0
		}

		goto _again
	tr65:
		it.state.cs = 2
//line ././uuid-grammar.rl:35

//line ././uuid-grammar.rl:47
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:26

		// OK, save the UUID
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:12
		n = specSep2Bits(it.state.data[p])
		if n < idx {
			// wrong UUID order; must be type-object-event-ref
			it.state.cs = (RON_error)
			{
				p++
				goto _out
			}
		} else {
			// start parsing the UUID
			idx = n
			i = it.uuids[idx].uint128
			digit = 0
		}

		goto _again
	tr73:
		it.state.cs = 2
//line ./op-grammar.rl:96
		had_end = true

//line ./op-grammar.rl:85
		idx = 0
		if had_end {
			// one op is done, so stop parsing for now
			// make sure the parser restarts with the next op
			p--
			it.state.cs = (RON_start)
			{
				p++
				goto _out
			}
		}

//line ./op-grammar.rl:12
		n = specSep2Bits(it.state.data[p])
		if n < idx {
			// wrong UUID order; must be type-object-event-ref
			it.state.cs = (RON_error)
			{
				p++
				goto _out
			}
		} else {
			// start parsing the UUID
			idx = n
			i = it.uuids[idx].uint128
			digit = 0
		}

		goto _again
	tr77:
		it.state.cs = 2
//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:78

//line ./op-grammar.rl:96
		had_end = true

//line ./op-grammar.rl:85
		idx = 0
		if had_end {
			// one op is done, so stop parsing for now
			// make sure the parser restarts with the next op
			p--
			it.state.cs = (RON_start)
			{
				p++
				goto _out
			}
		}

//line ./op-grammar.rl:12
		n = specSep2Bits(it.state.data[p])
		if n < idx {
			// wrong UUID order; must be type-object-event-ref
			it.state.cs = (RON_error)
			{
				p++
				goto _out
			}
		} else {
			// start parsing the UUID
			idx = n
			i = it.uuids[idx].uint128
			digit = 0
		}

		goto _again
	tr85:
		it.state.cs = 2
//line ./op-grammar.rl:78

//line ./op-grammar.rl:96
		had_end = true

//line ./op-grammar.rl:85
		idx = 0
		if had_end {
			// one op is done, so stop parsing for now
			// make sure the parser restarts with the next op
			p--
			it.state.cs = (RON_start)
			{
				p++
				goto _out
			}
		}

//line ./op-grammar.rl:12
		n = specSep2Bits(it.state.data[p])
		if n < idx {
			// wrong UUID order; must be type-object-event-ref
			it.state.cs = (RON_error)
			{
				p++
				goto _out
			}
		} else {
			// start parsing the UUID
			idx = n
			i = it.uuids[idx].uint128
			digit = 0
		}

		goto _again
	tr93:
		it.state.cs = 2
//line ./op-grammar.rl:52
		i[1] |= ATOM_INT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:78

//line ./op-grammar.rl:96
		had_end = true

//line ./op-grammar.rl:85
		idx = 0
		if had_end {
			// one op is done, so stop parsing for now
			// make sure the parser restarts with the next op
			p--
			it.state.cs = (RON_start)
			{
				p++
				goto _out
			}
		}

//line ./op-grammar.rl:12
		n = specSep2Bits(it.state.data[p])
		if n < idx {
			// wrong UUID order; must be type-object-event-ref
			it.state.cs = (RON_error)
			{
				p++
				goto _out
			}
		} else {
			// start parsing the UUID
			idx = n
			i = it.uuids[idx].uint128
			digit = 0
		}

		goto _again
	tr103:
		it.state.cs = 2
//line ././uuid-grammar.rl:38

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:78

//line ./op-grammar.rl:96
		had_end = true

//line ./op-grammar.rl:85
		idx = 0
		if had_end {
			// one op is done, so stop parsing for now
			// make sure the parser restarts with the next op
			p--
			it.state.cs = (RON_start)
			{
				p++
				goto _out
			}
		}

//line ./op-grammar.rl:12
		n = specSep2Bits(it.state.data[p])
		if n < idx {
			// wrong UUID order; must be type-object-event-ref
			it.state.cs = (RON_error)
			{
				p++
				goto _out
			}
		} else {
			// start parsing the UUID
			idx = n
			i = it.uuids[idx].uint128
			digit = 0
		}

		goto _again
	tr112:
		it.state.cs = 2
//line ./op-grammar.rl:56

//line ./op-grammar.rl:58
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:78

//line ./op-grammar.rl:96
		had_end = true

//line ./op-grammar.rl:85
		idx = 0
		if had_end {
			// one op is done, so stop parsing for now
			// make sure the parser restarts with the next op
			p--
			it.state.cs = (RON_start)
			{
				p++
				goto _out
			}
		}

//line ./op-grammar.rl:12
		n = specSep2Bits(it.state.data[p])
		if n < idx {
			// wrong UUID order; must be type-object-event-ref
			it.state.cs = (RON_error)
			{
				p++
				goto _out
			}
		} else {
			// start parsing the UUID
			idx = n
			i = it.uuids[idx].uint128
			digit = 0
		}

		goto _again
	tr121:
		it.state.cs = 2
//line ./op-grammar.rl:58
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:78

//line ./op-grammar.rl:96
		had_end = true

//line ./op-grammar.rl:85
		idx = 0
		if had_end {
			// one op is done, so stop parsing for now
			// make sure the parser restarts with the next op
			p--
			it.state.cs = (RON_start)
			{
				p++
				goto _out
			}
		}

//line ./op-grammar.rl:12
		n = specSep2Bits(it.state.data[p])
		if n < idx {
			// wrong UUID order; must be type-object-event-ref
			it.state.cs = (RON_error)
			{
				p++
				goto _out
			}
		} else {
			// start parsing the UUID
			idx = n
			i = it.uuids[idx].uint128
			digit = 0
		}

		goto _again
	tr129:
		it.state.cs = 2
//line ././uuid-grammar.rl:35

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:78

//line ./op-grammar.rl:96
		had_end = true

//line ./op-grammar.rl:85
		idx = 0
		if had_end {
			// one op is done, so stop parsing for now
			// make sure the parser restarts with the next op
			p--
			it.state.cs = (RON_start)
			{
				p++
				goto _out
			}
		}

//line ./op-grammar.rl:12
		n = specSep2Bits(it.state.data[p])
		if n < idx {
			// wrong UUID order; must be type-object-event-ref
			it.state.cs = (RON_error)
			{
				p++
				goto _out
			}
		} else {
			// start parsing the UUID
			idx = n
			i = it.uuids[idx].uint128
			digit = 0
		}

		goto _again
	tr140:
		it.state.cs = 2
//line ././uuid-grammar.rl:35

//line ././uuid-grammar.rl:47
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:78

//line ./op-grammar.rl:96
		had_end = true

//line ./op-grammar.rl:85
		idx = 0
		if had_end {
			// one op is done, so stop parsing for now
			// make sure the parser restarts with the next op
			p--
			it.state.cs = (RON_start)
			{
				p++
				goto _out
			}
		}

//line ./op-grammar.rl:12
		n = specSep2Bits(it.state.data[p])
		if n < idx {
			// wrong UUID order; must be type-object-event-ref
			it.state.cs = (RON_error)
			{
				p++
				goto _out
			}
		} else {
			// start parsing the UUID
			idx = n
			i = it.uuids[idx].uint128
			digit = 0
		}

		goto _again
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
//line dfa.go:790
		switch it.state.data[p] {
		case 32:
			goto st2
		case 43:
			goto tr4
		case 45:
			goto tr4
		case 91:
			goto tr5
		case 93:
			goto tr5
		case 96:
			goto tr7
		case 123:
			goto tr5
		case 125:
			goto tr5
		case 126:
			goto tr6
		}
		switch {
		case it.state.data[p] < 40:
			switch {
			case it.state.data[p] > 13:
				if 36 <= it.state.data[p] && it.state.data[p] <= 37 {
					goto tr4
				}
			case it.state.data[p] >= 9:
				goto st2
			}
		case it.state.data[p] > 41:
			switch {
			case it.state.data[p] < 65:
				if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
					goto tr6
				}
			case it.state.data[p] > 90:
				if 95 <= it.state.data[p] && it.state.data[p] <= 122 {
					goto tr6
				}
			default:
				goto tr6
			}
		default:
			goto tr5
		}
		goto st0
	tr4:
//line ././uuid-grammar.rl:5
		half = 0

//line ././uuid-grammar.rl:41
		half |= 1
		i[half] &= INT60_FULL
		i[half] |= uint64(uuidSep2Bits(it.state.data[p])) << 60

		goto st3
	tr56:
//line ././uuid-grammar.rl:35
//line ././uuid-grammar.rl:41
		half |= 1
		i[half] &= INT60_FULL
		i[half] |= uint64(uuidSep2Bits(it.state.data[p])) << 60

		goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
//line dfa.go:866
		switch it.state.data[p] {
		case 32:
			goto tr8
		case 33:
			goto tr9
		case 35:
			goto tr10
		case 39:
			goto tr11
		case 42:
			goto tr10
		case 44:
			goto tr9
		case 58:
			goto tr10
		case 59:
			goto tr9
		case 61:
			goto tr14
		case 62:
			goto tr15
		case 63:
			goto tr9
		case 64:
			goto tr10
		case 91:
			goto tr12
		case 93:
			goto tr12
		case 94:
			goto tr16
		case 95:
			goto tr13
		case 123:
			goto tr12
		case 125:
			goto tr12
		case 126:
			goto tr13
		}
		switch {
		case it.state.data[p] < 48:
			switch {
			case it.state.data[p] > 13:
				if 40 <= it.state.data[p] && it.state.data[p] <= 41 {
					goto tr12
				}
			case it.state.data[p] >= 9:
				goto tr8
			}
		case it.state.data[p] > 57:
			switch {
			case it.state.data[p] > 90:
				if 97 <= it.state.data[p] && it.state.data[p] <= 122 {
					goto tr13
				}
			case it.state.data[p] >= 65:
				goto tr13
			}
		default:
			goto tr13
		}
		goto st0
	tr8:
//line ./op-grammar.rl:26

		// OK, save the UUID
		it.uuids[idx] = UUID{uint128: i}
		idx++

		goto st4
	tr45:
//line ././uuid-grammar.rl:38

//line ./op-grammar.rl:26

		// OK, save the UUID
		it.uuids[idx] = UUID{uint128: i}
		idx++

		goto st4
	tr53:
//line ././uuid-grammar.rl:35

//line ./op-grammar.rl:26

		// OK, save the UUID
		it.uuids[idx] = UUID{uint128: i}
		idx++

		goto st4
	tr63:
//line ././uuid-grammar.rl:35

//line ././uuid-grammar.rl:47
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:26

		// OK, save the UUID
		it.uuids[idx] = UUID{uint128: i}
		idx++

		goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
//line dfa.go:980
		switch it.state.data[p] {
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
		if 9 <= it.state.data[p] && it.state.data[p] <= 13 {
			goto st4
		}
		goto st0
	tr9:
//line ./op-grammar.rl:26

		// OK, save the UUID
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:100
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		// not that necessary: op refs the frame
		it.frame = it.state.data

//line ./op-grammar.rl:81
		it.term = termSep2Bits(it.state.data[p])

		goto st22
	tr18:
//line ./op-grammar.rl:100
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		// not that necessary: op refs the frame
		it.frame = it.state.data

//line ./op-grammar.rl:81
		it.term = termSep2Bits(it.state.data[p])

		goto st22
	tr46:
//line ././uuid-grammar.rl:38

//line ./op-grammar.rl:26

		// OK, save the UUID
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:100
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		// not that necessary: op refs the frame
		it.frame = it.state.data

//line ./op-grammar.rl:81
		it.term = termSep2Bits(it.state.data[p])

		goto st22
	tr54:
//line ././uuid-grammar.rl:35

//line ./op-grammar.rl:26

		// OK, save the UUID
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:100
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		// not that necessary: op refs the frame
		it.frame = it.state.data

//line ./op-grammar.rl:81
		it.term = termSep2Bits(it.state.data[p])

		goto st22
	tr64:
//line ././uuid-grammar.rl:35

//line ././uuid-grammar.rl:47
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:26

		// OK, save the UUID
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:100
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		// not that necessary: op refs the frame
		it.frame = it.state.data

//line ./op-grammar.rl:81
		it.term = termSep2Bits(it.state.data[p])

		goto st22
	tr72:
//line ./op-grammar.rl:96
		had_end = true

		goto st22
	tr76:
//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:78

//line ./op-grammar.rl:81
		it.term = termSep2Bits(it.state.data[p])

		goto st22
	tr84:
//line ./op-grammar.rl:78
//line ./op-grammar.rl:81
		it.term = termSep2Bits(it.state.data[p])

		goto st22
	tr92:
//line ./op-grammar.rl:52
		i[1] |= ATOM_INT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:78

//line ./op-grammar.rl:81
		it.term = termSep2Bits(it.state.data[p])

		goto st22
	tr102:
//line ././uuid-grammar.rl:38

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:78

//line ./op-grammar.rl:81
		it.term = termSep2Bits(it.state.data[p])

		goto st22
	tr111:
//line ./op-grammar.rl:56
//line ./op-grammar.rl:58
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:78

//line ./op-grammar.rl:81
		it.term = termSep2Bits(it.state.data[p])

		goto st22
	tr120:
//line ./op-grammar.rl:58
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:78

//line ./op-grammar.rl:81
		it.term = termSep2Bits(it.state.data[p])

		goto st22
	tr128:
//line ././uuid-grammar.rl:35

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:78

//line ./op-grammar.rl:81
		it.term = termSep2Bits(it.state.data[p])

		goto st22
	tr139:
//line ././uuid-grammar.rl:35

//line ././uuid-grammar.rl:47
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:78

//line ./op-grammar.rl:81
		it.term = termSep2Bits(it.state.data[p])

		goto st22
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
//line dfa.go:1271
		switch it.state.data[p] {
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
		if 9 <= it.state.data[p] && it.state.data[p] <= 13 {
			goto tr72
		}
		goto st0
	tr71:
//line ./op-grammar.rl:109
		it.state.streaming = false

		goto st23
	tr74:
//line ./op-grammar.rl:96
		had_end = true

//line ./op-grammar.rl:109
		it.state.streaming = false

		goto st23
	tr79:
//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:78

//line ./op-grammar.rl:96
		had_end = true

//line ./op-grammar.rl:109
		it.state.streaming = false

		goto st23
	tr87:
//line ./op-grammar.rl:78
//line ./op-grammar.rl:96
		had_end = true

//line ./op-grammar.rl:109
		it.state.streaming = false

		goto st23
	tr95:
//line ./op-grammar.rl:52
		i[1] |= ATOM_INT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:78

//line ./op-grammar.rl:96
		had_end = true

//line ./op-grammar.rl:109
		it.state.streaming = false

		goto st23
	tr105:
//line ././uuid-grammar.rl:38

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:78

//line ./op-grammar.rl:96
		had_end = true

//line ./op-grammar.rl:109
		it.state.streaming = false

		goto st23
	tr114:
//line ./op-grammar.rl:56
//line ./op-grammar.rl:58
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:78

//line ./op-grammar.rl:96
		had_end = true

//line ./op-grammar.rl:109
		it.state.streaming = false

		goto st23
	tr123:
//line ./op-grammar.rl:58
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:78

//line ./op-grammar.rl:96
		had_end = true

//line ./op-grammar.rl:109
		it.state.streaming = false

		goto st23
	tr133:
//line ././uuid-grammar.rl:35

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:78

//line ./op-grammar.rl:96
		had_end = true

//line ./op-grammar.rl:109
		it.state.streaming = false

		goto st23
	tr142:
//line ././uuid-grammar.rl:35

//line ././uuid-grammar.rl:47
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:78

//line ./op-grammar.rl:96
		had_end = true

//line ./op-grammar.rl:109
		it.state.streaming = false

		goto st23
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
//line dfa.go:1478
		goto st0
	tr11:
//line ./op-grammar.rl:26

		// OK, save the UUID
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:100
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		// not that necessary: op refs the frame
		it.frame = it.state.data

//line ./op-grammar.rl:75
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st5
	tr20:
//line ./op-grammar.rl:100
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		// not that necessary: op refs the frame
		it.frame = it.state.data

//line ./op-grammar.rl:75
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st5
	tr48:
//line ././uuid-grammar.rl:38

//line ./op-grammar.rl:26

		// OK, save the UUID
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:100
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		// not that necessary: op refs the frame
		it.frame = it.state.data

//line ./op-grammar.rl:75
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st5
	tr57:
//line ././uuid-grammar.rl:35

//line ./op-grammar.rl:26

		// OK, save the UUID
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:100
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		// not that necessary: op refs the frame
		it.frame = it.state.data

//line ./op-grammar.rl:75
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st5
	tr66:
//line ././uuid-grammar.rl:35

//line ././uuid-grammar.rl:47
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:26

		// OK, save the UUID
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:100
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		// not that necessary: op refs the frame
		it.frame = it.state.data

//line ./op-grammar.rl:75
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st5
	tr78:
//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st5
	tr86:
//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st5
	tr94:
//line ./op-grammar.rl:52
		i[1] |= ATOM_INT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st5
	tr104:
//line ././uuid-grammar.rl:38

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st5
	tr113:
//line ./op-grammar.rl:56

//line ./op-grammar.rl:58
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st5
	tr122:
//line ./op-grammar.rl:58
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st5
	tr131:
//line ././uuid-grammar.rl:35

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st5
	tr141:
//line ././uuid-grammar.rl:35

//line ././uuid-grammar.rl:47
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
//line dfa.go:1741
		switch it.state.data[p] {
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
//line ./op-grammar.rl:62
		i[0] = uint64(p)

		goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
//line dfa.go:1764
		switch it.state.data[p] {
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
//line ./op-grammar.rl:62
		i[0] = uint64(p)

//line ./op-grammar.rl:65
		i[1] = uint64(p) | ATOM_STRING_62

		goto st24
	tr28:
//line ./op-grammar.rl:65
		i[1] = uint64(p) | ATOM_STRING_62

		goto st24
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
//line dfa.go:1797
		switch it.state.data[p] {
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
		if 9 <= it.state.data[p] && it.state.data[p] <= 13 {
			goto tr75
		}
		goto st0
	tr75:
//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:78

//line ./op-grammar.rl:96
		had_end = true

		goto st25
	tr83:
//line ./op-grammar.rl:78
//line ./op-grammar.rl:96
		had_end = true

		goto st25
	tr91:
//line ./op-grammar.rl:52
		i[1] |= ATOM_INT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:78

//line ./op-grammar.rl:96
		had_end = true

		goto st25
	tr101:
//line ././uuid-grammar.rl:38

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:78

//line ./op-grammar.rl:96
		had_end = true

		goto st25
	tr110:
//line ./op-grammar.rl:56
//line ./op-grammar.rl:58
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:78

//line ./op-grammar.rl:96
		had_end = true

		goto st25
	tr119:
//line ./op-grammar.rl:58
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:78

//line ./op-grammar.rl:96
		had_end = true

		goto st25
	tr127:
//line ././uuid-grammar.rl:35

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:78

//line ./op-grammar.rl:96
		had_end = true

		goto st25
	tr138:
//line ././uuid-grammar.rl:35

//line ././uuid-grammar.rl:47
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:78

//line ./op-grammar.rl:96
		had_end = true

		goto st25
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
//line dfa.go:1972
		switch it.state.data[p] {
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
		if 9 <= it.state.data[p] && it.state.data[p] <= 13 {
			goto tr83
		}
		goto st0
	tr14:
//line ./op-grammar.rl:26

		// OK, save the UUID
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:100
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		// not that necessary: op refs the frame
		it.frame = it.state.data

//line ./op-grammar.rl:75
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st7
	tr21:
//line ./op-grammar.rl:100
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		// not that necessary: op refs the frame
		it.frame = it.state.data

//line ./op-grammar.rl:75
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st7
	tr50:
//line ././uuid-grammar.rl:38

//line ./op-grammar.rl:26

		// OK, save the UUID
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:100
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		// not that necessary: op refs the frame
		it.frame = it.state.data

//line ./op-grammar.rl:75
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st7
	tr60:
//line ././uuid-grammar.rl:35

//line ./op-grammar.rl:26

		// OK, save the UUID
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:100
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		// not that necessary: op refs the frame
		it.frame = it.state.data

//line ./op-grammar.rl:75
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st7
	tr68:
//line ././uuid-grammar.rl:35

//line ././uuid-grammar.rl:47
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:26

		// OK, save the UUID
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:100
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		// not that necessary: op refs the frame
		it.frame = it.state.data

//line ./op-grammar.rl:75
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st7
	tr80:
//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st7
	tr88:
//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st7
	tr96:
//line ./op-grammar.rl:52
		i[1] |= ATOM_INT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st7
	tr107:
//line ././uuid-grammar.rl:38

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st7
	tr115:
//line ./op-grammar.rl:56

//line ./op-grammar.rl:58
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st7
	tr124:
//line ./op-grammar.rl:58
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st7
	tr135:
//line ././uuid-grammar.rl:35

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st7
	tr144:
//line ././uuid-grammar.rl:35

//line ././uuid-grammar.rl:47
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
//line dfa.go:2268
		switch it.state.data[p] {
		case 32:
			goto st7
		case 43:
			goto tr31
		case 45:
			goto tr31
		}
		switch {
		case it.state.data[p] > 13:
			if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
				goto tr32
			}
		case it.state.data[p] >= 9:
			goto st7
		}
		goto st0
	tr31:
//line ./op-grammar.rl:41

//line ./op-grammar.rl:43
		if it.state.data[p] == '-' {
			i[1] |= 1
		}

		goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
//line dfa.go:2302
		if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
			goto tr33
		}
		goto st0
	tr32:
//line ./op-grammar.rl:41

//line ./op-grammar.rl:48
		i[0] *= 10
		i[0] += uint64(int(it.state.data[p]) - int('0'))

		goto st26
	tr33:
//line ./op-grammar.rl:48
		i[0] *= 10
		i[0] += uint64(int(it.state.data[p]) - int('0'))

		goto st26
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
//line dfa.go:2329
		switch it.state.data[p] {
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
		case it.state.data[p] > 13:
			if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
				goto tr33
			}
		case it.state.data[p] >= 9:
			goto tr91
		}
		goto st0
	tr15:
//line ./op-grammar.rl:26

		// OK, save the UUID
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:100
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		// not that necessary: op refs the frame
		it.frame = it.state.data

//line ./op-grammar.rl:75
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st9
	tr22:
//line ./op-grammar.rl:100
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		// not that necessary: op refs the frame
		it.frame = it.state.data

//line ./op-grammar.rl:75
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st9
	tr51:
//line ././uuid-grammar.rl:38

//line ./op-grammar.rl:26

		// OK, save the UUID
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:100
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		// not that necessary: op refs the frame
		it.frame = it.state.data

//line ./op-grammar.rl:75
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st9
	tr61:
//line ././uuid-grammar.rl:35

//line ./op-grammar.rl:26

		// OK, save the UUID
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:100
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		// not that necessary: op refs the frame
		it.frame = it.state.data

//line ./op-grammar.rl:75
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st9
	tr69:
//line ././uuid-grammar.rl:35

//line ././uuid-grammar.rl:47
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:26

		// OK, save the UUID
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:100
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		// not that necessary: op refs the frame
		it.frame = it.state.data

//line ./op-grammar.rl:75
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st9
	tr81:
//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st9
	tr89:
//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st9
	tr97:
//line ./op-grammar.rl:52
		i[1] |= ATOM_INT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st9
	tr108:
//line ././uuid-grammar.rl:38

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st9
	tr116:
//line ./op-grammar.rl:56

//line ./op-grammar.rl:58
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st9
	tr125:
//line ./op-grammar.rl:58
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st9
	tr136:
//line ././uuid-grammar.rl:35

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st9
	tr145:
//line ././uuid-grammar.rl:35

//line ././uuid-grammar.rl:47
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
//line dfa.go:2630
		switch it.state.data[p] {
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
		case it.state.data[p] < 40:
			switch {
			case it.state.data[p] > 13:
				if 36 <= it.state.data[p] && it.state.data[p] <= 37 {
					goto tr35
				}
			case it.state.data[p] >= 9:
				goto st9
			}
		case it.state.data[p] > 41:
			switch {
			case it.state.data[p] < 65:
				if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
					goto tr37
				}
			case it.state.data[p] > 90:
				if 97 <= it.state.data[p] && it.state.data[p] <= 122 {
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
		half = 0

//line ././uuid-grammar.rl:41
		half |= 1
		i[half] &= INT60_FULL
		i[half] |= uint64(uuidSep2Bits(it.state.data[p])) << 60

		goto st27
	tr130:
//line ././uuid-grammar.rl:35
//line ././uuid-grammar.rl:41
		half |= 1
		i[half] &= INT60_FULL
		i[half] |= uint64(uuidSep2Bits(it.state.data[p])) << 60

		goto st27
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
//line dfa.go:2706
		switch it.state.data[p] {
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
		case it.state.data[p] < 48:
			switch {
			case it.state.data[p] > 13:
				if 40 <= it.state.data[p] && it.state.data[p] <= 41 {
					goto tr99
				}
			case it.state.data[p] >= 9:
				goto tr75
			}
		case it.state.data[p] > 57:
			switch {
			case it.state.data[p] > 90:
				if 97 <= it.state.data[p] && it.state.data[p] <= 122 {
					goto tr100
				}
			case it.state.data[p] >= 65:
				goto tr100
			}
		default:
			goto tr100
		}
		goto st0
	tr99:
//line ././uuid-grammar.rl:30
		digit = 0
		half |= 1

//line ././uuid-grammar.rl:10
		digit = prefixSep2Bits(it.state.data[p]) + 4
		i[half] &= INT60_FLAGS | PREFIX_MASKS[digit]

		goto st28
	tr100:
//line ././uuid-grammar.rl:30
		digit = 0
		half |= 1

//line ././uuid-grammar.rl:23
		i[half] &= INT60_FLAGS

//line ././uuid-grammar.rl:15
		i[half] |= uint64(ABC[it.state.data[p]]) << DIGIT_OFFSETS[digit]
		digit++
		if digit > 10 {
			{
				p++
				it.state.cs = 28
				goto _out
			}
		}

		goto st28
	tr106:
//line ././uuid-grammar.rl:15
		i[half] |= uint64(ABC[it.state.data[p]]) << DIGIT_OFFSETS[digit]
		digit++
		if digit > 10 {
			{
				p++
				it.state.cs = 28
				goto _out
			}
		}

		goto st28
	tr132:
//line ././uuid-grammar.rl:35

//line ././uuid-grammar.rl:30
		digit = 0
		half |= 1

//line ././uuid-grammar.rl:10
		digit = prefixSep2Bits(it.state.data[p]) + 4
		i[half] &= INT60_FLAGS | PREFIX_MASKS[digit]

		goto st28
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
//line dfa.go:2833
		switch it.state.data[p] {
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
		case it.state.data[p] < 48:
			if 9 <= it.state.data[p] && it.state.data[p] <= 13 {
				goto tr101
			}
		case it.state.data[p] > 57:
			switch {
			case it.state.data[p] > 90:
				if 97 <= it.state.data[p] && it.state.data[p] <= 122 {
					goto tr106
				}
			case it.state.data[p] >= 65:
				goto tr106
			}
		default:
			goto tr106
		}
		goto st0
	tr16:
//line ./op-grammar.rl:26

		// OK, save the UUID
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:100
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		// not that necessary: op refs the frame
		it.frame = it.state.data

//line ./op-grammar.rl:75
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st10
	tr23:
//line ./op-grammar.rl:100
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		// not that necessary: op refs the frame
		it.frame = it.state.data

//line ./op-grammar.rl:75
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st10
	tr52:
//line ././uuid-grammar.rl:38

//line ./op-grammar.rl:26

		// OK, save the UUID
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:100
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		// not that necessary: op refs the frame
		it.frame = it.state.data

//line ./op-grammar.rl:75
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st10
	tr62:
//line ././uuid-grammar.rl:35

//line ./op-grammar.rl:26

		// OK, save the UUID
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:100
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		// not that necessary: op refs the frame
		it.frame = it.state.data

//line ./op-grammar.rl:75
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st10
	tr70:
//line ././uuid-grammar.rl:35

//line ././uuid-grammar.rl:47
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:26

		// OK, save the UUID
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:100
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		// not that necessary: op refs the frame
		it.frame = it.state.data

//line ./op-grammar.rl:75
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st10
	tr82:
//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st10
	tr90:
//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st10
	tr98:
//line ./op-grammar.rl:52
		i[1] |= ATOM_INT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st10
	tr109:
//line ././uuid-grammar.rl:38

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st10
	tr118:
//line ./op-grammar.rl:56

//line ./op-grammar.rl:58
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st10
	tr126:
//line ./op-grammar.rl:58
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st10
	tr137:
//line ././uuid-grammar.rl:35

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st10
	tr146:
//line ././uuid-grammar.rl:35

//line ././uuid-grammar.rl:47
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st10
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
//line dfa.go:3147
		switch it.state.data[p] {
		case 32:
			goto st10
		case 43:
			goto st11
		case 45:
			goto st11
		}
		switch {
		case it.state.data[p] > 13:
			if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
				goto st12
			}
		case it.state.data[p] >= 9:
			goto st10
		}
		goto st0
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
			goto st12
		}
		goto st0
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		if it.state.data[p] == 46 {
			goto st13
		}
		if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
			goto st12
		}
		goto st0
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
			goto st29
		}
		goto st0
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
		switch it.state.data[p] {
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
		case it.state.data[p] > 13:
			if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
				goto st29
			}
		case it.state.data[p] >= 9:
			goto tr110
		}
		goto st0
	tr117:
//line ./op-grammar.rl:56

		goto st14
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
//line dfa.go:3253
		switch it.state.data[p] {
		case 43:
			goto st15
		case 45:
			goto st15
		}
		if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
			goto st30
		}
		goto st0
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
		if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
			goto st30
		}
		goto st0
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
		switch it.state.data[p] {
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
		case it.state.data[p] > 13:
			if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
				goto st30
			}
		case it.state.data[p] >= 9:
			goto tr119
		}
		goto st0
	tr36:
//line ././uuid-grammar.rl:5
		half = 0

//line ././uuid-grammar.rl:27

//line ././uuid-grammar.rl:10
		digit = prefixSep2Bits(it.state.data[p]) + 4
		i[half] &= INT60_FLAGS | PREFIX_MASKS[digit]

		goto st31
	tr134:
//line ././uuid-grammar.rl:15
		i[half] |= uint64(ABC[it.state.data[p]]) << DIGIT_OFFSETS[digit]
		digit++
		if digit > 10 {
			{
				p++
				it.state.cs = 31
				goto _out
			}
		}

		goto st31
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
//line dfa.go:3346
		switch it.state.data[p] {
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
		case it.state.data[p] < 43:
			switch {
			case it.state.data[p] < 36:
				if 9 <= it.state.data[p] && it.state.data[p] <= 13 {
					goto tr127
				}
			case it.state.data[p] > 37:
				if 40 <= it.state.data[p] && it.state.data[p] <= 41 {
					goto tr132
				}
			default:
				goto tr130
			}
		case it.state.data[p] > 45:
			switch {
			case it.state.data[p] < 65:
				if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
					goto tr134
				}
			case it.state.data[p] > 90:
				if 97 <= it.state.data[p] && it.state.data[p] <= 122 {
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
		half = 0

//line ././uuid-grammar.rl:27

//line ././uuid-grammar.rl:23
		i[half] &= INT60_FLAGS

//line ././uuid-grammar.rl:15
		i[half] |= uint64(ABC[it.state.data[p]]) << DIGIT_OFFSETS[digit]
		digit++
		if digit > 10 {
			{
				p++
				it.state.cs = 32
				goto _out
			}
		}

		goto st32
	tr143:
//line ././uuid-grammar.rl:15
		i[half] |= uint64(ABC[it.state.data[p]]) << DIGIT_OFFSETS[digit]
		digit++
		if digit > 10 {
			{
				p++
				it.state.cs = 32
				goto _out
			}
		}

		goto st32
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
//line dfa.go:3456
		switch it.state.data[p] {
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
		case it.state.data[p] < 43:
			switch {
			case it.state.data[p] < 36:
				if 9 <= it.state.data[p] && it.state.data[p] <= 13 {
					goto tr138
				}
			case it.state.data[p] > 37:
				if 40 <= it.state.data[p] && it.state.data[p] <= 41 {
					goto tr132
				}
			default:
				goto tr130
			}
		case it.state.data[p] > 45:
			switch {
			case it.state.data[p] < 65:
				if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
					goto tr143
				}
			case it.state.data[p] > 90:
				if 97 <= it.state.data[p] && it.state.data[p] <= 122 {
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
//line ./op-grammar.rl:62
		i[0] = uint64(p)

		goto st16
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
//line dfa.go:3541
		switch it.state.data[p] {
		case 10:
			goto st0
		case 13:
			goto st0
		}
		goto st6
	tr12:
//line ././uuid-grammar.rl:30
		digit = 0
		half |= 1

//line ././uuid-grammar.rl:10
		digit = prefixSep2Bits(it.state.data[p]) + 4
		i[half] &= INT60_FLAGS | PREFIX_MASKS[digit]

		goto st17
	tr13:
//line ././uuid-grammar.rl:30
		digit = 0
		half |= 1

//line ././uuid-grammar.rl:23
		i[half] &= INT60_FLAGS

//line ././uuid-grammar.rl:15
		i[half] |= uint64(ABC[it.state.data[p]]) << DIGIT_OFFSETS[digit]
		digit++
		if digit > 10 {
			{
				p++
				it.state.cs = 17
				goto _out
			}
		}

		goto st17
	tr49:
//line ././uuid-grammar.rl:15
		i[half] |= uint64(ABC[it.state.data[p]]) << DIGIT_OFFSETS[digit]
		digit++
		if digit > 10 {
			{
				p++
				it.state.cs = 17
				goto _out
			}
		}

		goto st17
	tr58:
//line ././uuid-grammar.rl:35

//line ././uuid-grammar.rl:30
		digit = 0
		half |= 1

//line ././uuid-grammar.rl:10
		digit = prefixSep2Bits(it.state.data[p]) + 4
		i[half] &= INT60_FLAGS | PREFIX_MASKS[digit]

		goto st17
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
//line dfa.go:3610
		switch it.state.data[p] {
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
		case it.state.data[p] < 48:
			if 9 <= it.state.data[p] && it.state.data[p] <= 13 {
				goto tr45
			}
		case it.state.data[p] > 57:
			switch {
			case it.state.data[p] > 90:
				if 97 <= it.state.data[p] && it.state.data[p] <= 122 {
					goto tr49
				}
			case it.state.data[p] >= 65:
				goto tr49
			}
		default:
			goto tr49
		}
		goto st0
	tr5:
//line ././uuid-grammar.rl:5
		half = 0

//line ././uuid-grammar.rl:27

//line ././uuid-grammar.rl:10
		digit = prefixSep2Bits(it.state.data[p]) + 4
		i[half] &= INT60_FLAGS | PREFIX_MASKS[digit]

		goto st18
	tr59:
//line ././uuid-grammar.rl:15
		i[half] |= uint64(ABC[it.state.data[p]]) << DIGIT_OFFSETS[digit]
		digit++
		if digit > 10 {
			{
				p++
				it.state.cs = 18
				goto _out
			}
		}

		goto st18
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
//line dfa.go:3690
		switch it.state.data[p] {
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
		case it.state.data[p] < 43:
			switch {
			case it.state.data[p] < 36:
				if 9 <= it.state.data[p] && it.state.data[p] <= 13 {
					goto tr53
				}
			case it.state.data[p] > 37:
				if 40 <= it.state.data[p] && it.state.data[p] <= 41 {
					goto tr58
				}
			default:
				goto tr56
			}
		case it.state.data[p] > 45:
			switch {
			case it.state.data[p] < 65:
				if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
					goto tr59
				}
			case it.state.data[p] > 90:
				if 97 <= it.state.data[p] && it.state.data[p] <= 122 {
					goto tr59
				}
			default:
				goto tr59
			}
		default:
			goto tr56
		}
		goto st0
	tr6:
//line ././uuid-grammar.rl:5
		half = 0

//line ././uuid-grammar.rl:27

//line ././uuid-grammar.rl:23
		i[half] &= INT60_FLAGS

//line ././uuid-grammar.rl:15
		i[half] |= uint64(ABC[it.state.data[p]]) << DIGIT_OFFSETS[digit]
		digit++
		if digit > 10 {
			{
				p++
				it.state.cs = 19
				goto _out
			}
		}

		goto st19
	tr67:
//line ././uuid-grammar.rl:15
		i[half] |= uint64(ABC[it.state.data[p]]) << DIGIT_OFFSETS[digit]
		digit++
		if digit > 10 {
			{
				p++
				it.state.cs = 19
				goto _out
			}
		}

		goto st19
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
//line dfa.go:3798
		switch it.state.data[p] {
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
		case it.state.data[p] < 43:
			switch {
			case it.state.data[p] < 36:
				if 9 <= it.state.data[p] && it.state.data[p] <= 13 {
					goto tr63
				}
			case it.state.data[p] > 37:
				if 40 <= it.state.data[p] && it.state.data[p] <= 41 {
					goto tr58
				}
			default:
				goto tr56
			}
		case it.state.data[p] > 45:
			switch {
			case it.state.data[p] < 65:
				if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
					goto tr67
				}
			case it.state.data[p] > 90:
				if 97 <= it.state.data[p] && it.state.data[p] <= 122 {
					goto tr67
				}
			default:
				goto tr67
			}
		default:
			goto tr56
		}
		goto st0
	tr7:
//line ./op-grammar.rl:6
		if idx != 0 {
			it.uuids[idx] = it.uuids[idx-1]
		}

		goto st20
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
//line dfa.go:3883
		switch it.state.data[p] {
		case 43:
			goto tr4
		case 45:
			goto tr4
		case 91:
			goto tr5
		case 93:
			goto tr5
		case 95:
			goto tr6
		case 123:
			goto tr5
		case 125:
			goto tr5
		case 126:
			goto tr6
		}
		switch {
		case it.state.data[p] < 48:
			switch {
			case it.state.data[p] > 37:
				if 40 <= it.state.data[p] && it.state.data[p] <= 41 {
					goto tr5
				}
			case it.state.data[p] >= 36:
				goto tr4
			}
		case it.state.data[p] > 57:
			switch {
			case it.state.data[p] > 90:
				if 97 <= it.state.data[p] && it.state.data[p] <= 122 {
					goto tr6
				}
			case it.state.data[p] >= 65:
				goto tr6
			}
		default:
			goto tr6
		}
		goto st0
	st_out:
	_test_eof21:
		it.state.cs = 21
		goto _test_eof
	_test_eof1:
		it.state.cs = 1
		goto _test_eof
	_test_eof2:
		it.state.cs = 2
		goto _test_eof
	_test_eof3:
		it.state.cs = 3
		goto _test_eof
	_test_eof4:
		it.state.cs = 4
		goto _test_eof
	_test_eof22:
		it.state.cs = 22
		goto _test_eof
	_test_eof23:
		it.state.cs = 23
		goto _test_eof
	_test_eof5:
		it.state.cs = 5
		goto _test_eof
	_test_eof6:
		it.state.cs = 6
		goto _test_eof
	_test_eof24:
		it.state.cs = 24
		goto _test_eof
	_test_eof25:
		it.state.cs = 25
		goto _test_eof
	_test_eof7:
		it.state.cs = 7
		goto _test_eof
	_test_eof8:
		it.state.cs = 8
		goto _test_eof
	_test_eof26:
		it.state.cs = 26
		goto _test_eof
	_test_eof9:
		it.state.cs = 9
		goto _test_eof
	_test_eof27:
		it.state.cs = 27
		goto _test_eof
	_test_eof28:
		it.state.cs = 28
		goto _test_eof
	_test_eof10:
		it.state.cs = 10
		goto _test_eof
	_test_eof11:
		it.state.cs = 11
		goto _test_eof
	_test_eof12:
		it.state.cs = 12
		goto _test_eof
	_test_eof13:
		it.state.cs = 13
		goto _test_eof
	_test_eof29:
		it.state.cs = 29
		goto _test_eof
	_test_eof14:
		it.state.cs = 14
		goto _test_eof
	_test_eof15:
		it.state.cs = 15
		goto _test_eof
	_test_eof30:
		it.state.cs = 30
		goto _test_eof
	_test_eof31:
		it.state.cs = 31
		goto _test_eof
	_test_eof32:
		it.state.cs = 32
		goto _test_eof
	_test_eof16:
		it.state.cs = 16
		goto _test_eof
	_test_eof17:
		it.state.cs = 17
		goto _test_eof
	_test_eof18:
		it.state.cs = 18
		goto _test_eof
	_test_eof19:
		it.state.cs = 19
		goto _test_eof
	_test_eof20:
		it.state.cs = 20
		goto _test_eof

	_test_eof:
		{
		}
		if p == eof {
			switch it.state.cs {
			case 22:
//line ./op-grammar.rl:96
				had_end = true

			case 25:
//line ./op-grammar.rl:78
//line ./op-grammar.rl:96
				had_end = true

			case 24, 27:
//line ./op-grammar.rl:36
				// TODO max size for int/float/string
				it.AddAtom(i)

//line ./op-grammar.rl:78

//line ./op-grammar.rl:96
				had_end = true

			case 31:
//line ././uuid-grammar.rl:35

//line ./op-grammar.rl:36

				// TODO max size for int/float/string
				it.AddAtom(i)

//line ./op-grammar.rl:78

//line ./op-grammar.rl:96
				had_end = true

			case 28:
//line ././uuid-grammar.rl:38

//line ./op-grammar.rl:36

				// TODO max size for int/float/string
				it.AddAtom(i)

//line ./op-grammar.rl:78

//line ./op-grammar.rl:96
				had_end = true

			case 26:
//line ./op-grammar.rl:52
				i[1] |= ATOM_INT_62

//line ./op-grammar.rl:36
				// TODO max size for int/float/string
				it.AddAtom(i)

//line ./op-grammar.rl:78

//line ./op-grammar.rl:96
				had_end = true

			case 30:
//line ./op-grammar.rl:58
				i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
				// TODO max size for int/float/string
				it.AddAtom(i)

//line ./op-grammar.rl:78

//line ./op-grammar.rl:96
				had_end = true

			case 32:
//line ././uuid-grammar.rl:35

//line ././uuid-grammar.rl:47
				i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:36

				// TODO max size for int/float/string
				it.AddAtom(i)

//line ./op-grammar.rl:78

//line ./op-grammar.rl:96
				had_end = true

			case 29:
//line ./op-grammar.rl:56
//line ./op-grammar.rl:58
				i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
				// TODO max size for int/float/string
				it.AddAtom(i)

//line ./op-grammar.rl:78

//line ./op-grammar.rl:96
				had_end = true

//line dfa.go:4094
			}
		}

	_out:
		{
		}
	}

//line dfa.rl:47

	it.state.incomplete = i
	it.state.idx = idx
	it.state.digit = digit
	it.state.half = half
	it.state.p = p

	if !it.state.streaming && it.state.cs < RON_first_final {
		it.state.cs = RON_error
	}

}

func (ctx_uuid UUID) Parse(data []byte) (UUID, error) {

//line dfa.rl:65
//line dfa.go:4122
	const UUID_start int = 1
	const UUID_first_final int = 2
	const UUID_error int = 0

	const UUID_en_main int = 1

//line dfa.rl:66
	var i uint128 = ctx_uuid.uint128
	digit := uint(0)
	half := 0

	cs, p, pe, eof := 0, 0, len(data), len(data)
	//var ts, te, act int
	_ = eof
	//_,_,_ = ts,te,act

//line dfa.go:4143
	{
		cs = UUID_start
	}

//line dfa.go:4148
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
		half = 0

//line ./uuid-grammar.rl:41
		half |= 1
		i[half] &= INT60_FULL
		i[half] |= uint64(uuidSep2Bits(data[p])) << 60

		goto st2
	tr7:
//line ./uuid-grammar.rl:35
//line ./uuid-grammar.rl:41
		half |= 1
		i[half] &= INT60_FULL
		i[half] |= uint64(uuidSep2Bits(data[p])) << 60

		goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
//line dfa.go:4242
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
//line ./uuid-grammar.rl:30
		digit = 0
		half |= 1

//line ./uuid-grammar.rl:10
		digit = prefixSep2Bits(data[p]) + 4
		i[half] &= INT60_FLAGS | PREFIX_MASKS[digit]

		goto st3
	tr5:
//line ./uuid-grammar.rl:30
		digit = 0
		half |= 1

//line ./uuid-grammar.rl:23
		i[half] &= INT60_FLAGS

//line ./uuid-grammar.rl:15
		i[half] |= uint64(ABC[data[p]]) << DIGIT_OFFSETS[digit]
		digit++
		if digit > 10 {
			{
				p++
				cs = 3
				goto _out
			}
		}

		goto st3
	tr6:
//line ./uuid-grammar.rl:15
		i[half] |= uint64(ABC[data[p]]) << DIGIT_OFFSETS[digit]
		digit++
		if digit > 10 {
			{
				p++
				cs = 3
				goto _out
			}
		}

		goto st3
	tr8:
//line ./uuid-grammar.rl:35

//line ./uuid-grammar.rl:30
		digit = 0
		half |= 1

//line ./uuid-grammar.rl:10
		digit = prefixSep2Bits(data[p]) + 4
		i[half] &= INT60_FLAGS | PREFIX_MASKS[digit]

		goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
//line dfa.go:4336
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
		half = 0

//line ./uuid-grammar.rl:27

//line ./uuid-grammar.rl:10
		digit = prefixSep2Bits(data[p]) + 4
		i[half] &= INT60_FLAGS | PREFIX_MASKS[digit]

		goto st4
	tr9:
//line ./uuid-grammar.rl:15
		i[half] |= uint64(ABC[data[p]]) << DIGIT_OFFSETS[digit]
		digit++
		if digit > 10 {
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
//line dfa.go:4385
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
		half = 0

//line ./uuid-grammar.rl:27

//line ./uuid-grammar.rl:23
		i[half] &= INT60_FLAGS

//line ./uuid-grammar.rl:15
		i[half] |= uint64(ABC[data[p]]) << DIGIT_OFFSETS[digit]
		digit++
		if digit > 10 {
			{
				p++
				cs = 5
				goto _out
			}
		}

		goto st5
	tr10:
//line ./uuid-grammar.rl:15
		i[half] |= uint64(ABC[data[p]]) << DIGIT_OFFSETS[digit]
		digit++
		if digit > 10 {
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
//line dfa.go:4463
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
//line ./uuid-grammar.rl:35

			case 3:
//line ./uuid-grammar.rl:38

			case 5:
//line ./uuid-grammar.rl:35
//line ./uuid-grammar.rl:47
				i[1] = UUID_NAME_UPPER_BITS

//line dfa.go:4530
			}
		}

	_out:
		{
		}
	}

//line dfa.rl:84

	if cs < UUID_first_final || digit > 10 {
		return ERROR_UUID, errors.New(fmt.Sprintf("parse error at pos %d", p))
	} else {
		return UUID{uint128: i}, nil
	}

}
