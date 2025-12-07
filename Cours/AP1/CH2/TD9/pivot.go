package main

import (
	"fmt"
	"math/rand"
)

var tailleMax = 30000
var valSup = 25

func main() {
	for taille := 10; taille < tailleMax; taille++ {

		for test := 0; test < 100; test++ {
			tab_1 := make([]int, taille)
			remplir(tab_1)

			tab_2 := make([]int, taille)
			copy(tab_2, tab_1)

			ixPivOrg := rand.Intn(taille)

			ixPivRes := pivot(tab_2, ixPivOrg)

			fmt.Println(testAutoPivot(tab_1, tab_2, ixPivOrg, ixPivRes))

		}
	}

}

func remplir(tab []int) {
	for i := 0; i < len(tab); i++ {
		tab[i] = rand.Intn(valSup)
	}
}

func afficherSousTab(tab []int, prem, dern int) {

	fmt.Print("[")
	for i := prem; i <= dern; i++ {
		fmt.Print(tab[i])
		fmt.Print(" ")
	}
	fmt.Print("]")
	fmt.Print("\n")
}
func afficher(tab []int) {
	afficherSousTab(tab, 0, len(tab)-1)
}

func pivot(tab []int, ixPiv int) int {

	tab[ixPiv], tab[0] = tab[0], tab[ixPiv]

	ixPiv = 0

	premier := 0
	dernier := len(tab) - 1

	for premier < dernier {
		if tab[ixPiv+1] <= tab[ixPiv] {
			tab[premier], tab[ixPiv+1] = tab[ixPiv+1], tab[premier]
			premier++
			ixPiv++
		} else {
			tab[dernier], tab[ixPiv+1] = tab[ixPiv+1], tab[dernier]
			dernier--
		}
	}
	return ixPiv
}

func drapeau(tab []int, ixPiv int) (int, int) {
	tab[ixPiv], tab[0] = tab[0], tab[ixPiv]

	ixPiv = 0

	premier := 0
	dernier := len(tab) - 1

	for premier <= dernier {
		if tab[premier] <= tab[ixPiv] {
			if tab[premier] == tab[ixPiv] {
				premier++
			} else {
				tab[premier], tab[ixPiv] = tab[ixPiv], tab[premier]
				premier++
				ixPiv++
			}
		} else {
			tab[premier], tab[dernier] = tab[dernier], tab[premier]
			dernier--
		}
	}
	return ixPiv, dernier
}

func dico(tab []int) []int {
	dico := make([]int, valSup)

	for i := 0; i < len(tab); i++ {
		dico[tab[i]]++
	}

	return dico
}

func egale_tab(tab1, tab2 []int) bool {

	for i := 0; i < len(tab1); i++ {
		if tab1[i] != tab2[i] {
			return false
		}
	}
	return true
}
func testAutoPivot(tabOrg, tabRes []int, ixPivOrg, ixPivRes int) bool {
	if len(tabOrg) != len(tabRes) {
		return false
	}

	if tabOrg[ixPivOrg] != tabRes[ixPivRes] {
		return false
	}
	if egale_tab(dico(tabOrg), dico(tabRes)) == false {
		return false
	}

	pivot := tabOrg[ixPivOrg]
	for i := 0; i < ixPivRes; i++ {
		if tabRes[i] > pivot {
			return false
		}
	}
	for i := ixPivRes + 1; i < len(tabRes); i++ {
		if tabRes[i] < pivot {
			return false
		}
	}

	return true
}
