
//line parser.rl:1
package RON


type parser struct {
    data []byte
    p, pe, cs int
    ts, te, act int
}

//func ParseOp(data []byte, op *Op) int {
//    return -1
//}

func ParseOp(data []byte, op *Op) int {

    
//line parser.rl:17
    
//line parser.go:22
const RON_start int = 1
const RON_first_final int = 32
const RON_error int = 0

const RON_en_main int = 1


//line parser.rl:18

    var uuid UUID
    var _, int60_off int
    var ret int

	cs, p, pe, eof := 0, 0, len(data), len(data)
	var ts, te, act int
    _ = eof
    _,_,_ = ts,te,act

	
//line parser.go:42
	{
	cs = RON_start
	}

//line parser.go:47
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
	case 32:
		goto st_case_32
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
	case 8:
		goto st_case_8
	case 9:
		goto st_case_9
	case 10:
		goto st_case_10
	case 11:
		goto st_case_11
	case 12:
		goto st_case_12
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
	case 18:
		goto st_case_18
	case 19:
		goto st_case_19
	case 20:
		goto st_case_20
	case 21:
		goto st_case_21
	case 22:
		goto st_case_22
	case 23:
		goto st_case_23
	case 24:
		goto st_case_24
	case 25:
		goto st_case_25
	case 26:
		goto st_case_26
	case 27:
		goto st_case_27
	case 28:
		goto st_case_28
	case 29:
		goto st_case_29
	case 30:
		goto st_case_30
	case 31:
		goto st_case_31
	}
	goto st_out
	st_case_1:
		switch data[p] {
		case 33:
			goto tr0
		case 35:
			goto tr2
		case 39:
			goto tr3
		case 46:
			goto tr4
		case 58:
			goto tr5
		case 61:
			goto tr6
		case 62:
			goto tr7
		case 63:
			goto tr0
		case 64:
			goto tr8
		}
		goto st0
st_case_0:
	st0:
		cs = 0
		goto _out
tr0:
//line parser.rl:75

            //uuid = zero
        
//line parser.rl:87

        
//line parser.rl:78

            //uuid = zero
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st2
tr10:
//line parser.rl:99

            ////fmt.Printf("ATOM\n");
        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st2
tr19:
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:99

            ////fmt.Printf("ATOM\n");
        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st2
tr27:
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:99

            ////fmt.Printf("ATOM\n");
        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st2
tr34:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:99

            ////fmt.Printf("ATOM\n");
        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st2
tr42:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:99

            ////fmt.Printf("ATOM\n");
        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st2
tr46:
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st2
tr55:
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st2
tr63:
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st2
tr70:
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st2
tr76:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st2
tr83:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st2
tr87:
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st2
tr95:
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st2
tr102:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st2
tr110:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st2
tr115:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st2
tr124:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st2
tr130:
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:87

        
//line parser.rl:78

            //uuid = zero
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st2
tr140:
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:87

        
//line parser.rl:78

            //uuid = zero
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st2
tr149:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:87

        
//line parser.rl:78

            //uuid = zero
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st2
tr159:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:87

        
//line parser.rl:78

            //uuid = zero
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
//line parser.go:738
		switch data[p] {
		case 10:
			goto tr9
		case 33:
			goto tr10
		case 39:
			goto tr11
		case 61:
			goto tr12
		case 62:
			goto tr13
		case 63:
			goto tr10
		}
		goto st0
tr9:
//line parser.rl:99

            ////fmt.Printf("ATOM\n");
        
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
	goto st32
tr18:
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:99

            ////fmt.Printf("ATOM\n");
        
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
	goto st32
tr26:
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:99

            ////fmt.Printf("ATOM\n");
        
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
	goto st32
tr33:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:99

            ////fmt.Printf("ATOM\n");
        
	goto st32
tr41:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:99

            ////fmt.Printf("ATOM\n");
        
	goto st32
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
//line parser.go:842
		goto st0
tr3:
//line parser.rl:75

            //uuid = zero
        
//line parser.rl:87

        
//line parser.rl:78

            //uuid = zero
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st3
tr11:
//line parser.rl:99

            ////fmt.Printf("ATOM\n");
        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st3
tr21:
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:99

            ////fmt.Printf("ATOM\n");
        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st3
tr28:
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:99

            ////fmt.Printf("ATOM\n");
        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st3
tr36:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:99

            ////fmt.Printf("ATOM\n");
        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st3
tr43:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:99

            ////fmt.Printf("ATOM\n");
        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st3
tr48:
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st3
tr56:
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st3
tr65:
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st3
tr71:
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st3
tr78:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st3
tr84:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st3
tr89:
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st3
tr96:
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st3
tr104:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st3
tr111:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st3
tr117:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st3
tr125:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st3
tr133:
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:87

        
//line parser.rl:78

            //uuid = zero
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st3
tr142:
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:87

        
//line parser.rl:78

            //uuid = zero
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st3
tr152:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:87

        
//line parser.rl:78

            //uuid = zero
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st3
tr161:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:87

        
//line parser.rl:78

            //uuid = zero
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
//line parser.go:1435
		if data[p] == 39 {
			goto st2
		}
		goto st3
