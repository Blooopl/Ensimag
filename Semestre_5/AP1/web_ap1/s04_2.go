package main

import (
	"errors"
	"fmt"
	"log"
)

func rot(lettre rune, decalage int32) (rune, error) {
	if decalage < 0 {
		return 0, errors.New("le décalage doit être positif")
	}
	// Les runes sont des int32 : on peut donc utliser les opérateurs classiques de comparaison
	estMaj := lettre >= 'A' && lettre <= 'Z'
	estMin := lettre >= 'a' && lettre <= 'z'
	if !estMaj && !estMin {
		return 0, errors.New("la lettre doit être une majuscule ou une minuscule")
	}
	var premLettre rune
	if estMaj {
		premLettre = 'A'
	} else {
		premLettre = 'a'
	}
	// % est l'opérateur modulo, ici du nombre de lettres dans l'alphabet
	return (lettre-premLettre+decalage)%26 + premLettre, nil
}

func testRot(lettre rune, decalage int32, verbeux bool) {
	code, err := rot(lettre, decalage)
	if err != nil {
		log.Fatal(err)
	}
	if verbeux {
		fmt.Print("rot", decalage, "(", string(lettre), ") = ", string(code), "\n")
	} else {
		fmt.Printf("%c", code)
	}
}

func main() {
	testRot('a', 1, true)
	testRot('Z', 1, true)
	testRot('x', 5, true)
	testRot('A', 3, true)
	testRot('b', 3, true)
	testRot('C', 3, true)

	testRot('n', 13, false)
	testRot('i', 13, false)
	testRot('r', 13, false)
	testRot('P', 13, false)
	testRot('r', 13, false)
	testRot('f', 13, false)
	testRot('n', 13, false)
	testRot('e', 13, false)
	fmt.Println()
}
