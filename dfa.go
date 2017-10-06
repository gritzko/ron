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

var SYNTAX_ERROR = NewError("BadSyntax")
var LIMIT_ERROR = NewError("SyntxLimit")
var EOF_ERROR = NewError("EOF")
var INCOMPLETE_ERROR = NewError("Incomplete")

//line dfa.rl:42
//line dfa.go:47
const RON_start int = 1
const RON_first_final int = 20
const RON_error int = 0

const RON_en_main int = 1

//line dfa.rl:43

//line dfa.rl:44
// Parse consumes one op, unless the buffer ends earlier.
func (it *Frame) Parse() error {

	//fmt.Println("GO");

	if it.IsLast() {
		it.Op = ZERO_OP
		return EOF_ERROR
	}

	if it.state.cs == 0 {
		//fmt.Println("INIT");

//line dfa.go:72
		{
			it.state.cs = RON_start
		}

//line dfa.rl:58
	} else if it.state.cs >= RON_first_final {
		it.state.cs = RON_start
	}

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
	//fmt.Println("GO!", it.state.cs, "at", p, "with", it.state.data[p]);

//line dfa.go:98
	{
		if p == pe {
			goto _test_eof
		}
		goto _resume

	_again:
		switch it.state.cs {
		case 1:
			goto st1
		case 0:
			goto st0
		case 2:
			goto st2
		case 3:
			goto st3
		case 4:
			goto st4
		case 5:
			goto st5
		case 20:
			goto st20
		case 21:
			goto st21
		case 6:
			goto st6
		case 7:
			goto st7
		case 22:
			goto st22
		case 23:
			goto st23
		case 8:
			goto st8
		case 9:
			goto st9
		case 24:
			goto st24
		case 10:
			goto st10
		case 25:
			goto st25
		case 26:
			goto st26
		case 11:
			goto st11
		case 12:
			goto st12
		case 27:
			goto st27
		case 28:
			goto st28
		case 29:
			goto st29
		case 13:
			goto st13
		case 14:
			goto st14
		case 30:
			goto st30
		case 31:
			goto st31
		case 32:
			goto st32
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
		}

		if p++; p == pe {
			goto _test_eof
		}
	_resume:
		switch it.state.cs {
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
		case 23:
			goto st_case_23
		case 8:
			goto st_case_8
		case 9:
			goto st_case_9
		case 24:
			goto st_case_24
		case 10:
			goto st_case_10
		case 25:
			goto st_case_25
		case 26:
			goto st_case_26
		case 11:
			goto st_case_11
		case 12:
			goto st_case_12
		case 27:
			goto st_case_27
		case 28:
			goto st_case_28
		case 29:
			goto st_case_29
		case 13:
			goto st_case_13
		case 14:
			goto st_case_14
		case 30:
			goto st_case_30
		case 31:
			goto st_case_31
		case 32:
			goto st_case_32
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
		}
		goto st_out
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		switch it.state.data[p] {
		case 32:
			goto st1
		case 35:
			goto st2
		case 42:
			goto st2
		case 58:
			goto st2
		case 64:
			goto st2
		}
		if 9 <= it.state.data[p] && it.state.data[p] <= 13 {
			goto st1
		}
		goto st0
	st_case_0:
	st0:
		it.state.cs = 0
		goto _out
	tr15:
//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

		goto st2
	tr49:
//line ././uuid-grammar.rl:39

//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

		goto st2
	tr57:
//line ././uuid-grammar.rl:36

//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

		goto st2
	tr67:
//line ././uuid-grammar.rl:36

//line ././uuid-grammar.rl:48
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

		goto st2
	tr74:
		it.state.cs = 2
//line ./op-grammar.rl:94

		//fmt.Println("END", it.state.cs, "AT", p)
		if p < pe {
			//fmt.Println("BACK")
			p--
			it.state.cs = (RON_start)
			{
				p++
				goto _out
			}
		}

		goto _again
	tr79:
		it.state.cs = 2
//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:80

//line ./op-grammar.rl:94

		//fmt.Println("END", it.state.cs, "AT", p)
		if p < pe {
			//fmt.Println("BACK")
			p--
			it.state.cs = (RON_start)
			{
				p++
				goto _out
			}
		}

		goto _again
	tr87:
		it.state.cs = 2
//line ./op-grammar.rl:80
//line ./op-grammar.rl:94

		//fmt.Println("END", it.state.cs, "AT", p)
		if p < pe {
			//fmt.Println("BACK")
			p--
			it.state.cs = (RON_start)
			{
				p++
				goto _out
			}
		}

		goto _again
	tr95:
		it.state.cs = 2
//line ./op-grammar.rl:53
		i[1] |= ATOM_INT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:80

//line ./op-grammar.rl:94

		//fmt.Println("END", it.state.cs, "AT", p)
		if p < pe {
			//fmt.Println("BACK")
			p--
			it.state.cs = (RON_start)
			{
				p++
				goto _out
			}
		}

		goto _again
	tr105:
		it.state.cs = 2
//line ././uuid-grammar.rl:39

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:80

//line ./op-grammar.rl:94

		//fmt.Println("END", it.state.cs, "AT", p)
		if p < pe {
			//fmt.Println("BACK")
			p--
			it.state.cs = (RON_start)
			{
				p++
				goto _out
			}
		}

		goto _again
	tr114:
		it.state.cs = 2
//line ./op-grammar.rl:57
//line ./op-grammar.rl:59
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:80

//line ./op-grammar.rl:94

		//fmt.Println("END", it.state.cs, "AT", p)
		if p < pe {
			//fmt.Println("BACK")
			p--
			it.state.cs = (RON_start)
			{
				p++
				goto _out
			}
		}

		goto _again
	tr125:
		it.state.cs = 2
//line ./op-grammar.rl:59
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:80

//line ./op-grammar.rl:94

		//fmt.Println("END", it.state.cs, "AT", p)
		if p < pe {
			//fmt.Println("BACK")
			p--
			it.state.cs = (RON_start)
			{
				p++
				goto _out
			}
		}

		goto _again
	tr133:
		it.state.cs = 2
//line ././uuid-grammar.rl:36

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:80

//line ./op-grammar.rl:94

		//fmt.Println("END", it.state.cs, "AT", p)
		if p < pe {
			//fmt.Println("BACK")
			p--
			it.state.cs = (RON_start)
			{
				p++
				goto _out
			}
		}

		goto _again
	tr144:
		it.state.cs = 2
//line ././uuid-grammar.rl:36

//line ././uuid-grammar.rl:48
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:80

//line ./op-grammar.rl:94

		//fmt.Println("END", it.state.cs, "AT", p)
		if p < pe {
			//fmt.Println("BACK")
			p--
			it.state.cs = (RON_start)
			{
				p++
				goto _out
			}
		}

		goto _again
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
//line dfa.go:534
		switch it.state.data[p] {
		case 32:
			goto tr3
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
				goto tr3
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
	tr3:
		it.state.cs = 3
//line ./op-grammar.rl:12

		//fmt.Println("UUID", it.state.data[p-1]);
		n = specSep2Bits(it.state.data[p-1])
		if n < idx {
			//fmt.Println("EARLY", n, idx, p)
			it.state.cs = (RON_error)
			{
				p++
				goto _out
			}
		} else {
			idx = n
			i = it.uuids[idx].uint128
			digit = 0
		}

		goto _again
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
//line dfa.go:604
		switch it.state.data[p] {
		case 32:
			goto st3
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
				goto st3
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
	tr4:
		it.state.cs = 4
//line ./op-grammar.rl:12

		//fmt.Println("UUID", it.state.data[p-1]);
		n = specSep2Bits(it.state.data[p-1])
		if n < idx {
			//fmt.Println("EARLY", n, idx, p)
			it.state.cs = (RON_error)
			{
				p++
				goto _out
			}
		} else {
			idx = n
			i = it.uuids[idx].uint128
			digit = 0
		}

//line ././uuid-grammar.rl:5

		//fmt.Println("START");
		half = 0

//line ././uuid-grammar.rl:42
		half |= 1
		i[half] &= INT60_FULL
		i[half] |= uint64(uuidSep2Bits(it.state.data[p])) << 60

		goto _again
	tr9:
//line ././uuid-grammar.rl:5
		//fmt.Println("START");
		half = 0

//line ././uuid-grammar.rl:42
		half |= 1
		i[half] &= INT60_FULL
		i[half] |= uint64(uuidSep2Bits(it.state.data[p])) << 60

		goto st4
	tr58:
//line ././uuid-grammar.rl:36
//line ././uuid-grammar.rl:42
		half |= 1
		i[half] &= INT60_FULL
		i[half] |= uint64(uuidSep2Bits(it.state.data[p])) << 60

		goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
//line dfa.go:709
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
	tr13:
//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

		goto st5
	tr47:
//line ././uuid-grammar.rl:39

//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

		goto st5
	tr55:
//line ././uuid-grammar.rl:36

//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

		goto st5
	tr65:
//line ././uuid-grammar.rl:36

//line ././uuid-grammar.rl:48
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

		goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
//line dfa.go:823
		switch it.state.data[p] {
		case 32:
			goto st5
		case 33:
			goto tr23
		case 35:
			goto st2
		case 39:
			goto tr24
		case 42:
			goto st2
		case 44:
			goto tr23
		case 58:
			goto st2
		case 59:
			goto tr23
		case 61:
			goto tr25
		case 62:
			goto tr26
		case 63:
			goto tr23
		case 64:
			goto st2
		case 94:
			goto tr27
		}
		if 9 <= it.state.data[p] && it.state.data[p] <= 13 {
			goto st5
		}
		goto st0
	tr14:
//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:104
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:90
		idx = 0

//line ./op-grammar.rl:83
		//fmt.Println("TERM", fc, it.state.cs, "AT", p);
		it.term = termSep2Bits(it.state.data[p])

		{
			p++
			it.state.cs = 20
			goto _out
		}

		goto st20
	tr23:
//line ./op-grammar.rl:104
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:90
		idx = 0

//line ./op-grammar.rl:83
		//fmt.Println("TERM", fc, it.state.cs, "AT", p);
		it.term = termSep2Bits(it.state.data[p])

		{
			p++
			it.state.cs = 20
			goto _out
		}

		goto st20
	tr48:
//line ././uuid-grammar.rl:39

//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:104
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:90
		idx = 0

//line ./op-grammar.rl:83
		//fmt.Println("TERM", fc, it.state.cs, "AT", p);
		it.term = termSep2Bits(it.state.data[p])

		{
			p++
			it.state.cs = 20
			goto _out
		}

		goto st20
	tr56:
//line ././uuid-grammar.rl:36

//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:104
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:90
		idx = 0

//line ./op-grammar.rl:83
		//fmt.Println("TERM", fc, it.state.cs, "AT", p);
		it.term = termSep2Bits(it.state.data[p])

		{
			p++
			it.state.cs = 20
			goto _out
		}

		goto st20
	tr66:
//line ././uuid-grammar.rl:36

//line ././uuid-grammar.rl:48
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:104
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:90
		idx = 0

//line ./op-grammar.rl:83
		//fmt.Println("TERM", fc, it.state.cs, "AT", p);
		it.term = termSep2Bits(it.state.data[p])

		{
			p++
			it.state.cs = 20
			goto _out
		}

		goto st20
	tr73:
		it.state.cs = 20
//line ./op-grammar.rl:94

		//fmt.Println("END", it.state.cs, "AT", p)
		if p < pe {
			//fmt.Println("BACK")
			p--
			it.state.cs = (RON_start)
			{
				p++
				goto _out
			}
		}

		goto _again
	tr78:
//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:80

//line ./op-grammar.rl:83

		//fmt.Println("TERM", fc, it.state.cs, "AT", p);
		it.term = termSep2Bits(it.state.data[p])

		{
			p++
			it.state.cs = 20
			goto _out
		}

		goto st20
	tr86:
//line ./op-grammar.rl:80
//line ./op-grammar.rl:83

		//fmt.Println("TERM", fc, it.state.cs, "AT", p);
		it.term = termSep2Bits(it.state.data[p])

		{
			p++
			it.state.cs = 20
			goto _out
		}

		goto st20
	tr94:
//line ./op-grammar.rl:53
		i[1] |= ATOM_INT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:80

//line ./op-grammar.rl:83

		//fmt.Println("TERM", fc, it.state.cs, "AT", p);
		it.term = termSep2Bits(it.state.data[p])

		{
			p++
			it.state.cs = 20
			goto _out
		}

		goto st20
	tr104:
//line ././uuid-grammar.rl:39

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:80

//line ./op-grammar.rl:83

		//fmt.Println("TERM", fc, it.state.cs, "AT", p);
		it.term = termSep2Bits(it.state.data[p])

		{
			p++
			it.state.cs = 20
			goto _out
		}

		goto st20
	tr113:
//line ./op-grammar.rl:57
//line ./op-grammar.rl:59
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:80

//line ./op-grammar.rl:83

		//fmt.Println("TERM", fc, it.state.cs, "AT", p);
		it.term = termSep2Bits(it.state.data[p])

		{
			p++
			it.state.cs = 20
			goto _out
		}

		goto st20
	tr124:
//line ./op-grammar.rl:59
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:80

//line ./op-grammar.rl:83

		//fmt.Println("TERM", fc, it.state.cs, "AT", p);
		it.term = termSep2Bits(it.state.data[p])

		{
			p++
			it.state.cs = 20
			goto _out
		}

		goto st20
	tr132:
//line ././uuid-grammar.rl:36

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:80

//line ./op-grammar.rl:83

		//fmt.Println("TERM", fc, it.state.cs, "AT", p);
		it.term = termSep2Bits(it.state.data[p])

		{
			p++
			it.state.cs = 20
			goto _out
		}

		goto st20
	tr143:
//line ././uuid-grammar.rl:36

//line ././uuid-grammar.rl:48
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:80

//line ./op-grammar.rl:83

		//fmt.Println("TERM", fc, it.state.cs, "AT", p);
		it.term = termSep2Bits(it.state.data[p])

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
//line dfa.go:1182
		switch it.state.data[p] {
		case 32:
			goto tr73
		case 35:
			goto tr74
		case 42:
			goto tr74
		case 46:
			goto tr75
		case 58:
			goto tr74
		case 64:
			goto tr74
		}
		if 9 <= it.state.data[p] && it.state.data[p] <= 13 {
			goto tr73
		}
		goto st0
	tr75:
		it.state.cs = 21
//line ./op-grammar.rl:94

		//fmt.Println("END", it.state.cs, "AT", p)
		if p < pe {
			//fmt.Println("BACK")
			p--
			it.state.cs = (RON_start)
			{
				p++
				goto _out
			}
		}

		goto _again
	tr81:
		it.state.cs = 21
//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:80

//line ./op-grammar.rl:94

		//fmt.Println("END", it.state.cs, "AT", p)
		if p < pe {
			//fmt.Println("BACK")
			p--
			it.state.cs = (RON_start)
			{
				p++
				goto _out
			}
		}

		goto _again
	tr89:
		it.state.cs = 21
//line ./op-grammar.rl:80
//line ./op-grammar.rl:94

		//fmt.Println("END", it.state.cs, "AT", p)
		if p < pe {
			//fmt.Println("BACK")
			p--
			it.state.cs = (RON_start)
			{
				p++
				goto _out
			}
		}

		goto _again
	tr97:
		it.state.cs = 21
//line ./op-grammar.rl:53
		i[1] |= ATOM_INT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:80

//line ./op-grammar.rl:94

		//fmt.Println("END", it.state.cs, "AT", p)
		if p < pe {
			//fmt.Println("BACK")
			p--
			it.state.cs = (RON_start)
			{
				p++
				goto _out
			}
		}

		goto _again
	tr107:
		it.state.cs = 21
//line ././uuid-grammar.rl:39

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:80

//line ./op-grammar.rl:94

		//fmt.Println("END", it.state.cs, "AT", p)
		if p < pe {
			//fmt.Println("BACK")
			p--
			it.state.cs = (RON_start)
			{
				p++
				goto _out
			}
		}

		goto _again
	tr122:
		it.state.cs = 21
//line ./op-grammar.rl:57
//line ./op-grammar.rl:59
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:80

//line ./op-grammar.rl:94

		//fmt.Println("END", it.state.cs, "AT", p)
		if p < pe {
			//fmt.Println("BACK")
			p--
			it.state.cs = (RON_start)
			{
				p++
				goto _out
			}
		}

		goto _again
	tr127:
		it.state.cs = 21
//line ./op-grammar.rl:59
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:80

//line ./op-grammar.rl:94

		//fmt.Println("END", it.state.cs, "AT", p)
		if p < pe {
			//fmt.Println("BACK")
			p--
			it.state.cs = (RON_start)
			{
				p++
				goto _out
			}
		}

		goto _again
	tr137:
		it.state.cs = 21
//line ././uuid-grammar.rl:36

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:80

//line ./op-grammar.rl:94

		//fmt.Println("END", it.state.cs, "AT", p)
		if p < pe {
			//fmt.Println("BACK")
			p--
			it.state.cs = (RON_start)
			{
				p++
				goto _out
			}
		}

		goto _again
	tr146:
		it.state.cs = 21
//line ././uuid-grammar.rl:36

//line ././uuid-grammar.rl:48
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:80

//line ./op-grammar.rl:94

		//fmt.Println("END", it.state.cs, "AT", p)
		if p < pe {
			//fmt.Println("BACK")
			p--
			it.state.cs = (RON_start)
			{
				p++
				goto _out
			}
		}

		goto _again
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
//line dfa.go:1417
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
	tr16:
//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:104
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:90
		idx = 0

//line ./op-grammar.rl:77
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st6
	tr24:
//line ./op-grammar.rl:104
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:90
		idx = 0

//line ./op-grammar.rl:77
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st6
	tr50:
//line ././uuid-grammar.rl:39

//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:104
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:90
		idx = 0

//line ./op-grammar.rl:77
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st6
	tr59:
//line ././uuid-grammar.rl:36

//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:104
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:90
		idx = 0

//line ./op-grammar.rl:77
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st6
	tr68:
//line ././uuid-grammar.rl:36

//line ././uuid-grammar.rl:48
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:104
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:90
		idx = 0

//line ./op-grammar.rl:77
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st6
	tr80:
//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st6
	tr88:
//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st6
	tr96:
//line ./op-grammar.rl:53
		i[1] |= ATOM_INT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st6
	tr106:
//line ././uuid-grammar.rl:39

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st6
	tr115:
//line ./op-grammar.rl:57

//line ./op-grammar.rl:59
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st6
	tr126:
//line ./op-grammar.rl:59
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st6
	tr135:
//line ././uuid-grammar.rl:36

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st6
	tr145:
//line ././uuid-grammar.rl:36

//line ././uuid-grammar.rl:48
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
//line dfa.go:1711
		switch it.state.data[p] {
		case 10:
			goto st0
		case 13:
			goto st0
		case 39:
			goto tr29
		case 92:
			goto tr30
		}
		goto tr28
	tr28:
//line ./op-grammar.rl:63
		i[0] = uint64(p)

		goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
//line dfa.go:1734
		switch it.state.data[p] {
		case 10:
			goto st0
		case 13:
			goto st0
		case 39:
			goto tr32
		case 92:
			goto st15
		}
		goto st7
	tr29:
//line ./op-grammar.rl:63
		i[0] = uint64(p)

//line ./op-grammar.rl:66
		//fmt.Println("STRING");
		i[1] = uint64(p) | ATOM_STRING_62

		goto st22
	tr32:
//line ./op-grammar.rl:66
		//fmt.Println("STRING");
		i[1] = uint64(p) | ATOM_STRING_62

		goto st22
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
//line dfa.go:1769
		switch it.state.data[p] {
		case 32:
			goto tr77
		case 33:
			goto tr78
		case 35:
			goto tr79
		case 39:
			goto tr80
		case 42:
			goto tr79
		case 44:
			goto tr78
		case 46:
			goto tr81
		case 58:
			goto tr79
		case 59:
			goto tr78
		case 61:
			goto tr82
		case 62:
			goto tr83
		case 63:
			goto tr78
		case 64:
			goto tr79
		case 94:
			goto tr84
		}
		if 9 <= it.state.data[p] && it.state.data[p] <= 13 {
			goto tr77
		}
		goto st0
	tr77:
		it.state.cs = 23
//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:80

//line ./op-grammar.rl:94

		//fmt.Println("END", it.state.cs, "AT", p)
		if p < pe {
			//fmt.Println("BACK")
			p--
			it.state.cs = (RON_start)
			{
				p++
				goto _out
			}
		}

		goto _again
	tr85:
		it.state.cs = 23
//line ./op-grammar.rl:80
//line ./op-grammar.rl:94

		//fmt.Println("END", it.state.cs, "AT", p)
		if p < pe {
			//fmt.Println("BACK")
			p--
			it.state.cs = (RON_start)
			{
				p++
				goto _out
			}
		}

		goto _again
	tr93:
		it.state.cs = 23
//line ./op-grammar.rl:53
		i[1] |= ATOM_INT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:80

//line ./op-grammar.rl:94

		//fmt.Println("END", it.state.cs, "AT", p)
		if p < pe {
			//fmt.Println("BACK")
			p--
			it.state.cs = (RON_start)
			{
				p++
				goto _out
			}
		}

		goto _again
	tr103:
		it.state.cs = 23
//line ././uuid-grammar.rl:39

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:80

//line ./op-grammar.rl:94

		//fmt.Println("END", it.state.cs, "AT", p)
		if p < pe {
			//fmt.Println("BACK")
			p--
			it.state.cs = (RON_start)
			{
				p++
				goto _out
			}
		}

		goto _again
	tr112:
		it.state.cs = 23
//line ./op-grammar.rl:57
//line ./op-grammar.rl:59
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:80

//line ./op-grammar.rl:94

		//fmt.Println("END", it.state.cs, "AT", p)
		if p < pe {
			//fmt.Println("BACK")
			p--
			it.state.cs = (RON_start)
			{
				p++
				goto _out
			}
		}

		goto _again
	tr123:
		it.state.cs = 23
//line ./op-grammar.rl:59
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:80

//line ./op-grammar.rl:94

		//fmt.Println("END", it.state.cs, "AT", p)
		if p < pe {
			//fmt.Println("BACK")
			p--
			it.state.cs = (RON_start)
			{
				p++
				goto _out
			}
		}

		goto _again
	tr131:
		it.state.cs = 23
//line ././uuid-grammar.rl:36

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:80

//line ./op-grammar.rl:94

		//fmt.Println("END", it.state.cs, "AT", p)
		if p < pe {
			//fmt.Println("BACK")
			p--
			it.state.cs = (RON_start)
			{
				p++
				goto _out
			}
		}

		goto _again
	tr142:
		it.state.cs = 23
//line ././uuid-grammar.rl:36

//line ././uuid-grammar.rl:48
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:80

//line ./op-grammar.rl:94

		//fmt.Println("END", it.state.cs, "AT", p)
		if p < pe {
			//fmt.Println("BACK")
			p--
			it.state.cs = (RON_start)
			{
				p++
				goto _out
			}
		}

		goto _again
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
//line dfa.go:2007
		switch it.state.data[p] {
		case 32:
			goto tr85
		case 33:
			goto tr86
		case 35:
			goto tr87
		case 39:
			goto tr88
		case 42:
			goto tr87
		case 44:
			goto tr86
		case 46:
			goto tr89
		case 58:
			goto tr87
		case 59:
			goto tr86
		case 61:
			goto tr90
		case 62:
			goto tr91
		case 63:
			goto tr86
		case 64:
			goto tr87
		case 94:
			goto tr92
		}
		if 9 <= it.state.data[p] && it.state.data[p] <= 13 {
			goto tr85
		}
		goto st0
	tr19:
//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:104
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:90
		idx = 0

//line ./op-grammar.rl:77
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st8
	tr25:
//line ./op-grammar.rl:104
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:90
		idx = 0

//line ./op-grammar.rl:77
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st8
	tr52:
//line ././uuid-grammar.rl:39

//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:104
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:90
		idx = 0

//line ./op-grammar.rl:77
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st8
	tr62:
//line ././uuid-grammar.rl:36

//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:104
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:90
		idx = 0

//line ./op-grammar.rl:77
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st8
	tr70:
//line ././uuid-grammar.rl:36

//line ././uuid-grammar.rl:48
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:104
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:90
		idx = 0

//line ./op-grammar.rl:77
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st8
	tr82:
//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st8
	tr90:
//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st8
	tr98:
//line ./op-grammar.rl:53
		i[1] |= ATOM_INT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st8
	tr109:
//line ././uuid-grammar.rl:39

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st8
	tr117:
//line ./op-grammar.rl:57

//line ./op-grammar.rl:59
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st8
	tr128:
//line ./op-grammar.rl:59
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st8
	tr139:
//line ././uuid-grammar.rl:36

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st8
	tr148:
//line ././uuid-grammar.rl:36

//line ././uuid-grammar.rl:48
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
//line dfa.go:2325
		switch it.state.data[p] {
		case 32:
			goto st8
		case 43:
			goto tr35
		case 45:
			goto tr35
		}
		switch {
		case it.state.data[p] > 13:
			if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
				goto tr36
			}
		case it.state.data[p] >= 9:
			goto st8
		}
		goto st0
	tr35:
//line ./op-grammar.rl:42

//line ./op-grammar.rl:44
		if it.state.data[p] == '-' {
			i[1] |= 1
		}

		goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
//line dfa.go:2359
		if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
			goto tr37
		}
		goto st0
	tr36:
//line ./op-grammar.rl:42

//line ./op-grammar.rl:49
		i[0] *= 10
		i[0] += uint64(int(it.state.data[p]) - int('0'))

		goto st24
	tr37:
//line ./op-grammar.rl:49
		i[0] *= 10
		i[0] += uint64(int(it.state.data[p]) - int('0'))

		goto st24
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
//line dfa.go:2386
		switch it.state.data[p] {
		case 32:
			goto tr93
		case 33:
			goto tr94
		case 35:
			goto tr95
		case 39:
			goto tr96
		case 42:
			goto tr95
		case 44:
			goto tr94
		case 46:
			goto tr97
		case 58:
			goto tr95
		case 59:
			goto tr94
		case 61:
			goto tr98
		case 62:
			goto tr99
		case 63:
			goto tr94
		case 64:
			goto tr95
		case 94:
			goto tr100
		}
		switch {
		case it.state.data[p] > 13:
			if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
				goto tr37
			}
		case it.state.data[p] >= 9:
			goto tr93
		}
		goto st0
	tr20:
//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:104
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:90
		idx = 0

//line ./op-grammar.rl:77
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st10
	tr26:
//line ./op-grammar.rl:104
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:90
		idx = 0

//line ./op-grammar.rl:77
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st10
	tr53:
//line ././uuid-grammar.rl:39

//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:104
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:90
		idx = 0

//line ./op-grammar.rl:77
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st10
	tr63:
//line ././uuid-grammar.rl:36

//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:104
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:90
		idx = 0

//line ./op-grammar.rl:77
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st10
	tr71:
//line ././uuid-grammar.rl:36

//line ././uuid-grammar.rl:48
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:104
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:90
		idx = 0

//line ./op-grammar.rl:77
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st10
	tr83:
//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st10
	tr91:
//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st10
	tr99:
//line ./op-grammar.rl:53
		i[1] |= ATOM_INT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st10
	tr110:
//line ././uuid-grammar.rl:39

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st10
	tr118:
//line ./op-grammar.rl:57

//line ./op-grammar.rl:59
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st10
	tr129:
//line ./op-grammar.rl:59
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st10
	tr140:
//line ././uuid-grammar.rl:36

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st10
	tr149:
//line ././uuid-grammar.rl:36

//line ././uuid-grammar.rl:48
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
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
//line dfa.go:2709
		switch it.state.data[p] {
		case 32:
			goto st10
		case 43:
			goto tr39
		case 45:
			goto tr39
		case 91:
			goto tr40
		case 93:
			goto tr40
		case 95:
			goto tr41
		case 123:
			goto tr40
		case 125:
			goto tr40
		case 126:
			goto tr41
		}
		switch {
		case it.state.data[p] < 40:
			switch {
			case it.state.data[p] > 13:
				if 36 <= it.state.data[p] && it.state.data[p] <= 37 {
					goto tr39
				}
			case it.state.data[p] >= 9:
				goto st10
			}
		case it.state.data[p] > 41:
			switch {
			case it.state.data[p] < 65:
				if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
					goto tr41
				}
			case it.state.data[p] > 90:
				if 97 <= it.state.data[p] && it.state.data[p] <= 122 {
					goto tr41
				}
			default:
				goto tr41
			}
		default:
			goto tr40
		}
		goto st0
	tr39:
