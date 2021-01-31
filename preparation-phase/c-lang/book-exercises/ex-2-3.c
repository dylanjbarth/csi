// Exercise 2 - 1. Write a program to determine the ranges of char, short, int, and long variables, both signed and unsigned, by printing appropriate values from standard headers and by direct computation.Harder if you compute them : determine the ranges of the various floating - point types.

#include <stdio.h>
#include <string.h>
#include <ctype.h>

int htoi(char s[]);

int main()
{
  printf("5 %d\n", htoi("5"));
  printf("5E %d\n", htoi("5E"));
  printf("3E8 %d\n", htoi("3E8"));
}

int htoi(char s[])
{
  int n = 0, hexdigit;
  for (size_t i = 0; i < strlen(s); i++)
  {
    if (s[i] >= '0' && s[i] <= '9')
    {
      hexdigit = s[i] - '0';
      n = 16 * n + hexdigit;
    }
    else if (tolower(s[i]) >= 'a' && tolower(s[i]) <= 'f')
    {
      hexdigit = tolower(s[i]) - 'a' + 10;
      n = 16 * n + hexdigit;
    }
  }
  return n;
}