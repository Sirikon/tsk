package cli

import (
	"fmt"
	"io"
)

func HandlePanic(result *int, out io.Writer) {
	if r := recover(); r != nil {
		_, _ = fmt.Fprintln(out, r.(error))
		*result = 1
	}
}

func HandleErr(err error)  {
	if err != nil {
		panic(err)
	}
}
