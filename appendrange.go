package piscine

func AppendRange(min, max int) []int {

	if min >= max {
		return nil //indique l'absence de valeur
	}
	var answer []int
	for i := min; i < max; i++ {
		answer = append(answer, i) //ajoute les elements a la tranche
	}

	return answer

}
