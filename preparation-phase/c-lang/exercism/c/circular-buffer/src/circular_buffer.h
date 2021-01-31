#ifndef CIRCULAR_BUFFER_H
#define CIRCULAR_BUFFER_H

typedef int buffer_value_t;
typedef struct
{
  int *buffer_value_t;
  int max;
  int head;
  int tail;
} circular_buffer_t;

int write(circular_buffer_t *cb, buffer_value_t v);
int overwrite(circular_buffer_t *cb, buffer_value_t v);
int read(circular_buffer_t *cb, buffer_value_t *v);
circular_buffer_t *new_circular_buffer(int capacity);
void delete_buffer(circular_buffer_t *b);
void clear_buffer(circular_buffer_t *b);
#endif
