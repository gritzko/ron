package rdt

import (
	"testing"
	"github.com/gritzko/ron"
)

func TestVV_Reduce(t *testing.T) {
	vvs := []string {
		"*vv#vec@1+a!@,",
		"*vv#vec@3+b!@2+a,@3+b,@1+c,",
		"*vv#vec@4+c!@3+b,@4+c,",
	}
	batch := ron.Batch{}
	for _, s := range vvs {
		batch = append(batch, ron.ParseFrameString(s))
	}
	vv := MakeVVReducer()
	res := vv.Reduce(batch)
	correct := "*vv#vec@4+c!@2+a,@3+b,@4+c,"
	if res.String() != correct {
		t.Logf("got \n%s != \n%s\n", res.String(), correct)
		t.Fail()
	}
}

// TODO epoch
