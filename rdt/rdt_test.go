package rdt

import (
	"github.com/gritzko/ron"
	"io/ioutil"
	"testing"
)

func RunRONTest(t *testing.T, reducer ron.Reducer, scriptFile string) {
	scriptBuf, err := ioutil.ReadFile(scriptFile)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
		return
	}
	scriptFrame := ron.ParseFrame(scriptBuf)
	script := scriptFrame.Split()
	if len(script) == 0 {
		t.Log("script not parsed")
		t.Fail()
		return
	} // todo parse err

	for len(script) != 0 {

		q := script[0]
		script = script[1:]
		if !q.IsComment() || !q.IsQuery() {
			t.Log("no test header")
			t.Fail()
			break
		}

		t.Log(q.GetString(0), "?..")

		l := 0
		for len(script) > l && !script[l].IsComment() {
			l++
		}
		inputs := script[:l]
		script = script[l:]

		output := reducer.Reduce(inputs)
		outputs := output.Split()
		//t.Log("split", output.String(), outputs.String())

		if len(script) == 0 || !script[0].IsHeader() {
			t.Log("no output specified")
			t.Fail()
			break
		}
		a := script[0]
		script = script[1:]

		l = 0
		for len(script) > l && !script[l].IsComment() {
			l++
		}
		correct := script[:l]
		script = script[l:]

		if !correct.Equal(outputs) {
			t.Fail()
			t.Logf("FAILS to produce %s:\n%s\nshould be\n%s\n", a.GetString(0), outputs.String(), correct.String())
			t.Logf("exact input, output:\n%s\n%s\n\n", inputs.String(), output.String())
		} else {
			t.Log("produces", a.GetString(0), "!\n")
		}

	}
}
