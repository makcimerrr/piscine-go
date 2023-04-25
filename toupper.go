package piscine

func ToUpper(s string) string {
	c := []rune(s)
	for i := 0; i < len([]rune(s)); i++ {
		if c[i] >= 'a' && c[i] <= 'z' {
			c[i] = c[i] - 32
		}
	}
	return string(c)
}