//line ././uuid-grammar.rl:5

		//fmt.Println("START");
		half = 0

//line ././uuid-grammar.rl:42
		half |= 1
		i[half] &= INT60_FULL
		i[half] |= uint64(uuidSep2Bits(it.state.data[p])) << 60

		goto st25
	tr134:
//line ././uuid-grammar.rl:36
//line ././uuid-grammar.rl:42
		half |= 1
		i[half] &= INT60_FULL
		i[half] |= uint64(uuidSep2Bits(it.state.data[p])) << 60

		goto st25
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
//line dfa.go:2786
		switch it.state.data[p] {
		case 32:
			goto tr77
		case 33:
			goto tr78
		case 35:
			goto tr79
		case 39:
			goto tr80
		case 42:
			goto tr79
		case 44:
			goto tr78
		case 46:
			goto tr81
		case 58:
			goto tr79
		case 59:
			goto tr78
		case 61:
			goto tr82
		case 62:
			goto tr83
		case 63:
			goto tr78
		case 64:
			goto tr79
		case 91:
			goto tr101
		case 93:
			goto tr101
		case 94:
			goto tr84
		case 95:
			goto tr102
		case 123:
			goto tr101
		case 125:
			goto tr101
		case 126:
			goto tr102
		}
		switch {
		case it.state.data[p] < 48:
			switch {
			case it.state.data[p] > 13:
				if 40 <= it.state.data[p] && it.state.data[p] <= 41 {
					goto tr101
				}
			case it.state.data[p] >= 9:
				goto tr77
			}
		case it.state.data[p] > 57:
			switch {
			case it.state.data[p] > 90:
				if 97 <= it.state.data[p] && it.state.data[p] <= 122 {
					goto tr102
				}
			case it.state.data[p] >= 65:
				goto tr102
			}
		default:
			goto tr102
		}
		goto st0
	tr101:
