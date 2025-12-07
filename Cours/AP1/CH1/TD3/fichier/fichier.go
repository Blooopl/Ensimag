package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	nomFichier := "victor.txt"
	contenu_rajout := "je suis là"
	creationFichier(nomFichier)

	lectureFichier(nomFichier)

	ajoutFichier(nomFichier, contenu_rajout)

}

func creationFichier(nomFichier string) {

	fich, err := os.Create(nomFichier)

	if err != nil {
		log.Fatal(err)
	}

	defer fich.Close()

	fmt.Fprintln(fich, "Demain dès l'aube")

}

func lectureFichier(nomFichier string) {

	file, err := os.Open(nomFichier)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	lecteur := bufio.NewScanner(file)
	lecteur.Scan()
	text := lecteur.Text()
	fmt.Fprintln(os.Stdout, text)

}

func ajoutFichier(nomFichier, contenu string) {

	file, err := os.OpenFile(nomFichier, os.O_WRONLY|os.O_APPEND, 0)

	if err != nil {
		log.Fatal(err)
	}

	file.WriteString(contenu)

	file.Close()
}
