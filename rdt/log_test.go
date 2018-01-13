package rdt

import (
	"github.com/gritzko/ron"
	"testing"
)

func TestLog_Reduce(t *testing.T) {
	vvs := []string{
		"*lww#id!@2+B:b=2@1+A:a=1",
		//"*lww#id!",
		"*lww#id@3+C:c=3@1+A:a=1",
	}

	// FIXME  *log  vs  *lww
	// ack:   time-orig
	// monopatch

	batch := ron.ParseStringBatch(vvs)
	log := MakeLogReducer()
	res := log.Reduce(batch)
	correct := "*lww#id@3+C!:c=3@2+B:b=2@1+A:a=1"
	if res.String() != correct {
		t.Logf("got \n%s != \n%s\n", res.String(), correct)
		t.Fail()
	}
}
