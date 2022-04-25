package post

import (
	"database/sql"
	"fmt"
	"github.com/eyebrow-fish/stupid-simple-blog/db"
	"github.com/eyebrow-fish/stupid-simple-blog/page"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func oneHandler(r *http.Request) (*page.Page[post], error) {
	id := mux.Vars(r)["id"]
	if _, err := strconv.Atoi(id); err != nil {
		return nil, fmt.Errorf("id cannot be \"%s\"", id)
	}

	rows, err := db.DB.Query(`
		select p.id, p.title, p.text, pu.id, pu.email, c.id, c.post_id, c.text, cu.id, cu.email from posts p
		join users pu
		    on p.user_id = pu.id
		left join comments c
			on p.id = c.post_id
		left join users cu
			on c.user_id = cu.id
		where p.id = $1 order by p.id desc;
	`, id)
	if err != nil {
		return nil, err
	}
	defer func(r *sql.Rows) { _ = r.Close() }(rows)

	return page.WrapWithPageAndError[post](buildPost(rows))
}
