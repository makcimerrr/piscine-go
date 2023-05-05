package piscine

func Unmatch(a []int) int {
	liste := make(map[int]int)

	for _, c := range a {
		liste[c]++
	}

	for num, v := range liste {
		if v%2 != 0 {
			return num
		}
	}

	return -1
}

/*func main() {
	a := []int{1, 2, 3, 1, 2, 3, 4}
	unmatch := Unmatch(a)
	fmt.Println(unmatch)
}*/
