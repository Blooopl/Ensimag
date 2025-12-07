package main

import (
	"fmt"
	"slices"
)

type ensembleEntiers map[int]bool

func créerEnsemble() ensembleEntiers {
	ens := make(ensembleEntiers)
	return ens
}

func insérer(val int, ens ensembleEntiers) {
	ens[val] = true
}

func supprimer(val int, ens ensembleEntiers) {
	delete(ens, val)
}

func afficher(ens ensembleEntiers, nom string) {
	lesClés := make([]int, 0)
	for clé := range ens {
		lesClés = append(lesClés, clé)
	}
	slices.Sort(lesClés)
	fmt.Print(nom, " = ", "{ ")
	for _, valeur := range lesClés {
		if ens[valeur] {
			fmt.Print(valeur, " ")
		}
	}
	fmt.Println("}")
}

func appartient(val int, ens ensembleEntiers) bool {
	return ens[val]
}
func union(ens1, ens2 ensembleEntiers) ensembleEntiers {

	union := make(map[int]bool)

	for cle := range ens1 {
		insérer(cle, union)
	}

	for cle := range ens2 {
		insérer(cle, union)
	}

	return union
}

func intersection(ens1, ens2 ensembleEntiers) ensembleEntiers {
	inter := make(map[int]bool)

	for cle := range ens1 {
		if appartient(cle, ens2) {
			insérer(cle, inter)
		}
	}

	return inter
}

func différence(ens1, ens2 ensembleEntiers) ensembleEntiers {
	diff := union(ens1, ens2)

	for cle := range ens2 {
		if appartient(cle, diff) {
			supprimer(cle, diff)
		}
	}
	return diff
}

func différenceSymétrique(ens1, ens2 ensembleEntiers) ensembleEntiers {
	return différence(union(ens1, ens2), intersection(ens1, ens2))
}

func main() {
	A := créerEnsemble()
	B := créerEnsemble()
	insérer(1, A)
	insérer(2, A)
	insérer(3, A)
	insérer(4, A)
	afficher(A, "A")

	insérer(3, B)
	insérer(4, B)
	insérer(5, B)
	insérer(6, B)
	afficher(B, "B")

	afficher(union(A, B), "A | B")
	afficher(intersection(A, B), "A & B")
	afficher(différence(A, B), "A - B")
	afficher(différence(B, A), "B - A")
	afficher(différenceSymétrique(A, B), "A ^ B")
}
