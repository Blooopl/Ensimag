package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func lesMotsDuFichier(nomFichier string) []string {
	var temp []string
	file, err := os.Open(nomFichier)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lecteur := bufio.NewScanner(file)

	for lecteur.Scan() {
		for _, valeur := range strings.FieldsFunc(lecteur.Text(), toutSaufLettre) {
			temp = append(temp, strings.ToLower(valeur))
		}
	}
	return temp
}

func toutSaufLettre(car rune) bool {
	return !unicode.IsLetter(car)
}

func main() {
	nom_fichier := "s19_3_joey.txt"

	mots_fichiers := lesMotsDuFichier(nom_fichier)
	couples_mots := lesCouplesDeMots(mots_fichiers)
	nombre_couples := lesNombresDeCouples(couples_mots)
	génèreGraphe(nom_fichier, nombre_couples)
}

type couplesDeMots struct {
	mot1 string
	mot2 string
}

func lesCouplesDeMots(lesMots []string) []couplesDeMots {
	var liste_couples []couplesDeMots
	var new_couple couplesDeMots
	for cle := 0; cle < len(lesMots)-1; cle++ {
		new_couple.mot1 = lesMots[cle]
		new_couple.mot2 = lesMots[cle+1]

		liste_couples = append(liste_couples, new_couple)
	}

	return liste_couples
}

func lesNombresDeCouples(lesCouples []couplesDeMots) map[string]map[string]int {

	new_map := make(map[string]map[string]int)

	for _, valeur := range lesCouples {
		m1 := valeur.mot1
		m2 := valeur.mot2

		if new_map[m1] == nil {
			map_1 := make(map[string]int)
			map_1[m2] = 0
			new_map[m1] = map_1
		}

		new_map[m1][m2]++
	}

	return new_map
}

func génèreGraphe(nomFichier string, lesNombres map[string]map[string]int) {
	fmt.Println("digraph {")
	defer fmt.Println("}")

	for mot1, couple := range lesNombres {

		for mot2, valeur := range couple {
			fmt.Print(mot1, " -> ", mot2, " [label = ", valeur, "]")
			fmt.Println("")
		}
	}

}
