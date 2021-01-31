// Exercise 1 - 13. Write a program to print a histogram of the lengths of words in its input.It is easy to draw the histogram with the bars horizontal;
// a vertical orientation is more challenging.

#include <stdio.h>

#define MAX_WORD_LEN 10

int is_blank(char c);

int main()
{
  static int wf[MAX_WORD_LEN] = {0};
  // Calculate word freq
  int c;
  int cur_word_len = 0;
  int words_gt = 0;
  char prev_char = '\0';
  while ((c = getchar()) && c != EOF)
  {
    if (is_blank(c) && !is_blank(prev_char))
    {
      if (cur_word_len >= MAX_WORD_LEN)
      {
        words_gt += 1;
      }
      else
      {
        wf[cur_word_len] += 1;
      }
      cur_word_len = 0;
    }
    else if (!is_blank(c))
    {
      cur_word_len += 1;
    }
    prev_char = c;
  }

  // Print horizontal hist
  printf("Horizontal WF Histogram\n\n");
  int i, j;
  for (i = 0; i < MAX_WORD_LEN; i++)
  {
    printf("%5d:", i);
    for (j = 0; j < wf[i]; j++)
    {
      printf("|");
    }
    printf("%d\n", j);
  }
  printf("  >%d:", MAX_WORD_LEN);
  for (i = 0; i < words_gt; i++)
  {
    printf("|");
  }

  printf("%d\n\n", i);

  // Print vertical hist
  // Strategy: Columns + 1 for each character. Figure out max height and work downward from there.
  printf("Vertical WF Histogram\n\n");
  int col, row;
  int max = words_gt;
  // Get max
  for (size_t i = 0; i < MAX_WORD_LEN; i++)
  {
    if (wf[i] > max)
    {
      max = wf[i];
    }
  }
  // Print out histogram
  for (row = max; row >= 0; row--)
  {
    for (col = 0; col <= MAX_WORD_LEN; col++)
    {
      int total = col < MAX_WORD_LEN ? wf[col] : words_gt;
      if (total >= row)
      {
        printf(" -- ");
      }
      else
      {
        printf("    ");
      }
    }
    printf("\n");
  }
  for (size_t i = 0; i <= MAX_WORD_LEN; i++)
  {
    printf(" %2d ", i);
  }
}

int is_blank(char c)
{
  return c == ' ' || c == '\t' || c == '\n' || c == '\0';
}