#include <string.h>
#include <stdio.h>
#include "httprequest.h"

int httprequest_deserialize(int srcc, char *srcv, HttpRequest *dest)
{
  // Resolve method.
  HttpMethod method;
  int err = get_method(srcc, srcv, &method);
  if (err != 0)
    return err;
  dest->method = method;

  // Resolve uri.
  int start, end;
  err = get_uri_range(srcc, srcv, &start, &end);
  if (err != 0)
    return err;
  char uri[1024];
  strncpy(uri, srcv + start, end - start);
  dest->uri = uri;

  return 0;
}

int get_method(int srcc, char *srcv, HttpMethod *method)
{
  if (srcc < 14) // No header available.
    return -1;

  switch (srcv[0])
  {
  case 'G':
    if (srcv[1] == 'E' && srcv[2] == 'T')
    {
      (*method) = HTTP_GET;
      return 0;
    }
    break;
  case 'P':
    if (srcv[1] == 'U' && srcv[2] == 'T')
    {
      (*method) = HTTP_PUT;
      return 0;
    }
    if (srcv[1] == 'O' && srcv[2] == 'S' && srcv[3] == 'T')
    {
      (*method) = HTTP_POST;
      return 0;
    }
    break;
  case 'D':
    if (srcv[1] == 'E' && srcv[2] == 'L' && srcv[3] == 'E' && srcv[5] == 'T' && srcv[6] == 'E')
    {
      (*method) = HTTP_DELETE;
      return 0;
    }
    break;
  case 'O':
    if (srcv[1] == 'P' && srcv[2] == 'T' && srcv[3] == 'I' && srcv[4] == 'O' && srcv[5] == 'T' && srcv[6] == 'S')
    {
      (*method) = HTTP_OPTIONS;
      return 0;
    }
    break;
  }

  return -1;
}

int get_uri_range(int srcc, char *srcv, int *start, int *end)
{
  int s = 0;
  int e = 0;
  for (int i = 0; i < srcc; i++)
  {
    char c = srcv[i];
    if (c == ' ')
    {
      if (s == 0 && srcc > i + 1)
      {
        s = i + 1;
      }
      else if (e == 0)
      {
        e = i;
        break;
      }
    }
  }

  if ((s == 0 || e == 0) && s > e) // Could not find start or end of URI.
    return -1;

  (*start) = s;
  (*end) = e;

  return 0;
}
