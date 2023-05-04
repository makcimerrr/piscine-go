package main

import "fmt"

const AIRCRAFT1 = 1

type Pilot struct {
	Name     string
	Life     float64
	Age      int
	Aircraft int
}

func main() {
	donnie := Pilot{
		Name:     "Donnie",
		Life:     100.0,
		Age:      24,
		Aircraft: AIRCRAFT1,
	}
	fmt.Println(donnie)
}
