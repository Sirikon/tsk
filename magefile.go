// +build mage

package main

import (
	"bytes"
	"fmt"
	"github.com/magefile/mage/sh"
	"strings"
)

func Build() {
	err := sh.RunV("go", "build", "-ldflags", "-s -w", "-o", "./out/tsk", "./cmd/tsk"); check(err)
	err = sh.RunV("ls", "-lh", "./out"); check(err)
}

func Test() {
	err := sh.RunV("go", "test", "./test")
	check(err)
}

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

	err := sh.RunV("goreleaser", goreleaserArgs...); check(err)
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

func check(err error) {
	if err != nil {
		panic(err)
	}
}
