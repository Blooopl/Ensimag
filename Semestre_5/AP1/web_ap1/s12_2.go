package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"slices"
	"strconv"
	"strings"
)

// Lit des entiers signés depuis stdin et renvoie un tableau les contenant.
func lectureDonnees() []int {
	lecteur := bufio.NewScanner(os.Stdin)
	tab := make([]int, 0)
	for lecteur.Scan() {
		lesNombres := strings.Fields(lecteur.Text())
		// Fields renvoie un tableau vide s'il n'y a que des séparateurs (ligne vide)
		//   => on ne rentre pas dans la boucle dans ce cas
		for idx := 0; idx < len(lesNombres); idx++ {
			val, err := strconv.Atoi(lesNombres[idx])
			if err != nil {
				log.Fatal("erreur : \"%s\" n'est pas un entier", lesNombres[idx])
			}
			tab = append(tab, val)
		}
	}
	return tab
}

// Inverse le contenu du sous-tableau tab[0..dern] (indices inclus).
// Précondition : dern est dans [0..len(tab)[
func inverser(tab []int, dern int) {
	for prem := 0; prem < dern; prem, dern = prem+1, dern-1 {
		tab[prem], tab[dern] = tab[dern], tab[prem]
	}
}

// Renvoie l'indice de la valeur maximale dans le sous-tableau tab[0..sup[.
// Précondition : sup est dans [0..len(tab)] et len(tab) > 0
func rechercherIndiceValMax(tab []int, sup int) int {
	idxMax := 0
	for idx := 1; idx < sup; idx++ {
		if tab[idx] > tab[idxMax] {
			idxMax = idx
		}
	}
	return idxMax
}

// Tri le tableau selon l'algorithme de la crêpe.
func trierCrepe(tab []int) {
	for sup := len(tab); sup > 1; sup-- {
		idxMax := rechercherIndiceValMax(tab, sup)
		if idxMax != sup-1 {
			if idxMax != 0 {
				inverser(tab, idxMax)
			}
			inverser(tab, sup-1)
		}
	}
}

// Génère un tableau d'une taille donnée d'entiers signés tirés aléatoirement.
func genererTab(taille int) []int {
	const amplitude = 21
	tab := make([]int, taille)
	for idx := 0; idx < len(tab); idx++ {
		tab[idx] = rand.Intn(amplitude) - amplitude/2
	}
	return tab
}

// Effectue beaucoup de tests du tri sur des tableaux de «grandes » tailles.
// On affiche les traces sur stderr au cas où on voudrait afficher les tableaux et rediriger stdout dans un fichier.
func testAutoTri() {
	fmt.Fprintln(os.Stderr, "Test automatique du tri de la crêpe :")
	const tailleMax = 1000
	const nbrIter = 100
	for taille := 0; taille <= tailleMax; taille++ {
		fmt.Fprintf(os.Stderr, "\r=> tableau de %v éléments (max : %v)", taille, tailleMax)
		for idx := 0; idx < nbrIter; idx++ {
			tab := genererTab(taille)
			ref := slices.Clone(tab)
			trierCrepe(tab)
			slices.Sort(ref)
			if !slices.Equal(ref, tab) {
				log.Fatal("erreur : le tableau résultat est faux")
			}
		}
	}
	fmt.Fprintln(os.Stderr, "\n=> OK !")
	fmt.Fprintln(os.Stderr)
}

// Programme principal.
func main() {
	// testAutoTri()
	tab := lectureDonnees()
	trierCrepe(tab)
	fmt.Println(tab)
}
