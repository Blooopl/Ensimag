package main

import (
	"fmt"
	"math/rand"
)

// Taille du tableau, c'est-à-dire l'intervalle des résultats de la fonction de hachage [0..3]
const taille = 4

// On manipule des chiffres dans [0..9]
const valSup = 10

// Nombre d'éléments qu'on insère pour les tests
const nbrElemTest = 8

// La table de hachage est une structure avec :
// - un champ qui garde le nombre d'éléments dans la table
// - un tableau de listes qui contiennent les éléments
type tabListes struct {
	nbrElem int
	tab     [taille]liste
}

// Une liste est une structure avec :
// - le nombre d'éléments dans la liste ;
// - un pointeur sur la cellule de tête de la liste, qui est un élément fictif
type liste struct {
	nbrElem int
	tete    *cellule
}

// Cellule classique
type cellule struct {
	val  int
	suiv *cellule
}

// La fonction de hachage prend la valeur d'un élément en paramètre et renvoie une valeur de hachage dans [0..3]
func hachage(val int) int {
	return val % taille
}

func créerTabListe() *tabListes {
	var tabL *tabListes = new(tabListes)
	// Ce type de for range est équivalent à for idx := 0; idx < taille; idx ++
	for idx := range taille {
		// On crée le fictif en tête de chaque liste (on l'alloue, et par défaut val et suiv seront nuls)
		tabL.tab[idx].tete = new(cellule)
	}
	return tabL
}

func afficherListe(lsc *cellule) {
	for ; lsc != nil; lsc = lsc.suiv {
		fmt.Print(lsc.val, " -> ")
	}
	fmt.Println("FIN")
}

func afficherTabListes(tabL *tabListes) {
	fmt.Println("Nombre total d'éléments =", tabL.nbrElem)
	for idx := range taille {
		fmt.Printf("[%v] nombre d'éléments = %v, valeurs = ", idx, tabL.tab[idx].nbrElem)
		afficherListe(tabL.tab[idx].tete.suiv)
	}
}

func insérerTabListes(tabL *tabListes, val int) {
	h := hachage(val)
	cell := new(cellule)
	cell.val, cell.suiv = val, tabL.tab[h].tete.suiv
	tabL.tab[h].tete.suiv = cell
	tabL.tab[h].nbrElem++
	tabL.nbrElem++
}

func supprimerTabListes(tabL *tabListes, val int) bool {
	h := hachage(val)
	var prec *cellule
	for prec = tabL.tab[h].tete; (prec.suiv != nil) && (prec.suiv.val != val); prec = prec.suiv {
	}
	if prec.suiv != nil {
		prec.suiv = prec.suiv.suiv
		tabL.tab[h].nbrElem--
		tabL.nbrElem--
		return true
	}
	return false
}

func main() {
	tabL := créerTabListe()
	afficherTabListes(tabL)
	fmt.Println()
	// Ici, on n'a même pas besoin d'indice : on veut juste que la boucle s'exécute nbrElemTest fois
	for range nbrElemTest {
		val := rand.Intn(valSup)
		fmt.Println("***** Insertion de la valeur", val, "*****")
		insérerTabListes(tabL, val)
		afficherTabListes(tabL)
		fmt.Println()
	}
	for tabL.nbrElem > 0 {
		val := rand.Intn(valSup)
		if supprimerTabListes(tabL, val) {
			fmt.Println("***** Suppression de la valeur", val, "*****")
			afficherTabListes(tabL)
			fmt.Println()
		}
	}
}
