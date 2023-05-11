package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) > 1 {
		arg := os.Args[1]
		for i := 0; i < len(os.Args); i++ {
			z01.PrintRune(rune(arg[i]))
		}
		z01.PrintRune('\n')
	}
}
