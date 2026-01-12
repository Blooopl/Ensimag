package main

import "fmt"

func main() {
	var x int = 5 // c'est le x externe à la boucle
	for x != 10 { // c'est le x externe à la boucle !!
		x := 10        // c'est le x interne à la boucle (qui masque le x externe)
		fmt.Println(x) // c'est le x interne à la boucle
	} // => boucle infinie
	fmt.Println(x)

	for idx := 0; idx < 5; idx++ { // idx est une variable locale à cette boucle imbriquante
		for idx = 4; idx < 5; idx++ { // c'est bien le même idx qu'au dessus, on change sa valeur
			fmt.Print(idx, " ")
		}
		// ici idx vaut 5, donc on va sortir de la boucle imbriquante
	}
	fmt.Println()
	// affichage : 4

	for idx := 0; idx < 5; idx++ { // idx est une variable locale à cette boucle imbriquante
		for idx := 4; idx < 5; idx++ { // on alloue une nouvelle variable locale à la boucle
			fmt.Print(idx, " ") //          imbriquée, qui masque donc le idx du dessus
		} // => le idx externe n'est pas modifié par la boucle interne
	}
	fmt.Println()
	// affichage : 4 4 4 4 4
}
