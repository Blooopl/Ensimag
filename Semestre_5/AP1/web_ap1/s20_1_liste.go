package main

import "fmt"

type cellule struct {
	val  int
	prec *cellule
	suiv *cellule
}

type liste struct {
	tete  *cellule
	queue *cellule
}

func initListe() liste {
	var ldc = liste{new(cellule), new(cellule)}
	ldc.tete.prec, ldc.tete.suiv = nil, ldc.queue
	ldc.queue.prec, ldc.queue.suiv = ldc.tete, nil
	return ldc
}

func remplirListe(ldc liste, tab []int) {
	for idx := range len(tab) {
		cell := new(cellule)
		cell.val, cell.prec, cell.suiv = tab[idx], ldc.queue.prec, ldc.queue
		ldc.queue.prec.suiv = cell
		ldc.queue.prec = cell
	}
}

func afficherListe(ldc liste) {
	fmt.Print("T <-> ")
	for cour := ldc.tete.suiv; cour != ldc.queue; cour = cour.suiv {
		fmt.Print(cour.val, " <-> ")
	}
	fmt.Println("Q")
}

func echangerCellules(cell *cellule) *cellule {
	prec := cell.prec
	cour := cell
	suiv := cell.suiv
	suiv_suiv := suiv.suiv
	prec.suiv = suiv
	cour.prec = suiv
	cour.suiv = suiv_suiv
	suiv.prec = prec
	suiv.suiv = cour
	suiv_suiv.prec = cour
	return suiv
}

func listeVideOuSingleton(ldc liste) bool {
	return ldc.tete.suiv == ldc.queue || ldc.tete.suiv.suiv == ldc.queue
}

func parcourirEchangerListe(ldc liste, inv bool) bool {
	// Exactement le même principe que pour les tableaux
	var cour, fin *cellule
	if inv {
		cour, fin = ldc.queue.prec.prec, ldc.tete
	} else {
		cour, fin = ldc.tete.suiv, ldc.queue.prec
	}
	modif := false
	for cour != fin {
		if cour.val > cour.suiv.val {
			// Erreur classique : oublier de mettre à jour cour avec le pointeur renvoyé par echangerCellules
			cour = echangerCellules(cour)
			modif = true
		} else {
			if inv {
				cour = cour.prec
			} else {
				cour = cour.suiv
			}
		}
	}
	return modif
}

func trierListe(ldc liste) {
	if listeVideOuSingleton(ldc) {
		return
	}
	for inv := false; parcourirEchangerListe(ldc, inv); inv = !inv {
	}
}
