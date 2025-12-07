package main

import "fmt"

func main() {
	var hauteur float64 = 200
	var largeur float64 = 200
	fmt.Println(debutImageSVG(largeur, hauteur))

	fmt.Println(debutGroupeSVG("red", "red", 5))

	var centre pointSVG
	centre.x = 10
	centre.y = 10

	var rayon float64
	rayon = 30

	fmt.Println(cercleSVG(centre, rayon))

	fmt.Println(finGroupeSVG())

	fmt.Println(debutGroupeSVG("red", "black", 5))

	centre.x, centre.y = 100, 100
	rayon = 60

	fmt.Println(cercleSVG(centre, rayon))

	centre.x, centre.y = 180, 50
	rayon = 10

	fmt.Println(cercleSVG(centre, rayon))

	fmt.Println(finGroupeSVG())

	fmt.Println(finImageSVG())

}
