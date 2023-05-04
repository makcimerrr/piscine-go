package piscine

func Abort(a, b, c, d, e int) int {
	var tab []int
	tab = append(tab, a, b, c, d, e)
	SortIntegerTable(tab)
	res := tab[2]
	return res
}

func SortIntegerTable(table []int) {
	a := len(table)
	for i := 1; i < a; {
		if table[i-1] > table[i] {
			a := table[i]
			table[i] = table[i-1]
			table[i-1] = a
			i = 1
		} else {
			i++
		}
	}
}

/*func main() {
	middle := Abort(2, 3, 8, 5, 7)
	fmt.Println(middle)
}*/
