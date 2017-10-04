//line dfa.rl:1
package RON

import "fmt"
import "errors"

const trace = false

/*
 * [ ] use parserState inside XParse
 * [ ] move pS to Iter, XP -> Iter.Next
 * [ ] ron.hpp structure
 * [ ] Cursor, separate atom-grammar.rl
 * [ ] Cursor.Integer()... (same as uuid-..., ragel prepares a slice)
 *
 * */

type OpParserPos struct {
	// int60 idx, base64 digit idx
	idx, half, digit uint
}

type OpParserState struct {
	OpParserPos
	// the RON frame (for the streaming mode, probably a bit less or a bit more)
	data []byte
	// parser position
	p int
	// ragel state
	cs int
	// ts, te, act int
	// incomplete uuid/atom data
	incomplete uint128
	// streaming mode switch
	streaming bool
}

/*func (state *OpParserState) atom_slice () (ret []byte) {
    if len(cur_atom) > 0 {
        return append(cur_atom, data[:p])
    } else {
        return data[atom_start:p]
    }
}*/

const (
	PARSED_ERROR = iota
	PARSED_OP
	PARSED_INCOMPLETE
	PARSED_EOF
)

// Parse consumes one op, unless the buffer ends earlier.
func (it *Iterator) Parse() int {

	if it.IsLast() {
		it.Op = ZERO_OP
		return PARSED_EOF
	}

//line dfa.rl:60
//line dfa.go:65
	const RON_start int = 1
	const RON_first_final int = 18
	const RON_error int = 0

	const RON_en_main int = 1

//line dfa.rl:61

//line dfa.rl:62
	if it.state.cs == 0 {
		it.Reset()
		it.frame = it.state.data
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
	}

	p, pe, eof := it.state.p, len(it.state.data), len(it.state.data)
	n := uint(0)
	done := false
	_ = done
	_ = eof

	if it.state.streaming {
		eof = -1
	}

	i := it.state.incomplete
	idx := it.state.idx
	half := it.state.half
	digit := it.state.digit

//line dfa.go:101
	{
		it.state.cs = RON_start
	}

//line dfa.go:106
	{
		if p == pe {
			goto _test_eof
		}
		switch it.state.cs {
		case 1:
			goto st_case_1
		case 0:
			goto st_case_0
		case 18:
			goto st_case_18
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		case 4:
			goto st_case_4
		case 5:
			goto st_case_5
		case 19:
			goto st_case_19
		case 20:
			goto st_case_20
		case 21:
			goto st_case_21
		case 6:
			goto st_case_6
		case 7:
			goto st_case_7
		case 22:
			goto st_case_22
		case 8:
			goto st_case_8
		case 23:
			goto st_case_23
		case 24:
			goto st_case_24
		case 9:
			goto st_case_9
		case 10:
			goto st_case_10
		case 25:
			goto st_case_25
		case 26:
			goto st_case_26
		case 27:
			goto st_case_27
		case 11:
			goto st_case_11
		case 12:
			goto st_case_12
		case 28:
			goto st_case_28
		case 29:
			goto st_case_29
		case 30:
			goto st_case_30
		case 13:
			goto st_case_13
		case 14:
			goto st_case_14
		case 15:
			goto st_case_15
		case 16:
			goto st_case_16
		case 17:
			goto st_case_17
		}
		goto st_out
	tr13:
//line ./op-grammar.rl:20

		goto st1
	tr41:
//line ././uuid-grammar.rl:38

//line ./op-grammar.rl:20

		goto st1
	tr49:
//line ././uuid-grammar.rl:35

//line ./op-grammar.rl:20

		goto st1
	tr59:
//line ././uuid-grammar.rl:35

//line ././uuid-grammar.rl:47
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:20

		goto st1
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
//line dfa.go:214
		switch it.state.data[p] {
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
		if 9 <= it.state.data[p] && it.state.data[p] <= 13 {
			goto st1
		}
		goto st0
	st_case_0:
	st0:
		it.state.cs = 0
		goto _out
	tr2:
//line ./op-grammar.rl:71
		it.term = termSep2Bits(it.state.data[p])

		goto st18
	tr14:
//line ./op-grammar.rl:20
//line ./op-grammar.rl:71
		it.term = termSep2Bits(it.state.data[p])

		goto st18
	tr42:
//line ././uuid-grammar.rl:38

//line ./op-grammar.rl:20

//line ./op-grammar.rl:71
		it.term = termSep2Bits(it.state.data[p])

		goto st18
	tr50:
//line ././uuid-grammar.rl:35

//line ./op-grammar.rl:20

//line ./op-grammar.rl:71
		it.term = termSep2Bits(it.state.data[p])

		goto st18
	tr60:
//line ././uuid-grammar.rl:35

//line ././uuid-grammar.rl:47
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:20

//line ./op-grammar.rl:71
		it.term = termSep2Bits(it.state.data[p])

		goto st18
	tr67:
//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 18
			goto _out
		}

		goto st18
	tr68:
//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 18
			goto _out
		}

//line ./op-grammar.rl:71
		it.term = termSep2Bits(it.state.data[p])

		goto st18
	tr76:
//line ./op-grammar.rl:27
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:68

//line ./op-grammar.rl:71
		it.term = termSep2Bits(it.state.data[p])

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 18
			goto _out
		}

		goto st18
	tr84:
//line ./op-grammar.rl:68
//line ./op-grammar.rl:71
		it.term = termSep2Bits(it.state.data[p])

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 18
			goto _out
		}

		goto st18
	tr93:
//line ./op-grammar.rl:43
		i[1] |= ATOM_INT_62

//line ./op-grammar.rl:27
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:68

//line ./op-grammar.rl:71
		it.term = termSep2Bits(it.state.data[p])

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 18
			goto _out
		}

		goto st18
	tr103:
//line ././uuid-grammar.rl:38

//line ./op-grammar.rl:27

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:68

//line ./op-grammar.rl:71
		it.term = termSep2Bits(it.state.data[p])

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 18
			goto _out
		}

		goto st18
	tr112:
//line ./op-grammar.rl:47
//line ./op-grammar.rl:49
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:27
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:68

//line ./op-grammar.rl:71
		it.term = termSep2Bits(it.state.data[p])

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 18
			goto _out
		}

		goto st18
	tr123:
//line ./op-grammar.rl:49
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:27
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:68

//line ./op-grammar.rl:71
		it.term = termSep2Bits(it.state.data[p])

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 18
			goto _out
		}

		goto st18
	tr131:
