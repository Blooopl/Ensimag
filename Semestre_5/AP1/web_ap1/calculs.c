#include <math.h>
#include <stdio.h>

unsigned fact(unsigned n) {
    unsigned res;
    // On rappelle qu'en C (comme en Go), les paramètres sont passés par copie : on peut donc modifier n comme une variable locale.
    for (res = 1; n > 1; n--) {
        res *= n;
    }
    return res;
}

void test_fact(void) {
    for (unsigned i = 0; i < 6; i++) {
        printf("%u! = %u\n", i, fact(i));
    }
    puts("");
}

double racine(double x) {
    if (x < 0.0) {
        return -NAN;
    }
    double res = 1.0;
    for (unsigned nbr_iter = 10; nbr_iter > 0; nbr_iter--) {
        res = (res + (x / res)) / 2.0;
    }
    return res;
}

void test_racine(void) {
    for (double val = -1.0; val < 10.0; val += 1.0) {
        printf("racine(%g) = %g ~ %g\n", val, racine(val), sqrt(val));
    }
    puts("");
}

int pgcd(unsigned a, unsigned b) {
    if ((a == 0) || (b == 0)) {
        return -1;
    }
    while (a != b) {
        if (a < b) {
            b -= a;
        } else {
            a -= b;
        }
    }
    return a;
}

void afficher_res_ou_erreur(unsigned a, unsigned b, int res) {
    // Si on ne met pas de '\n' à la fin d'une chaîne, il est possible qu'elle ne soit pas affichée immédiatement, mais ça dépend du système.
    // Pour forcer l'affichage, on utilise fflush sur stdout.
    printf("PGCD(%u, %u) ", a, b); fflush(stdout);
    if (res == -1) {
        fprintf(stderr, ": erreur, au moins des paramètres est nul !\n");
    } else {
        printf("= %d\n", res);
    }
}

void test_pgcd(void) {
    unsigned a, b;
    // En C, une affectation est une expression valant la valeur affectée : donc b = 0 vaut 0, qu'on affecte ensuite à a.
    a = b = 0;
    int res = pgcd(a, b);
    afficher_res_ou_erreur(a, b, res);
    b = 1;
    res = pgcd(a, b);
    afficher_res_ou_erreur(a, b, res);
    res = pgcd(b, a);
    afficher_res_ou_erreur(a, b, res);
    // Attention : a, b = 15, 10; n'est pas possible en C !
    a = 15;
    b = 10;
    res = pgcd(a, b);
    afficher_res_ou_erreur(a, b, res);
    res = pgcd(b, a);
    afficher_res_ou_erreur(a, b, res);
}

int main(void) {
    test_fact();
    test_racine();
    test_pgcd();
}
