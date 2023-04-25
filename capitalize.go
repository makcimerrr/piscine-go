package piscine

import (
	"unicode"
)

func Capitalize(s string) string {
	var res string
	for i := 0; i < len(s); i++ {
		if i == 0 || !unicode.IsLetter(rune(s[i-1])) {
			res += string(unicode.ToUpper(rune(s[i])))
		} else {
			res += string(unicode.ToLower(rune(s[i])))
		}
	}
	return res
}
