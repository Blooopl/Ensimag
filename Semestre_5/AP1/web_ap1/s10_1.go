package main

import (
	"fmt"
	"log"
	"math/rand"
	"slices"
	"time"
)

const valSup = 10

const tailleMax = 6

const tailleMaxAuto = 1000

const tailleMinPerfs = 100_000

const tailleMaxPerfs = 400_000

const nombreItérationsTest = 100

func remplir(tab []int) {
	for idx := 0; idx < len(tab); idx++ {
		tab[idx] = rand.Intn(valSup)
	}
}

// Meilleur cas : tab est déjà trié par ordre croissant : O(n)
// Pire cas : tab est déjà trié par ordre décroissant : O(n^2)
func trierNain(tab []int) {
	for idx := 0; idx < len(tab)-1; {
		if tab[idx] > tab[idx+1] {
			tab[idx], tab[idx+1] = tab[idx+1], tab[idx]
			if idx > 0 {
				idx--
			}
		} else {
			idx++
		}
	}
}

// Meilleur et pire cas : O(n^2)
func trierMax(tab []int) {
	for idx := 0; idx < len(tab)-1; idx++ {
		ixMax := idx
		for idxRech := idx + 1; idxRech < len(tab); idxRech++ {
			if tab[idxRech] > tab[ixMax] {
				ixMax = idxRech
			}
		}
		tab[idx], tab[ixMax] = tab[ixMax], tab[idx]
	}
}

// Meilleur cas : tab est déjà trié par ordre croissant : O(n)
// Pire cas : tab est déjà trié par ordre décroissant : O(n^2)
func trierIns(tab []int) {
	for idx := 1; idx < len(tab); idx++ {
		val := tab[idx]
		var idxIns int
		for idxIns = idx - 1; (idxIns >= 0) && (tab[idxIns] > val); idxIns-- {
			tab[idxIns+1] = tab[idxIns]
		}
		tab[idxIns+1] = val
	}
}

func testVisuel() {
	for taille := 0; taille <= tailleMax; taille++ {
		fmt.Println("***** Tableau de taille", taille, "*****")
		tab := make([]int, taille)
		fmt.Print("Tableau initial            : ")
		remplir(tab)
		fmt.Println(tab)
		fmt.Print("Tableau trié par la nain   : ")
		trierNain(tab)
		fmt.Println(tab)
		fmt.Print("Tableau initial            : ")
		remplir(tab)
		fmt.Println(tab)
		fmt.Print("Tableau trié par sélection : ")
		trierMax(tab)
		fmt.Println(tab)
		fmt.Print("Tableau initial            : ")
		remplir(tab)
		fmt.Println(tab)
		fmt.Print("Tableau trié par insertion : ")
		trierIns(tab)
		fmt.Println(tab)
		fmt.Println()
	}
}

func testAuto() {
	for taille := 0; taille <= tailleMaxAuto; taille++ {
		fmt.Print("\rTests avec un tableau de taille ", taille, " : ")
		for cpt := 0; cpt < nombreItérationsTest; cpt++ {
			tabOrg := make([]int, taille)
			remplir(tabOrg)
			tabGo := slices.Clone(tabOrg)
			slices.Sort(tabGo)
			tabRes := slices.Clone(tabOrg)
			trierNain(tabRes)
			if !slices.Equal(tabRes, tabGo) {
				log.Fatal("erreur dans le tri du nain")
			}
			copy(tabRes, tabOrg)
			trierMax(tabRes)
			slices.Reverse(tabRes)
			if !slices.Equal(tabRes, tabGo) {
				log.Fatal("erreur dans le tri par sélection")
			}
			copy(tabRes, tabOrg)
			trierIns(tabRes)
			if !slices.Equal(tabRes, tabGo) {
				log.Fatal("erreur dans le tri par insertion")
			}

		}
		fmt.Print("OK !")
	}
	fmt.Println()
	fmt.Println()
}

// Récupère l'heure courante sous la forme d'un entier long de 64 bits
func heureCourante() int64 {
	return time.Now().UnixMilli()
}

// Soustrait deux temps en les convertissant en float64 pour ne pas perdre en précision
func diffTemps(deb, fin int64) float64 {
	return (float64(fin) - float64(deb)) / 1000.0
}

func testPerfs() {
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
		fmt.Print("- tableau trié par insertion : ")
		deb = heureCourante()
		trierIns(tabOrg)
		fin = heureCourante()
		fmt.Println(diffTemps(deb, fin), "secondes")
		copy(tabOrg, tabSav)
		fmt.Print("- tableau trié par Go        : ")
		deb = heureCourante()
		slices.Sort(tabOrg)
		fin = heureCourante()
		fmt.Println(diffTemps(deb, fin), "secondes")
	}
}

func main() {
	testVisuel()
	testAuto()
	testPerfs()
}
