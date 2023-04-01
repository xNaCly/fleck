#include "util.h"
#include <stdio.h>
#include <stdlib.h>

void util_print_err(char *m) {
  printf("err: %s\n", m);
  exit(1);
}
void util_print_wrn(char *m) { printf("wrn: %s\n", m); }
void util_print_inf(char *m) { printf("inf: %s\n", m); }
void util_print(char *m) { printf("%s\n", m); }
