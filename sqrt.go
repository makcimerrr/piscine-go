package piscine

import "math"

func Sqrt(nb int) int {
	if nb < 0 || math.MaxInt16 < nb {
		return 0
	} else {
		for i := 0; i < 101; i++ {
			if nb == ii {
				nb = i
				break
			} else if ii > nb {
				nb = 0
				break
			}
		}
	}
	return nb
}
