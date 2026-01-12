package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	// for i := 0; i < len(os.Args); i++ {
	// 	fmt.Println(os.Args[i])
	// }
	if len(os.Args) == 2 {
		fmt.Println(os.Args[1])
	} else if (len(os.Args) == 3) && (os.Args[1] == "--inv") {
		// certains peuvent être tentés de modifier la chaîne pour inverser ses caractères...
		// ... ça ne peut pas marcher car les strings en Go ne peuvent pas être modifiées !
		// Go fournit plusieurs façons de faire ça efficacement (par exemple en convertissant la chaîne en tableau de runes ou en utilisant strings.Builder)
		// mais ici de toute façon on veut juste afficher la chaîne en sens inverse
		for idx := len(os.Args[2]) - 1; idx >= 0; idx-- {
			fmt.Printf("%c", os.Args[2][idx])
		}
		fmt.Println()
	} else {
		log.Fatal("usage :", path.Base(os.Args[0]), "[--inv] chaine")
	}
}
