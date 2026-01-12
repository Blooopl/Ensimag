package main

import (
	"errors"
	"fmt"
	"log"
)

// Cette fonction est une fonction de calcul : elle ne gère pas l'erreur elle-même, elle la renvoie.
func division(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("on ne peut pas diviser par 0") // O comme quotient par convention s'il y a une erreur
	} else {
		return x / y, nil
	}
}

// Cette fonction est une fonction de test unitaire : elle teste la fonction de calcul une fois avec les paramètres donnés, et gère l'éventuelle erreur.
func testDiv(x, y float64) {
	// La chaîne "\n" permet d'aller à la ligne.
	// Attention : ce sont bien des guillemets, surtout pas des apostrophes.
	fmt.Print("Appel de testDiv(", x, ", ", y, ") :\n")
	res, err := division(x, y)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(x, "/", y, " = ", res, "\n\n")
}

func main() {
	testDiv(3.0, 4.0)
	testDiv(3.0, 0.0)
}
