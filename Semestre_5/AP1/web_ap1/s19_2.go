package main

import (
	"fmt"
	"maps"
	"slices"
)

// Un ensemble est juste une map de booléen : un élement est dans l'ensemble ou n'y est pas.
type ensembleEntiers map[int]bool

func créerEnsemble() ensembleEntiers {
	return make(ensembleEntiers)
}

func insérer(val int, ens ensembleEntiers) {
	ens[val] = true
}

func supprimer(val int, ens ensembleEntiers) {
	delete(ens, val)
}

func afficher(ens ensembleEntiers, nom string) {
	fmt.Print(nom, " = { ")
	// En Go 1.22, il n'y a pas encore d'itérateur, on fait donc à la main :
	lesClés := make([]int, 0)
	for clé := range ens {
		lesClés = append(lesClés, clé)
	}
	slices.Sort(lesClés)
	for _, clé := range lesClés {
		fmt.Print(clé, " ")
	}
	// A partir de Go 1.23, on peut faire directement for _, val := range slices.Sorted(maps.Keys(ens)) sans passer par un slice !
	fmt.Println("}")
}

func appartient(val int, ens ensembleEntiers) bool {
	return ens[val]
}

func union(ens1, ens2 ensembleEntiers) ensembleEntiers {
	union := créerEnsemble()
	maps.Copy(union, ens1)
	// maps.Copy va écraser les éléments déjà présents, c'est ce qu'on veut
	maps.Copy(union, ens2)
	return union
}

func intersection(ens1, ens2 ensembleEntiers) ensembleEntiers {
	intersection := créerEnsemble()
	for val := range maps.Keys(ens1) {
		if appartient(val, ens2) {
			insérer(val, intersection)
		}
	}
	return intersection
}

func différence(ens1, ens2 ensembleEntiers) ensembleEntiers {
	différence := créerEnsemble()
	maps.Copy(différence, ens1)
	for val := range maps.Keys(ens2) {
		// Ce n'est pas forcément très efficace de dupliquer puis de supprimer, mais c'est simple !
		supprimer(val, différence)
	}
	return différence
}

func différenceSymétrique(ens1, ens2 ensembleEntiers) ensembleEntiers {
	union := union(ens1, ens2)
	for val := range intersection(ens1, ens2) {
		supprimer(val, union)
	}
	return union
}

func main() {
	ensA := créerEnsemble()
	for x := 1; x < 5; x++ {
		insérer(x, ensA)
	}
	afficher(ensA, "A")
	ensB := créerEnsemble()
	for x := 3; x < 7; x++ {
		insérer(x, ensB)
	}
	afficher(ensB, "B")
	afficher(union(ensA, ensB), "A | B")
	afficher(intersection(ensA, ensB), "A & B")
	afficher(différence(ensA, ensB), "A - B")
	afficher(différence(ensB, ensA), "B - A")
	afficher(différenceSymétrique(ensA, ensB), "A ^ B")
}
