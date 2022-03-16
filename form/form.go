package form

import (
	"fmt"
	"github.com/gorilla/schema"
	"net/http"
)

var decoder = schema.NewDecoder()

type Form interface {
	Handle(r *http.Request) error
}

func NewFormHandler(form Form) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writeErr := func(code int, err error) {
			w.WriteHeader(code)
			_, _ = fmt.Fprintf(w, err.Error())
		}

		// Parse and decode the form against the request.
		if err := r.ParseForm(); err != nil {
			writeErr(http.StatusBadRequest, err)
			return
		}
		if err := decoder.Decode(form, r.PostForm); err != nil {
			writeErr(http.StatusBadRequest, err)
			return
		}

		// Form validation stage.
		if v, ok := form.(Validate); ok {
			err := v.Validate(w, r)
			if err != nil {
				writeErr(http.StatusBadRequest, err)
				return
			}
		}

		if err := form.Handle(r); err != nil {
			writeErr(http.StatusInternalServerError, err)
			return
		}

		// After stage.
		var err error
		if a, ok := form.(After); ok {
			err = a.After(w, r)
		} else {
			err = RefererRedirectAfter{}.After(w, r)
		}
		if err != nil {
			writeErr(http.StatusInternalServerError, err)
			return
		}
	}
}
