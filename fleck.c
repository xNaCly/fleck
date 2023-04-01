#include <stdio.h>
#include <stdlib.h>

#include "fleck.h"
#include "util.h"

int main(int argc, char **argv) {
  util_print("fleck - " FLECK_VERSION);
  if (argc < 2) {
    util_print_err("not enough arguments, exiting");
  }

  char *inFile = argv[1];
  if (inFile == NULL) {
    util_print_err("input file is null");
  }
  return EXIT_SUCCESS;
}
