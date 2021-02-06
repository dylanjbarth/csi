#include <stdio.h>
#include <stdlib.h>
#include <sys/stat.h>
#include <dirent.h>

#define PROGRAM "ls-clone"

int is_file(struct stat *f);
int is_dir(struct stat *f);
void print_dir_or_file(char *s);
void print_file(char *file, struct stat *s);
void print_dir(char *dir);

int main(int argc, char *argv[])
{
  while (--argc > 0)
  {
    char *cur_arg = *++argv;
    print_dir_or_file(cur_arg);
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
    print_file(dir_or_file, &statbuf);
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

// See https://man7.org/linux/man-pages/man7/inode.7.html
int is_file(struct stat *f)
{
  return S_ISREG(f->st_mode);
}
int is_dir(struct stat *f)
{
  return S_ISDIR(f->st_mode);
}

void print_file(char *name, struct stat *st)
{
  fprintf(stdout, "File %s\n", name);
  fprintf(stdout, "Mode %u\n", st->st_mode);
  fprintf(stdout, "User ID %u\n", st->st_uid);
  fprintf(stdout, "Group ID %u\n", st->st_gid);
  fprintf(stdout, "File Size %lld\n", st->st_size);
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
    struct stat statbuf;
    int stat_result = stat(dp->d_name, &statbuf);
    if (stat_result == -1)
    {
      fprintf(stderr, "Unable to access %s.\n", dp->d_name);
    }
    print_file(dp->d_name, &statbuf);
  }
  closedir(dirfd);
}
