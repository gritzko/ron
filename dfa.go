//line dfa.rl:1
package RON

import "fmt"
import "errors"

//line dfa.rl:7
//line dfa.go:12
const RON_start int = 22
const RON_first_final int = 22
const RON_error int = 0

const RON_en_main int = 22

//line dfa.rl:8

//line dfa.rl:9

// Parse consumes one op, unless the buffer ends earlier.
func (it *Frame) Parse() {

	//fmt.Println("GO");

	if it.state.p >= len(it.state.data) {
		if !it.state.streaming {
			it.Op = ZERO_OP
			it.state.cs = RON_error
		}
		return
	}

	if it.state.cs == 0 && it.state.p == 0 {

//line dfa.go:40
		{
			it.state.cs = RON_start
		}

//line dfa.rl:26
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
	//fmt.Println("GO!", it.state.cs, "at", p, "with", it.state.data[p]);

//line dfa.go:65
	{
		if p == pe {
			goto _test_eof
		}
		goto _resume

	_again:
		switch it.state.cs {
		case 22:
			goto st22
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
		case 5:
			goto st5
		case 23:
			goto st23
		case 24:
			goto st24
		case 6:
			goto st6
		case 7:
			goto st7
		case 25:
			goto st25
		case 26:
			goto st26
		case 8:
			goto st8
		case 9:
			goto st9
		case 27:
			goto st27
		case 10:
			goto st10
		case 28:
			goto st28
		case 29:
			goto st29
		case 11:
			goto st11
		case 12:
			goto st12
		case 13:
			goto st13
		case 14:
			goto st14
		case 30:
			goto st30
		case 15:
			goto st15
		case 16:
			goto st16
		case 31:
			goto st31
		case 32:
			goto st32
		case 33:
			goto st33
		case 17:
			goto st17
		case 18:
			goto st18
		case 19:
			goto st19
		case 20:
			goto st20
		case 21:
			goto st21
		}

		if p++; p == pe {
			goto _test_eof
		}
	_resume:
		switch it.state.cs {
		case 22:
			goto st_case_22
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
		case 5:
			goto st_case_5
		case 23:
			goto st_case_23
		case 24:
			goto st_case_24
		case 6:
			goto st_case_6
		case 7:
			goto st_case_7
		case 25:
			goto st_case_25
		case 26:
			goto st_case_26
		case 8:
			goto st_case_8
		case 9:
			goto st_case_9
		case 27:
			goto st_case_27
		case 10:
			goto st_case_10
		case 28:
			goto st_case_28
		case 29:
			goto st_case_29
		case 11:
			goto st_case_11
		case 12:
			goto st_case_12
		case 13:
			goto st_case_13
		case 14:
			goto st_case_14
		case 30:
			goto st_case_30
		case 15:
			goto st_case_15
		case 16:
			goto st_case_16
		case 31:
			goto st_case_31
		case 32:
			goto st_case_32
		case 33:
			goto st_case_33
		case 17:
			goto st_case_17
		case 18:
			goto st_case_18
		case 19:
			goto st_case_19
		case 20:
			goto st_case_20
		case 21:
			goto st_case_21
		}
		goto st_out
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
		switch it.state.data[p] {
		case 32:
			goto st1
		case 35:
			goto tr2
		case 42:
			goto tr2
		case 46:
			goto tr76
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
//line ./op-grammar.rl:91

		//fmt.Println("OP START", it.state.cs, "AT", p)
		idx = 0
		if had_end {
			//fmt.Println("BACK")
			p--
			it.state.cs = (RON_start)
			{
				p++
				goto _out
			}
		}

		goto _again
	tr15:
//line ./op-grammar.rl:26
		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

		goto st2
	tr52:
//line ././uuid-grammar.rl:39

//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

		goto st2
	tr60:
//line ././uuid-grammar.rl:36

//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

		goto st2
	tr70:
//line ././uuid-grammar.rl:36

//line ././uuid-grammar.rl:48
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

		goto st2
	tr78:
		it.state.cs = 2
//line ./op-grammar.rl:102
		had_end = true
		//fmt.Println("OP END", it.state.cs, "AT", p)

//line ./op-grammar.rl:91
		//fmt.Println("OP START", it.state.cs, "AT", p)
		idx = 0
		if had_end {
			//fmt.Println("BACK")
			p--
			it.state.cs = (RON_start)
			{
				p++
				goto _out
			}
		}

		goto _again
	tr82:
		it.state.cs = 2
//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:83

//line ./op-grammar.rl:102
		had_end = true
		//fmt.Println("OP END", it.state.cs, "AT", p)

//line ./op-grammar.rl:91
		//fmt.Println("OP START", it.state.cs, "AT", p)
		idx = 0
		if had_end {
			//fmt.Println("BACK")
			p--
			it.state.cs = (RON_start)
			{
				p++
				goto _out
			}
		}

		goto _again
	tr90:
		it.state.cs = 2
//line ./op-grammar.rl:83
//line ./op-grammar.rl:102
		had_end = true
		//fmt.Println("OP END", it.state.cs, "AT", p)

//line ./op-grammar.rl:91
		//fmt.Println("OP START", it.state.cs, "AT", p)
		idx = 0
		if had_end {
			//fmt.Println("BACK")
			p--
			it.state.cs = (RON_start)
			{
				p++
				goto _out
			}
		}

		goto _again
	tr98:
		it.state.cs = 2
//line ./op-grammar.rl:55
		//fmt.Println("INT END");
		i[1] |= ATOM_INT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:83

//line ./op-grammar.rl:102
		had_end = true
		//fmt.Println("OP END", it.state.cs, "AT", p)

//line ./op-grammar.rl:91
		//fmt.Println("OP START", it.state.cs, "AT", p)
		idx = 0
		if had_end {
			//fmt.Println("BACK")
			p--
			it.state.cs = (RON_start)
			{
				p++
				goto _out
			}
		}

		goto _again
	tr108:
		it.state.cs = 2
//line ././uuid-grammar.rl:39

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:83

//line ./op-grammar.rl:102
		had_end = true
		//fmt.Println("OP END", it.state.cs, "AT", p)

//line ./op-grammar.rl:91
		//fmt.Println("OP START", it.state.cs, "AT", p)
		idx = 0
		if had_end {
			//fmt.Println("BACK")
			p--
			it.state.cs = (RON_start)
			{
				p++
				goto _out
			}
		}

		goto _again
	tr117:
		it.state.cs = 2
//line ./op-grammar.rl:60
//line ./op-grammar.rl:62
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:83

//line ./op-grammar.rl:102
		had_end = true
		//fmt.Println("OP END", it.state.cs, "AT", p)

//line ./op-grammar.rl:91
		//fmt.Println("OP START", it.state.cs, "AT", p)
		idx = 0
		if had_end {
			//fmt.Println("BACK")
			p--
			it.state.cs = (RON_start)
			{
				p++
				goto _out
			}
		}

		goto _again
	tr126:
		it.state.cs = 2
//line ./op-grammar.rl:62
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:83

//line ./op-grammar.rl:102
		had_end = true
		//fmt.Println("OP END", it.state.cs, "AT", p)

//line ./op-grammar.rl:91
		//fmt.Println("OP START", it.state.cs, "AT", p)
		idx = 0
		if had_end {
			//fmt.Println("BACK")
			p--
			it.state.cs = (RON_start)
			{
				p++
				goto _out
			}
		}

		goto _again
	tr134:
		it.state.cs = 2
//line ././uuid-grammar.rl:36

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:83

//line ./op-grammar.rl:102
		had_end = true
		//fmt.Println("OP END", it.state.cs, "AT", p)

//line ./op-grammar.rl:91
		//fmt.Println("OP START", it.state.cs, "AT", p)
		idx = 0
		if had_end {
			//fmt.Println("BACK")
			p--
			it.state.cs = (RON_start)
			{
				p++
				goto _out
			}
		}

		goto _again
	tr145:
		it.state.cs = 2
//line ././uuid-grammar.rl:36

//line ././uuid-grammar.rl:48
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:83

//line ./op-grammar.rl:102
		had_end = true
		//fmt.Println("OP END", it.state.cs, "AT", p)

//line ./op-grammar.rl:91
		//fmt.Println("OP START", it.state.cs, "AT", p)
		idx = 0
		if had_end {
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
//line dfa.go:597
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
//line dfa.go:667
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
	tr61:
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
//line dfa.go:772
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
	tr50:
//line ././uuid-grammar.rl:39

//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

		goto st5
	tr58:
//line ././uuid-grammar.rl:36

//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

		goto st5
	tr68:
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
//line dfa.go:886
		switch it.state.data[p] {
		case 32:
			goto st5
		case 33:
			goto tr23
		case 35:
			goto st2
		case 39:
			goto tr25
		case 42:
			goto st2
		case 44:
			goto tr23
		case 58:
			goto st2
		case 59:
			goto tr23
		case 61:
			goto tr26
		case 62:
			goto tr27
		case 63:
			goto tr23
		case 64:
			goto st2
		case 94:
			goto tr28
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

//line ./op-grammar.rl:107
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:86
		//fmt.Println("TERM", fc, it.state.cs, "AT", p);
		it.term = termSep2Bits(it.state.data[p])

		goto st23
	tr23:
//line ./op-grammar.rl:107
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:86
		//fmt.Println("TERM", fc, it.state.cs, "AT", p);
		it.term = termSep2Bits(it.state.data[p])

		goto st23
	tr51:
//line ././uuid-grammar.rl:39

//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:107
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:86
		//fmt.Println("TERM", fc, it.state.cs, "AT", p);
		it.term = termSep2Bits(it.state.data[p])

		goto st23
	tr59:
//line ././uuid-grammar.rl:36

//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:107
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:86
		//fmt.Println("TERM", fc, it.state.cs, "AT", p);
		it.term = termSep2Bits(it.state.data[p])

		goto st23
	tr69:
//line ././uuid-grammar.rl:36

//line ././uuid-grammar.rl:48
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:107
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:86
		//fmt.Println("TERM", fc, it.state.cs, "AT", p);
		it.term = termSep2Bits(it.state.data[p])

		goto st23
	tr77:
//line ./op-grammar.rl:102
		had_end = true
		//fmt.Println("OP END", it.state.cs, "AT", p)

		goto st23
	tr81:
//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:83

//line ./op-grammar.rl:86

		//fmt.Println("TERM", fc, it.state.cs, "AT", p);
		it.term = termSep2Bits(it.state.data[p])

		goto st23
	tr89:
//line ./op-grammar.rl:83
//line ./op-grammar.rl:86

		//fmt.Println("TERM", fc, it.state.cs, "AT", p);
		it.term = termSep2Bits(it.state.data[p])

		goto st23
	tr97:
//line ./op-grammar.rl:55
		//fmt.Println("INT END");
		i[1] |= ATOM_INT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:83

//line ./op-grammar.rl:86

		//fmt.Println("TERM", fc, it.state.cs, "AT", p);
		it.term = termSep2Bits(it.state.data[p])

		goto st23
	tr107:
//line ././uuid-grammar.rl:39

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:83

//line ./op-grammar.rl:86

		//fmt.Println("TERM", fc, it.state.cs, "AT", p);
		it.term = termSep2Bits(it.state.data[p])

		goto st23
	tr116:
//line ./op-grammar.rl:60
//line ./op-grammar.rl:62
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:83

//line ./op-grammar.rl:86

		//fmt.Println("TERM", fc, it.state.cs, "AT", p);
		it.term = termSep2Bits(it.state.data[p])

		goto st23
	tr125:
//line ./op-grammar.rl:62
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:83

//line ./op-grammar.rl:86

		//fmt.Println("TERM", fc, it.state.cs, "AT", p);
		it.term = termSep2Bits(it.state.data[p])

		goto st23
	tr133:
//line ././uuid-grammar.rl:36

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:83

//line ./op-grammar.rl:86

		//fmt.Println("TERM", fc, it.state.cs, "AT", p);
		it.term = termSep2Bits(it.state.data[p])

		goto st23
	tr144:
//line ././uuid-grammar.rl:36

//line ././uuid-grammar.rl:48
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:83

//line ./op-grammar.rl:86

		//fmt.Println("TERM", fc, it.state.cs, "AT", p);
		it.term = termSep2Bits(it.state.data[p])

		goto st23
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
//line dfa.go:1194
		switch it.state.data[p] {
		case 32:
			goto tr77
		case 35:
			goto tr78
		case 42:
			goto tr78
		case 46:
			goto tr79
		case 58:
			goto tr78
		case 64:
			goto tr78
		}
		if 9 <= it.state.data[p] && it.state.data[p] <= 13 {
			goto tr77
		}
		goto st0
	tr76:
//line ./op-grammar.rl:115

		//fmt.Println("FRAME END")
		it.state.streaming = false

		goto st24
	tr79:
//line ./op-grammar.rl:102
		had_end = true
		//fmt.Println("OP END", it.state.cs, "AT", p)

//line ./op-grammar.rl:115

		//fmt.Println("FRAME END")
		it.state.streaming = false

		goto st24
	tr84:
//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:83

//line ./op-grammar.rl:102
		had_end = true
		//fmt.Println("OP END", it.state.cs, "AT", p)

//line ./op-grammar.rl:115

		//fmt.Println("FRAME END")
		it.state.streaming = false

		goto st24
	tr92:
//line ./op-grammar.rl:83
//line ./op-grammar.rl:102
		had_end = true
		//fmt.Println("OP END", it.state.cs, "AT", p)

//line ./op-grammar.rl:115

		//fmt.Println("FRAME END")
		it.state.streaming = false

		goto st24
	tr100:
//line ./op-grammar.rl:55
		//fmt.Println("INT END");
		i[1] |= ATOM_INT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:83

//line ./op-grammar.rl:102
		had_end = true
		//fmt.Println("OP END", it.state.cs, "AT", p)

//line ./op-grammar.rl:115

		//fmt.Println("FRAME END")
		it.state.streaming = false

		goto st24
	tr110:
//line ././uuid-grammar.rl:39

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:83

//line ./op-grammar.rl:102
		had_end = true
		//fmt.Println("OP END", it.state.cs, "AT", p)

//line ./op-grammar.rl:115

		//fmt.Println("FRAME END")
		it.state.streaming = false

		goto st24
	tr119:
//line ./op-grammar.rl:60
//line ./op-grammar.rl:62
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:83

//line ./op-grammar.rl:102
		had_end = true
		//fmt.Println("OP END", it.state.cs, "AT", p)

//line ./op-grammar.rl:115

		//fmt.Println("FRAME END")
		it.state.streaming = false

		goto st24
	tr128:
//line ./op-grammar.rl:62
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:83

//line ./op-grammar.rl:102
		had_end = true
		//fmt.Println("OP END", it.state.cs, "AT", p)

//line ./op-grammar.rl:115

		//fmt.Println("FRAME END")
		it.state.streaming = false

		goto st24
	tr138:
//line ././uuid-grammar.rl:36

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:83

//line ./op-grammar.rl:102
		had_end = true
		//fmt.Println("OP END", it.state.cs, "AT", p)

//line ./op-grammar.rl:115

		//fmt.Println("FRAME END")
		it.state.streaming = false

		goto st24
	tr147:
//line ././uuid-grammar.rl:36

//line ././uuid-grammar.rl:48
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:83

//line ./op-grammar.rl:102
		had_end = true
		//fmt.Println("OP END", it.state.cs, "AT", p)

//line ./op-grammar.rl:115

		//fmt.Println("FRAME END")
		it.state.streaming = false

		goto st24
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
//line dfa.go:1428
		goto st0
	tr16:
//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:107
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:80
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st6
	tr25:
//line ./op-grammar.rl:107
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:80
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st6
	tr53:
//line ././uuid-grammar.rl:39

//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:107
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:80
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st6
	tr62:
//line ././uuid-grammar.rl:36

//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:107
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:80
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st6
	tr71:
//line ././uuid-grammar.rl:36

//line ././uuid-grammar.rl:48
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:107
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:80
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st6
	tr83:
//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st6
	tr91:
//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st6
	tr99:
//line ./op-grammar.rl:55

		//fmt.Println("INT END");
		i[1] |= ATOM_INT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st6
	tr109:
//line ././uuid-grammar.rl:39

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st6
	tr118:
//line ./op-grammar.rl:60

//line ./op-grammar.rl:62
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st6
	tr127:
//line ./op-grammar.rl:62
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st6
	tr136:
//line ././uuid-grammar.rl:36

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st6
	tr146:
//line ././uuid-grammar.rl:36

//line ././uuid-grammar.rl:48
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
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
//line dfa.go:1694
		switch it.state.data[p] {
		case 10:
			goto st0
		case 13:
			goto st0
		case 39:
			goto tr30
		case 92:
			goto tr31
		}
		goto tr29
	tr29:
//line ./op-grammar.rl:66
		i[0] = uint64(p)

		goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
//line dfa.go:1717
		switch it.state.data[p] {
		case 10:
			goto st0
		case 13:
			goto st0
		case 39:
			goto tr33
		case 92:
			goto st17
		}
		goto st7
	tr30:
//line ./op-grammar.rl:66
		i[0] = uint64(p)

//line ./op-grammar.rl:69
		//fmt.Println("STRING");
		i[1] = uint64(p) | ATOM_STRING_62

		goto st25
	tr33:
//line ./op-grammar.rl:69
		//fmt.Println("STRING");
		i[1] = uint64(p) | ATOM_STRING_62

		goto st25
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
//line dfa.go:1752
		switch it.state.data[p] {
		case 32:
			goto tr80
		case 33:
			goto tr81
		case 35:
			goto tr82
		case 39:
			goto tr83
		case 42:
			goto tr82
		case 44:
			goto tr81
		case 46:
			goto tr84
		case 58:
			goto tr82
		case 59:
			goto tr81
		case 61:
			goto tr85
		case 62:
			goto tr86
		case 63:
			goto tr81
		case 64:
			goto tr82
		case 94:
			goto tr87
		}
		if 9 <= it.state.data[p] && it.state.data[p] <= 13 {
			goto tr80
		}
		goto st0
	tr80:
//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:83

//line ./op-grammar.rl:102
		had_end = true
		//fmt.Println("OP END", it.state.cs, "AT", p)

		goto st26
	tr88:
//line ./op-grammar.rl:83
//line ./op-grammar.rl:102
		had_end = true
		//fmt.Println("OP END", it.state.cs, "AT", p)

		goto st26
	tr96:
//line ./op-grammar.rl:55
		//fmt.Println("INT END");
		i[1] |= ATOM_INT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:83

//line ./op-grammar.rl:102
		had_end = true
		//fmt.Println("OP END", it.state.cs, "AT", p)

		goto st26
	tr106:
//line ././uuid-grammar.rl:39

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:83

//line ./op-grammar.rl:102
		had_end = true
		//fmt.Println("OP END", it.state.cs, "AT", p)

		goto st26
	tr115:
//line ./op-grammar.rl:60
//line ./op-grammar.rl:62
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:83

//line ./op-grammar.rl:102
		had_end = true
		//fmt.Println("OP END", it.state.cs, "AT", p)

		goto st26
	tr124:
//line ./op-grammar.rl:62
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:83

//line ./op-grammar.rl:102
		had_end = true
		//fmt.Println("OP END", it.state.cs, "AT", p)

		goto st26
	tr132:
//line ././uuid-grammar.rl:36

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:83

//line ./op-grammar.rl:102
		had_end = true
		//fmt.Println("OP END", it.state.cs, "AT", p)

		goto st26
	tr143:
//line ././uuid-grammar.rl:36

//line ././uuid-grammar.rl:48
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:83

//line ./op-grammar.rl:102
		had_end = true
		//fmt.Println("OP END", it.state.cs, "AT", p)

		goto st26
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
//line dfa.go:1943
		switch it.state.data[p] {
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
			goto tr93
		case 62:
			goto tr94
		case 63:
			goto tr89
		case 64:
			goto tr90
		case 94:
			goto tr95
		}
		if 9 <= it.state.data[p] && it.state.data[p] <= 13 {
			goto tr88
		}
		goto st0
	tr19:
//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:107
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:80
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st8
	tr26:
//line ./op-grammar.rl:107
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:80
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st8
	tr55:
//line ././uuid-grammar.rl:39

//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:107
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:80
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st8
	tr65:
//line ././uuid-grammar.rl:36

//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:107
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:80
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st8
	tr73:
//line ././uuid-grammar.rl:36

//line ././uuid-grammar.rl:48
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:107
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:80
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st8
	tr85:
//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st8
	tr93:
//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st8
	tr101:
//line ./op-grammar.rl:55

		//fmt.Println("INT END");
		i[1] |= ATOM_INT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st8
	tr112:
//line ././uuid-grammar.rl:39

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st8
	tr120:
//line ./op-grammar.rl:60

//line ./op-grammar.rl:62
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st8
	tr129:
//line ./op-grammar.rl:62
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st8
	tr140:
//line ././uuid-grammar.rl:36

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st8
	tr149:
//line ././uuid-grammar.rl:36

//line ././uuid-grammar.rl:48
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
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
//line dfa.go:2242
		switch it.state.data[p] {
		case 32:
			goto st8
		case 43:
			goto tr36
		case 45:
			goto tr36
		}
		switch {
		case it.state.data[p] > 13:
			if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
				goto tr37
			}
		case it.state.data[p] >= 9:
			goto st8
		}
		goto st0
	tr36:
//line ./op-grammar.rl:42

		//fmt.Println("INT START");

//line ./op-grammar.rl:45
		if it.state.data[p] == '-' {
			i[1] |= 1
		}

		goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
//line dfa.go:2277
		if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
			goto tr38
		}
		goto st0
	tr37:
//line ./op-grammar.rl:42

		//fmt.Println("INT START");

//line ./op-grammar.rl:50

		//fmt.Println("INT DGT", p, fc);
		i[0] *= 10
		i[0] += uint64(int(it.state.data[p]) - int('0'))

		goto st27
	tr38:
//line ./op-grammar.rl:50
		//fmt.Println("INT DGT", p, fc);
		i[0] *= 10
		i[0] += uint64(int(it.state.data[p]) - int('0'))

		goto st27
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
//line dfa.go:2307
		switch it.state.data[p] {
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
			goto tr101
		case 62:
			goto tr102
		case 63:
			goto tr97
		case 64:
			goto tr98
		case 94:
			goto tr103
		}
		switch {
		case it.state.data[p] > 13:
			if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
				goto tr38
			}
		case it.state.data[p] >= 9:
			goto tr96
		}
		goto st0
	tr20:
//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:107
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:80
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st10
	tr27:
//line ./op-grammar.rl:107
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:80
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st10
	tr56:
//line ././uuid-grammar.rl:39

//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:107
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:80
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st10
	tr66:
//line ././uuid-grammar.rl:36

//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:107
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:80
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st10
	tr74:
//line ././uuid-grammar.rl:36

//line ././uuid-grammar.rl:48
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:107
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:80
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st10
	tr86:
//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st10
	tr94:
//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st10
	tr102:
//line ./op-grammar.rl:55

		//fmt.Println("INT END");
		i[1] |= ATOM_INT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st10
	tr113:
//line ././uuid-grammar.rl:39

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st10
	tr121:
//line ./op-grammar.rl:60

//line ./op-grammar.rl:62
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st10
	tr130:
//line ./op-grammar.rl:62
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st10
	tr141:
//line ././uuid-grammar.rl:36

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st10
	tr150:
//line ././uuid-grammar.rl:36

//line ././uuid-grammar.rl:48
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
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
//line dfa.go:2611
		switch it.state.data[p] {
		case 32:
			goto st10
		case 43:
			goto tr40
		case 45:
			goto tr40
		case 91:
			goto tr41
		case 93:
			goto tr41
		case 95:
			goto tr42
		case 123:
			goto tr41
		case 125:
			goto tr41
		case 126:
			goto tr42
		}
		switch {
		case it.state.data[p] < 40:
			switch {
			case it.state.data[p] > 13:
				if 36 <= it.state.data[p] && it.state.data[p] <= 37 {
					goto tr40
				}
			case it.state.data[p] >= 9:
				goto st10
			}
		case it.state.data[p] > 41:
			switch {
			case it.state.data[p] < 65:
				if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
					goto tr42
				}
			case it.state.data[p] > 90:
				if 97 <= it.state.data[p] && it.state.data[p] <= 122 {
					goto tr42
				}
			default:
				goto tr42
			}
		default:
			goto tr41
		}
		goto st0
	tr40:
//line ././uuid-grammar.rl:5

		//fmt.Println("START");
		half = 0

//line ././uuid-grammar.rl:42
		half |= 1
		i[half] &= INT60_FULL
		i[half] |= uint64(uuidSep2Bits(it.state.data[p])) << 60

		goto st28
	tr135:
//line ././uuid-grammar.rl:36
//line ././uuid-grammar.rl:42
		half |= 1
		i[half] &= INT60_FULL
		i[half] |= uint64(uuidSep2Bits(it.state.data[p])) << 60

		goto st28
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
//line dfa.go:2688
		switch it.state.data[p] {
		case 32:
			goto tr80
		case 33:
			goto tr81
		case 35:
			goto tr82
		case 39:
			goto tr83
		case 42:
			goto tr82
		case 44:
			goto tr81
		case 46:
			goto tr84
		case 58:
			goto tr82
		case 59:
			goto tr81
		case 61:
			goto tr85
		case 62:
			goto tr86
		case 63:
			goto tr81
		case 64:
			goto tr82
		case 91:
			goto tr104
		case 93:
			goto tr104
		case 94:
			goto tr87
		case 95:
			goto tr105
		case 123:
			goto tr104
		case 125:
			goto tr104
		case 126:
			goto tr105
		}
		switch {
		case it.state.data[p] < 48:
			switch {
			case it.state.data[p] > 13:
				if 40 <= it.state.data[p] && it.state.data[p] <= 41 {
					goto tr104
				}
			case it.state.data[p] >= 9:
				goto tr80
			}
		case it.state.data[p] > 57:
			switch {
			case it.state.data[p] > 90:
				if 97 <= it.state.data[p] && it.state.data[p] <= 122 {
					goto tr105
				}
			case it.state.data[p] >= 65:
				goto tr105
			}
		default:
			goto tr105
		}
		goto st0
	tr104:
//line ././uuid-grammar.rl:31
		digit = 0
		half |= 1

//line ././uuid-grammar.rl:11
		digit = prefixSep2Bits(it.state.data[p]) + 4
		i[half] &= INT60_FLAGS | PREFIX_MASKS[digit]

		goto st29
	tr105:
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
				it.state.cs = 29
				goto _out
			}
		}

		goto st29
	tr111:
