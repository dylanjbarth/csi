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
      // TODO why does this not work when I've already typed something? I have to type it twice.
      // TODO might be nice to also handle the keyboard interrupt signal?
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