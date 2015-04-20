package main

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

func main() {
	var uid uint64 = 0
	var gid uint64 = 0

	if path.Base(os.Args[0]) == "echo" {
		fmt.Print(strings.Join(os.Args[1:], " "))
		fmt.Print("\n")
		os.Exit(0)
	}

	if len(os.Args) > 1 {
		uid, _ = strconv.ParseUint(os.Args[1], 10, 0)
	}
	if len(os.Args) > 2 {
		gid, _ = strconv.ParseUint(os.Args[2], 10, 0)
	}

	err := os.Chown("/persistent", int(uid), int(gid))
	if err != nil {
		panic(err)
	}
}
