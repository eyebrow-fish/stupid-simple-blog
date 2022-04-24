package main

import (
	"github.com/eyebrow-fish/stupid-simple-blog/pages/comment"
	"github.com/eyebrow-fish/stupid-simple-blog/pages/post"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.Methods(http.MethodGet).Path("/create").Handler(post.CreateForm)
	r.Methods(http.MethodPost).Path("/create").Handler(post.Create)

	r.Methods(http.MethodGet).Path("/").Handler(post.All)
	r.Methods(http.MethodGet).Path("/{id}").Handler(post.One)

	r.Methods(http.MethodPost).Path("/reply").Handler(comment.Reply)

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
