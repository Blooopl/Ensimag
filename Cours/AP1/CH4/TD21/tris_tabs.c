#include <time.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define VAL_SUP 10

#define TAILLE_MAX 20

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

void tri_ins(int tab[], unsigned taille){
    for (unsigned i = 1;i < taille; i++){
        for(unsigned j = i;j>0 && tab[j-1]>tab[j];j--){
            echanger(&tab[j],&tab[j-1]);
        }
    }
}

void tri_nain(int tab[],unsigned taille){

    for (unsigned id_nain = 0;id_nain != taille-1;){
        if (tab[id_nain+1]>=tab[id_nain]){
            id_nain++;
        }else{
            echanger(&tab[id_nain+1],&tab[id_nain]);
            if (id_nain != 0){
                id_nain--;
            }
        }
    }
}

void tri_max(int tab[],unsigned taille){
    for (unsigned i = 0;i<taille;i++){
        unsigned i_max = i;
        for (unsigned j=i;j!=taille;j++){
            if (tab[j]>tab[i_max]){
                i_max = j;
            }
        }
        echanger(&tab[i_max],&tab[i]);
    }
}

int main(void) {
    srandom(time(NULL));

    int *tab = malloc(TAILLE_MAX*sizeof(int));
    int *tab2 = malloc(TAILLE_MAX*sizeof(int));
    int *tab3 = malloc(TAILLE_MAX*sizeof(int));
    remplir(tab,8);
    remplir(tab2,8);
    remplir(tab3,8);

    afficher(tab,8);
    afficher(tab2,8);
    afficher(tab3,8);

    tri_nain(tab,8);
    tri_ins(tab2,8);
    tri_max(tab3,8);

    afficher(tab,8);
    afficher(tab2,8);
    afficher(tab3,8);
}

