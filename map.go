package piscine

func Map(f func(int) bool, a []int) []bool {
	lenght := len(a)
	j := make([]bool, lenght)
	for i := 0; i < lenght; i++ {
		j[i] = f(a[i])
	}
	return j
}
