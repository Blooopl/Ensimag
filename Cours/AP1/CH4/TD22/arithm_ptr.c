#include <stdio.h>
#include <inttypes.h>

void test_tab(unsigned taille) {
    // Attention : on alloue le tableau dans la pile de la fonction test_tab, qui a une taille limitée (en général quelques méga-octets): ça ne marchera pas si la taille est grande !
    int32_t tab[taille];
    for (unsigned i = 0; i < taille; i++) {
        tab[i] = i * i;
    }

    printf("[ ");
    for (unsigned i = 0; i < taille; i++) {
        printf("%u ", *(tab + i));
    }
    puts("]");
    printf("Adresse de tab : %p\n", tab);
    for (int32_t *ptr = tab; ptr < tab + taille; ptr++) {
        printf("*%p == %u\n", ptr, *ptr);
    }
}

int main(void) {
    test_tab(10);
}