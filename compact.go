package piscine

func Compact(ptr *[]string) int {
	slice := *ptr
	count := 0
	lenght := len(slice)
	for i := 0; i < lenght; i++ {
		if slice[i] != "" {
			slice[count] = slice[i]
			count++
		}
	}
	*ptr = slice[:count]
	return count
}

/*const N = 6

func main() {
	a := make([]string, N)
	a[0] = "a"
	a[2] = "b"
	a[4] = "c"

	for _, v := range a {
		fmt.Println(v)
	}

	fmt.Println("Size after compacting:", Compact(&a))

	for _, v := range a {
		fmt.Println(v)
	}
}*/
