package main

import (
	"fmt"
	"math"
)

// Les types ci-dessous sont définis globalement : ils sont donc visibles dans toutes les fonctions
type infos struct {
	nomPrenom string
	naissance date
	taille    float64
	motard    bool
}

type date struct {
	jour  uint8
	mois  lesMois
	annee int
}

type lesMois int

const (
	janvier lesMois = iota + 1 // iota vaut 0
	février                    // Si on ne précise pas de valeur, on ajoute un par rapport à la ligne précédente
	mars
	avril
	mai
	juin
	juillet
	aout
	septembre
	octobre
	novembre
	décembre
)

func afficher(moi infos) {
	fmt.Println("Je m'appelle", moi.nomPrenom)
	fmt.Print("Je suis né le ", moi.naissance.jour, "/", moi.naissance.mois, "/", moi.naissance.annee, "\n")
	entier, frac := math.Modf(moi.taille)
	fmt.Println("Je mesure", int(entier), "mètre", int(frac*100))
	if moi.motard {
		fmt.Println("Gaaaaaz !")
	} else {
		fmt.Println("C'est quoi une moto ?")
	}
}

func main() {
	eddy := infos{
		nomPrenom: "Édouard Bracame",
		naissance: date{
			jour:  15,
			mois:  juin,
			annee: 1950,
		},
		taille: 1.78,
		motard: true,
	}
	obelix := infos{
		nomPrenom: "Obelix",
		naissance: date{
			jour:  4,
			mois:  septembre,
			annee: -70,
		},
		taille: 1.83,
		motard: false,
	}
	fmt.Println(eddy)
	fmt.Println(obelix)
	fmt.Println()
	afficher(eddy)
	fmt.Println()
	afficher(obelix)
}
