package cli

import (
	"fmt"
	"github.com/sirikon/tsk/src/application"
	"github.com/sirikon/tsk/src/utils"
	"github.com/logrusorgru/aurora"
	"io"
	"path/filepath"
)

const baseSpacing = "  "

type printer struct {
	out           io.Writer
	colorsEnabled bool
}

func (p *printer) getAurora() aurora.Aurora {
	return aurora.NewAurora(p.colorsEnabled)
}

func (p *printer) header(tskFile *application.TskFile) {
	au := p.getAurora()

	_, _ = fmt.Fprintln(p.out)

	_, _ = fmt.Fprint(p.out, baseSpacing)
	_, _ = fmt.Fprintln(p.out, au.Bold(au.Magenta("tsk")), au.Bold(tskFile.Name))

	_, _ = fmt.Fprint(p.out, baseSpacing)
	_, _ = fmt.Fprintln(p.out, au.Faint("Usage: tsk [command] <subcommands...>"))

	_, _ = fmt.Fprintln(p.out)
}

// PrintCommand .
func (p *printer) command(command *application.Command, project *application.Project, level int) {
	au := p.getAurora()

	_, _ = fmt.Fprint(p.out, baseSpacing)
	_, _ = fmt.Fprint(p.out, utils.PadLeft("", level+1, "  "))
	_, _ = fmt.Fprint(p.out, au.Bold(au.Cyan(command.Name)))

	relativePath, err := filepath.Rel(project.RootFolder, command.Path)
	if err != nil {
		relativePath = command.Path
	}

	_, _ = fmt.Fprint(p.out, au.Faint(" "+relativePath))
	_, _ = fmt.Fprintln(p.out)

	for _, c := range command.SubCommands {
		p.command(c, project, level+1)
	}
}

func (p *printer) line(text string) {
	_, _ = fmt.Fprintln(p.out, text)
}

func (p *printer) emptyLine() {
	_, _ = fmt.Fprintln(p.out)
}
