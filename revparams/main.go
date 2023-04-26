package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := os.Args
	lenght := len(argument)
	for i := lenght - 1; i >= 0; i-- {
		if i != 0 {
			for j := range argument[i] {
				lenght := []rune(argument[i])
				z01.PrintRune(lenght[j])
			}
			z01.PrintRune('\n')
		}
	}
}
