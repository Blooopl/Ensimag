package main

import (
	"fmt"
	"math/rand"
)

const valSup = 10

const taille = 12

type cellule struct {
	val  int
	suiv *cellule
}

type liste *cellule

func créer() liste {
	// On crée l'élément fictif dont le suiv est lui-même !
	cell := new(cellule)
	cell.suiv = cell
	return cell
}

func estVide(lct liste) bool {
	return lct.suiv == lct
}

func afficher(lct liste) {
	for cour := lct.suiv; cour != lct; cour = cour.suiv {
		fmt.Print(cour.val, " -> ")
	}
	fmt.Println("FIN")
}

func insérerListeCircTriée(lct liste, val int) {
	cell := new(cellule)
	cell.val = val
	prec := lct
	// On recherche l'endroit où insérer la nouvelle cellule pour préserver l'ordre
	// Au pire, on s'arrêtera sur la dernière cellule de la liste
	for ; (prec.suiv != lct) && (prec.suiv.val <= val); prec = prec.suiv {
	}
	cell.suiv = prec.suiv
	prec.suiv = cell
}

func supprimerListeCircTriée(lct liste, val int) bool {
	prec := lct
	// Même principe
	for ; (prec.suiv != lct) && (prec.suiv.val < val); prec = prec.suiv {
	}
	if (prec.suiv != lct) && (prec.suiv.val == val) { // On a pu s'arrêter car on a trouvé une cellule > val : il faut bien tester la valeur
		prec.suiv = prec.suiv.suiv
		return true
	}
	return false
}

func supprimerDoublonsListeCircTriée(lct liste) {
	for prec := lct; (prec.suiv != lct) && (prec.suiv.suiv != lct); prec = prec.suiv {
		// Tant que la valeur du suiv (prec.suiv.suiv) est égale à la valeur du cour (prec.suiv), on l'enlève de la liste
		for prec.suiv.val == prec.suiv.suiv.val {
			prec.suiv = prec.suiv.suiv
		}
	}
}

func découperListeCircTriée(lct liste) (liste, liste) {
	var lcts, derns [2]liste
	// Bien noter qu'ici, on ne crée que deux fictifs : on ne recopier pas les valeurs significatives de la liste
	lcts[0], lcts[1] = créer(), créer()
	derns[0], derns[1] = lcts[0], lcts[1]
	num := 0
	for prec := lct; prec.suiv != lct; prec.suiv = prec.suiv.suiv {
		derns[num].suiv = prec.suiv
		derns[num] = prec.suiv
		num = (num + 1) % 2
	}
	derns[0].suiv, derns[1].suiv = lcts[0], lcts[1]
	return lcts[0], lcts[1]
}

// Tests

func testIns(lct liste) {
	for range taille {
		val := rand.Intn(valSup)
		fmt.Print("Insertion de la valeur ", val, " : ")
		insérerListeCircTriée(lct, val)
		afficher(lct)
	}
}

func testDoublons() {
	fmt.Println("Suppression des doublons dans une liste vide :")
	lct := créer()
	fmt.Print(" liste initiale : ")
	afficher(lct)
	supprimerDoublonsListeCircTriée(lct)
	fmt.Print(" liste filtrée  : ")
	afficher(lct)
	fmt.Println("Suppression des doublons dans une liste singleton :")
	lct = créer()
	insérerListeCircTriée(lct, rand.Intn(valSup))
	fmt.Print(" liste initiale : ")
	afficher(lct)
	supprimerDoublonsListeCircTriée(lct)
	fmt.Print(" liste filtrée  : ")
	afficher(lct)
	fmt.Println("Suppression des doublons dans une liste à deux éléments distincts :")
	lct = créer()
	insérerListeCircTriée(lct, 1)
	insérerListeCircTriée(lct, 2)
	fmt.Print(" liste initiale : ")
	afficher(lct)
	supprimerDoublonsListeCircTriée(lct)
	fmt.Print(" liste filtrée  : ")
	afficher(lct)
	fmt.Println("Suppression des doublons dans une liste à deux éléments égaux :")
	lct = créer()
	insérerListeCircTriée(lct, 1)
	insérerListeCircTriée(lct, 1)
	fmt.Print(" liste initiale : ")
	afficher(lct)
	supprimerDoublonsListeCircTriée(lct)
	fmt.Print(" liste filtrée  : ")
	afficher(lct)
}

