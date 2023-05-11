package main

import (
	"fmt"
	"os"
)

var base = map[int]rune{
	0:  '0',
	1:  '1',
	2:  '2',
	3:  '3',
	4:  '4',
	5:  '5',
	6:  '6',
	7:  '7',
	8:  '8',
	9:  '9',
	10: 'A',
	11: 'B',
	12: 'C',
	13: 'D',
	14: 'E',
	15: 'F',
}

func main() {
	n := AtoI(os.Args[1])
	b := AtoI(os.Args[2])

	arr := []rune{}

	convToBase(n, b, &arr)

	fmt.Println(string(arr))
}

func convToBase(n int, b int, arr *[]rune) {
	if n > b {
		convToBase(n/b, b, arr)
	}

	*arr = append(*arr, base[n%b])
}

func AtoI(s string) int {
	str := []rune(s)
	final := 0

	for i, c := range str {
		final += int(c - '0')

		if i != len(str)-1 {
			final *= 10
		}
	}

	return final
}
