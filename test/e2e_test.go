package test

import (
	"bytes"
	"github.com/sirikon/tsk/src/cli"
	"gotest.tools/assert"
	"io"
	"os"
	"path"
	"testing"
)

func TestCLI_Index_When_No_Args(t *testing.T) {
	assertRun(t, []string{}, "", 0, `
  tsk Testing
  Usage: tsk [command] <subcommands...>

    build scripts/build.sh
    docs scripts/docs
      build scripts/docs/build.sh
    fail scripts/fail.sh
    vim scripts/vim.sh

`)
}

func TestCLI_Run_Command_Success(t *testing.T) {
	cwd, err := os.Getwd()
	handleError(err)

	assertRun(t, []string{"build"}, "", 0, `Started command 'build'
Testing TEST_VAR
Hello World!
`+path.Join(cwd, "test-folder")+`
`)
}

func TestCLI_Run_Command_Fail(t *testing.T) {
	assertRun(t, []string{"fail"}, "", 1, "I'm gonna fail\n")
}

func TestCLI_Run_Deep_Command(t *testing.T) {
	assertRun(t, []string{"docs", "build"}, "", 0, "docs build\n")
}

func TestCLI_Run_Unknown_Command(t *testing.T) {
	assertRun(t, []string{"unknown"}, "", 1, "Command not found\n")
}

func TestCLI_Run_NonRunnable_Command(t *testing.T) {
	assertRun(t, []string{"docs"}, "", 1, "Command not found\n")
}

func buildSut(out io.Writer, err io.Writer, in io.Reader) *cli.CLI {
	cwd, error := os.Getwd()
	handleError(error)
	return &cli.CLI{
		ColorsEnabled: false,
		Out:           out,
		Err:           err,
		In:            in,
		CWD:           path.Join(cwd, "test-folder"),
	}
}

func assertRun(t *testing.T, args []string, input string, expectedResultCode int, expectedOutput string) {
	out := &bytes.Buffer{}
	err := &bytes.Buffer{}
	in := &bytes.Buffer{}
	in.Write([]byte(input))
	sut := buildSut(out, err, in)
	code := sut.Run(args)
	assert.Equal(t, expectedResultCode, code)
	assert.Equal(t, expectedOutput, out.String())
}

func handleError(err error)  {
	if err != nil {
		panic(err)
	}
}
