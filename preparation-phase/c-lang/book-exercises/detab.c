// Exercise 1 - 20. Write a program detab that replaces tabs in the input with the proper number of blanks to space to the next tab stop.Assume a fixed set of tab stops, say every n columns.Should n be a variable or a symbolic parameter

#include <stdio.h>

#define COLUMNSIZE 10;

int main()
{
  int c;
  int n = COLUMNSIZE;
  while ((c = getchar()) != EOF)
  {
    switch (c)
    {
    case '\t':
      printf("\t");
      break;
    case '\b':
      printf("\b");
      break;
    case '\\':
      printf("\\");
      break;
    default:
      printf("%c", c);
      break;
    }
  }
}