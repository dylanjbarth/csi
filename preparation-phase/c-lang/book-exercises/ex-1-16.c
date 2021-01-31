#include <stdio.h>

#define MAX_LINE 488

int custom_getline(char c[], int lim);
void custom_copy(char to[], char from[]);

int main()
{
  char curr_line[MAX_LINE];
  char longest_line[MAX_LINE];
  int max = 0;
  int len;
  int pmax = MAX_LINE;
  while ((len = custom_getline(curr_line, MAX_LINE)) > 0)
  {
    if (len > max)
    {
      max = len;
      custom_copy(longest_line, curr_line);
    }
  }
  printf("Longest line was %d characters long.\n", max);
  printf("Longest line: %s\n", longest_line);
  if (max > MAX_LINE)
  {
    printf("Please note, output was truncated because it was longer than %d.", pmax);
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

void custom_copy(char to[], char from[])
{
  int i = 0;
  while ((to[i] = from[i]) != '\0')
  {
    i++;
  }
}
