package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		// Pas d'arguments, lire depuis stdin
		bytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
			return
		}
		fmt.Print(string(bytes))
	} else {
		// Lire depuis chaque fichier donné en argument
		for _, filename := range os.Args[1:] {
			bytes, err := ioutil.ReadFile(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
				continue
			}
			fmt.Print(string(bytes))
		}
	}
}
