#ifndef UTIL_H
#define UTIL_H

/*
 * prints m, prefixed with 'err: ', exists with code EXIT_FAILURE
 */
void util_print_err(char *m);

/*
 * prints m prefixed with 'wrn: '
 */
void util_print_wrn(char *m);

/*
 * prints m prefixed with 'inf: '
 */
void util_print_inf(char *m);

/*
 * prints m
 */
void util_print(char *m);

#endif
