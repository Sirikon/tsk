#!/usr/bin/env bash

curl -sL https://git.io/goreleaser | bash -s -- --snapshot --skip-publish --rm-dist
