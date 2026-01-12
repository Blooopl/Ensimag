#include <ctype.h>
#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

#define TAILLE_TABLE 4

// Une cellule assez classique.
typedef struct _cell_t {
    char val;
    struct _cell_t *suiv;
} cell_t;

// Ici une liste est une cellule, PAS un pointeur vers une cellule.
// Pourquoi ? Parce que cette cellule sera notre élément fictif en tête, qu'on n'aura donc pas besoin d'allouer.
typedef struct {
    cell_t tete;
} liste_t;

typedef struct {
    unsigned nbr_elem;
    liste_t table[TAILLE_TABLE];
} tablehash_t;

unsigned hash(char c)
{
    return c % TAILLE_TABLE;
}

tablehash_t *creer(void)
{
    // En faisant ça, on met à 0 :
    // - le nbr_elem de la table
    // - son tableau table, qui contient lui-même TAILLE_TABLE liste_t, c'est-à-dire des cellules tete :
    // => on a donc initialisé tous les éléments fictifs en mettant leur champ val à 0 (pas indispensable en fait) et surtout leur chap suiv à NULL (qui vaut 0)
    return calloc(1, sizeof(tablehash_t));
}

void inserer(char val, tablehash_t *th)
{
    unsigned ix = hash(val);
    cell_t *cell = malloc(sizeof(cell_t));
    cell->val = val;
    cell->suiv = th->table[ix].tete.suiv;
    th->table[ix].tete.suiv = cell;
    th->nbr_elem++;
}

void afficher(tablehash_t *th)
{
    printf("Nombre d'elements : %u\n", th->nbr_elem);
    for (unsigned i = 0; i < TAILLE_TABLE; i++) {
        printf("[%u] : ", i);
        for (cell_t *cour = th->table[i].tete.suiv; cour != NULL; cour = cour->suiv) {
            printf("%c -> ", cour->val);
        }
        printf("FIN\n");
    }
}

void detruire(tablehash_t *th)
{
    for (unsigned i = 0; i < TAILLE_TABLE; i++) {
        for (cell_t *cour = th->table[i].tete.suiv; cour != NULL; ) {
            cell_t *suiv = cour->suiv;
            free(cour);
            cour = suiv;
        }
    }
    free(th);
}

typedef struct {
    FILE *fichier;
    char car;
} scanner_t;

scanner_t *demarrer(bool clavier, char *nom_fichier) {
    scanner_t *scanner = malloc(sizeof(scanner_t));
    if (clavier) {
        scanner->fichier = stdin;
    } else {
        scanner->fichier = fopen(nom_fichier, "r");
        if (scanner->fichier == NULL) {
            perror(nom_fichier);
            exit(2);
        }
    }
    return scanner;
}

bool avancer(scanner_t *scanner) {
    scanner->car = fgetc(scanner->fichier);
    return !feof(scanner->fichier);
}

char courant(scanner_t *scanner) {
    return scanner->car;
}

void terminer(scanner_t *scanner) {
    if (scanner->fichier != stdin) {
        fclose(scanner->fichier);
    }
    free(scanner);
}

void lire(bool clavier, char *nom_fichier, tablehash_t *th) {
    scanner_t *scanner = demarrer(clavier, nom_fichier);
    while (avancer(scanner)) {
        char c = toupper(courant(scanner));
        if ((c >= 'A') && (c <= 'Z')) {
            inserer(c, th);
        }
    }
    terminer(scanner);
}

int main(int argc, char *argv[]) {
    if (argc > 2) {
        fprintf(stderr, "usage : %s [fichier]\n", argv[0]);
        exit(1);
    }
    tablehash_t *th = creer();
    // La syntaxe "? :" (qui n'existe pas en Go) est un if condensé :
    // - si la condition avant le ? est vrai, on renvoie la valeur avant les :
    // - sinon on renvoie la valeur après les :
    lire(argc == 1, argc == 1 ? NULL : argv[1], th);
    afficher(th);
    detruire(th);
}
