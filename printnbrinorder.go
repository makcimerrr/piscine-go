package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}

	// Calcul du nombre de chiffres dans l'entier
	numDigits := 1
	for temp := n / 10; temp != 0; temp /= 10 {
		numDigits++
	}

	// Création d'un tableau pour stocker les chiffres
	digits := make([]int, numDigits)

	// Stockage des chiffres dans le tableau
	for i := numDigits - 1; i >= 0; i-- {
		digits[i] = n % 10
		n /= 10
	}

	// Tri du tableau : SWAP des valeurs
	for i := 0; i < numDigits-1; i++ {
		for j := i + 1; j < numDigits; j++ {
			if digits[i] > digits[j] {
				temp := digits[i]
				digits[i] = digits[j]
				digits[j] = temp
			}
		}
	}

	// Impression des chiffres
	for _, digit := range digits {
		z01.PrintRune(rune(digit + '0'))
	}
}
