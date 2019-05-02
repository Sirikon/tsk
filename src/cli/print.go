package cli

import (
	"fmt"

	"github.com/Sirikon/tsk/src"
	"github.com/Sirikon/tsk/src/utils"
)

// PrintCommand .
func printCommand(command src.Command, level int) {
	fmt.Println(utils.PadLeft("", level, "  ") + command.Name)
	for _, c := range command.Subcommands {
		printCommand(c, level+1)
	}
}
