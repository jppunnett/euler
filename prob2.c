//	By considering the terms in the Fibonacci sequence whose values do not
//	exceed four million, find the sum of the even-valued terms.

#include <stdio.h>

int main(int argc, char const *argv[])
{
	const int MAX_FIB = 4000000;

	int term1 = 0;
	int term2 = 1;
	int sum = 0;
	int sumeven = 0;

	while ((sum = term1 + term2) < MAX_FIB) {
		if (sum % 2 == 0)
			sumeven += sum;
		term1 = term2;
		term2 = sum;
	}

	printf("%d\n", sumeven);

	return 0;
}