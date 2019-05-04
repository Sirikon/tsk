package cli

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/Sirikon/tsk/src/info"
)

func getCommands(tskfile *info.TskFile) []*info.Command {
	commands, err := info.GetCommands(tskfile)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return commands
}

func getTskFile() *info.TskFile {
	tskFile, err := info.ReadTskFile()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return tskFile
}

// Index .
func Index() {
	tskFile := getTskFile()
	commands := getCommands(tskFile)

	printHeader(tskFile)

	for _, command := range commands {
		printCommand(command, 0)
	}

	fmt.Println()
}

// Run .
func Run(args []string) {
	tskFile := getTskFile()
	commands := getCommands(tskFile)
	command := findCommand(args, commands)
	if command != nil && command.IsRunnable() {
		runCommand(command, tskFile)
	} else {
		fmt.Println("Command not found")
	}
}

func runCommand(command *info.Command, tskFile *info.TskFile) {
	cmd := exec.Command("sh", command.Path)
	cmd.Dir = tskFile.CWD
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = append(os.Environ(), tskFile.BuildEnvVars()...)
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
