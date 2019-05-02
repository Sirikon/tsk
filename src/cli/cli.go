package cli

import (
	"log"

	"github.com/Sirikon/tsk/src"
)

// Index .
func Index() {
	commands, err := src.GetCommands()
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, command := range commands {
		printCommand(command, 0)
	}
}
