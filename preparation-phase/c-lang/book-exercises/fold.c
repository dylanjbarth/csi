// Exercise 1 - 22. Write a program to “fold” long input lines into two or more shorter lines after the last non - blank character that occurs before the n - th column of input.Make sure your program does something intelligent with very long lines, and if there are no blanks or tabs before the specified column.

#include <stdio.h>

#define MAX_CHAR_PER_LINE 5;

int isBlank(char c);

int main()
{
  int c;
  int n = MAX_CHAR_PER_LINE;
  int lc = 0;
  while ((c = getchar()) != EOF)
  {
    if (lc > n && isBlank(c) == 1)
    {
      lc = 0;
      printf("\n");
    }
    else
    {
      printf("%c", c);
      ++lc;
    }
  }
}

int isBlank(char c)
{
  if (c == '\t' || c == '\n' || c == ' ')
  {
    return 1;
  }
  return 0;
}