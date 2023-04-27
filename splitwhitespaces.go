package main

func SplitWhiteSpaces(s string) []string {
	chaine := make([]string, 0)
	mot := ""
	flag := false

	for _, char := range s {
		if char == ' ' || char == '\t' || char == '\n' {
			if flag {

				chaine = append(chaine, mot)
				mot = ""
			}
			flag = false
		} else {
			mot += string(char)
			flag = true
		}
	}

	if flag {
		chaine = append(chaine, mot)
	}

	return chaine
}
