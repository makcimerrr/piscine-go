package piscine

func BasicAtoi(s string) int {
	result := 0
	for _, c := range s { // variable c utilisé pour stocker chauque caractere de la chaine
		digit := int(c - '0') // conversion de c en un chiffre
		result = result*10 + digit
	}
	return result
}
