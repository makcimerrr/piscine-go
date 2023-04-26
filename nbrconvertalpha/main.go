package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	upper := false
	if len(args) > 0 && args[0] == "--upper" { // si il y a la condition upper mettre le flag a true
		upper = true
		args = args[1:] // et prends en compte l'argument suivant pour ignorer upper
	}
	if len(args) == 0 { // si aucune demande renvoie rien
		return
	}
	for _, arg := range args {
		n := 0
		for _, r := range arg { // adaptation pour les lettres de plus de 9
			if r > '9' {
				break
			}
			digit := int(r - '0')
			n = n*10 + digit // conversion en
		}
		if n < 1 || n > 26 {
			z01.PrintRune(' ') // si caractere non valide renvoie un espace (Nv = hors alphabet)
		} else if upper {
			z01.PrintRune(rune('A' + n - 1))
		} else {
			z01.PrintRune(rune('a' + n - 1))
		}
	}
	z01.PrintRune('\n')
}
