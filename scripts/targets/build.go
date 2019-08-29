package targets

import (
	"fmt"
	"github.com/magefile/mage/sh"
	"github.com/sirikon/tsk/scripts/targets/utils"
	"os"
)

const entrypoint = "./cmd/tsk"
const output = "./out/tsk"

// Build Generates a new build in `out` folder
func Build() {
	utils.Check(sh.RunV("go", "build", "-ldflags", "-s -w", "-o", output, entrypoint))
	stat, err := os.Stat(output); utils.Check(err)
	fmt.Println("[" + formatBytes(stat.Size()) +"]", output)
}

// https://yourbasic.org/golang/formatting-byte-size-to-human-readable-format/
func formatBytes(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "KMGTPE"[exp])
}
