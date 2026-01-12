package main

import (
	"fmt"
	"log"
	"math/rand"
	"slices"
)

// On travaille sur des chiffres dans [0..9]
const valSup = 10

// Taille max des listes pour les tests visuels
const tailleVisu = 7

// Taille max des listes pour les tests automatiques
const tailleMax = 100

// Nombre d'itérations pour les tests automatiques
const nbrMaxIter = 1000

// Une cellule est composée :
// - d'une valeur entière
// - d'une pointeur vers la cellule suivante
type cellule struct {
	val  int
	suiv *cellule
}

// Une liste est simplement un pointeur vers la première cellule (sa tête)
type liste *cellule

// Fonctions de base

func afficher(lsc liste) {
	// On rappelle que les paramètres sont passés par copie : on peut modifier lsc sans impacter la liste dans la fonction appelante
	for ; lsc != nil; lsc = lsc.suiv {
		fmt.Print(lsc.val, " -> ")
	}
	fmt.Println("FIN")
}

func insérerTête(lsc liste, val int) liste {
	cell := new(cellule)
	// Pas de besoin de dissocier ici le cas de la liste vide : si lsc est nil, cell.suiv le deviendra
	cell.val, cell.suiv = val, lsc
	return cell
}

func insérerQueue(lsc liste, val int) liste {
	cell := new(cellule)
	cell.val = val
	// Là on doit bien dissocier le cas de la liste vide : dans ce cas, il faut renvoyer cell en tant que nouvelle liste
	if lsc == nil {
		return cell
	}
	cour := lsc
	// On cherche la dernière cellule, c'est-à-dire celle dont le suiv est nil
	for cour.suiv != nil {
		cour = cour.suiv
	}
	cour.suiv = cell
	return lsc
}

// Version où on dissocie chaque cas particulier
func supprimerPremièreOccurrence(lsc liste, val int) (liste, bool) {
	// Si la liste est vide, rien à supprimer
	if lsc == nil {
		return lsc, false
	}
	// Sinon si la première cellule contient la valeur à supprimer, on renvoie la deuxième cellule
	if lsc.val == val {
		return lsc.suiv, true
	}
	// Cas général : on sait que la première cellule ne contenait pas la valeur à supprimer, donc on peut commencer à tester à partir de la deuxième
	prec := lsc
	for (prec.suiv != nil) && (prec.suiv.val != val) {
		prec = prec.suiv
	}
	// Attention : tester prec.suiv.val serait faux si la liste ne contient pas la valeur à supprimer (prec.suiv vaudrait nil dans ce cas)
	if prec.suiv != nil {
		prec.suiv = prec.suiv.suiv
		return lsc, true
	}
	return lsc, false
}

// Version avec le fictif en tête : plus simple non ?
func supprimerPremièreOccurrenceFictif(lsc liste, val int) (liste, bool) {
	fictif := new(cellule)
	fictif.suiv = lsc
	prec := fictif
	for (prec.suiv != nil) && (prec.suiv.val != val) {
		prec = prec.suiv
	}
	if prec.suiv != nil {
		prec.suiv = prec.suiv.suiv
		return fictif.suiv, true
	}
	return lsc, false
}

func testDeBase() {
	fmt.Println("***** Tests de base *****")
	fmt.Println("--- Insertions en tête ---")
	var lsc liste // en Go on ne peut pas écrire lsc := nil car l'inférence de type ne peut pas déduire le type la variable lsc à partir de nil
	afficher(lsc)
	for range tailleVisu { // Une forme pratique du for range : on veut juste exécuter cette boucle tailleVisu fois
		lsc = insérer(lsc, true)
	}
	fmt.Println("--- Suppression sans fictif ---")
	afficher(lsc)
	for lsc != nil {
		lsc = supprimer(lsc, false)
	}
	supprimerVide(false)
	fmt.Println("--- Insertions en queue ---")
	lsc = nil
	afficher(lsc)
	for range tailleVisu {
		lsc = insérer(lsc, false)
	}
	fmt.Println("--- Suppression avec fictif ---")
	afficher(lsc)
	for lsc != nil {
		lsc = supprimer(lsc, true)
	}
	supprimerVide(true)
	fmt.Println()
}

func insérer(lsc liste, tete bool) liste {
	val := rand.Intn(valSup)
	fmt.Print("+", val, " = ")
	if tete {
		lsc = insérerTête(lsc, val)
	} else {
		lsc = insérerQueue(lsc, val)
	}
	afficher(lsc)
	return lsc
}

func supprimer(lsc liste, fictif bool) liste {
	val := rand.Intn(valSup)
	var supp bool
	if fictif {
		lsc, supp = supprimerPremièreOccurrenceFictif(lsc, val)
	} else {
		lsc, supp = supprimerPremièreOccurrence(lsc, val)
	}
	if supp { // pas de trace si on n'a pas trouvé la valeur pour rendre l'affichage plus clair, on peut en ajouter si on veut être sûr
		fmt.Print("-", val, " = ")
		afficher(lsc)
	}
	return lsc
}

// Vue notre façon de tester, il n'est pas certain qu'on essaie de supprimer dans une liste vide
// On le fait à la main car c'est un test important
func supprimerVide(fictif bool) {
	var lsc liste
	var supp bool
	if fictif {
		lsc, supp = supprimerPremièreOccurrenceFictif(lsc, 1024) // valeur sans importance
	} else {
		lsc, supp = supprimerPremièreOccurrence(lsc, 1024)
	}
	if supp || lsc != nil {
		log.Fatal("suppression dans une liste vide incorrecte")
	}
}

// Inversions et tris

func créer(taille int) liste {
	var lsc liste
	for range taille {
		lsc = insérerTête(lsc, rand.Intn(valSup))
	}
	return lsc
}

