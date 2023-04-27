package piscine

import (
	"fmt"
)

func SplitWhiteSpaces(s string) []string {
	mot := []string{}
	chaine := ""
	flag := false

	for _, c := range s {
		if c == ' ' || c == '\t' || c == '\n' {
			if flag {

				mot = append(mot, chaine)
				chaine = ""
			}
			flag = false
		} else {
			chaine += string(c)
			flag = true
		}
	}

	if flag {
		mot = append(mot, chaine)
	}

	return mot
}

func main() {
	fmt.Printf("%#v\n", SplitWhiteSpaces("Hello how are you?"))
}
