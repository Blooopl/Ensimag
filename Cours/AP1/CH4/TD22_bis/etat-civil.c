#include <stdint.h>
#include <stdbool.h>
#include <stdio.h>

typedef char *chaine_t;


typedef enum {
    JANVIER = 1, FEVRIER, MARS, AVRIL, MAI, JUIN, JUILLET, AOUT, SEPTEMBRE, OCTOBRE, NOVEMBRE, DECEMBRE
} mois_t;

typedef struct {
    uint8_t jour;
    mois_t mois;
    int annee;
} date_t;

typedef struct {
    chaine_t nomPrenom;
    date_t naissance;
    double taille;
    bool motard;
} infos_t;


typedef struct {
    int entiere;
    int fract;
} resultat_modf_t;

resultat_modf_t ma_modf(double val) {
    resultat_modf_t res;
    res.entiere = (int)val;
    res.fract = (int)((val - res.entiere) * 100);
    return res;
}

void afficher(infos_t moi){
    printf("Je m'apelle %s \n", moi.nomPrenom);
    printf("Je suis né le %d/%d/%d \n", moi.naissance.jour, moi.naissance.mois, moi.naissance.annee);
    printf("Je mesure %g \n", moi.taille);

    if (moi.motard){
        printf("GAAAAAAZ \n ");
    }else{
        printf("C'est quoi une moto ? \n");
    }
}

int main(void) {
    infos_t eddy = {"Édouard Bracame", {15, JUIN, 1950}, 1.78, true};
    infos_t obelix = {.taille = 1.83, .naissance = {4, SEPTEMBRE, -70}, .nomPrenom = "Obelix" };
    obelix.motard = false;
    afficher(eddy);
    puts("");
    afficher(obelix);
}