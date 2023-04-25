package piscine

func Capitalize(s string) string {
	c := []rune(s)
	for i := 0; i < len(c); i++ {
		if (c[i] >= 'a' && c[i] <= 'z') && (i == 0 || c[i-1] == ' ' || c[i-1] == '+') {
			c[i] = c[i] - 32
		} else if c[i] >= 'A' && c[i] <= 'Z' {
			c[i] = c[i] + 32
		}
	}
	return string(c)
}
