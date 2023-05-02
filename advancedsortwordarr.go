package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	lenght := len(a)
	for i := 0; i < lenght; i++ {
		for j := i + 1; j < lenght; j++ {
			if f(a[i], a[j]) == 1 {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}
