package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if len(base) < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	for i := 0; i < len(base); i++ {
		if base[i] == '+' || base[i] == '-' {
			z01.PrintRune('N')
			z01.PrintRune('V')
			return
		}
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			}
		}
	}
	var result string
	negative := nbr < 0
	if negative {
		nbr = -nbr
	}
	for nbr > 0 {
		digit := nbr % len(base)
		result = string(base[digit]) + result
		nbr = nbr / len(base)
	}
	if negative {
		result = "-" + result
	}

	for _, r := range result {
		z01.PrintRune(rune(r))
	}
}
