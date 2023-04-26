package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	lenght := len(arguments)
	for i := 0; i < lenght; i++ {
		for j := i + 1; j < lenght; j++ {
			if arguments[j] < arguments[i] {
				arguments[i], arguments[j] = arguments[j], arguments[i]
			}
		}
	}
	for _, c := range arguments {
		for _, r := range c {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
