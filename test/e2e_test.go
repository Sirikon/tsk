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

func TestCLI_Index(t *testing.T) {
	b := &bytes.Buffer{}
	sut := buildSut(b)

	sut.Index()

	assert.Equal(t, `
  tsk Testing
  Usage: tsk [command] <subcommands...>

    build scripts/build.sh
    docs scripts/docs
      build scripts/docs/build.sh
`, b.String())
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
