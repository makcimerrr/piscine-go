package piscine

func ActiveBits(n int) int {
	count := 0
	for n != 0 {
		if n&1 == 1 { //si trouve le bit 1 ajoute 1 au compteur
			count++
		}
		n = n >> 1 //decale d'un crant pour lire les 1
	}
	return count
}

/*func main() {
	fmt.Println(ActiveBits(7))
}*/
