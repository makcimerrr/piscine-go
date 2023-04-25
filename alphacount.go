package pisicne

func AlphaCount(s string) int {
	count := 0
	c := len([]rune(s))
	for i := 0; i < c; i++ {
		if rune(s[i]) >= 'c' && rune(s[i]) <= 'z' || rune(s[i]) >= 'A' && rune(s[i]) <= 'Z' {
			count = count + 1
		}
	}
	return count
}