//line ././uuid-grammar.rl:16
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
	tr137:
//line ././uuid-grammar.rl:36

//line ././uuid-grammar.rl:31
		digit = 0
		half |= 1

//line ././uuid-grammar.rl:11
		digit = prefixSep2Bits(it.state.data[p]) + 4
		i[half] &= INT60_FLAGS | PREFIX_MASKS[digit]

		goto st29
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
//line dfa.go:2815
		switch it.state.data[p] {
		case 32:
			goto tr106
		case 33:
			goto tr107
		case 35:
			goto tr108
		case 39:
			goto tr109
		case 42:
			goto tr108
		case 44:
			goto tr107
		case 46:
			goto tr110
		case 58:
			goto tr108
		case 59:
			goto tr107
		case 61:
			goto tr112
		case 62:
			goto tr113
		case 63:
			goto tr107
		case 64:
			goto tr108
		case 94:
			goto tr114
		case 95:
			goto tr111
		case 126:
			goto tr111
		}
		switch {
		case it.state.data[p] < 48:
			if 9 <= it.state.data[p] && it.state.data[p] <= 13 {
				goto tr106
			}
		case it.state.data[p] > 57:
			switch {
			case it.state.data[p] > 90:
				if 97 <= it.state.data[p] && it.state.data[p] <= 122 {
					goto tr111
				}
			case it.state.data[p] >= 65:
				goto tr111
			}
		default:
			goto tr111
		}
		goto st0
	tr21:
