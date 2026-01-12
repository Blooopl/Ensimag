package main

import (
	"bufio"
	"fmt"
	"os"
)

type cellule struct {
	val int
}

type ptrCellule *cellule

func parCopie(cell cellule) {
	fmt.Printf("Adresse du paramètre = %p\n", &cell)
}

func main() {
	// Manipulations de pointeurs
	var entier int = 5
	fmt.Printf("Valeur de l'entier = %v, adresse de l'entier = %p\n", entier, &entier)
	var pointeur *int = &entier
	fmt.Printf("Valeur du pointeur = %v, adresse du pointeur = %p\n", pointeur, &pointeur)
	var pointeurNul *int
	fmt.Printf("Valeur du pointeur = %v, adresse du pointeur = %p\n", pointeurNul, &pointeurNul)
	*pointeur = 10
	fmt.Printf("Valeur de l'entier = %v, adresse de l'entier = %p\n", entier, &entier)
	fmt.Printf("Valeur du pointeur = %v, adresse du pointeur = %p\n", pointeur, &pointeur)
	fmt.Println()

	// Pointeurs sur des structures
	var cell cellule
	fmt.Printf("Adresse de la cellule = %p\n", &cell)
	parCopie(cell)
	fmt.Println()

	cell.val = 5
	fmt.Println("Valeur dans la structure =", cell.val)
	var ptrCell ptrCellule = &cell
	fmt.Println("Valeur dans la structure =", (*ptrCell).val)
	fmt.Println("Valeur dans la structure =", ptrCell.val)
	ptrCell.val = 10
	fmt.Println("Valeur dans la structure =", cell.val)
	fmt.Println()

	ptrCell = new(cellule)
	fmt.Printf("Valeur du pointeur = %p, adresse de cell = %p\n", ptrCell, &cell)
	ptrCell.val = 20
	fmt.Println("Valeur dans la structure =", ptrCell.val)
	fmt.Println()

	var ptrPtrCell *ptrCellule = &ptrCell
	fmt.Println("Valeur dans la structure via le pointeur de pointeur =", (*ptrPtrCell).val)
	fmt.Println()

	// Chaînes de caractères
	chaine := "abc"
	for idx := 0; idx < len(chaine); idx++ {
		fmt.Println(chaine[idx], string(chaine[idx]))
	}
	fmt.Println()
	chaine2 := "10€"
	fmt.Println(len(chaine2))
	for idx := 0; idx < len(chaine2); idx++ {
		fmt.Println(chaine2[idx], string(chaine2[idx]))
	}
	fmt.Println()
	tabRunes := []rune("10€")
	fmt.Println(len(tabRunes))
	for idx := 0; idx < len(tabRunes); idx++ {
		fmt.Println(tabRunes[idx], string(tabRunes[idx]))
	}
	fmt.Println()
	for _, code := range chaine2 {
		fmt.Println(code, string(code))
	}
	fmt.Println()

	// Type d'un lecteur
	lecteur := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type de lecteur : %T\n", lecteur)
}
