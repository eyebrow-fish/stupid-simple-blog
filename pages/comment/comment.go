package comment

import (
	"errors"
	"github.com/eyebrow-fish/gosp"
	"github.com/eyebrow-fish/stupid-simple-blog/db"
	"github.com/eyebrow-fish/stupid-simple-blog/pages/user"
)

type Comment struct {
	Id     uint64
	PostId uint64
	User   user.User
	Text   string
}

type commentForm struct {
	Id   uint64
	Text string
}

var Reply = gosp.NewFormHandler[commentForm](
	func(c *commentForm) error {
		if c.Text == "" {
			return errors.New("comment cannot be empty")
		}

		_, err := db.DB.Exec(`
			insert into comments(user_id, post_id, text)
			values(1, $1, $2)
		`, c.Id, c.Text)
		if err != nil {
			return err
		}
		return nil
	},
)
