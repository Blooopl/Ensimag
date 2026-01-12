package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

const valSup = 10

// Un tableau de 3 entiers défini en variable globale.
// Attention : le système de typage de Go est vraiment très strict : la taille fait partie du type !
var tab [3]int

// Exercices de base

func remplir() {
	// Si VSCode souligne idx en disant quelque-chose du genre "for loop can be modernized using range over int"...
	// ... vous pouvez l'ignorer, on verra plus tard dans le cours comment écrire des boucles avec le mot-clé range !

	// C'est un parcours de type 1 comme vu en TD
	for idx := 0; idx < len(tab); idx++ {
		tab[idx] = rand.Intn(valSup)
	}
}

func afficher() {
	// C'est un parcours de type 1 comme vu en TD
	fmt.Print("[")
	for idx := 0; idx < len(tab); idx++ {
		fmt.Print(" ", tab[idx])
	}
	fmt.Println(" ]")
}

func valDansTab(val int) bool {
	// C'est une recherche comme vue en TD
	var idx int
	for idx = 0; (idx < len(tab)) && (tab[idx] != val); idx++ {
	}
	return idx < len(tab)
}

func inverser() {
	// On pourrait faire encore un parcours en divisant la taille par 2 mais on a envie de changer un peu !
	for deb, fin := 0, len(tab)-1; deb < fin; deb, fin = deb+1, fin-1 {
		tab[deb], tab[fin] = tab[fin], tab[deb]
	}
}

// Pile bornées

// Par convention, le sommet de pile pointe sur la dernière case occupée...
// c'est à dire -1 quand la pile est vide (-1 n'est pas une case valide du tableau bien sûr !)
var sommetPile int = -1

func afficherPile() {
	fmt.Printf("(sommet de pile = %v) [", sommetPile)
	for idx := 0; idx <= sommetPile; idx++ {
		fmt.Print(" ", tab[idx])
	}
	fmt.Println(" ]")
}

func empiler(val int) error {
	if sommetPile >= len(tab)-1 {
		return errors.New("erreur : la pile est pleine")
	}
	sommetPile++
	tab[sommetPile] = val
	return nil
}

func depiler() (int, error) {
	if sommetPile < 0 {
		return -1, errors.New("erreur : la pile est vide") // -1 : valeur non significative
	}
	val := tab[sommetPile]
	sommetPile--
	return val, nil
}

// File bornées

var ixPlusAncien = 0

var nbrElem = 0

// On peut factoriser l'opération "(idx + incr) % len(tab)" qu'on utilise souvent
func plusModulo(idx, incr int) int {
	return (idx + incr) % len(tab)
}

func afficherFile() {
	fmt.Printf("(indice du plus ancien = %v, nombre d'élément = %v) [", ixPlusAncien, nbrElem)
	for cpt, idx := 0, ixPlusAncien; cpt < nbrElem; cpt++ {
		fmt.Print(" ", tab[idx])
		idx = plusModulo(idx, 1)
	}
	fmt.Println(" ]")
}

func inserer(val int) error {
	if nbrElem == len(tab) {
		return errors.New("erreur : la file est pleine")
	}
	tab[plusModulo(ixPlusAncien, nbrElem)] = val
	nbrElem++
	return nil
}

func retirer() (int, error) {
	if nbrElem == 0 {
		return -1, errors.New("erreur : la file est vide") // -1 : valeur non significative
	}
	val := tab[ixPlusAncien]
	ixPlusAncien = plusModulo(ixPlusAncien, 1)
	nbrElem--
	return val, nil
}

// Test général de la pile et de la file

func testStruct(pile bool) {
	var nomStruct string
	if pile {
		nomStruct = "Pile"
	} else {
		nomStruct = "File"
	}
	fmt.Printf("\n%s bornée : entrez les valeur une par une :\n", nomStruct)
	fmt.Println("- les chiffres dans [0..9] seront insérés")
	fmt.Println("- -1 signifie retirer une valeur")
	fmt.Println("- toutes les autres valeurs seront ignorées")
	fmt.Println("- ctrl-d pour terminer")
	lecteur := bufio.NewScanner(os.Stdin)
	if pile {
		afficherPile()
	} else {
		afficherFile()
	}
	for lecteur.Scan() {
		val, err := strconv.Atoi(lecteur.Text())
		if err != nil || val < -1 || val >= valSup {
			fmt.Fprintln(os.Stderr, "erreur : valeur incorrecte")
			continue
		}
		if val == -1 {
			if pile {
				val, err = depiler()
			} else {
				val, err = retirer()
			}
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			fmt.Println("-> valeur retirée :", val)
		} else {
			if pile {
				err = empiler(val)
			} else {
				err = inserer(val)
			}
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
		}
		if pile {
			afficherPile()
		} else {
			afficherFile()
		}
	}
}

func main() {
	fmt.Print("Tableau initial : ")
	remplir()
	afficher()
	fmt.Print("Tableau inversé : ")
	inverser()
	afficher()
	for idx := 0; idx < 5; idx++ {
		val := rand.Intn(valSup*2) - valSup
		if valDansTab(val) {
			fmt.Println(val, "est dans tab ")
		} else {
			fmt.Println(val, "n'est pas dans tab ")
		}
	}
	testStruct(true)
	testStruct(false)
}
