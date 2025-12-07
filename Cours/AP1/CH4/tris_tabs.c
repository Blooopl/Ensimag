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

void tri_nain(int tab[], unsigned taille){
    
}

int main(void) {
    srandom(time(NULL));
}