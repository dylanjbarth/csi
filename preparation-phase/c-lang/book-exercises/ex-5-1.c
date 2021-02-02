// Exercise 5 - 1. As written, getint treats a + or -not followed by a digit as a valid representation of zero.Fix it to push such a character back on the input.

#include <ctype.h>
#include <stdio.h>

#define BUFSIZE 100
#define SIZE 10

int getch(void);
void ungetch(int);
char buf[BUFSIZE]; /* buffer for ungetch */
int bufp = 0;      /* next free position in buf */

int main()
{
  int n, array[SIZE], getint(int *);
  for (n = 0; n < SIZE && getint(&array[n]) != EOF; n++)
  {
    printf("pos %d value %d\n", n, array[n]);
  };
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

/* getint:  get next integer from input into *pn */
int getint(int *pn)
{
  int c, sign, lookahead;
  while (isspace(c = getch())) /* skip white space */
    ;
  if (!isdigit(c) && c != EOF && c != '+' && c != '-')
  {
    ungetch(c); /* it's not a number */
    return 0;
  }
  sign = (c == '-') ? -1 : 1;
  if (c == '+' || c == '-')
    c = getch();
  printf("1 Value of pn %d\n", *pn);
  for (*pn = 0; isdigit(c); c = getch())
    *pn = 10 * *pn + (c - '0');
  printf("2 Value of pn %d\n", *pn);
  *pn *= sign;
  printf("3 Value of pn %d\n", *pn);
  if (c != EOF)
    ungetch(c);
  return c;
}
