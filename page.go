package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

type Page interface {
	Template() (*template.Template, error)
	Render(map[string]string) any
}

func pageHandler(p Page) *mux.Router {
	r := mux.NewRouter()

	r.Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t, err := p.Template()
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = fmt.Fprint(w, t.Execute(w, p.Render(mux.Vars(r))))
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

	return r
}