//line ././uuid-grammar.rl:31
		digit = 0
		half |= 1

//line ././uuid-grammar.rl:11
		digit = prefixSep2Bits(it.state.data[p]) + 4
		i[half] &= INT60_FLAGS | PREFIX_MASKS[digit]

		goto st26
	tr102:
//line ././uuid-grammar.rl:31
		digit = 0
		half |= 1

//line ././uuid-grammar.rl:24
		i[half] &= INT60_FLAGS

//line ././uuid-grammar.rl:16
		i[half] |= uint64(ABC[it.state.data[p]]) << DIGIT_OFFSETS[digit]
		digit++
		if digit > 10 {
			{
				p++
				it.state.cs = 26
				goto _out
			}
		}

		goto st26
	tr108:
//line ././uuid-grammar.rl:16
		i[half] |= uint64(ABC[it.state.data[p]]) << DIGIT_OFFSETS[digit]
		digit++
		if digit > 10 {
			{
				p++
				it.state.cs = 26
				goto _out
			}
		}

		goto st26
	tr136:
//line ././uuid-grammar.rl:36

//line ././uuid-grammar.rl:31
		digit = 0
		half |= 1

//line ././uuid-grammar.rl:11
		digit = prefixSep2Bits(it.state.data[p]) + 4
		i[half] &= INT60_FLAGS | PREFIX_MASKS[digit]

		goto st26
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
//line dfa.go:2913
		switch it.state.data[p] {
		case 32:
			goto tr103
		case 33:
			goto tr104
		case 35:
			goto tr105
		case 39:
			goto tr106
		case 42:
			goto tr105
		case 44:
			goto tr104
		case 46:
			goto tr107
		case 58:
			goto tr105
		case 59:
			goto tr104
		case 61:
			goto tr109
		case 62:
			goto tr110
		case 63:
			goto tr104
		case 64:
			goto tr105
		case 94:
			goto tr111
		case 95:
			goto tr108
		case 126:
			goto tr108
		}
		switch {
		case it.state.data[p] < 48:
			if 9 <= it.state.data[p] && it.state.data[p] <= 13 {
				goto tr103
			}
		case it.state.data[p] > 57:
			switch {
			case it.state.data[p] > 90:
				if 97 <= it.state.data[p] && it.state.data[p] <= 122 {
					goto tr108
				}
			case it.state.data[p] >= 65:
				goto tr108
			}
		default:
			goto tr108
		}
		goto st0
	tr21:
//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:104
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:90
		idx = 0

//line ./op-grammar.rl:77
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st11
	tr27:
//line ./op-grammar.rl:104
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:90
		idx = 0

//line ./op-grammar.rl:77
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st11
	tr54:
//line ././uuid-grammar.rl:39

//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:104
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:90
		idx = 0

//line ./op-grammar.rl:77
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st11
	tr64:
//line ././uuid-grammar.rl:36

//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:104
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:90
		idx = 0

//line ./op-grammar.rl:77
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st11
	tr72:
//line ././uuid-grammar.rl:36

//line ././uuid-grammar.rl:48
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:104
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:90
		idx = 0

//line ./op-grammar.rl:77
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st11
	tr84:
//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st11
	tr92:
//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st11
	tr100:
//line ./op-grammar.rl:53
		i[1] |= ATOM_INT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st11
	tr111:
//line ././uuid-grammar.rl:39

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st11
	tr120:
//line ./op-grammar.rl:57

//line ./op-grammar.rl:59
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st11
	tr130:
//line ./op-grammar.rl:59
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st11
	tr141:
//line ././uuid-grammar.rl:36

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st11
	tr150:
//line ././uuid-grammar.rl:36

//line ././uuid-grammar.rl:48
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st11
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
//line dfa.go:3249
		switch it.state.data[p] {
		case 32:
			goto st11
		case 43:
			goto st12
		case 45:
			goto st12
		}
		switch {
		case it.state.data[p] > 13:
			if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
				goto st27
			}
		case it.state.data[p] >= 9:
			goto st11
		}
		goto st0
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
			goto st27
		}
		goto st0
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
		switch it.state.data[p] {
		case 32:
			goto tr112
		case 33:
			goto tr113
		case 35:
			goto tr114
		case 39:
			goto tr115
		case 42:
			goto tr114
		case 44:
			goto tr113
		case 46:
			goto tr116
		case 58:
			goto tr114
		case 59:
			goto tr113
		case 61:
			goto tr117
		case 62:
			goto tr118
		case 63:
			goto tr113
		case 64:
			goto tr114
		case 69:
			goto tr119
		case 94:
			goto tr120
		case 101:
			goto tr119
		}
		switch {
		case it.state.data[p] > 13:
			if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
				goto st27
			}
		case it.state.data[p] >= 9:
			goto tr112
		}
		goto st0
	tr116:
		it.state.cs = 28