tr6:
//line parser.rl:75

            //uuid = zero
        
//line parser.rl:87

        
//line parser.rl:78

            //uuid = zero
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st4
tr12:
//line parser.rl:99

            ////fmt.Printf("ATOM\n");
        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st4
tr24:
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:99

            ////fmt.Printf("ATOM\n");
        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st4
tr31:
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:99

            ////fmt.Printf("ATOM\n");
        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st4
tr39:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:99

            ////fmt.Printf("ATOM\n");
        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st4
tr44:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:99

            ////fmt.Printf("ATOM\n");
        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st4
tr52:
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st4
tr60:
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st4
tr68:
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st4
tr74:
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st4
tr81:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st4
tr85:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st4
tr93:
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st4
tr100:
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st4
tr108:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st4
tr113:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st4
tr121:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st4
tr127:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st4
tr137:
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:87

        
//line parser.rl:78

            //uuid = zero
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st4
tr146:
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:87

        
//line parser.rl:78

            //uuid = zero
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st4
tr156:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:87

        
//line parser.rl:78

            //uuid = zero
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st4
tr163:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:87

        
//line parser.rl:78

            //uuid = zero
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
//line parser.go:2031
		switch data[p] {
		case 43:
			goto st5
		case 45:
			goto st5
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr17
		}
		goto st0
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		if 48 <= data[p] && data[p] <= 57 {
			goto tr17
		}
		goto st0
tr17:
//line parser.rl:108

            ////fmt.Printf("DIGIT %c\n", fc);
        
	goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
//line parser.go:2062
		switch data[p] {
		case 10:
			goto tr9
		case 33:
			goto tr10
		case 39:
			goto tr11
		case 61:
			goto tr12
		case 62:
			goto tr13
		case 63:
			goto tr10
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr17
		}
		goto st0
tr7:
//line parser.rl:75

            //uuid = zero
        
//line parser.rl:87

        
//line parser.rl:78

            //uuid = zero
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st7
tr13:
//line parser.rl:99

            ////fmt.Printf("ATOM\n");
        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st7
tr25:
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:99

            ////fmt.Printf("ATOM\n");
        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st7
tr32:
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:99

            ////fmt.Printf("ATOM\n");
        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st7
tr40:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:99

            ////fmt.Printf("ATOM\n");
        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st7
tr45:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:99

            ////fmt.Printf("ATOM\n");
        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st7
tr53:
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st7
tr61:
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st7
tr69:
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st7
tr75:
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st7
tr82:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st7
tr86:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st7
tr94:
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st7
tr101:
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st7
tr109:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st7
tr114:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st7
tr122:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st7
tr128:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st7
tr138:
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:87

        
//line parser.rl:78

            //uuid = zero
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st7
tr147:
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:87

        
//line parser.rl:78

            //uuid = zero
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st7
tr157:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:87

        
//line parser.rl:78

            //uuid = zero
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st7
tr164:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:87

        
//line parser.rl:78

            //uuid = zero
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
//line parser.rl:93

        
//line parser.rl:96

            ////fmt.Printf("ATOM_START %c\n", fc);
        
	goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
//line parser.go:2672
		switch data[p] {
		case 10:
			goto tr18
		case 33:
			goto tr19
		case 39:
			goto tr21
		case 43:
			goto tr20
		case 45:
			goto tr20
		case 61:
			goto tr24
		case 62:
			goto tr25
		case 63:
			goto tr19
		case 91:
			goto tr22
		case 93:
			goto tr22
		case 95:
			goto tr23
		case 123:
			goto tr22
		case 125:
			goto tr22
		case 126:
			goto tr23
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 37:
				if 40 <= data[p] && data[p] <= 41 {
					goto tr22
				}
			case data[p] >= 36:
				goto tr20
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr23
				}
			case data[p] >= 65:
				goto tr23
			}
		default:
			goto tr23
		}
		goto st0
tr20:
//line parser.rl:60

            uuid.Sign = data[p];
            _ = 11;
        
	goto st8
tr35:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:60

            uuid.Sign = data[p];
            _ = 11;
        
	goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
//line parser.go:2750
		switch data[p] {
		case 10:
			goto tr26
		case 33:
			goto tr27
		case 39:
			goto tr28
		case 61:
			goto tr31
		case 62:
			goto tr32
		case 63:
			goto tr27
		case 91:
			goto tr29
		case 93:
			goto tr29
		case 95:
			goto tr30
		case 123:
			goto tr29
		case 125:
			goto tr29
		case 126:
			goto tr30
		}
		switch {
		case data[p] < 48:
			if 40 <= data[p] && data[p] <= 41 {
				goto tr29
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr30
				}
			case data[p] >= 65:
				goto tr30
			}
		default:
			goto tr30
		}
		goto st0
tr29:
//line parser.rl:30
 
            ////fmt.Printf("INT60_PREFIX: %c\n", fc);
            int60_off = 4
        
	goto st9
tr30:
//line parser.rl:34

            ////fmt.Printf("INT60_DIGIT: %c\n", fc);
            //uuid[_+int60_off] = fc
            uuid.Value = (uuid.Value<<6) | uint64(ABC[data[p]])
            int60_off++
        
	goto st9
