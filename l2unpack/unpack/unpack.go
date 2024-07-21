package unpack

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

func Unpack(str string) string {
	var acc string
	var escape bool
	for _, c := range str {
		if string(c) == `\` && escape == false {
			escape = true
			continue
		} else if escape == true || !unicode.IsNumber(c) {
			acc += string(c)
		} else {
			last, _ := utf8.DecodeLastRuneInString(acc)
			if last == utf8.RuneError {
				continue
			}
			acc = acc + strings.Repeat(string(last), int(c-'0')-1)
		}
		escape = false
	}

	return acc
}
