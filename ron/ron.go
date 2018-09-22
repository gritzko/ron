package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		return
	}
	cmd := args[0]
	args = args[1:]
	switch cmd {
	case "convert":
		convert(args)
	default:
		Misunderstood("command unknown")
	}
}

func Misunderstood(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(7)
}
