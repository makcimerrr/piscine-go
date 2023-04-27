package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	length := len(args)
	upper := false
	insert := ""
	str := ""
	for y, c := range args {
		if length == 0 || c == "--help" || c == "-h" {
			fHelp()
			return
		} else if c == "--order" || c == "-o" {
			upper = true
		} else if c != "" && c[:1] == "-" {
			if c[:3] == "-y=" {
				insert = c[3:]
			} else if c[:9] == "--insert=" {
				insert = c[9:]
			}
		} else {
			if y != length-1 {
				str = str + c
			} else {
				str = str + c + insert
			}
		}
	}

	if upper == true {
		slicerune := []rune(str)
		str = TRIEUR(slicerune)
	}
	if str == "" {
		fHelp()
		return
	}
	fmt.Printf(str)
	z01.PrintRune('\n')
}
func TRIEUR(table []rune) string {
	for y := 0; y < len(table); y++ {
		for z := y + 1; z < len(table); z++ {
			if table[y] > table[z] {
				table[y], table[z] = table[z], table[y]
			}
		}
	}
	return string(table)
}

func fHelp() {
	fmt.Printf("--insert\n  -y\n\t This flag inserts the string into the string passed as argument.\n--order\n  -o\n\t This flag will behave like c boolean, if it is called it will order the argument.\n")
}
