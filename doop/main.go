package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	arg1, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		return
	}

	arg2 := os.Args[2]

	arg3, err := strconv.ParseInt(os.Args[3], 10, 64)
	if err != nil {
		return
	}

	var result int64

	switch arg2 {
	case "+":
		result = arg1 + arg3
	case "-":
		result = arg1 - arg3
	case "*":
		result = arg1 * arg3
	case "/":
		if arg3 == 0 {
			fmt.Println("No division by 0")
			return
		}
		result = arg1 / arg3
	case "%":
		if arg3 == 0 {
			fmt.Println("No modulo by 0")
			return
		}
		result = arg1 % arg3
	default:
		return
	}

	if result > 9223372036854775807 || result < -9223372036854775808 {
		return
	}

	fmt.Println(result)
}
