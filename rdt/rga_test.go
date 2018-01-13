package rdt

import (
	"github.com/gritzko/ron"
	"testing"
	//	"fmt"
	//	"math/rand"
	//	"time"
)

func TestRGA_Primers(t *testing.T) {
	RunRONTest(
		t,
		MakeRGAReducer(),
		"test/00-rga-basic.ron",
	)
}

func TestRGA_Mapper(t *testing.T) {
	frame := "*rga#1UQ8p+bart@1UQ8yk+lisa!" +
		"@(s+bart'H'@[r'e'@(t'l'@[T'l'@[i'o'" +
		"@(w+lisa' '@(x'w'@(y'o'@[1'r'@{a'l'@[2'd'@[k'!'"
	right := "Hello world!"
	var txt TxtMapper
	hello := txt.Map(ron.Batch{ron.ParseFrameString(frame)})
	if hello != right {
		t.Fail()
		t.Logf("'%s' != '%s'", hello, right)
	}
}

/*
func TestHelloWorld(t *testing.T) {

	var src = [13]string {
		"*rga#1UQ8p+bart!",
		"*rga#1UQ8p+bart@1UQ8s+bart:0'H'",
		"*rga#1UQ8p+bart@1UQ8sr+bart:1UQ8s+bart'e'",
		"*rga#1UQ8p+bart@1UQ8t+bart:1UQ8sr+bart'l'",
		"*rga#1UQ8p+bart@1UQ8tT+bart:1UQ8t+bart'l'",
		"*rga#1UQ8p+bart@1UQ8ti+bart:1UQ8tT+bart'o'",
		"*rga#1UQ8p+bart@1UQ8w+lisa:1UQ8ti+bart' '",
		"*rga#1UQ8p+bart@1UQ8x+lisa:1UQ8w+lisa'w'",
		"*rga#1UQ8p+bart@1UQ8y+lisa:1UQ8x+lisa'o'",
		"*rga#1UQ8p+bart@1UQ8y1+lisa:1UQ8y+lisa'r'",
		"*rga#1UQ8p+bart@1UQ8y1a+lisa:1UQ8y1+lisa'l'",
		"*rga#1UQ8p+bart@1UQ8y2+lisa:1UQ8y1a+lisa'd'",
		"*rga#1UQ8p+bart@1UQ8yk+lisa:1UQ8y2+lisa'!'",
	}

	count := 1000

	for i:=0; i<count; i++ {

		seed := time.Now().UnixNano()
		//fmt.Printf("ACI test seed %d\n", seed)
		seed = 1512325615325939065
		r := rand.New(rand.NewSource(seed))

		data := make([]string, 13)
		perm := r.Perm(len(src))
		for i, v := range perm { // this way we test COMMUTATIVITY
			data[v] = src[i]
		}
		frames := []ron.Frame{}
		for i := 0; i < len(data); i++ {
			frames = append(frames, ron.ParseFrameString(data[i]))
		}

		rga := MakeRGAReducer()

		for len(frames) > 1 {
			from := int(r.Uint32()) % len(frames)
			till := int(r.Uint32()) % (len(frames) - from)
			till += from + 1
			//fmt.Printf("\nReduce %d..%d of %d\n", from, till, len(frames))
			for _, f := range frames[from:till] {
				fmt.Printf("+ %s\n", f.String())
			}
			// this way we test ASSOCIATIVITYz
			frameC := rga.Reduce(frames[from:till]).Reformat(ron.FRAME_FORMAT_LIST)
			fmt.Printf("---\n%s\n\n", frameC.String())
			f := make(ron.Batch, 0, len(frames))
			f = append(f, frames[:from]...)
			f = append(f, frameC)
			f = append(f, frames[till:]...)
			frames = f
		}

		right := "Hello world!"
		var txt TxtMapper
		hello := txt.Map(ron.Batch{frames[0]})
		if hello != right {
			t.Fail()
			t.Logf("'%s' != '%s', seed %d", hello, right, seed)
			break
		} else {
			t.Logf("%d %d %s\n", i, seed, hello)
		}

	}
}
*/
