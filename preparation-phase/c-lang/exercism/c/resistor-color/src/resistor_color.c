#include "resistor_color.h"

int color_code(resistor_band_t v)
{
  return v;
}

int *colors()
{
  static int colors[10];
  for (int i = BLACK; i < WHITE + 1; i++)
  {
    colors[i] = i;
  }
  return colors;
}
