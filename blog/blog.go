package blog

import (
	"database/sql"
	_ "embed"
	"fmt"
	"github.com/eyebrow-fish/stupid-simple-blog/db"
	_ "github.com/eyebrow-fish/stupid-simple-blog/db"
	"html/template"
	"strconv"
)

//go:embed page.html
var pageTemplate string

var Template = template.Must(template.New("blog").Parse(pageTemplate))

type Comment struct {
	Text string
}

type Post struct {
	Title    string
	Text     string
	Comments []Comment
}

func Handler(v map[string]string) (*Post, error) {
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

func buildPost(r *sql.Rows) (*Post, error) {
	if !r.Next() {
		return nil, nil
	}

	var b Post
	for r.Next() {
		var c Comment
		err := r.Scan(&b.Title, &b.Text, &c.Text)
		if err != nil {
			return nil, err
		}
		b.Comments = append(b.Comments, c)
	}

	return &b, nil
}
