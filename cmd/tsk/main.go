package main

import (
	"fmt"
	"log"

	"github.com/Sirikon/tsk/src"
)

func main() {
	commands, err := src.GetCommands()
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, command := range commands {
		fmt.Println(command)
	}
}
