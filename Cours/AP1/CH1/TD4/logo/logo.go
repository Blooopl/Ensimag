package main

import (
	"fmt"
	"math"
)

var positionTortue = pointSVG{0, 0}
var angleTortue float64 = 90.0
var crayonEnBas bool = false

var hauteur float64 = 200
var largeur float64 = 200

func avanceTortue(distance float64) {
	equi_rad := angleTortue * math.Pi / 180
	new_x := positionTortue.x + distance*math.Cos(equi_rad)
	new_y := positionTortue.y - distance*math.Sin(equi_rad)

	new_point := pointSVG{new_x, new_y}
	if crayonEnBas {
		fmt.Println(segmentSVG(positionTortue, new_point))
	}

	positionTortue = new_point
}

func tourneTortueGauche(deltaAngle float64) {
	angleTortue = angleTortue + deltaAngle
}

func tourneTortueDroite(deltaAngle float64) {
	angleTortue = angleTortue - deltaAngle
}

func baisseCrayon() {
	crayonEnBas = true
}

func leveCrayon() {
	crayonEnBas = false
}