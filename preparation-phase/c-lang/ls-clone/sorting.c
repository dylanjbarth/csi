#include "ls-clone.h"

int sort_lexagraphic(struct dirfile *df1, struct dirfile *df2)
{
  int res = strcmp(df2->filename, df1->filename);
  return res;
}

int sort_size(struct dirfile *df1, struct dirfile *df2)
{
  return df1->s.st_size - df2->s.st_size;
}

int sort_no_op(struct dirfile *df1, struct dirfile *df2)
{
  // No-op is just a pass through
  return 1;
}