tr37:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:30
 
            ////fmt.Printf("INT60_PREFIX: %c\n", fc);
            int60_off = 4
        
	goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
//line parser.go:2828
		switch data[p] {
		case 10:
			goto tr26
		case 33:
			goto tr27
		case 39:
			goto tr28
		case 61:
			goto tr31
		case 62:
			goto tr32
		case 63:
			goto tr27
		case 95:
			goto tr30
		case 126:
			goto tr30
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr30
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr30
			}
		default:
			goto tr30
		}
		goto st0
tr22:
//line parser.rl:30
 
            ////fmt.Printf("INT60_PREFIX: %c\n", fc);
            int60_off = 4
        
	goto st10
tr38:
//line parser.rl:34

            ////fmt.Printf("INT60_DIGIT: %c\n", fc);
            //uuid[_+int60_off] = fc
            uuid.Value = (uuid.Value<<6) | uint64(ABC[data[p]])
            int60_off++
        
	goto st10
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
//line parser.go:2881
		switch data[p] {
		case 10:
			goto tr33
		case 33:
			goto tr34
		case 39:
			goto tr36
		case 43:
			goto tr35
		case 45:
			goto tr35
		case 61:
			goto tr39
		case 62:
			goto tr40
		case 63:
			goto tr34
		case 91:
			goto tr37
		case 93:
			goto tr37
		case 95:
			goto tr38
		case 123:
			goto tr37
		case 125:
			goto tr37
		case 126:
			goto tr38
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 37:
				if 40 <= data[p] && data[p] <= 41 {
					goto tr37
				}
			case data[p] >= 36:
				goto tr35
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr38
				}
			case data[p] >= 65:
				goto tr38
			}
		default:
			goto tr38
		}
		goto st0
tr23:
//line parser.rl:34

            ////fmt.Printf("INT60_DIGIT: %c\n", fc);
            //uuid[_+int60_off] = fc
            uuid.Value = (uuid.Value<<6) | uint64(ABC[data[p]])
            int60_off++
        
	goto st11
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
//line parser.go:2949
		switch data[p] {
		case 10:
			goto tr41
		case 33:
			goto tr42
		case 39:
			goto tr43
		case 43:
			goto tr35
		case 45:
			goto tr35
		case 61:
			goto tr44
		case 62:
			goto tr45
		case 63:
			goto tr42
		case 91:
			goto tr37
		case 93:
			goto tr37
		case 95:
			goto tr23
		case 123:
			goto tr37
		case 125:
			goto tr37
		case 126:
			goto tr23
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 37:
				if 40 <= data[p] && data[p] <= 41 {
					goto tr37
				}
			case data[p] >= 36:
				goto tr35
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr23
				}
			case data[p] >= 65:
				goto tr23
			}
		default:
			goto tr23
		}
		goto st0
tr2:
//line parser.rl:75

            //uuid = zero
        
//line parser.rl:87

        
//line parser.rl:78

            //uuid = zero
        
	goto st12
tr131:
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:87

        
//line parser.rl:78

            //uuid = zero
        
	goto st12
tr141:
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:87

        
//line parser.rl:78

            //uuid = zero
        
	goto st12
tr150:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:87

        
//line parser.rl:78

            //uuid = zero
        
	goto st12
tr160:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:87

        
//line parser.rl:78

            //uuid = zero
        
	goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
//line parser.go:3095
		switch data[p] {
		case 33:
			goto tr46
		case 39:
			goto tr48
		case 43:
			goto tr47
		case 45:
			goto tr47
		case 58:
			goto tr51
		case 61:
			goto tr52
		case 62:
			goto tr53
		case 63:
			goto tr46
		case 64:
			goto tr54
		case 91:
			goto tr49
		case 93:
			goto tr49
		case 95:
			goto tr50
		case 123:
			goto tr49
		case 125:
			goto tr49
		case 126:
			goto tr50
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 37:
				if 40 <= data[p] && data[p] <= 41 {
					goto tr49
				}
			case data[p] >= 36:
				goto tr47
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr50
				}
			case data[p] >= 65:
				goto tr50
			}
		default:
			goto tr50
		}
		goto st0
tr47:
//line parser.rl:60

            uuid.Sign = data[p];
            _ = 11;
        
	goto st13
tr116:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:60

            uuid.Sign = data[p];
            _ = 11;
        
	goto st13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
//line parser.go:3175
		switch data[p] {
		case 33:
			goto tr55
		case 39:
			goto tr56
		case 58:
			goto tr59
		case 61:
			goto tr60
		case 62:
			goto tr61
		case 63:
			goto tr55
		case 64:
			goto tr62
		case 91:
			goto tr57
		case 93:
			goto tr57
		case 95:
			goto tr58
		case 123:
			goto tr57
		case 125:
			goto tr57
		case 126:
			goto tr58
		}
		switch {
		case data[p] < 48:
			if 40 <= data[p] && data[p] <= 41 {
				goto tr57
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr58
				}
			case data[p] >= 65:
				goto tr58
			}
		default:
			goto tr58
		}
		goto st0
