package piscine

import "fmt"

func Sqrt(nb int) int {
	for i := 0; i <= nb; i++ {
		if i*i == nb {
			return i
		}
	}
	return 0
}

/*func main() {
	fmt.Println(Sqrt(1))
	fmt.Println(Sqrt(3))
	fmt.Println(Sqrt(4))
}
