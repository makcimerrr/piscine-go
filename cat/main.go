package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 1 {
		// Pas d'arguments, lire depuis stdin
		bytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			printError(err.Error())
			return
		}
		printString(string(bytes))
	} else {
		// Lire depuis chaque fichier donné en argument
		for _, filename := range os.Args[1:] {
			bytes, err := ioutil.ReadFile(filename)
			if err != nil {
				printError(err.Error())
				continue
			}
			printString(string(bytes))
		}
	}
}

// Écrit une chaîne de caractères sur la sortie standard.
func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

// Écrit un message d'erreur sur la sortie standard d'erreur.
func printError(msg string) {
	for _, r := range msg {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
