package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path"
	"slices"
	"strconv"
	"strings"
)

const valSup = 10

type pile struct {
	tab []int
}

func creerPile() *pile {
	lifo := new(pile)         // on alloue la struct pile
	lifo.tab = make([]int, 0) // ET on doit créer aussi le tableau qu'elle contient
	return lifo
}

func afficherPile(lifo *pile) {
	fmt.Printf("{len/cap = %v/%v} (taille = %v) %v\n", len(lifo.tab), cap(lifo.tab), len(lifo.tab), lifo.tab)
}

func empiler(val int, lifo *pile) {
	// typage : lifo est un pointeur vers une pile =
	// - on pourrait avoir envie d'écrire "*lifo.tab" ce qui serait FAUX car l'opérateur . est prioritaire sur *
	// - pour l'écrire correctement, il faudrait plutôt "(*lifo).tab"
	// - mais Go nous facilite la vie en sachant déréférencer tout seul dans le cas d'un pointeur vers une struct
	lifo.tab = append(lifo.tab, val)
}

func depiler(lifo *pile) (int, error) {
	if len(lifo.tab) == 0 {
		return -1, errors.New("la pile est vide") // -1 : non-significatif
	}
	val := lifo.tab[len(lifo.tab)-1]
	lifo.tab = slices.Delete(lifo.tab, len(lifo.tab)-1, len(lifo.tab))
	return val, nil
}

type file struct {
	ixPlusAncien int
	nbrElem      int
	tab          []int
}

func creerFile() *file {
	fifo := new(file)
	fifo.ixPlusAncien = 0
	fifo.nbrElem = 0
	fifo.tab = make([]int, 0)
	return fifo
}

func afficherFile(fifo *file) {
	fmt.Printf("{len/cap = %v/%v} (indice du plus ancien = %v, nombre d'élément = %v) [", len(fifo.tab), cap(fifo.tab), fifo.ixPlusAncien, fifo.nbrElem)
	for idx := fifo.ixPlusAncien; idx < fifo.ixPlusAncien+fifo.nbrElem; idx++ {
		fmt.Print(" ", fifo.tab[idx])
	}
	fmt.Println(" ]")
}

func inserer(val int, fifo *file) {
	fifo.tab = append(fifo.tab, val)
	fifo.nbrElem++
}

func retirer(fifo *file) (int, error) {
	if fifo.nbrElem == 0 {
		return -1, errors.New("la file est vide") // -1 : non-significatif
	}
	val := fifo.tab[fifo.ixPlusAncien]
	fifo.ixPlusAncien++
	fifo.nbrElem--
	if fifo.nbrElem == 0 { // retassage gratuit
		fifo.ixPlusAncien = 0
		fifo.tab = fifo.tab[:0] // la tranche [:0] (ou [0:0] : c'est exactement pareil) est vide, car la borne sup est exclue
	}
	return val, nil
}

func afficher(lifo *pile, fifo *file) {
	if lifo != nil {
		afficherPile(lifo)
	} else {
		afficherFile(fifo)
	}
}

func extraire(lifo *pile, fifo *file) {
	var val int
	var err error
	if lifo != nil {
		val, err = depiler(lifo)
	} else {
		val, err = retirer(fifo)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Println("-> valeur retirée :", val)
	}
	afficher(lifo, fifo)
}

func ajouter(lifo *pile, fifo *file, val int) {
	if lifo != nil {
		empiler(val, lifo)
	} else {
		inserer(val, fifo)
	}
	afficher(lifo, fifo)
}

func testStruct(estPile bool) {
	var nomStruct string
	var lifo *pile
	var fifo *file
	if estPile {
		nomStruct = "Pile"
		lifo = creerPile()
	} else {
		nomStruct = "File"
		fifo = creerFile()
	}
	fmt.Printf("\n%s bornée : entrez les valeur une par une :\n", nomStruct)
	fmt.Println("- les chiffres dans [0..9] seront insérés")
	fmt.Println("- -1 signifie retirer une valeur")
	fmt.Println("- toutes les autres valeurs seront ignorées")
	fmt.Println("- ctrl-d pour terminer")
	lecteur := bufio.NewScanner(os.Stdin)
	for lecteur.Scan() {
		val, err := strconv.Atoi(lecteur.Text())
		if err != nil || val < -1 || val >= valSup {
			fmt.Fprintln(os.Stderr, "erreur : valeur incorrecte")
			afficher(lifo, fifo)
			continue
		}
		if val == -1 {
			extraire(lifo, fifo)
		} else {
			ajouter(lifo, fifo, val)
		}
	}
}

func main() {
	opts := []string{"--pile", "--file"}
	if len(os.Args) != 2 || !slices.Contains(opts, os.Args[1]) {
		fmt.Fprintln(os.Stderr, "usage :", path.Base(os.Args[0]), "(", strings.Join(opts, " | "), ")")
		os.Exit(1)
	}
	testStruct(os.Args[1] == opts[0])
}
