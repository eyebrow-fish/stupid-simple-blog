#include <stdio.h>
#include "test.h"

void assert(Test *test, int assertion, char *message)
{
	if (assertion)
		return;

	char failure[1024];
	sprintf(failure, "%s", message);
	test->failures[test->num_failures] = failure;
	test->num_failures = test->num_failures + 1;
}

int main(void)
{
	test_f all_tests[] = {
			can_deserialize_httprequest,
	};
	int num_all_tests = sizeof(all_tests) / sizeof(all_tests[0]);

	Test test = {};

	for (int i = 0; i < num_all_tests; i++)
	{
		test.current = i;
		all_tests[i](&test);
	}

	if (test.num_failures == 0)
	{
		printf("All tests passed.\n");
	} else
	{
		printf("There were %d failures:\n", test.num_failures);
	}

	for (int i = 0; i < test.num_failures; i++)
	{
		printf("%s\n", test.failures[i]);
	}
}
