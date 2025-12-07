package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	fmt.Println(fichier("s06_1_suite1.txt"))
	fmt.Println(fichier("s06_1_suite2.txt"))
	fmt.Println(fichier("s06_1_suite3.txt"))
	fmt.Println(fichier("s06_1_suite4.txt"))
	fmt.Println(fichier("s06_1_suite5.txt"))
}

func fichier(nomFichier string) int {

	var compteur int

	var prec int
	var cour int
	var suiv int

	var consta bool
	file, err := os.Open(nomFichier)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	lecteur := bufio.NewScanner(file)

	for {
		if lecteur.Scan() != false {
			suiv, _ = strconv.Atoi(lecteur.Text())
		} else {
			break
		}
		if cour == 0 || suiv == 0 {
			prec = cour
			cour = suiv
			continue
		}
		consta = (cour == suiv)

		if prec < cour && cour > suiv {
			compteur = compteur + 1
		}

		if prec > cour && cour < suiv {
			compteur = compteur + 1
		}

		if suiv != cour {
			prec = cour
			cour = suiv
		}
	}
	if cour == suiv && suiv == prec && suiv == 0 {
		return 0
	}
	if consta {
		return 1
	}

	return compteur + 2
}
