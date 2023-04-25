package main

func LastRune(s string) rune {
	longueur := (len(s)) - 1
	return rune([]rune(s)[longueur])
}
