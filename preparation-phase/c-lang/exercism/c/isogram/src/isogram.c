#include <stdio.h>
#include <ctype.h>
#include <string.h>
#include "isogram.h"

// 26 characters in the alphabet
// multiple spaces and hypens are allowed
// strategy is to create an array to contain occurences of 26 letters of the alphabet.
bool is_isogram(const char phrase[])
{
  if (phrase[0] == '\0')
  {
    return false;
  }
  int freq[26] = {0};
  for (unsigned int i = 0; i < strlen(phrase); i++)
  {
    int idx = tolower(phrase[i]) - 'a';
    printf("The phrase: %c %d index: %d\n", phrase[i], phrase[i], phrase[i] - 'a');
    if (!freq[idx])
    {
      freq[idx] = 1;
    }
    else
    {
      freq[idx] += 1;
    }
    if (freq[idx] > 1)
    {
      printf("Freq %d", freq[idx]);
      return false;
    }
  }
  return true;
}
