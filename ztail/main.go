package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	lenght := len(os.Args)
	if lenght < 4 {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}

	numero, fichier := numberOfBytes(os.Args[1:])

	printName := len(fichier) > 1

	for j, f := range fichier {
		fi, err := os.Open(f)
		if err != nil {
			fmt.Printf("tail: cannot open '%s' for reading: No such file or directory\n", f)
			os.Exit(1)
		}
		if printName {
			fmt.Printf("==> %s <==\n", f)
		}
		table := make([]byte, int(numero))
		_, er := fi.ReadAt(table, TailleFichier(fi)-int64(numero))
		if er != nil {
			fmt.Println(er.Error())
		}

		for _, c := range table {
			z01.PrintRune(rune(c))
		}

		if j < len(fichier)-1 {
			z01.PrintRune('\n')
		}

		fi.Close()
	}
}

func numberOfBytes(args []string) (int, []string) {
	lenght := len(args)
	numero := 0
	var fichier []string
	for i, v := range args {
		var err error
		_, err = strconv.Atoi(v)
		if v == "-c" {
			if i >= lenght-1 {
				fmt.Printf("tail: option requires an argument -- 'c'\nTry 'tail --help' for more information.")
				os.Exit(1)
			}
			arg := args[i+1]

			numero, err = strconv.Atoi(arg)

			if err != nil {
				fmt.Printf("tail: invalid number of bytes: %s\n", arg)
				os.Exit(1)
			}
			continue
		}

		if err != nil {
			fichier = append(fichier, v)
		}

	}
	return numero, fichier
}

func TailleFichier(fi *os.File) int64 {
	fil, err := fi.Stat()
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}

	return fil.Size()
}
