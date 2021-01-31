// Exercise 2 - 4. Write an alternate version of squeeze(s1, s2) that deletes each character in s1 that matches any character in the string s2.

#include <stdio.h>
#include <string.h>
#include <ctype.h>

void squeeze(char s1[], char s2[]);

int main()
{
  char s1[] = "Does this match...";
  char s2[] = "oea";
  squeeze(s1, s2);
  printf("Squeezed %s", s1);
}

void squeeze(char s1[], char s2[])
{
  int i, k;
  for (i = k = 0; s1[i] != '\0'; i++)
  {
    int match = 0;
    for (size_t j = 0; j < strlen(s2); j++)
    {
      if (s1[i] == s2[j])
      {
        printf("Caught a match %c\n", s1[i]);
        match = 1;
        break;
      }
    }
    if (!match)
    {
      s1[k] = s1[i];
      k++;
    }
  }
  s1[k] = '\0';
}