package main

import (
	"fmt"
	"github.com/gritzko/ron"
	"bufio"
	"os"
)

func convert (args []string) {
	if len(args)==0 {
		Misunderstood("nothing to convert")
		return
	}
	noun := args[0]
	args = args[1:]
	switch noun {
	case "int":
		convert_int(args);
	case "int4":
		convert_int4(args);
	case "uuid":
		convert_uuid(args);
	default:
		Misunderstood("convert what?")
	}
}

func convert_int (args []string) {
	if len(args)==0 {
		Misunderstood("nothing to convert")
		return
	}
	if args[0]=="-" {
		args = read_stdin()
	}
	for _, a := range args {
		var i uint64
		_, err := fmt.Sscanf(a, "%d", &i)
		if err != nil {
			Misunderstood(err.Error())
			return
		}
		var out [12]byte
		bi := ron.FormatInt(out[:0], i)
		fmt.Println(string(bi))
	}
}

func convert_int4 (args []string) {
	if len(args)==0 {
		Misunderstood("nothing to convert")
		return
	}
	if args[0]=="-" {
		args = read_stdin()
	}
	for len(args)>0 {
		var ints [4]uint64
		i := 0
		for ; i<4 && i<len(args); i++ {
			_, err := fmt.Sscanf(args[i], "%d", &ints[i])
			if err != nil {
				Misunderstood(err.Error())
				return
			}
		}
		args = args[i:]
		uuid := ron.NewRonUUID(uint(ints[2]), uint(ints[0]), ints[1], ints[3])
		fmt.Println(uuid.String())
	}
}

func read_stdin () (ret []string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}
	return
}

func convert_uuid (args []string) {
	if len(args)==0 {
		Misunderstood("nothing to convert")
		return
	}
	uuids := args[0:1]
	args = args[1:]
	if uuids[0]=="-" {
		uuids = read_stdin()
	}
	to := "int"
	if len(args)>1 && args[0]=="to" {
		to = args[1]
		args = args[2:]
	}
	for _, a := range uuids {
		uuid, err := ron.ParseUUIDString(a)
		if err != nil {
			Misunderstood(err.Error()+" with "+a)
			return
		}
		switch to {
		case "int":
			fmt.Printf("%d %d\n", uuid[0], uuid[1])
		case "int4":
			fmt.Printf("%d %d %d %d\n", uuid.Variety(), uuid.Value(), uuid.Scheme(), uuid.Origin())
		default:
			Misunderstood("convert to what?")
			return
		}
	}
}
