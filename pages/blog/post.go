package blog

import (
	"database/sql"
	_ "embed"
	"fmt"
	"github.com/eyebrow-fish/stupid-simple-blog/db"
	"html/template"
	"strconv"
)

//go:embed post.html
var postTemplateStr string

var postTemplate = template.Must(template.New("post").Parse(postTemplateStr))

type comment struct {
	Text string
}

type post struct {
	Title    string
	Text     string
	Comments []comment
}

func getPostHandler(v map[string]string) (*post, error) {
	id := v["id"]
	if _, err := strconv.Atoi(id); err != nil {
		return nil, fmt.Errorf("id cannot be \"%s\"", id)
	}

	r, err := db.DB.Query(`
		select p.title, p.text, c.text from posts p
		join comments c
			on p.id = c.post_id
		where p.id = ` + id + `;`,
	)
	if err != nil {
		return nil, err
	}
	defer func(r *sql.Rows) { _ = r.Close() }(r)

	return buildPost(r)
}

func buildPost(r *sql.Rows) (*post, error) {
	if !r.Next() {
		return nil, nil
	}

	var b post
	for r.Next() {
		var c comment
		err := r.Scan(&b.Title, &b.Text, &c.Text)
		if err != nil {
			return nil, err
		}
		b.Comments = append(b.Comments, c)
	}

	return &b, nil
}