//line ./op-grammar.rl:57

//line ./op-grammar.rl:59
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADDING", i);
		it.AddAtom(i)

//line ./op-grammar.rl:80

//line ./op-grammar.rl:94

		//fmt.Println("END", it.state.cs, "AT", p)
		if p < pe {
			//fmt.Println("BACK")
			p--
			it.state.cs = (RON_start)
			{
				p++
				goto _out
			}
		}

		goto _again
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
//line dfa.go:3358
		switch it.state.data[p] {
		case 32:
			goto st21
		case 46:
			goto st21
		}
		switch {
		case it.state.data[p] > 13:
			if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
				goto st29
			}
		case it.state.data[p] >= 9:
			goto st21
		}
		goto st0
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
		switch it.state.data[p] {
		case 32:
			goto tr112
		case 33:
			goto tr113
		case 35:
			goto tr114
		case 39:
			goto tr115
		case 42:
			goto tr114
		case 44:
			goto tr113
		case 46:
			goto tr122
		case 58:
			goto tr114
		case 59:
			goto tr113
		case 61:
			goto tr117
		case 62:
			goto tr118
		case 63:
			goto tr113
		case 64:
			goto tr114
		case 69:
			goto tr119
		case 94:
			goto tr120
		case 101:
			goto tr119
		}
		switch {
		case it.state.data[p] > 13:
			if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
				goto st29
			}
		case it.state.data[p] >= 9:
			goto tr112
		}
		goto st0
	tr119:
