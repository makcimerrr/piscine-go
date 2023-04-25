package piscine

func BasicJoin(elems []string) string {
	full := ""
	for _, c := range elems {
		full = full + c
	}
	return full
}
