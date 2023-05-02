package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1:]
	lenarg := len(arg)
	erreur := false
	sauterr := false
	saut := true
	if len(arg) > 2 {
		index := BasicAtoi(arg[1])
		for i := 2; i < lenarg; i++ {
			data, err := os.ReadFile(arg[i])
			if err != nil {
				fmt.Printf(err.Error())
				fmt.Printf("\n")
				erreur = true
				sauterr = false
			} else if index <= len(data)+1 {
				if sauterr == false && saut == true && i != 2 {
					fmt.Printf("\n")
				}
				fmt.Printf("==> " + string(arg[i]) + " <==")
				fmt.Printf("\n")
				fmt.Printf(string(data[len(data)-index:]))
				saut = true
			} else if index > len(data) {
				if sauterr == false && saut == true {
					fmt.Printf("\n")
				}
				fmt.Printf("==> " + string(arg[i]) + " <==")
				fmt.Printf("\n")
				fmt.Printf(string(data))
			}
		}
	} else {
		os.Exit(1)
	}
	if erreur == true {
		os.Exit(1)
	}
}

func BasicAtoi(s string) int {
	a := len(s)
	b := []rune(s)
	var ret int
	for i := 0; i < a; i++ {
		ret = ret*10 + int((b[i])-'0')
	}
	return ret
}
