#include "vendor/unity.h"

#include "vec.h"
#include <time.h>

extern data_t dotproduct(vec_ptr, vec_ptr);

void setUp(void)
{
}

void tearDown(void)
{
}

void test_empty(void)
{
  vec_ptr u = new_vec(0);
  vec_ptr v = new_vec(0);

  TEST_ASSERT_EQUAL(0, dotproduct(u, v));

  free_vec(u);
  free_vec(v);
}

void test_basic(void)
{
  vec_ptr u = new_vec(3);
  vec_ptr v = new_vec(3);

  set_vec_element(u, 0, 1);
  set_vec_element(u, 1, 2);
  set_vec_element(u, 2, 3);
  set_vec_element(v, 0, 4);
  set_vec_element(v, 1, 5);
  set_vec_element(v, 2, 6);

  TEST_ASSERT_EQUAL(32, dotproduct(u, v));

  free_vec(u);
  free_vec(v);
}

void test_longer(void)
{
  long n = 1000000;
  vec_ptr u = new_vec(n);
  vec_ptr v = new_vec(n);

  for (long i = 0; i < n; i++)
  {
    set_vec_element(u, i, i + 1);
    set_vec_element(v, i, i + 1);
  }

  long expected = (2 * n * n * n + 3 * n * n + n) / 6;
  TEST_ASSERT_EQUAL(expected, dotproduct(u, v));

  free_vec(u);
  free_vec(v);
}

int main(void)
{
  clock_t test_start, test_end;
  double clocks_elapsed, time_elapsed;

  UNITY_BEGIN();

  test_start = clock();
  RUN_TEST(test_empty);
  test_end = clock();
  clocks_elapsed = test_end - test_start;
  time_elapsed = clocks_elapsed / CLOCKS_PER_SEC;
  printf("%.2fs to run test_empty (%.2fns per test)\n", time_elapsed, time_elapsed * 1e9);

  test_start = clock();
  RUN_TEST(test_basic);
  test_end = clock();
  clocks_elapsed = test_end - test_start;
  time_elapsed = clocks_elapsed / CLOCKS_PER_SEC;
  printf("%.2fs to run test_start (%.2fns per test)\n", time_elapsed, time_elapsed * 1e9);

  test_start = clock();
  RUN_TEST(test_longer);
  test_end = clock();
  clocks_elapsed = test_end - test_start;
  time_elapsed = clocks_elapsed / CLOCKS_PER_SEC;
  printf("%.2fs to run test_longer (%.2fns per test)\n", time_elapsed, time_elapsed * 1e9);

  return UNITY_END();
}
