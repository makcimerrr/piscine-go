package piscine

func Swap(a *int, b *int) {
	i := *a
	j := *b
	*a = j
	*b = i
}
