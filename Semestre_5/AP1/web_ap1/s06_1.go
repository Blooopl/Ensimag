package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func estExtremum(xim1, xi, xip1 int) bool {
	if ((xim1 < xi) && (xi > xip1)) || ((xim1 > xi) && (xi < xip1)) {
		// On affiche une trace au passage pour faciliter la mise au point
		fmt.Println("=> extremum :", xi)
		return true
	}
	return false
}

// C'est bien de factoriser ce genre de traitements pour alléger le programme principal
func decodeEntier(chaine string) int {
	val, err := strconv.Atoi(chaine)
	if err != nil {
		log.Fatal(err)
	}
	return val
}

func main() {
	fmt.Println("Entrez chaque valeur de la suite suivie d'un retour-chariot (ctrl-d pour finir) :")
	lecteur := bufio.NewScanner(os.Stdin)
	if !lecteur.Scan() {
		// EOF dès le début : suite vide
		fmt.Println("Nombre d'extremums : 0 (suite vide)")
		return
	}
	prec := decodeEntier(lecteur.Text())
	var cour int
	// Cette boucle sert à avaler les entiers égaux en tête de la suite, qui peuvent indiquer :
	// - soit une suite constante si on arrive à EOF sans avoir trouvé de valeur différente ;
	// - soit une répétition en tête de la suite
	for {
		if !lecteur.Scan() {
			// EOF après des valeurs toutes égales : suite constante
			fmt.Println("Nombre d'extremums : 1 (suite constante)")
			return
		}
		cour = decodeEntier(lecteur.Text())
		if cour != prec {
			// On sort de la boucle dès qu'on a trouvé une valeur différente de la ou les premières de la suite
			break
		}
	}
	// On a éliminé les cas des suites vides et constantes : il y aura donc toujours au moins 2 extremums (les bornes)
	nbr := 2
	for lecteur.Scan() {
		suiv := decodeEntier(lecteur.Text())
		if suiv == cour {
			// On avale les répétitions en redémarrant la boucle sans traiter la valeur égale à celle d'avant
			continue
		}
		if estExtremum(prec, cour, suiv) {
			nbr += 1
		}
		// "Fenêtre glissante" : on perd l'ancienne valeur de prec (qui ne sert plus à rien) et on prépare l'itération suivante
		prec, cour = cour, suiv
	}
	fmt.Println("Nombre d'extremums :", nbr)
}
