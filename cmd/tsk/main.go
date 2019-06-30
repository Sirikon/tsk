package main

import (
	"fmt"
	"github.com/sirikon/tsk/src/cli"
	"os"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	c := &cli.CLI{
		CWD:           cwd,
		Out:           os.Stdout,
		ColorsEnabled: true,
	}

	os.Exit(c.Run(os.Args[1:]))
}
