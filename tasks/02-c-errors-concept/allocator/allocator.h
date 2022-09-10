#include <errno.h>

extern int errno;

#define ADMIN 777
#define MIN_MEMORY_BLOCK 1024

void *allocate(int user_id, size_t size)
{
    if (user_id != ADMIN) {
        errno = EPERM;

        return NULL;
    } else if (size < 1024) {
        errno = EDOM;

        return NULL;
    }

    errno = 0;
    void *res = malloc(size);

    if (res == NULL) {
        errno = ENOMEM;

        return NULL;
    }

    return res;
}
