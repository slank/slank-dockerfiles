package main

import (
	"os"
	"strconv"
)

func main() {
	var uid uint64 = 0
	var gid uint64 = 0

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
