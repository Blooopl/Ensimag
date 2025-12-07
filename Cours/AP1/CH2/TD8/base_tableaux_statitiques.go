package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func main() {
	inserer(4)
	inserer(5)
	inserer(8)

	afficherFile()
	retirer()
	retirer()
	inserer(5)
	inserer(6)
	afficherFile()
}

var taille int = 12
var tab [12]int
var valSup = 10

func remplir() {
	for i := 0; i < len(tab); i++ {
		tab[i] = rand.Intn(valSup)
	}
}

func afficher() {
	fmt.Print("[ ")
	for i := 0; i < len(tab); i++ {
		fmt.Print(tab[i])
		fmt.Print(" ")
	}
	fmt.Print("] ")
}

func valDansTab(val int) bool {
	for i := 0; i < len(tab); i++ {
		if tab[i] == val {
			return true
		}
	}
	return false
}

func inverser() {
	for i := 0; i < ((len(tab) - len(tab)%2) / 2); i++ {
		tab[i], tab[len(tab)-1-i] = tab[len(tab)-1-i], tab[i]
	}
}

var sommetPile int = -1

func afficherPile() {
	for i := 0; i <= sommetPile; i++ {
		fmt.Print(tab[i])
		fmt.Print(" ")
	}
	fmt.Println()
}

func empiler(val int) error {
	error_full := errors.New("Pile Pleine")

	if sommetPile >= 11 {
		return error_full
	}
	sommetPile++
	tab[sommetPile] = val
	return nil

}

func depiler() (int, error) {
	error_empty := errors.New("Pile vide")
	if sommetPile == 0 {
		return 0, error_empty
	}
	res := tab[sommetPile]
	sommetPile--
	return res, nil
}

var ixPlusAncien int = -1
var nbrElem int = 0

func plusModulo(ix, incr int) int {
	return ix + incr%taille
}

func afficherFile() {
	fmt.Print("[")
	for i := ixPlusAncien; i != plusModulo(ixPlusAncien, nbrElem); i = (i + 1) % taille {
		fmt.Print(" ")
		fmt.Print(tab[i])
	}
	fmt.Print("]")
}

func inserer(val int) error {
	error_full := errors.New("File Pleine")

	if nbrElem == 0 {
		ixPlusAncien = 0
	}
	if nbrElem >= taille {
		return error_full
	}

	tab[plusModulo(ixPlusAncien, nbrElem)] = val
	nbrElem++
	return nil

}

func retirer() (int, error) {
	error_empty := errors.New("File vide")
	if nbrElem == 0 {
		return 0, error_empty
	}

	res := tab[ixPlusAncien]
	ixPlusAncien++
	nbrElem--
	return res, nil
}
