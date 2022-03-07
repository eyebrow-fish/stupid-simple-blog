package main

import (
	"github.com/eyebrow-fish/stupid-simple-blog/pages"
	"github.com/eyebrow-fish/stupid-simple-blog/pages/post"
	"github.com/gorilla/mux"
	"net/http"
)

var handlers = map[string]http.HandlerFunc{
	"/":     pages.PageHandler(post.All),
	"/{id}": pages.PageHandler(post.One),
}

func main() {
	r := mux.NewRouter()

	for k, v := range handlers {
		r.HandleFunc(k, v)
	}

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
