package rdt

import (
	"fmt"
	"github.com/gritzko/ron"
	"testing"
)

// 3-part tables: first all inserts, then all deletes
var lww_3_tests = [][3]string{
	{ // 0+o
		"*lww#test!",
		"*lww#test@time:a'A'",
		"*lww#test@time!:a'A'",
	},
	{ // s+o
		"*lww#test@1!:a'A'",
		"*lww#test@2:b'B'",
		"*lww#test@2!@1:a'A'@2:b'B'",
	},
	{ // o+o
		"*lww#test@1:a'A1'",
		"*lww#test@2:a'A2'",
		"*lww#test@2:d!:a'A2'",
	},
	{ // p+p
		"*lww#test@1:d! :a'A1':b'B1':c'C1'",
		"*lww#test@2:d! :a'A2':b'B2'",
		"*lww#test@2:d!:a'A2':b'B2'@1:c'C1'",
	},
	{
		"*lww#test@0ld!@new:key'new_value'",
		"*lww#test@new:key'new_value'",
		"*lww#test@new!:key'new_value'",
	},
}

func TestLWW_Reduce(t *testing.T) {
	for i := 0; i < len(lww_3_tests); i++ {
		test := lww_3_tests[i]
		C := test[2]
		frameA := ron.ParseFrameString(test[0])
		frameB := ron.ParseFrameString(test[1])
		var lww LWW
		frameC, err := lww.Reduce(frameA, frameB)
		fmt.Println(frameA.String(), frameB.String(), frameC.String())
		if err != ron.ZERO_UUID {
			t.Fail()
			fmt.Printf("reduction error at %d: %s\n", i, err.String())
		} else if frameC.String() != C {
			t.Fail()
			fmt.Printf("\n-------------------------\nwrong result at %d: \nhave [ %s ]\nneed [ %s ]\n\n", i, frameC.String(), C)
		}

	}
}

func TestLWW_FallbackReduce(t *testing.T) {
	frameA := ron.ParseFrameString("*lww#id@1!:b=2:a=1:c=1")
	frameB := ron.ParseFrameString("*lww#id@2!:d'4':c=3")
	var lww LWW
	frameC, err := lww.FallbackReduce(ron.Batch{frameA, frameB})
	if !err.IsZero() || frameC.String()!="*lww#id@2!@1:a=1:b=2@2:c=3:d'4'" {
		t.Fail()
		t.Log(frameC.String())
	}
}
