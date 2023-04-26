package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func fHelp() {
	fmt.Printf("--insert\n  -i\n\t This upper inserts the string into the string passed as argument.\n--order\n  -o\n\t This upper will behave like a boolean, if it is called it will order the argument.\n")
} //imprime le message d'aide

func TRIEUR(table []rune) string { //trie une table de runes
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
		if length == 0 || a == "--help" || a == "-h" { //si presence de la commande help, envoie le message d'aide
			fHelp()
			return
		} else if a == "--order" || a == "-o" { //si presence de la commande order, active le upper
			upper = true
		} else if a != "" && a[:1] == "-" {
			if a[:3] == "-i=" { //si presence de la commande insert, petite ou grande, insert la chaine de caractere
				insert = a[3:] //ici apres -i= ajout de la chaine
			} else if a[:9] == "--insert=" {
				insert = a[9:] //ici apres --insert= ajout de la chaine
			}
		} else {
			if i != length-1 { //si ce n'est pas le dernier argument, ajout a la chaine secondaire et ainsi de suite
				str = str + a
			} else {
				str = str + a + insert // si c'est .e dernier caractere, ajout a la chaine de la chaine secondaire precedemment creer, et l'ajoute a la principale
			}
		}
	}

	if upper == true {
		resrune := []rune(str) //conversion chaine de caractere en une tranche runes
		str = TRIEUR(resrune)  //triache de la tranche rune
	}
	if str == "" { //si chaine de caractere vide, envoie l'affiche de l'aide
		fHelp()
		return
	}
	fmt.Printf(str) //affichage de la chaine, trier si besoin
	z01.PrintRune('\n')
}
