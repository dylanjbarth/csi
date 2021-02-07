#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <dirent.h>
#include <sys/stat.h>
#include <sys/ioctl.h>
#include <unistd.h>

#define PROGRAM "ls-clone"
#define COLGUTTER 5
#define MAX_DIRENT 65535  // https://stackoverflow.com/a/466596 ?
#define MAX_FILENAME 1024 // meh https://www.systutorials.com/maximum-allowed-file-path-length-for-c-programming-on-linux/

struct dirfile
{
  char filename[MAX_FILENAME];
  struct stat s;
};
struct dirfile *entries[MAX_DIRENT];

int is_file(struct stat *f);
int is_dir(struct stat *f);
void print_dir_or_file(char *s);
void print_file(struct dirfile *d, int colsize);
void print_dir(char *dir);
int make_dirfile(char *filename, char *fullpath, struct dirfile *df);
void shift_entries(int idx, int curr_len);
void insert_to_entries(struct dirfile *f, int entries_len, int (*fn)(struct dirfile *df1, struct dirfile *df2));
int entries_bin_search(int low, int high, struct dirfile *f, int (*compare)(struct dirfile *df1, struct dirfile *df2));
int compare_lexagraphic(struct dirfile *df1, struct dirfile *df2);

int win_cols;

int main(int argc, char *argv[])
{
  // Grab window size
  struct winsize w;
  ioctl(STDOUT_FILENO, TIOCGWINSZ, &w);
  win_cols = w.ws_col;
  if (argc > 1)
  {
    while (--argc > 0)
    {
      char *cur_arg = *++argv;
      print_dir_or_file(cur_arg);
    }
  }
  else
  {
    print_dir_or_file(".");
  }
}

void print_dir_or_file(char *dir_or_file)
{
  struct stat statbuf;
  int stat_result = stat(dir_or_file, &statbuf);
  if (stat_result == -1)
  {
    fprintf(stderr, "Unable to access %s.\n", dir_or_file);
  }
  else if (is_file(&statbuf))
  {
    struct dirfile *f = (struct dirfile *)malloc(sizeof(struct dirfile));
    int access = make_dirfile(dir_or_file, dir_or_file, f);
    print_file(f, win_cols);
    free(f);
  }
  else if (is_dir(&statbuf))
  {
    print_dir(dir_or_file);
  }
  else
  {
    fprintf(stderr, "File type of %s is not supported by %s.\n", dir_or_file, PROGRAM);
  }
}

int make_dirfile(char *filename, char *full_path, struct dirfile *df)
{
  int stat_result = stat(full_path, &df->s);
  if (stat_result == -1)
  {
    fprintf(stderr, "Unable to access %s.\n", filename);
    return -1;
  }
  strcpy(df->filename, filename);
  return 1;
}

// See https://man7.org/linux/man-pages/man7/inode.7.html
int is_file(struct stat *f)
{
  return S_ISREG(f->st_mode);
}
int is_dir(struct stat *f)
{
  return S_ISDIR(f->st_mode);
}

void print_file(struct dirfile *d, int colsize)
{
  // Assume regular mode
  // fprintf(stdout, "%" COLWIDTH "s", name);
  fprintf(stdout, "%s", d->filename);
  for (size_t i = strlen(d->filename); i < colsize; i++)
  {
    fprintf(stdout, " ");
  }
  // fprintf(stdout, "Mode %u\n", st->st_mode);
  // fprintf(stdout, "User ID %u\n", st->st_uid);
  // fprintf(stdout, "Group ID %u\n", st->st_gid);
  // fprintf(stdout, "File Size %lld\n", st->st_size);
}

void print_dir(char *dir)
{
  DIR *dirfd;
  struct dirent *dp;
  if ((dirfd = opendir(dir)) == NULL)
  {
    fprintf(stderr, "Cannot access dir %s\n", dir);
    return;
  }
  int idx = 0, longest = 0;
  while ((dp = readdir(dirfd)) != NULL)
  {
    char *fmt = dir[-1] == '/' ? "%s%s" : "%s/%s";
    char *full_path = (char *)malloc(sizeof(dir) + sizeof(fmt) + sizeof(dp->d_name));
    sprintf(full_path, fmt, dir, dp->d_name);
    // Insert into sorted array
    struct dirfile *f = (struct dirfile *)malloc(sizeof(struct dirfile));
    make_dirfile(dp->d_name, full_path, f);
    insert_to_entries(f, idx, compare_lexagraphic);
    free(full_path);
    int f_len = strlen(dp->d_name);
    if (longest < f_len)
    {
      longest = f_len;
    }
    idx++;
  }
  int col_size = longest + COLGUTTER;
  int n_cols = win_cols / col_size;
  int n_rows = idx / n_cols;
  for (size_t row = 0; row < n_rows; row++)
  {
    for (size_t col = 0; col < n_cols; col++)
    {
      print_file(entries[(n_rows * col) + row], col_size);
    }
    fprintf(stdout, "\n");
  }
  closedir(dirfd);
}

void insert_to_entries(struct dirfile *f, int entries_len, int (*strategy)(struct dirfile *df1, struct dirfile *df2))
{
  // basically just insertion sort here, since we can expect a pretty small dataset based on dirsize max.
  // null case
  if (entries_len == 0)
  {
    entries[entries_len] = f;
    return;
  }
  int idx = 0;
  // Search until the end of the list OR we find a place where the previous value is less than current is equal to or greater than.
  while (idx < entries_len && !(strategy(f, entries[idx - 1]) < 0 && strategy(f, entries[idx]) >= 0))
  {
    idx++;
  }
  if (idx < entries_len)
  {
    shift_entries(idx, entries_len);
  }
  entries[idx] = f;
}

int compare_lexagraphic(struct dirfile *df1, struct dirfile *df2)
{
  int res = strcmp(df2->filename, df1->filename);
  return res;
}

void shift_entries(int idx, int curr_len)
{
  for (size_t i = curr_len; i >= idx; i--)
  {
    entries[i] = entries[i - 1];
  }
}

/*
  Next steps: 

  - determine column size of the terminal window AND determine max length of a filename.
  - sort in lexagraphical order
  - add support for varioous flags, some goals
  -- add a -h flag even though ls doesn't come with one, to substitute the man entry
  -- 1 / C
  -- a 
  -- c / S   
  -- i
*/