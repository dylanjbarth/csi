// Exercise 7 - 1. Write a program that converts upper case to lower or lower case to upper, depending on the name it is invoked with, as found in argv[0].
#include <stdio.h>
#include <string.h>
#include <ctype.h>
#include <stdlib.h>

#define UPPER "upper"
#define LOWER "lower"

int main(int argc, char *argv[])
{
  int c;
  // printf("%s", argv[0]);
  // printf("%s", argv[1]);
  // printf("%s", argv[1]);
  if (!argv[1])
  {
    printf("Must be called with %s or %s", UPPER, LOWER);
    exit(0);
  }
  if (strcmp(argv[1], UPPER) == 0)
  {
    while ((c = getchar()) != EOF)
    {
      putchar(toupper(c));
    }
  }

  else if (strcmp(argv[1], LOWER) == 0)
  {
    while ((c = getchar()) != EOF)
    {
      putchar(tolower(c));
    }
  }
}