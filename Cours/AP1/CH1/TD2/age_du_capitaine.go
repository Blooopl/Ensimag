package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	lecteur := bufio.NewScanner(os.Stdin)

	fmt.Fprintln(os.Stdout, "Entrez votre année de naissance")

	lecteur.Scan()
	const annee int = 2025
	entree := lecteur.Text()

	var naissance int
	var erreur error

	naissance, erreur = strconv.Atoi(entree)

	if erreur == nil {
		println("Vous avez", annee-naissance, "années")
	} else {
		log.SetFlags(log.Flags() | log.Lshortfile)
		log.Fatal("naissance invalide")
	}

}
