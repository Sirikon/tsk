package utils

// https://github.com/willf/pad/blob/3bc964e5ac6d49201d42be066f5480a2d7401cdc/utf8/pad.go

import "unicode/utf8"

func times(str string, n int) (out string) {
	for i := 0; i < n; i++ {
		out += str
	}
	return
}

// PadLeft .
func PadLeft(str string, len int, pad string) string {
	return times(pad, len-utf8.RuneCountInString(str)) + str
}
