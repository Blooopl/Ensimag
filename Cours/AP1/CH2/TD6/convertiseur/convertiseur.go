package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	var x int
	var y int

	format_error := errors.New("Pas de fichier pbm")

	file, err := os.Open("s06_2_image.pbm")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	lecteur := bufio.NewScanner(file)

	lecteur.Scan()
	format := lecteur.Text()

	lecteur.Scan()
	largeur, _ := strconv.Atoi(lecteur.Text())

	lecteur.Scan()
	hauteur, _ := strconv.Atoi(lecteur.Text())

	if format != "P1" {
		log.Fatal(format_error)
	}

	fmt.Println(debutImageSVG(float64(largeur), float64(hauteur)))
	fmt.Println(debutGroupeSVG("black", "white", 1))

	for y = 0; y < hauteur; y++ {
		for x = 0; x < largeur; x++ {
			lecteur.Scan()
			bit, _ := strconv.Atoi(lecteur.Text())
			if bit == 1 {
				var centre pointSVG
				centre.x = float64(x)
				centre.y = float64(y)

				fmt.Println(cercleSVG(centre, 1))
			}
		}
	}

	fmt.Println(finGroupeSVG())
	fmt.Println(finImageSVG())

}
