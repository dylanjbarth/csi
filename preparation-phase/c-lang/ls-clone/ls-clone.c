#include <stdio.h>
#include <stdlib.h>
#include <sys/stat.h>
#include <dirent.h>

#define PROGRAM "ls-clone"
#define MAXFILENAME 100

int is_file(struct stat *f);
int is_dir(struct stat *f);
void print_dir_or_file(char *s);
void print_file(char *file);
void print_dir(char *dir);

int main(int argc, char *argv[])
{
  while (--argc > 0)
  {
    char *cur_arg = *++argv;
    printf("Arg %s\n", cur_arg);
    print_dir_or_file(cur_arg);
  }
}

void print_dir_or_file(char *dir_or_file)
{
  struct stat statbuf;
  int stat_result = stat(dir_or_file, &statbuf);
  if (stat_result == -1)
  {
    printf("Unable to access %s.\n", dir_or_file);
  }
  else if (is_file(&statbuf))
  {
    print_file(dir_or_file);
  }
  else if (is_dir(&statbuf))
  {
    print_dir(dir_or_file);
  }
  else
  {
    printf("File type of %s is not supported by %s.\n", dir_or_file, PROGRAM);
  }
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

void print_file(char *file)
{
  fprintf(stdout, "File %s\n", file);
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
  fprintf(stdout, "Able to access dir %s\n", dir);
  while ((dp = readdir(dirfd)) != NULL)
  {
    print_file(dp->d_name);
  }
  closedir(dirfd);
  // Open dir, read each file, print it to stdout, close dir.
}
