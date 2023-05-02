package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1:]
	lenght := len(arg)
	flag := false
	flag2 := false
	flag3 := true
	if len(arg) > 2 {
		index := BasicAtoi(arg[1])
		for i := 2; i < lenght; i++ {
			data, err := os.ReadFile(arg[i])
			if err != nil {
				fmt.Printf(err.Error())
				fmt.Printf("\n")
				flag = true
				flag2 = false
			} else if index <= len(data)+1 {
				if flag2 == false && flag3 == true && i != 2 {
					fmt.Printf("\n")
				}
				fmt.Printf("==> " + string(arg[i]) + " <==")
				fmt.Printf("\n")
				fmt.Printf(string(data[len(data)-index:]))
				flag3 = true
			} else if index > len(data) {
				if flag2 == false && flag3 == true {
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
	if flag == true {
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
