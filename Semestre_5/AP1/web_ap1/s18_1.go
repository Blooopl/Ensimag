package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

const valSup = 10

const tailleMax = 7

// Dans cette séance les cellules contiennent un pointeur vers la cellule précédente, en plus d'un vers la suivante et de la valeur habituelle.
type cellule struct {
	val  int
	prec *cellule
	suiv *cellule
}

// On définit un type desc comme une structure contenant un pointeur vers la première cellule de la liste ainsi qu'un autre vers la dernière.
// Ces deux cellules sont en fait des éléments fictifs en tête et en queue.
type desc struct {
	tete  *cellule
	queue *cellule
}

// Et une liste est donc juste un pointeur vers une desc.
type liste *desc

func créer() liste {
	// Là on alloue une structure avec les pointeurs de tête et de queue : on rappelle que new renvoie un pointeurs vers ce qu'il alloue
	var ldc liste = new(desc)
	// Il faut donc penser à allouer les deux éléments fictifs
	ldc.tete = new(cellule)
	ldc.queue = new(cellule)
	// Et à les chaîner proprement l'un à l'autre (c'est une liste vide, il n'y a que les éléments fictifs)
	ldc.tete.suiv = ldc.queue
	ldc.queue.prec = ldc.tete
	return ldc
}

func afficher(ldc liste, teteVersQueue bool) {
	var cour, fin *cellule
	// On n'affiche pas le contenu d'une cellule fictive par convention (en fait on ne doit jamais y accéder)
	if teteVersQueue {
		cour = ldc.tete.suiv
		fin = ldc.queue
		fmt.Print("T <-> ")
	} else {
		cour = ldc.queue.prec
		fin = ldc.tete
		fmt.Print("Q <-> ")
	}
	// On peut factoriser la boucle d'affichage grâce aux variables ci-dessus
	for cour != fin {
		fmt.Print(cour.val, " <-> ")
		if teteVersQueue {
			cour = cour.suiv
		} else {
			cour = cour.prec
		}
	}
	if teteVersQueue {
		fmt.Println("Q")
	} else {
		fmt.Println("T")
	}
}

func insérer(ldc liste, val int, enTete bool) {
	// C'est là qu'on vous recommande vraiment de faire des dessins !
	cell := new(cellule)
	if enTete { // Insertion en tête
		cell.val, cell.prec, cell.suiv = val, ldc.tete, ldc.tete.suiv
		ldc.tete.suiv.prec = cell
		ldc.tete.suiv = cell
	} else { // ou en queue
		cell.val, cell.prec, cell.suiv = val, ldc.queue.prec, ldc.queue
		ldc.queue.prec.suiv = cell
		ldc.queue.prec = cell
	}
}

// prérequis : ldc n'est pas vide
func extraire(ldc liste, enTete bool) int {
	var val int
	if enTete { // Extraction en tête
		val = ldc.tete.suiv.val
		deux := ldc.tete.suiv.suiv
		deux.prec = ldc.tete
		ldc.tete.suiv = deux
	} else { // ou en queue
		val = ldc.queue.prec.val
		deux := ldc.queue.prec.prec
		deux.suiv = ldc.queue
		ldc.queue.prec = deux
	}
	return val
}

// Prérequis : ptr et ptr.suiv désignent des cellules significatives (ni tête ni queue ni queue.prec)
func échanger(ptr *cellule) *cellule {
	// L'idée est de commencer par stocker les pointeurs vers les 4 cellules impliquées dans l'échange, pour être sûr de n'en perdre aucune
	prec := ptr.prec
	cour := ptr
	suiv := ptr.suiv
	suivDuSuiv := ptr.suiv.suiv
	// Une fois qu'on a fait ça, on peut faire un dessin pour voir qu'il y a en fait 6 chaînages à affecter
	prec.suiv = suiv
	cour.prec = suiv
	cour.suiv = suivDuSuiv
	suiv.prec = prec
	suiv.suiv = cour
	suivDuSuiv.prec = cour
	// En faisant comme ci-dessus, on n'a pas besoin de se préoccuper de l'ordre dans lequel on écrase les chaînages existants : on fait comme si les cellules n'étaient pas déjà chaînées
	return suiv
}

