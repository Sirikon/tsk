// +build mage

package main

import (
	"github.com/magefile/mage/sh"
)

func Build()  {
	sh.RunV("go", "build", "-ldflags", "-s -w", "-o", "./out/tsk", "./cmd/tsk")
	sh.RunV("ls", "-lh", "./out")
}

func Test()  {
	sh.RunV("go", "test", "./test")
}

func Release()  {
	currentTag := getCurrentTag()

	goreleaserArgs := []string{"--rm-dist"}

	if currentTag == "" {
		goreleaserArgs = append(goreleaserArgs, "--snapshot", "--skip-publish")
	}

	sh.RunV("goreleaser", goreleaserArgs...)
}

func getCurrentTag() string {
	tag, err := sh.Output("git", "describe", "--tags"); handleError(err)
	return tag
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
