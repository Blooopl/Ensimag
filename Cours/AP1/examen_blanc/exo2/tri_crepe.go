package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"slices"
	"strconv"
	"strings"
)

func extraction() []int {
	tab := make([]int, 0)

	lecteur := bufio.NewScanner(os.Stdin)
	for lecteur.Scan() {
		entree := lecteur.Text()
		parse := strings.Fields(entree)
		for i := 0; i < len(parse); i++ {
			temp, err := strconv.Atoi(parse[i])
			if err != nil {
				log.Fatal(err)
			}
			tab = append(tab, temp)
		}

	}

	return tab
}

func tri_crepe(liste []int) {
	taille := len(liste)
	sup := taille - 1

	id := 0
	for sup = taille; sup >= 0; sup-- {
		id = id_max(liste, sup)
		if id == sup-1 {
		} else {
			if sup != 0 {
				inverse(liste, 0, id)
				inverse(liste, 0, sup-1) //
			}
		}
	}
}
func main() {
	//liste := extraction()
	//tri_crepe(liste)

	fmt.Println(test_auto(1000, 100))
}

func inverse(liste []int, debut, fin int) {

	temp := make([]int, fin-debut+1)

	j := 0
	for i := fin; i >= debut; i-- {
		temp[j] = liste[i]
		j++
	}

	for i := 0; i <= len(temp)-1; i++ {
		liste[debut+i] = temp[i]
	}

}

func id_max(liste []int, sup int) int {

	if len(liste) <= 0 {
		return -1
	}

	id := 0
	max := liste[0]

	for i := 0; i < sup; i++ {
		if liste[i] > max {
			max = liste[i]
			id = i
		}
	}

	return id
}

func remplir(temp []int) {

	for i := 0; i < len(temp); i++ {
		temp[i] = rand.IntN(11)
		minus := rand.IntN(2)
		if minus == 0 {
			temp[i] = -temp[i]
		}
	}
}

func test_auto(nb_test, taille_test int) bool {

	for test := 0; test < 100; test++ {
		temp := make([]int, taille_test)
		remplir(temp)
		tri_crepe(temp)

		copie := make([]int, taille_test)
		copy(copie, temp)
		slices.Sort(copie)

		equal := slices.Equal(copie, temp)
		if !equal {
			return false
		}
	}
	return true
}