tr57:
//line parser.rl:30
 
            ////fmt.Printf("INT60_PREFIX: %c\n", fc);
            int60_off = 4
        
	goto st14
tr58:
//line parser.rl:34

            ////fmt.Printf("INT60_DIGIT: %c\n", fc);
            //uuid[_+int60_off] = fc
            uuid.Value = (uuid.Value<<6) | uint64(ABC[data[p]])
            int60_off++
        
	goto st14
tr118:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:30
 
            ////fmt.Printf("INT60_PREFIX: %c\n", fc);
            int60_off = 4
        
	goto st14
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
//line parser.go:3255
		switch data[p] {
		case 33:
			goto tr55
		case 39:
			goto tr56
		case 58:
			goto tr59
		case 61:
			goto tr60
		case 62:
			goto tr61
		case 63:
			goto tr55
		case 64:
			goto tr62
		case 95:
			goto tr58
		case 126:
			goto tr58
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr58
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr58
			}
		default:
			goto tr58
		}
		goto st0
tr5:
//line parser.rl:75

            //uuid = zero
        
//line parser.rl:87

        
//line parser.rl:78

            //uuid = zero
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
	goto st15
tr51:
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
	goto st15
tr59:
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
	goto st15
tr92:
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
	goto st15
tr99:
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
	goto st15
tr107:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
	goto st15
tr112:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
	goto st15
tr120:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
	goto st15
tr126:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
	goto st15
tr136:
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:87

        
//line parser.rl:78

            //uuid = zero
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
	goto st15
tr145:
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:87

        
//line parser.rl:78

            //uuid = zero
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
	goto st15
tr155:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:87

        
//line parser.rl:78

            //uuid = zero
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
	goto st15
tr162:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:87

        
//line parser.rl:78

            //uuid = zero
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
//line parser.rl:91

        
//line parser.rl:84

            //uuid = zero
        
	goto st15
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
//line parser.go:3627
		switch data[p] {
		case 33:
			goto tr63
		case 39:
			goto tr65
		case 43:
			goto tr64
		case 45:
			goto tr64
		case 61:
			goto tr68
		case 62:
			goto tr69
		case 63:
			goto tr63
		case 91:
			goto tr66
		case 93:
			goto tr66
		case 95:
			goto tr67
		case 123:
			goto tr66
		case 125:
			goto tr66
		case 126:
			goto tr67
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 37:
				if 40 <= data[p] && data[p] <= 41 {
					goto tr66
				}
			case data[p] >= 36:
				goto tr64
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr67
				}
			case data[p] >= 65:
				goto tr67
			}
		default:
			goto tr67
		}
		goto st0
tr64:
//line parser.rl:60

            uuid.Sign = data[p];
            _ = 11;
        
	goto st16
tr77:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:60

            uuid.Sign = data[p];
            _ = 11;
        
	goto st16
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
//line parser.go:3703
		switch data[p] {
		case 33:
			goto tr70
		case 39:
			goto tr71
		case 61:
			goto tr74
		case 62:
			goto tr75
		case 63:
			goto tr70
		case 91:
			goto tr72
		case 93:
			goto tr72
		case 95:
			goto tr73
		case 123:
			goto tr72
		case 125:
			goto tr72
		case 126:
			goto tr73
		}
		switch {
		case data[p] < 48:
			if 40 <= data[p] && data[p] <= 41 {
				goto tr72
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr73
				}
			case data[p] >= 65:
				goto tr73
			}
		default:
			goto tr73
		}
		goto st0
tr72:
//line parser.rl:30
 
            ////fmt.Printf("INT60_PREFIX: %c\n", fc);
            int60_off = 4
        
	goto st17
tr73:
//line parser.rl:34

            ////fmt.Printf("INT60_DIGIT: %c\n", fc);
            //uuid[_+int60_off] = fc
            uuid.Value = (uuid.Value<<6) | uint64(ABC[data[p]])
            int60_off++
        
	goto st17
tr79:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:30
 
            ////fmt.Printf("INT60_PREFIX: %c\n", fc);
            int60_off = 4
        
	goto st17
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
//line parser.go:3779
		switch data[p] {
		case 33:
			goto tr70
		case 39:
			goto tr71
		case 61:
			goto tr74
		case 62:
			goto tr75
		case 63:
			goto tr70
		case 95:
			goto tr73
		case 126:
			goto tr73
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr73
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr73
			}
		default:
			goto tr73
		}
		goto st0
tr66:
//line parser.rl:30
 
            ////fmt.Printf("INT60_PREFIX: %c\n", fc);
            int60_off = 4
        
	goto st18
tr80:
//line parser.rl:34

            ////fmt.Printf("INT60_DIGIT: %c\n", fc);
            //uuid[_+int60_off] = fc
            uuid.Value = (uuid.Value<<6) | uint64(ABC[data[p]])
            int60_off++
        
	goto st18
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
//line parser.go:3830
		switch data[p] {
		case 33:
			goto tr76
		case 39:
			goto tr78
		case 43:
			goto tr77
		case 45:
			goto tr77
		case 61:
			goto tr81
		case 62:
			goto tr82
		case 63:
			goto tr76
		case 91:
			goto tr79
		case 93:
			goto tr79
		case 95:
			goto tr80
		case 123:
			goto tr79
		case 125:
			goto tr79
		case 126:
			goto tr80
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 37:
				if 40 <= data[p] && data[p] <= 41 {
					goto tr79
				}
			case data[p] >= 36:
				goto tr77
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr80
				}
			case data[p] >= 65:
				goto tr80
			}
		default:
			goto tr80
		}
		goto st0
