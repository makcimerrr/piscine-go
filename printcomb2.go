package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := 0; i <= 8; i++ {
		for j := i + 1; j <= 9; j++ {
			z01.PrintRune(rune(i + '0'))
			z01.PrintRune(rune(j + '0'))
			if i != 8 || j != 9 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
