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

func (c *CLI) Run(args []string) int {
	if len(args) == 0 {
		c.index()
		return 0
	}
	return c.runCommand(args)
}

// Index .
func (c *CLI) index() {
	tskFile := c.getTskFile()
	commands := getCommands(tskFile)
	p := c.getPrinter()

	p.header(tskFile)

	for _, command := range commands {
		p.command(command, tskFile, 0)
	}

	fmt.Fprintln(c.Out)
}

// RunCommand .
func (c *CLI) runCommand(args []string) int {
	tskFile := c.getTskFile()
	commands := getCommands(tskFile)
	command := findCommand(args, commands)

	if command == nil || !command.IsRunnable() {
		_, _ = fmt.Fprintln(c.Out, "Command not found")
	}

	return c.execCommand(command, tskFile)
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

func (c *CLI) execCommand(command *info.Command, tskFile *info.TskFile) int {
	cmd := exec.Command("sh", command.Path)
	cmd.Dir = tskFile.CWD
	cmd.Stdout = c.Out
	cmd.Stderr = os.Stderr
	cmd.Env = append(os.Environ(), tskFile.BuildEnvVars()...)
	err := cmd.Run()
	if err != nil {
		return 1
	}
	return 0
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
