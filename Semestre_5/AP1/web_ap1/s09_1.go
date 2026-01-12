package main

import (
	"fmt"
	"log"
	"math/rand"
	"slices"
)

const valSup = 10

const tailleMax = 12

// On peut ajuster selon la vitesse de la machine
const tailleMaxAuto = 3000

const nombreItérationsTests = 100

func remplir(tab []int) {
	for idx := 0; idx < len(tab); idx++ {
		tab[idx] = rand.Intn(valSup)
	}
}

// N'utilisez pas la syntaxe "à la Python" tab[2:5] : elle fonctionne différemment en Go
// (et en plus le but de l'exercice, c'est de faire des boucles sur des tableaux !)
func afficherSousTab(tab []int, prem, dern int) {
	fmt.Print("[")
	for idx := prem; idx <= dern; idx++ {
		fmt.Print(" ", tab[idx])
	}
	fmt.Println(" ]")
}

func afficher(tab []int) {
	afficherSousTab(tab, 0, len(tab)-1)
}

func pivot(tab []int, ixPiv int) int {
	// Pas indispensable, mais c'est toujours mieux de vérifier proprement !
	if len(tab) < 1 {
		log.Fatal("Erreur : le tableau doit contenir au moins élément !")
	}
	// On place le pivot dans la case 0 pour se ramener à l'algo vu en TD
	tab[0], tab[ixPiv] = tab[ixPiv], tab[0]
	ixPiv = 0
	prem, dern := 1, len(tab)-1
	for prem <= dern {
		if tab[prem] <= tab[ixPiv] {
			prem++
		} else {
			tab[prem], tab[dern] = tab[dern], tab[prem]
			dern--
		}
	}
	tab[dern], tab[0] = tab[0], tab[dern]
	return dern
}

func drapeau(tab []int, ixPiv int) (int, int) {
	if len(tab) < 1 {
		log.Fatal("Erreur : le tableau doit contenir au moins élément !")
	}
	tab[0], tab[ixPiv] = tab[ixPiv], tab[0]
	ixPiv = 0
	prem, dern := 1, len(tab)-1
	for prem <= dern {
		if tab[prem] < tab[ixPiv] {
			tab[prem], tab[ixPiv] = tab[ixPiv], tab[prem]
			ixPiv++
			prem++
		} else if tab[prem] == tab[ixPiv] {
			prem++
		} else {
			tab[prem], tab[dern] = tab[dern], tab[prem]
			dern--
		}
	}
	return ixPiv, dern
}

func testVisuelPivot() {
	tab := make([]int, tailleMax)
	fmt.Print("Tableau pour le pivot  : ")
	remplir(tab)
	afficher(tab)
	ixPiv := rand.Intn(len(tab))
	fmt.Println(" indice du pivot =", ixPiv, "valeur =", tab[ixPiv])
	ixPiv = pivot(tab, ixPiv)
	fmt.Print("Tableau après le pivot : ")
	afficher(tab)
	fmt.Println(" indice du pivot =", ixPiv, "valeur =", tab[ixPiv])
	fmt.Print(" sous-tableau des <= : ")
	afficherSousTab(tab, 0, ixPiv)
	fmt.Print(" sous-tableau des >  : ")
	afficherSousTab(tab, ixPiv+1, len(tab)-1)
}

func testVisuelDrapeau() {
	tab := make([]int, tailleMax)
	fmt.Print("Tableau pour le drapeau  : ")
	remplir(tab)
	afficher(tab)
	ixPiv := rand.Intn(len(tab))
	fmt.Println(" indice du pivot =", ixPiv, "valeur =", tab[ixPiv])
	prem, dern := drapeau(tab, ixPiv)
	fmt.Print("Tableau après le drapeau : ")
	afficher(tab)
	fmt.Println(" indices du pivot dans [", prem, "...", dern, "]")
	fmt.Print(" sous-tableau des < : ")
	afficherSousTab(tab, 0, prem-1)
	fmt.Print(" sous-tableau des = : ")
	afficherSousTab(tab, prem, dern)
	fmt.Print(" sous-tableau des > : ")
	afficherSousTab(tab, dern+1, len(tab)-1)
}

// Fonctions de tests automatiques

func erreur(tabOrg, tabRes []int, ixPivOrg, ixPivResOuPrem, ixDern int, msg string) {
	fmt.Println("!! Erreur !!")
	fmt.Print(" tableau original : ")
	afficher(tabOrg)
	fmt.Println(" indice original du pivot =", ixPivOrg, "valeur =", tabOrg[ixPivOrg])
	fmt.Print(" tableau résultat : ")
	afficher(tabRes)
	if ixDern == -1 { // pivot
		fmt.Println(" indice résultat du pivot =", ixPivResOuPrem, "valeur =", tabRes[ixPivResOuPrem])
	} else {
		fmt.Println(" indices résultat du pivot = [", ixPivResOuPrem, "..", ixDern, "], valeur =", tabRes[ixPivResOuPrem])
	}
	log.Fatal(msg)
}

