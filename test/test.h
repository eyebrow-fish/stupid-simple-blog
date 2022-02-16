#ifndef TEST_H
#define TEST_H

typedef struct
{
  int num_failures;
  char *failures[1024];
  int current;
} Test;

typedef void (*test_f)(Test *);

extern void assert_eq(Test *, void *, void *);
extern void assert_neq(Test *, void *, void *);

// Our tests.
void can_deserialize_httprequest(Test *);

#endif