//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:107
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:80
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st11
	tr28:
//line ./op-grammar.rl:107
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:80
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st11
	tr57:
//line ././uuid-grammar.rl:39

//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:107
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:80
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st11
	tr67:
//line ././uuid-grammar.rl:36

//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:107
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:80
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st11
	tr75:
//line ././uuid-grammar.rl:36

//line ././uuid-grammar.rl:48
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:26

		//fmt.Println("UUID#", idx);
		it.uuids[idx] = UUID{uint128: i}
		idx++

//line ./op-grammar.rl:107
		if it.term != TERM_RAW {
			it.term = TERM_REDUCED
		}
		it.Reset()
		it.frame = it.state.data

//line ./op-grammar.rl:80
		idx = 0

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st11
	tr87:
//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st11
	tr95:
//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st11
	tr103:
//line ./op-grammar.rl:55

		//fmt.Println("INT END");
		i[1] |= ATOM_INT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st11
	tr114:
//line ././uuid-grammar.rl:39

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st11
	tr123:
//line ./op-grammar.rl:60

//line ./op-grammar.rl:62
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st11
	tr131:
//line ./op-grammar.rl:62
		i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st11
	tr142:
//line ././uuid-grammar.rl:36

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
		it.AddAtom(i)

//line ./op-grammar.rl:32
		digit = 0
		i = uint128{0, 0}

		goto st11
	tr151:
