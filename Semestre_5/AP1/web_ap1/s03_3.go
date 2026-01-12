package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	chaine1, chaine2, entier, booleen, flottant := "Orni", "thorynque", 42, true, math.Pi
	fmt.Println(chaine1, chaine2, entier, booleen, flottant)
	fmt.Print(chaine1, chaine2, entier, booleen, flottant, "\n")
	fmt.Fprintln(os.Stdout, chaine1, chaine2, entier, booleen, flottant)
	fmt.Print(fmt.Sprintln(chaine1, chaine2, entier, booleen, flottant))
	fmt.Printf("%6s %5s %03d %08b %t %.4f\n", chaine1, chaine2, entier, entier, booleen, flottant)
	fmt.Printf("%T %T %T %T %T\n", chaine1, chaine2, entier, booleen, flottant)
}