//line ././uuid-grammar.rl:35

//line ./op-grammar.rl:27

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:68

//line ./op-grammar.rl:71
		it.term = termSep2Bits(it.state.data[p])

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 18
			goto _out
		}

		goto st18
	tr142:
//line ././uuid-grammar.rl:35

//line ././uuid-grammar.rl:47
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:27

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:68

//line ./op-grammar.rl:71
		it.term = termSep2Bits(it.state.data[p])

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 18
			goto _out
		}

		goto st18
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
//line dfa.go:504
		switch it.state.data[p] {
		case 32:
			goto tr67
		case 33:
			goto tr68
		case 35:
			goto tr69
		case 39:
			goto tr70
		case 42:
			goto tr69
		case 44:
			goto tr68
		case 46:
			goto tr71
		case 58:
			goto tr69
		case 59:
			goto tr68
		case 61:
			goto tr72
		case 62:
			goto tr73
		case 63:
			goto tr68
		case 64:
			goto tr69
		case 94:
			goto tr74
		}
		if 9 <= it.state.data[p] && it.state.data[p] <= 13 {
			goto tr67
		}
		goto st0
	tr3:
//line ./op-grammar.rl:12
		n = specSep2Bits(it.state.data[p]) << 1
		if n <= idx {
			{
				p++
				it.state.cs = 2
				goto _out
			}
		}
		idx = n

		goto st2
	tr15:
//line ./op-grammar.rl:20
//line ./op-grammar.rl:12
		n = specSep2Bits(it.state.data[p]) << 1
		if n <= idx {
			{
				p++
				it.state.cs = 2
				goto _out
			}
		}
		idx = n

		goto st2
	tr43:
//line ././uuid-grammar.rl:38

//line ./op-grammar.rl:20

//line ./op-grammar.rl:12
		n = specSep2Bits(it.state.data[p]) << 1
		if n <= idx {
			{
				p++
				it.state.cs = 2
				goto _out
			}
		}
		idx = n

		goto st2
	tr51:
//line ././uuid-grammar.rl:35

//line ./op-grammar.rl:20

//line ./op-grammar.rl:12
		n = specSep2Bits(it.state.data[p]) << 1
		if n <= idx {
			{
				p++
				it.state.cs = 2
				goto _out
			}
		}
		idx = n

		goto st2
	tr61:
//line ././uuid-grammar.rl:35

//line ././uuid-grammar.rl:47
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:20

//line ./op-grammar.rl:12
		n = specSep2Bits(it.state.data[p]) << 1
		if n <= idx {
			{
				p++
				it.state.cs = 2
				goto _out
			}
		}
		idx = n

		goto st2
	tr69:
//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 2
			goto _out
		}

//line ./op-grammar.rl:12
		n = specSep2Bits(it.state.data[p]) << 1
		if n <= idx {
			{
				p++
				it.state.cs = 2
				goto _out
			}
		}
		idx = n

		goto st2
	tr77:
//line ./op-grammar.rl:27

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 2
			goto _out
		}

//line ./op-grammar.rl:12
		n = specSep2Bits(it.state.data[p]) << 1
		if n <= idx {
			{
				p++
				it.state.cs = 2
				goto _out
			}
		}
		idx = n

		goto st2
	tr85:
//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 2
			goto _out
		}

//line ./op-grammar.rl:12
		n = specSep2Bits(it.state.data[p]) << 1
		if n <= idx {
			{
				p++
				it.state.cs = 2
				goto _out
			}
		}
		idx = n

		goto st2
	tr94:
//line ./op-grammar.rl:43
		i[1] |= ATOM_INT_62

//line ./op-grammar.rl:27
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 2
			goto _out
		}

//line ./op-grammar.rl:12
		n = specSep2Bits(it.state.data[p]) << 1
		if n <= idx {
			{
				p++
				it.state.cs = 2
				goto _out
			}
		}
		idx = n

		goto st2
	tr104:
//line ././uuid-grammar.rl:38

//line ./op-grammar.rl:27

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 2
			goto _out
		}

//line ./op-grammar.rl:12
		n = specSep2Bits(it.state.data[p]) << 1
		if n <= idx {
			{
				p++
				it.state.cs = 2
				goto _out
			}
		}
		idx = n

		goto st2
	tr113:
//line ./op-grammar.rl:47

//line ./op-grammar.rl:49
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:27
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 2
			goto _out
		}

//line ./op-grammar.rl:12
		n = specSep2Bits(it.state.data[p]) << 1
		if n <= idx {
			{
				p++
				it.state.cs = 2
				goto _out
			}
		}
		idx = n

		goto st2
	tr124:
//line ./op-grammar.rl:49
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:27
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 2
			goto _out
		}

//line ./op-grammar.rl:12
		n = specSep2Bits(it.state.data[p]) << 1
		if n <= idx {
			{
				p++
				it.state.cs = 2
				goto _out
			}
		}
		idx = n

		goto st2
	tr132:
//line ././uuid-grammar.rl:35

//line ./op-grammar.rl:27

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 2
			goto _out
		}

//line ./op-grammar.rl:12
		n = specSep2Bits(it.state.data[p]) << 1
		if n <= idx {
			{
				p++
				it.state.cs = 2
				goto _out
			}
		}
		idx = n

		goto st2
	tr143:
//line ././uuid-grammar.rl:35

//line ././uuid-grammar.rl:47
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:27

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 2
			goto _out
		}

//line ./op-grammar.rl:12
		n = specSep2Bits(it.state.data[p]) << 1
		if n <= idx {
			{
				p++
				it.state.cs = 2
				goto _out
			}
		}
		idx = n

		goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
//line dfa.go:841
		switch it.state.data[p] {
		case 32:
			goto st2
		case 43:
			goto tr9
		case 45:
			goto tr9
		case 91:
			goto tr10
		case 93:
			goto tr10
		case 96:
			goto tr12
		case 123:
			goto tr10
		case 125:
			goto tr10
		case 126:
			goto tr11
		}
		switch {
		case it.state.data[p] < 40:
			switch {
			case it.state.data[p] > 13:
				if 36 <= it.state.data[p] && it.state.data[p] <= 37 {
					goto tr9
				}
			case it.state.data[p] >= 9:
				goto st2
			}
		case it.state.data[p] > 41:
			switch {
			case it.state.data[p] < 65:
				if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
					goto tr11
				}
			case it.state.data[p] > 90:
				if 95 <= it.state.data[p] && it.state.data[p] <= 122 {
					goto tr11
				}
			default:
				goto tr11
			}
		default:
			goto tr10
		}
		goto st0
	tr9:
