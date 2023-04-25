package piscine

func LastRune(s string) rune {
	nb := (len(s)) - 1
	return rune([]rune(s)[nb])
}
