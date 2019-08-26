package utils

import (
	"bytes"
	"github.com/sirikon/tsk/src/cli"
	"gotest.tools/assert"
	"io"
	"os"
	"path"
	"testing"
)

type RunAssertion struct {
	t *testing.T
	args []string
	input string
	expectedResultCode int
	expectedOutput string
}

func (a *RunAssertion) Args(args ...string) *RunAssertion {
	a.args = args
	return a
}

func (a *RunAssertion) Input(input string) *RunAssertion {
	a.input = input
	return a
}

func (a *RunAssertion) ExpectResultCode(expectedResultCode int) *RunAssertion {
	a.expectedResultCode = expectedResultCode
	return a
}

func (a *RunAssertion) ExpectedOutput(expectedOutput string) *RunAssertion {
	a.expectedOutput = expectedOutput
	return a
}

func (a *RunAssertion) Run() {
	out := &bytes.Buffer{}
	err := &bytes.Buffer{}
	in := &bytes.Buffer{}

	in.Write([]byte(a.input))
	sut := buildSut(out, err, in)
	code := sut.Run(a.args)
	assert.Equal(a.t, a.expectedResultCode, code)
	assert.Equal(a.t, a.expectedOutput, out.String())
}

func AssertRun(t *testing.T) *RunAssertion {
	return &RunAssertion{
		t:                  t,
		args:               []string{},
		input:              "",
		expectedResultCode: 0,
		expectedOutput:     "",
	}
}

func buildSut(out io.Writer, errOut io.Writer, in io.Reader) *cli.CLI {
	cwd, err := os.Getwd()
	HandleError(err)
	return &cli.CLI{
		ColorsEnabled: false,
		Out:           out,
		Err:           errOut,
		In:            in,
		CWD:           path.Join(cwd, "test-folder"),
	}
}
