#include <stdio.h>
#include <stdint.h>
#include <stdbool.h>

typedef char *chaine_t;

// Si on écrivait la même chose sans le typedef, ça marcherait aussi mais on devrait écrire "enum mois_t" à chaque fois qu'on utilise le type mois_t.
typedef enum {
    JANVIER = 1, FEVRIER, MARS, AVRIL, MAI, JUIN, JUILLET, AOUT, SEPTEMBRE, OCTOBRE, NOVEMBRE, DECEMBRE
} mois_t;

// Même remarque que pour le enum ci-dessus.
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

void afficher(infos_t moi) {
    printf("Je m'appelle %s\n", moi.nomPrenom);
    printf("Je suis né le %02u/%02u/%d\n", moi.naissance.jour, moi.naissance.mois, moi.naissance.annee);
    resultat_modf_t taille = ma_modf(moi.taille);
    printf("Je mesure %d mètre %d\n", taille.entiere, taille.fract);
	if (moi.motard) {
		puts("Gaaaaaz !");
	} else {
		puts("C'est quoi une moto ?");
	}
}

int main(void) {
    // Quand on écrit {15, JUIN, 1950} le compilateur comprend tout seul qu'il s'agit d'une struct, sans qu'on ait besoin de préciser laquelle comme en Go :
    // le C est beaucoup moins fortement typé que Go, on peut facilement faire n'importe-quoi !
    infos_t eddy = {"Édouard Bracame", {15, JUIN, 1950}, 1.78, true};
    infos_t obelix = {.taille = 1.83, .naissance = {4, SEPTEMBRE, -70}, .nomPrenom = "Obelix" };
    obelix.motard = false;
    afficher(eddy);
    puts("");
    afficher(obelix);
}
