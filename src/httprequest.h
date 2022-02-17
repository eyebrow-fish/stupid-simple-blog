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

int get_method(int, char *, HttpMethod *);
int get_uri_range(int, char *, int *, int *);
int get_version(int, char *, int, HttpVersion *);
int httprequest_deserialize(int, char *, HttpRequest *);

#endif
