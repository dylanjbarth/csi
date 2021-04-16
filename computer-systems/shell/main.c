#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/wait.h>
#include <string.h>

#define PROMPT "$> "
#define BYE "BYEEEEEEEE"
#define MAXCHAR 100
#define MAXARGS 5

void split(char *str, char **output);

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
    int cpid = fork();
    int cstat;
    if (cpid > 0)
    {
      // parent process, wait for child to complete then reap
      waitpid(cpid, &cstat, WUNTRACED);
    }
    else
    {
      // split strings on whitespace into char array
      // exec if command else command not found
      char **args = malloc(MAXARGS * sizeof(char *));
      split(input, args);
      int err = execvp(args[0], args);
      if (err < 0)
      {
        printf("Exec failed");
      }
    }
  };
  return 0;
};

void split(char *str, char **output)
{
  char *ptr;
  size_t i = 0;
  ptr = strtok(str, " ");
  while (ptr != NULL)
  {
    output[i] = ptr;
    ptr = strtok(NULL, " ");
    i++;
  }
}