#include "vec.h"

data_t dotproduct(vec_ptr u, vec_ptr v)
{
   data_t sum1 = 0, sum2 = 0;

   long len = vec_length(u); // moved this out of the loop so it's evaluated once, did not affect perf in instruments

   // allow direct data access, removing two procedure calls
   data_t *ustart = get_vec_start(u);
   data_t *vstart = get_vec_start(v);

   // 2x1 loop unrolling
   long i;
   for (i = 0; i < len; i += 2)
   { // we can assume both vectors are same length
      sum1 += ustart[i] * vstart[i];
      sum2 += ustart[i + 1] * vstart[i + 1];
   }
   // finish off the loop
   for (; i < len; i++)
   { // we can assume both vectors are same length
      sum1 += ustart[i] * vstart[i];
   }
   return sum1 + sum2;
}