//line ././uuid-grammar.rl:5
		half = 0

//line ././uuid-grammar.rl:41
		half |= 1
		i[half] &= INT60_FULL
		i[half] |= uint64(uuidSep2Bits(it.state.data[p])) << 60

		goto st3
	tr52:
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
//line dfa.go:917
		switch it.state.data[p] {
		case 32:
			goto tr13
		case 33:
			goto tr14
		case 35:
			goto tr15
		case 39:
			goto tr16
		case 42:
			goto tr15
		case 44:
			goto tr14
		case 58:
			goto tr15
		case 59:
			goto tr14
		case 61:
			goto tr19
		case 62:
			goto tr20
		case 63:
			goto tr14
		case 64:
			goto tr15
		case 91:
			goto tr17
		case 93:
			goto tr17
		case 94:
			goto tr21
		case 95:
			goto tr18
		case 123:
			goto tr17
		case 125:
			goto tr17
		case 126:
			goto tr18
		}
		switch {
		case it.state.data[p] < 48:
			switch {
			case it.state.data[p] > 13:
				if 40 <= it.state.data[p] && it.state.data[p] <= 41 {
					goto tr17
				}
			case it.state.data[p] >= 9:
				goto tr13
			}
		case it.state.data[p] > 57:
			switch {
			case it.state.data[p] > 90:
				if 97 <= it.state.data[p] && it.state.data[p] <= 122 {
					goto tr18
				}
			case it.state.data[p] >= 65:
				goto tr18
			}
		default:
			goto tr18
		}
		goto st0
	tr4:
//line ./op-grammar.rl:66

//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

		goto st4
	tr16:
//line ./op-grammar.rl:20
//line ./op-grammar.rl:66

//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

		goto st4
	tr44:
//line ././uuid-grammar.rl:38

//line ./op-grammar.rl:20

//line ./op-grammar.rl:66

//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

		goto st4
	tr53:
//line ././uuid-grammar.rl:35

//line ./op-grammar.rl:20

//line ./op-grammar.rl:66

//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

		goto st4
	tr62:
//line ././uuid-grammar.rl:35

//line ././uuid-grammar.rl:47
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:20

//line ./op-grammar.rl:66

//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

		goto st4
	tr70:
//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 4
			goto _out
		}

//line ./op-grammar.rl:66
//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

		goto st4
	tr78:
//line ./op-grammar.rl:27
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 4
			goto _out
		}

//line ./op-grammar.rl:66

		goto st4
	tr86:
//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 4
			goto _out
		}

//line ./op-grammar.rl:66

		goto st4
	tr95:
//line ./op-grammar.rl:43
		i[1] |= ATOM_INT_62

//line ./op-grammar.rl:27
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 4
			goto _out
		}

//line ./op-grammar.rl:66

		goto st4
	tr105:
//line ././uuid-grammar.rl:38

//line ./op-grammar.rl:27

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 4
			goto _out
		}

//line ./op-grammar.rl:66

		goto st4
	tr114:
//line ./op-grammar.rl:47
//line ./op-grammar.rl:49
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:27
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 4
			goto _out
		}

//line ./op-grammar.rl:66

		goto st4
	tr125:
//line ./op-grammar.rl:49
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:27
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 4
			goto _out
		}

//line ./op-grammar.rl:66

		goto st4
	tr134:
//line ././uuid-grammar.rl:35

//line ./op-grammar.rl:27

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 4
			goto _out
		}

//line ./op-grammar.rl:66

		goto st4
	tr144:
//line ././uuid-grammar.rl:35

//line ././uuid-grammar.rl:47
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:27

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 4
			goto _out
		}

//line ./op-grammar.rl:66

		goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
//line dfa.go:1283
		switch it.state.data[p] {
		case 10:
			goto st0
		case 13:
			goto st0
		case 39:
			goto tr23
		case 92:
			goto tr24
		}
		goto tr22
	tr22:
//line ./op-grammar.rl:53
		i[0] = uint64(p)

		goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
//line dfa.go:1306
		switch it.state.data[p] {
		case 10:
			goto st0
		case 13:
			goto st0
		case 39:
			goto tr26
		case 92:
			goto st13
		}
		goto st5
	tr23:
//line ./op-grammar.rl:53
		i[0] = uint64(p)

//line ./op-grammar.rl:56
		i[1] = uint64(p) | ATOM_STRING_62

		goto st19
	tr26:
//line ./op-grammar.rl:56
		i[1] = uint64(p) | ATOM_STRING_62

		goto st19
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
//line dfa.go:1339
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
//line ./op-grammar.rl:27

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 20
			goto _out
		}

		goto st20
	tr83:
//line ./op-grammar.rl:68
//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 20
			goto _out
		}

		goto st20
	tr92:
//line ./op-grammar.rl:43
		i[1] |= ATOM_INT_62

//line ./op-grammar.rl:27
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 20
			goto _out
		}

		goto st20
	tr102:
//line ././uuid-grammar.rl:38

//line ./op-grammar.rl:27

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 20
			goto _out
		}

		goto st20
	tr111:
//line ./op-grammar.rl:47
//line ./op-grammar.rl:49
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:27
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 20
			goto _out
		}

		goto st20
	tr122:
//line ./op-grammar.rl:49
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:27
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 20
			goto _out
		}

		goto st20
	tr130:
//line ././uuid-grammar.rl:35

//line ./op-grammar.rl:27

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 20
			goto _out
		}

		goto st20
	tr141:
//line ././uuid-grammar.rl:35

//line ././uuid-grammar.rl:47
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:27

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 20
			goto _out
		}

		goto st20
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
//line dfa.go:1522
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
	tr71:
//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 21
			goto _out
		}

		goto st21
	tr79:
//line ./op-grammar.rl:27
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 21
			goto _out
		}

		goto st21
	tr87:
//line ./op-grammar.rl:68
//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 21
			goto _out
		}

		goto st21
	tr96:
//line ./op-grammar.rl:43
		i[1] |= ATOM_INT_62

//line ./op-grammar.rl:27
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 21
			goto _out
		}

		goto st21
	tr106:
//line ././uuid-grammar.rl:38

//line ./op-grammar.rl:27

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 21
			goto _out
		}

		goto st21
	tr121:
//line ./op-grammar.rl:47
//line ./op-grammar.rl:49
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:27
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 21
			goto _out
		}

		goto st21
	tr126:
//line ./op-grammar.rl:49
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:27
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 21
			goto _out
		}

		goto st21
	tr136:
