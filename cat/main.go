package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	lenarg := len(arg)
	if lenarg > 0 {
		for i := 0; i < lenarg; i++ {
			data, err := ioutil.ReadFile(arg[i])
			if err != nil {
				errorp := "ERROR: open " + string(arg[i]) + ": no such file or directory"
				PrintStr(errorp)
				z01.PrintRune('\n')
				os.Exit(1)
			}
			PrintStr(string(data))
		}
	} else {
		io.Copy(os.Stdout, os.Stdin)
	}
}

func PrintStr(s string) {
	a := len(s)
	for i := 0; i < a; i++ {
		z01.PrintRune(rune(s[i]))
	}
}
