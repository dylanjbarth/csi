#include "vec.h"

data_t dotproduct(vec_ptr u, vec_ptr v)
{
   data_t sum = 0;

   long len = vec_length(u); // moved this out of the loop so it's evaluated once, did not affect perf in instruments

   // allow direct data access, removing two procedure calls
   data_t *ustart = get_vec_start(u);
   data_t *vstart = get_vec_start(v);

   for (long i = 0; i < len; i++)
   { // we can assume both vectors are same length
      sum += ustart[i] * vstart[i];
   }
   return sum;
}
