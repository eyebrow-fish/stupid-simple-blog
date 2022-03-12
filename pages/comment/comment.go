package comment

import (
	"errors"
	"github.com/eyebrow-fish/stupid-simple-blog/db"
	"github.com/eyebrow-fish/stupid-simple-blog/pages"
	"github.com/eyebrow-fish/stupid-simple-blog/pages/user"
	"github.com/gorilla/mux"
	"net/http"
)

type Comment struct {
	Id     uint64
	PostId uint64
	User   user.User
	Text   string
}

type commentForm struct {
	Text string
}

func (c commentForm) Handle(r *http.Request) error {
	if c.Text == "" {
		return errors.New("comment cannot be empty")
	}

	id := mux.Vars(r)["id"]
	rs, err := db.DB.Query(`
			insert into comments(user_id, post_id, text) 
			values(1, $1, $2)
		`, id, c.Text)
	if err != nil {
		return err
	}
	defer func() { _ = rs.Close() }()

	return nil
}

var Reply = pages.FormHandler(commentForm{})
