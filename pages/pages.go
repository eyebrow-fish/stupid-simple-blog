package pages

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

type handler[T any] func(map[string]string) (*T, error)

type Page[T any] struct {
	template    *template.Template
	handlerFunc handler[T]
}

func NewPage[T any](t *template.Template, h handler[T]) Page[T] { return Page[T]{t, h} }

func NewEmptyPage(t *template.Template) Page[struct{}] {
	return NewPage(t, func(_ map[string]string) (*struct{}, error) { return &struct{}{}, nil })
}

func PageHandler[T any](p Page[T]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		o, err := p.handlerFunc(mux.Vars(r))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = fmt.Fprint(w, err.Error())
			return
		}

		if o == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		err = p.template.Execute(w, newPageData(o))
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