func testDécoupage() {
	var lcts [2]liste
	fmt.Println("Découpage sur une liste vide :")
	lct := créer()
	lcts[0], lcts[1] = découperListeCircTriée(lct)
	fmt.Print(" liste initiale : ")
	afficher(lct)
	fmt.Print(" liste initiale : ")
	afficher(lct)
	fmt.Print(" sous-liste 1   : ")
	afficher(lcts[0])
	fmt.Print(" sous-liste 2   : ")
	afficher(lcts[1])
	fmt.Println("Découpage sur une liste singleton :")
	lct = créer()
	insérerListeCircTriée(lct, rand.Intn(valSup))
	fmt.Print(" liste initiale : ")
	afficher(lct)
	lcts[0], lcts[1] = découperListeCircTriée(lct)
	fmt.Print(" liste initiale : ")
	afficher(lct)
	fmt.Print(" sous-liste 1   : ")
	afficher(lcts[0])
	fmt.Print(" sous-liste 2   : ")
	afficher(lcts[1])
	fmt.Println("Découpage sur une liste à deux éléments :")
	lct = créer()
	insérerListeCircTriée(lct, rand.Intn(valSup))
	insérerListeCircTriée(lct, rand.Intn(valSup))
	fmt.Print(" liste initiale : ")
	afficher(lct)
	lcts[0], lcts[1] = découperListeCircTriée(lct)
	fmt.Print(" liste initiale : ")
	afficher(lct)
	fmt.Print(" sous-liste 1   : ")
	afficher(lcts[0])
	fmt.Print(" sous-liste 2   : ")
	afficher(lcts[1])
}

func testSupp() {
	fmt.Println("Suppression dans une liste vide :")
	lct := créer()
	fmt.Print(" liste initiale : ")
	afficher(lct)
	val := rand.Intn(valSup)
	fmt.Print(" suppression de la valeur ", val, " : ")
	fmt.Println(supprimerListeCircTriée(lct, val))
	fmt.Print(" liste filtrée  : ")
	afficher(lct)
	fmt.Println("Suppression d'une valeur présente dans une liste singleton :")
	lct = créer()
	val = rand.Intn(valSup)
	insérerListeCircTriée(lct, val)
	fmt.Print(" liste initiale : ")
	afficher(lct)
	fmt.Print(" suppression de la valeur ", val, " : ")
	fmt.Println(supprimerListeCircTriée(lct, val))
	fmt.Print(" liste filtrée  : ")
	afficher(lct)
	fmt.Println("Suppression d'une valeur absente dans une liste singleton :")
	lct = créer()
	val = rand.Intn(valSup)
	insérerListeCircTriée(lct, val)
	fmt.Print(" liste initiale : ")
	afficher(lct)
	fmt.Print(" suppression de la valeur ", (val+1)%valSup, " : ")
	fmt.Println(supprimerListeCircTriée(lct, (val+1)%valSup))
	fmt.Print(" liste filtrée : ")
	afficher(lct)
	fmt.Println("Suppression d'une valeur présente dans une liste à deux éléments :")
	lct = créer()
	val = rand.Intn(valSup)
	insérerListeCircTriée(lct, val)
	insérerListeCircTriée(lct, (val+1)%valSup)
	fmt.Print(" liste initiale : ")
	afficher(lct)
	fmt.Print(" suppression de la valeur ", val, " : ")
	fmt.Println(supprimerListeCircTriée(lct, val))
	fmt.Print(" liste filtrée  : ")
	afficher(lct)
	fmt.Println("Suppression d'une valeur absente dans une liste à deux éléments :")
	lct = créer()
	val = rand.Intn(valSup)
	insérerListeCircTriée(lct, val)
	insérerListeCircTriée(lct, (val+1)%valSup)
	fmt.Print(" liste initiale : ")
	afficher(lct)
	fmt.Print(" suppression de la valeur ", (val+2)%valSup, " : ")
	fmt.Println(supprimerListeCircTriée(lct, (val+2)%valSup))
	fmt.Print(" liste filtrée  : ")
	afficher(lct)
}

func testCasGénéraux() {
	lct := créer()
	fmt.Print(" liste initiale : ")
	afficher(lct)
	testIns(lct)
	fmt.Println("Suppression des doublons :")
	fmt.Print(" liste initiale : ")
	afficher(lct)
	supprimerDoublonsListeCircTriée(lct)
	fmt.Print(" liste filtrée  : ")
	afficher(lct)
	fmt.Println("Découpage en sous-listes :")
	fmt.Print(" liste initiale : ")
	afficher(lct)
	var lcts [2]liste
	lcts[0], lcts[1] = découperListeCircTriée(lct)
	fmt.Print(" liste initiale : ")
	afficher(lct)
	fmt.Print(" sous-liste 1   : ")
	afficher(lcts[0])
	fmt.Print(" sous-liste 2   : ")
	afficher(lcts[1])
	for num := 0; num < 2; num++ {
		for !estVide(lcts[num]) {
			val := rand.Intn(valSup)
			if supprimerListeCircTriée(lcts[num], val) {
				fmt.Print("Suppression de la valeur ", val, " : ")
				afficher(lcts[num])
			}
		}
	}
}

func main() {
	// Ici on a fait des fonctions spécifiques pour tester les cas particuliers (liste vide, singleton, etc.)
	testDoublons()
	testDécoupage()
	testSupp()
	testCasGénéraux()
}
