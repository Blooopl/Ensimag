package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func initialiser(taille int) []bool {

	lesPremiers := make([]bool, taille)
	for i := 0; i < taille; i++ {
		lesPremiers[i] = true
	}
	lesPremiers[0] = false

	fmt.Println(lesPremiers)

	return lesPremiers
}

func filtrer(lesPremiers []bool, prem int) {
	valMax := len(lesPremiers)
	for i := 2 * prem; i <= valMax; i = i + prem {
		lesPremiers[i-1] = false
	}
}

func afficher(lesPremiers []bool, valMax int) {
	fmt.Print("Les nombres premiers inférieurs ou égaux à 10 sont : ")

	for i := 1; i < valMax; i++ {
		if lesPremiers[i] == true {
			fmt.Print(" ")
			fmt.Print(i + 1)
			fmt.Print(" ")
		}
	}
}

func main() {
	fmt.Print("Jusqu'à quelle valeur voulez vous les nombres premiers ?")

	lecteur := bufio.NewScanner(os.Stdin)
	lecteur.Scan()
	entree := lecteur.Text()

	taille, err := strconv.Atoi(entree)

	if err != nil {
		log.Fatal(err)
	}

	if taille < 2 {
		erreur := errors.New("La taille demandée est trop petite")
		log.Fatal(erreur)
	}

	lesPremiers := initialiser(taille)

	for test := 2; test <= int(math.Floor(math.Sqrt(float64(taille))))+1; test++ {
		if lesPremiers[test-1] == true {
			filtrer(lesPremiers, test)
		}
	}

	afficher(lesPremiers, len(lesPremiers))
}
