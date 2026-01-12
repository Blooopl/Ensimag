package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"slices"
	"strings"
)

// On a implanté des fonctions similaires plusieurs fois en TP
func lirePaquetsMots(nomFichier string) map[int][]string {
	// Toujours penser à gérer les erreurs
	fich, err := os.Open(nomFichier)
	if err != nil {
		log.Fatal(err)
	}
	// Et la fermeture du fichier
	defer fich.Close()
	lecteur := bufio.NewScanner(fich)
	dico := make(map[int][]string)
	numPaquet := 0
	// Ce tableau dynamique nous sert à stocker les lignes du paquet courant
	lignesDuPaquet := make([]string, 0)
	for lecteur.Scan() {
		ligne := strings.TrimSpace(lecteur.Text())
		if len(ligne) > 0 {
			// C'est une vraie ligne, on l'ajoute au paquet
			lignesDuPaquet = append(lignesDuPaquet, ligne)
		} else {
			// C'est une ligne vide, c'est-à-dire un séparateur de paquets :
			// 1) on ajoute le paquet au dico
			dico[numPaquet] = lignesDuPaquet
			// 2) on passe au paquet suivant (note : on disait bien dans l'énoncé que les paquets sont séparés par UNE ligne vide, jamais deux ou plus à la suite)
			numPaquet++
			// 3) on ré-initialise le tableau qui stocke les lignes du paquet : le plus simple est d'en allouer un nouveau et laisser le GC libérer l'ancien
			lignesDuPaquet = make([]string, 0)
		}
	}
	return dico
}

func genererPoemes3Vers(paquet0, paquet1, paquet2 []string) [][]string {
	poemes := make([][]string, 0)
	for _, mot0 := range paquet0 {
		for _, mot1 := range paquet1 {
			for _, mot2 := range paquet2 {
				poemes = append(poemes, []string{mot0, mot1, mot2})
			}
		}
	}
	return poemes
}

// C'était la question dure de l'examen : pour bien comprendre ce qu'on fait, il faut vraiment prendre un exemple et le dérouler
func plusUn(cour []int, indicesMax []int) []int {
	suiv := make([]int, len(indicesMax))
	copy(suiv, cour)
	for idx := len(suiv) - 1; idx >= 0; idx-- {
		if suiv[idx] < indicesMax[idx] {
			suiv[idx]++
			return suiv
		}
		for i := idx; i < len(suiv); i++ {
			suiv[i] = 0
		}
	}
	return nil
}

func genererIndices(indicesMax []int) [][]int {
	res := make([][]int, 0)
	res = append(res, make([]int, len(indicesMax)))
	for {
		suiv := plusUn(res[len(res)-1], indicesMax)
		if suiv == nil {
			break
		}
		res = append(res, suiv)
	}
	return res
}

func tousLesPoemes(paquets map[int][]string) [][]string {
	var res [][]string
	var indicesMax []int
	for idx := range len(paquets) {
		indicesMax = append(indicesMax, len(paquets[idx])-1)
	}
	for _, indices := range genererIndices(indicesMax) {
		poeme := make([]string, 0)
		for idx, indice := range indices {
			poeme = append(poeme, paquets[idx][indice])
		}
		res = append(res, poeme)
	}
	return res
}

func main() {
	var nomFichier string
	switch {
	case len(os.Args) == 1:
		nomFichier = "poemes1.txt"
	case len(os.Args) == 2:
		nomFichier = os.Args[1]
	default:
		log.Fatal("usage : go run", path.Base(os.Args[0])+".go", "[ nom du fichier .txt ]")
	}
	fmt.Println("On travaille avec le fichier", nomFichier)
	fmt.Println()

	fmt.Println("*** Question 1 ***")
	paquetsDeMots := lirePaquetsMots(nomFichier)
	fmt.Println("=> les paquets de mots :")
	// Version pour ceux qui ont Go >= 1.23
	// for _, numPaquet := range slices.Sorted(maps.Keys(paquetsDeMots)) {
	// 	fmt.Println("Paquet", numPaquet, ":")
	// 	for numVers, vers := range paquetsDeMots[numPaquet] {
	// 		fmt.Printf("- vers %d : \"%s\"\n", numVers, vers)
	// 	}
	// }
	// Version qui marche sur les machines de l'école
	lesClés := make([]int, 0)
	for clé := range paquetsDeMots {
		lesClés = append(lesClés, clé)
	}
	slices.Sort(lesClés)
	for _, numPaquet := range lesClés {
		fmt.Println("Paquet", numPaquet, ":")
		for numVers, vers := range paquetsDeMots[numPaquet] {
			fmt.Printf("- vers %d : \"%s\"\n", numVers, vers)
		}
	}
	fmt.Println()

	fmt.Println("*** Question 2 ***")
	poemesDe3Vers := genererPoemes3Vers(paquetsDeMots[0], paquetsDeMots[1], paquetsDeMots[2])
	fmt.Println("=> tous les poèmes de trois vers :")
	for _, poeme := range poemesDe3Vers {
		fmt.Println("-", strings.Join(poeme, " "))
	}
	fmt.Println()

	fmt.Println("*** Question 3 ***")
	indicesMax := []int{2, 1, 0}
	fmt.Println("=> la séquence d'indices inférieurs à", indicesMax, ":")
	for _, indices := range genererIndices(indicesMax) {
		fmt.Println(indices)
	}
	fmt.Println()

	fmt.Println("*** Question 4 ***")
	fmt.Println("=> tous les poèmes :")
	for _, poeme := range tousLesPoemes(paquetsDeMots) {
		fmt.Println("-", strings.Join(poeme, " "))
	}
	fmt.Println()
}
