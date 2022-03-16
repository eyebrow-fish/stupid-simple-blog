package form

import "net/http"

type After interface {
	After(w http.ResponseWriter, r *http.Request) error
}

type RefererRedirectAfter struct{}

func (a RefererRedirectAfter) After(w http.ResponseWriter, r *http.Request) error {
	http.Redirect(w, r, r.Header.Get("referer"), http.StatusFound)
	return nil
}

type Validate interface {
	Validate(w http.ResponseWriter, r *http.Request) error
}
