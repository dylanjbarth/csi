#include <stdio.h>

int main()
{
  int blanks, c;
  blanks = 0;
  while ((c = getchar()) != EOF)
  {
    if (c == '\n' || c == '\t' || c == ' ')
    {
      ++blanks;
    }
  }
  printf("So many blanks: %d", blanks);
}