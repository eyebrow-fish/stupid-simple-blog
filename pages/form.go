package pages

import (
	"fmt"
	"github.com/gorilla/schema"
	"net/http"
)

var decoder = schema.NewDecoder()

func FormHandler[T any](w http.ResponseWriter, r *http.Request, t T, f func(t T) error) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, err.Error())
		return
	}

	err := decoder.Decode(&t, r.PostForm)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, err.Error())
		return
	}

	if err := f(t); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, err.Error())
		return
	}

	http.Redirect(w, r, r.Header.Get("referer"), http.StatusFound)
}
