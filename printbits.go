package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	n := 0
	for _, c := range os.Args[1] {
		if c < '0' || c > '9' {
			fmt.Print("00000000")
			return
		}
		n = n*10 + int(c-'0')
	}
	for i := 7; i >= 0; i-- {
		bit := (n / (1 << i)) % 2
		fmt.Print(bit)
	}
}
