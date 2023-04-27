package main

func Split(s, sep string) []string {

	var decodage []string
	chaine := len(s)
	separateur := len(sep)
	i := 0
	for j := 0; j < chaine; j++ {
		if j+separateur > chaine || s[j:j+separateur] != sep {
			continue
		}

		decodage = append(decodage, s[i:j])

		i = j + separateur
	}

	decodage = append(decodage, s[i:])

	return decodage
}
