#include <time.h>
#include <stdio.h>
#include <assert.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>

#define CAPACITE_INIT 1

#define TAILLE_MAX_TRIER 10000

#define TAILLE_MAX_REMPLIR (TAILLE_MAX_TRIER * 50000)

#define VAL_SUP 10

typedef struct {
    unsigned nbr_elem;
    unsigned capacite;
    int *valeurs;
} tabdyn_t;


tabdyn_t *creer(void){
    
    tabdyn_t *tab = malloc(sizeof(tabdyn_t));
    tab -> capacite = CAPACITE_INIT;
    tab -> nbr_elem = 0;
    int zone_valeurs = calloc(sizeof(int),CAPACITE_INIT);
    tab -> valeurs = zone_valeurs;

    return &tab;
}

void detruire(tabdyn_t *tab){
    free(tab->valeurs);
    free(tab);
}

void afficher_tab_stat(int tab[], unsigned taille) {
    printf("[ ");
    for (unsigned i = 0; i < taille; i++) {
        printf("%d ", tab[i]);
    }
    puts("]");
}

void afficher(tabdyn_t *tab){
    printf("(",tab->nbr_elem,"/",tab->capacite,")");
    afficher_tab_stat(tab->valeurs,tab->nbr_elem);
}

void verifier_capacite(tabdyn_t *tab){
    if (tab->capacite<=tab->nbr_elem){
        realloc(tab->valeurs,tab->capacite);
        tab->capacite = tab->capacite * 2;
    }
    bzero(tab->capacite+10,)
}