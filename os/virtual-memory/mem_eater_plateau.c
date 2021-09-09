#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

int main()
{
	unsigned int i = 0;
	while (1)
	{
		char *addr = malloc(1024 * 1024);
		addr = "nom";
		if (i % 10000 == 0)
		{
			printf("Iteration %d\n", i);
			sleep(1);
		}
		i++;
	}
}