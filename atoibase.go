package piscine

func indexInBase(c rune, base string) int {
	for i, b := range base {
		if b == c {
			return i
		}
	}
	return -1
}

func AtoiBase(s string, base string) int {
	//1-Vérification de la validité de la base
	if len(base) < 2 { //Une base doit contenir au moins 2 caractères
		return 0
	}
	seen := make(map[rune]bool) //sert a s'assurer que chaque carectere est unique
	//la carte est utilisée pour stocker les caractères de la base que nous avons déjà rencontrés.
	for _, c := range base {
		if c == '+' || c == '-' || seen[c] { //Une base ne doit pas contenir de caractères + ou - et ne pas avoir de caractere en double.
			return 0
		}
		seen[c] = true
	}
	/*La fonction parcourt chaque caractère de la base et elle vérifie si elle a déjà été rencontrée en consultant la carte "seen".
	Si le caractère est déjà présent dans la carte, donc deja rencontré donc la fonction cela retourne 0.
	Si le caractère n'a pas encore été rencontré, la fonction l'ajoute à la carte en l'associant à la valeur true.*/

	//2-Conversion du nombre dans la base donnée en entier décimal
	result := 0
	baseLEN := len(base)
	for _, c := range s {
		digit := indexInBase(c, base) //permet de récupérer chaque caractere de de la chaine "c" a chaque itération
		if digit == -1 {
			return 0
		}
		result = result*baseLEN + digit //revient a faire lindex de la base correspondante plus la base demandé multiplié par le resultat précédent, sachant qu'on part de 0
	}
	return result
}
