package cli

import (
	"fmt"
	"github.com/Sirikon/tsk/src/info"
	"github.com/Sirikon/tsk/src/utils"
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

func (p *printer) header(tskFile *info.TskFile) {
	au := p.getAurora()

	_, _ = fmt.Fprintln(p.out)

	_, _ = fmt.Fprint(p.out, baseSpacing)
	_, _ = fmt.Fprintln(p.out, au.Bold(au.Magenta("tsk")), au.Bold(tskFile.Name))

	_, _ = fmt.Fprint(p.out, baseSpacing)
	_, _ = fmt.Fprintln(p.out, au.Faint("Usage: tsk [command] <subcommands...>"))

	_, _ = fmt.Fprintln(p.out)
}

// PrintCommand .
func (p *printer) command(command *info.Command, tskFile *info.TskFile, level int) {
	au := p.getAurora()

	_, _ = fmt.Fprint(p.out, baseSpacing)
	_, _ = fmt.Fprint(p.out, utils.PadLeft("", level+1, "  "))
	_, _ = fmt.Fprint(p.out, au.Bold(au.Cyan(command.Name)))

	relativePath, err := filepath.Rel(tskFile.CWD, command.Path)
	if err != nil {
		relativePath = command.Path
	}

	_, _ = fmt.Fprint(p.out, au.Faint(" "+relativePath))
	_, _ = fmt.Fprintln(p.out)

	for _, c := range command.Subcommands {
		p.command(c, tskFile, level+1)
	}
}
