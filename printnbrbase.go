package piscine

import (
	"fmt"
)

func PrintNbrBase(nbr int, base string) {
	if len(base) < 2 { //A base must contain at least 2 characters. (condition 1)
		fmt.Println("NV")
		return
	}
	for i, c := range base { //A base should not contain + or - characters.(condition 3)
		if c == '-' || c == '+' {
			fmt.Println("NV")
			return
		}
		for j, s := range base {
			if i != j && c == s { //Each character of a base must be unique. (condition 2)

				/*Si i et j sont différents et c et c2 sont identiques,
				il y a des caractères en double dans "base" donc invalide.*/

				fmt.Println("NV")
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
		result = string(base[digit]) + result //construis la chaine en ajoutant caracteres par caracteres.
		nbr = nbr / len(base)                 //determine le prochain chiffre a convertir.
	}
	if negative {
		result = "-" + result //si negatif affiche le signe "-" devant.
	}

	fmt.Println(result)
}
