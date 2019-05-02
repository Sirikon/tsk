package main

import (
	"fmt"
	"os"

	"github.com/Sirikon/tsk/src/cli"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		cli.Index()
		return
	}

	fmt.Println(args)
}
