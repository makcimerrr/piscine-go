package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 1 {
		// Pas d'arguments, lire depuis stdin
		bytes, _ := ioutil.ReadAll(os.Stdin)
		for _, b := range bytes {
			z01.PrintRune(rune(b))
		}
	} else {
		// Lire depuis chaque fichier donné en argument
		for _, filename := range os.Args[1:] {
			bytes, err := ioutil.ReadFile(filename)
			if err != nil {
				printError("open " + filename + ": no such file or directory")
				os.Exit(1)
			}
			for _, b := range bytes {
				z01.PrintRune(rune(b))
			}
		}
	}
}

func printError(msg string) {
	for _, c := range "ERROR: " + msg {
		z01.PrintRune(rune(c))
	}
	z01.PrintRune('\n')
}
