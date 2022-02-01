package hangman

//Cette fonction permet de vérifier si le mot donner est complet, il affichera en cas de victore deux phrases.
func Victory() bool {
	var compt int
	for i := 0; i < len(Hidewords); i++ {
		if Hidewords[i] == Tab[i] {
			compt++
		}
	}
	return compt == len(Tab)
}
