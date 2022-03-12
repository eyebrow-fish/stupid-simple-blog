package comment

import (
	"database/sql"
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

func ReplyHandler(w http.ResponseWriter, r *http.Request) {
	pages.FormHandler[commentForm](w, r, commentForm{}, func(c commentForm) error {
		if c.Text == "" {
			return errors.New("comment cannot be empty")
		}

		id := mux.Vars(r)["id"]
		r, err := db.DB.Query(`
			insert into comments(user_id, post_id, text) 
			values(1, $1, $2)
		`, id, c.Text)
		if err != nil {
			return err
		}
		defer func(r *sql.Rows) { _ = r.Close() }(r)

		return nil
	})
}
