package hangman

import "fmt"

//Cette fonction permet de prendre seulement une lettre ou un mot, il vérifit aussi les caractères entrées par l'utilisateur.
//Transforme les majuscules en minuscules.
func Letter() {
	a := 0
	for a <= 0 {
		var number []rune
		fmt.Println("Choisissez une lettre :")
		for i := 0; i < len(Interaction); i++ {
			number = append(number, rune(Interaction[i]))
		}
		switch Interaction {
		case Interaction:
			if len(number) == 1 {
				if number[0] >= 97 && number[0] <= 122 {
					a++
					Response(number[0])
				} else if number[0] >= 65 && number[0] <= 90 {
					a++
					number[0] += 32
					Response(number[0])
				} else {
					a++
				}
			} else if len(Interaction) > 1 {
				for i := 0; i < len(number); i++ {
					if number[i] >= 97 && number[i] <= 122 {
						a++
					} else if number[i] >= 65 && number[i] <= 90 {
						a++
						number[i] += 32
					}
				}
				if a == len(number) {
					Verifword(string(number))
				} else {
					a++
				}
			}
		}
	}
}

//Cette fonction permet de vérifier la lettre dans le mot, si elle est présente il retourne "true" au sinon il retourne "faux".
func VerifLetter(number string) bool {
	var index int
	number = string(number)
	for i := 0; i < len(Words); i++ {
		if Tab[i] == number {
			Hidewords[i] = number
			index++
		}
	}
	if index >= 1 {
		return true
	} else {
		return false
	}
}

func PrintLetter() {
	fmt.Print("\n")
	for i := 0; i < len(Hidewords); i++ {
		fmt.Print(Hidewords[i], " ")
	}
	fmt.Print("\n", "\n")
}
