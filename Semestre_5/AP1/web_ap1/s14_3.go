package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Lit la ligne contenant la largeur et la hauteur de l'image, et les renvoie comme entiers.
func lireLargeurHauteur(lecteur *bufio.Scanner) (int, int) {
	lecteur.Scan()
	largHaut := strings.Fields(lecteur.Text())
	larg, err := strconv.Atoi(largHaut[0])
	if err != nil {
		log.Fatal(fmt.Errorf("erreur : la largeur \"%s\" de l'image n'est pas un entier", largHaut[0]))
	}
	haut, err := strconv.Atoi(largHaut[1])
	if err != nil {
		log.Fatal(fmt.Errorf("erreur : la hauteur \"%s\" de l'image n'est pas un entier", largHaut[1]))
	}
	return larg, haut
}

// Renvoie '0' ssi etat vaut '1' et '1' sinon.
func inverserEtat(etat rune) rune {
	if etat == '0' {
		return '1'
	}
	return '0'
}

// Compresse l'entrée en écrivant le résultat sur la sortie standard.
func compresser(lecteur *bufio.Scanner) {
	fmt.Println("RLE")
	cptPoints := 0
	etat := '0'
	nbrRepetitions := 0
	somRepetitions := 0
	larg, haut := lireLargeurHauteur(lecteur)
	fmt.Println(larg, haut)
	for lecteur.Scan() {
		for _, car := range lecteur.Text() {
			if car == etat {
				cptPoints++
				nbrRepetitions++
			} else if car == inverserEtat(etat) {
				cptPoints++
				etat = inverserEtat(etat)
				fmt.Println(nbrRepetitions)
				somRepetitions += nbrRepetitions
				nbrRepetitions = 1
			} // On ignore tous les autres caractères
		}
	}
	fmt.Println(nbrRepetitions)
	somRepetitions += nbrRepetitions
	if cptPoints != larg*haut {
		log.Fatal(fmt.Errorf("erreur : le nom de points (%v) est différent de largeur x hauteur (%v)", cptPoints, larg*haut))
	}
	if somRepetitions != larg*haut {
		log.Fatal(fmt.Errorf("erreur : la somme des répétitions (%v) est différente de largeurs x hauteur (%v)", somRepetitions, larg*haut))
	}
}

// Décompresse l'entrée en écrivant le résultat sur la sortie standard.
func decompresser(lecteur *bufio.Scanner) {
	fmt.Println("P1")
	etat := '0'
	larg, haut := lireLargeurHauteur(lecteur)
	fmt.Println(larg, haut)
	for lecteur.Scan() {
		nbr, _ := strconv.Atoi(lecteur.Text()) // On considère que le fichier est correct (énoncé)
		for range nbr {
			fmt.Printf("%c", etat)
		}
		etat = inverserEtat(etat)
	}
	fmt.Println()
}

func main() {
	lecteur := bufio.NewScanner(os.Stdin)
	lecteur.Scan()
	if lecteur.Text() == "P1" {
		compresser(lecteur)
	} else if lecteur.Text() == "RLE" {
		decompresser(lecteur)
	} else {
		log.Fatal("erreur : format d'entrée inconnu")
	}
}
