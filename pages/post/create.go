package post

import (
	"github.com/eyebrow-fish/stupid-simple-blog/pages"
	"net/http"
)

type postForm struct {
	Title string
	Text  string
}

func Create(w http.ResponseWriter, r *http.Request) {
	pages.FormHandler[postForm](w, r, postForm{}, func(p postForm) error {
		return nil
	})
}
