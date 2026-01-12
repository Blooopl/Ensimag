#include <stdio.h>
#include <stdint.h>

int main(void) {
    int32_t tmp = 5; (void)tmp;

    int32_t val;
    int32_t *ptr = &val;
    printf("val == %d\n", *ptr);

    ptr++;
    printf("val == %d\n", *ptr);

    ptr = NULL;
    printf("val == %d\n", *ptr);

}
