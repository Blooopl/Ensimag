package main

import (
	"fmt"
	"math"
)

func somme(n int) int {

	if n < 0 {
		return 0
	} else {
		somme := 0
		for i := 0; i <= n; i++ {
			somme += i
		}

		return somme
	}

}

func testSomme() {
	for test := -1; test <= 5; test++ {
		fmt.Println(somme(test))
	}
}

func fact(n uint) uint {
	temp := 1
	if n == 0 || n == 1 {
		return 1
	}

	for i := n; i > 0; i-- {
		temp = temp * int(i)
	}

	return uint(temp)
}

func testFact() {
	for test := 0; test <= 5; test++ {
		fmt.Println(fact(uint(test)))
	}
}

func pgcd(a, b uint) int {
	if a == 0 || b == 0 {
		return -1
	}

	for {
		if a == b {
			break
		}

		if a > b {
			a = a - b
		} else {
			b = b - a
		}
	}

	return int(a)
}

func testPGCD() {
	fmt.Println(pgcd(120, 60))
	fmt.Println(pgcd(0, 12))
	fmt.Println(pgcd(50, 43))
}

func racine(x float64) float64 {
	if x < 0 {
		return math.NaN()
	}
	r_i := 1.0
	for i := 1; i <= 10; i++ {
		r_i = (r_i + x/r_i) / 2
	}

	return r_i

}

func testRacine() {
	for test := -1; test <= 9; test++ {
		fmt.Println(math.Sqrt(float64(test)), racine(float64(test)))
	}
}

func main() {
	testRacine()
}
