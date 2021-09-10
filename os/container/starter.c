#define _GNU_SOURCE
#include <sched.h>
#include <sys/wait.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

#define STACK_SIZE 65536
#define MAX_PROCS 5

void update_cgroup();

struct child_config
{
  int argc;
  char **argv;
};

/* Entry point for child after `clone` */
int child(void *arg)
{
  update_cgroup();
  struct child_config *config = arg;
  if (execvpe(config->argv[0], config->argv, NULL))
  {
    fprintf(stderr, "execvpe failed %m.\n");
    return -1;
  }
  return 0;
}

void update_cgroup()
{
  // Write pid to cgroup file
  char *cgroup_path = "/sys/fs/cgroup/pids/contained/cgroup.procs";
  FILE *f = fopen(cgroup_path, "w");
  if (f == NULL)
  {
    printf("Error opening file %s. Please ensure the directory exists.\n", cgroup_path);
    exit(1);
  }
  fprintf(f, "%d", getpid());
  fclose(f);
  // Control number of allowed processes for group
  f = fopen("/sys/fs/cgroup/pids/contained/pids.max", "w");
  fprintf(f, "%d", MAX_PROCS);
  fclose(f);
}

int main(int argc, char **argv)
{
  struct child_config config = {0};
  // All the namespace flags https://man7.org/linux/man-pages/man7/namespaces.7.html
  int flags = SIGCHLD         // signal to parent on exit
              | CLONE_NEWNET  // try `ifconfig` in the container - returns nothing.
                              // | CLONE_NEWIPC // I can still kill another process -- is this just related to message queues?
              | CLONE_NEWNS   // isolated namespace for mount points.
              | CLONE_NEWPID  // isolated process space. I can still see new processes starting soutside of container but can't send signals to the pids I don't start.
              | CLONE_NEWUSER // isolated userspace  (try `users`/`groups` in and out of container) -- notice that bash starts as "nobody" instead of root
              // also sudo: effective uid is not 0, is sudo installed setuid root???
              | CLONE_NEWUTS // isolated system identifiers  (try setting `hostname` in and out of container)
      ;
  pid_t child_pid = 0;

  // Prepare child configuration
  config.argc = argc - 1;
  config.argv = &argv[1];

  // Allocate stack for child
  char *stack = 0;
  if (!(stack = malloc(STACK_SIZE)))
  {
    fprintf(stderr, "Malloc failed");
    exit(1);
  }

  // Clone parent, enter child code
  // NB flag requires invoking as sudo.
  if ((child_pid = clone(child, stack + STACK_SIZE, flags, &config)) == -1)
  {
    fprintf(stderr, "Clone failed");
    exit(2);
  }

  wait(NULL);

  return 0;
}
