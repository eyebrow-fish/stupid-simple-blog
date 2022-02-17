#include <string.h>
#include <stdarg.h>
#include <stdio.h>
#include "logger.h"

void log_prefixed(char *prefix, char *msg, va_list args)
{
	char s[strlen(msg) + strlen(prefix) + 2];
	strcpy(s, prefix);
	strcat(s, " ");
	strcat(s, msg);
	strcat(s, "\n");

	vprintf(s, args);
}

void log_info(char *msg, ...)
{
	va_list args;
	va_start(args, msg);
	log_prefixed("\e[32m[INFO]\e[0m", msg, args);
	va_end(args);
}

void log_warn(char *msg, ...)
{
	va_list args;
	va_start(args, msg);
	log_prefixed("\e[33m[WARN]\e[0m", msg, args);
	va_end(args);
}

void log_error(char *msg, ...)
{
	va_list args;
	va_start(args, msg);
	log_prefixed("\e[31m[ERROR]\e[0m", msg, args);
	va_end(args);
}
