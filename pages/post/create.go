package post

import (
	"errors"
	"github.com/eyebrow-fish/gosp"
	"github.com/eyebrow-fish/stupid-simple-blog/db"
	"net/http"
)

var Create = gosp.NewFormHandler(
	func(p *postForm) error {
		if p.Title == "" {
			return errors.New("title cannot be empty")
		}
		if p.Text == "" {
			return errors.New("text cannot be empty")
		}

		_, err := db.DB.Exec(`
			insert into posts(title, text, user_id)
			values($1, $2, 1)
		`, p.Title, p.Text)
		if err != nil {
			return err
		}

		return nil
	},
	func(f *gosp.FormHandler[postForm]) {
		f.RedirectHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/", http.StatusFound)
		})
	},
)

type postForm struct {
	Title string
	Text  string
}
