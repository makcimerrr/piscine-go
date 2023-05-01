package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("File name missing")
		return
	} else if len(args) > 2 {
		fmt.Println("Too many arguments")
		return
	}

	filename := args[1]
	message, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	os.Stdout.Write(message)
}
