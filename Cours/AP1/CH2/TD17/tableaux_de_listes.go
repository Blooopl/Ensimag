package main

import "fmt"

var taille = 4

type tabListes struct {
	nbelm int
	tab   [4]liste
}

type liste struct {
	nbelm int
	tete  *cellule
}

type cellule struct {
	val  int
	suiv *cellule
}

func créerTabListe() *tabListes {
	var tabL *tabListes = new(tabListes)
	for i := 0; i < taille; i++ {
		tabL.tab[i].tete = new(cellule)
	}
	return tabL
}

func afficherTabListes(tabL *tabListes) {
	fmt.Println("Nombre total d'élements = ", tabL.nbelm)

	for i := 0; i < taille; i++ {
		afficherListe(tabL.tab[i])
	}
}

func afficherListe(lst liste) {
	fmt.Print("Nombre d'élements = ", lst.nbelm, " valeurs = ")
	tete := lst.tete
	tete = tete.suiv
	for tete != nil {
		fmt.Print(tete.val, " -> ")
		tete = tete.suiv
	}
	fmt.Println("FIN")
}

func insérerTabListes(tabL *tabListes, val int) {
	idx := hash(val)
	insérerListe(tabL.tab[idx], val)
	tabL.tab[idx].nbelm++
	tabL.nbelm++
}

func insérerListe(lst liste, val int) liste {
	tete := lst.tete

	for tete.suiv != nil {
		tete = tete.suiv
	}

	cell := new(cellule)
	cell.val = val

	tete.suiv = cell

	return lst

}

func hash(val int) int {
	return val % taille
}

func main() {
	tabL := créerTabListe()
	insérerTabListes(tabL, 1)
	insérerTabListes(tabL, 5)
	insérerTabListes(tabL, 9)
	insérerTabListes(tabL, 2)
	insérerTabListes(tabL, 8)
	insérerTabListes(tabL, 4)
	insérerTabListes(tabL, 7)
	afficherTabListes(tabL)
	supprimerTabListes(tabL, 4)
	afficherTabListes(tabL)
}

func supprimerTabListes(tabL *tabListes, val int) bool {

	idx := hash(val)
	res := supprimerListe(tabL.tab[idx], val)
	if res == true {
		tabL.nbelm--
		tabL.tab[idx].nbelm--
	}

	return res
}

func supprimerListe(lst liste, val int) bool {

	tete := lst.tete
	prec := tete
	for tete != nil {
		tete = tete.suiv
		if tete.val == val {
			if tete.suiv != nil {
				prec.suiv = tete.suiv
			} else {
				prec.suiv = nil
			}
			return true
		}
		prec = prec.suiv
	}

	return false
}
