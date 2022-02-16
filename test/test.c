#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "test.h"

void assert_eq(Test *test, void *expected, void *actual)
{
  if (expected == actual)
    return;

  char failure[1024];
  sprintf(failure, "Expected [%p] did not equal actual [%p].", expected, actual);
  test->failures[test->num_failures] = failure;
  test->num_failures++;
}

void assert_neq(Test *test, void *expected, void *actual)
{
  if (expected != actual)
    return;

  char failure[1024];
  sprintf(failure, "Expected [%p] was equal to actual [%p].", expected, actual);
  test->failures[test->num_failures] = failure;
  test->num_failures++;
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

  if (test.num_failures > 0)
  {
    printf("There were %d failures:\n", test.num_failures);
  }

  for (int i = 0; i < test.num_failures; i++)
  {
    printf("%s\n", test.failures[i]);
  }
}
