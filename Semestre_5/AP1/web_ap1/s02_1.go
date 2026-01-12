package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const annéeCourante = 2025

// Exercice d'intro
func intro() {
	fmt.Println("J'ecris sur stdout.")
	fmt.Fprintln(os.Stderr, "J'ecris sur stderr.")
}

// Version 1 : on ignore les erreurs
func v1() {
	fmt.Print("Entrez votre nom : ")
	lecteur := bufio.NewScanner(os.Stdin) // On crée un objet Scanner qui se branche sur le flux d'entrée (par défaut le clavier)
	lecteur.Scan()                        // On avance d'une ligne dans le "fichier" (ici c'est le clavier, donc jusqu'à ce qu'on tape entrée)
	nom := lecteur.Text()                 // On récupère la chaîne lue lors du Scan précédent
	fmt.Print("Bonjour ", nom, ", entrez votre année de naissance : ")
	lecteur.Scan()
	annee, _ := strconv.Atoi(lecteur.Text()) // On ignore l'erreur potentiellement renvoyée par Atoi en la "stockant" dans _
	fmt.Println("Re-bonjour", nom, "vous avez (environ)", annéeCourante-annee, "ans !")
}

// Version 2 : gestion des erreurs à la main
func v2() {
	fmt.Print("Entrez votre nom : ")
	lecteur := bufio.NewScanner(os.Stdin)
	lecteur.Scan()
	nom := lecteur.Text()
	fmt.Print("Bonjour ", nom, ", entrez votre année de naissance : ")
	lecteur.Scan()
	annee, err := strconv.Atoi(lecteur.Text()) // Cette fois-ci, on récupère vraiment l'erreur potentielle dans la variable err
	if err != nil {                            // Si err n'est pas nil, il y a eu une erreur : on la traite tout de suite
		fmt.Fprintln(os.Stderr, "Erreur : vous n'avez vraisemblablement pas saisi un entier !")
		os.Exit(1) // Cette fonction tue le programme directement, par convention on renvoie 1 à Unix pour dire qu'il y a eu une erreur
	}
	fmt.Println("Re-bonjour", nom, "vous avez (environ)", annéeCourante-annee, "ans !")
}

// Version 3 : gestion des erreurs à la Go
func v3() {
	fmt.Print("Entrez votre nom : ")
	lecteur := bufio.NewScanner(os.Stdin)
	lecteur.Scan()
	nom := lecteur.Text()
	fmt.Print("Bonjour ", nom, ", entrez votre année de naissance : ")
	lecteur.Scan()
	annee, err := strconv.Atoi(lecteur.Text())
	if err != nil {
		log.SetFlags(log.Flags() | log.Lshortfile) // Cette ligne ésotérique permet d'avoir des messages d'erreur plus détaillés quand on utilise log.Fatal
		log.Fatal(err)
	}
	fmt.Println("Re-bonjour", nom, "vous avez (environ)", annéeCourante-annee, "ans !")
}

func main() {
	// intro()
	// v1()
	// v2()
	v3()
}
