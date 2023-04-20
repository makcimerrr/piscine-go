package piscine

func Atoi(s string) int {
	result := 0
	sign := 1
	started := false
	for _, c := range s {
		if !started {
			if c == '-' {
				sign = -1
				started = true
				continue
			} else if c == '+' {
				started = true
				continue
			}
		}
		if c < '0' || c > '9' {
			return 0
		}
		digit := int(c - '0')
		result = result*10 + digit
		started = true
	}
	return result * sign
}
