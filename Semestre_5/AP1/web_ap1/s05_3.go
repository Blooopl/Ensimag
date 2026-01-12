package main

import "fmt"

const coteCase = 25.0

const nbrCases = 4

func couleurSuivante(couleurCourante string) string {
	if couleurCourante == "black" {
		return "white"
	}
	return "black"
}

func dessinerLigne(ligne int, couleurPremiereCase string) {
	for colonne, couleur := 0, couleurPremiereCase; colonne < nbrCases; colonne++ {
		fmt.Println(rectangleSVG(pointSVG{float64(colonne) * coteCase, float64(ligne) * coteCase}, coteCase, coteCase, couleur))
		couleur = couleurSuivante(couleur)
	}
}

func echiquier() {
	fmt.Println(debutImageSVG(coteCase*nbrCases, coteCase*nbrCases))
	for ligne, couleur := 0, "black"; ligne < nbrCases; ligne++ {
		dessinerLigne(ligne, couleur)
		couleur = couleurSuivante(couleur)
	}
	fmt.Println(finImageSVG())
}

func main() {
	echiquier()
}
