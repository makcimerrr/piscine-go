package piscine

func NRune(s string, n int) rune {
	if n <= 0 {
		return 0
	}
	if len([]rune(s)) < n {
		return 0
	} else {
		return rune([]rune(s)[n-1])
	}
}
