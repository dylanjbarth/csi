#include <stdio.h>
#include <string.h>

#define MAX_LINE 488

int custom_getline(char c[], int lim);
void reverse(char str[], int len, char reversed[]);

int main()
{
  char curr_line[MAX_LINE];
  char reversed[MAX_LINE];
  int max = 0;
  int len;
  int pmax = MAX_LINE;
  while ((len = custom_getline(curr_line, MAX_LINE)) > 0)
  {
    printf("%s\n", curr_line);
    int tail = len > MAX_LINE ? MAX_LINE : len;
    reverse(curr_line, tail, reversed);
    printf("Reversed: %s\n", reversed);
  }
}

int custom_getline(char line[], int lim)
{
  int c;
  int i = 0;
  int limit_hit = 0;
  while ((c = getchar()) != EOF && c != '\n')
  {
    if (i < lim - 1)
    {
      line[i] = c;
    }
    else if (!limit_hit)
    {
      line[i] = '\0';
      limit_hit = 1;
    }
    i++;
  }
  if (c == '\n' && i < lim - 2)
  {
    line[i] = c;
    line[i + 1] = '\0';
  }
  return i;
}

void reverse(char str[], int len, char reversed[])
{
  int j = 0;
  // Don't copy the EOF
  for (int i = len - 2; i >= 0; i--)
  {
    reversed[j] = str[i];
    j++;
  }
  reversed[j] = '\0';
}