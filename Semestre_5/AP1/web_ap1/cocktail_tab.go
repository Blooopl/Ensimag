package main

import (
	"math/rand"
)

// On travaille sur des chiffres dans [0..9]
const valSup = 10

// Alloue et initialise un tableau avec des chiffres aléatoires
func initTab(taille int) []int {
	var tab = make([]int, taille)
	for idx := 0; idx < len(tab); idx++ {
		tab[idx] = rand.Intn(valSup)
	}
	return tab
}

func parcourirEchangerTab(tab []int, inv bool) bool {
	// supprimer cette ligne quand vous aurez implanté la fonction !
	return false
}

func trierTab(tab []int) {
}
