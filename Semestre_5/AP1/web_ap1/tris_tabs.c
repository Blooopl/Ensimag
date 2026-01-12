#include <time.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define VAL_SUP 10

#define TAILLE_MAX 6

void remplir(int tab[], unsigned taille) {
    for (unsigned i = 0; i < taille; i++) {
        tab[i] = random() % VAL_SUP;
    }
}

void echanger(int *a, int *b) {
    int tmp = *a;
    *a = *b;
    *b = tmp;
}

void afficher(int tab[], unsigned taille) {
    printf("[ ");
    for (unsigned i = 0; i < taille; i++) {
        printf("%d ", tab[i]);
    }
    puts("]");
}

void trier_nain(int tab[], int taille) {
    for (int i = 0; i < taille - 1; ) {
        if (tab[i] > tab[i+1]) {
            echanger(&tab[i], &tab[i + 1]);
            if (i > 0) {
                i--;
            }
        } else {
            i++;
        }
    }
}

void trier_max(int tab[], int taille) {
    for (int i = 0; i < taille - 1; i++) {
        unsigned ix_max = i;
        for (int ix = i + 1; ix < taille; ix++) {
            if (tab[ix] > tab[ix_max]) {
                ix_max = ix;
            }
        }
        echanger(&tab[i], &tab[ix_max]);
    }
}

void trier_ins(int tab[], unsigned taille) {
    for (unsigned i = 1; i < taille; i++) {
        int val = tab[i];
        int ix;
        for (ix = i - 1; (ix >= 0) && (tab[ix] > val); ix--) {
            tab[ix + 1] = tab[ix];
        }
        tab[ix + 1] = val;
    }
}

int main(void) {
    srandom(time(NULL));
    for (unsigned taille = 0; taille <= TAILLE_MAX; taille++) {
        printf("***** Tableau de taille %u *****\n", taille);
        int *init = malloc(TAILLE_MAX * sizeof(int)); // erreur très classique : oublier le "* sizeof(int)"
        remplir(init, taille);
        printf("Tableau initial            : ");
        afficher(init, taille);
        int *tab = malloc(TAILLE_MAX * sizeof(int));

        memcpy(tab, init, taille * sizeof(int)); // Là aussi : memcpy prend en paramètre le nombre d'octets à copier, pas le nombre d'éléments
        printf("Tableau trié par le nain   : ");
        trier_nain(tab, taille);
        afficher(tab, taille);

        memcpy(tab, init, taille * sizeof(int));
        printf("Tableau trié par sélection : ");
        trier_max(tab, taille);
        afficher(tab, taille);

        memcpy(tab, init, taille * sizeof(int));
        printf("Tableau trié par insertion : ");
        trier_ins(tab, taille);
        afficher(tab, taille);

        free(init);
        free(tab);
        puts("");
    }
}