//line ./op-grammar.rl:57

		goto st13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
//line dfa.go:3432
		switch it.state.data[p] {
		case 43:
			goto st14
		case 45:
			goto st14
		}
		if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
			goto st30
		}
		goto st0
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
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
			goto tr123
		case 33:
			goto tr124
		case 35:
			goto tr125
		case 39:
			goto tr126
		case 42:
			goto tr125
		case 44:
			goto tr124
		case 46:
			goto tr127
		case 58:
			goto tr125
		case 59:
			goto tr124
		case 61:
			goto tr128
		case 62:
			goto tr129
		case 63:
			goto tr124
		case 64:
			goto tr125
		case 94:
			goto tr130
		}
		switch {
		case it.state.data[p] > 13:
			if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
				goto st30
			}
		case it.state.data[p] >= 9:
			goto tr123
		}
		goto st0
	tr40:
//line ././uuid-grammar.rl:5

		//fmt.Println("START");
		half = 0

//line ././uuid-grammar.rl:28

//line ././uuid-grammar.rl:11
		digit = prefixSep2Bits(it.state.data[p]) + 4
		i[half] &= INT60_FLAGS | PREFIX_MASKS[digit]

		goto st31
	tr138:
//line ././uuid-grammar.rl:16
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
//line dfa.go:3526
		switch it.state.data[p] {
		case 32:
			goto tr131
		case 33:
			goto tr132
		case 35:
			goto tr133
		case 39:
			goto tr135
		case 42:
			goto tr133
		case 44:
			goto tr132
		case 46:
			goto tr137
		case 58:
			goto tr133
		case 59:
			goto tr132
		case 61:
			goto tr139
		case 62:
			goto tr140
		case 63:
			goto tr132
		case 64:
			goto tr133
		case 91:
			goto tr136
		case 93:
			goto tr136
		case 94:
			goto tr141
		case 95:
			goto tr138
		case 123:
			goto tr136
		case 125:
			goto tr136
		case 126:
			goto tr138
		}
		switch {
		case it.state.data[p] < 43:
			switch {
			case it.state.data[p] < 36:
				if 9 <= it.state.data[p] && it.state.data[p] <= 13 {
					goto tr131
				}
			case it.state.data[p] > 37:
				if 40 <= it.state.data[p] && it.state.data[p] <= 41 {
					goto tr136
				}
			default:
				goto tr134
			}
		case it.state.data[p] > 45:
			switch {
			case it.state.data[p] < 65:
				if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
					goto tr138
				}
			case it.state.data[p] > 90:
				if 97 <= it.state.data[p] && it.state.data[p] <= 122 {
					goto tr138
				}
			default:
				goto tr138
			}
		default:
			goto tr134
		}
		goto st0
	tr41:
//line ././uuid-grammar.rl:5

		//fmt.Println("START");
		half = 0

//line ././uuid-grammar.rl:28

//line ././uuid-grammar.rl:24
		i[half] &= INT60_FLAGS

//line ././uuid-grammar.rl:16
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
	tr147:
//line ././uuid-grammar.rl:16
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
//line dfa.go:3637
		switch it.state.data[p] {
		case 32:
			goto tr142
		case 33:
			goto tr143
		case 35:
			goto tr144
		case 39:
			goto tr145
		case 42:
			goto tr144
		case 44:
			goto tr143
		case 46:
			goto tr146
		case 58:
			goto tr144
		case 59:
			goto tr143
		case 61:
			goto tr148
		case 62:
			goto tr149
		case 63:
			goto tr143
		case 64:
			goto tr144
		case 91:
			goto tr136
		case 93:
			goto tr136
		case 94:
			goto tr150
		case 95:
			goto tr147
		case 123:
			goto tr136
		case 125:
			goto tr136
		case 126:
			goto tr147
		}
		switch {
		case it.state.data[p] < 43:
			switch {
			case it.state.data[p] < 36:
				if 9 <= it.state.data[p] && it.state.data[p] <= 13 {
					goto tr142
				}
			case it.state.data[p] > 37:
				if 40 <= it.state.data[p] && it.state.data[p] <= 41 {
					goto tr136
				}
			default:
				goto tr134
			}
		case it.state.data[p] > 45:
			switch {
			case it.state.data[p] < 65:
				if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
					goto tr147
				}
			case it.state.data[p] > 90:
				if 97 <= it.state.data[p] && it.state.data[p] <= 122 {
					goto tr147
				}
			default:
				goto tr147
			}
		default:
			goto tr134
		}
		goto st0
	tr30:
//line ./op-grammar.rl:63
		i[0] = uint64(p)

		goto st15
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
//line dfa.go:3722
		switch it.state.data[p] {
		case 10:
			goto st0
		case 13:
			goto st0
		}
		goto st7
	tr17:
//line ././uuid-grammar.rl:31
		digit = 0
		half |= 1

//line ././uuid-grammar.rl:11
		digit = prefixSep2Bits(it.state.data[p]) + 4
		i[half] &= INT60_FLAGS | PREFIX_MASKS[digit]

		goto st16
	tr18:
//line ././uuid-grammar.rl:31
		digit = 0
		half |= 1

//line ././uuid-grammar.rl:24
		i[half] &= INT60_FLAGS

//line ././uuid-grammar.rl:16
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
	tr51:
//line ././uuid-grammar.rl:16
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
	tr60:
//line ././uuid-grammar.rl:36

//line ././uuid-grammar.rl:31
		digit = 0
		half |= 1

//line ././uuid-grammar.rl:11
		digit = prefixSep2Bits(it.state.data[p]) + 4
		i[half] &= INT60_FLAGS | PREFIX_MASKS[digit]

		goto st16
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
//line dfa.go:3791
		switch it.state.data[p] {
		case 32:
			goto tr47
		case 33:
			goto tr48
		case 35:
			goto tr49
		case 39:
			goto tr50
		case 42:
			goto tr49
		case 44:
			goto tr48
		case 58:
			goto tr49
		case 59:
			goto tr48
		case 61:
			goto tr52
		case 62:
			goto tr53
		case 63:
			goto tr48
		case 64:
			goto tr49
		case 94:
			goto tr54
		case 95:
			goto tr51
		case 126:
			goto tr51
		}
		switch {
		case it.state.data[p] < 48:
			if 9 <= it.state.data[p] && it.state.data[p] <= 13 {
				goto tr47
			}
		case it.state.data[p] > 57:
			switch {
			case it.state.data[p] > 90:
				if 97 <= it.state.data[p] && it.state.data[p] <= 122 {
					goto tr51
				}
			case it.state.data[p] >= 65:
				goto tr51
			}
		default:
			goto tr51
		}
		goto st0
	tr5:
		it.state.cs = 17
//line ./op-grammar.rl:12

		//fmt.Println("UUID", it.state.data[p-1]);
		n = specSep2Bits(it.state.data[p-1])
		if n < idx {
			//fmt.Println("EARLY", n, idx, p)
			it.state.cs = (RON_error)
			{
				p++
				goto _out
			}
		} else {
			idx = n
			i = it.uuids[idx].uint128
			digit = 0
		}

//line ././uuid-grammar.rl:5

		//fmt.Println("START");
		half = 0

