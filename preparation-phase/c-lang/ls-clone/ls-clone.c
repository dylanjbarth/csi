#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <dirent.h>
#include <sys/stat.h>
#include <sys/ioctl.h>
#include <unistd.h>

#define PROGRAM "ls-clone"
#define COLWIDTH "20"
#define MAXDIRSIZE 65535  // https://stackoverflow.com/a/466596 ?
#define MAX_FILENAME 1024 // meh https://www.systutorials.com/maximum-allowed-file-path-length-for-c-programming-on-linux/

struct dirfile
{
  char filename[MAX_FILENAME];
  struct stat s;
};

int is_file(struct stat *f);
int is_dir(struct stat *f);
void print_dir_or_file(char *s);
void print_file(struct dirfile *d);
void print_dir(char *dir);
int make_dirfile(char *filename, struct dirfile *df);
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
    int access = make_dirfile(dir_or_file, f);
    print_file(f);
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

int make_dirfile(char *filename, struct dirfile *df)
{
  int stat_result = stat(filename, &df->s);
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

void print_file(struct dirfile *d)
{
  // Assume regular mode
  // fprintf(stdout, "%" COLWIDTH "s", name);
  fprintf(stdout, "File %s\n", d->filename);
  // fprintf(stdout, "Mode %u\n", st->st_mode);
  // fprintf(stdout, "User ID %u\n", st->st_uid);
  // fprintf(stdout, "Group ID %u\n", st->st_gid);
  // fprintf(stdout, "File Size %lld\n", st->st_size);
}

void print_dir(char *dir)
{
  int idx = 0;
  DIR *dirfd;
  struct dirent *dp;
  if ((dirfd = opendir(dir)) == NULL)
  {
    fprintf(stderr, "Cannot access dir %s\n", dir);
    return;
  }
  while ((dp = readdir(dirfd)) != NULL)
  {
    char *fmt = dir[-1] == '/' ? "%s%s" : "%s/%s";
    char *full_path = (char *)malloc(sizeof(dir) + sizeof(fmt) + sizeof(dp->d_name));
    sprintf(full_path, fmt, dir, dp->d_name);
    struct dirfile *d = (struct dirfile *)malloc(sizeof(struct dirfile));
    make_dirfile(full_path, d);
    print_file(d);
    free(d);
    idx++;
  }
  closedir(dirfd);
}

/*
  Next steps: 

  - determine column size of the terminal window AND determine max length of a filename.
  - sort in lexagraphical order
*/