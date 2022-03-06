package blog

import (
	_ "embed"
	"github.com/eyebrow-fish/stupid-simple-blog/db"
	_ "github.com/eyebrow-fish/stupid-simple-blog/db"
	"html/template"
)

//go:embed page.html
var pageTemplate string

var Template = template.Must(template.New("blog").Parse(pageTemplate))

type Blog struct {
	Title string
}

func Render(m map[string]string) (*Blog, error) {
	r, err := db.DB.Query("select * from posts")
	if err != nil {
		return nil, err
	}

	if !r.Next() {
		return nil, nil
	}

	var b Blog
	if err = r.Scan(&b.Title); err != nil {
		return nil, err
	}

	return &b, nil
}
