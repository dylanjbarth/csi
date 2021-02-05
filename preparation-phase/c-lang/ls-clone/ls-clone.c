#include <stdio.h>

#define PROGRAM "ls-clone"

int main(int argc, char *argv[])
{
  printf("Hello, world! Welcome to %s\n", PROGRAM);
  printf("N args %d\n", argc);
  while (--argc > 0)
  {
    printf("Arg %s\n", *++argv);
  }
}