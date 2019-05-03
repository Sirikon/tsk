package cli

import (
	"fmt"
	"log"
	"os"
	"os/exec"

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
	if command != nil && command.IsRunnable() {
		runCommand(command)
	} else {
		fmt.Println("Command not found")
	}
}

func runCommand(command *info.Command) {
	cmd := exec.Command("sh", command.Path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		os.Exit(1)
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
