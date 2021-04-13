/*

Two different ways to loop over an array of arrays.

Spotted at:
http://stackoverflow.com/questions/9936132/why-does-the-order-of-the-loops-affect-performance-when-iterating-over-a-2d-arra

*/

#define DIM 10000

void option_one()
{
  int i, j;
  static int x[DIM][DIM];
  for (i = 0; i < DIM; i++)
  {
    for (j = 0; j < DIM; j++)
    {
      x[i][j] = i + j;
    }
  }
}

void option_two()
{
  int i, j;
  static int x[DIM][DIM];
  for (i = 0; i < DIM; i++)
  {
    for (j = 0; j < DIM; j++)
    {
      x[j][i] = i + j;
    }
  }
}

int main()
{
  option_one();
  // option_two();
  return 0;
}
