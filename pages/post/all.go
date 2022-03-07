package post

import (
	"database/sql"
	"github.com/eyebrow-fish/stupid-simple-blog/db"
)

func allHandler(_ map[string]string) (*[]post, error) {
	r, err := db.DB.Query(`
		select p.id, p.title, p.text, c.text from posts p
		left join comments c
			on p.id = c.post_id
	`)
	if err != nil {
		return nil, err
	}
	defer func(r *sql.Rows) { _ = r.Close() }(r)

	ps, err := buildPosts(r)
	return &ps, err
}
