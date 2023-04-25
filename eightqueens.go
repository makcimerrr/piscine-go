package main

import (
	"github.com/01-edu/z01"
)

func EightQueens() {
	// Créer un tableau de 8x8 pour stocker le plateau de jeu
	tableau := make([][]bool, 8)
	for i := range tableau {
		tableau[i] = make([]bool, 8)
	}
	// Appeler la fonction recursive qui résoudra le puzzle
	solve(tableau, 0, []int{})
}

func solve(tableau [][]bool, lignes int, colonnes []int) {
	// Si toutes les reines sont placées, imprimer la solution et terminer la récursion
	if lignes == 8 {
		printBoard(colonnes)
		return
	}
	// Essayer de placer une reine dans chaque colonne de la ligne actuelle
	for colonne := 0; colonne < 8; colonne++ {
		// Vérifier si la position est valide (pas d'autres reines menaçant cette position)
		if isValid(colonnes, lignes, colonne) {
			// Si c'est le cas, ajouter cette colonne à la liste des colonnes et marquer la position sur le plateau
			colonnes = append(colonnes, colonne)
			tableau[lignes][colonne] = true
			// Appeler récursivement la fonction pour placer les reines restantes sur les lignes suivantes
			solve(tableau, lignes+1, colonnes)
			// Retirer la colonne de la liste et démarquer la position pour revenir en arrière dans la récursion
			colonnes = colonnes[:len(colonnes)-1]
			tableau[lignes][colonne] = false
		}
	}
}

func isValid(colonnes []int, lignes int, colonne int) bool {
	for r, c := range colonnes {
		// Vérifier si la reine sur la ligne r et colonne c menace la position de la reine qu'on essaye de placer
		if c == colonne || r-c == lignes-colonne || r+c == lignes+colonne {
			return false
		}
	}
	return true
}

func printBoard(colonnes []int) {
	for _, colonne := range colonnes {
		for i := 0; i < 8; i++ {
			if i == colonne {
				z01.PrintRune(rune(i + 49)) // convertir le chiffre en caractère ASCII correspondant
			}
		}
	}
	z01.PrintRune('\n')
}

func main() {
	EightQueens()
}
