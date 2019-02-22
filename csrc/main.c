#include <stdlib.h>
#include <unistd.h>
#include <stdio.h>

#include "libgoc.h"

int
main(int argc, char **argv) {
  printf("%d\n", getpid());

  DebugGoGC();
  // void *m = GetGoMap();

  while (1) {
    sleep(1);
  }

  return 0;
}
