package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

// Initialise le tableau des booléens avec toutes les cases à true, sauf la case 0 qui est false par défaut.
func initialiser(taille int) []bool {
	lesPremiers := make([]bool, taille)
	for idx := 1; idx < taille; idx++ {
		lesPremiers[idx] = true
	}
	return lesPremiers
}

// Marque à false tous les multiples de prem.
func filtrer(lesPremiers []bool, prem int) {
	for num := prem * 2; num <= len(lesPremiers); num += prem {
		// Attention au -1 car l'entier 1 est représenté par la case d'indice 0 et ainsi de suite
		lesPremiers[num-1] = false
	}
}

// Affiche les nombres premiers inférieurs ou égaux à valMax.
// Note : on pourrait aussi arrêter l'affichage à Sqrt(valMax) mais on affiche tout le tableau pour être sûrs qu'on ne s'est pas trompé.
func afficher(lesPremiers []bool, valMax int) {
	fmt.Print("Les nombres premiers inférieurs ou égaux à ", valMax, " sont : ")
	for idx := 0; idx < len(lesPremiers); idx++ {
		if lesPremiers[idx] {
			// Idem dans l'autre sens
			fmt.Print(idx+1, " ")
		}
	}
	fmt.Println()
}

// Programme principal.
func main() {
	fmt.Print("Entrez la valeur maximale : ")
	lecteur := bufio.NewScanner(os.Stdin)
	lecteur.Scan()
	valMax, err := strconv.Atoi(lecteur.Text())
	if (err != nil) || (valMax < 2) {
		log.Fatal("erreur : la valeur limite doit être un entier supérieur ou égal à 2")
	}
	lesPremiers := initialiser(valMax)
	// L'appel à math.Floor n'est pas nécessaire en fait car valMax >= 0 donc int(valMax) donne directement le même résultat
	for num := 2; num <= int(math.Floor(math.Sqrt(float64(valMax)))); num++ {
		if lesPremiers[num-1] {
			filtrer(lesPremiers, num)
		}
	}
	afficher(lesPremiers, valMax)
}
