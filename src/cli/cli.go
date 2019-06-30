package cli

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/sirikon/tsk/src/info"
)

type CLI struct {
	CWD           string
	Out           io.Writer
	ColorsEnabled bool
}

func (c *CLI) Run(args []string) int {
	if len(args) == 0 {
		return c.index()
	}
	return c.runCommand(args)
}

// Index .
func (c *CLI) index() int {
	tskFile, err := c.getTskFile()
	if err != nil {
		_, _ = fmt.Fprintln(c.Out, err)
		return 1
	}

	commands, err := getCommands(tskFile)
	if err != nil {
		_, _ = fmt.Fprintln(c.Out, err)
		return 1
	}

	p := c.getPrinter()

	p.header(tskFile)

	for _, command := range commands {
		p.command(command, tskFile, 0)
	}

	_, _ = fmt.Fprintln(c.Out)
	return 0
}

func (c *CLI) runCommand(args []string) int {
	tskFile, err := c.getTskFile()
	if err != nil {
		_, _ = fmt.Fprintln(c.Out, err)
		return 1
	}

	commands, err := getCommands(tskFile)
	if err != nil {
		_, _ = fmt.Fprintln(c.Out, err)
		return 1
	}

	command := findCommand(args, commands)

	if command == nil || !command.IsRunnable() {
		_, _ = fmt.Fprintln(c.Out, "Command not found")
		return 1
	}

	return c.execCommand(command, tskFile)
}

func (c *CLI) getPrinter() *printer {
	return &printer{
		colorsEnabled: c.ColorsEnabled,
		out:           c.Out,
	}
}

func getCommands(tskfile *info.TskFile) ([]*info.Command, error) {
	commands, err := info.GetCommands(tskfile)
	if err != nil {
		return nil, err
	}
	return commands, nil
}

func (c *CLI) getTskFile() (*info.TskFile, error) {
	tskFile, err := info.ReadTskFile(c.CWD)
	if err != nil {
		return nil, err
	}
	return tskFile, nil
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
