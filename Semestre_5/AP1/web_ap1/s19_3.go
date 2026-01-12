package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"unicode"
)

// Une fonction utilitaire qui va nous permettre de filtrer les caractères qui ne sont pas des lettres.
func toutSaufLettre(car rune) bool {
	return !unicode.IsLetter(car)
}

func lesMotsDuFichier(nomFichier string) []string {
	fich, err := os.Open(nomFichier)
	if err != nil {
		log.Fatal("erreur : impossible d'ouvrir le fichier")
	}
	defer fich.Close()
	lecteur := bufio.NewScanner(fich)
	var lesMots []string
	for lecteur.Scan() {
		lesMotsDeLaLigne := strings.FieldsFunc(lecteur.Text(), toutSaufLettre) // Cette fonction FieldsFunc prend une autre fonction en paramètre, voir la doc Go
		for _, mot := range lesMotsDeLaLigne {
			lesMots = append(lesMots, strings.ToLower(mot)) // on stocke les mots en minuscules au passage
		}
	}
	if len(lesMots) < 2 {
		log.Fatal("erreur : le fichier contient moins de deux mots")
	}
	return lesMots
}

// Deux mots qui se suivent dans le texte initial
type coupleDeMots struct {
	mot1 string
	mot2 string
}

// On construit le tableau des mots qui se suivent à partir du tableau des mots
func lesCouplesDeMots(lesMots []string) []coupleDeMots {
	var lesCouples []coupleDeMots
	// Le premier mot du texte n'a pas de prédécesseur, c'est le prédécesseur du deuxième mot du texte
	motPrec := lesMots[0]
	for _, motCour := range lesMots[1:] { // on parcours les mots à partir du 2ème (tranche de slice de la case 1 inclue jusqu'à la fin du slice)
		lesCouples = append(lesCouples, coupleDeMots{motPrec, motCour})
		motPrec = motCour // à chaque tour on avance d'un mot, donc le courant devient le précédent d'un mot qui sera le courant au tour de boucle suivant
	}
	return lesCouples
}

// Cette fonction sert à calculer les étiquettes des arêtes du graphe qui sont les compteurs du nombre de fois où mot1 est suivi par mot2 dans le texte initial
func lesNombresDeCouples(lesCouples []coupleDeMots) map[string]map[string]int {
	// Ceci est une map dont :
	// - les clés sont des chaînes de caractères (qui correspondent à tous les mot1)
	// - les valeurs sont des maps dont :
	// -- les clés sont des chaînes de caractères (qui correspondent à tous les mot2)
	// -- les valeurs sont des int (qui correspondent aux compteurs de couples mot1-mot2)
	lesNombres := make(map[string]map[string]int)
	for _, couple := range lesCouples {
		// Est-ce que j'ai déjà une arête allant de mot1 à un autre mot ?
		if nombreMot, trouve := lesNombres[couple.mot1]; trouve {
			// Si oui, je récupère le compteur d'arêtes entre mot1 et mot2 pour l'incrémenter : mais peut-être qu'il n'y avait pas encore d'arête entre mot1 et mot2...
			nombreMot[couple.mot2] += 1 // nombreMot[couple.mot2] renvoie 0 si couple.mot2 n'est pas dans la map, ça nous va bien, c'est le cas envisagé ci-dessus
		} else {
			// Si non, c'est la première fois que je vois mot1, donc là je dois créer la map qui va contenir les compteurs entre mot1 et tous ses mots suivants
			lesNombres[couple.mot1] = make(map[string]int)
			lesNombres[couple.mot1][couple.mot2] = 1 // et initialiser le compteur d'arête entre mot1 et mot2 à 1 parce que je viens d'en trouver une justement !
		}
	}
	return lesNombres
}

func génèreGraphe(nomFichier string, lesNombres map[string]map[string]int) {
	fich, err := os.Create(nomFichier) // le fichier sera écrasé s'il existe déjà
	if err != nil {
		log.Fatal(err)
	}
	defer fich.Close()
	// La syntaxe d'un fichier .gv est vraiment simple :
	// 1) on commence par écrire la chaine ci-dessous et on va à la ligne
	fmt.Fprintln(fich, "digraph {")
	// 2) puis on écrit une ligne par arête sous le format mot1 -> mot2 [label = compteur]
	for mot1, nbrs := range lesNombres {
		for mot2, nbr := range nbrs {
			fmt.Fprintf(fich, "%s -> %s [label = %d]\n", mot1, mot2, nbr)
		}

	}
	// 3) et on finit en fermant l'accolade
	fmt.Fprintln(fich, "}")
}

func main() {
	if len(os.Args) != 3 {
		log.Fatal(fmt.Errorf("usage : %s <fichier .txt> <fichier .gv>", path.Base(os.Args[0])))
	}
	génèreGraphe(os.Args[2], lesNombresDeCouples(lesCouplesDeMots(lesMotsDuFichier(os.Args[1]))))
}
