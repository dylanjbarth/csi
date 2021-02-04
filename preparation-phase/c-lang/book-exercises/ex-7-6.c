// Exercise 7 - 6. Write a program to compare two files, printing the first line where they differ.
#include <stdlib.h>
#include <stdio.h>

int main(int argc, char *argv[])
{
  if (argc != 3)
  {
    printf("Must provide two files as input args.");
    exit(1);
  }
  FILE *fp1 = fopen(argv[1], "r");
  FILE *fp2 = fopen(argv[2], "r");
  char *line1, *line2;
  while (line1 = fgets(fp1))
}