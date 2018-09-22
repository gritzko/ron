package rdt

import (
	"testing"

	"github.com/gritzko/ron"
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
			"*set#test1@3:1;",
			"*set#test1@4:2;",
			"*set#test1@4:d!:2,@3:1,",
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
		{
			"*set#mice@1YKDY54a01+1YKDY5!>mouse$1YKDY5",
			"*set#mice@1YKDXO3201+1YKDXO?!@>mouse$1YKDXO@(WBF901(WBY>mouse$1YKDWBY@[67H01[6>mouse$1YKDW6@(Uh4j01(Uh>mouse$1YKDUh@(S67V01(S6>mouse$1YKDS6@(Of(N3:1YKDN3DS01+1YKDN3,@(MvBV01(IuJ:0>mouse$1YKDIuJ@(LF:1YKDIuEY01+1YKDIuJ,:{A601,@(Io5l01[oA:0>mouse$1YKDIoA@[l7_01[l>mouse$1YKDIl@(57(4B:1YKD4B3f01+1YKD4B,@(0bB401+1YKCsd:0>mouse$1YKCsd@1YKCu6+:1YKCsd7Q01+1YKCsd,",
			"*set#mice@1YKDXO3201+1YKDXO!@(Y54a01(Y5>mouse$1YKDY5@(XO3201(XO>mouse$1YKDXO@(WBF901(WBY>mouse$1YKDWBY@[67H01[6>mouse$1YKDW6@(Uh4j01(Uh>mouse$1YKDUh@(S67V01(S6>mouse$1YKDS6@(Of(N3:1YKDN3DS01+1YKDN3,@(MvBV01(IuJ:0>mouse$1YKDIuJ@(LF:1YKDIuEY01+1YKDIuJ,:{A601,@(Io5l01[oA:0>mouse$1YKDIoA@[l7_01[l>mouse$1YKDIl@(57(4B:1YKD4B3f01+1YKD4B,@(0bB401+1YKCsd:0>mouse$1YKCsd@1YKCu6+:1YKCsd7Q01+1YKCsd,",
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

func TestSet_Reduce_Basic(t *testing.T) {
	RunRONTest(
		t,
		MakeSetReducer(),
		"test/01-set-basic.ron",
	)
}
