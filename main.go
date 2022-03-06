package main

import (
	"github.com/eyebrow-fish/stupid-simple-blog/blog"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.PathPrefix("/blog").Handler(pageHandler(blog.Page{}))

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
