// Exercise 2 - 1. Write a program to determine the ranges of char, short, int, and long variables, both signed and unsigned, by printing appropriate values from standard headers and by direct computation.Harder if you compute them : determine the ranges of the various floating - point types.

#include <stdio.h>
#include <limits.h>
#include <float.h>

int main()
{
  printf("signed char min: %d; max: %d\n", SCHAR_MIN, SCHAR_MAX);
  printf("unsigned char min: %d; max: %d\n", 0, UCHAR_MAX);
  printf("signed short min: %d; max: %d\n", SHRT_MIN, SHRT_MAX);
  printf("unsigned short min: %d; max: %d\n", 0, USHRT_MAX);
  printf("signed int min: %d; max: %d\n", INT_MIN, INT_MAX);
  printf("unsigned int min: %d; max: %d\n", 0, UINT_MAX);
  printf("signed long min: %ld; max: %ld\n", LONG_MIN, LONG_MAX);
  printf("unsigned long min: %d; max: %lu\n", 0, ULONG_MAX);
}