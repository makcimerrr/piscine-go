package piscine

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
	resolution(tableau, 0, []int{})
}

func resolution(tableau [][]bool, ligne int, colonne []int) {
	// Si toutes les reines sont placées, imprimer la solution et terminer la récursion
	if ligne == 8 {
		printBoard(colonne)
		return
	}

	// Essayer de placer une reine dans chaque colonne de la ligne actuelle
	for col := 0; col < 8; col++ {
		// Vérifier si la position est valide (pas d'autres reines menaçant cette position)
		if isValid(colonne, ligne, col) {
			// Si c'est le cas, ajouter cette colonne à la liste des colonnes et marquer la position sur le plateau
			colonne = append(colonne, col)
			tableau[ligne][col] = true

			// Appeler récursivement la fonction pour placer les reines restantes sur les lignes suivantes
			resolution(tableau, ligne+1, colonne)

			// Retirer la colonne de la liste et démarquer la position pour revenir en arrière dans la récursion
			colonne = colonne[:len(colonne)-1]
			tableau[ligne][col] = false
		}
	}
}

func isValid(colonne []int, ligne int, col int) bool {
	for r, c := range colonne {
		// Vérifier si la reine sur la ligne r et colonne c menace la position de la reine qu'on essaye de placer
		if c == col || r-c == ligne-col || r+c == ligne+col {
			return false
		}
	}
	return true
}

func printBoard(colonne []int) {
	for _, col := range colonne {
		for i := 0; i < 8; i++ {
			if i == col {
				z01.PrintRune(rune(i + 49)) // convertir le chiffre en caractère ASCII correspondant
			}
		}
	}
	z01.PrintRune('\n')
}
