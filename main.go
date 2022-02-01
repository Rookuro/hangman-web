package main

import (
	"hangman"
	"html/template"
	"net/http"
	"strings"
)

type User struct {
	Interaction string
	Success     bool
	Hidewords   string
	Life        int
	UseWords    []string
	VarLines    string
	Win         bool
	Loose       bool
}

func main() {
	hangman.Mot()
	hangman.Random(hangman.Tab)
	hangman.Position()
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	http.HandleFunc("/", Accueil)
	http.HandleFunc("/hangman", Execute)
	http.ListenAndServe(":80", nil)
}

func Accueil(w http.ResponseWriter, r *http.Request) {
	var details User
	tmpl1 := template.Must(template.ParseFiles("index.html"))
	if hangman.Victory() {
		details = User{
			Win: true,
		}
	} else if hangman.Vie == 0 {
		details = User{
			Loose: true,
		}
	} else {
		details = User{
			Interaction: r.FormValue("lettre"),
			Hidewords:   strings.Join(hangman.Hidewords, " "),
			Life:        hangman.Vie,
			UseWords:    hangman.UseWords,
			VarLines:    strings.Join(hangman.VarLines, ""),
			Success:     true,
			Win:         false,
			Loose:       false,
		}
	}
	tmpl1.Execute(w, details)
}

func Execute(w http.ResponseWriter, r *http.Request) {
	details := User{
		Interaction: r.FormValue("lettre"),
		Success:     true,
	}
	hangman.Interaction = details.Interaction
	hangman.Letter()
	http.Redirect(w, r, "/", http.StatusFound)
}
