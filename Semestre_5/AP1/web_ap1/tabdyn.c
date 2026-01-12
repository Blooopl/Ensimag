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

tabdyn_t *creer(void) {
    tabdyn_t *tab = malloc(sizeof(tabdyn_t));
    tab->nbr_elem = 0;
    tab->capacite = CAPACITE_INIT;
    tab->valeurs = calloc(CAPACITE_INIT, sizeof(int));
    return tab;
}

void detruire(tabdyn_t *tab) {
    free(tab->valeurs);
    free(tab);
}

void afficher(tabdyn_t *tab) {
    printf("(%u/%u) [ ", tab->nbr_elem, tab->capacite);
    for (unsigned i = 0; i < tab->nbr_elem; i++) {
        printf("%d ", tab->valeurs[i]);
    }
    printf("]\n");
}

void verifier_capacite(tabdyn_t *tab) {
    if (tab->nbr_elem < tab->capacite) {
        return;
    }
    unsigned ancienne_capacite = tab->capacite;
    tab->capacite <<= 1; // décaler un entier d'un bit vers la gauche est équivalent à le multiplier par 2
    tab->valeurs = realloc(tab->valeurs, tab->capacite * sizeof(int));
    // realloc ne met pas à 0 la nouvelle zone
    bzero(tab->valeurs + ancienne_capacite, ancienne_capacite * sizeof(int));
    // inutile de libérer l'ancienne zone : realloc s'en charge s'il n'arrive pas à agrandir
}

void inserer(tabdyn_t *tab, unsigned idx, int val) {
    assert(idx <= tab->nbr_elem); // attention : c'est vraiment <= (et pas <) car insérer à len(tab) veut dire append en fait
    verifier_capacite(tab);
    memmove(tab->valeurs + idx + 1, tab->valeurs + idx, (tab->nbr_elem - idx) * sizeof(int));
    tab->valeurs[idx] = val;
    tab->nbr_elem++;
}

void ajouter(tabdyn_t *tab, int val) {
    inserer(tab, tab->nbr_elem, val);
}

int extraire(tabdyn_t *tab, unsigned idx) {
    assert(idx < tab->nbr_elem);
    int val = tab->valeurs[idx];
    tab->nbr_elem--;
    memmove(tab->valeurs + idx, tab->valeurs + idx + 1, (tab->nbr_elem - idx) * sizeof(int));
    return val;
}

void test_visuel(void) {
    tabdyn_t *tab;
    tab = creer();
    afficher(tab);
    for (unsigned i = 1; i <= 5; i+=2) {
        ajouter(tab, i);
        afficher(tab);
    }
    for (unsigned i = 0; i <= 6; i+=2) {
        inserer(tab, i, i);
        afficher(tab);
    }
    for (unsigned i = 0; i <= 3; i++) {
        printf("-> %d\n", extraire(tab, i));
        afficher(tab);
    }
    for (unsigned i = 0; i < 3; i++) {
        printf("-> %d\n", extraire(tab, i % 2));
        afficher(tab);
    }
    detruire(tab);
    puts("");
}

void remplir_tabdyn(tabdyn_t *tab, unsigned nbr) {
    for (unsigned i = 0; i < nbr; i++) {
        ajouter(tab, random() % VAL_SUP);
    }
}

void remplir_tabstat(int tab[], unsigned nbr) {
    for (unsigned i = 0; i < nbr; i++) {
        tab[i] = random() % VAL_SUP;
    }
}

void test_perf_remplir(void) {
    // On compare avec le remplissage d'un tableau de taille fixe
    int *tab_fixe = malloc(TAILLE_MAX_REMPLIR * sizeof(int));
    printf("Remplissage d'un tableau de taille fixe de %u éléments...\n", TAILLE_MAX_REMPLIR);
    clock_t debut = clock();
    remplir_tabstat(tab_fixe, TAILLE_MAX_REMPLIR);
    clock_t fin = clock();
    printf("=> temps d'execution = %f sec.\n", (double)(fin - debut) / CLOCKS_PER_SEC);
    free(tab_fixe);
    printf("Remplissage d'un tableau dynamique de %u éléments...\n", TAILLE_MAX_REMPLIR);
    tabdyn_t *tab = creer();
    debut = clock();
    remplir_tabdyn(tab, TAILLE_MAX_REMPLIR);
    fin = clock();
    detruire(tab);
    printf("=> temps d'execution = %f sec.\n", (double)(fin - debut) / CLOCKS_PER_SEC);
    puts("");
}

void trier(tabdyn_t *tab) {
    unsigned sup = tab->nbr_elem;
    bool modif;
    do {
        modif = false;
        for (unsigned i = 1; i < sup; i++) {
            if (tab->valeurs[i - 1] > tab->valeurs[i]) {
                int tmp = tab->valeurs[i - 1];
                tab->valeurs[i - 1] = tab->valeurs[i];
                tab->valeurs[i] = tmp;
                modif = true;
            }
        }
        sup--;
    } while (modif);
}

void trier_bete(tabdyn_t *tab) {
    unsigned sup = tab->nbr_elem;
    bool modif;
    do {
        modif = false;
        for (unsigned i = 1; i < sup; i++) {
            if (tab->valeurs[i - 1] > tab->valeurs[i]) {
                int tmp = extraire(tab, i - 1);
                inserer(tab, i, tmp);
                modif = true;
            }
        }
        sup--;
    } while (modif);
}

void test_perf_trier(void) {
    tabdyn_t *tab = creer();
    remplir_tabdyn(tab, TAILLE_MAX_TRIER);
    printf("Tri à bulles classique sur %u éléments...\n", TAILLE_MAX_TRIER);
    clock_t debut = clock();
    trier(tab);
    clock_t fin = clock();
    printf("=> temps d'execution = %f sec.\n", (double)(fin - debut) / CLOCKS_PER_SEC);
    remplir_tabdyn(tab, TAILLE_MAX_TRIER);
    printf("Tri à bulles très bête sur %u éléments...\n", TAILLE_MAX_TRIER);
    debut = clock();
    trier_bete(tab);
    fin = clock();
    printf("=> temps d'execution = %f sec.\n", (double)(fin - debut) / CLOCKS_PER_SEC);
    detruire(tab);
}

int main(void) {
   srandom(time(NULL));
   test_visuel();
   test_perf_remplir();
   test_perf_trier();
}