//line ././uuid-grammar.rl:36

//line ././uuid-grammar.rl:48
		i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:36

		// TODO max size for int/float/string
		//fmt.Println("ADD ATOM", i);
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
//line dfa.go:3132
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
				goto st13
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
			goto st13
		}
		goto st0
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		if it.state.data[p] == 46 {
			goto st14
		}
		if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
			goto st13
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
			goto tr120
		case 62:
			goto tr121
		case 63:
			goto tr116
		case 64:
			goto tr117
		case 69:
			goto tr122
		case 94:
			goto tr123
		case 101:
			goto tr122
		}
		switch {
		case it.state.data[p] > 13:
			if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
				goto st30
			}
		case it.state.data[p] >= 9:
			goto tr115
		}
		goto st0
	tr122:
//line ./op-grammar.rl:60

		goto st15
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
//line dfa.go:3238
		switch it.state.data[p] {
		case 43:
			goto st16
		case 45:
			goto st16
		}
		if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
			goto st31
		}
		goto st0
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
			goto st31
		}
		goto st0
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
		switch it.state.data[p] {
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
			goto tr129
		case 62:
			goto tr130
		case 63:
			goto tr125
		case 64:
			goto tr126
		case 94:
			goto tr131
		}
		switch {
		case it.state.data[p] > 13:
			if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
				goto st31
			}
		case it.state.data[p] >= 9:
			goto tr124
		}
		goto st0
	tr41:
