package piscine

func ToLower(s string) string {
	c := []rune(s)
	for i := 0; i < len([]rune(s)); i++ {
		if c[i] >= 'A' && c[i] <= 'Z' {
			c[i] = c[i] + 32
		}
	}
	return string(c)
}
