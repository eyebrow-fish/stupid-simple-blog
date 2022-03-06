package main

import (
	"github.com/eyebrow-fish/stupid-simple-blog/blog"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

var handlers = map[string]*mux.Router{
	"/blog": pageHandler(page[blog.Blog]{blog.Template, blog.Render}),
}

type page[T any] struct {
	Template *template.Template
	Renderer render[T]
}

type render[T any] func(map[string]string) T

func pageHandler[T any](p page[T]) *mux.Router {
	r := mux.NewRouter()

	r.Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := p.Template.Execute(w, p.Renderer(mux.Vars(r)))
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

	return r
}