//line ././uuid-grammar.rl:28

//line ././uuid-grammar.rl:11
		digit = prefixSep2Bits(it.state.data[p]) + 4
		i[half] &= INT60_FLAGS | PREFIX_MASKS[digit]

		goto _again
	tr10:
//line ././uuid-grammar.rl:5
		//fmt.Println("START");
		half = 0

//line ././uuid-grammar.rl:28

//line ././uuid-grammar.rl:11
		digit = prefixSep2Bits(it.state.data[p]) + 4
		i[half] &= INT60_FLAGS | PREFIX_MASKS[digit]

		goto st17
	tr61:
//line ././uuid-grammar.rl:16
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
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
//line dfa.go:3902
		switch it.state.data[p] {
		case 32:
			goto tr55
		case 33:
			goto tr56
		case 35:
			goto tr57
		case 39:
			goto tr59
		case 42:
			goto tr57
		case 44:
			goto tr56
		case 58:
			goto tr57
		case 59:
			goto tr56
		case 61:
			goto tr62
		case 62:
			goto tr63
		case 63:
			goto tr56
		case 64:
			goto tr57
		case 91:
			goto tr60
		case 93:
			goto tr60
		case 94:
			goto tr64
		case 95:
			goto tr61
		case 123:
			goto tr60
		case 125:
			goto tr60
		case 126:
			goto tr61
		}
		switch {
		case it.state.data[p] < 43:
			switch {
			case it.state.data[p] < 36:
				if 9 <= it.state.data[p] && it.state.data[p] <= 13 {
					goto tr55
				}
			case it.state.data[p] > 37:
				if 40 <= it.state.data[p] && it.state.data[p] <= 41 {
					goto tr60
				}
			default:
				goto tr58
			}
		case it.state.data[p] > 45:
			switch {
			case it.state.data[p] < 65:
				if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
					goto tr61
				}
			case it.state.data[p] > 90:
				if 97 <= it.state.data[p] && it.state.data[p] <= 122 {
					goto tr61
				}
			default:
				goto tr61
			}
		default:
			goto tr58
		}
		goto st0
	tr6:
		it.state.cs = 18
//line ./op-grammar.rl:12

		//fmt.Println("UUID", it.state.data[p-1]);
		n = specSep2Bits(it.state.data[p-1])
		if n < idx {
			//fmt.Println("EARLY", n, idx, p)
			it.state.cs = (RON_error)
			{
				p++
				goto _out
			}
		} else {
			idx = n
			i = it.uuids[idx].uint128
			digit = 0
		}

//line ././uuid-grammar.rl:5

		//fmt.Println("START");
		half = 0

//line ././uuid-grammar.rl:28

//line ././uuid-grammar.rl:24
		i[half] &= INT60_FLAGS

//line ././uuid-grammar.rl:16
		i[half] |= uint64(ABC[it.state.data[p]]) << DIGIT_OFFSETS[digit]
		digit++
		if digit > 10 {
			{
				p++
				goto _out
			}
		}

		goto _again
	tr11:
//line ././uuid-grammar.rl:5
		//fmt.Println("START");
		half = 0

//line ././uuid-grammar.rl:28

//line ././uuid-grammar.rl:24
		i[half] &= INT60_FLAGS

//line ././uuid-grammar.rl:16
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
	tr69:
//line ././uuid-grammar.rl:16
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
//line dfa.go:4048
		switch it.state.data[p] {
		case 32:
			goto tr65
		case 33:
			goto tr66
		case 35:
			goto tr67
		case 39:
			goto tr68
		case 42:
			goto tr67
		case 44:
			goto tr66
		case 58:
			goto tr67
		case 59:
			goto tr66
		case 61:
			goto tr70
		case 62:
			goto tr71
		case 63:
			goto tr66
		case 64:
			goto tr67
		case 91:
			goto tr60
		case 93:
			goto tr60
		case 94:
			goto tr72
		case 95:
			goto tr69
		case 123:
			goto tr60
		case 125:
			goto tr60
		case 126:
			goto tr69
		}
		switch {
		case it.state.data[p] < 43:
			switch {
			case it.state.data[p] < 36:
				if 9 <= it.state.data[p] && it.state.data[p] <= 13 {
					goto tr65
				}
			case it.state.data[p] > 37:
				if 40 <= it.state.data[p] && it.state.data[p] <= 41 {
					goto tr60
				}
			default:
				goto tr58
			}
		case it.state.data[p] > 45:
			switch {
			case it.state.data[p] < 65:
				if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
					goto tr69
				}
			case it.state.data[p] > 90:
				if 97 <= it.state.data[p] && it.state.data[p] <= 122 {
					goto tr69
				}
			default:
				goto tr69
			}
		default:
			goto tr58
		}
		goto st0
	tr7:
		it.state.cs = 19
//line ./op-grammar.rl:12

		//fmt.Println("UUID", it.state.data[p-1]);
		n = specSep2Bits(it.state.data[p-1])
		if n < idx {
			//fmt.Println("EARLY", n, idx, p)
			it.state.cs = (RON_error)
			{
				p++
				goto _out
			}
		} else {
			idx = n
			i = it.uuids[idx].uint128
			digit = 0
		}

//line ./op-grammar.rl:6
		if idx != 0 {
			it.uuids[idx] = it.uuids[idx-1]
		}

		goto _again
	tr12:
