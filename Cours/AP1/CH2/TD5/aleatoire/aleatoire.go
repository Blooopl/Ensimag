package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand/v2"
	"os"
	"strconv"
)

func devinette() {
	fmt.Println("J'ai tiré aléatoirement un entier en 1 et 10 inclus, il faut le deviner")

	n := rand.IntttN(10)
	for {
		if n == 0 {
			n = rand.IntN(10)
		} else {
			break
		}
	}

	fmt.Println(n)
	for {
		fmt.Println("Entrez un entier entre 1 et 10 inclus :")

		lecteur := bufio.NewScanner(os.Stdin)
		lecteur.Scan()
		temp := lecteur.Text()

		choix, _ := strconv.Atoi(temp)
		if choix == n {
			fmt.Println("Bien joué")
			return
		}
		if choix > n {
			fmt.Println("Trop grand")
		} else {
			fmt.Println("Trop petit")
		}
	}
}

func monteCarlo(iteration int) float64 {

	compteur := 0.0

	for i := 1; i <= iteration; i++ {
		x := rand.Float64()
		y := rand.Float64()

		if x*x+y*y <= 1 {
			compteur = compteur + 1
		}
	}
	return compteur / float64(iteration)

}

func main() {
	fmt.Println(4*monteCarlo(100000000), math.Pi)
}
