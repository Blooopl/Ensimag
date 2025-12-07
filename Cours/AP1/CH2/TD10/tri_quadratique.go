package main

import (
	"fmt"
	"math/rand"
	"slices"
	"time"
)

func main() {
	testPerfs()
}

var tailleMax = 30000
var valSup = 25

func remplir(tab []int) {
	for i := 0; i < len(tab); i++ {
		tab[i] = rand.Intn(valSup)
	}
}

func trierNain(tab []int) {

	id_nain := 0
	n := len(tab)
	for id_nain != n-1 {

		if tab[id_nain+1] >= tab[id_nain] {
			id_nain++
		} else {
			tab[id_nain], tab[id_nain+1] = tab[id_nain+1], tab[id_nain]
			if id_nain != 0 {
				id_nain--
			}
		}
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

func maxi_id(i int, tab []int) int {

	n := len(tab)
	maxi := tab[i]
	id_max := i
	for j := i; j < n; j++ {
		if tab[j] > maxi {
			maxi = tab[j]
			id_max = j
		}

	}

	return id_max
}

func trierMax(tab []int) {

	n := len(tab)
	for i := 0; i < n; i++ {
		id_max := maxi_id(i, tab)
		tab[i], tab[id_max] = tab[id_max], tab[i]
	}
}

func trierIns(tab []int) {

	n := len(tab)

	for i := 1; i < n; i++ {
		for j := i; j > 0 && tab[j-1] > tab[j]; j-- {
			tab[j-1], tab[j] = tab[j], tab[j-1]
		}
	}

}

func testAuto() bool {

	tailleMaxAuto := 1000
	nombreItérationsTest := 100

	for taille := 1; taille <= tailleMaxAuto; taille++ {
		for cpt := 0; cpt < nombreItérationsTest; cpt++ {

			tab_1 := make([]int, taille)
			tab_2 := slices.Clone(tab_1)
			tab_3 := slices.Clone(tab_1)

			slices.Sort(tab_2)

			trierNain(tab_1)
			test_1 := slices.Equal(tab_2, tab_1)
			copy(tab_1, tab_3)

			trierMax(tab_1)
			slices.Reverse(tab_1)
			test_2 := slices.Equal(tab_2, tab_1)
			copy(tab_1, tab_3)

			trierMax(tab_1)
			test_3 := slices.Equal(tab_2, tab_1)

			if !test_1 || !test_2 || !test_3 {
				return false
			}
		}
	}
	return true
}

func heureCourante() int64 {
	return time.Now().UnixMilli()
}

func diffTemps(deb, fin int64) float64 {
	return (float64(fin) - float64(deb)) / 1000.0
}

func testPerfs() {
	tailleMinPerfs := 100000
	tailleMaxPerfs := 400000

	var deb, fin int64
	for taille := tailleMinPerfs; taille <= tailleMaxPerfs; taille *= 2 {
		fmt.Println("Tris d'un tableau de", taille, "éléments")
		tabOrg := make([]int, taille)
		remplir(tabOrg)
		// On sauvegarde le tableau avant de le trier pour le réutiliser dans les tris suivants
		tabSav := slices.Clone(tabOrg)
		fmt.Print("- tableau trié par le nain   : ")
		deb = heureCourante()
		trierNain(tabOrg)
		fin = heureCourante()
		fmt.Println(diffTemps(deb, fin), "secondes")
		// On restaure le tableau initial pour être sûr qu'on donne bien le même tableau à chaque tri
		copy(tabOrg, tabSav)
		fmt.Print("- tableau trié par sélection : ")
		deb = heureCourante()
		trierMax(tabOrg)
		fin = heureCourante()
		fmt.Println(diffTemps(deb, fin), "secondes")

		copy(tabOrg, tabSav)
		fmt.Print("- tableau trié par sélection : ")
		deb = heureCourante()
		trierIns(tabOrg)
		fin = heureCourante()
		fmt.Println(diffTemps(deb, fin), "secondes")
		// Et ainsi de suite pour les autres tris

		copy(tabOrg, tabSav)
		fmt.Print("- tableau trié par Go : ")
		deb = heureCourante()
		slices.Sort(tabOrg)
		fin = heureCourante()
		fmt.Println(diffTemps(deb, fin), "secondes")

	}
}
