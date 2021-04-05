#include <stdint.h>
#include <stdio.h>
#include <time.h>

#define TEST_LOOPS 10000000

uint64_t pagecount(uint64_t memory_size, uint8_t page_size)
{
  return memory_size >> page_size;
}

int main(int argc, char **argv)
{
  clock_t baseline_start, baseline_end, test_start, test_end;
  uint64_t memory_size, page_size;
  uint8_t ppower;
  double clocks_elapsed, time_elapsed;
  int i, ignore = 0;

  uint64_t msizes[] = {1L << 32, 1L << 40, 1L << 52};
  uint8_t psize_powers[] = {12, 16, 32};
  uint64_t psizes[] = {1L << psize_powers[0], 1L << psize_powers[1], 1L << psize_powers[2]};

  // this is snarky but since this is a contrived example we could just precompute this entirely
  // and get constant time lookup
  // uint64_t pagecounts[] = {msizes[0]/psizes[0], msizes[1]/psizes[1], msizes[2]/psizes[2]};

  baseline_start = clock();
  for (i = 0; i < TEST_LOOPS; i++)
  {
    memory_size = msizes[i % 3];
    page_size = psizes[i % 3];
    ignore += 1 + memory_size +
              page_size; // so that this loop isn't just optimized away
  }
  baseline_end = clock();

  test_start = clock();
  for (i = 0; i < TEST_LOOPS; i++)
  {
    memory_size = msizes[i % 3];
    ppower = psize_powers[i % 3];
    ignore += pagecount(memory_size, ppower) + memory_size + ppower;
  }
  test_end = clock();

  clocks_elapsed = test_end - test_start - (baseline_end - baseline_start);
  time_elapsed = clocks_elapsed / CLOCKS_PER_SEC;

  printf("%.2fs to run %d tests (%.2fns per test)\n", time_elapsed, TEST_LOOPS,
         time_elapsed * 1e9 / TEST_LOOPS);
  return ignore;
}
