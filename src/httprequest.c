#include <string.h>
#include "httprequest.h"
#include "http.h"

int httprequest_method(int srcc, const char *srcv, HttpMethod *method)
{
	if (srcc < 14) // No header available.
		return -1;

	switch (srcv[0])
	{
		case 'G':
			if (srcv[1] == 'E' && srcv[2] == 'T') // GET
			{
				(*method) = HTTP_GET;
				return 0;
			}
			break;
		case 'P':
			if (srcv[1] == 'U' && srcv[2] == 'T') // PUT
			{
				(*method) = HTTP_PUT;
				return 0;
			}
			if (srcv[1] == 'O' && srcv[2] == 'S' && srcv[3] == 'T') // POST
			{
				(*method) = HTTP_POST;
				return 0;
			}
			break;
		case 'D':
			if (srcv[1] == 'E' && srcv[2] == 'L' && srcv[3] == 'E' && srcv[5] == 'T' && srcv[6] == 'E') // DELETE
			{
				(*method) = HTTP_DELETE;
				return 0;
			}
			break;
		case 'O':
			if (srcv[1] == 'P' && srcv[2] == 'T' && srcv[3] == 'I' && srcv[4] == 'O' && srcv[5] == 'T' &&
			    srcv[6] == 'S') // OPTIONS
			{
				(*method) = HTTP_OPTIONS;
				return 0;
			}
			break;
	}

	return -1;
}

int httprequest_uri_range(int srcc, const char *srcv, int *start, int *end)
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
			else
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

int httprequest_version(int srcc, const char *srcv, int uri_end, HttpVersion *version)
{
	int vs = uri_end + 6;
	if (srcc < vs || srcv[vs + 1] != '.')
		return -1;

	switch (srcv[vs])
	{
		case '0':
			if (srcv[vs + 2] == '9') // 0.9
			{
				(*version) = VERSION_0_9;
				return 0;
			}
			break;
		case '1':
			if (srcv[vs + 2] == '0') // 1.0
			{
				(*version) = VERSION_1_0;
				return 0;
			}
			if (srcv[vs + 2] == '1') // 1.1
			{
				(*version) = VERSION_1_1;
				return 0;
			}
			break;
		case '2':
			if (srcv[vs + 2] == '0') // 2.0
			{
				(*version) = VERSION_2_0;
				return 0;
			}
			break;
	}

	return 0;
}

int httprequest_header_str_range(int srcc, const char *srcv, int *start, int *end)
{
	int s = 0;
	int e = 0;
	for (int i = 0; i < srcc; i++)
	{
		if (srcv[i] == '\r' && srcv[i + 1] == '\n')
		{
			if (s == 0 && srcc > i + 2)
			{
				s = i + 2;
			}
			else if (srcv[i] == '\r' && srcv[i + 1] == '\n' && srcv[i + 2] == '\r' && srcv[i + 3] == '\n' )
			{
				e = i;
				break;
			}
		}
	}

	if ((s == 0 || e == 0) && s > e) // Could not find start or end of headers.
		return -1;

	(*start) = s;
	(*end) = e;

	return 0;
}

int httprequest_deserialize(int srcc, char *srcv, HttpRequest *dest)
{
	// Resolve method.
	HttpMethod method;
	int err = httprequest_method(srcc, srcv, &method);
	if (err != 0)
		return err;
	dest->method = method;

	// Resolve uri.
	int uri_s, uri_e;
	err = httprequest_uri_range(srcc, srcv, &uri_s, &uri_e);
	if (err != 0)
		return err;
	char uri[1024] = {};
	strncpy(uri, srcv + uri_s, uri_e - uri_s);
	dest->uri = uri;

	// Resolve version.
	HttpVersion version;
	err = httprequest_version(srcc, srcv, uri_e, &version);
	if (err != 0)
		return err;
	dest->version = version;

	// Resolve header_str.
	int header_s, header_e;
	err = httprequest_header_str_range(srcc, srcv, &header_s, &header_e);
	if (err != 0)
		return err;
	char header_str[65536] = {};
	strncpy(header_str, srcv + header_s, header_e - header_s);
	dest->header_str = header_str;

	// Resolve body_str.
	int body_s = header_e + 4; // Prepended with \r\n\r\n.
	char body_str[65536] = {};
	strncpy(body_str, srcv + body_s, srcc - body_s);
	dest->body_str = body_str;

	return 0;
}
