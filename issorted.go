package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	n := len(a)
	if n == 0 || n == 1 {
		return true
	}
	if f(a[0], a[n-1]) > 0 {
		for i := 1; i < n; i++ {
			if f(a[i-1], a[i]) < 0 {
				return false
			}
		}
	} else {
		for i := 1; i < n; i++ {
			if f(a[i-1], a[i]) > 0 {
				return false
			}
		}
	}
	return true
}
