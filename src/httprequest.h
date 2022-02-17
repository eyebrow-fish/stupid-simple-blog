#ifndef HTTP_REQUEST_H
#define HTTP_REQUEST_H

#include "http.h"

typedef enum
{
	HTTP_GET,
	HTTP_PUT,
	HTTP_POST,
	HTTP_DELETE,
	HTTP_OPTIONS,
} HttpMethod;

typedef struct
{
	HttpMethod method;
	char *uri;
	HttpVersion version;
	char *header_str;
	char *body_str;
} HttpRequest;

int httprequest_deserialize(int, char *, HttpRequest *);

#endif
