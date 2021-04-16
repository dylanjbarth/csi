#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/wait.h>
#include <string.h>
#include <signal.h>

#define SHELL "turtlsh"
#define PROMPT "$> "
#define BYE "Remember, it's turtles all the way down."
#define MAXCHAR 100
#define MAXARGS 5

#define BUILTIN_EXIT "exit"
#define BUILTIN_CD "cd"

char *ascii_turtle = "\
  _____     ____\n\
 /      \\  |  o |\n\
|        |/ ___\\|\n\
|_________/     \n\
|_|_| |_|_|\n\
Welcome to turtlsh!\n";
char *utf8_turtle = "\xF0\x9F\x90\xA2";
void split(char *str, char **output);
void bye();
void interrupt_handler(int sig);
int cpid = 0;

int main(int argc, char **argv)
{
  printf("%s", ascii_turtle);
  signal(SIGINT, interrupt_handler);
  while (1)
  {
    char input[MAXCHAR] = "";
    char ch;
    printf("%s $>", utf8_turtle);
    int i = 0;
    while ((ch = fgetc(stdin)))
    {
      if (ch == EOF)
      {
        printf("^D");
        bye();
      }
      else if (ch == '\n')
      {
        break;
      }
      input[i] = ch;
      i++;
    }
    cpid = fork();
    if (cpid > 0)
    {
      // parent process, wait for child to complete then reap
      int cstat;
      waitpid(cpid, &cstat, WUNTRACED);
      cpid = 0;
    }
    else
    {
      // split strings on whitespace into char array
      // exec if command else command not found
      char **args = malloc(MAXARGS * sizeof(char *));
      split(input, args);
      // check for builtins
      if (strcmp(args[0], BUILTIN_CD) == 0)
      {
        int err = chdir(args[1]);
        if (err)
        {
          printf("-%s: cd: %s: No such file or directory\n", SHELL, args[1]);
        }
      }
      else if (strcmp(args[0], BUILTIN_EXIT) == 0)
      {
        bye();
      }
      else
      {
        // assume not builtin
        int err = execvp(args[0], args);
        if (err < 0)
        {
          printf("'%s': command not found.\n", args[0]);
        }
      }
      exit(0);
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

// handle a sigint
void interrupt_handler(int sig)
{
  if (cpid != 0)
  {
    kill(cpid, SIGKILL);
  }
  else
  {
    bye();
  }
  cpid = 0;
}

void bye()
{
  printf("\n%s\n", BYE);
  exit(0);
}
