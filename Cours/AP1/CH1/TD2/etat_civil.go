package main

import (
	"math"
	"strconv"
)

type lesMois int64

const (
	janvier lesMois = iota + 1
	fevrier
	mars
	avril
	mais
	juin
	juillet
	aout
	septembre
	octobre
	novembre
	decembre
)

type date struct {
	jour  int8
	mois  lesMois
	annee uint
}

type info struct {
	prenom    string
	nom       string
	naissance date
	taille    float64
	motarde   bool
}

func main() {

	gens1 := info{
		prenom: "Bastien",
		nom:    "Pichet",
		naissance: date{
			jour:  22,
			mois:  juillet,
			annee: 2005,
		},

		taille:  1.70,
		motarde: false,
	}

	gens2 := info{
		prenom: "Chloe",
		nom:    "Ramon",
		naissance: date{
			jour:  11,
			mois:  mars,
			annee: 2005,
		},

		taille:  1.70,
		motarde: false,
	}

	affichage(gens1)
	affichage(gens2)
}

func affichage(gens info) {

	println("Je m'apelle", gens.prenom, gens.nom)
	println(strconv.Itoa(int(gens.naissance.jour)) + "/" + strconv.Itoa(int(gens.naissance.mois)) + "/" + strconv.Itoa(int(gens.naissance.annee)))
	var entier float64
	var fraction float64
	entier, fraction = math.Modf(gens.taille)
	fraction = fraction * 100
	println("Je mesure", strconv.Itoa(int(entier)), "metre", strconv.Itoa(int(fraction)))
}
