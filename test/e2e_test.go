package test

import (
	"github.com/sirikon/tsk/test/utils"
	"os"
	"path"
	"testing"
)

func TestCLI_Index_When_No_Args(t *testing.T) {

	expectedOutput := `
  tsk Testing
  Usage: tsk [command] <subcommands...>

    build scripts/build.sh
    docs scripts/docs
      build scripts/docs/build.sh
    fail scripts/fail.sh
    interactive scripts/interactive.sh
    params scripts/params.sh

`

	utils.AssertRun(t).
		ExpectedOutput(expectedOutput).
		Run()
}

func TestCLI_Run_Command_Success(t *testing.T) {
	cwd, err := os.Getwd()
	utils.HandleError(err)

	expectedOutput := `Started command 'build'
Testing TEST_VAR
Hello World!
`+path.Join(cwd, "test-folder")+`
`

	utils.AssertRun(t).
		Args("build").
		ExpectedOutput(expectedOutput).
		Run()
}

func TestCLI_Run_Command_Fail(t *testing.T) {
	utils.AssertRun(t).
		Args("fail").
		ExpectResultCode(1).
		ExpectedOutput("I'm gonna fail\n").
		Run()
}

func TestCLI_Run_Deep_Command(t *testing.T) {
	utils.AssertRun(t).
		Args("docs", "build").
		ExpectedOutput("docs build\n").
		Run()
}

func TestCLI_Run_Unknown_Command(t *testing.T) {
	utils.AssertRun(t).
		Args("unknown").
		ExpectResultCode(1).
		ExpectedOutput("Command not found\n").
		Run()
}

func TestCLI_Run_NonRunnable_Command(t *testing.T) {
	utils.AssertRun(t).
		Args("docs").
		ExpectResultCode(1).
		ExpectedOutput("Command not found\n").
		Run()
}

func TestCLI_Run_Interactive_Command(t *testing.T) {
	utils.AssertRun(t).
		Args("interactive").
		Input("hello world").
		ExpectedOutput("hello world").
		Run()
}

/*func TestCLI_Run_Command_With_Parameters(t *testing.T)  {
	assertRun(t, []string{"params", "hello"}, "", 0, "hello")
}*/
