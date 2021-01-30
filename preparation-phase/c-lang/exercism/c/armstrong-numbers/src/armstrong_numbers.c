#include "armstrong_numbers.h"

bool is_armstrong_number(int candidate)
{
  int n_dig = num_digits(candidate);
  int counter = candidate;
  int sum = 0;
  while (counter)
  {
    int digit = counter % 10;
    sum += power(digit, n_dig);
    counter /= 10;
  }
  return sum == candidate;
}

int num_digits(int n)
{
  int count = 0;
  while (n != 0)
  {
    n /= 10;
    count++;
  }
  return count;
}

int power(int base, int exponent)
{
  int total = 1;
  for (int i = 0; i < exponent; i++)
  {
    total *= base;
  }
  return total;
}
