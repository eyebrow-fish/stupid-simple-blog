#include "test.h"
#include "httprequest.h"

void can_deserialize_httprequest(Test *test)
{
  char *req = "GET / HTTP/1.1\r\nContent-Type:application/json\r\n\r\n{\"foo\":\"bar\"}";
  HttpRequest x;

  httprequest_deserialize(req, &x);

  assert_eq(test, (void*)x.method, (void*)HTTP_GET);
  assert_eq(test, (void*)x.uri, (void*)"/");
  assert_eq(test, (void*)x.version, (void*)VERSION_1_1);
  assert_eq(test, (void*)x.header_str, (void*)"Content-Type:application/json");
  assert_eq(test, (void*)x.body_str, (void*)"{\"foo\":\"bar\"}");
}
