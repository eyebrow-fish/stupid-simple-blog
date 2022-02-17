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

int httprequest_method(int, const char *, HttpMethod *);

int httprequest_uri_range(int, const char *, int *, int *);

int httprequest_version(int, const char *, int, HttpVersion *);

int httprequest_header_str_range(int, const char *, int *, int *);

#endif
