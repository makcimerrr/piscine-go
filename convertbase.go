package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	// Conversion de nbr en base 10
	lenght := len(baseFrom)
	s := make(map[byte]int)
	for i := 0; i < lenght; i++ {
		s[baseFrom[i]] = i
	}

	num := 0
	for i := 0; i < len(nbr); i++ {
		num = num*lenght + s[nbr[i]]
	}

	// Conversion de num en base demande (baseTo)
	c := []byte(baseTo)
	lenght2 := len(c)
	result := ""
	for num > 0 {
		digit := num % lenght2
		result = string(c[digit]) + result
		num = num / lenght2
	}

	return result
}
