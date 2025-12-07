package main

import "fmt"

type cellule struct {
	val  int
	suiv *cellule
}

type liste *cellule

func créer() liste {
	fic := new(cellule)
	fic.suiv = fic
	return fic
}

func estVide(lct liste) bool {
	return lct.suiv == lct
}

func afficher(lct liste) {
	fmt.Print("Début : ")
	for pointeur := lct.suiv; pointeur != lct; pointeur = pointeur.suiv {
		fmt.Print(pointeur.val, " -> ")
	}
	fmt.Println("Fin")
}

func insérerListeCircTriée(lct liste, val int) {
	cell := new(cellule)
	cell.val = val

	pointeur := lct

	for ; pointeur.suiv.val < val && pointeur.suiv != lct; pointeur = pointeur.suiv {

	}
	cell.suiv = pointeur.suiv
	pointeur.suiv = cell

}

func supprimerListeCircTriée(lct liste, val int) bool {
	if estVide(lct) {
		return false
	}
	prec := lct
	for ; prec.suiv.val != val && prec.suiv != lct; prec = prec.suiv {
	}

	if prec.suiv == lct {
		return false
	} else {
		prec.suiv = prec.suiv.suiv
		return true
	}

}

func supprimerDoublonsListeCircTriée(lct liste) {
	prec := lct
	for ; prec.suiv != lct; prec = prec.suiv {
		if prec.suiv.val == prec.val {
			supprimerListeCircTriée(lct, prec.val)
		}
	}
}

func découperListeCircTriée(lct liste) (liste, liste) {
	l1 := créer()
	l2 := créer()

	if estVide(lct) {
		return l1, l2
	}

	i := 0
	for pointeur := lct.suiv; pointeur != lct; i++ {
		if i%2 == 0 {
			insérerListeCircTriée(l1, pointeur.val)
		} else {
			insérerListeCircTriée(l2, pointeur.val)
		}
		pointeur = pointeur.suiv
	}

	return l1, l2
}
func main() {
	lct := créer()
	afficher(lct)
	insérerListeCircTriée(lct, 1)
	afficher(lct)
	insérerListeCircTriée(lct, 4)
	afficher(lct)
	insérerListeCircTriée(lct, 9)
	afficher(lct)
	insérerListeCircTriée(lct, 4)
	afficher(lct)
	insérerListeCircTriée(lct, 9)
	afficher(lct)
	insérerListeCircTriée(lct, 4)
	afficher(lct)
	insérerListeCircTriée(lct, 12)
	afficher(lct)
	insérerListeCircTriée(lct, 7)
	afficher(lct)
	insérerListeCircTriée(lct, 8)
	afficher(lct)
	insérerListeCircTriée(lct, 2)
	afficher(lct)

	var l1 liste
	var l2 liste

	l1, l2 = découperListeCircTriée(lct)
	afficher(l1)
	afficher(l2)
}
