package main

import "fmt"

const coteCase = 25.0

const nbrCases = 4

func echiquier() {
	fmt.Println(debutImageSVG(coteCase*nbrCases, coteCase*nbrCases))
	// Ligne 1

	for mul := 0; mul <= 3; mul++ {
		for ligne := 0; ligne <= 3; ligne++ {
			if (ligne+mul)%2 == 0 {
				fmt.Println(rectangleSVG(pointSVG{float64(ligne) * coteCase, float64(coteCase * mul)}, coteCase, coteCase, "black"))
			} else {
				fmt.Println(rectangleSVG(pointSVG{float64(ligne) * coteCase, float64(coteCase * mul)}, coteCase, coteCase, "white"))
			}
		}
	}
	fmt.Println(finImageSVG())
}

func main() {
	echiquier()
}

