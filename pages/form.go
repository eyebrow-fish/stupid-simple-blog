package pages

import (
	"fmt"
	"github.com/gorilla/schema"
	"net/http"
)

var decoder = schema.NewDecoder()

type Form interface {
	Handle(r *http.Request) error
}

func FormHandler(f Form) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = fmt.Fprintf(w, err.Error())
			return
		}

		err := decoder.Decode(&f, r.PostForm)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = fmt.Fprintf(w, err.Error())
			return
		}

		if err := f.Handle(r); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = fmt.Fprintf(w, err.Error())
			return
		}

		http.Redirect(w, r, r.Header.Get("referer"), http.StatusFound)
	}
}
