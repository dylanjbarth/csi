#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

int main()
{
	unsigned int i = 0;
	while (1)
	{
		char *addr = malloc(1024); // 1kib
		// explicitly fill every single byte with a char
		for (size_t i = 0; i < 1024; i++)
		{
			addr[i] = '0';
		}

		if (i % 1000 == 0)
		{
			printf("Iteration %d\n", i);
		}
		if (i % 100000 == 0)
		{
			sleep(1);
		}
		i++;
	}
}
