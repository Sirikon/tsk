package test

import (
	"bytes"
	"github.com/Sirikon/tsk/src/cli"
	"gotest.tools/assert"
	"io"
	"log"
	"os"
	"path"
	"testing"
)

func TestCLI_Index_When_No_Args(t *testing.T) {
	assertRun(t, []string{}, 0, `
  tsk Testing
  Usage: tsk [command] <subcommands...>

    build scripts/build.sh
    docs scripts/docs
      build scripts/docs/build.sh
    fail scripts/fail.sh

`)
}

func TestCLI_Run_Command_Success(t *testing.T) {
	cwd, err := os.Getwd()
	handleError(err)

	assertRun(t, []string{"build"}, 0, `Started command 'build'
Testing TEST_VAR
Hello World!
`+path.Join(cwd, "test-folder")+`
`)
}

func TestCLI_Run_Command_Fail(t *testing.T) {
	assertRun(t, []string{"fail"}, 1, "I'm gonna fail\n")
}

func TestCLI_Run_Deep_Command(t *testing.T) {
	assertRun(t, []string{"docs", "build"}, 0, "docs build\n")
}

func buildSut(out io.Writer) *cli.CLI {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return &cli.CLI{
		ColorsEnabled: false,
		Out:           out,
		CWD:           path.Join(cwd, "test-folder"),
	}
}

func assertRun(t *testing.T, args []string, expectedResultCode int, expectedOutput string) {
	b := &bytes.Buffer{}
	sut := buildSut(b)
	code := sut.Run(args)
	assert.Equal(t, expectedResultCode, code)
	assert.Equal(t, expectedOutput, b.String())
}

func handleError(err error)  {
	if err != nil {
		panic(err)
	}
}
