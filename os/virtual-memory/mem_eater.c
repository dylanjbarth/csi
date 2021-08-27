#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

int main()
{
	unsigned int i = 0;
	while (1)
	{
		char *addr = malloc(sizeof "nomnomnom " * 1024); // should be 1 byte per char * len of 10 = 10 bytes... * 1024 = 10kb per iteration, every 1k iterations = 1MB, every 1M iterations = 1GB
		addr = "nom nom nom";
		if (i % 1000 == 0)
		{
			printf("Iteration %d\n", i);
			sleep(1);
		}
		i++;
	}
}
