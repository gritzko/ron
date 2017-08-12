package RDT

import (
	"testing"
	"fmt"
	"github.com/gritzko/RON"
)

// 3-part tables: first all inserts, then all deletes
var lww_3_tests = [][3]string{
	{ // s+o
		"*lww#test@1!:a'A'",
		"*lww#test@2:b'B'.",
		"*lww#test@2!@1:a'A'@2:b'B'",
	},
	{ // o+o
		"*lww#test@1:a'A1'.",
		"*lww#test@2:a'A2'.",
		"*lww#test@2;:a'A2'",
	},
	{ // p+p
		"*lww#test@1; :a'A1' :b'B1' :c'C1'",
		"*lww#test@2; :a'A2' :b'B2'",
		"*lww#test@2;:a'A2':b'B2'@1:c'C1'",
	},
}

func TestLWW_Reduce(t *testing.T) {
	for i := 0; i < len(lww_3_tests); i++ {
		test := lww_3_tests[i]
		C := test[2]
		frameA, _ := RON.Parse(test[0])
		frameB, _ := RON.Parse(test[1])
		var lww LWW
		frameC, err := lww.Reduce(frameA, frameB)
		if err != RON.ZERO_UUID {
			t.Fail()
			fmt.Printf("reduction error at %d: %s\n", i, err.String())
		} else if frameC.String() != C {
			t.Fail()
			fmt.Printf("\n-------------------------\nwrong result at %d: \nhave [ %s ]\nneed [ %s ]\n\n", i, frameC.String(), C)
		}

	}
 }