//line ././uuid-grammar.rl:35

//line ./op-grammar.rl:27

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 21
			goto _out
		}

		goto st21
	tr145:
//line ././uuid-grammar.rl:35

//line ././uuid-grammar.rl:47
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:27

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 21
			goto _out
		}

		goto st21
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
//line dfa.go:1712
		switch it.state.data[p] {
		case 32:
			goto st21
		case 46:
			goto st21
		}
		if 9 <= it.state.data[p] && it.state.data[p] <= 13 {
			goto st21
		}
		goto st0
	tr5:
//line ./op-grammar.rl:66

//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

		goto st6
	tr19:
//line ./op-grammar.rl:20
//line ./op-grammar.rl:66

//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

		goto st6
	tr46:
//line ././uuid-grammar.rl:38

//line ./op-grammar.rl:20

//line ./op-grammar.rl:66

//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

		goto st6
	tr56:
//line ././uuid-grammar.rl:35

//line ./op-grammar.rl:20

//line ./op-grammar.rl:66

//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

		goto st6
	tr64:
//line ././uuid-grammar.rl:35

//line ././uuid-grammar.rl:47
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:20

//line ./op-grammar.rl:66

//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

		goto st6
	tr72:
//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 6
			goto _out
		}

//line ./op-grammar.rl:66
//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

		goto st6
	tr80:
//line ./op-grammar.rl:27
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 6
			goto _out
		}

//line ./op-grammar.rl:66

		goto st6
	tr88:
//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 6
			goto _out
		}

//line ./op-grammar.rl:66

		goto st6
	tr97:
//line ./op-grammar.rl:43
		i[1] |= ATOM_INT_62

//line ./op-grammar.rl:27
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 6
			goto _out
		}

//line ./op-grammar.rl:66

		goto st6
	tr108:
//line ././uuid-grammar.rl:38

//line ./op-grammar.rl:27

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 6
			goto _out
		}

//line ./op-grammar.rl:66

		goto st6
	tr116:
//line ./op-grammar.rl:47
//line ./op-grammar.rl:49
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:27
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 6
			goto _out
		}

//line ./op-grammar.rl:66

		goto st6
	tr127:
//line ./op-grammar.rl:49
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:27
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 6
			goto _out
		}

//line ./op-grammar.rl:66

		goto st6
	tr138:
//line ././uuid-grammar.rl:35

//line ./op-grammar.rl:27

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 6
			goto _out
		}

//line ./op-grammar.rl:66

		goto st6
	tr147:
//line ././uuid-grammar.rl:35

//line ././uuid-grammar.rl:47
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:27

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 6
			goto _out
		}

//line ./op-grammar.rl:66

		goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
//line dfa.go:2025
		switch it.state.data[p] {
		case 32:
			goto st6
		case 43:
			goto tr29
		case 45:
			goto tr29
		}
		switch {
		case it.state.data[p] > 13:
			if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
				goto tr30
			}
		case it.state.data[p] >= 9:
			goto st6
		}
		goto st0
	tr29:
//line ./op-grammar.rl:32

//line ./op-grammar.rl:34
		if it.state.data[p] == '-' {
			i[1] |= 1
		}

		goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
//line dfa.go:2059
		if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
			goto tr31
		}
		goto st0
	tr30:
//line ./op-grammar.rl:32

//line ./op-grammar.rl:39
		i[0] *= 10
		i[0] += uint64(int(it.state.data[p]) - int('0'))

		goto st22
	tr31:
//line ./op-grammar.rl:39
		i[0] *= 10
		i[0] += uint64(int(it.state.data[p]) - int('0'))

		goto st22
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
//line dfa.go:2086
		switch it.state.data[p] {
		case 32:
			goto tr92
		case 33:
			goto tr93
		case 35:
			goto tr94
		case 39:
			goto tr95
		case 42:
			goto tr94
		case 44:
			goto tr93
		case 46:
			goto tr96
		case 58:
			goto tr94
		case 59:
			goto tr93
		case 61:
			goto tr97
		case 62:
			goto tr98
		case 63:
			goto tr93
		case 64:
			goto tr94
		case 94:
			goto tr99
		}
		switch {
		case it.state.data[p] > 13:
			if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
				goto tr31
			}
		case it.state.data[p] >= 9:
			goto tr92
		}
		goto st0
	tr6:
//line ./op-grammar.rl:66

//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

		goto st8
	tr20:
//line ./op-grammar.rl:20
//line ./op-grammar.rl:66

//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

		goto st8
	tr47:
//line ././uuid-grammar.rl:38

//line ./op-grammar.rl:20

//line ./op-grammar.rl:66

//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

		goto st8
	tr57:
//line ././uuid-grammar.rl:35

//line ./op-grammar.rl:20

//line ./op-grammar.rl:66

//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

		goto st8
	tr65:
//line ././uuid-grammar.rl:35

//line ././uuid-grammar.rl:47
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:20

//line ./op-grammar.rl:66

//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

		goto st8
	tr73:
//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 8
			goto _out
		}

//line ./op-grammar.rl:66
//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

		goto st8
	tr81:
//line ./op-grammar.rl:27
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 8
			goto _out
		}

//line ./op-grammar.rl:66

		goto st8
	tr89:
//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 8
			goto _out
		}

//line ./op-grammar.rl:66

		goto st8
	tr98:
//line ./op-grammar.rl:43
		i[1] |= ATOM_INT_62

//line ./op-grammar.rl:27
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 8
			goto _out
		}

//line ./op-grammar.rl:66

		goto st8
	tr109:
//line ././uuid-grammar.rl:38

//line ./op-grammar.rl:27

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 8
			goto _out
		}

//line ./op-grammar.rl:66

		goto st8
	tr117:
//line ./op-grammar.rl:47
//line ./op-grammar.rl:49
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:27
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 8
			goto _out
		}

//line ./op-grammar.rl:66

		goto st8
	tr128:
//line ./op-grammar.rl:49
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:27
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 8
			goto _out
		}

//line ./op-grammar.rl:66

		goto st8
	tr139:
//line ././uuid-grammar.rl:35

//line ./op-grammar.rl:27

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 8
			goto _out
		}

//line ./op-grammar.rl:66

		goto st8
	tr148:
//line ././uuid-grammar.rl:35

//line ././uuid-grammar.rl:47
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:27

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 8
			goto _out
		}

//line ./op-grammar.rl:66

		goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
