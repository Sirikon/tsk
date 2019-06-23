package main

import (
	"github.com/Sirikon/tsk/src/cli"
	"log"
	"os"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	c := &cli.CLI{
		CWD:           cwd,
		Out:           os.Stdout,
		ColorsEnabled: true,
	}

	c.Run(os.Args[1:])
}
