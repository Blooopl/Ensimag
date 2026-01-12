package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"unicode"
)

const carParLigne = 5

// Affichage de 5 colonnes par ligne
func afficher(car rune, nbr uint, num *int) {
	if *num == carParLigne {
		fmt.Println()
		*num = 0
	} else {
		sep := "    "
		if *num == carParLigne-1 {
			sep = ""
		}
		fmt.Printf("'%c' : %02d%s", car, nbr, sep)
		*num++
	}
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal(fmt.Sprint("usage : ", path.Base(os.Args[0]), " <fichier .txt>"))
	}
	fich, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer fich.Close()
	dico := make(map[rune]uint)
	for lecteur := bufio.NewScanner(fich); lecteur.Scan(); {
		for _, car := range lecteur.Text() { // On utilise un range sur la string Text pour récupérer les runes (et pas les bytes qui les composent !)
			dico[car]++
		}
	}
	fmt.Println("Affichage dans un ordre aléatoire :")
	num := 0
	for car, nbr := range dico {
		afficher(car, nbr, &num) // on passe un pointeur sur num car afficher change ce paramètre
	}
	fmt.Println("\nAffichage dans l'ordre des caractères :")
	// Là l'idée est d'énumérer tous les caractères possibles entre 0 et le dernier caractère latin : on aurait pu faire plus efficace, d'ailleurs c'est comme ça qu'on fera au TP suivant !
	for car, num := '\u0000', 0; car <= unicode.MaxLatin1; car++ {
		// Si l'accès à la valeur associée à la clé car est 0, c'est que le caractère n'apparait pas dans le texte (valeur nulle renvoyée en cas d'absence de la clé cherchée)
		if nbr := dico[car]; nbr > 0 {
			afficher(car, nbr, &num)
		}
	}
	fmt.Println()
}
