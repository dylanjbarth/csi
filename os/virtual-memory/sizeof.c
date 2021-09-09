#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

int main()
{
	printf("%lu\n", sizeof '0');
	printf("%lu\n", sizeof " ");
	printf("%lu\n", sizeof "  ");
	printf("%lu\n", sizeof "   ");
}
