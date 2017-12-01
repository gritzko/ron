package rdt

import (
	"github.com/gritzko/ron"
	"testing"
	"fmt"
)


// 3-part tables: first all inserts, then all deletes
var rga_3_tests = [][3]string{

	{ // 0+o
		"*rga#textB!",
		"*rga#textB@time'A'",
		"*rga#textB@time!@'A'",
	},
	{ // s+o
		"*rga#test@1!@'A'",
		"*rga#test@2:1'B'",
		"*rga#test@2!@1'A'@2'B'",
	},
	{ // o+o
		"*rga#test@2:1'B'",
		"*rga#test@3:2'C'",
		"*rga#test@3:1!@2:0'B'@3'C'",
	},
	{ // s+p
		"*rga#test@1!@'A'",
		"*rga#test@2:1!:0'B'",
		"*rga#test@2!@1'A'@2'B'",
	},
	{ // 4) p+p
		"*rga#test@2:1!:0'B'",
		"*rga#test@3:2!:0'C'",
		"*rga#test@3:1!@2:0'B'@3'C'",
	},

	{ // s+s
		"*rga#test@1!@'A'",
		"*rga#test@2!@1'A'@2'B'",
		"*rga#test@2!@1'A'@2'B'",
	},
	{ // 6) s1+s2 merge
		"*rga#test@2!@1'A'@2'B'",
		"*rga#test@3!@1'A'@3'C'",
		"*rga#test@3!@1'A'@3'C'@2'B'",
	},
	{ // s1+s(rm) merge
		"*rga#test@2!@1'A'@2'B'",
		"*rga#test@4!@1:4'A'@3:0'C'",
		"*rga#test@4!@1:4'A'@3:0'C'@2'B'",
	},

	{ // 8) o+rm
		"*rga#test@2:1'B'",
		"*rga#test@3:2;",
		"*rga#test@3:1!@2:3'B'",
	},

	{ // p+rm
		"*rga#test@3:1!@2:0'B'@3'C'",
		"*rga#test@4:2;",
		"*rga#test@4:1!@2:4'B'@3:0'C'",
	},
	{ // 10 s+rms
		"*rga#test@2!@1'A'@2'B'",
		"*rga#test@4:rm!@3:1,@4:2,",
		"*rga#test@4!@1:3'A'@2:4'B'",
	},
	{ // s(rm)+s(rm) merge
		"*rga#test@5!@1:4a'A'@2:5'B'",
		"*rga#test@4!@1:4'A'@3:0'C'",
		"*rga#test@4!@1:4a'A'@3:0'C'@2:5'B'",
	},
	{ // 12 s(rm)+s(rm) merge
		"*rga#test@3!@1:4a'A'@3:0'C'@2:5'B'",
		"*rga#test@4!@1:4a'A'@3:0'C'@4:0'D'@2:5'B'",
		"*rga#test@4!@1:4a'A'@3:0'C'@4'D'@2:5'B'",
	},
	{ // s(rm)+s(rm) merge
		"*rga#test@5!@1:4a'A'@5:0'E'@3:0'C'@2:5'B'",
		"*rga#test@7!@1:4a'A'@6:0'F'@3:7'C'@4:0'D'@2:5'B'",
		"*rga#test@7!@1:4a'A'@6:0'F'@5'E'@3:7'C'@4:0'D'@2:5'B'",
	},
	{ // 14 s+ins
		"*rga#test@2!@1'A'@2'B'",
		"*rga#test@3:1'-';",
		"*rga#test@3!@1'A'@3'-'@2'B'",
	},
	{ // 15 eclipsed rm
		"*rga#test@4!@1'A'@2:4'B'",
		"*rga#test@3:2;",
		"*rga#test@3!@1'A'@2:4'B'",
	},
	// CUT BRANCHES
	{ // 16 unapplied remove
		"*rga#test@2!@1'A'@2'B'",
		"*rga#test@4:3;",
		"*rga#test@4!@1'A'@2'B'@4:rm!:3,",
	},
	{ // 17 unapplied remove applied
		"*rga#test@4!@1'A'@2'B'@4:rm!:3,",
		"*rga#test@3:2'C'",
		"*rga#test@3!@1'A'@2'B'@3:4'C'",
	},
	{ // 18 unapplied patch
		"*rga#test@2!@1'A'@2'B'",
		"*rga#test@5:3!@4:0'D'@5'E'",
		"*rga#test@5!@1'A'@2'B'@5:3!@4:0'D'@5'E'",
	},
	{ // 19 unapplied patch applied
		"*rga#test@2!@1'A'@2'B'@5:3!@4:0'D'@5'E'",
		"*rga#test@3:2'C';",
		"*rga#test@3!@1'A'@2'B'@3'C'@4'D'@5'E'",
	},
	{ // 20 unapplied patch - ins+rm
		"*rga#test@2!@1'A'@2'B'",
		"*rga#test@6:3!@4:0'D'@5'E'@6:rm!:3,",
		"*rga#test@6!@1'A'@2'B'@6:3!@4:0'D'@5'E'@6:rm!:3,",
	},
	{ // 21 unapplied ins+rm patch applied
		"*rga#test@6!@1'A'@2'B'@6:3!@4:0'D'@5'E'@6:rm!:3,",
		"*rga#test@3:2!@'C'",
		"*rga#test@3!@1'A'@2'B'@3:6'C'@4:0'D'@5'E'",
	},
}

func TestRGA_Reduce(t *testing.T) {
	for i := 0; i < len(rga_3_tests); i++ {
		test := rga_3_tests[i]
		C := test[2]
		frames := [2]ron.Frame{
			ron.ParseFrameString(test[0]),
			ron.ParseFrameString(test[1]),
		}
		rga := MakeRGAReducer()
		//fmt.Printf("%d\n%s\n%s\n",i, test[0], test[1])
		frameC := rga.Reduce(frames[0:2])
		if frameC.String() != C {
			t.Fail()
			fmt.Printf("\n-------------------------\nwrong result at %d: \nhave [ %s ]\nneed [ %s ]\n\n", i, frameC.String(), C)
		}

	}
}


func TestRGA_Mapper(t *testing.T) {
	frame := "*rga#1UQ8p+bart@1UQ8yk+lisa!"+
		"@(s+bart'H'@[r'e'@(t'l'@[T'l'@[i'o'"+
		"@(w+lisa' '@(x'w'@(y'o'@[1'r'@{a'l'@[2'd'@[k'!'"
	right := "Hello world!"
	var txt TxtMapper
	hello := txt.Map(ron.Batch{ron.ParseFrameString(frame)})
	if hello != right {
		t.Fail()
		t.Logf("'%s' != '%s'", hello, right)
	}
}

// reduceAll: 4-line tables (state, ch1, ch2, result)
