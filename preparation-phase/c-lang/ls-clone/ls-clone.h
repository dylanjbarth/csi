#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <dirent.h>
#include <sys/stat.h>
#include <sys/ioctl.h>
#include <unistd.h>
#include <math.h>

#define PROGRAM "ls-clone"
#define COLGUTTER 5
#define MAX_DIRENT 65535  // https://stackoverflow.com/a/466596 ?
#define MAX_FILENAME 1024 // meh https://www.systutorials.com/maximum-allowed-file-path-length-for-c-programming-on-linux/

// The -1, -C, -x, and -l options all override each other; the last one specified determines the format used.
#define FLAG_COLUMNS 'C'   // Force multi-column output; this is the default when output is to a terminal.
#define FLAG_LINES '1'     //  (The numeric digit ``one''.)  Force output to be one entry per line.  This is the default when output is not to a terminal.
#define FLAG_LONG 'l'      // (The lowercase letter ``ell''.)  List in long format.  (See below.)  A total sum for all the file sizes is output on a line before the long listing.
#define FLAG_ALL 'a'       // Include directory entries whose names begin with a dot (.)
#define FLAG_NO_SORT 'f'   // Output is not sorted.  This option turns on the -a option.
#define FLAG_SORT_SIZE 'S' // Sort files by size

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
int no_op(struct dirfile *df1, struct dirfile *df2);
int compare_lexagraphic(struct dirfile *df1, struct dirfile *df2);
int compare_size(struct dirfile *df1, struct dirfile *df2);
int parse_flags(char *arg);
