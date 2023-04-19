package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	// Vérifie que n est compris entre 1 et 9 inclus
	if n < 1 || n > 9 {
		return
	}

	// Cas spécial pour n = 1
	if n == 1 {
		// Pour chaque chiffre de 0 à 9
		for i := 0; i < 10; i++ {
			// Imprime le chiffre
			z01.PrintRune(rune(i) + '0')

			// Si ce n'est pas le dernier chiffre, imprime une virgule et un espace
			if i != 9 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}

		// Imprime un retour à la ligne à la fin
		z01.PrintRune('\n')
	} else {
		// Pour chaque chiffre de 0 à 10-n inclus
		for i := 0; i <= 9-n+1; i++ {
			// Appelle la fonction récursive avec n-1, i+1 et un préfixe vide
			PrintCombNRec(n-1, i+1, concat(string(rune(i)+'0'), ""), i == 9-n+1)
		}
	}
}

func PrintCombNRec(n, last int, prefix string, isLast bool) {
	// Cas de base: si n est égal à 0, imprime le préfixe et un retour à la ligne si c'est la dernière combinaison
	if n == 0 {
		for _, r := range prefix {
			z01.PrintRune(r)
		}
		if !isLast {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		} else {
			z01.PrintRune('\n')
		}
		return
	}

	// Pour chaque chiffre de last à 9 inclus
	for i := last; i <= 9; i++ {
		// Appelle la fonction récursive avec n-1, i+1 et le préfixe concaténé avec le chiffre courant
		PrintCombNRec(n-1, i+1, concat(prefix, string(rune(i)+'0')), isLast)
	}
}

func concat(s1 string, s2 string) string {
	// Concatène (mise bout a bout de) deux chaînes de caractères
	return s1 + s2
}
