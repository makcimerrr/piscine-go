package piscine

func Split(s, sep string) []string {

	var decodage []string
	str := len(s)
	separateur := len(sep)
	i := 0
	for j := 0; j < chaines; j++ {
		if j+separateur > chaines || s[j:j+separateur] != sep {
			continue
		}

		decodage = append(decodage, s[i:j])

		i = j + separateur
	}

	decodage = append(decodage, s[i:])

	return decodage
}
