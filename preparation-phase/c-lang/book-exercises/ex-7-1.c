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
  char *first_arg = argv[1];
  if (!first_arg)
  {
    printf("Must be called with %s or %s", UPPER, LOWER);
    exit(0);
  }
  if (strcmp(first_arg, UPPER) == 0)
  {
    while ((c = getchar()) != EOF)
    {
      putchar(toupper(c));
    }
  }

  else if (strcmp(first_arg, LOWER) == 0)
  {
    while ((c = getchar()) != EOF)
    {
      putchar(tolower(c));
    }
  }
}