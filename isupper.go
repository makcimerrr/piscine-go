package piscine

func IsUpper(s string) bool {
	flag := 0
	for i := 0; i < len([]rune(s)); i++ {
		if rune(s[i]) >= 'A' && rune(s[i]) <= 'Z' {
			flag = 1
		} else {
			flag = 0
		}
	}
	if flag == 1 {
		return true
	}
	return false
}
