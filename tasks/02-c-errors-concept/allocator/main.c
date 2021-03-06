#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "allocator.h"

int main()
{
    size_t size = 0;
    int uid = 0;

    errno = 0;
    if (scanf("%d %zu", &uid, &size) != 2) {
        perror("scanf failed");
        exit(1);
    }

    errno = 0;
    void *p = allocate(uid, size);
    if (p == NULL) {
        printf("allocation error: %s\n", strerror(errno));
    } else {
        printf("allocation was successful for %zu bytes\n", size);
        free(p);
    }

    return 0;
}
