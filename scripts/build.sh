#!/usr/bin/env bash

set -e

go build -ldflags "-s -w" -o ./out/tsk ./cmd/tsk
ls -lh ./out