tr67:
//line parser.rl:34

            ////fmt.Printf("INT60_DIGIT: %c\n", fc);
            //uuid[_+int60_off] = fc
            uuid.Value = (uuid.Value<<6) | uint64(ABC[data[p]])
            int60_off++
        
	goto st19
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
//line parser.go:3896
		switch data[p] {
		case 33:
			goto tr83
		case 39:
			goto tr84
		case 43:
			goto tr77
		case 45:
			goto tr77
		case 61:
			goto tr85
		case 62:
			goto tr86
		case 63:
			goto tr83
		case 91:
			goto tr79
		case 93:
			goto tr79
		case 95:
			goto tr67
		case 123:
			goto tr79
		case 125:
			goto tr79
		case 126:
			goto tr67
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 37:
				if 40 <= data[p] && data[p] <= 41 {
					goto tr79
				}
			case data[p] >= 36:
				goto tr77
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr67
				}
			case data[p] >= 65:
				goto tr67
			}
		default:
			goto tr67
		}
		goto st0
tr8:
//line parser.rl:75

            //uuid = zero
        
//line parser.rl:87

        
//line parser.rl:78

            //uuid = zero
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
	goto st20
tr54:
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
	goto st20
tr62:
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
	goto st20
tr123:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
	goto st20
tr129:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
	goto st20
tr139:
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:87

        
//line parser.rl:78

            //uuid = zero
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
	goto st20
tr148:
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:87

        
//line parser.rl:78

            //uuid = zero
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
	goto st20
tr158:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:52


        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:87

        
//line parser.rl:78

            //uuid = zero
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
	goto st20
tr165:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:66

            ////fmt.Printf("uuid: %s\n", string(uuid[:]))
            _ = 0
        
//line parser.rl:87

        
//line parser.rl:78

            //uuid = zero
        
//line parser.rl:89

        
//line parser.rl:81

            //uuid = zero
        
	goto st20
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
//line parser.go:4149
		switch data[p] {
		case 33:
			goto tr87
		case 39:
			goto tr89
		case 43:
			goto tr88
		case 45:
			goto tr88
		case 58:
			goto tr92
		case 61:
			goto tr93
		case 62:
			goto tr94
		case 63:
			goto tr87
		case 91:
			goto tr90
		case 93:
			goto tr90
		case 95:
			goto tr91
		case 123:
			goto tr90
		case 125:
			goto tr90
		case 126:
			goto tr91
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 37:
				if 40 <= data[p] && data[p] <= 41 {
					goto tr90
				}
			case data[p] >= 36:
				goto tr88
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr91
				}
			case data[p] >= 65:
				goto tr91
			}
		default:
			goto tr91
		}
		goto st0
tr88:
//line parser.rl:60

            uuid.Sign = data[p];
            _ = 11;
        
	goto st21
tr103:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:60

            uuid.Sign = data[p];
            _ = 11;
        
	goto st21
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
//line parser.go:4227
		switch data[p] {
		case 33:
			goto tr95
		case 39:
			goto tr96
		case 58:
			goto tr99
		case 61:
			goto tr100
		case 62:
			goto tr101
		case 63:
			goto tr95
		case 91:
			goto tr97
		case 93:
			goto tr97
		case 95:
			goto tr98
		case 123:
			goto tr97
		case 125:
			goto tr97
		case 126:
			goto tr98
		}
		switch {
		case data[p] < 48:
			if 40 <= data[p] && data[p] <= 41 {
				goto tr97
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr98
				}
			case data[p] >= 65:
				goto tr98
			}
		default:
			goto tr98
		}
		goto st0
tr97:
//line parser.rl:30
 
            ////fmt.Printf("INT60_PREFIX: %c\n", fc);
            int60_off = 4
        
	goto st22
tr98:
//line parser.rl:34

            ////fmt.Printf("INT60_DIGIT: %c\n", fc);
            //uuid[_+int60_off] = fc
            uuid.Value = (uuid.Value<<6) | uint64(ABC[data[p]])
            int60_off++
        
	goto st22
tr105:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:30
 
            ////fmt.Printf("INT60_PREFIX: %c\n", fc);
            int60_off = 4
        
	goto st22
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
//line parser.go:4305
		switch data[p] {
		case 33:
			goto tr95
		case 39:
			goto tr96
		case 58:
			goto tr99
		case 61:
			goto tr100
		case 62:
			goto tr101
		case 63:
			goto tr95
		case 95:
			goto tr98
		case 126:
			goto tr98
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr98
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr98
			}
		default:
			goto tr98
		}
		goto st0
tr90:
//line parser.rl:30
 
            ////fmt.Printf("INT60_PREFIX: %c\n", fc);
            int60_off = 4
        
	goto st23
