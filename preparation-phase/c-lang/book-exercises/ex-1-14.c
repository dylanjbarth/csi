// Exercise 1 - 14. Write a program to print a histogram of the frequencies of different characters in its input.

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
    printf("Evaluting: %c\n", c);
    if (is_blank(c) && !is_blank(prev_char))
    {
      printf("End of word, adding count to %d\n", cur_word_len);
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
      printf("Cur word length increasing %d\n", cur_word_len);
    }
    prev_char = c;
  }

  // Print hist
  int i, j;
  for (i = 0; i < MAX_WORD_LEN; i++)
  {
    printf("%2d:", i);
    for (j = 0; j < wf[i]; j++)
    {
      printf("|");
    }
    printf("\n");
  }
  printf("Words with length greater than %d: %d", MAX_WORD_LEN, words_gt);
}

int is_blank(char c)
{
  return c == ' ' || c == '\t' || c == '\n' || c == '\0';
}