package main

import "fmt"

func déclarationsDeBase() {
	fmt.Println("déclarationsDeBase :")
	// Les lignes ci-dessous sont toutes équivalentes : elles déclarent 3 variables locales avec une valeur de base.
	var a int = 1   // la façon la plus explicite de déclarer une variable
	b := 1          // la façon condensée : on laisse le compilateur inférer (deviner) que b est un int vue la valeur qu'on lui affecte.
	var c = uint(1) // une façon intermédiaire : en transtypant la constante 1 (int par défaut) en uint, on indique au compilateur que c est un uint.
	fmt.Println(a, b, c)
}

func déclarationsDeStruct() {
	fmt.Println("déclarationsDeStruct :")
	// On définit un type uneStruct qui est une structure : on le fait souvent en global, mais ça marche aussi en local si le type n'est utile que dans la fonction.
	type uneStruct struct {
		i int
		j float64
		k bool
	}
	// Maintenant on va déclarer une variable de type uneStruct et initialiser ses champs de plusieurs façons possibles.
	// La plus évidente est en fait fausse car Go ne sait pas ce qu'est "{2, 3.14}" : il faut lui expliquer que c'est uneStruct.
	// var s uneStruct = {2, 3.14, true}
	// On fera donc comme suit :
	var s uneStruct = uneStruct{2, 3.14, true}
	// La forme condensée marche aussi.
	t := uneStruct{2, 3.14, true}
	// Au passage, on peut ne pas initialiser tous les champs si on veut, ou les initialiser dans le désordre.
	u := uneStruct{k: true, i: 2} // j n'est pas initialisé, il vaut donc 0.0 par défaut
	// On peut aussi d'abord déclarer la variable sans initialiser les champs : tous les champs sont donc initialisés à la valeur nulle correspondant au type de chaque champ.
	var v uneStruct // v.i = 0, v.j = 0.0 et v.k = false
	// Puis initialiser les champs à la main.
	v.i, v.j, v.k = 2, 3.14, true
	// Autre syntaxe équivalente sous la forme condensée.
	w := uneStruct{} // les {} sont obligatoires : la syntaxe est comme ça !
	w.i, w.j, w.k = 2, 3.14, true
	fmt.Println(s, t, u, v, w)
}

func uneFonction() {
	fmt.Println("uneFonction :")
	// Là c'est simple, on affiche les valeurs de la variable globale x et de la constante globale y.
	fmt.Println(x)
	fmt.Println(y)
}

func uneAutreFonction() {
	fmt.Println("uneAutreFonction :")
	// Cette constante locale masque la variable globale x, le x "extérieur" n'est simplement plus accessible dans cette fonction.
	const x int = 7
	fmt.Println(x)
	// Même principe pour la constante globale y, qu'on masque ici avec une variable locale initialisée à 0 par défaut.
	var y int = 8
	fmt.Println(y)
	// On voit au passage que les constantes et les variables sont tellement traitées de façon similaire qu'on peut masquer une const par une var et vice-versa.
}

// Une variable globale : elle est visible partout dans le programme, y compris dans les fonctions écritent avant cette ligne.
// Si on ne met pas le "= 5", elle sera initialisée à 0 comme toutes les autres variables.
var x int = 5

// La syntaxe condensée n'est pas acceptée pour les variables globales, c'est une limitation du langage, donc la ligne ci-dessous est fausse.
// x := 5

// Par contre, on peut ne pas préciser le type si on veut qu'il l'infère tout seul : la ligne ci-dessous est correcte.
// var x = 5

// Une constante n'est en fait pas tellement différente d'une variable : on ne peut juste pas changer la valeur fixée à la déclaration.
// Attention on ne peut pas ne pas mettre le "= 0" en pensant que ça sera initialisé automatiquement, dans le cas d'une constante il faut une valeur explicite.
const y int = 0

func encoreUneFonction(y int) {
	fmt.Println("encoreUneFonction :")
	// La ligne ci-dessous est fausse : les paramètres sont considérés comme des variables locales et on ne peut pas déclarer deux variables portant le même nom dans le même bloc.
	// var y int = 2
	{
		// Par contre, là c'est juste, car on est dans un sous-bloc : on masque le paramètre avec la variable locale. Bouh, c'est piégeux, il ne faut pas faire ça !
		var y int = 2
		fmt.Println(y)
	}
	// Là, c'est bien le paramètre qu'on affiche, car le y local au sous-bloc n'existe pas en dehors du sous-bloc.
	fmt.Println(y)
}

func unPeuLaMêmeAvecUnIf(y int) {
	fmt.Println("unPeuLaMêmeAvecUnIf :")
	// Là on affiche le paramètre y, tout va bien.
	fmt.Println(y)
	if y := 9; y > 5 { // aïe, on masque le paramètre y avec une variable y locale (interne) au if : c'est MAL de faire ça ! Mais c'est tout à fait légal quand même...
		fmt.Println(y)
	}
	// Et là c'est de nouveau le paramètre qu'on affiche, parce que la variable locale au if n'existe plus.
	fmt.Println(y)
}

func main() {
	déclarationsDeBase()
	déclarationsDeStruct()
	uneFonction()
	uneAutreFonction()
	encoreUneFonction(4)
	unPeuLaMêmeAvecUnIf(4)
}
