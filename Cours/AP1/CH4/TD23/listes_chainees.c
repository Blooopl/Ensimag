#include <time.h>
#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

#define VAL_SUP 10
#define TAILLE_MAX 6

typedef struct _cell_t {
    int val;
    struct _cell_t *suiv;
} cellule_t;

typedef cellule_t *liste_t;

void afficher(liste_t lsc){
    while (lsc -> suiv != NULL){
        printf("%d",lsc -> val);
        printf("-> ");
        lsc = lsc -> suiv;
    }
    printf("FIN");
}

liste_t inserer_tete(liste_t lsc, int val){
    liste_t new = malloc(sizeof(liste_t));
    new -> val = val;
    new -> suiv = lsc;

    return new;
}

bool supprimer_premiere_occurence(liste_t *lsc, int val){
    
    if (*lsc == NULL){ //La liste est vide
        return false;
    }

    if ((*lsc)->val == val){ //La valeur est dans le première cellule.
        cellule_t *tmp = *lsc;
        *lsc = (*lsc)->suiv;
        free(tmp);
        return true;
    }


    cellule_t *prev = NULL;
    cellule_t *curr = *lsc;

    while(curr -> suiv != NULL){
        if (curr -> val == val){
           prev -> suiv = curr -> suiv;
           free(curr);
           return true;
        }else {
            prev = curr;
            curr = curr -> suiv;
        }
    }

    return false;
}




void main(){
    liste_t lsc = NULL;
    lsc = inserer_tete(lsc,1);
    lsc = inserer_tete(lsc,2);
    lsc = inserer_tete(lsc,3);
    lsc = inserer_tete(lsc,4);
    lsc = inserer_tete(lsc,5);
    lsc = inserer_tete(lsc,6);
    lsc = inserer_tete(lsc,7);
    lsc = inserer_tete(lsc,8);
    lsc = inserer_tete(lsc,9);
    lsc = inserer_tete(lsc,10);
    lsc = inserer_tete(lsc,11);
    lsc = inserer_tete(lsc,12);
    afficher(lsc);

    liste_t *pointeur;
    
    supprimer_premiere_occurence(&lsc,9);
    afficher(lsc);
}


