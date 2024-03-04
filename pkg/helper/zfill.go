package helper

import "strings"

func Zfill(s string, filler string, maxDigit int) string {
	/*
		Zfill("5432", "0", 6)
		output: 005432
	*/
	l := maxDigit - len(s)
	return strings.Repeat(filler, l) + s
}