tr106:
//line parser.rl:34

            ////fmt.Printf("INT60_DIGIT: %c\n", fc);
            //uuid[_+int60_off] = fc
            uuid.Value = (uuid.Value<<6) | uint64(ABC[data[p]])
            int60_off++
        
	goto st23
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
//line parser.go:4358
		switch data[p] {
		case 33:
			goto tr102
		case 39:
			goto tr104
		case 43:
			goto tr103
		case 45:
			goto tr103
		case 58:
			goto tr107
		case 61:
			goto tr108
		case 62:
			goto tr109
		case 63:
			goto tr102
		case 91:
			goto tr105
		case 93:
			goto tr105
		case 95:
			goto tr106
		case 123:
			goto tr105
		case 125:
			goto tr105
		case 126:
			goto tr106
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 37:
				if 40 <= data[p] && data[p] <= 41 {
					goto tr105
				}
			case data[p] >= 36:
				goto tr103
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr106
				}
			case data[p] >= 65:
				goto tr106
			}
		default:
			goto tr106
		}
		goto st0
tr91:
//line parser.rl:34

            ////fmt.Printf("INT60_DIGIT: %c\n", fc);
            //uuid[_+int60_off] = fc
            uuid.Value = (uuid.Value<<6) | uint64(ABC[data[p]])
            int60_off++
        
	goto st24
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
//line parser.go:4426
		switch data[p] {
		case 33:
			goto tr110
		case 39:
			goto tr111
		case 43:
			goto tr103
		case 45:
			goto tr103
		case 58:
			goto tr112
		case 61:
			goto tr113
		case 62:
			goto tr114
		case 63:
			goto tr110
		case 91:
			goto tr105
		case 93:
			goto tr105
		case 95:
			goto tr91
		case 123:
			goto tr105
		case 125:
			goto tr105
		case 126:
			goto tr91
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 37:
				if 40 <= data[p] && data[p] <= 41 {
					goto tr105
				}
			case data[p] >= 36:
				goto tr103
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr91
				}
			case data[p] >= 65:
				goto tr91
			}
		default:
			goto tr91
		}
		goto st0
tr49:
//line parser.rl:30
 
            ////fmt.Printf("INT60_PREFIX: %c\n", fc);
            int60_off = 4
        
	goto st25
tr119:
//line parser.rl:34

            ////fmt.Printf("INT60_DIGIT: %c\n", fc);
            //uuid[_+int60_off] = fc
            uuid.Value = (uuid.Value<<6) | uint64(ABC[data[p]])
            int60_off++
        
	goto st25
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
//line parser.go:4501
		switch data[p] {
		case 33:
			goto tr115
		case 39:
			goto tr117
		case 43:
			goto tr116
		case 45:
			goto tr116
		case 58:
			goto tr120
		case 61:
			goto tr121
		case 62:
			goto tr122
		case 63:
			goto tr115
		case 64:
			goto tr123
		case 91:
			goto tr118
		case 93:
			goto tr118
		case 95:
			goto tr119
		case 123:
			goto tr118
		case 125:
			goto tr118
		case 126:
			goto tr119
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 37:
				if 40 <= data[p] && data[p] <= 41 {
					goto tr118
				}
			case data[p] >= 36:
				goto tr116
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr119
				}
			case data[p] >= 65:
				goto tr119
			}
		default:
			goto tr119
		}
		goto st0
tr50:
//line parser.rl:34

            ////fmt.Printf("INT60_DIGIT: %c\n", fc);
            //uuid[_+int60_off] = fc
            uuid.Value = (uuid.Value<<6) | uint64(ABC[data[p]])
            int60_off++
        
	goto st26
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
//line parser.go:4571
		switch data[p] {
		case 33:
			goto tr124
		case 39:
			goto tr125
		case 43:
			goto tr116
		case 45:
			goto tr116
		case 58:
			goto tr126
		case 61:
			goto tr127
		case 62:
			goto tr128
		case 63:
			goto tr124
		case 64:
			goto tr129
		case 91:
			goto tr118
		case 93:
			goto tr118
		case 95:
			goto tr50
		case 123:
			goto tr118
		case 125:
			goto tr118
		case 126:
			goto tr50
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 37:
				if 40 <= data[p] && data[p] <= 41 {
					goto tr118
				}
			case data[p] >= 36:
				goto tr116
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr50
				}
			case data[p] >= 65:
				goto tr50
			}
		default:
			goto tr50
		}
		goto st0
tr4:
//line parser.rl:75

            //uuid = zero
        
	goto st27
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
//line parser.go:4638
		switch data[p] {
		case 33:
			goto tr130
		case 35:
			goto tr131
		case 39:
			goto tr133
		case 43:
			goto tr132
		case 45:
			goto tr132
		case 58:
			goto tr136
		case 61:
			goto tr137
		case 62:
			goto tr138
		case 63:
			goto tr130
		case 64:
			goto tr139
		case 91:
			goto tr134
		case 93:
			goto tr134
		case 95:
			goto tr135
		case 123:
			goto tr134
		case 125:
			goto tr134
		case 126:
			goto tr135
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 37:
				if 40 <= data[p] && data[p] <= 41 {
					goto tr134
				}
			case data[p] >= 36:
				goto tr132
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr135
				}
			case data[p] >= 65:
				goto tr135
			}
		default:
			goto tr135
		}
		goto st0
