package rdt

import (
	"github.com/gritzko/ron"
	"testing"
)

func TestSet_Reduce(t *testing.T) {
	tests := [][]string{
		{
			"*set#test1@1=1",
			"*set#test1@2=2",
			"*set#test1@2:d!:0=2@1=1",
		},
		{
			"*set#test1@1!@=1",
			"*set#test1@2:1;",
			"*set#test1@2!:1,",
		},
		{
			"*set#test1@2!@=2@1=1",
			"*set#test1@5!@=5@3:2,@4:1,",
			"*set#test1@5!@=5@3:2,@4:1,",
		},
		{
			"*set#test1@2!@=2@1=1",
			"*set#test1@3!@:2,@4:1,",
			"*set#test1@5!@=5",
			"*set#test1@5!@=5@3:2,@4:1,",
		},
		{
			"*set#test1@3!@:2,@4:1,",
			"*set#test1@5!@=5",
			"*set#test1@2!@=2@1=1",
			"*set#test1@2!@5=5@3:2,@4:1,",
		},
	}
	cs := MakeSetReducer()
	for i, test := range tests {
		inputs := test[:len(test)-1]
		batch := ron.ParseStringBatch(inputs)
		result := cs.Reduce(batch)
		if result.String() != test[len(test)-1] {
			t.Logf("%d set reduce fail, got\n'%s' want\n'%s'\n", i, result.String(), test[len(test)-1])
			t.Fail()
		}
	}
}
