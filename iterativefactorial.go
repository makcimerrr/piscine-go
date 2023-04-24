package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 20 {
		return
	}
	if nb == 0 {
		return 1
	} else {
		return nb * IterativeFactorial(nb-1)
	}
}
