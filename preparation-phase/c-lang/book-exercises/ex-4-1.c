// Exercise 4 - 1. Write the function strrindex(s, t), which returns the position of the rightmost occurrence of t in s, or −1 if there is none.

#include <stdio.h>
#include <string.h>
#include <ctype.h>

int strindex(char s[], char t[]);

int main()
{
  printf("%s %d\n", "help", strindex("help", "e"));
  printf("%s %d\n", "yee", strindex("yee", "k"));
  char test1[] = "this is a bigger one";
  printf("%s %d\n", test1, strindex(test1, "igge"));
}

int strindex(char s[], char t[])
{
  int rindex = -1;
  int i, j, k;
  for (i = 0; s[i] != '\0'; i++)
  {
    for (j = i, k = 0; t[k] != '\0' && s[j] == t[k]; j++, k++)
    {
      // checking to see if we get past 0 index on t
    }
    if (k > 0 && t[k] == '\0')
    {
      rindex = i;
    }
  }
  return rindex;
};
