package piscine

import (
	"github.com/01-edu/z01"
)

func ForEach(f func(int), a []int) {
	for _, c := range a {
		f(c)
	}
	z01.PrintRune('\n')
}

func PrintNbr(n int) {
	if n == -9223372036854775808 {
		PrintNbr(-9223372036854775807 / 10)
		z01.PrintRune('8')
		return
	}

	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	if n >= 10 {
		PrintNbr(n / 10)
	}
	z01.PrintRune(rune(n%10 + '0'))
}
