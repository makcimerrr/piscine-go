package main

import "fmt"

func ActiveBits(n int) int {
	count := 0
	for n != 0 {
		if n%2 == 1 {
			count = count + 1
		}
		n = n / 2
	}
	return count
}

func main() {
	fmt.Println(ActiveBits(7))
}
