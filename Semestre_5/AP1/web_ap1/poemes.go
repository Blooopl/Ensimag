package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"slices"
	"strings"
)

func lirePaquetsMots(nomFichier string) map[int][]string {
	// supprimer cette ligne quand vous aurez implanté la fonction !
	return nil
}

func genererPoemes3Vers(paquet0, paquet1, paquet2 []string) [][]string {
	// supprimer cette ligne quand vous aurez implanté la fonction !
	return nil
}

func genererIndices(indicesMax []int) [][]int {
	// supprimer cette ligne quand vous aurez implanté la fonction !
	return nil
}

func tousLesPoemes(paquets map[int][]string) [][]string {
	// supprimer cette ligne quand vous aurez implanté la fonction !
	return nil
}

// Programme principal qui fait les tests attendus
func main() {
	// Si on ne précise pas le nom du fichier sur la ligne de commande, c'est "poemes1.txt" par défaut
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
