#include <stdio.h>
#include <string.h>
#include "test.h"
#include "logger.h"

void assert(Test *test, int assertion, char *message)
{
	if (assertion)
		return;

	char failure[1024] = {};
	sprintf(failure, "[Test #%d] %s\n", test->current, message);
	strcat(test->failure_str, failure);
	test->num_failures = test->num_failures + 1;
}

int main(void)
{
	test_f all_tests[] = {
			can_deserialize_httprequest,
			can_deserialize_httprequest_simple_rest_post,
			can_serialize_httpresponse,
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
		log_info("All tests passed.");
	}
	else
	{
		log_error("\e[31mThere were %d failures:\n", test.num_failures);
		log_error("\e[31m%s\e[0m", test.failure_str);
	}
}
