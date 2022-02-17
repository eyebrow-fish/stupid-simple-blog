#include "test.h"
#include "string.h"
#include <stdio.h>
#include "httprequest.h"

void can_deserialize_httprequest(Test *test)
{
	char *req = "GET / HTTP/1.1\r\nContent-Type:application/json\r\n\r\n{\"foo\":\"bar\"}";
	HttpRequest x;

	httprequest_deserialize(72, req, &x);

	assert(test, x.method == HTTP_GET, "Method was not GET");
	assert(test, strcmp(x.uri, "/") == 0, "URI was not /");
	assert(test, x.version == VERSION_1_1, "Version was not 1.1");
	// assert(test, strcmp(x.header_str, "Content-Type:application/json"), "Header string was wrong");
	// assert(test, strcmp(x.body_str, "{\"foo\":\"bar\"}"), "Body string was wrong");
}
