package piscine

func Capitalize(s string) string {
	c := []rune(s)
	flag := 0
	for i := 0; i < len(c); i++ {
		if flag == 0 {
			if (c[i] >= 'A' && c[i] <= 'Z') || (c[i] >= '0' && c[i] <= '9') {
				flag = 1
			} else if c[i] >= 'a' && c[i] <= 'z' {
				c[i] = c[i] - 32
				flag = 1
			}
		} else if c[i] >= 'A' && c[i] <= 'Z' && flag == 1 {
			c[i] = c[i] + 32
		} else if !(c[i] >= 'a' && c[i] <= 'z') && !(c[i] >= '0' && c[i] <= '9') {
			flag = 0
		}
	}
	return string(c)
}
