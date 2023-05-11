package main

import "fmt"

func Chunk(slice []int, size int) {
	if size > 0 && size < len(slice) {
		var chunk [][]int

		for size < len(slice) {
			slice, chunk = slice[size:], append(chunk, slice[:size])
		}

		fmt.Print(append(chunk, slice))
	}
	if len(slice) == 0 {
		fmt.Print(slice)
	}
	fmt.Println()
}

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}
