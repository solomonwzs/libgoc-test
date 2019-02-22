#include <stdlib.h>
#include <unistd.h>
#include <stdio.h>

#include "libgoc.h"

int
main(int argc, char **argv) {
  printf("%d\n", getpid());

  GoFoo();
  // void *m = GoBar();

  while (1) {
    sleep(1);
  }

  return 0;
}
