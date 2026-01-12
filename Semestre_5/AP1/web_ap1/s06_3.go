package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

type lesÉtats int

const (
	début lesÉtats = iota
	unBon
	deuxBons
	troisBons
)

func entreeValide(chaine string) (int, error) {
	val, err := strconv.Atoi(chaine)
	if err != nil || val < 0 || val > 9 {
		// -1 : valeur non significative
		return -1, errors.New("erreur : vous devez saisir des chiffres")
	}
	return val, nil
}

// Pour avoir un affichage plus clair
func étatVersChaîne(état lesÉtats) string {
	var chaîne string
	switch état {
	case début:
		chaîne = "début"
	case unBon:
		chaîne = "unBon"
	case deuxBons:
		chaîne = "deuxBons"
	case troisBons:
		chaîne = "troisBons"
	default:
		log.Fatal("erreur : cet état n'existe pas")
	}
	// Pour que les chaînes aient toutes la même taille
	return fmt.Sprintf("%9s", chaîne)
}

func main() {
	fmt.Println("Vous êtes face à une porte, équipée d'un digicode.")
	lecteur := bufio.NewScanner(os.Stdin)
	for état, ouverture := début, false; !ouverture; {
		fmt.Print("[", étatVersChaîne(état), "] Entrez un chiffre : ")
		lecteur.Scan()
		touche, err := entreeValide(lecteur.Text())
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		switch état {
		case début:
			if touche == 4 {
				état = unBon
			}
		case unBon:
			if touche == 0 {
				état = deuxBons
			} else if touche != 4 {
				état = début
			}
		case deuxBons:
			switch touche {
			case 9:
				état = troisBons
			case 4:
				état = unBon
			default:
				état = début
			}
		case troisBons:
			switch touche {
			case 6:
				ouverture = true
			case 4:
				état = unBon
			default:
				état = début
			}
		default:
			log.Fatal("erreur : cet état n'existe pas")
		}
	}
	fmt.Println("Vous pouvez entrer, bienvenue !")
}
