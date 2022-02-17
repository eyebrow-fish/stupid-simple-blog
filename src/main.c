#include <stdio.h>
#include <stdlib.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <unistd.h>

#include "main.h"
#include "httprequest.h"
#include "httpresponse.h"

int sock_fd;

void log_error(char *msg)
{
	printf("[ERROR] %s\n", msg); // I was typing the same thing everywhere.
}

void log_info(char *msg)
{
	printf("[INFO] %s\n", msg); // I was typing the same thing everywhere.
}

void handle_conn(int conn_fd)
{
	char req_b[65536] = {};
	int rec_n = recv(conn_fd, req_b, 65536, 0);
	if (rec_n < 0)
	{
		log_error("Could not read from socket.");
		return;
	}

	HttpRequest req;
	int err = httprequest_deserialize(rec_n, req_b, &req);
	if (err != 0)
	{
		log_error("Could not deserialize request.");
		return;
	}

	log_info(req.uri);

	char res_b[65536] = {};
	HttpResponse res = {
			.version = VERSION_1_1,
			.status = STATUS_OK,
	};

	int send_n = httpresponse_serialize(&res, res_b);
	if (send_n < 0)
	{
		log_error("Could not serialize response.");
		return;
	}

	err = send(conn_fd, res_b, send_n, 0);
	if (err < 0)
	{
		log_error("Could not write to socket.");
		return;
	}

	err = close(conn_fd);
	if (err != 0)
	{
		log_error("Could not close connection.");
		return;
	}
}


void cleanup(void)
{
	if (sock_fd < 0)
		return;

	close(sock_fd);
}

int main(void)
{
	atexit(cleanup);

	sock_fd = socket(AF_INET, SOCK_STREAM, 0);
	if (sock_fd < 0)
	{
		perror("Could not create socket.");
		exit(-1);
	}

	struct sockaddr_in server;
	server.sin_family = AF_INET;
	server.sin_addr.s_addr = htonl(INADDR_ANY);
	server.sin_port = htons(PORT);

	int err = bind(sock_fd, (struct sockaddr *) &server, sizeof(server));
	if (err != 0)
	{
		perror("Could not bind socket.");
		exit(-1);
	}

	err = listen(sock_fd, 8);
	if (err != 0)
	{
		perror("Could not listen on socket.");
		exit(-1);
	}

	while (1)
	{
		struct sockaddr cli;
		socklen_t cli_len = sizeof(cli);
		int conn_fd = accept(sock_fd, (struct sockaddr *) &cli, &cli_len);
		if (conn_fd < 0)
		{
			log_error("Could not accept connection.");
			break;
		}

		handle_conn(conn_fd);
	}

	return 0;
}
