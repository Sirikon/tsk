package targets

import (
	"github.com/magefile/mage/sh"
	"github.com/sirikon/tsk/scripts/targets/utils"
)

// Test Runs the tests
func Test() {
	utils.Check(sh.RunV("go", "test", "./test"))
}