//line ././uuid-grammar.rl:5

		//fmt.Println("START");
		half = 0

//line ././uuid-grammar.rl:28

//line ././uuid-grammar.rl:11
		digit = prefixSep2Bits(it.state.data[p]) + 4
		i[half] &= INT60_FLAGS | PREFIX_MASKS[digit]

		goto st32
	tr139:
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
//line dfa.go:3332
		switch it.state.data[p] {
		case 32:
			goto tr132
		case 33:
			goto tr133
		case 35:
			goto tr134
		case 39:
			goto tr136
		case 42:
			goto tr134
		case 44:
			goto tr133
		case 46:
			goto tr138
		case 58:
			goto tr134
		case 59:
			goto tr133
		case 61:
			goto tr140
		case 62:
			goto tr141
		case 63:
			goto tr133
		case 64:
			goto tr134
		case 91:
			goto tr137
		case 93:
			goto tr137
		case 94:
			goto tr142
		case 95:
			goto tr139
		case 123:
			goto tr137
		case 125:
			goto tr137
		case 126:
			goto tr139
		}
		switch {
		case it.state.data[p] < 43:
			switch {
			case it.state.data[p] < 36:
				if 9 <= it.state.data[p] && it.state.data[p] <= 13 {
					goto tr132
				}
			case it.state.data[p] > 37:
				if 40 <= it.state.data[p] && it.state.data[p] <= 41 {
					goto tr137
				}
			default:
				goto tr135
			}
		case it.state.data[p] > 45:
			switch {
			case it.state.data[p] < 65:
				if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
					goto tr139
				}
			case it.state.data[p] > 90:
				if 97 <= it.state.data[p] && it.state.data[p] <= 122 {
					goto tr139
				}
			default:
				goto tr139
			}
		default:
			goto tr135
		}
		goto st0
	tr42:
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
				it.state.cs = 33
				goto _out
			}
		}

		goto st33
	tr148:
