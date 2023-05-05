package piscine

func Unmatch(a []int) int {
	for _, c := range a {
		count := 0
		for _, v := range a {
			if c == v {
				count = count + 1
			}
		}
		if count%2 != 0 {
			return c
		}
	}
	return -1
}

/*func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	unmatch := Unmatch(a)
	fmt.Println(unmatch)
}*/
