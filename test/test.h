#ifndef TEST_H
#define TEST_H

typedef struct
{
  int num_failures;
  char *failures[1024];
  int current;
} Test;

typedef void (*test_f)(Test *);

extern void assert(Test *, int, char *);

// Our tests.
void can_deserialize_httprequest(Test *);

#endif
