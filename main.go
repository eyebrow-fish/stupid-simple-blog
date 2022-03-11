package main

import (
	"github.com/eyebrow-fish/stupid-simple-blog/pages"
	"github.com/eyebrow-fish/stupid-simple-blog/pages/comment"
	"github.com/eyebrow-fish/stupid-simple-blog/pages/post"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.Methods(http.MethodGet).Path("/").Handler(pages.PageHandler(post.All))
	r.Methods(http.MethodGet).Path("/{id}").Handler(pages.PageHandler(post.One))
	r.Methods(http.MethodPost).Path("/{id}/reply").HandlerFunc(comment.ReplyHandler)

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
