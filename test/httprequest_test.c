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
	assert(test, strcmp(x.header_str, "Content-Type:application/json") == 0, "Header string was wrong");
	assert(test, strcmp(x.body_str, "{\"foo\":\"bar\"}") == 0, "Body string was wrong");
}

void can_deserialize_httprequest_simple_rest_post(Test *test)
{
	char *req = "POST /foo/bar/123 HTTP/2.0\r\nContent-Type:application/json\r\nAccess-Control-Allow-Origin:*\r\nAuthorization:qwerty\r\n\r\n{\"foo\":\"bar\",\n\"fizz\":123}";
	HttpRequest x;

	httprequest_deserialize(156, req, &x);

	assert(test, x.method == HTTP_POST, "Method was not POST");
	assert(test, strcmp(x.uri, "/foo/bar/123") == 0, "URI was not /foo/bar");
	assert(test, x.version == VERSION_2_0, "Version was not 2.0");
	assert(test, strcmp(x.header_str, "Content-Type:application/json\r\nAccess-Control-Allow-Origin:*\r\nAuthorization:qwerty") == 0, "Header string was wrong");
	assert(test, strcmp(x.body_str, "{\"foo\":\"bar\",\n\"fizz\":123}") == 0, "Body string was wrong");
}
