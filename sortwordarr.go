package piscine

func SortWordArr(a []string) {
	lenght := len(a)
	for i := 1; i < lenght; {
		if a[i-1] > a[i] {
			lenght2 := a[i]
			a[i] = a[i-1]
			a[i-1] = lenght2
			i = 1
		} else {
			i++
		}
	}
}
