#include "ls-clone.h"

// Formatting globals
int win_cols;

// Flag globals
enum format_opts
{
  columns,
  lines,
  lines_long
};
// Force output to be one entry per line.  This is the default when output is not to a terminal.
enum format_opts format = lines;
enum sort_opts
{
  lexagraphic,
  none,
  size
};
// Force output to be one entry per line.  This is the default when output is not to a terminal.
enum sort_opts sort = lexagraphic;
int flag_all = 0;

int main(int argc, char *argv[])
{
  // Grab window size
  struct winsize w;
  ioctl(STDOUT_FILENO, TIOCGWINSZ, &w);
  win_cols = w.ws_col == 0 ? COL_DEFAULT : w.ws_col;
  if (isatty(fileno(stdout))) // Force multi-column output; this is the default when output is to a terminal.
  {
    format = columns;
  };
  if (argc > 1)
  {
    while (--argc > 0)
    {
      char *cur_arg = *++argv;
      if (!parse_flags(cur_arg))
      {
        print_dir_or_file(cur_arg);
      };
    }
  }
  else
  {
    print_dir_or_file(".");
  }
}

// Return 1 if valid flag, 0 if not flag.
int parse_flags(char *arg)
{
  if (arg[0] == '-')
  {
    for (size_t i = 1; i < strlen(arg); i++)
    {
      switch (arg[i])
      {
      case FLAG_COLUMNS:
        format = columns;
        break;
      case FLAG_LINES:
        format = lines;
        break;
      case FLAG_LONG:
        format = lines_long;
        break;
      case FLAG_ALL:
        flag_all = 1;
        break;
      case FLAG_NO_SORT:
        // Output is not sorted.  This option turns on the -a option.
        flag_all = 1;
        sort = none;
        break;
      case FLAG_SORT_SIZE:
        sort = size;
        break;
      default:
        fprintf(stderr, "Unrecognized flag %c.\n", arg[i]);
        break;
      }
    }
    return 1;
  }
  return 0;
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
  fprintf(stdout, "%s", d->filename);
  for (size_t i = strlen(d->filename); i < colsize; i++)
  {
    fprintf(stdout, " ");
  }
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
    // Skip if hidden unless all is set.
    if (dp->d_name[0] != '.' || flag_all)
    {
      char *fmt = dir[-1] == '/' ? "%s%s" : "%s/%s";
      char *full_path = (char *)malloc(sizeof(dir) + sizeof(fmt) + sizeof(dp->d_name));
      sprintf(full_path, fmt, dir, dp->d_name);
      // Insert into sorted array
      struct dirfile *f = (struct dirfile *)malloc(sizeof(struct dirfile));
      make_dirfile(dp->d_name, full_path, f);
      sort_strategy sorter = sort == lexagraphic ? &sort_lexagraphic : sort == size ? &sort_size
                                                                                    : &sort_no_op;
      insert_to_entries(f, idx, sorter);
      free(full_path);
      int f_len = strlen(dp->d_name);
      if (longest < f_len)
      {
        longest = f_len;
      }
      idx++;
    }
  }
  int col_size = longest + COL_GUTTER;
  int n_cols = win_cols / col_size;
  int n_rows = ceil(idx / (double)n_cols);
  if (format == columns)
  {
    for (size_t row = 0; row < n_rows; row++)
    {
      for (size_t col = 0; col < n_cols; col++)
      {
        int entry_idx = (n_rows * col) + row;
        if (entry_idx >= idx)
        {
          break;
        }
        print_file(entries[entry_idx], col_size);
      }
      fprintf(stdout, "\n");
    }
  }
  else if (format == lines)
  {
    for (size_t i = 0; i < idx; i++)
    {
      print_file(entries[i], col_size);
      fprintf(stdout, "\n");
    }
  }
  else if (format == lines_long)
  {
    fprintf(stdout, "TODO add support for long format\n.");
  }
  closedir(dirfd);
}

void insert_to_entries(struct dirfile *f, int entries_len, sort_strategy strategy)
{
  // basically just insertion sort here, since we can expect a pretty small dataset based on dirsize max.
  // null case
  if (entries_len == 0)
  {
    entries[entries_len] = f;
    return;
  }
  int idx = 0;
  // Search until we find a place where the previous value is less than and current is equal to or greater than OR the end of the list.
  while (idx < entries_len)
  {
    int check_curr = (*strategy)(f, entries[idx]);
    if (idx == 0 && check_curr >= 0)
    {
      break;
    }
    else if (idx > 0)
    {
      int check_back = (*strategy)(f, entries[idx - 1]);
      if (check_back < 0 && check_curr >= 0)
      {
        break;
      }
    }
    idx++;
  }
  // Need to shift if inserting into middle of array.
  if (idx < entries_len)
  {
    for (size_t i = entries_len; i > idx; i--)
    {
      entries[i] = entries[i - 1];
    }
  }
  entries[idx] = f;
}
