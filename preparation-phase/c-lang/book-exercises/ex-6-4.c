// Exercise 6 - 4. Write a program that prints the distinct words in its input sorted into decreasing order of frequency of occurrence. Precede each word by its count.

#include <stdio.h>
#include <ctype.h>
#include <string.h>
#include <stdlib.h>

#define MAXWORDS 1000
#define MAXWORD 100

#define BUFSIZE 100
char buf[BUFSIZE]; /* buffer for ungetch */
int bufp = 0;      /* next free position in buf */
int getch(void);
void ungetch(int c);
struct word_fr
{
  char *word;
  size_t count;
};

void find_or_add_word(char *w, struct word_fr *words);
int size = 0;
int custom_getword(char *, int);

int main()
{
  struct word_fr words[MAXWORDS];
  char word[MAXWORD];
  // printf("words before loop 0 %s", words[0]);
  // printf("words before loop 1 %s", words[1]);
  // printf("words before loop 2 %s", words[2]);
  // printf("words before loop 3 %s", words[3]);
  while (custom_getword(word, MAXWORD) != EOF)
  {
    if (isalpha(word[0]))
    {
      // printf("Evaluating word %s\n", word);
      find_or_add_word(word, words);
    }
  }
  printf("Printing final output\n");
  for (size_t i = 0; i < size; i++)
  {
    printf("Word: %s; Count: %lu\n", words[i].word, words[i].count);
  }

  // for each word, add it too the struct array if it doesn't exist and resort the list.
}

void find_or_add_word(char *w, struct word_fr *words)
{
  int found = 0;
  for (size_t i = 0; i < size; i++)
  {
    // printf("Iterating through struct array, pos: %lu; word: %s, count: %lu\n", i, words[i].word, words[i].count);
    if (size >= MAXWORDS)
    {
      printf("Error: Max word limit reached.");
      exit(0);
    }
    else if (strcmp(words[i].word, w) == 0)
    {
      // printf("Found %s at idx %lu. %s == %s\n", w, i, words[i].word, w);
      found = 1;
      words[i].count += 1;
      break;
    }
  }
  if (!found)
  {
    // printf("Adding %s at idx %d\n", w, size);
    words[size].word = malloc(sizeof w);
    strcpy(words[size].word, w);
    words[size].count = 1;
    size++;
  }
  found = 0;
}

int custom_getword(char *word, int lim)
{
  int c, getch(void);
  void ungetch(int);
  char *w = word;
  while (isspace(c = getch()))
    ;
  if (c != EOF)
    *w++ = c;
  if (!isalpha(c))
  {
    *w = '\0';
    return c;
  }
  for (; --lim > 0; w++)
    if (!isalnum(*w = getch()))
    {
      ungetch(*w);
      break;
    }
  *w = '\0';
  return word[0];
}

int getch(void) /* get a (possibly pushed back) character */
{
  return (bufp > 0) ? buf[--bufp] : getchar();
}
void ungetch(int c) /* push character back on input */
{
  if (bufp >= BUFSIZE)
    printf("ungetch: too many characters\n");
  else
    buf[bufp++] = c;
}