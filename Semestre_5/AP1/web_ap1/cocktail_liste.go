package main

import "fmt"

// Cellule doublement chaînée
type cellule struct {
	val  int
	prec *cellule
	suiv *cellule
}

// Liste doublement chaînée avec pointeurs de tête et de queue
type liste struct {
	tete  *cellule
	queue *cellule
}

// Alloue les fictifs en tête et en queue et les chaine l'un à l'autre
func initListe() liste {
	var ldc = liste{new(cellule), new(cellule)}
	ldc.tete.prec, ldc.tete.suiv = nil, ldc.queue
	ldc.queue.prec, ldc.queue.suiv = ldc.tete, nil
	return ldc
}

// Construit une liste à partir des valeurs d'un tableau, par insertion en queue pour préserver l'ordre
func remplirListe(ldc liste, tab []int) {
	for idx := range len(tab) {
		cell := new(cellule)
		cell.val, cell.prec, cell.suiv = tab[idx], ldc.queue.prec, ldc.queue
		ldc.queue.prec.suiv = cell
		ldc.queue.prec = cell
	}
}

// Affiche une liste de la tête à la queue
func afficherListe(ldc liste) {
	fmt.Print("T <-> ")
	for cour := ldc.tete.suiv; cour != ldc.queue; cour = cour.suiv {
		fmt.Print(cour.val, " <-> ")
	}
	fmt.Println("Q")
}

// Echange (par modification de chaînage) les cellules cell et cell.suiv
func echangerCellules(cell *cellule) *cellule {
	// On définit des pointeurs vers les 4 cellules impactées
	prec := cell.prec
	cour := cell
	suiv := cell.suiv
	suiv_suiv := suiv.suiv
	// On modifie les 6 chaînages concernés
	prec.suiv = suiv
	cour.prec = suiv
	cour.suiv = suiv_suiv
	suiv.prec = prec
	suiv.suiv = cour
	suiv_suiv.prec = cour
	// On renvoie la "nouvelle" cell, c'est-à-dire la cell.suiv initiale
	return suiv
}

func listeVideOuSingleton(ldc liste) bool {
	// supprimer cette ligne quand vous aurez implanté la fonction !
	return false
}

func parcourirEchangerListe(ldc liste, inv bool) bool {
	// supprimer cette ligne quand vous aurez implanté la fonction !
	return false
}

func trierListe(ldc liste) {
}
