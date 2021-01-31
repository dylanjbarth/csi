#include <stdio.h>

main()
{
  int c;
  while ((c = getchar()) && c != 'q')
  {
    if (c == EOF)
    {
      printf("Character is EOF, represented as %d", c);
      break;
    }
    else
    {
      printf("%d", c);
    }
  }
}