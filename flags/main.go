package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func fHelp() {
	fmt.Printf("--insert\n  -i\n\t This flag inserts the string into the string passed as argument.\n--order\n  -o\n\t This flag will behave like a boolean, if it is called it will order the argument.\n")
}

func TRIEUR(table []rune) string {
	for i := 0; i < len(table); i++ {
		for j := i + 1; j < len(table); j++ {
			if table[i] > table[j] {
				table[i], table[j] = table[j], table[i]
			}
		}
	}
	return string(table)
}

func main() {
	args := os.Args[1:]
	length := len(args)
	upper := false
	insert := ""
	str := ""
	for i, a := range args {
		if length == 0 || a == "--help" || a == "-h" {
			fHelp()
			return
		} else if a == "--order" || a == "-o" {
			upper = true
		} else if a != "" && a[:1] == "-" {
			if a[:3] == "-i=" {
				insert = a[3:]
			} else if a[:9] == "--insert=" {
				insert = a[9:]
			}
		} else {
			if i != length-1 {
				str = str + a
			} else {
				str = str + a + insert
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
