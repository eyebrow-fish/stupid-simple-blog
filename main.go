package main

import (
	"github.com/eyebrow-fish/stupid-simple-blog/comment"
	post2 "github.com/eyebrow-fish/stupid-simple-blog/post"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.Methods(http.MethodGet).Path("/create").Handler(post2.CreateForm)
	r.Methods(http.MethodPost).Path("/create").Handler(post2.Create)

	r.Methods(http.MethodGet).Path("/").Handler(post2.All)
	r.Methods(http.MethodGet).Path("/{id}").Handler(post2.One)

	r.Methods(http.MethodPost).Path("/{id}/reply").Handler(comment.Reply)

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
