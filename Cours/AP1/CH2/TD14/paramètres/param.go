package main

import (
	"fmt"
	"os"
)

func main() {
	afficher()
}

func afficher() {

	n := len(os.Args)

	if n == 1 {
		fmt.Println(os.Args[0])
	} else {
		n++
	}

}
