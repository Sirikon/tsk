package cli

import (
	"fmt"

	"github.com/Sirikon/tsk/src/info"
	"github.com/Sirikon/tsk/src/utils"
	"github.com/logrusorgru/aurora"
)

const baseSpacing = "  "

func printHeader(tskFile *info.TskFile) {
	fmt.Println()

	fmt.Print(baseSpacing)
	fmt.Println(aurora.Bold(aurora.Magenta("tsk")), aurora.Bold(tskFile.Name))

	fmt.Print(baseSpacing)
	fmt.Println(aurora.Faint("Usage: tsk [command] <subcommands...>"))

	fmt.Println()
}

// PrintCommand .
func printCommand(command *info.Command, level int) {
	fmt.Print(baseSpacing)
	fmt.Print(utils.PadLeft("", level+1, "  "))
	fmt.Print(aurora.Bold(aurora.Cyan(command.Name)))
	fmt.Print(aurora.Faint(" " + command.Path))
	fmt.Println()

	for _, c := range command.Subcommands {
		printCommand(c, level+1)
	}
}
