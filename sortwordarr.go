package piscine

func SortWordArr(a []string) {
	lenght := len(a)
	for i := 0; i < lenght; i++ {
		for j := i + 1; j < lenght; j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}
