#ifndef HTTP_H
#define HTTP_H

typedef enum
{
	HTTP_GET,
	HTTP_PUT,
	HTTP_POST,
	HTTP_DELETE,
	HTTP_OPTIONS,
} HttpMethod;

typedef enum
{
	VERSION_0_9,
	VERSION_1_0,
	VERSION_1_1,
	VERSION_2_0,
} HttpVersion;

typedef struct
{
	HttpMethod method;
	char *uri;
	HttpVersion version;
	char *header_str;
	char *body_str;
} HttpRequest;

int get_method(int, const char *, HttpMethod *);

int get_uri_range(int, const char *, int *, int *);

int get_version(int, const char *, int, HttpVersion *);

int get_header_str_range(int, const char *, int *, int *);

int httprequest_deserialize(int, char *, HttpRequest *);

#endif
