#include <stdio.h>
#include <sys/stat.h>

#define PROGRAM "ls-clone"
#define MAXFILENAME 100

int is_file(struct stat *f);
int is_dir(struct stat *f);
void print_dir_or_file(char *s);

int main(int argc, char *argv[])
{
  printf("Hello, world! Welcome to %s\n", PROGRAM);
  printf("N args %d\n", argc);
  while (--argc > 0)
  {
    // If a file, just print the file
    // TODO some arg parsing first.
    char *cur_arg = *++argv;
    printf("Arg %s\n", cur_arg);
    print_dir_or_file(cur_arg);
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

void print_dir_or_file(char *dir_or_file)
{
  struct stat statbuf;
  int stat_result = stat(dir_or_file, &statbuf);
  if (stat_result == -1)
  {
    printf("Unable to open %s.\n", dir_or_file);
  }
  else if (is_file(&statbuf))
  {
    printf("%s is a file.\n", dir_or_file);
  }
  else if (is_dir(&statbuf))
  {
    printf("%s is a directory.\n", dir_or_file);
  }
  else
  {
    printf("File type of %s is not supported by %s.\n", dir_or_file, PROGRAM);
  }
}