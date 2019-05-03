#!/bin/sh

go build -ldflags "-s -w" -o ./out/tsk ./cmd/tsk
ls -lh ./out
