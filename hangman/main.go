package hangman

var Words string       //Cette variable va stocker le mot choisi par le programme aléatoirement.
var Hidewords []string //Cette variable permet de stocker les lettres de la variable "words" et les remplacer par des underscores(_)
var Tab []string
var Interaction string //Cette variable permet de stocker la lettre ou le mot rentrer par l'utilisateur.
var UseWords []string  //Cette variable permet d'afficher dans le programme en cours les lettres ou les mots utiliser par l'utilisateur.
var Vie int = 10       //Cette variable stock le nombre de l'utilisateur qui sera soustrait par 1 ou par 2 si la lettre ou le mot rentrer ne correspond pas.
var Lines []string
var VarLines []string

func Exec() {
	Mot()
	Random(Tab)
	Letter()
}
