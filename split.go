package piscine

func Split(s, sep string) []string {
	var decodage []string
	a := 0
	b := 0
	lenght := len(s)
	lenght2 := len(sep)
	for i := 0; i < lenght; i++ {
		b = Index(s[a:], sep)
		if b < 0 {
			break
		}
		decodage = append(decodage, s[a:a+b])
		a = a + b + lenght2
	}
	decodage = append(decodage, s[a:])
	return decodage
}