//line ./op-grammar.rl:6
		if idx != 0 {
			it.uuids[idx] = it.uuids[idx-1]
		}

		goto st19
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
//line dfa.go:4156
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
	_test_eof23:
		it.state.cs = 23
		goto _test_eof
	_test_eof8:
		it.state.cs = 8
		goto _test_eof
	_test_eof9:
		it.state.cs = 9
		goto _test_eof
	_test_eof24:
		it.state.cs = 24
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
	_test_eof11:
		it.state.cs = 11
		goto _test_eof
	_test_eof12:
		it.state.cs = 12
		goto _test_eof
	_test_eof27:
		it.state.cs = 27
		goto _test_eof
	_test_eof28:
		it.state.cs = 28
		goto _test_eof
	_test_eof29:
		it.state.cs = 29
		goto _test_eof
	_test_eof13:
		it.state.cs = 13
		goto _test_eof
	_test_eof14:
		it.state.cs = 14
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
	_test_eof15:
		it.state.cs = 15
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

	_test_eof:
		{
		}
		if p == eof {
			switch it.state.cs {
			case 20:
//line ./op-grammar.rl:94

				//fmt.Println("END", it.state.cs, "AT", p)
				if p < pe {
					//fmt.Println("BACK")
					p--
					it.state.cs = (RON_start)
					{
						p++
						it.state.cs = 0
						goto _out
					}
				}

			case 23:
//line ./op-grammar.rl:80
//line ./op-grammar.rl:94

				//fmt.Println("END", it.state.cs, "AT", p)
				if p < pe {
					//fmt.Println("BACK")
					p--
					it.state.cs = (RON_start)
					{
						p++
						it.state.cs = 0
						goto _out
					}
				}

			case 22, 25:
//line ./op-grammar.rl:36
				// TODO max size for int/float/string
				//fmt.Println("ADDING", i);
				it.AddAtom(i)

//line ./op-grammar.rl:80

//line ./op-grammar.rl:94

				//fmt.Println("END", it.state.cs, "AT", p)
				if p < pe {
					//fmt.Println("BACK")
					p--
					it.state.cs = (RON_start)
					{
						p++
						it.state.cs = 0
						goto _out
					}
				}

			case 31:
//line ././uuid-grammar.rl:36

//line ./op-grammar.rl:36

				// TODO max size for int/float/string
				//fmt.Println("ADDING", i);
				it.AddAtom(i)

//line ./op-grammar.rl:80

//line ./op-grammar.rl:94

				//fmt.Println("END", it.state.cs, "AT", p)
				if p < pe {
					//fmt.Println("BACK")
					p--
					it.state.cs = (RON_start)
					{
						p++
						it.state.cs = 0
						goto _out
					}
				}

			case 26:
//line ././uuid-grammar.rl:39

//line ./op-grammar.rl:36

				// TODO max size for int/float/string
				//fmt.Println("ADDING", i);
				it.AddAtom(i)

//line ./op-grammar.rl:80

//line ./op-grammar.rl:94

				//fmt.Println("END", it.state.cs, "AT", p)
				if p < pe {
					//fmt.Println("BACK")
					p--
					it.state.cs = (RON_start)
					{
						p++
						it.state.cs = 0
						goto _out
					}
				}

			case 24:
//line ./op-grammar.rl:53
				i[1] |= ATOM_INT_62

//line ./op-grammar.rl:36
				// TODO max size for int/float/string
				//fmt.Println("ADDING", i);
				it.AddAtom(i)

//line ./op-grammar.rl:80

//line ./op-grammar.rl:94

				//fmt.Println("END", it.state.cs, "AT", p)
				if p < pe {
					//fmt.Println("BACK")
					p--
					it.state.cs = (RON_start)
					{
						p++
						it.state.cs = 0
						goto _out
					}
				}

			case 30:
//line ./op-grammar.rl:59
				i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
				// TODO max size for int/float/string
				//fmt.Println("ADDING", i);
				it.AddAtom(i)

//line ./op-grammar.rl:80

//line ./op-grammar.rl:94

				//fmt.Println("END", it.state.cs, "AT", p)
				if p < pe {
					//fmt.Println("BACK")
					p--
					it.state.cs = (RON_start)
					{
						p++
						it.state.cs = 0
						goto _out
					}
				}

			case 32:
//line ././uuid-grammar.rl:36

//line ././uuid-grammar.rl:48
				i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:36

				// TODO max size for int/float/string
				//fmt.Println("ADDING", i);
				it.AddAtom(i)

//line ./op-grammar.rl:80

//line ./op-grammar.rl:94

				//fmt.Println("END", it.state.cs, "AT", p)
				if p < pe {
					//fmt.Println("BACK")
					p--
					it.state.cs = (RON_start)
					{
						p++
						it.state.cs = 0
						goto _out
					}
				}

			case 27, 29:
//line ./op-grammar.rl:57
//line ./op-grammar.rl:59
				i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
				// TODO max size for int/float/string
				//fmt.Println("ADDING", i);
				it.AddAtom(i)

//line ./op-grammar.rl:80

//line ./op-grammar.rl:94

				//fmt.Println("END", it.state.cs, "AT", p)
				if p < pe {
					//fmt.Println("BACK")
					p--
					it.state.cs = (RON_start)
					{
						p++
						it.state.cs = 0
						goto _out
					}
				}

//line dfa.go:4428
			}
		}

	_out:
		{
		}
	}

//line dfa.rl:83

	//fmt.Println("DONE", it.state.cs, "at", p);

	it.state.incomplete = i
	it.state.idx = idx
	it.state.digit = digit
	it.state.half = half
	it.state.p = p

	if p >= pe && !it.state.streaming && it.state.cs != RON_start && it.state.cs < RON_first_final {
		it.state.cs = RON_error
		//fmt.Println("BAD", p, pe, it.state.cs)
	}

	if it.state.cs == RON_start || it.state.cs >= RON_first_final || p == pe {
		return nil
	} else if it.state.cs == RON_error {
		//fmt.Println("DONE1", p);
		it.Op = ZERO_OP
		return SYNTAX_ERROR
	} else {
		//fmt.Println("DONE2", p);
		return INCOMPLETE_ERROR
	}
}

func (frame Frame) EOF() bool {
	return frame.state.cs == RON_error
}

func (frame Frame) Offset() int {
	return frame.state.p
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

//line dfa.rl:133
//line dfa.go:4488
	const UUID_start int = 1
	const UUID_first_final int = 2
	const UUID_error int = 0

	const UUID_en_main int = 1

//line dfa.rl:134
	var i uint128 = ctx_uuid.uint128
	digit := uint(0)
	half := 0

	cs, p, pe, eof := 0, 0, len(data), len(data)
	//var ts, te, act int
	_ = eof
	//_,_,_ = ts,te,act

//line dfa.go:4509
	{
		cs = UUID_start
	}

//line dfa.go:4514
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

		//fmt.Println("START");
		half = 0

//line ./uuid-grammar.rl:42
		half |= 1
		i[half] &= INT60_FULL
		i[half] |= uint64(uuidSep2Bits(data[p])) << 60

		goto st2
	tr7:
//line ./uuid-grammar.rl:36
//line ./uuid-grammar.rl:42
		half |= 1
		i[half] &= INT60_FULL
		i[half] |= uint64(uuidSep2Bits(data[p])) << 60

		goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
//line dfa.go:4609
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
//line ./uuid-grammar.rl:31
		digit = 0
		half |= 1

//line ./uuid-grammar.rl:11
		digit = prefixSep2Bits(data[p]) + 4
		i[half] &= INT60_FLAGS | PREFIX_MASKS[digit]

		goto st3
	tr5:
//line ./uuid-grammar.rl:31
		digit = 0
		half |= 1

//line ./uuid-grammar.rl:24
		i[half] &= INT60_FLAGS

//line ./uuid-grammar.rl:16
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
//line ./uuid-grammar.rl:16
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
//line ./uuid-grammar.rl:36

//line ./uuid-grammar.rl:31
		digit = 0
		half |= 1

//line ./uuid-grammar.rl:11
		digit = prefixSep2Bits(data[p]) + 4
		i[half] &= INT60_FLAGS | PREFIX_MASKS[digit]

		goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
//line dfa.go:4703
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

		//fmt.Println("START");
		half = 0

//line ./uuid-grammar.rl:28

//line ./uuid-grammar.rl:11
		digit = prefixSep2Bits(data[p]) + 4
		i[half] &= INT60_FLAGS | PREFIX_MASKS[digit]

		goto st4
	tr9:
//line ./uuid-grammar.rl:16
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
//line dfa.go:4753
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

		//fmt.Println("START");
		half = 0

//line ./uuid-grammar.rl:28

//line ./uuid-grammar.rl:24
		i[half] &= INT60_FLAGS

//line ./uuid-grammar.rl:16
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
//line ./uuid-grammar.rl:16
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
//line dfa.go:4832
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
//line ./uuid-grammar.rl:36

			case 3:
//line ./uuid-grammar.rl:39

			case 5:
//line ./uuid-grammar.rl:36
//line ./uuid-grammar.rl:48
				i[1] = UUID_NAME_UPPER_BITS

//line dfa.go:4899
			}
		}

	_out:
		{
		}
	}

//line dfa.rl:152

	if cs < UUID_first_final || digit > 10 {
		return ERROR_UUID, errors.New(fmt.Sprintf("parse error at pos %d", p))
	} else {
		return UUID{uint128: i}, nil
	}

}
