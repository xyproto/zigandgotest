#include <stdio.h> // for printf
#include <stdlib.h> // for free

// implemented in add.zig
extern int add(int a, int b);

static void myprint(char *fmt, char* s) {
  printf(fmt, s);
}