func trierNain(ldc liste) {
	if (ldc.tete.suiv == ldc.queue) || (ldc.tete.suiv.suiv == ldc.queue) { // liste vide ou singleton
		return
	}
	for cour := ldc.tete.suiv; cour.suiv != ldc.queue; {
		if cour.val > cour.suiv.val {
			cour = échanger(cour) // attention : si vous ne mettez pas à jour cour avec la valeur renvoyée, ça ne marchera pas !
			if cour.prec != ldc.tete {
				cour = cour.prec
			}
		} else {
			cour = cour.suiv
		}
	}
}

func testVisuel() {
	for taille := 0; taille <= tailleMax; taille++ {
		fmt.Println("***** Liste de taille", taille, "*****")
		ldc := créer()
		for i := 0; i < taille; i++ {
			val := rand.Intn(valSup)
			enTete := rand.Intn(2) == 1
			insérer(ldc, val, enTete)
			var ch string
			if enTete {
				ch = "au début"
			} else {
				ch = "à la fin"
			}
			fmt.Println("Insertion de la valeur", val, ch, "de la liste :")
			afficher(ldc, true)
			afficher(ldc, false)
		}
		trierNain(ldc)
		fmt.Println("Liste triée :")
		afficher(ldc, true)
		if taille > 0 {
			for i := 0; i < taille; i++ {
				enTete := rand.Intn(2) == 1
				val := extraire(ldc, enTete)
				var ch string
				if enTete {
					ch = "au début"
				} else {
					ch = "à la fin"
				}
				fmt.Println("Extraction de la valeur", val, ch, "de la liste :")
				afficher(ldc, true)
			}
		}
		fmt.Println()
	}
}

type xile struct {
	nbrElem int
	vals    liste
}

func testXile(pile bool) {
	var donnees xile
	// Il n'y a pas de différence entre une liste chaînée gérée comme une pile ou une file : c'est la même structure en dessous, c'est juste qu'on ne fait pas la même chose avec
	donnees.vals = créer()
	var nomStruct string
	if pile {
		nomStruct = "Pile"
	} else {
		nomStruct = "File"
	}
	fmt.Printf("\n%s bornée : entrez les valeur une par une :\n", nomStruct)
	fmt.Println("- les chiffres dans [0..9] seront insérés")
	fmt.Println("- -1 signifie retirer une valeur")
	fmt.Println("- toutes les autres valeurs seront ignorées")
	fmt.Println("- ctrl-d pour terminer")
	lecteur := bufio.NewScanner(os.Stdin)
	fmt.Print(nomStruct, " de ", donnees.nbrElem, " éléments : ")
	afficher(donnees.vals, true)
	for lecteur.Scan() {
		val, err := strconv.Atoi(lecteur.Text())
		if err != nil || val < -1 || val >= valSup {
			fmt.Fprintln(os.Stderr, "Erreur : valeur incorrecte !")
			continue
		}
		if val == -1 {
			if donnees.nbrElem > 0 {
				donnees.nbrElem--
				val = extraire(donnees.vals, !pile)
				fmt.Println("-> valeur retirée :", val)
			} else {
				fmt.Fprintln(os.Stderr, fmt.Sprint(nomStruct, " vide !"))
				continue
			}
		} else {
			donnees.nbrElem++
			insérer(donnees.vals, val, !pile)
		}
		fmt.Print(nomStruct, " de ", donnees.nbrElem, " éléments : ")
		afficher(donnees.vals, true)
	}
	fmt.Println()
}

func main() {
	testVisuel()
	testXile(true)
	testXile(false)
}
