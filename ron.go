package ron

import (
	"flag"
	"fmt"
	"os"
)

var mode_stamp = flag.Bool("stamp", false, "timestamp frames ($1, $2 etc)")
var mode_now = flag.Bool("now", false, "print a timestamp")
var mode_unzip = flag.Bool("unzip", true, "don't zip the resulting frame")
var mode_trim = flag.Bool("trim", false, "trim the resulting frame (if unzipped)")
var see_stamp = flag.String("see", "", "see a timestamp")
var me_stamp = flag.String("me", "", "set the replica id")

const MAXFILES int = 10

func readStdin() [][]byte {
	return [][]byte{}
}

func readFiles(inputs []string) [][]byte {
	var files [MAXFILES]*os.File
	var err error
	for i, name := range inputs {
		files[i], err = os.Open(name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "file '%s' open error :%s", err.Error())
			os.Exit(1)
		}
	}

	var frames [][]byte = make([][]byte, len(inputs))
	for i, file := range files {
		stat, _ := file.Stat()
		frames[i] = make([]byte, stat.Size())
		l, err := file.Read(frames[i])
		if err != nil || int64(l) != stat.Size() {
			fmt.Fprintf(os.Stderr, "file %s read fail %s", inputs[i], err.Error())
			os.Exit(2)
		}
	}

	return frames
}

// Usage:  ron file1 file2 > file3 (reduces k frames)
//			ron <stdin, double-newline separated>
func main() {
	/*
		flag.Parse()
		inputs := flag.Args()
		var bufs [][]byte

		if len(inputs) > 0 {
			bufs = readFiles(inputs)
		} else {
			bufs = readStdin()
		}

		var frames []Frame = make([]Frame, len(bufs))
		for i, b := range bufs {
			frames[i] = Frame{body: b}
		}

		omni := OmniReducer{}

		result, err := omni.ReduceAll(frames)

		if err != ZERO_UUID {
			fmt.Fprintf(os.Stderr, "reducer error: %s", err.String)
		} else {
			os.Stdout.Write(result.body)
		}
	*/
}
