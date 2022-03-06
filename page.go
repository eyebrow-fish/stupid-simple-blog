package main

import (
	"fmt"
	"github.com/eyebrow-fish/stupid-simple-blog/blog"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

var handlers = map[string]http.HandlerFunc{
	"/{id}": pageHandler(page[blog.Post]{blog.Template, blog.Handler}),
}

type page[T any] struct {
	Template    *template.Template
	HandlerFunc handler[T]
}

type handler[T any] func(map[string]string) (*T, error)

func pageHandler[T any](p page[T]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		o, err := p.HandlerFunc(mux.Vars(r))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = fmt.Fprint(w, err.Error())
			return
		}

		if o == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		err = p.Template.Execute(w, o)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
