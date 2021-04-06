#include "vec.h"

data_t dotproduct(vec_ptr u, vec_ptr v)
{
   data_t sum = 0, u_val, v_val;

   long len = vec_length(u); // moved this out of the loop so it's evaluated once, did not affect perf in instruments

   // allow direct data access, removing two procedure calls
   data_t *ustart = get_vec_start(u);
   data_t *vstart = get_vec_start(v);

   for (long i = 0; i < len; i++)
   {                     // we can assume both vectors are same length
      u_val = ustart[i]; // get_vec_element(u, i, &u_val);
      v_val = vstart[i]; // get_vec_element(v, i, &v_val);
      sum += u_val * v_val;
   }
   return sum;
}