// Il est indispensable de renvoyer la nouvelle liste, car lsc est une COPIE du pointeur vers la liste initiale
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

// On veut trier la liste par ordre croissant
func trierMax(lsc liste) liste {
	fictif := new(cellule)
	fictif.suiv = lsc
	var res liste = nil
	for fictif.suiv != nil {
		precMax := fictif
		prec := fictif.suiv
		for prec.suiv != nil {
			if prec.suiv.val > precMax.suiv.val {
				precMax = prec
			}
			prec = prec.suiv
		}
		suiv := precMax.suiv.suiv
		precMax.suiv.suiv = res
		res = precMax.suiv
		precMax.suiv = suiv
	}
	return res
}

// On veut trier la liste par ordre décroissant
func trierIns(lsc liste) liste {
	fictif := new(cellule)
	for lsc != nil {
		suiv := lsc.suiv
		prec := fictif
		for (prec.suiv != nil) && (prec.suiv.val >= lsc.val) {
			prec = prec.suiv
		}
		lsc.suiv = prec.suiv
		prec.suiv = lsc
		lsc = suiv
	}
	return fictif.suiv
}

func testInvEtTris() {
	fmt.Println("***** Tests de l'inversion et des tris *****")
	for taille := 0; taille <= tailleMax; taille++ {
		for cpt := range nbrMaxIter {
			lsc := créer(taille)
			lscOrg := dupliquer(lsc)
			afficherSi(taille <= tailleVisu && cpt == 0, "Liste initiale    : ", lsc)
			tabOrg := listeVersTab(lsc)
			slices.Reverse(tabOrg)
			lsc = inverser(lsc)
			afficherSi(taille <= tailleVisu && cpt == 0, "Liste inversée    : ", lsc)
			tabRes := listeVersTab(lsc)
			if !slices.Equal(tabOrg, tabRes) {
				log.Fatal("la liste résultat n'est pas l'inverse de l'initiale")
			}
			slices.Sort(tabOrg)
			lsc = trierMax(lsc)
			afficherSi(taille <= tailleVisu && cpt == 0, "Liste triée (max) : ", lsc)
			tabRes = listeVersTab(lsc)
			if !slices.Equal(tabOrg, tabRes) {
				log.Fatal("la liste n'est pas triée par ordre croissant")
			}
			slices.Reverse(tabOrg)
			lsc = lscOrg
			lsc = trierIns(lsc)
			afficherSi(taille <= tailleVisu && cpt == 0, "Liste triée (ins) : ", lsc)
			tabRes = listeVersTab(lsc)
			if !slices.Equal(tabOrg, tabRes) {
				log.Fatal("la liste n'est pas triée par ordre décroissant")
			}
			if taille == tailleVisu+1 && cpt == 0 {
				fmt.Print("On continue à tester sans trace jusqu'à ", tailleMax, "...")
			}
		}
	}
	fmt.Println("OK")
	fmt.Println()
}

// Prend une liste en paramètre et renvoie le tableau correspondant
func listeVersTab(lsc liste) []int {
	var tab []int
	for lsc != nil {
		tab = append(tab, lsc.val)
		lsc = lsc.suiv
	}
	return tab
}

// Affiche une liste si elle n'est pas trop longue, ne fait rien sinon
func afficherSi(affiche bool, msg string, lsc liste) {
	if affiche {
		fmt.Print(msg)
		afficher(lsc)
	}
}

// Duplique une liste, c'est-à-dire en crée une nouvelle avec les mêmes valeurs dans l'ordre
func dupliquer(lsc liste) liste {
	fictif := new(cellule)
	for dern := fictif; lsc != nil; lsc, dern = lsc.suiv, dern.suiv {
		dern.suiv = new(cellule)
		dern.suiv.val = lsc.val
	}
	return fictif.suiv
}

// Le pivot et le découpage vus en TD

// Algo du pivot
func pivot(lsc liste) liste {
	pivot := lsc
	fictifInf, fictifSup := new(cellule), new(cellule)
	dernInf, dernSup := fictifInf, fictifSup
	for cour := lsc.suiv; cour != nil; cour = cour.suiv {
		if cour.val <= pivot.val {
			dernInf.suiv, dernInf = cour, cour
		} else {
			dernSup.suiv, dernSup = cour, cour
		}
	}
	dernInf.suiv, pivot.suiv, dernSup.suiv = pivot, fictifSup.suiv, nil
	return fictifInf.suiv
}

// Découpage "au milieu" d'une liste en deux sous-listes
func lièvreTortue(lsc liste) (liste, liste) {
	if lsc == nil {
		return nil, nil
	}
	lievre, tortue := lsc, lsc
	for avanceTortue := false; lievre.suiv != nil; lievre = lievre.suiv {
		if avanceTortue {
			tortue = tortue.suiv
		}
		avanceTortue = !avanceTortue
	}
	lievre = tortue.suiv
	tortue.suiv = nil
	return lsc, lievre
}

func testPivotLièvreTortue() {
	fmt.Println("***** Tests du pivot et du découpage *****")
	for taille := 0; taille <= tailleVisu; taille++ {
		lsc := créer(taille)
		afficherSi(true, "Liste initiale : ", lsc)
		if taille > 0 { // le pivot n'a pas de sens sur une liste vide
			valPivot := lsc.val
			lsc = pivot(lsc)
			afficherSi(true, fmt.Sprint("Pivot (", valPivot, ")      : "), lsc)
		}
		lsc1, lsc2 := lièvreTortue(lsc)
		afficherSi(true, "Liste1         : ", lsc1)
		afficherSi(true, "Liste2         : ", lsc2)
	}
	fmt.Println()
}

func main() {
	testDeBase()
	testInvEtTris()
	testPivotLièvreTortue()
}
