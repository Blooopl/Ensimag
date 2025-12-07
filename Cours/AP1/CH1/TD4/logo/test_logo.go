package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(debutImageSVG(100, 100))
	fmt.Println(debutGroupeSVG("black", "white", 3))
	smiley()
	fmt.Println(finGroupeSVG())
	fmt.Println(finImageSVG())
}

func trait() {
	tourneTortueDroite(135)
	baisseCrayon()
	avanceTortue(math.Sqrt2 * 100.0)
}

func carre() {
	tourneTortueDroite(135)
	avanceTortue(math.Sqrt2 * 25)
	//tourneTortueGauche(45)
	baisseCrayon()
	tourneTortueGauche(45)
	avanceTortue(50)
	tourneTortueDroite(90)
	avanceTortue(50)
	tourneTortueDroite(90)
	avanceTortue(50)
	tourneTortueDroite(90)
	avanceTortue(50)
}

func triangle() {
	tourneTortueDroite(135)
	avanceTortue(50 * math.Sqrt2)

	tourneTortueGauche(135)
	avanceTortue(25)

	baisseCrayon()

	tourneTortueDroite(150)
	avanceTortue(50)

	tourneTortueDroite(120)
	avanceTortue(50)

	tourneTortueDroite(120)
	avanceTortue(50)

}

func etoile() {

	pas := 40.0
	angle := 35.5
	tourneTortueDroite(135)
	avanceTortue(50 * math.Sqrt2)

	baisseCrayon()

	tourneTortueGauche(135)
	avanceTortue(25)

	tourneTortueDroite(180 - angle)
	avanceTortue(pas)

	tourneTortueDroite(180 - angle)
	avanceTortue(pas)

	tourneTortueDroite(180 - angle)
	avanceTortue(pas)

	tourneTortueDroite(180 - angle)
	avanceTortue(pas)

	tourneTortueDroite(180 - angle)
	avanceTortue(pas)

	tourneTortueDroite(180 - angle)
	avanceTortue(pas)
}


func smiley(){

	leveCrayon()
	tourneTortueDroite(135)
	avanceTortue(math.Sqrt2 * 25)

	tourneTortueGauche(45)
	baisseCrayon()

	//tourneTortueGauche(45)
	avanceTortue(20)
	tourneTortueDroite(90)
	avanceTortue(20)
	tourneTortueDroite(90)
	avanceTortue(20)
	tourneTortueDroite(90)
	avanceTortue(20)

	leveCrayon()

	tourneTortueDroite(90)
	avanceTortue(30)

	baisseCrayon()
	avanceTortue(20)
	tourneTortueDroite(90)
	avanceTortue(20)
	tourneTortueDroite(90)
	avanceTortue(20)
	tourneTortueDroite(90)
	avanceTortue(20)


	tourneTortueDroite(180)

	leveCrayon()
	avanceTortue(60)
	baisseCrayon()

	tourneTortueGauche(135)
	avanceTortue(math.Sqrt2 * 25)

	tourneTortueDroite(180)
	avanceTortue(math.Sqrt2 * 25)

	tourneTortueDroite(45)
	avanceTortue(10)

	tourneTortueDroite(45)
	avanceTortue(math.Sqrt2 * 25)







}