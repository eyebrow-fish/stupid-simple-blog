# stupid-simple-blog

Simple server-side rendering with Go HTML template, and backed by a PSQL.
It doesn't get much more simple.

# dev

I develop against a local postgres container managed by `docker-compose`, everything you need is in
[build](./build).

The rest is all written in Go 1.18, so just compile the source with a compatible `go` binary.
