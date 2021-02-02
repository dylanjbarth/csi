// Exercise 5 - 3. Write a pointer version of the function strcat that we showed in Chapter 2 : strcat(s, t) copies the string t to the end of s.
#include <stdio.h>

void custom_strcat(char *s, char *t);

int main()
{
  char str1[100] = "This is a string pointer";
  char str2[] = "and here's a second one";
  printf("Str1: %s\n", str1);
  printf("Str2: %s\n", str2);
  custom_strcat(str1, str2);
  printf("Cat: %s", str1);
}

void custom_strcat(char *s, char *t)
{
  while (*s) // find the end of s
  {
    s++;
  }
  while ((*s = *t) != '\0')
  {
    s++;
    t++;
  }
}