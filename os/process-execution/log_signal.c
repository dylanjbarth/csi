#include <stdio.h>
#include <signal.h>
#include <unistd.h>

void signal_handler(int sig)
{
  printf("Got signal %d\n", sig);
  return;
}

int main()
{

  // Create signal handlers
  int max_sig = 31;
  for (int i = 1; i <= max_sig; i++)
  {
    if (signal(i, signal_handler) == SIG_ERR)
    {
      printf("Error creating signal handler for %d\n", i);
    };
  }
  printf("Hi there I'm process # %d. Sleeping for awhile while you send me signals.\n", getpid());
  sleep(60);
}