//line dfa.go:2428
		switch it.state.data[p] {
		case 32:
			goto st8
		case 43:
			goto tr33
		case 45:
			goto tr33
		case 91:
			goto tr34
		case 93:
			goto tr34
		case 95:
			goto tr35
		case 123:
			goto tr34
		case 125:
			goto tr34
		case 126:
			goto tr35
		}
		switch {
		case it.state.data[p] < 40:
			switch {
			case it.state.data[p] > 13:
				if 36 <= it.state.data[p] && it.state.data[p] <= 37 {
					goto tr33
				}
			case it.state.data[p] >= 9:
				goto st8
			}
		case it.state.data[p] > 41:
			switch {
			case it.state.data[p] < 65:
				if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
					goto tr35
				}
			case it.state.data[p] > 90:
				if 97 <= it.state.data[p] && it.state.data[p] <= 122 {
					goto tr35
				}
			default:
				goto tr35
			}
		default:
			goto tr34
		}
		goto st0
	tr33:
//line ././uuid-grammar.rl:5
		half = 0

//line ././uuid-grammar.rl:41
		half |= 1
		i[half] &= INT60_FULL
		i[half] |= uint64(uuidSep2Bits(it.state.data[p])) << 60

		goto st23
	tr133:
//line ././uuid-grammar.rl:35
//line ././uuid-grammar.rl:41
		half |= 1
		i[half] &= INT60_FULL
		i[half] |= uint64(uuidSep2Bits(it.state.data[p])) << 60

		goto st23
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
//line dfa.go:2504
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
			goto tr100
		case 93:
			goto tr100
		case 94:
			goto tr82
		case 95:
			goto tr101
		case 123:
			goto tr100
		case 125:
			goto tr100
		case 126:
			goto tr101
		}
		switch {
		case it.state.data[p] < 48:
			switch {
			case it.state.data[p] > 13:
				if 40 <= it.state.data[p] && it.state.data[p] <= 41 {
					goto tr100
				}
			case it.state.data[p] >= 9:
				goto tr75
			}
		case it.state.data[p] > 57:
			switch {
			case it.state.data[p] > 90:
				if 97 <= it.state.data[p] && it.state.data[p] <= 122 {
					goto tr101
				}
			case it.state.data[p] >= 65:
				goto tr101
			}
		default:
			goto tr101
		}
		goto st0
	tr100:
//line ././uuid-grammar.rl:30
		digit = 0
		half |= 1

//line ././uuid-grammar.rl:10
		digit = prefixSep2Bits(it.state.data[p]) + 4
		i[half] &= INT60_FLAGS | PREFIX_MASKS[digit]

		goto st24
	tr101:
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
				it.state.cs = 24
				goto _out
			}
		}

		goto st24
	tr107:
//line ././uuid-grammar.rl:15
		i[half] |= uint64(ABC[it.state.data[p]]) << DIGIT_OFFSETS[digit]
		digit++
		if digit > 10 {
			{
				p++
				it.state.cs = 24
				goto _out
			}
		}

		goto st24
	tr135:
//line ././uuid-grammar.rl:35

//line ././uuid-grammar.rl:30
		digit = 0
		half |= 1

//line ././uuid-grammar.rl:10
		digit = prefixSep2Bits(it.state.data[p]) + 4
		i[half] &= INT60_FLAGS | PREFIX_MASKS[digit]

		goto st24
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
//line dfa.go:2631
		switch it.state.data[p] {
		case 32:
			goto tr102
		case 33:
			goto tr103
		case 35:
			goto tr104
		case 39:
			goto tr105
		case 42:
			goto tr104
		case 44:
			goto tr103
		case 46:
			goto tr106
		case 58:
			goto tr104
		case 59:
			goto tr103
		case 61:
			goto tr108
		case 62:
			goto tr109
		case 63:
			goto tr103
		case 64:
			goto tr104
		case 94:
			goto tr110
		case 95:
			goto tr107
		case 126:
			goto tr107
		}
		switch {
		case it.state.data[p] < 48:
			if 9 <= it.state.data[p] && it.state.data[p] <= 13 {
				goto tr102
			}
		case it.state.data[p] > 57:
			switch {
			case it.state.data[p] > 90:
				if 97 <= it.state.data[p] && it.state.data[p] <= 122 {
					goto tr107
				}
			case it.state.data[p] >= 65:
				goto tr107
			}
		default:
			goto tr107
		}
		goto st0
	tr7:
//line ./op-grammar.rl:66

//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

		goto st9
	tr21:
//line ./op-grammar.rl:20
//line ./op-grammar.rl:66

//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

		goto st9
	tr48:
//line ././uuid-grammar.rl:38

//line ./op-grammar.rl:20

//line ./op-grammar.rl:66

//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

		goto st9
	tr58:
//line ././uuid-grammar.rl:35

//line ./op-grammar.rl:20

//line ./op-grammar.rl:66

//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

		goto st9
	tr66:
//line ././uuid-grammar.rl:35

//line ././uuid-grammar.rl:47
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:20

//line ./op-grammar.rl:66

//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

		goto st9
	tr74:
//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 9
			goto _out
		}

//line ./op-grammar.rl:66
//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

		goto st9
	tr82:
//line ./op-grammar.rl:27
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 9
			goto _out
		}

//line ./op-grammar.rl:66

		goto st9
	tr90:
//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 9
			goto _out
		}

//line ./op-grammar.rl:66

		goto st9
	tr99:
//line ./op-grammar.rl:43
		i[1] |= ATOM_INT_62

//line ./op-grammar.rl:27
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 9
			goto _out
		}

//line ./op-grammar.rl:66

		goto st9
	tr110:
//line ././uuid-grammar.rl:38

//line ./op-grammar.rl:27

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 9
			goto _out
		}

//line ./op-grammar.rl:66

		goto st9
	tr119:
//line ./op-grammar.rl:47
//line ./op-grammar.rl:49
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:27
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 9
			goto _out
		}

//line ./op-grammar.rl:66

		goto st9
	tr129:
//line ./op-grammar.rl:49
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:27
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 9
			goto _out
		}

//line ./op-grammar.rl:66

		goto st9
	tr140:
//line ././uuid-grammar.rl:35

//line ./op-grammar.rl:27

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 9
			goto _out
		}

//line ./op-grammar.rl:66

		goto st9
	tr149:
//line ././uuid-grammar.rl:35

//line ././uuid-grammar.rl:47
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:27

		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:23
		digit = 0
		i = uint128{0, 0}

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 9
			goto _out
		}

//line ./op-grammar.rl:66

		goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
