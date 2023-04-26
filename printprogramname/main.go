package piscine

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	var filename string

	str, _ := os.Executable()
	lenght := len(str)

	for i := lenght - 1; i >= 0; i-- {
		if str[i] == '/' || str[i] == '\\' {
			filename = str[i+1:]
			break
		}
	}

	for _, c := range filename {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
