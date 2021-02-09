#include "ls-clone.h"

int sort_lexagraphic(struct dirfile *df1, struct dirfile *df2)
{
  return strcmp(df2->filename, df1->filename);
}

int sort_size(struct dirfile *df1, struct dirfile *df2)
{
  long long res = df1->s.st_size - df2->s.st_size;
  if (res == 0)
  {
    // Interestingly, it looks like ls will fallback to lexagraphic sort if size is the same.
    return sort_lexagraphic(df1, df2);
  }
  return res;
}

int sort_no_op(struct dirfile *df1, struct dirfile *df2)
{
  // No-op is just a pass through
  return 1;
}
