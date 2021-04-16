#include <stdio.h>

#define PROMPT "$> "
#define MAXCHAR 100

int main(int argc, char **argv)
{
  char input[MAXCHAR];
  while (1)
  {
    printf("%s", PROMPT);
    fgets(input, MAXCHAR, stdin);
    printf("%s", input);
  };
  return 0;
};