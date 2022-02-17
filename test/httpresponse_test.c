#include <stdio.h>
#include <string.h>
#include "httpresponse.h"
#include "test.h"

void can_serialize_httpresponse(Test *test)
{
	char *exp = "HTTP/1.1 200 OK\r\nContent-Type:plain/text\r\n\r\nHello, World!";
	HttpResponse resp = {
			.version = VERSION_1_1,
			.status = STATUS_OK,
			.header_str = "Content-Type:plain/text",
			.body_str = "Hello, World!",
	};

	char x[248] = {};
	httpresponse_serialize(&resp, x);

	char msg[248] = {};
	sprintf(msg, "HTTP response is not \"%s\"", exp);
	assert(test, strcmp(x, exp) == 1, x);
}
