package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const SIGNATURE = "P1"

func main() {
	lecteur := bufio.NewScanner(os.Stdin)
	// On vérifie juste la signature
	if !lecteur.Scan() || lecteur.Text() != SIGNATURE {
		log.Fatal("erreur : ce fichier n'est pas au format PBM")
	}
	// À partir de là, on considère que le fichier est bien formaté sans vérifier
	lecteur.Scan()
	largeur, _ := strconv.Atoi(lecteur.Text())
	lecteur.Scan()
	hauteur, _ := strconv.Atoi(lecteur.Text())
	fmt.Println(debutImageSVG(float64(largeur), float64(hauteur)))
	// On crée un fond blanc sur lequel on dessinera en noir
	fmt.Println(debutGroupeSVG("black", "white", 1.0))
	for lig := 0; lig < hauteur; lig++ {
		for col := 0; col < largeur; col++ {
			lecteur.Scan()
			// Un point noir == un cercle de rayon 1
			if lecteur.Text() == "1" {
				fmt.Println(cercleSVG(pointSVG{float64(col), float64(lig)}, 1))
			}
		}
	}
	fmt.Println(finGroupeSVG())
	fmt.Println(finImageSVG())
}
