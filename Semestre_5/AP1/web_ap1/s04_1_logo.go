package main

import (
	"fmt"
	"math"
)

// En Go, il y a plein de façon d'initialiser des variables.
// On a vu en cours qu'on ne peut pas utiliser la syntaxe condensée avec := quand il s'agit de variables globales.
// Par contre, toutes les syntaxes ci-dessous sont correctes.

// Version longue :
// var positionTortue pointSVG = pointSVG{0, 0}
// Version courte :
var positionTortue = pointSVG{0, 0}

// On n'a même pas besoin de préciser le type quand il s'agit d'un type inférable par le compilateur
var angleTortue = 90.0

// Et bien sûr, la version classique
var crayonEnBas bool = false

func avanceTortue(distance float64) {
	destination := pointSVG{
		x: positionTortue.x + math.Cos(angleTortue*math.Pi/180.0)*distance,
		y: positionTortue.y - math.Sin(angleTortue*math.Pi/180.0)*distance,
	}
	if crayonEnBas {
		fmt.Println(segmentSVG(positionTortue, destination))
	}
	positionTortue = destination
}

func tourneTortueGauche(deltaAngle float64) {
	angleTortue += deltaAngle
}

func tourneTortueDroite(deltaAngle float64) {
	angleTortue -= deltaAngle
}

func baisseCrayon() {
	crayonEnBas = true
}

func leveCrayon() {
	crayonEnBas = false
}
