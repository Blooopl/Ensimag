package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
)

func devinette() {
	const max = 10
	myst := rand.Intn(max) + 1
	fmt.Println("J'ai tiré aléatoirement un entier entre 1 et", max, "inclus, il faut le deviner !")
	lecteur := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Entrez un entier entre 1 et ", max, " inclus : ")
		lecteur.Scan()
		val, err := strconv.Atoi(lecteur.Text())
		if err != nil {
			log.Fatal(err)
		}
		if val < myst {
			fmt.Println("Trop petit, essayez encore une fois !")
		} else if val > myst {
			fmt.Println("Trop grand, essayez encore une fois !")
		} else {
			fmt.Println("Bravo, c'était ça !")
			break
		}
	}
	fmt.Println()
}

const nbrIter = 100_000_000.0

func monteCarlo() {
	fmt.Println("Je vais travailler en", nbrIter, "itérations...")
	fmt.Println("Valeur de référence    :", math.Pi)
	cpt := 0.0
	for idx := 0; idx < nbrIter; idx++ {
		x := rand.Float64()
		y := rand.Float64()
		if x*x+y*y <= 1.0 {
			cpt += 1.0
		}
	}
	fmt.Println("Approximation calculée :", (4.0*cpt)/nbrIter)
}

func main() {
	devinette()
	monteCarlo()
}
