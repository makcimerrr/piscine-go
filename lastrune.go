package piscine

func LastRune(s string) rune {
	lenght := (len(s)) - 1
	return rune([]rune(s)[lenght])
}
