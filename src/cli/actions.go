package cli

import (
	"strings"
	"io"
	"os"
	"os/exec"
	"runtime"
	"log"
	"bufio"

	"github.com/sirikon/tsk/src/application"
)

type CLI struct {
	CWD           string
	Out           io.Writer
	Err           io.Writer
	In            io.Reader
	ColorsEnabled bool
}

func (c *CLI) Run(args []string) int {
	if len(args) == 0 {
		return c.index()
	}
	return c.runCommand(args)
}

// Index .
func (c *CLI) index() (result int) {
	defer HandlePanic(&result, c.Out)
	p := c.getPrinter()
	project := c.getProject()

	p.header(project.TskFile)
	for _, command := range project.Commands {
		p.command(command, project, 0)
	}
	p.emptyLine()

	return 0
}

func (c *CLI) runCommand(args []string) (result int) {
	defer HandlePanic(&result, c.Out)
	p := c.getPrinter()
	project := c.getProject()

	command, remainingArgs := findCommand(args, project.Commands)

	if command == nil || !command.Runnable {
		p.line("Command not found")
		return 1
	}

	return c.execCommand(command, remainingArgs, project)
}

func (c *CLI) getProject() *application.Project {
	project, err := application.GetProject(c.CWD); HandleErr(err)
	return project
}

func (c *CLI) getPrinter() *printer {
	return &printer{
		colorsEnabled: c.ColorsEnabled,
		out:           c.Out,
	}
}

func (c *CLI) execCommand(command *application.Command, args []string, project *application.Project) int {
	completeArgs := append([]string{command.Path}, args...)

	cli := "sh"

	if runtime.GOOS == "windows" {
		cli = "cmd.exe"
		completeArgs = append([]string{"/c"}, completeArgs...)
	}

	if runtime.GOOS != "windows" {
		interp := findInterpreter(command)
		cli = interp.Command
		completeArgs = append(interp.Arguments, completeArgs...)
	}

	cmd := exec.Command(cli, completeArgs...)
	cmd.Dir = project.RootFolder
	cmd.Stdout = c.Out
	cmd.Stderr = c.Err
	cmd.Stdin = c.In
	cmd.Env = append(os.Environ(), buildEnvVars(project.TskFile)...)
	err := cmd.Run()
	if err != nil {
		return 1
	}
	return 0
}


// BuildEnvVars .
func buildEnvVars(tskFile *application.TskFile) []string {
	result := make([]string, 0)
	for key, value := range tskFile.Env {
		result = append(result, key+"="+value)
	}
	return result
}

func findCommand(args []string, commands []*application.Command) (command *application.Command, remainingArgs []string) {
	if len(args) == 0 {
		return nil, args
	}

	commandName := args[0]

	for _, c := range commands {
		if c.Name == commandName {
			if c.Runnable {
				return c, args[1:]
			}
			return findCommand(args[1:], c.SubCommands)
		}
	}

	return nil, args
}


type interpreter struct {
	Command string
	Arguments []string
}
// findInterpreter is used to determine the executable to use for a choosen script (Linux/Unix).
func findInterpreter(command *application.Command) interpreter {

	interp := "sh"
	arguments := []string{}

	file, err := os.Open(command.Path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	// Parsing line
	if line[2:] == "#!" {
		line = strings.TrimLeft(line, "#!")
		parts := strings.Split(line, " ")
		interp = parts[0]

		if len(parts) >= 2 {
			arguments = append(arguments, parts[2:]...)
		}
	}


	// Return interpreter
	return interpreter{interp, arguments}
}