package piscine

func SplitWhiteSpaces(s string) []string {
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
