#ifndef TEST_H
#define TEST_H

typedef struct
{
	int num_failures;
	char failure_str[65536];
	int current;
} Test;

typedef void (*test_f)(Test *);

extern void assert(Test *, int, char *);

// Our tests.
void can_deserialize_httprequest(Test *);

void can_deserialize_httprequest_simple_rest_post(Test *);

void can_serialize_httpresponse(Test *);

#endif
