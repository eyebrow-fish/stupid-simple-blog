package post

import (
	"errors"
	"github.com/eyebrow-fish/stupid-simple-blog/db"
	"github.com/eyebrow-fish/stupid-simple-blog/form"
	"net/http"
)

var Create = form.NewFormHandler(&postForm{})

type postForm struct {
	Title string
	Text  string
}

func (p postForm) Handle(_ *http.Request) error {
	_, err := db.DB.Exec(`
			insert into posts(title, text, user_id)
			values($1, $2, 1)
		`, p.Title, p.Text)
	if err != nil {
		return err
	}

	return nil
}

func (p postForm) Validate(w http.ResponseWriter, r *http.Request) error {
	if p.Title == "" {
		return errors.New("title cannot be empty")
	}
	if p.Text == "" {
		return errors.New("text cannot be empty")
	}
	return nil
}

func (p postForm) After(w http.ResponseWriter, r *http.Request) error {
	http.Redirect(w, r, "/", http.StatusFound)
	return nil
}
