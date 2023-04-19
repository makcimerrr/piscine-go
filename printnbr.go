package piscine

import "github.com/01-edu/z01"

func printNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	for n >= 10 {
		digit := n % 10
		n /= 10
		z01.PrintRune(rune(digit + '0'))
	}
	z01.PrintRune(rune(n + '0'))
}
