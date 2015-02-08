// Calculate the sum of all natural numbers < 1000 that are multiples of 3 & 5.
#include <stdio.h>

int main(int argc, char const *argv[])
{
	int sum = 0;
	
	int i;
	for (i = 3; i < 1000; ++i)
		if (i % 3 == 0 || i % 5 == 0)
			sum += i;

	printf("Sum of natural numbers below 1,000 is <%d>.\n", sum);

	return 0;
}