//line dfa.go:2986
		switch it.state.data[p] {
		case 32:
			goto st9
		case 43:
			goto st10
		case 45:
			goto st10
		}
		switch {
		case it.state.data[p] > 13:
			if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
				goto st25
			}
		case it.state.data[p] >= 9:
			goto st9
		}
		goto st0
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
			goto st25
		}
		goto st0
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
		switch it.state.data[p] {
		case 32:
			goto tr111
		case 33:
			goto tr112
		case 35:
			goto tr113
		case 39:
			goto tr114
		case 42:
			goto tr113
		case 44:
			goto tr112
		case 46:
			goto tr115
		case 58:
			goto tr113
		case 59:
			goto tr112
		case 61:
			goto tr116
		case 62:
			goto tr117
		case 63:
			goto tr112
		case 64:
			goto tr113
		case 69:
			goto tr118
		case 94:
			goto tr119
		case 101:
			goto tr118
		}
		switch {
		case it.state.data[p] > 13:
			if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
				goto st25
			}
		case it.state.data[p] >= 9:
			goto tr111
		}
		goto st0
	tr115:
//line ./op-grammar.rl:47

//line ./op-grammar.rl:49
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:27
		// TODO max size for int/float/string
		it.AddAtom(i)

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
		done = true
		{
			p++
			it.state.cs = 26
			goto _out
		}

		goto st26
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
//line dfa.go:3088
		switch it.state.data[p] {
		case 32:
			goto st21
		case 46:
			goto st21
		}
		switch {
		case it.state.data[p] > 13:
			if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
				goto st27
			}
		case it.state.data[p] >= 9:
			goto st21
		}
		goto st0
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
		switch it.state.data[p] {
		case 32:
			goto tr111
		case 33:
			goto tr112
		case 35:
			goto tr113
		case 39:
			goto tr114
		case 42:
			goto tr113
		case 44:
			goto tr112
		case 46:
			goto tr121
		case 58:
			goto tr113
		case 59:
			goto tr112
		case 61:
			goto tr116
		case 62:
			goto tr117
		case 63:
			goto tr112
		case 64:
			goto tr113
		case 69:
			goto tr118
		case 94:
			goto tr119
		case 101:
			goto tr118
		}
		switch {
		case it.state.data[p] > 13:
			if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
				goto st27
			}
		case it.state.data[p] >= 9:
			goto tr111
		}
		goto st0
	tr118:
//line ./op-grammar.rl:47

		goto st11
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
//line dfa.go:3162
		switch it.state.data[p] {
		case 43:
			goto st12
		case 45:
			goto st12
		}
		if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
			goto st28
		}
		goto st0
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
			goto st28
		}
		goto st0
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
		switch it.state.data[p] {
		case 32:
			goto tr122
		case 33:
			goto tr123
		case 35:
			goto tr124
		case 39:
			goto tr125
		case 42:
			goto tr124
		case 44:
			goto tr123
		case 46:
			goto tr126
		case 58:
			goto tr124
		case 59:
			goto tr123
		case 61:
			goto tr127
		case 62:
			goto tr128
		case 63:
			goto tr123
		case 64:
			goto tr124
		case 94:
			goto tr129
		}
		switch {
		case it.state.data[p] > 13:
			if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
				goto st28
			}
		case it.state.data[p] >= 9:
			goto tr122
		}
		goto st0
	tr34:
//line ././uuid-grammar.rl:5
		half = 0

//line ././uuid-grammar.rl:27

//line ././uuid-grammar.rl:10
		digit = prefixSep2Bits(it.state.data[p]) + 4
		i[half] &= INT60_FLAGS | PREFIX_MASKS[digit]

		goto st29
	tr137:
//line ././uuid-grammar.rl:15
		i[half] |= uint64(ABC[it.state.data[p]]) << DIGIT_OFFSETS[digit]
		digit++
		if digit > 10 {
			{
				p++
				it.state.cs = 29
				goto _out
			}
		}

		goto st29
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
//line dfa.go:3255
		switch it.state.data[p] {
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
		case it.state.data[p] < 43:
			switch {
			case it.state.data[p] < 36:
				if 9 <= it.state.data[p] && it.state.data[p] <= 13 {
					goto tr130
				}
			case it.state.data[p] > 37:
				if 40 <= it.state.data[p] && it.state.data[p] <= 41 {
					goto tr135
				}
			default:
				goto tr133
			}
		case it.state.data[p] > 45:
			switch {
			case it.state.data[p] < 65:
				if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
					goto tr137
				}
			case it.state.data[p] > 90:
				if 97 <= it.state.data[p] && it.state.data[p] <= 122 {
					goto tr137
				}
			default:
				goto tr137
			}
		default:
			goto tr133
		}
		goto st0
	tr35:
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
				it.state.cs = 30
				goto _out
			}
		}

		goto st30
	tr146:
//line ././uuid-grammar.rl:15
		i[half] |= uint64(ABC[it.state.data[p]]) << DIGIT_OFFSETS[digit]
		digit++
		if digit > 10 {
			{
				p++
				it.state.cs = 30
				goto _out
			}
		}

		goto st30
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
//line dfa.go:3365
		switch it.state.data[p] {
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
			goto tr147
		case 62:
			goto tr148
		case 63:
			goto tr142
		case 64:
			goto tr143
		case 91:
			goto tr135
		case 93:
			goto tr135
		case 94:
			goto tr149
		case 95:
			goto tr146
		case 123:
			goto tr135
		case 125:
			goto tr135
		case 126:
			goto tr146
		}
		switch {
		case it.state.data[p] < 43:
			switch {
			case it.state.data[p] < 36:
				if 9 <= it.state.data[p] && it.state.data[p] <= 13 {
					goto tr141
				}
			case it.state.data[p] > 37:
				if 40 <= it.state.data[p] && it.state.data[p] <= 41 {
					goto tr135
				}
			default:
				goto tr133
			}
		case it.state.data[p] > 45:
			switch {
			case it.state.data[p] < 65:
				if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
					goto tr146
				}
			case it.state.data[p] > 90:
				if 97 <= it.state.data[p] && it.state.data[p] <= 122 {
					goto tr146
				}
			default:
				goto tr146
			}
		default:
			goto tr133
		}
		goto st0
	tr24:
//line ./op-grammar.rl:53
		i[0] = uint64(p)

		goto st13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
//line dfa.go:3450
		switch it.state.data[p] {
		case 10:
			goto st0
		case 13:
			goto st0
		}
		goto st5
	tr17:
