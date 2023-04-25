package piscine

func LastRune(s string) rune {
	number := (len(s)) - 1
	return rune([]rune(s)[number])
}
