package post

import (
	"errors"
	"github.com/eyebrow-fish/stupid-simple-blog/db"
	"github.com/eyebrow-fish/stupid-simple-blog/pages"
	"net/http"
)

type postForm struct {
	Title string
	Text  string
}

func (p postForm) Handle(_ *http.Request) error {
	if p.Title == "" {
		return errors.New("title cannot be empty")
	}
	if p.Text == "" {
		return errors.New("text cannot be empty")
	}

	rs, err := db.DB.Query(`
			insert into posts(title, text, user_id)
			values($1, $2, 1)
		`, p.Title, p.Text)
	if err != nil {
		return err
	}
	defer func() { _ = rs.Close() }()

	return nil
}

var Create = pages.FormHandler(postForm{})
