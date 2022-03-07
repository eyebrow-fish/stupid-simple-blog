package post

import (
	"database/sql"
	"fmt"
	"github.com/eyebrow-fish/stupid-simple-blog/db"
	"strconv"
)

func oneHandler(v map[string]string) (*post, error) {
	id := v["id"]
	if _, err := strconv.Atoi(id); err != nil {
		return nil, fmt.Errorf("id cannot be \"%s\"", id)
	}

	r, err := db.DB.Query(`
		select p.id, p.title, p.text, c.text from posts p
		left join comments c
			on p.id = c.post_id
		where p.id = ` + id + `;`,
	)
	if err != nil {
		return nil, err
	}
	defer func(r *sql.Rows) { _ = r.Close() }(r)

	return buildPost(r)
}
