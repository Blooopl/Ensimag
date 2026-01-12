package main

import (
	"fmt"
)

// Astuce : si les types ou fonctions du module SVG (pointSVG, debutImageSVG, etc.) sont soulignés dans VSCode :
// - ce n'est pas une erreur, c'est VSCode qui est bête !
// - fermez VSCode et ré-ouvrez les deux fichiers d'un coup depuis le terminal : code test_svg.go svg.go
// - attention : il faut bien ouvrir le fichier qui utilise les fonctions avant celui dans lequel elles sont définies !

func main() {
	pt1 := pointSVG{10, 10}
	pt2 := pointSVG{100, 100}
	pt3 := pointSVG{180, 50}
	// Il faut bien afficher les chaînes renvoyées par les fonctions du module SVG
	fmt.Println(debutImageSVG(200, 200))
	fmt.Println(debutGroupeSVG("black", "red", 5))
	fmt.Println(cercleSVG(pt1, 30))
	fmt.Println(finGroupeSVG())
	fmt.Println(debutGroupeSVG("red", "black", 5))
	fmt.Println(cercleSVG(pt2, 60))
	fmt.Println(cercleSVG(pt3, 10))
	fmt.Println(finGroupeSVG())
	fmt.Println(finImageSVG())
}
