// Exercise 1 - 20. Write a program detab that replaces tabs in the input with the proper number of blanks to space to the next tab stop.Assume a fixed set of tab stops, say every n columns.Should n be a variable or a symbolic parameter

#include <stdio.h>

#define COLUMNSIZE 10;

int main()
{
  int c;
  while ((c = getchar()) != EOF)
  {
    printf("%c%d", c, c);
  }
}