//line ././uuid-grammar.rl:16
		i[half] |= uint64(ABC[it.state.data[p]]) << DIGIT_OFFSETS[digit]
		digit++
		if digit > 10 {
			{
				p++
				it.state.cs = 33
				goto _out
			}
		}

		goto st33
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
//line dfa.go:3443
		switch it.state.data[p] {
		case 32:
			goto tr143
		case 33:
			goto tr144
		case 35:
			goto tr145
		case 39:
			goto tr146
		case 42:
			goto tr145
		case 44:
			goto tr144
		case 46:
			goto tr147
		case 58:
			goto tr145
		case 59:
			goto tr144
		case 61:
			goto tr149
		case 62:
			goto tr150
		case 63:
			goto tr144
		case 64:
			goto tr145
		case 91:
			goto tr137
		case 93:
			goto tr137
		case 94:
			goto tr151
		case 95:
			goto tr148
		case 123:
			goto tr137
		case 125:
			goto tr137
		case 126:
			goto tr148
		}
		switch {
		case it.state.data[p] < 43:
			switch {
			case it.state.data[p] < 36:
				if 9 <= it.state.data[p] && it.state.data[p] <= 13 {
					goto tr143
				}
			case it.state.data[p] > 37:
				if 40 <= it.state.data[p] && it.state.data[p] <= 41 {
					goto tr137
				}
			default:
				goto tr135
			}
		case it.state.data[p] > 45:
			switch {
			case it.state.data[p] < 65:
				if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
					goto tr148
				}
			case it.state.data[p] > 90:
				if 97 <= it.state.data[p] && it.state.data[p] <= 122 {
					goto tr148
				}
			default:
				goto tr148
			}
		default:
			goto tr135
		}
		goto st0
	tr31:
