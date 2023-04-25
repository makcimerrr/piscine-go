package piscine

func IsUpper(s string) bool {
	for i := 0; i < len([]rune(s)); i++ {
		if rune(s[i]) < 'A' || rune(s[i]) > 'Z' {
			return false
		}
	}
	return true
}
