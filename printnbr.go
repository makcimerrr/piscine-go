package piscine

import "github.com/01-edu/z01"

func printNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	if n >= 10 {
		piscine.printNbr(n / 10)
	}
	z01.PrintRune(rune(n%10 + '0'))
}
