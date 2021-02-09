#include "ls-clone.h"

// These sorting functions should return 0 if the strings are equal, positive if item 2 is greater than item 1 or negative if item 2 is less than item 1.
int sort_lexagraphic(struct dirfile *df1, struct dirfile *df2)
{
  return strcmp(df2->filename, df1->filename);
}

int sort_size(struct dirfile *df1, struct dirfile *df2)
{
  return df2->s.st_size - df1->s.st_size;
}

int sort_no_op(struct dirfile *df1, struct dirfile *df2)
{
  // No-op is just a pass through
  return 1;
}
