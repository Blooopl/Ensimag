package main

import (
	"fmt"
	"math/rand"
)

const tailleMaxTab = 10

const tailleDonnées = 1_000_000

const premierChar = 'A'
const dernierChar = 'Z'

type entrée struct {
	cle     *rune
	données *[tailleDonnées]byte
}

type tableau []*entrée

func main() {
	tab := make(tableau, tailleMaxTab)
	initTab(tab)
	afficherTab(tab)

	fmt.Println("")
	triBulle(tab)
	afficherTab(tab)
}

func creerEntrée() *entrée {

	pointeur_rune := new(rune)
	*pointeur_rune = rune('A' + rand.Intn(25))

	pointeur_entrée := new(entrée)
	pointeur_entrée.cle = pointeur_rune

	return pointeur_entrée
}

func initTab(tab tableau) {
	for i := 0; i < len(tab); i++ {
		tab[i] = creerEntrée()
	}
}

func afficherTab(tab tableau) {
	for i := 0; i < len(tab); i++ {
		runee := *((*(tab[i])).cle)
		fmt.Println(string(runee))
	}
}

func triBulle(tab tableau) {

	compteur := 1
	for compteur != 0 {
		compteur = 0
		for i := 0; i < len(tab)-1; i++ {
			runee_i := *((*(tab[i])).cle)
			rune_i_1 := *((*(tab[i+1])).cle)

			if rune_i_1 < runee_i {
				tab[i], tab[i+1] = tab[i+1], tab[i]
				compteur++
			}

		}
	}
}