func testAutoPivot(tabOrg, tabRes []int, ixPivOrg, ixPivRes int) {
	// On vérifie que les tableaux ont la même taille
	if len(tabOrg) != len(tabRes) {
		erreur(tabOrg, tabRes, ixPivOrg, ixPivRes, -1, "Erreur : les deux tableaux n'ont pas la même taille !")
	}
	// On vérifie que le pivot est au bon endroit
	valPiv := tabOrg[ixPivOrg]
	if tabRes[ixPivRes] != valPiv {
		erreur(tabOrg, tabRes, ixPivOrg, ixPivRes, -1, "Erreur : le pivot est mal-placé !")
	}
	// On compte le nombre d'occurrences de chaque valeur du tableau original...
	var nbrs [valSup]int
	for idx := 0; idx < len(tabOrg); idx++ {
		nbrs[tabOrg[idx]]++
	}
	// et on vérifie qu'on a bien les mêmes nombres dans le tableau résultat
	for idx := 0; idx < len(tabRes); idx++ {
		nbrs[tabRes[idx]]--
	}
	for idx := 0; idx < len(nbrs); idx++ {
		if nbrs[idx] != 0 {
			erreur(tabOrg, tabRes, ixPivOrg, ixPivRes, -1, "Erreur : les contenus des tableaux sont incohérents !")
		}
	}
	// On vérifie que tous les éléments à gauche du pivot sont <= à la valeur du pivot...
	for idx := 0; idx <= ixPivRes; idx++ {
		if tabRes[idx] > valPiv {
			erreur(tabOrg, tabRes, ixPivOrg, ixPivRes, -1, "Erreur : la partie gauche contient une valeur > au pivot !")
		}
	}
	// ... et que ceux à droite sont strictement supérieurs
	for idx := ixPivRes + 1; idx < len(tabRes); idx++ {
		if tabRes[idx] <= valPiv {
			erreur(tabOrg, tabRes, ixPivOrg, ixPivRes, -1, "Erreur : la partie gauche contient une valeur > au pivot !")
		}
	}
}

func testAutoDrapeau(tabOrg, tabRes []int, ixPiv, ixPrem, ixDern int) {
	if len(tabOrg) != len(tabRes) {
		erreur(tabOrg, tabRes, ixPiv, ixPrem, ixDern, "Erreur : les deux tableaux n'ont pas la même taille !")
	}
	var nbrs [valSup]int
	for idx := 0; idx < len(tabOrg); idx++ {
		nbrs[tabOrg[idx]]++
	}
	for idx := 0; idx < len(tabRes); idx++ {
		nbrs[tabRes[idx]]--
	}
	for idx := 0; idx < len(nbrs); idx++ {
		if nbrs[idx] != 0 {
			erreur(tabOrg, tabRes, ixPiv, ixPrem, ixDern, "Erreur : les contenus des tableaux sont incohérents !")
		}
	}
	valPiv := tabOrg[ixPiv]
	for idx := 0; idx < ixPrem; idx++ {
		if tabRes[idx] >= valPiv {
			erreur(tabOrg, tabRes, ixPiv, ixPrem, ixDern, "Erreur : la partie gauche contient une valeur >= au pivot !")
		}
	}
	for idx := ixPrem; idx <= ixDern; idx++ {
		if tabRes[idx] != valPiv {
			erreur(tabOrg, tabRes, ixPiv, ixPrem, ixDern, "Erreur : la partie centrale contient une valeur != du pivot !")
		}
	}
	for idx := ixDern + 1; idx < len(tabRes); idx++ {
		if tabRes[idx] <= valPiv {
			erreur(tabOrg, tabRes, ixPiv, ixPrem, ixDern, "Erreur : la partie gauche contient une valeur <= au pivot !")
		}
	}
}

func main() {
	testVisuelPivot()
	// On test le pivot avec des tableaux de taille 1 à tailleMax inclus...
	for taille := 1; taille <= tailleMaxAuto; taille++ {
		fmt.Print("\rTest du pivot avec ", nombreItérationsTests, " tableaux aléatoires de taille ", taille, " : ")
		// ... et on teste chaque taille avec 100 tableauw remplis aléatoirement !
		for idx := 0; idx < nombreItérationsTests; idx++ {
			tabRes := make([]int, taille)
			remplir(tabRes)
			// slices.Clone() alloue un nouveau tableau (comme make) puis recopie dedans les valeurs de celui passé en paramètre (comme copy)
			tabOrg := slices.Clone(tabRes)
			ixPivOrg := rand.Intn(len(tabRes))
			ixPivRes := pivot(tabRes, ixPivOrg)
			testAutoPivot(tabOrg, tabRes, ixPivOrg, ixPivRes)
		}
		fmt.Print("OK !")
	}
	fmt.Println()
	testVisuelDrapeau()
	for taille := 1; taille <= tailleMaxAuto; taille++ {
		fmt.Print("\rTest du drapeau avec ", nombreItérationsTests, " tableaux aléatoires de taille ", taille, " : ")
		for idx := 0; idx < nombreItérationsTests; idx++ {
			tabRes := make([]int, taille)
			remplir(tabRes)
			tabOrg := slices.Clone(tabRes)
			ixPiv := rand.Intn(len(tabRes))
			ixPrem, ixDern := drapeau(tabRes, ixPiv)
			testAutoDrapeau(tabOrg, tabRes, ixPiv, ixPrem, ixDern)
		}
		fmt.Print("OK !")
	}
	fmt.Println()
}
