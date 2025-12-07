package main

import (
	"fmt"
	"math/rand"
)

const valSup = 10

type cellule struct {
	val  int
	suiv *cellule
}

type liste *cellule

func main() {
	var lst liste
	lst = insérerTête(lst, 4)
	lst = insérerTête(lst, 2)
	lst = insérerTête(lst, 1)
	lst = insérerTête(lst, 8)
	lst = insérerTête(lst, 9)
	afficher(lst)
	lst = trierIns(lst)
	afficher(lst)
}

func afficher(lsc liste) {
	if lsc != nil {
		pointeur := lsc
		for pointeur.suiv != nil {
			fmt.Print(pointeur.val)
			fmt.Print(" -> ")
			pointeur = pointeur.suiv
		}
		fmt.Print(pointeur.val)
		fmt.Print(" -> ")
		fmt.Print(" FIN ")
	}
	fmt.Println("")
}

func insérerTête(lsc liste, val int) liste {
	var cell cellule
	cell.val = val
	if lsc == nil {
		lsc = &cell
	} else {
		cell.suiv = lsc
		lsc = &cell
	}

	return lsc
}

func insérerQueue(lsc liste, val int) liste {
	var cell cellule
	cell.val = val
	if lsc == nil {
		lsc = &cell
	} else {
		pointeur := lsc
		for pointeur.suiv != nil {
			pointeur = pointeur.suiv
		}
		pointeur.suiv = &cell
	}

	return lsc
}

func supprimerPremièreOccurence(lsc liste, val int) (liste, bool) {

	if lsc == nil {
		return lsc, false
	} else {
		pointeur := lsc
		if pointeur.val == val {
			lsc = pointeur.suiv
			return lsc, true
		}

		for pointeur.suiv != nil {
			if pointeur.suiv.val == val {
				pointeur.suiv = pointeur.suiv.suiv
				return lsc, true
			}
			pointeur = pointeur.suiv
		}
	}
	return lsc, false
}

func créer(taille int) liste {
	var lsc liste
	for i := 1; i <= taille; i++ {
		lsc = insérerTête(lsc, rand.Intn(valSup))
	}
	return lsc
}

func inverser(lsc liste) liste {
	var res liste

	for lsc != nil {
		suiv := lsc.suiv
		lsc.suiv = res
		res = lsc
		lsc = suiv
	}

	return res
}

func trierMax(lsc liste) liste {
	var res liste

	for lsc != nil {
		var prec_max *cellule
		max := lsc
		prec := lsc

		for prec.suiv != nil {
			if prec.suiv.val > max.val {
				max = prec.suiv
				prec_max = prec
			}
			prec = prec.suiv
		}

		if prec_max == nil {
			lsc = lsc.suiv
		} else {
			prec_max.suiv = max.suiv
		}

		max.suiv = res
		res = max
	}
	return res
}

func trierIns(lsc liste) liste {
	//Prendre le premier élements
	//modifer lsc
	//Insérer dans res
	var res liste
	var prems *cellule

	for lsc != nil {

		prems = lsc.suiv

		res = Insérer(res, prems)

		lsc = lsc.suiv
	}
	return res
}

func Insérer(res liste, pointeur *cellule) liste {
	if res == nil {
		res = pointeur
	} else {
		temp := res
		for temp != nil && temp.val > pointeur.val {
			temp = temp.suiv
		}
		if temp == nil {
			temp = pointeur
		} else {
			pointeur.suiv = temp.suiv
			temp = pointeur.suiv
		}
	}
	return res
}
