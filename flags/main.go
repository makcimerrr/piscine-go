package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func fHelp() {
	fmt.Printf("--insert\n  -i\n\t This flag inserts the string into the string passed as argument.\n--order\n  -o\n\t This flag will behave like a boolean, if it is called it will order the argument.\n")
}
func fOrder(table []rune) string {
	for i := 0; i < len(table); i++ {
		for j := i + 1; j < len(table); j++ {
			if table[i] > table[j] {
				table[i], table[j] = table[j], table[i]
			}
		}
	}
	return string(table)
}
func fInsert() {
}
func main() {
	toto := os.Args[1:]
	lentoto := len(toto)
	flago := false
	valins := ""
	resfinal := ""
	for _, a := range toto {
		if lentoto == 0 || a == "--help" || a == "-h" { // si on trouve un help = msg d'aide + quit
			fHelp()
			return
		} else if a == "--order" || a == "-o" { // test si présence d'un order (ou +)
			flago = true
		} else if a != "" && a[:1] == "-" {
			if a[:3] == "-i=" { // ajout valeur après -i=
				valins = valins + a[3:]
			} else if a[:9] == "--insert=" {
				valins = valins + a[9:] // ajout valeur après --insert=
			}
		} else {
			resfinal = resfinal + a // on regroupe les args
		}
	}
	resfinal = resfinal + valins // ajout valeur insert, si aucun insert resfinal = resfinal + 0
	if flago == true {           // sort des valeurs
		resrune := []rune(resfinal)
		resfinal = fOrder(resrune)
	}
	if resfinal == "" {
		fHelp()
		return
	}
	fmt.Printf(resfinal)
	z01.PrintRune('\n')
}