//line ././uuid-grammar.rl:30
		digit = 0
		half |= 1

//line ././uuid-grammar.rl:10
		digit = prefixSep2Bits(it.state.data[p]) + 4
		i[half] &= INT60_FLAGS | PREFIX_MASKS[digit]

		goto st14
	tr18:
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
				it.state.cs = 14
				goto _out
			}
		}

		goto st14
	tr45:
//line ././uuid-grammar.rl:15
		i[half] |= uint64(ABC[it.state.data[p]]) << DIGIT_OFFSETS[digit]
		digit++
		if digit > 10 {
			{
				p++
				it.state.cs = 14
				goto _out
			}
		}

		goto st14
	tr54:
//line ././uuid-grammar.rl:35

//line ././uuid-grammar.rl:30
		digit = 0
		half |= 1

//line ././uuid-grammar.rl:10
		digit = prefixSep2Bits(it.state.data[p]) + 4
		i[half] &= INT60_FLAGS | PREFIX_MASKS[digit]

		goto st14
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
//line dfa.go:3519
		switch it.state.data[p] {
		case 32:
			goto tr41
		case 33:
			goto tr42
		case 35:
			goto tr43
		case 39:
			goto tr44
		case 42:
			goto tr43
		case 44:
			goto tr42
		case 58:
			goto tr43
		case 59:
			goto tr42
		case 61:
			goto tr46
		case 62:
			goto tr47
		case 63:
			goto tr42
		case 64:
			goto tr43
		case 94:
			goto tr48
		case 95:
			goto tr45
		case 126:
			goto tr45
		}
		switch {
		case it.state.data[p] < 48:
			if 9 <= it.state.data[p] && it.state.data[p] <= 13 {
				goto tr41
			}
		case it.state.data[p] > 57:
			switch {
			case it.state.data[p] > 90:
				if 97 <= it.state.data[p] && it.state.data[p] <= 122 {
					goto tr45
				}
			case it.state.data[p] >= 65:
				goto tr45
			}
		default:
			goto tr45
		}
		goto st0
	tr10:
//line ././uuid-grammar.rl:5
		half = 0

//line ././uuid-grammar.rl:27

//line ././uuid-grammar.rl:10
		digit = prefixSep2Bits(it.state.data[p]) + 4
		i[half] &= INT60_FLAGS | PREFIX_MASKS[digit]

		goto st15
	tr55:
//line ././uuid-grammar.rl:15
		i[half] |= uint64(ABC[it.state.data[p]]) << DIGIT_OFFSETS[digit]
		digit++
		if digit > 10 {
			{
				p++
				it.state.cs = 15
				goto _out
			}
		}

		goto st15
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
//line dfa.go:3599
		switch it.state.data[p] {
		case 32:
			goto tr49
		case 33:
			goto tr50
		case 35:
			goto tr51
		case 39:
			goto tr53
		case 42:
			goto tr51
		case 44:
			goto tr50
		case 58:
			goto tr51
		case 59:
			goto tr50
		case 61:
			goto tr56
		case 62:
			goto tr57
		case 63:
			goto tr50
		case 64:
			goto tr51
		case 91:
			goto tr54
		case 93:
			goto tr54
		case 94:
			goto tr58
		case 95:
			goto tr55
		case 123:
			goto tr54
		case 125:
			goto tr54
		case 126:
			goto tr55
		}
		switch {
		case it.state.data[p] < 43:
			switch {
			case it.state.data[p] < 36:
				if 9 <= it.state.data[p] && it.state.data[p] <= 13 {
					goto tr49
				}
			case it.state.data[p] > 37:
				if 40 <= it.state.data[p] && it.state.data[p] <= 41 {
					goto tr54
				}
			default:
				goto tr52
			}
		case it.state.data[p] > 45:
			switch {
			case it.state.data[p] < 65:
				if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
					goto tr55
				}
			case it.state.data[p] > 90:
				if 97 <= it.state.data[p] && it.state.data[p] <= 122 {
					goto tr55
				}
			default:
				goto tr55
			}
		default:
			goto tr52
		}
		goto st0
	tr11:
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
				it.state.cs = 16
				goto _out
			}
		}

		goto st16
	tr63:
//line ././uuid-grammar.rl:15
		i[half] |= uint64(ABC[it.state.data[p]]) << DIGIT_OFFSETS[digit]
		digit++
		if digit > 10 {
			{
				p++
				it.state.cs = 16
				goto _out
			}
		}

		goto st16
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
//line dfa.go:3707
		switch it.state.data[p] {
		case 32:
			goto tr59
		case 33:
			goto tr60
		case 35:
			goto tr61
		case 39:
			goto tr62
		case 42:
			goto tr61
		case 44:
			goto tr60
		case 58:
			goto tr61
		case 59:
			goto tr60
		case 61:
			goto tr64
		case 62:
			goto tr65
		case 63:
			goto tr60
		case 64:
			goto tr61
		case 91:
			goto tr54
		case 93:
			goto tr54
		case 94:
			goto tr66
		case 95:
			goto tr63
		case 123:
			goto tr54
		case 125:
			goto tr54
		case 126:
			goto tr63
		}
		switch {
		case it.state.data[p] < 43:
			switch {
			case it.state.data[p] < 36:
				if 9 <= it.state.data[p] && it.state.data[p] <= 13 {
					goto tr59
				}
			case it.state.data[p] > 37:
				if 40 <= it.state.data[p] && it.state.data[p] <= 41 {
					goto tr54
				}
			default:
				goto tr52
			}
		case it.state.data[p] > 45:
			switch {
			case it.state.data[p] < 65:
				if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
					goto tr63
				}
			case it.state.data[p] > 90:
				if 97 <= it.state.data[p] && it.state.data[p] <= 122 {
					goto tr63
				}
			default:
				goto tr63
			}
		default:
			goto tr52
		}
		goto st0
	tr12:
