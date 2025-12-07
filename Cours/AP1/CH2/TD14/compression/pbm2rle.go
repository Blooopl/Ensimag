package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	lecteur := bufio.NewScanner(os.Stdin)

	lecteur.Scan()
	src := lecteur.Text()
	dest := "res" + src

	creationFichier(dest)

	file, err := os.Open(src)

	if err != nil {
		log.Fatal(err)
	}

	lecteur = bufio.NewScanner(file)
	lecteur.Scan()
	format := lecteur.Text()

	if format == "P1" {
		decompression(src, dest)
	} else if format == "RLE" {
		compression(src, dest)
	} else {
		fmt.Fprintln(os.Stderr, "Le format du fichier n'est pas correcte")
	}

	file.Close()

}

func compression(src, dest string) {
	lecteur = bufio.NewScanner(file)
	lecteur.Scan()
	format := lecteur.Text()
}

func decompression(src, dest string) {}

func creationFichier(nomFichier string) {

	fich, err := os.Create(nomFichier)

	if err != nil {
		log.Fatal(err)
	}

	fich.Close()

}

func ajoutFichier(nomFichier, contenu string) {

	file, err := os.OpenFile(nomFichier, os.O_WRONLY|os.O_APPEND, 0)

	if err != nil {
		log.Fatal(err)
	}

	file.WriteString(contenu)

	file.Close()
}
