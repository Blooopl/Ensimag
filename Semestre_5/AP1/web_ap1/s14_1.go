package main

import (
	"fmt"
	"math/rand"
)

const tailleMaxTab = 10

const tailleDonnées = 1_000_000

const premierChar = 'A'
const dernierChar = 'Z'

type entrée struct {
	cle     *rune
	données *[tailleDonnées]byte
}

type tableau []*entrée

func creerEntrée() *entrée {
	val := new(entrée)                                                    // là on alloue l'entrée, c'est à dire la structure
	val.cle = new(rune)                                                   // puis on alloue le champ rune de l'entrée...
	*val.cle = rune(rand.Int()%(dernierChar-premierChar+1) + premierChar) // ...et on affecte une valeur à ce champ
	val.données = new([tailleDonnées]byte)                                // on finit par allouer le tableau STATIQUE contenant les données
	return val
}

func initTab(tab tableau) {
	for idx := 0; idx < len(tab); idx++ {
		tab[idx] = creerEntrée()
	}
}

func afficherTab(tab tableau) {
	fmt.Print("[ ")
	for idx := 0; idx < len(tab); idx++ {
		// attention au typage :
		// - tab[idx] est un pointeur sur une entrée (une struct)
		// - on devrait mettre une * pour déréférencer et accéder à son champ cle MAIS Go sait le faire tout seul pour les pointeurs sur des structs
		// - le * porte en fait sur la cle, qui est elle-même un pointeur vers la rune qu'on veut afficher
		// - un problème courant concerne la priorité des opérateurs : est-ce qu'il considère que [] est prioritaire sur *, qu'en est-il du . etc ?
		// - l'expression ci-dessous est en fait équivalente à :
		// *(
		//     (
		//        *(
		//            tab[idx]
		//         )
		//     ).cle
		//  )
		// -- tab[idx] est calculée en premier, c'est une *entree
		// -- le * que Go fait tout seul est le plus interne, qui produit une entree
		// -- ensuite Go accède à .cle qui est un champ de la struct entree
		// -- et à la fin, il déréférence cette cle pour obtenir la rune pointée
		fmt.Print(string(*tab[idx].cle), " ")
	}
	fmt.Println("]")
}

func triBulle(tab tableau) {
	borneSup := len(tab)
	for {
		echange := false
		for idx := 1; idx < borneSup; idx++ {
			if *tab[idx-1].cle > *tab[idx].cle { // attention au typage : ce sont bien les VALEURS des runes qu'on veut comparer (idem que ci-dessus pour le principe)
				tab[idx-1], tab[idx] = tab[idx], tab[idx-1] // là par contre, on échange bien les POINTEURS vers les tableaux statiques
				echange = true
			}
		}
		borneSup--
		if !echange {
			break
		}
	}
}

func testVisuel() {
	for tailleTab := 0; tailleTab <= tailleMaxTab; tailleTab++ {
		fmt.Println("***** taille ==", tailleTab, "*****")
		tab := make(tableau, tailleTab)
		initTab(tab)
		afficherTab(tab)
		triBulle(tab)
		afficherTab(tab)
		fmt.Println()
	}
}

func main() {
	testVisuel()
}
