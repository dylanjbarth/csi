#include "circular_buffer.h"

circular_buffer_t *new_circular_buffer(int capacity)
{
  return &(circular_buffer_t){0, capacity, 0, 0};
}
int write(circular_buffer_t *cb, buffer_value_t v){};
int overwrite(circular_buffer_t *cb, buffer_value_t v);
int read(circular_buffer_t *cb, buffer_value_t *v);
void delete_buffer(circular_buffer_t *b);
void clear_buffer(circular_buffer_t *b);
