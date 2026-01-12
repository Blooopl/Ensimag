#include <time.h>
#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

#define VAL_SUP 10
#define TAILLE_MAX 6

// En C, on doit utiliser un "nom temporaire" _cell_t pour pouvoir définir suiv commme un pointeur vers ce type (définition circulaire).
typedef struct _cell_t {
    int val;
    struct _cell_t *suiv;
} cellule_t; // Mais dans le code, on utilisera toujours cellule_t qui est le nom définitif du type cellule.

typedef cellule_t *liste_t;

void afficher(liste_t lsc) {
    while (lsc != NULL) {
        printf("%d -> ", lsc->val);
        lsc = lsc->suiv;
    }
    printf("FIN\n");
}

liste_t inserer_tete(liste_t lsc, int val) {
    liste_t cell = malloc(sizeof(cellule_t));
    cell->val = val;
    cell->suiv = lsc;
    return cell;
}

bool supprimer_premiere_occurrence(liste_t *lsc, int val) {
    // Ici on définit le fictif comme une variable locale, qui sera donc désallouée automatiquement à la fin de la fonction :
    // ça tombe bien, c'est exactement ce qu'on veut !
    cellule_t fictif;
    fictif.suiv = *lsc;
    liste_t prec = &fictif;
    // Pour une boucle sans corps, on n'est pas obligé de mettre les {} comme en Go, un ; suffit
    for (; (prec->suiv != NULL) && (prec->suiv->val != val); prec = prec->suiv);
    bool trouvee = prec->suiv != NULL;
    if (trouvee) {
        // Si on fait free tout de suite, on va être bien ennuyé pour la ligne suivante !
        liste_t a_detruire = prec->suiv;
        prec->suiv = prec->suiv->suiv;
        free(a_detruire);
    }
    *lsc = fictif.suiv;
    return trouvee;
}

void tester_ins_supp(void) {
    printf("--- Test de l'insertion en tête ---\n");
    liste_t lsc = NULL;
	printf("lsc = "); fflush(stdout); afficher(lsc);
	for (unsigned i = 0; i <= TAILLE_MAX; i++) {
		int val = random() % VAL_SUP;
        printf("+ val = %d\n", val);
		lsc = inserer_tete(lsc, val);
		printf("lsc = "); fflush(stdout); afficher(lsc);
	}
	puts("");
    printf("--- Test de la suppression ---\n");
	printf("lsc = "); fflush(stdout); afficher(lsc);
    while (lsc != NULL) {
		int val = random() % VAL_SUP;
        if (supprimer_premiere_occurrence(&lsc, val)) {
            printf("- val = %d\n", val);
    		printf("lsc = "); fflush(stdout); afficher(lsc);            
        }
    }
	puts("");
}

liste_t creer(unsigned taille) {
    liste_t lsc = NULL; // il faut VRAIMENT se souvenir qu'en C, les variables ne sont pas pré-inialisées, on DOIT affecter à NULL ici
    for (unsigned i = 0; i < taille; i++) {
        lsc = inserer_tete(lsc, random() % VAL_SUP);
    }
    return lsc;
}

void detruire(liste_t lsc) {
    while (lsc != NULL) {
        liste_t suiv = lsc->suiv;
        free(lsc);
        lsc = suiv;
    }
}

liste_t inverser(liste_t lsc) {
    liste_t res = NULL;
    while (lsc != NULL) {
        liste_t suiv = lsc->suiv;
        lsc->suiv = res;
        res = lsc;
        lsc = suiv;
    }
    return res;
}

liste_t trier_max(liste_t lsc) {
	cellule_t fictif;
	fictif.suiv = lsc;
	liste_t res = NULL;
	while (fictif.suiv != NULL) {
		liste_t prec_max = &fictif;
		liste_t prec = fictif.suiv;
		while (prec->suiv != NULL) {
			if (prec->suiv->val > prec_max->suiv->val) {
				prec_max = prec;
			}
			prec = prec->suiv;
		}
		liste_t suiv = prec_max->suiv->suiv;
		prec_max->suiv->suiv = res;
		res = prec_max->suiv;
		prec_max->suiv = suiv;
	}
	return res;    
}

liste_t trier_insertion(liste_t lsc) {
	cellule_t fictif;
    fictif.suiv = NULL;
	while (lsc != NULL) {
		liste_t suiv = lsc->suiv;
		liste_t prec = &fictif;
		for (; (prec->suiv != NULL) && (prec->suiv->val >= lsc->val); prec = prec->suiv);
		lsc->suiv = prec->suiv;
		prec->suiv = lsc;
		lsc = suiv;
	}
	return fictif.suiv;
}

void tester_reorganisation(char *texte, liste_t (*reorganiser)(liste_t)) {
    printf("--- Test %s --\n", texte);
    for (unsigned taille = 0; taille <= TAILLE_MAX; taille++) {
        liste_t lsc = creer(taille);
  		printf("lsc originale    = "); fflush(stdout); afficher(lsc);
        lsc = reorganiser(lsc);
  		printf("lsc réorganisée  = "); fflush(stdout); afficher(lsc);
        detruire(lsc);
    }
    puts("");
}

int main(void) {
    srandom(time(NULL));
    tester_ins_supp();
    tester_reorganisation("de l'inversion", inverser);
    tester_reorganisation("du tri par sélection du maximum (ordre croissant)", trier_max);
    tester_reorganisation("du tri par insertion (ordre décroissant)", trier_insertion);
}
