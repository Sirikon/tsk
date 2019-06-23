package cli

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/Sirikon/tsk/src/info"
)

type CLI struct {
	CWD           string
	Out           io.Writer
	ColorsEnabled bool
}

func (c *CLI) Run(args []string) {
	if len(args) == 0 {
		c.Index()
		return
	}
	c.RunCommand(args)
}

// Index .
func (c *CLI) Index() {
	tskFile := c.getTskFile()
	commands := getCommands(tskFile)
	p := c.getPrinter()

	p.header(tskFile)

	for _, command := range commands {
		p.command(command, tskFile, 0)
	}

	fmt.Println()
}

// RunCommand .
func (c *CLI) RunCommand(args []string) {
	tskFile := c.getTskFile()
	commands := getCommands(tskFile)
	command := findCommand(args, commands)
	if command != nil && command.IsRunnable() {
		runCommand(command, tskFile)
	} else {
		fmt.Println("Command not found")
	}
}

func (c *CLI) getPrinter() *printer {
	return &printer{
		colorsEnabled: c.ColorsEnabled,
		out:           c.Out,
	}
}

func getCommands(tskfile *info.TskFile) []*info.Command {
	commands, err := info.GetCommands(tskfile)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return commands
}

func (c *CLI) getTskFile() *info.TskFile {
	tskFile, err := info.ReadTskFile(c.CWD)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return tskFile
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
