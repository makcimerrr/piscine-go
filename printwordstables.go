package piscine

import "github.com/01-edu/z01"

func SplitWhiteSpace(s string) []string {
	chaine := make([]string, 0)
	mot := ""
	flag := false

	for _, c := range s {
		if c == ' ' || c == '\t' || c == '\n' {
			if flag == true {

				chaine = append(chaine, mot)
				mot = ""
			}
			flag = false
		} else {
			mot += string(c)
			flag = true
		}
	}

	if flag == true {
		chaine = append(chaine, mot)
	}

	return chaine
}

func PrintWordsTables(a []string) {
	for _, y := range a {
		for _, z := range y {
			z01.PrintRune(z)
		}
		z01.PrintRune('\n')
	}
}