//line ./op-grammar.rl:6
		if idx != 0 {
			it.uuids[idx] = it.uuids[idx-1]
		}

		goto st17
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
//line dfa.go:3792
		switch it.state.data[p] {
		case 43:
			goto tr9
		case 45:
			goto tr9
		case 91:
			goto tr10
		case 93:
			goto tr10
		case 95:
			goto tr11
		case 123:
			goto tr10
		case 125:
			goto tr10
		case 126:
			goto tr11
		}
		switch {
		case it.state.data[p] < 48:
			switch {
			case it.state.data[p] > 37:
				if 40 <= it.state.data[p] && it.state.data[p] <= 41 {
					goto tr10
				}
			case it.state.data[p] >= 36:
				goto tr9
			}
		case it.state.data[p] > 57:
			switch {
			case it.state.data[p] > 90:
				if 97 <= it.state.data[p] && it.state.data[p] <= 122 {
					goto tr11
				}
			case it.state.data[p] >= 65:
				goto tr11
			}
		default:
			goto tr11
		}
		goto st0
	st_out:
	_test_eof1:
		it.state.cs = 1
		goto _test_eof
	_test_eof18:
		it.state.cs = 18
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
	_test_eof5:
		it.state.cs = 5
		goto _test_eof
	_test_eof19:
		it.state.cs = 19
		goto _test_eof
	_test_eof20:
		it.state.cs = 20
		goto _test_eof
	_test_eof21:
		it.state.cs = 21
		goto _test_eof
	_test_eof6:
		it.state.cs = 6
		goto _test_eof
	_test_eof7:
		it.state.cs = 7
		goto _test_eof
	_test_eof22:
		it.state.cs = 22
		goto _test_eof
	_test_eof8:
		it.state.cs = 8
		goto _test_eof
	_test_eof23:
		it.state.cs = 23
		goto _test_eof
	_test_eof24:
		it.state.cs = 24
		goto _test_eof
	_test_eof9:
		it.state.cs = 9
		goto _test_eof
	_test_eof10:
		it.state.cs = 10
		goto _test_eof
	_test_eof25:
		it.state.cs = 25
		goto _test_eof
	_test_eof26:
		it.state.cs = 26
		goto _test_eof
	_test_eof27:
		it.state.cs = 27
		goto _test_eof
	_test_eof11:
		it.state.cs = 11
		goto _test_eof
	_test_eof12:
		it.state.cs = 12
		goto _test_eof
	_test_eof28:
		it.state.cs = 28
		goto _test_eof
	_test_eof29:
		it.state.cs = 29
		goto _test_eof
	_test_eof30:
		it.state.cs = 30
		goto _test_eof
	_test_eof13:
		it.state.cs = 13
		goto _test_eof
	_test_eof14:
		it.state.cs = 14
		goto _test_eof
	_test_eof15:
		it.state.cs = 15
		goto _test_eof
	_test_eof16:
		it.state.cs = 16
		goto _test_eof
	_test_eof17:
		it.state.cs = 17
		goto _test_eof

	_test_eof:
		{
		}
		if p == eof {
			switch it.state.cs {
			case 18:
//line ./op-grammar.rl:75
				done = true
				{
					p++
					it.state.cs = 0
					goto _out
				}

			case 20:
//line ./op-grammar.rl:68
//line ./op-grammar.rl:75
				done = true
				{
					p++
					it.state.cs = 0
					goto _out
				}

			case 19, 23:
//line ./op-grammar.rl:27
				// TODO max size for int/float/string
				it.AddAtom(i)

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
				done = true
				{
					p++
					it.state.cs = 0
					goto _out
				}

			case 29:
//line ././uuid-grammar.rl:35

//line ./op-grammar.rl:27

				// TODO max size for int/float/string
				it.AddAtom(i)

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
				done = true
				{
					p++
					it.state.cs = 0
					goto _out
				}

			case 24:
//line ././uuid-grammar.rl:38

//line ./op-grammar.rl:27

				// TODO max size for int/float/string
				it.AddAtom(i)

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
				done = true
				{
					p++
					it.state.cs = 0
					goto _out
				}

			case 22:
//line ./op-grammar.rl:43
				i[1] |= ATOM_INT_62

//line ./op-grammar.rl:27
				// TODO max size for int/float/string
				it.AddAtom(i)

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
				done = true
				{
					p++
					it.state.cs = 0
					goto _out
				}

			case 28:
//line ./op-grammar.rl:49
				i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:27
				// TODO max size for int/float/string
				it.AddAtom(i)

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
				done = true
				{
					p++
					it.state.cs = 0
					goto _out
				}

			case 30:
//line ././uuid-grammar.rl:35

//line ././uuid-grammar.rl:47
				i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:27

				// TODO max size for int/float/string
				it.AddAtom(i)

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
				done = true
				{
					p++
					it.state.cs = 0
					goto _out
				}

			case 25, 27:
//line ./op-grammar.rl:47
//line ./op-grammar.rl:49
				i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:27
				// TODO max size for int/float/string
				it.AddAtom(i)

//line ./op-grammar.rl:68

//line ./op-grammar.rl:75
				done = true
				{
					p++
					it.state.cs = 0
					goto _out
				}

//line dfa.go:4010
			}
		}

	_out:
		{
		}
	}

//line dfa.rl:93

	it.state.incomplete = i
	it.state.idx = idx
	it.state.digit = digit
	it.state.half = half
	it.state.p = p

	if it.state.cs == RON_error {
		return PARSED_ERROR
	} else if it.state.cs >= RON_first_final {
		return PARSED_EOF
	} else if p < pe {
		return PARSED_OP
	} else {
		return PARSED_INCOMPLETE
	}
}

var DIGIT_OFFSETS [11]uint8
var PREFIX_MASKS [11]uint64

func init() {
	var one uint64 = 1
	for i := 0; i < 11; i++ {
		var bitoff uint8 = uint8(60 - i*6)
		DIGIT_OFFSETS[i] = bitoff - 6
		PREFIX_MASKS[i] = ((one << 60) - 1) - ((one << bitoff) - 1)
	}
}

func (ctx_uuid UUID) Parse(data []byte) (UUID, error) {

//line dfa.rl:128
//line dfa.go:4055
	const UUID_start int = 1
	const UUID_first_final int = 2
	const UUID_error int = 0

	const UUID_en_main int = 1

//line dfa.rl:129
	var i uint128 = ctx_uuid.uint128
	digit := uint(0)
	half := 0

	cs, p, pe, eof := 0, 0, len(data), len(data)
	//var ts, te, act int
	_ = eof
	//_,_,_ = ts,te,act

//line dfa.go:4076
	{
		cs = UUID_start
	}

//line dfa.go:4081
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
//line dfa.go:4175
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
//line dfa.go:4269
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
//line dfa.go:4318
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
//line dfa.go:4396
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

//line dfa.go:4463
			}
		}

	_out:
		{
		}
	}

//line dfa.rl:147

	if cs < UUID_first_final || digit > 10 {
		return ERROR_UUID, errors.New(fmt.Sprintf("parse error at pos %d", p))
	} else {
		return UUID{uint128: i}, nil
	}

}
