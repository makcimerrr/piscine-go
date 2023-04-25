package piscine

func IsAlpha(s string) bool {
	for i := 0; i < len([]rune(s)); i++ {
		if (rune(s[i]) < 'a' || rune(s[i]) > 'z') && (rune(s[i]) < 'A' || rune(s[i]) > 'Z') && (rune(s[i]) < '0' || rune(s[i]) > '9') {
			return false
		}
	}
	return true
}
