package piscine

func IsUpper(s string) bool {
	for i := 0; i < len([]rune(s)); i++ {
		if rune(s[i]) < 'a' || rune(s[i]) > 'z' {
			return false
		}
	}
	return true
}
