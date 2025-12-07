package main

import "fmt"

func main() {

	chaine2 := "10€"
	for idx := 0; idx < len(chaine2); idx++ {
		fmt.Println(chaine2[idx], string(chaine2[idx]))
	}

	fmt.Println("")
	tabRunes := []rune("10€")
	for idx := 0; idx < len(tabRunes); idx++ {
		fmt.Println(tabRunes[idx], string(tabRunes[idx]))
	}
}
