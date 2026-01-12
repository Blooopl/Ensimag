package main

import (
	"fmt"
	"math"
	"os"
)

func somme(n int) int {
	var res = 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}

func testSomme() {
	for n := -1; n <= 5; n++ {
		fmt.Print("Somme des entiers naturels dans [0..", n, "] : ", somme(n))
		fmt.Println()
	}
	fmt.Println()
}

func fact(n uint) uint {
	var res uint
	for res = 1; n > 1; n-- {
		res *= n
	}
	return res
}

func testFact() {
	var i uint
	for i = 0; i < 6; i++ {
		fmt.Print(i, "! = ", fact(i), "\n")
	}
	fmt.Println()
}

func pgcd(a, b uint) int {
	if (a == 0) || (b == 0) {
		return -1
	}
	for a != b {
		if a < b {
			b -= a
		} else {
			a -= b
		}
	}
	return int(a)
}

func testPGCD() {
	var a uint = 0
	var b uint = 0
	res := pgcd(a, b)
	fmt.Print("PGCD(", a, ",", b, ") = ")
	if res == -1 {
		fmt.Fprintln(os.Stderr, "Erreur : a et b sont nuls !")
	} else {
		fmt.Println(res)
	}
	b = 1
	res = pgcd(a, b)
	fmt.Print("PGCD(", a, ",", b, ") = ")
	if res == -1 {
		fmt.Fprintln(os.Stderr, "Erreur : a est nul !")
	} else {
		fmt.Println(res)
	}
	res = pgcd(b, a)
	fmt.Print("PGCD(", b, ",", a, ") = ")
	if res == -1 {
		fmt.Fprintln(os.Stderr, "Erreur : b est nul !")
	} else {
		fmt.Println(res)
	}
	a = 15
	b = 10
	fmt.Print("PGCD(", a, ",", b, ") = ", pgcd(a, b), "\n")
	fmt.Print("PGCD(", b, ",", a, ") = ", pgcd(b, a), "\n")
	fmt.Println()
}

func racine(x float64) float64 {
	if x < 0.0 {
		return math.NaN()
	}
	nbrIter := 10
	var res float64 = 1.0
	for nbrIter > 0 {
		res = (res + (x / res)) / 2.0
		nbrIter -= 1
	}
	return res
}

func testRacine() {
	for val := -1.0; val < 10.0; val += 1.0 {
		fmt.Print("racine(", val, ") = ", racine(val), " ~ ", math.Sqrt(val), "\n")
	}
}

func main() {
	testSomme()
	testFact()
	testPGCD()
	testRacine()
}
