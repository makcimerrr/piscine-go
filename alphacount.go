package piscine

func AlphaCount(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		if rune(s[i]) >= 'A' && rune(s[i]) <= 'Z' || rune(s[i]) >= 'a' && rune(s[i]) <= 'z' {
			count = count + 1
		}
	}
	return count
}
