package targets

import (
	"bytes"
	"fmt"
	"github.com/magefile/mage/sh"
	"github.com/sirikon/tsk/scripts/targets/utils"
	"strings"
)

// Release Creates a new release in `dist` folder using goreleaser, and
// publishes to GitHub if finds a valid version tag.
func Release() {
	fmt.Println("Release process started")
	currentTag := getCurrentTag()

	goreleaserArgs := []string{"--rm-dist"}

	if currentTag == "" {
		fmt.Println("Couldn't find any git tag. Skipping goreleaser publication.")
		goreleaserArgs = append(goreleaserArgs, "--snapshot", "--skip-publish")
	} else {
		fmt.Println("Found git tag:", currentTag)
	}

	utils.Check(sh.RunV("goreleaser", goreleaserArgs...))
}

func getCurrentTag() string {
	stdOut := &bytes.Buffer{}
	stdErr := &bytes.Buffer{}
	ran, err := sh.Exec(nil, stdOut, stdErr, "git", "describe", "--tags", "--exact-match")

	if !ran {
		panic("Could not run git command: git describe --tags --exact-match")
	}

	if err != nil {
		return ""
	}

	return strings.TrimSuffix(stdOut.String(), "\n")
}
