package piscine

func Unmatch(a []int) int {
	liste := make(map[int]int)

	for _, count := range a {
		liste[count]++
	}

	for num, count := range liste {
		if count%2 != 0 {
			return num
		}
	}

	return -1
}

/*func main() {
	a := []int{1, 2, 1, 1, 4, 5, 5, 4, 1, 7}
	unmatch := Unmatch(a)
	fmt.Println(unmatch)
}*/
