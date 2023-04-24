package main

import (
	"github.com/01-edu/z01"
)

func EightQueens() {
	// Créer un tableau de 8x8 pour stocker le plateau de jeu
	board := make([][]bool, 8)
	for i := range board {
		board[i] = make([]bool, 8)
	}

	// Appeler la fonction recursive qui résoudra le puzzle
	solve(board, 0, []int{})
}

func solve(board [][]bool, row int, columns []int) {
	// Si toutes les reines sont placées, imprimer la solution et terminer la récursion
	if row == 8 {
		printBoard(columns)
		return
	}

	// Essayer de placer une reine dans chaque colonne de la ligne actuelle
	for col := 0; col < 8; col++ {
		// Vérifier si la position est valide (pas d'autres reines menaçant cette position)
		if isValid(columns, row, col) {
			// Si c'est le cas, ajouter cette colonne à la liste des colonnes et marquer la position sur le plateau
			columns = append(columns, col)
			board[row][col] = true

			// Appeler récursivement la fonction pour placer les reines restantes sur les lignes suivantes
			solve(board, row+1, columns)

			// Retirer la colonne de la liste et démarquer la position pour revenir en arrière dans la récursion
			columns = columns[:len(columns)-1]
			board[row][col] = false
		}
	}
}

func isValid(columns []int, row int, col int) bool {
	for r, c := range columns {
		// Vérifier si la reine sur la ligne r et colonne c menace la position de la reine qu'on essaye de placer
		if c == col || r-c == row-col || r+c == row+col {
			return false
		}
	}
	return true
}

func printBoard(columns []int) {
	for _, col := range columns {
		for i := 0; i < 8; i++ {
			if i == col {
				z01.PrintRune(rune(i + 49)) // convertir le chiffre en caractère ASCII correspondant
			} else {
				z01.PrintRune('.')
			}
		}
	}
	z01.PrintRune('\n')
}

func main() {
	EightQueens()
}
