package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		// Pas d'arguments, lire depuis stdin
		for {
			var text string
			_, err := fmt.Scanln(&text)
			if err != nil {
				break
			}
			fmt.Println(text)
		}
	} else {
		// Lire depuis chaque fichier donné en argument
		for _, filename := range os.Args[1:] {
			file, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
				continue
			}
			defer file.Close()
			var text string
			for {
				_, err := fmt.Fscanln(file, &text)
				if err != nil {
					break
				}
				fmt.Println(text)
			}
		}
	}
}
