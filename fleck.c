#include <stdio.h>
#include <stdlib.h>

#include "fleck.h"

void die(char *msg) {
  printf("err: %s\n", msg);
  exit(EXIT_FAILURE);
}

int main(int argc, char **argv) {
  printf("fleck - %s\n", FLECK_VERSION);
  if (argc < 2) {
    die("not enough arguments");
  }

  char *inFile = argv[1];
  if (inFile == NULL) {
    die("input file is null");
  }

  return EXIT_SUCCESS;
}
