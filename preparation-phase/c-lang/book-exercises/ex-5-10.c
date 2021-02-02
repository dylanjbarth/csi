// Exercise 5 - 10. Write the program expr, which evaluates a reverse Polish expression from the command line, where each operator or operand is a separate argument.
// For example, expr 2 3 4 + * evaluates 2 × (3 + 4).

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>

#define MAX_ARGS 1000
#define PLUS '+'
#define MINUS '-'
#define TIMES 'x'

int isoperator(char a);

int main(int argc, char *argv[])
{
  printf("argc %d\n", argc);
  if (argc > MAX_ARGS)
  {
    printf("Max of %d args allowed", MAX_ARGS);
    return 0;
  }
  int next_idx, i, total = 0;
  int first_iter = 1;
  while (argv[i] != NULL)
  {
    if (isoperator(*argv[i]) && first_iter)
    {
      switch (*argv[i])
      {
      case PLUS:
        printf("Adding %s + %s\n", argv[i - 1], argv[i - 2]);
        total = atoi(argv[i - 1]) + atoi(argv[i - 2]);
        break;
      case MINUS:
        printf("Subtracting %s - %s\n", argv[i - 1], argv[i - 2]);
        total = atoi(argv[i - 1]) - atoi(argv[i - 2]);
        break;
      case TIMES:
        printf("Multiplying %s * %s\n", argv[i - 1], argv[i - 2]);
        total = atoi(argv[i - 1]) * atoi(argv[i - 2]);
        break;
      default:
        break;
      }
      next_idx = i - 3;
      first_iter = 0;
    }
    else if (isoperator(*argv[i]))
    {
      switch (*argv[i])
      {
      case PLUS:
        printf("Adding %d + %s\n", total, argv[next_idx]);
        total += atoi(argv[next_idx]);
        break;
      case MINUS:
        printf("Subtracting %d - %s\n", total, argv[next_idx]);
        total -= atoi(argv[next_idx]);
        break;
      case TIMES:
        printf("Multiplying %d * %s\n", total, argv[next_idx]);
        total *= atoi(argv[next_idx]);
        break;
      default:
        break;
      }
      next_idx -= 1;
    }
    i++;
  };
  printf("Result: %d", total);
}

int isoperator(char a)
{
  return a == PLUS || a == MINUS || a == TIMES;
}