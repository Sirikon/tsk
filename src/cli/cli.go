package cli

import (
	"fmt"
	"log"

	"github.com/Sirikon/tsk/src/info"
)

func getCommands() []*info.Command {
	commands, err := info.GetCommands()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return commands
}

// Index .
func Index() {
	commands, err := info.GetCommands()
	if err != nil {
		log.Fatal(err)
		return
	}

	printHeader()

	for _, command := range commands {
		printCommand(command, 0)
	}

	fmt.Println()
}

// Run .
func Run(args []string) {
	commands := getCommands()
	command := findCommand(args, commands)
	if command != nil {
		fmt.Println(command.Path)
	} else {
		fmt.Println("Command not found")
	}
}

func findCommand(args []string, commands []*info.Command) *info.Command {
	if len(args) == 0 {
		return nil
	}
	commandName := args[0]
	isLast := len(args) == 1
	for _, c := range commands {
		if c.Name == commandName {
			if isLast {
				return c
			}

			return findCommand(args[1:], c.Subcommands)
		}
	}

	return nil
}
