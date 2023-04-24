package main

import (
	"fmt"
)

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	}

	i := nb
	if i%2 == 0 {
		i++
	}

	for {
		if izPrime(i) {
			return i
		}
		i += 2
	}
}

func izPrime(n int) bool {
	if n <= 1 {
		return false
	}

	if n == 2 {
		return true
	}

	if n%2 == 0 {
		return false
	}

	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(FindNextPrime(5))
	fmt.Println(FindNextPrime(4))
	fmt.Println(FindNextPrime(2))
}
