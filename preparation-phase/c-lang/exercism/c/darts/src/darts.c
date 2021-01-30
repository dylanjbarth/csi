#include <stdio.h>
#include "darts.h"

// If the dart lands outside the target, player earns no points (0 points).
// If the dart lands in the outer circle of the target, player earns 1 point.
// If the dart lands in the middle circle of the target, player earns 5 points.
// If the dart lands in the inner circle of the target, player earns 10 points.
// The outer circle has a radius of 10 units (This is equivalent to the total radius for the entire target), the middle circle a radius of 5 units, and the inner circle a radius of 1. Of course, they are all centered to the same point (That is, the circles are concentric) defined by the coordinates (0, 0).

int score(coordinate_t pt)
{
  if (inside_circle_r(1.0, pt))
  {
    return 10;
  }
  else if (inside_circle_r(5.0, pt))
  {
    return 5;
  }
  else if (inside_circle_r(10.0, pt))
  {
    return 1;
  }
  return 0;
}

int inside_circle_r(float radius, coordinate_t pt)
{
  // https://stackoverflow.com/a/481150 Thanks
  // In general, x and y must satisfy (x - center_x)^2 + (y - center_y)^2 < radius^2.
  if ((pt[0] * pt[0]) + (pt[1] * pt[1]) <= radius * radius)
  {
    return 1;
  }
  return 0;
}
