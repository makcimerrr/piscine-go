package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, i := range s {
		z01.PrintRune(i)
	}
}

func main() {
	file := os.Args[1:]
	if len(file) > 0 {
		for i := 0; i < len(file); i++ {
			content, err := os.ReadFile(file[i])
			if err != nil {
				msgerreur := "ERROR: open " + string(file[i]) + ": no such file or directory\n"
				printStr(msgerreur)
				os.Exit(1)
			}
			printStr(string(content))
		}
	} else {
		if _, err := io.Copy(os.Stdout, os.Stdin); err != nil {
			panic(err)
		}
	}
}
