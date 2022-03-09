package post

import (
	"database/sql"
	"github.com/eyebrow-fish/stupid-simple-blog/db"
)

func allHandler(_ map[string]string) (*[]post, error) {
	r, err := db.DB.Query(`
		select p.id, p.title, p.text, pu.id, pu.email, c.id, c.post_id, c.text, cu.id, cu.email from posts p
		join users pu
		    on p.user_id = pu.id
		left join comments c
			on p.id = c.post_id
		left join users cu
			on c.user_id = cu.id
		order by p.id desc
	`)
	if err != nil {
		return nil, err
	}
	defer func(r *sql.Rows) { _ = r.Close() }(r)

	ps, err := buildPosts(r)
	return &ps, err
}
