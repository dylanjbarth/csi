// Exercise 7 - 6. Write a program to compare two files, printing the first line where they differ.
#include <stdlib.h>
#include <stdio.h>
#include <string.h>

#define MAXLINE 1000

int main(int argc, char *argv[])
{
  if (argc != 3)
  {
    printf("Must provide two files as input args.");
    exit(1);
  }
  FILE *fp1 = fopen(argv[1], "r");
  FILE *fp2 = fopen(argv[2], "r");
  char *line1 = malloc(MAXLINE * (sizeof(char)));
  char *line2 = malloc(MAXLINE * (sizeof(char)));
  int i = 0;
  while ((fgets(line1, MAXLINE, fp1) != NULL) && (fgets(line2, MAXLINE, fp2)) != NULL)
  {
    if (strcmp(line1, line2) != 0)
    {
      printf("Found differing line:\n");
      printf("File 1, line %d: %s\n", i, line1);
      printf("File 2, line %d: %s\n", i, line2);
      exit(0);
    }
    i++;
  }
  printf("These files are the same.\n");
  exit(0);
}