package main

import "fmt"

type cellule struct {
	val int
}

type ptrCellule *cellule

func main() {
	var cell cellule
	fmt.Printf("Adresse de la cellule = %p\n", &cell)
	parCopie(cell)

	cell.val = 5
	fmt.Println("Valeur dans la structure =", cell.val)
	var ptrCell ptrCellule = &cell
	fmt.Println("Valeur dans la structure =", (*ptrCell).val)
	fmt.Println("Valeur dans la structure =", ptrCell.val)
	ptrCell.val = 10
	fmt.Println("Valeur dans la structure =", cell.val)

	ptrCell = new(cellule)
	fmt.Printf("Valeur du pointeur = %p, adresse de cell = %p\n", ptrCell, &cell)
	ptrCell.val = 20
	fmt.Println("Valeur dans la structure =", ptrCell.val)

	var ptrPtrCell *ptrCellule
	ptrPtrCell = &ptrCell
	fmt.Println("Valeur dans la structure via le pointeur de pointeur =", (*ptrPtrCell).val)

}

func parCopie(cell cellule) {
	fmt.Printf("Adresse du paramètre = %p\n", &cell)
}