//line ./op-grammar.rl:66
		i[0] = uint64(p)

		goto st17
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
//line dfa.go:3528
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

		goto st18
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
				it.state.cs = 18
				goto _out
			}
		}

		goto st18
	tr54:
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
	tr63:
//line ././uuid-grammar.rl:36

//line ././uuid-grammar.rl:31
		digit = 0
		half |= 1

//line ././uuid-grammar.rl:11
		digit = prefixSep2Bits(it.state.data[p]) + 4
		i[half] &= INT60_FLAGS | PREFIX_MASKS[digit]

		goto st18
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
//line dfa.go:3597
		switch it.state.data[p] {
		case 32:
			goto tr50
		case 33:
			goto tr51
		case 35:
			goto tr52
		case 39:
			goto tr53
		case 42:
			goto tr52
		case 44:
			goto tr51
		case 58:
			goto tr52
		case 59:
			goto tr51
		case 61:
			goto tr55
		case 62:
			goto tr56
		case 63:
			goto tr51
		case 64:
			goto tr52
		case 94:
			goto tr57
		case 95:
			goto tr54
		case 126:
			goto tr54
		}
		switch {
		case it.state.data[p] < 48:
			if 9 <= it.state.data[p] && it.state.data[p] <= 13 {
				goto tr50
			}
		case it.state.data[p] > 57:
			switch {
			case it.state.data[p] > 90:
				if 97 <= it.state.data[p] && it.state.data[p] <= 122 {
					goto tr54
				}
			case it.state.data[p] >= 65:
				goto tr54
			}
		default:
			goto tr54
		}
		goto st0
	tr5:
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

		goto st19
	tr64:
//line ././uuid-grammar.rl:16
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
//line dfa.go:3708
		switch it.state.data[p] {
		case 32:
			goto tr58
		case 33:
			goto tr59
		case 35:
			goto tr60
		case 39:
			goto tr62
		case 42:
			goto tr60
		case 44:
			goto tr59
		case 58:
			goto tr60
		case 59:
			goto tr59
		case 61:
			goto tr65
		case 62:
			goto tr66
		case 63:
			goto tr59
		case 64:
			goto tr60
		case 91:
			goto tr63
		case 93:
			goto tr63
		case 94:
			goto tr67
		case 95:
			goto tr64
		case 123:
			goto tr63
		case 125:
			goto tr63
		case 126:
			goto tr64
		}
		switch {
		case it.state.data[p] < 43:
			switch {
			case it.state.data[p] < 36:
				if 9 <= it.state.data[p] && it.state.data[p] <= 13 {
					goto tr58
				}
			case it.state.data[p] > 37:
				if 40 <= it.state.data[p] && it.state.data[p] <= 41 {
					goto tr63
				}
			default:
				goto tr61
			}
		case it.state.data[p] > 45:
			switch {
			case it.state.data[p] < 65:
				if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
					goto tr64
				}
			case it.state.data[p] > 90:
				if 97 <= it.state.data[p] && it.state.data[p] <= 122 {
					goto tr64
				}
			default:
				goto tr64
			}
		default:
			goto tr61
		}
		goto st0
	tr6:
		it.state.cs = 20
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
				it.state.cs = 20
				goto _out
			}
		}

		goto st20
	tr72:
