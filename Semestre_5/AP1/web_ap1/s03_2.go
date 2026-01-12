package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const nomFichier = "victor.txt"

func creationFichier() {
	fich, err := os.Create(nomFichier)
	if err != nil {
		log.Fatal(err)
	}
	// Erreur classique : utiliser Println au lieu de Fprintln...
	fmt.Fprintln(fich, "Demain, dès l'aube, à l'heure où blanchit la campagne,")
	fmt.Fprintln(fich, "Je partirai. Vois-tu, je sais que tu m'attends.")
	fich.Close()
}

func lectureFichier() {
	fich, err := os.Open(nomFichier)
	if err != nil {
		log.Fatal(err)
	}
	defer fich.Close()
	lecteur := bufio.NewScanner(fich)
	// On sait qu'il n'y a que deux lignes à lire dans le fichier, pas besoin de boucle
	// Scan renvoie un booléen qui vaut true ssi il a réussi à lire une ligne (auquel cas, on l'affiche ; sinon, on ignore)
	if lecteur.Scan() {
		fmt.Println(lecteur.Text())
	}
	if lecteur.Scan() {
		fmt.Println(lecteur.Text())
	}
}

func ajoutFichier() {
	fich, err := os.OpenFile(nomFichier, os.O_WRONLY|os.O_APPEND, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer fich.Close()
	fmt.Fprintln(fich, "J'irai par la forêt, j'irai par la montagne.")
	fmt.Fprintln(fich, "Je ne puis demeurer loin de toi plus longtemps.")
}

func main() {
	log.SetFlags(log.Flags() | log.Lshortfile)
	creationFichier()
	lectureFichier()
	ajoutFichier()
}
