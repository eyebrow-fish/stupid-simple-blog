#include <stdio.h>
#include <string.h>
#include "httpresponse.h"

unsigned long httpresponse_serialize(HttpResponse *src, char *dest)
{
	// Version mapping.
	char *version;
	switch (src->version)
	{
		case VERSION_0_9:
			version = "0.9";
			break;
		case VERSION_1_0:
			version = "1.0";
			break;
		case VERSION_1_1:
			version = "1.1";
			break;
		case VERSION_2_0:
			version = "2.0";
			break;
		default:
			return -1;
	}

	// Status mapping.
	char *status_str;
	switch (src->status)
	{
		case STATUS_OK:
			status_str = "200 OK";
			break;
		case STATUS_BAD_REQUEST:
			status_str = "400 Bad Request";
			break;
		case STATUS_UNAUTHORIZED:
			status_str = "401 Unauthorized";
			break;
		case STATUS_NOT_FOUND:
			status_str = "404 Not Found";
			break;
		case STATUS_INTERNAL_SERVER_ERROR:
			status_str = "500 Internal Server Error";
			break;
		case STATUS_I_AM_A_TEAPOT:
			status_str = "418 I Am A Teapot";
			break;
		default:
			return -1;
	}

	char header_str[65536] = {};
	if (src->header_str != NULL)
	{
		strcpy(header_str, "\r\n");
		strcat(header_str, src->header_str);
	}

	char body_str[65536] = {};
	if (src->body_str != NULL)
		strcpy(body_str, src->body_str);

	sprintf(dest, "HTTP/%s %s%s\r\n\r\n%s", version, status_str, header_str, body_str);

	return strlen(dest);
}
