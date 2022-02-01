package hangman

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

//Cette fonction permet de lier le fichier importjose.go au fichier hangman.txt.
func Position() []string {
	readFile, err := os.Open("./Annexe/hangman.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		Lines = append(Lines, "\n"+fileScanner.Text())
	}
	readFile.Close()
	if Vie == 10 {
		VarLines = Lines[0:0]
		return Lines
	}
	if Vie == 9 {
		VarLines = Lines[0:7]
		return Lines
	} else if Vie == 8 {
		VarLines = Lines[8:15]
		fmt.Println("GlaDoS : Le mot est pourtant simple, réfléchissez.")
		return Lines
	} else if Vie == 7 {
		VarLines = Lines[16:23]
		return Lines
	} else if Vie == 6 {
		VarLines = Lines[24:31]
		fmt.Println("GlaDoS : Y aurait-il une défaillance dans vôtre programme ?.")
		return Lines
	} else if Vie == 5 {
		VarLines = Lines[32:39]
		return Lines
	} else if Vie == 4 {
		VarLines = Lines[40:47]
		fmt.Println("GlaDoS : Je crains que vous vous soyez en position de faiblesse.")
		return Lines
	} else if Vie == 3 {
		VarLines = Lines[48:55]
		return Lines
	} else if Vie == 2 {
		VarLines = Lines[56:63]
		fmt.Println("GlaDoS : Je vais ajouter le mot échec à vôtre base de données.")
		return Lines
	} else if Vie == 1 {
		VarLines = Lines[64:71]
		return Lines
	} else if Vie == 0 {
		VarLines = Lines[72:79]
		fmt.Println("Le mot était :", Words)
		fmt.Println("GlaDoS : On dirai que mon intelligence artificielle est supérieur à la vôtre.")
	}
	return VarLines
}
