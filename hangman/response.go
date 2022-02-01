package hangman

import "fmt"

//Cette fonction est une fonction de transition qui permet d'utiliser deux fonction "Victory" qui affiche un texte si la lettre est
//présente et d'afficher un texte si le mot est complet.
//Et la fonction "position" qui permet d'afficher les positions du pendu, il affiche aussi le nombre de vie restante à l'utilisateur.
func Response(number rune) {
	var a int
	for i := 0; i < len(UseWords); i++ {
		if string(number) == UseWords[i] {
			a++
		}
	}
	if a == 0 {
		UseWords = append(UseWords, string(number))
	} else {
		fmt.Println("Vous avez déjà écris cette lettre.")
		fmt.Println("Lettres déjà testées :", UseWords)
		fmt.Println("Il vous reste :", Vie, "vies")
		fmt.Println("-----------------------------------------------")
	}

	if a == 0 {
		if VerifLetter(string(number)) {
			fmt.Println("Bravo, vous avez trouvé une lettre !")
		} else {
			fmt.Println("Dommage cette lettre n'est pas dans le mot.")
			Vie--
			Position()
		}
		if Vie > 1 {
			fmt.Println("Il vous reste :", Vie, "vies")
		} else {
			fmt.Println("Il vous reste :", Vie, "vie")
		}
		PrintLetter()
		Victory()
		fmt.Println("-----------------------------------------------")
	}
}