//line ././uuid-grammar.rl:16
		i[half] |= uint64(ABC[it.state.data[p]]) << DIGIT_OFFSETS[digit]
		digit++
		if digit > 10 {
			{
				p++
				it.state.cs = 20
				goto _out
			}
		}

		goto st20
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
//line dfa.go:3854
		switch it.state.data[p] {
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
		case 91:
			goto tr63
		case 93:
			goto tr63
		case 94:
			goto tr75
		case 95:
			goto tr72
		case 123:
			goto tr63
		case 125:
			goto tr63
		case 126:
			goto tr72
		}
		switch {
		case it.state.data[p] < 43:
			switch {
			case it.state.data[p] < 36:
				if 9 <= it.state.data[p] && it.state.data[p] <= 13 {
					goto tr68
				}
			case it.state.data[p] > 37:
				if 40 <= it.state.data[p] && it.state.data[p] <= 41 {
					goto tr63
				}
			default:
				goto tr61
			}
		case it.state.data[p] > 45:
			switch {
			case it.state.data[p] < 65:
				if 48 <= it.state.data[p] && it.state.data[p] <= 57 {
					goto tr72
				}
			case it.state.data[p] > 90:
				if 97 <= it.state.data[p] && it.state.data[p] <= 122 {
					goto tr72
				}
			default:
				goto tr72
			}
		default:
			goto tr61
		}
		goto st0
	tr7:
		it.state.cs = 21
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

		goto st21
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
//line dfa.go:3962
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
	_test_eof22:
		it.state.cs = 22
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
	_test_eof5:
		it.state.cs = 5
		goto _test_eof
	_test_eof23:
		it.state.cs = 23
		goto _test_eof
	_test_eof24:
		it.state.cs = 24
		goto _test_eof
	_test_eof6:
		it.state.cs = 6
		goto _test_eof
	_test_eof7:
		it.state.cs = 7
		goto _test_eof
	_test_eof25:
		it.state.cs = 25
		goto _test_eof
	_test_eof26:
		it.state.cs = 26
		goto _test_eof
	_test_eof8:
		it.state.cs = 8
		goto _test_eof
	_test_eof9:
		it.state.cs = 9
		goto _test_eof
	_test_eof27:
		it.state.cs = 27
		goto _test_eof
	_test_eof10:
		it.state.cs = 10
		goto _test_eof
	_test_eof28:
		it.state.cs = 28
		goto _test_eof
	_test_eof29:
		it.state.cs = 29
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
	_test_eof14:
		it.state.cs = 14
		goto _test_eof
	_test_eof30:
		it.state.cs = 30
		goto _test_eof
	_test_eof15:
		it.state.cs = 15
		goto _test_eof
	_test_eof16:
		it.state.cs = 16
		goto _test_eof
	_test_eof31:
		it.state.cs = 31
		goto _test_eof
	_test_eof32:
		it.state.cs = 32
		goto _test_eof
	_test_eof33:
		it.state.cs = 33
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
	_test_eof21:
		it.state.cs = 21
		goto _test_eof

	_test_eof:
		{
		}
		if p == eof {
			switch it.state.cs {
			case 23:
//line ./op-grammar.rl:102
				had_end = true
				//fmt.Println("OP END", it.state.cs, "AT", p)

			case 26:
//line ./op-grammar.rl:83
//line ./op-grammar.rl:102
				had_end = true
				//fmt.Println("OP END", it.state.cs, "AT", p)

			case 25, 28:
//line ./op-grammar.rl:36
				// TODO max size for int/float/string
				//fmt.Println("ADD ATOM", i);
				it.AddAtom(i)

//line ./op-grammar.rl:83

//line ./op-grammar.rl:102
				had_end = true
				//fmt.Println("OP END", it.state.cs, "AT", p)

			case 32:
//line ././uuid-grammar.rl:36

//line ./op-grammar.rl:36

				// TODO max size for int/float/string
				//fmt.Println("ADD ATOM", i);
				it.AddAtom(i)

//line ./op-grammar.rl:83

//line ./op-grammar.rl:102
				had_end = true
				//fmt.Println("OP END", it.state.cs, "AT", p)

			case 29:
//line ././uuid-grammar.rl:39

//line ./op-grammar.rl:36

				// TODO max size for int/float/string
				//fmt.Println("ADD ATOM", i);
				it.AddAtom(i)

//line ./op-grammar.rl:83

//line ./op-grammar.rl:102
				had_end = true
				//fmt.Println("OP END", it.state.cs, "AT", p)

			case 27:
//line ./op-grammar.rl:55
				//fmt.Println("INT END");
				i[1] |= ATOM_INT_62

//line ./op-grammar.rl:36
				// TODO max size for int/float/string
				//fmt.Println("ADD ATOM", i);
				it.AddAtom(i)

//line ./op-grammar.rl:83

//line ./op-grammar.rl:102
				had_end = true
				//fmt.Println("OP END", it.state.cs, "AT", p)

			case 31:
//line ./op-grammar.rl:62
				i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
				// TODO max size for int/float/string
				//fmt.Println("ADD ATOM", i);
				it.AddAtom(i)

//line ./op-grammar.rl:83

//line ./op-grammar.rl:102
				had_end = true
				//fmt.Println("OP END", it.state.cs, "AT", p)

			case 33:
//line ././uuid-grammar.rl:36

//line ././uuid-grammar.rl:48
				i[1] = UUID_NAME_UPPER_BITS

//line ./op-grammar.rl:36

				// TODO max size for int/float/string
				//fmt.Println("ADD ATOM", i);
				it.AddAtom(i)

//line ./op-grammar.rl:83

//line ./op-grammar.rl:102
				had_end = true
				//fmt.Println("OP END", it.state.cs, "AT", p)

			case 30:
//line ./op-grammar.rl:60
//line ./op-grammar.rl:62
				i[1] |= ATOM_FLOAT_62

//line ./op-grammar.rl:36
				// TODO max size for int/float/string
				//fmt.Println("ADD ATOM", i);
				it.AddAtom(i)

//line ./op-grammar.rl:83

//line ./op-grammar.rl:102
				had_end = true
				//fmt.Println("OP END", it.state.cs, "AT", p)

//line dfa.go:4191
			}
		}

	_out:
		{
		}
	}

//line dfa.rl:50

	//fmt.Println("DONE", it.state.cs, "at", p);

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

//line dfa.rl:69
//line dfa.go:4220
	const UUID_start int = 1
	const UUID_first_final int = 2
	const UUID_error int = 0

	const UUID_en_main int = 1

//line dfa.rl:70
	var i uint128 = ctx_uuid.uint128
	digit := uint(0)
	half := 0

	cs, p, pe, eof := 0, 0, len(data), len(data)
	//var ts, te, act int
	_ = eof
	//_,_,_ = ts,te,act

//line dfa.go:4241
	{
		cs = UUID_start
	}

//line dfa.go:4246
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
//line dfa.go:4341
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
//line dfa.go:4435
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
//line dfa.go:4485
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
//line dfa.go:4564
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

//line dfa.go:4631
			}
		}

	_out:
		{
		}
	}

//line dfa.rl:88

	if cs < UUID_first_final || digit > 10 {
		return ERROR_UUID, errors.New(fmt.Sprintf("parse error at pos %d", p))
	} else {
		return UUID{uint128: i}, nil
	}

}
