package rdt

import (
	"github.com/gritzko/ron"
	"testing"
)

func TestCausalSet_Reduce(t *testing.T) {
	tests := [][]string{
		{
			"*cas#test1@1=1",
			"*cas#test1@2=2",
			"*cas#test1@2:d!:0=2@1=1",
		},
		{
			"*cas#test1@1=1",
			"*cas#test1@2:1;",
			"*cas#test1@2:d!",
		},
		{
			"*cas#test1@2!@=2@1=1",
			"*cas#test1@5!@=5@3:2,@4:1,",
			"*cas#test1@5!@=5",
		},
		{
			"*cas#test1@2!@=2@1=1",
			"*cas#test1@3!@:2,@4:1,",
			"*cas#test1@5!@=5",
			"*cas#test1@5!@=5",
		},
	}
	cs := MakeCausalSetReducer()
	for i, test := range tests {
		inputs := test[:len(test)-1]
		batch := ron.ParseStringBatch(inputs)
		result := cs.Reduce(batch)
		if result.String() != test[len(test)-1] {
			t.Logf("%d cset reduce fail, got\n'%s' want\n'%s'\n", i, result.String(), test[len(test)-1])
			t.Fail()
		}
	}
}
