package main

import (
	"fmt"
	"math"
)

func main() {

	chaine1 := "Orni"
	chaine2 := "thorynque"
	entier := 42
	boole := true
	flottant := math.Pi

	//fmt.Println(chaine1, chaine2, entier, boole, flottant)
	//fmt.Print(chaine1, "\n", chaine2, "\n", entier, "\n", boole, "\n", flottant)

	//str := fmt.Sprintln(chaine1, chaine2, entier, boole, flottant)
	//fmt.Println(str)

	//fmt.Printf("%d est un nombre pair ? %t\n", entier, entier%2 == 0)

	fmt.Printf("%6s %6s %.3d %.8b %t %.4f \n", chaine1, chaine2, entier, entier, boole, flottant)
	fmt.Printf("%T %T %T %T %T", chaine1, chaine2, entier, boole, flottant)
}
