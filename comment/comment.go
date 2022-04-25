package comment

import (
	"errors"
	"github.com/eyebrow-fish/gosp"
	"github.com/eyebrow-fish/stupid-simple-blog/db"
	"github.com/eyebrow-fish/stupid-simple-blog/user"
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

var Reply = gosp.NewFormHandler[commentForm](
	func(r *http.Request, c *commentForm) error {
		if c.Text == "" {
			return errors.New("comment cannot be empty")
		}

		_, err := db.DB.Exec(`
			insert into comments(user_id, post_id, text)
			values(1, $1, $2)
		`, mux.Vars(r)["id"], c.Text)
		return err
	},
)
