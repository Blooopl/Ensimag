package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type etat int

const (
	Debut etat = iota
	UnBon
	DeuxBons
	TroisBons
)

func main() {
	state := Debut

	lecteur := bufio.NewScanner(os.Stdin)

	for {
		lecteur.Scan()
		choix, _ := strconv.Atoi(lecteur.Text())

		switch state {

		case Debut:
			switch choix {
			case 4:
				state = UnBon
			}

		case UnBon:
			switch choix {
			case 4:
				state = UnBon
			case 0:
				state = DeuxBons
			default:
				state = Debut
			}

		case DeuxBons:
			switch choix {
			case 4:
				state = UnBon
			case 9:
				state = TroisBons
			default:
				state = Debut
			}

		case TroisBons:
			switch choix {
			case 4:
				state = UnBon
			case 6:
				fmt.Println("Porte ouverte")
				break
			default:
				state = Debut
			}
		}
	}
}
