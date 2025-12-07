package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	testRot('n', 39, false)
	testRot('i', 13, false)
	testRot('r', 13, false)
	testRot('P', 13, false)
	testRot('r', 13, false)
	testRot('f', 13, false)
	testRot('n', 13, false)
	testRot('e', 13, false)
}

func rot(lettre rune, decalage int32) (rune, error) {
	decalage = decalage % 26
	var vide rune = ' '
	neg := errors.New("Decalage négatif")
	if decalage < 0 {
		return vide, neg
	} else if ((lettre+decalage) > 90 && (lettre+decalage) < 97) || (lettre+decalage > 122) {
		return lettre + decalage - 26, nil
	} else {
		return lettre + decalage, nil
	}

}

func testRot(lettre rune, decalage int32, verbeux bool) {

	res, err := rot(lettre, decalage)

	if err != nil {
		log.Fatal(err)
	}

	if verbeux {
		fmt.Println("rot(", string(lettre), ")=", string(res))
	} else {
		fmt.Println(string(res))
	}
}
