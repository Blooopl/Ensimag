package main

import "fmt"

type cellule struct {
	val  int
	prec *cellule
	suiv *cellule
}

type desc struct {
	tete  *cellule
	queue *cellule
}

type liste *desc

func créer() liste {
	queue := new(cellule)
	tete := new(cellule)

	queue.suiv = tete
	queue.prec = tete
	tete.suiv = queue
	tete.prec = queue

	var dsc desc
	dsc.tete = tete
	dsc.queue = queue

	return &dsc
}

func afficher(ldc liste, teteVersQueue bool) {
	var pointeur *cellule
	if teteVersQueue {
		fmt.Print("T <-> ")
		for pointeur = ldc.tete.suiv; pointeur != ldc.queue; pointeur = pointeur.suiv {
			fmt.Print(pointeur.val, " <-> ")
		}
		fmt.Println("Q")
	} else {
		fmt.Print("Q <-> ")
		for pointeur = ldc.queue.prec; pointeur != ldc.tete; pointeur = pointeur.prec {
			fmt.Print(pointeur.val, " <-> ")
		}
		fmt.Println("T")
	}
}

func insérer(ldc liste, val int, enTete bool) {
	var cell cellule
	cell.val = val

	if enTete {
		cell.suiv = ldc.tete.suiv
		cell.prec = ldc.tete
		cell.suiv.prec = &cell
		ldc.tete.suiv = &cell
	} else {
		cell.prec = ldc.queue.prec
		cell.suiv = ldc.queue
		cell.prec.suiv = &cell
		ldc.queue.prec = &cell
	}

}

func extraire(ldc liste, enTete bool) int {
	var val int
	if enTete {
		val = ldc.tete.suiv.val

		ldc.tete.suiv.suiv.prec = ldc.tete
		ldc.tete.suiv = ldc.tete.suiv.suiv

		return val
	} else {
		val = ldc.queue.prec.val

		ldc.queue.prec.prec.suiv = ldc.queue
		ldc.queue.prec = ldc.queue.prec.prec

		return val
	}
}
func main() {
	ldc := créer()

	insérer(ldc, 1, true)
	insérer(ldc, 2, true)
	insérer(ldc, 3, true)
	insérer(ldc, 4, true)
	insérer(ldc, 5, true)
	insérer(ldc, 6, true)
	insérer(ldc, 7, true)
	insérer(ldc, 8, true)
	insérer(ldc, 9, true)
	insérer(ldc, 10, true)
	insérer(ldc, 11, true)
	insérer(ldc, 12, true)

	échanger(ldc.tete.suiv.suiv.suiv.suiv.suiv)

	afficher(ldc, true)
}

func échanger(ptr *cellule) *cellule {

	cell_1 := ptr.prec
	cell_2 := cell_1.suiv
	cell_3 := cell_2.suiv
	cell_4 := cell_3.suiv

	cell_1.suiv = cell_3
	cell_3.prec = cell_1
	cell_3.suiv = cell_2
	cell_2.prec = cell_3
	cell_2.suiv = cell_4
	cell_4.prec = cell_2

	return ptr.suiv
}

func trierNain(ldc liste) {

}
