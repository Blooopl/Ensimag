package main

import "fmt"

// Le type representant un point, c'est à dire ses coordonnées.
type pointSVG struct {
	x float64
	y float64
}

// Retourne la chaîne indiquant le debut du fichier SVG et precisant la taille de l'image.
func debutImageSVG(largeur, hauteur float64) string {
	return fmt.Sprint("<svg xmlns=\"http://www.w3.org/2000/svg\" version=\"1.1\" width=\"", largeur, "\" height=\"", hauteur, "\">")
}

// Retourne la chaîne indiquant la fin du fichier SVG.
func finImageSVG() string {
	return "</svg>"
}

// Retourne la chaîne indiquant le début d'un groupe d'objets SVG.
// Un groupe SVG permet de factoriser certains paramètres, comme la couleur et l'épaisseur du trait ou la couleur de remplissage des formes.
func debutGroupeSVG(couleurLigne, couleurRemplissage string, epaisseurLigne float64) string {
	return fmt.Sprint("<g stroke=\"", couleurLigne, "\" fill=\"", couleurRemplissage, "\" stroke-width=\"", epaisseurLigne, "\">")
}

// Retourne la chaîne indiquant le début d'un groupe d'objets SVG à fond transparent.
func debutGroupeTransparentSVG(opacite float64) string {
	return fmt.Sprint("<g opacity=\"", opacite, "\">")
}

// Retourne la chaîne indiquant la fin d'un groupe d'objet SVG.
func finGroupeSVG() string {
	return "</g>"
}

// Retourne la chaîne decrivant un cercle SVG, en précisant les coordonnées de son centre, ainsi que son rayon.
func cercleSVG(centre pointSVG, rayon float64) string {
	return fmt.Sprint("<circle cx=\"", centre.x, "\" cy=\"", centre.y, "\" r=\"", rayon, "\" />")
}

// Retourne la chaîne decrivant un segment SVG, en précisant les points de départ et d'arrivée.
func segmentSVG(depart, arrivee pointSVG) string {
	return fmt.Sprint("<line x1=\"", depart.x, "\" y1=\"", depart.y, "\" x2=\"", arrivee.x, "\" y2=\"", arrivee.y, "\" />")
}

// Retourne la chaîne décrivant un rectangle SVG, en précisant le point en haut à gauche, la largeur et hauteur du rectangle, ainsi que la couleur du fond.
// Si on veut que le fond soit transparent, il suffit de passer la chaîne "" comme couleur.
// Si le fond n'est pas transparent, les contours du rectangle seront invisibles.
func rectangleSVG(origine pointSVG, largeur, hauteur float64, couleur string) string {
	if couleur != "" {
		couleur = fmt.Sprint(" stroke=\"none\" fill=\"", couleur, "\"")
	}
	return fmt.Sprint("<rect x=\"", origine.x, "\" y=\"", origine.y, "\" width=\"", largeur, "\" height=\"", hauteur, "\"", couleur, " />")
}

// Retourne la chaîne décrivant un triangle SVG, en précisant ses trois points.
func triangleSVG(sommet1, sommet2, sommet3 pointSVG) string {
	return fmt.Sprint("<polygon points=\"", sommet1.x, ",", sommet1.y, " ", sommet2.x, ",", sommet2.y, " ", sommet3.x, ",", sommet3.y, "\" />")
}

// Retourne la chaîne décrivant un texte SVG, en précisant sa position, la taille de la police et bien sûr le texte en question.
func texteSVG(origine pointSVG, texte string, taille float64) string {
	return fmt.Sprint("<text x=\"", origine.x, "\" y=\"", origine.y, "\" font-size=\"", taille, "\">", texte, "</text>")
}
