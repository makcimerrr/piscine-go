package piscine

func Max(a []int) int {
	lenght := len(a)
	if lenght == 0 {
		return 0
	}
	max := a[0]
	for i := 1; i < lenght; i++ {
		if a[i] > max {
			max = a[i]
		}
	}
	return max
}

/*func main() {
	a := []int{23, 123, 1, 11, 55, 93}
	max := Max(a)
	fmt.Println(max)
}*/
