package piscine

func Sqrt(nb int) int {
	if nb == 0 {
		return 0
	}
	racine := nb
	for racine*racine > nb {
		racine = (racine + nb/racine) / 2
	}
	if racine*racine == nb {
		return racine
	}
	return racine - 1
}
