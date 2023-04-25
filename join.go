package main

import (
	"fmt"
)

func Join(strs []string, sep string) string {
	step := ""
	count := 0
	for i := range strs {
		count = count + 1
		i = i + 1
	}
	for j := range strs {
		step = step + strs[j]
		if j < count-1 {
			step = step + sep
		}
	}
	return step
}

func main() {
	toConcat := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(Join(toConcat, ":"))
}
