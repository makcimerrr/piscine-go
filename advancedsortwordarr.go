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

/*func main() {
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	AdvancedSortWordArr(result, Compare)

	fmt.Println(result)
}

func Compare(a, b string) int {
	if a < b {
		return -1
	}
	if a == b {
		return 0
	}
	if a > b {
		return 1
	}
	return 0
}*/
