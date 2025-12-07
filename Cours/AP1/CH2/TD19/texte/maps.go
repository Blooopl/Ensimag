package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	maps := analysefichier()
	affichage_maps(maps)
}

func analysefichier() map[string]float64 {
	maps := make(map[string]float64)

	file, err := os.Open("s19_1_victor.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	lecteur := bufio.NewScanner(file)

	for lecteur.Scan() {
		texte := lecteur.Text()
		if texte != "" {
			for _, temp := range texte {
				maps[string(temp)]++
			}
		}
	}

	return maps
}

func affichage_maps(maps map[string]float64) {
	for cle, valeur := range maps {
		fmt.Print(cle, valeur)
	}
}
