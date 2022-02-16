#include "test.h"

void assert_true(Test *test)
{
  assert_eq(test, (void *)"bar", (void*)"bar"); 
}