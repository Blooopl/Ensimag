package main

import (
	"fmt"
	"log"
	"os"
	"slices"
)

// Nombre maximal d'éléments pour les tests visuels
const tailleMaxVisu = 10

// Nombre maximal d'éléments pour les tests bruts
const tailleMax = 100

// Nombre d'itérations pour les tests bruts
const nbrIter = 1000

// Option de la ligne de commande pour lancer les tests bruts
const opt = "--brut"

// Tests visuels sur des tailles de 0 à tailleMaxVisu inclus
func testVisuel() {
	for taille := 0; taille <= tailleMaxVisu; taille++ {
		fmt.Println("*** Taille =", taille, "***")
		tab := initTab(taille)
		fmt.Print("Tableau initial : ")
		fmt.Println(tab)
		liste := initListe()
		remplirListe(liste, tab)
		fmt.Print("Liste initiale  : ")
		afficherListe(liste)
		trierTab(tab)
		fmt.Print("Tableau trié    : ")
		fmt.Println(tab)
		trierListe(liste)
		fmt.Print("Liste triée     : ")
		afficherListe(liste)
		fmt.Println()
	}
}

// Vérifie les tableaux et listes après le tri
func verifier(liste liste, tab []int, ref []int) {
	// Comparaison de notre tableau trié avec le tableau de référence
	if !slices.Equal(tab, ref) {
		fmt.Println()
		log.Fatal("erreur dans le tri du tableau")
	}
	// Comparaison que la liste trié contient les mêmes éléments dans le même ordre que dans le tableau de référence
	idx := 0
	for cour := liste.tete.suiv; cour != liste.queue; cour, idx = cour.suiv, idx+1 {
		if cour.val != ref[idx] {
			fmt.Println()
			log.Fatal("erreur dans le tri de la liste")
		}
	}
	// Vérification que la taille de la liste est bien égale à celle du tableau de référence
	if idx != len(ref) {
		fmt.Println()
		log.Fatal("liste de taille incorrecte")
	}
	// Parcours de la liste en sens inverse : ça ne fait rien, mais ça sautera si un chaînage est faux
	for cour := liste.queue; cour != liste.tete; cour = cour.prec {
	}
}

// Tests bruts sur nbrIter tableaux et listes de tailles de 0 à tailleMax inclus
func testBrut() {
	for taille := 0; taille <= tailleMax; taille++ {
		fmt.Print("\rTest en cours : ", taille, "%")
		for range nbrIter {
			tab := initTab(taille)
			ref := slices.Clone(tab)
			trierTab(tab)
			liste := initListe()
			remplirListe(liste, ref)
			trierListe(liste)
			slices.Sort(ref)
			verifier(liste, tab, ref)
		}
	}
	fmt.Println("\r100% des tests validés !")
	fmt.Println()
}

// Programme principal
func main() {
	switch {
	case len(os.Args) == 1:
		testVisuel()
	case len(os.Args) == 2 && os.Args[1] == opt:
		testBrut()
	default:
		log.Fatal("usage : go run cocktail*.go", "[", opt, "]")
	}
}
