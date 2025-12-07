#include <math.h>
#include <stdio.h>

unsigned fact(unsigned n) {
    unsigned res;
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
}

int pgcd(unsigned a, unsigned b){
    if ((a == 0) || (b==0)){
        return -1;
    }

    for (; a!=b;) {

        if (a > b) {
            a = a - b;
        }else{
            b = b - a;
        }       
    }

    return a;
}

void test_pgcd(void){
    printf("Test pgcd %d",pgcd(0, 0));
    printf("Test pgcd %d",pgcd(1, 0));
    printf("Test pgcd %d",pgcd(15, 10));
    printf("Test pgcd %d",pgcd(10, 15));
}

int main(void) {
    printf("Test");
    test_fact();
    test_racine();
    test_pgcd();
}
