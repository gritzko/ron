package rdt

import (
	"github.com/gritzko/ron"
	"testing"
	"fmt"
)

func TestIMultiMap_Take(t *testing.T) {
	frame := ron.ParseFrame([]byte("*rga#test@2:1'B'"))
	mm := MakeUUIDFrameMultiMap()
	mm.Put(ron.ZERO_UUID, &frame)
	b2, next := mm.Take(ron.ZERO_UUID)
	if &frame != b2 || next != ron.ZERO_UUID {
		t.Fail()
	}
}

// 3-part tables: first all inserts, then all deletes
var rga_3_tests = [][3]string{
	{ // 0+o
		"*rga#textB!",
		"*rga#textB@time'A'",
		"*rga#textB@time!'A'",
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
	{ // s+rms
		"*rga#test@2!@1'A'@2'B'",
		"*rga#test@4!@3:1,@4:2,",
		"*rga#test@4!@1:3'A'@2:4'B'",
	},
	{ // s(rm)+s(rm) merge
		"*rga#test@5!@1:4a'A'@2:5'B'",
		"*rga#test@4!@1:4'A'@3:0'C'",
		"*rga#test@4!@1:4a'A'@3:0'C'@2:5'B'",
	},
	{ // s(rm)+s(rm) merge
		"*rga#test@3!@1:4a'A'@3:0'C'@2:5'B'",
		"*rga#test@4!@1:4a'A'@3:0'C'@4:0'D'@2:5'B'",
		"*rga#test@4!@1:4a'A'@3:0'C'@4'D'@2:5'B'",
	},
	{ // s(rm)+s(rm) merge
		"*rga#test@5!@1:4a'A'@5:0'E'@3:0'C'@2:5'B'",
		"*rga#test@7!@1:4a'A'@6:0'F'@3:7'C'@4:0'D'@2:5'B'",
		"*rga#test@7!@1:4a'A'@6:0'F'@5'E'@3:7'C'@4:0'D'@2:5'B'",
	},
	// TODo concurrent, eclipsed removes
	// TODO: real mess, trees and orphans
}

func TestRGA_Reduce(t *testing.T) {
	for i := 0; i < len(rga_3_tests); i++ {
		test := rga_3_tests[i]
		C := test[2]
		frameA := ron.ParseFrameString(test[0])
		frameB := ron.ParseFrameString(test[1])
		rga := MakeRGAReducer()
		frameC, err := rga.Reduce(frameA, frameB)
		//fmt.Println(frameA.String(), frameB.String(), frameC.String())
		if err != ron.ZERO_UUID {
			t.Fail()
			fmt.Printf("reduction error at %d: %s\n", i, err.String())
		} else if frameC.String() != C {
			t.Fail()
			fmt.Printf("\n-------------------------\nwrong result at %d: \nhave [ %s ]\nneed [ %s ]\n\n", i, frameC.String(), C)
		}

	}
}

// reduceAll: 4-line tables (state, ch1, ch2, result)