tr132:
//line parser.rl:60

            uuid.Sign = data[p];
            _ = 11;
        
	goto st28
tr151:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:60

            uuid.Sign = data[p];
            _ = 11;
        
	goto st28
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
//line parser.go:4720
		switch data[p] {
		case 33:
			goto tr140
		case 35:
			goto tr141
		case 39:
			goto tr142
		case 58:
			goto tr145
		case 61:
			goto tr146
		case 62:
			goto tr147
		case 63:
			goto tr140
		case 64:
			goto tr148
		case 91:
			goto tr143
		case 93:
			goto tr143
		case 95:
			goto tr144
		case 123:
			goto tr143
		case 125:
			goto tr143
		case 126:
			goto tr144
		}
		switch {
		case data[p] < 48:
			if 40 <= data[p] && data[p] <= 41 {
				goto tr143
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr144
				}
			case data[p] >= 65:
				goto tr144
			}
		default:
			goto tr144
		}
		goto st0
tr143:
//line parser.rl:30
 
            ////fmt.Printf("INT60_PREFIX: %c\n", fc);
            int60_off = 4
        
	goto st29
tr144:
//line parser.rl:34

            ////fmt.Printf("INT60_DIGIT: %c\n", fc);
            //uuid[_+int60_off] = fc
            uuid.Value = (uuid.Value<<6) | uint64(ABC[data[p]])
            int60_off++
        
	goto st29
tr153:
//line parser.rl:103

            ////fmt.Printf("VALUE\n");
            ret = p
        
//line parser.rl:30
 
            ////fmt.Printf("INT60_PREFIX: %c\n", fc);
            int60_off = 4
        
	goto st29
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
//line parser.go:4802
		switch data[p] {
		case 33:
			goto tr140
		case 35:
			goto tr141
		case 39:
			goto tr142
		case 58:
			goto tr145
		case 61:
			goto tr146
		case 62:
			goto tr147
		case 63:
			goto tr140
		case 64:
			goto tr148
		case 95:
			goto tr144
		case 126:
			goto tr144
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr144
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr144
			}
		default:
			goto tr144
		}
		goto st0
tr134:
//line parser.rl:30
 
            ////fmt.Printf("INT60_PREFIX: %c\n", fc);
            int60_off = 4
        
	goto st30
tr154:
//line parser.rl:34

            ////fmt.Printf("INT60_DIGIT: %c\n", fc);
            //uuid[_+int60_off] = fc
            uuid.Value = (uuid.Value<<6) | uint64(ABC[data[p]])
            int60_off++
        
	goto st30
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
//line parser.go:4859
		switch data[p] {
		case 33:
			goto tr149
		case 35:
			goto tr150
		case 39:
			goto tr152
		case 43:
			goto tr151
		case 45:
			goto tr151
		case 58:
			goto tr155
		case 61:
			goto tr156
		case 62:
			goto tr157
		case 63:
			goto tr149
		case 64:
			goto tr158
		case 91:
			goto tr153
		case 93:
			goto tr153
		case 95:
			goto tr154
		case 123:
			goto tr153
		case 125:
			goto tr153
		case 126:
			goto tr154
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 37:
				if 40 <= data[p] && data[p] <= 41 {
					goto tr153
				}
			case data[p] >= 36:
				goto tr151
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr154
				}
			case data[p] >= 65:
				goto tr154
			}
		default:
			goto tr154
		}
		goto st0
tr135:
//line parser.rl:34

            ////fmt.Printf("INT60_DIGIT: %c\n", fc);
            //uuid[_+int60_off] = fc
            uuid.Value = (uuid.Value<<6) | uint64(ABC[data[p]])
            int60_off++
        
	goto st31
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
//line parser.go:4931
		switch data[p] {
		case 33:
			goto tr159
		case 35:
			goto tr160
		case 39:
			goto tr161
		case 43:
			goto tr151
		case 45:
			goto tr151
		case 58:
			goto tr162
		case 61:
			goto tr163
		case 62:
			goto tr164
		case 63:
			goto tr159
		case 64:
			goto tr165
		case 91:
			goto tr153
		case 93:
			goto tr153
		case 95:
			goto tr135
		case 123:
			goto tr153
		case 125:
			goto tr153
		case 126:
			goto tr135
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 37:
				if 40 <= data[p] && data[p] <= 41 {
					goto tr153
				}
			case data[p] >= 36:
				goto tr151
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr135
				}
			case data[p] >= 65:
				goto tr135
			}
		default:
			goto tr135
		}
		goto st0
	st_out:
	_test_eof2: cs = 2; goto _test_eof
	_test_eof32: cs = 32; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
	_test_eof16: cs = 16; goto _test_eof
	_test_eof17: cs = 17; goto _test_eof
	_test_eof18: cs = 18; goto _test_eof
	_test_eof19: cs = 19; goto _test_eof
	_test_eof20: cs = 20; goto _test_eof
	_test_eof21: cs = 21; goto _test_eof
	_test_eof22: cs = 22; goto _test_eof
	_test_eof23: cs = 23; goto _test_eof
	_test_eof24: cs = 24; goto _test_eof
	_test_eof25: cs = 25; goto _test_eof
	_test_eof26: cs = 26; goto _test_eof
	_test_eof27: cs = 27; goto _test_eof
	_test_eof28: cs = 28; goto _test_eof
	_test_eof29: cs = 29; goto _test_eof
	_test_eof30: cs = 30; goto _test_eof
	_test_eof31: cs = 31; goto _test_eof

	_test_eof: {}
	_out: {}
	}

