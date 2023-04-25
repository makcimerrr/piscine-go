package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if len(base) < 2 { //A base must contain at least 2 characters. (condition 1)
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	for i := 0; i < len(base); i++ {
		if base[i] == '+' || base[i] == '-' { //A base should not contain + or - characters.(condition 3)
			z01.PrintRune('N')
			z01.PrintRune('V')
			return
		}
		for j := i + 1; j < len(base); j++ { //Each character of a base must be unique. (condition 2)
			if base[i] == base[j] {
				/*Si i et j sont sont identiques,
				il y a des caractères en double dans "base" donc invalide.*/
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			}
		}
	}
	var result string
	negative := nbr < 0
	if negative {
		nbr = -nbr //convertit nbr en positif
	}
	for nbr > 0 {
		digit := nbr % len(base)              //correspond au chiffre courant dans la nouvelle base
		result = string(base[digit]) + result //construis la chaine en ajoutant caractère par caractère
		nbr = nbr / len(base)                 //determine le prochain chiffre a convertir.
	}
	if negative {
		result = "-" + result //si negatif affiche le signe "-" devant.
	}

	for _, r := range result { //parcourt chaque caractère de la variable "result" contenant la conversion, et affiche chaque caractère
		z01.PrintRune(rune(r)) //affiche le résultat de la conversion de l'entier dans la base, caractère par caractère
	
}
