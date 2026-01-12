package main

import (
	"math/rand"
)

const valSup = 10

func initTab(taille int) []int {
	var tab = make([]int, taille)
	// Si le jour de l'examen vous utiliser un for idx := 0; idx < len(tab); ... : ça sera tout aussi valable !
	for idx := range len(tab) {
		tab[idx] = rand.Intn(valSup)
	}
	return tab
}

func parcourirEchangerTab(tab []int, inv bool) bool {
	// On va factoriser un peu le code en utilisant des variables bien choisies
	var debut, fin, incr int
	// Toute la difficulté de la question consiste à bien initialiser ces variables : c'est là où un dessin est utile pour ne pas se tromper !
	if inv {
		debut, fin, incr = len(tab)-2, -1, -1
	} else {
		debut, fin, incr = 0, len(tab)-1, 1
	}
	modif := false
	for idx := debut; idx != fin; idx += incr {
		if tab[idx] > tab[idx+1] {
			tab[idx], tab[idx+1] = tab[idx+1], tab[idx]
			modif = true
		}
	}
	return modif
}

func trierTab(tab []int) {
	// Cas du tableau vide et du singleton
	if len(tab) < 2 {
		return
	}
	for inv := false; parcourirEchangerTab(tab, inv); inv = !inv {
	}
}