//line parser.rl:118


        return ret
}



// BIG FIXME  ERROR HANDLING, TESTS
func ParseUUID(data []byte, context UUID) (uuid UUID, length int) {

    
//line parser.rl:129
    
//line parser.go:5040
const UUID_start int = 1
const UUID_first_final int = 1
const UUID_error int = 0

const UUID_en_main int = 1


//line parser.rl:130

    uuid = context
    var i uint64 = context.Value
    var digits uint
    length = -1

	cs, p, pe, eof := 0, 0, len(data), len(data)
	var ts, te, act int
    _ = eof
    _,_,_ = ts,te,act


	
//line parser.go:5062
	{
	cs = UUID_start
	}

//line parser.go:5067
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
//line parser.rl:172

            uuid.Sign = data[p]
            i = context.Origin
        
	goto st2
tr6:
//line parser.rl:158

            if digits>0 {
                uuid.Value = i << (60-digits)
                digits = 0
            }
            i = context.Origin
        
//line parser.rl:172

            uuid.Sign = data[p]
            i = context.Origin
        
	goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
//line parser.go:5160
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
//line parser.rl:144

            digits = uint(-ABC[data[p]]-10+4) * 6
            i >>= (60-digits)  // FIXME
        
	goto st3
tr5:
//line parser.rl:149

            if digits==0 {
                i = 0
            }
            digits+=6
            i <<= 6
            i |= uint64(ABC[data[p]])
        
	goto st3
tr7:
//line parser.rl:158

            if digits>0 {
                uuid.Value = i << (60-digits)
                digits = 0
            }
            i = context.Origin
        
//line parser.rl:144

            digits = uint(-ABC[data[p]]-10+4) * 6
            i >>= (60-digits)  // FIXME
        
	goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
//line parser.go:5231
		switch data[p] {
		case 95:
			goto tr5
		case 126:
			goto tr5
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr5
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr5
			}
		default:
			goto tr5
		}
		goto st0
tr2:
//line parser.rl:144

            digits = uint(-ABC[data[p]]-10+4) * 6
            i >>= (60-digits)  // FIXME
        
	goto st4
tr8:
//line parser.rl:149

            if digits==0 {
                i = 0
            }
            digits+=6
            i <<= 6
            i |= uint64(ABC[data[p]])
        
	goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
//line parser.go:5274
		switch data[p] {
		case 43:
			goto tr6
		case 45:
			goto tr6
		case 91:
			goto tr7
		case 93:
			goto tr7
		case 95:
			goto tr8
		case 123:
			goto tr7
		case 125:
			goto tr7
		case 126:
			goto tr8
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 37:
				if 40 <= data[p] && data[p] <= 41 {
					goto tr7
				}
			case data[p] >= 36:
				goto tr6
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr8
				}
			case data[p] >= 65:
				goto tr8
			}
		default:
			goto tr8
		}
		goto st0
tr3:
//line parser.rl:149

            if digits==0 {
                i = 0
            }
            digits+=6
            i <<= 6
            i |= uint64(ABC[data[p]])
        
	goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
//line parser.go:5332
		switch data[p] {
		case 43:
			goto tr6
		case 45:
			goto tr6
		case 91:
			goto tr7
		case 93:
			goto tr7
		case 95:
			goto tr3
		case 123:
			goto tr7
		case 125:
			goto tr7
		case 126:
			goto tr3
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 37:
				if 40 <= data[p] && data[p] <= 41 {
					goto tr7
				}
			case data[p] >= 36:
				goto tr6
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
	st_out:
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 1:
//line parser.rl:177

            length = pe
        
		case 5:
//line parser.rl:158

            if digits>0 {
                uuid.Value = i << (60-digits)
                digits = 0
            }
            i = context.Origin
        
//line parser.rl:177

            length = pe
        
		case 2, 3:
//line parser.rl:166

            if digits>0 {
                uuid.Origin = i << (60-digits)
            }
        
//line parser.rl:177

            length = pe
        
		case 4:
//line parser.rl:158

            if digits>0 {
                uuid.Value = i << (60-digits)
                digits = 0
            }
            i = context.Origin
        
//line parser.rl:166

            if digits>0 {
                uuid.Origin = i << (60-digits)
            }
        
//line parser.rl:177

            length = pe
        
//line parser.go:5431
		}
	}

	_out: {}
	}

//line parser.rl:186


    // FIXME checkk all input is parsed

    return
}

