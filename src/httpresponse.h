#ifndef HTTPRESPONSE_H
#define HTTPRESPONSE_H

#include "http.h"

typedef enum
{
	STATUS_OK,
	STATUS_BAD_REQUEST,
	STATUS_UNAUTHORIZED,
	STATUS_NOT_FOUND,
	STATUS_INTERNAL_SERVER_ERROR,
	STATUS_I_AM_A_TEAPOT,
} HttpStatus;

typedef struct
{
	HttpVersion version;
	HttpStatus status;
	char *header_str;
	char *body_str;
} HttpResponse;

int httpresponse_serialize(HttpResponse *, char *);

#endif
