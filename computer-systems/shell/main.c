#include <stdio.h>
#include <stdlib.h>

#define PROMPT "$> "
#define BYE "BYEEEEEEEE"
#define MAXCHAR 100

int main(int argc, char **argv)
{
  while (1)
  {
    char input[MAXCHAR] = "";
    char ch;
    printf("%s", PROMPT);
    int i = 0;
    while ((ch = fgetc(stdin)))
    {
      if (ch == EOF)
      {
        printf("^D\n%s\n", BYE);
        exit(0);
      }
      else if (ch == '\n')
      {
        break;
      }
      input[i] = ch;
      i++;
    }
    printf("%s\n", input);
  };
  return 0;
};