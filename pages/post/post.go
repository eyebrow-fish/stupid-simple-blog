package post

import (
	"database/sql"
	_ "embed"
	"fmt"
	"github.com/eyebrow-fish/stupid-simple-blog/pages"
	"html/template"
	"sort"
)

var One = pages.NewPage(onePostTemplate, oneHandler)
var All = pages.NewPage(allPostTemplate, allHandler)

//go:embed post.html
var onePostStr string

//go:embed all.html
var allPostStr string

var onePostTemplate = template.Must(template.New("post/one").Parse(onePostStr))
var allPostTemplate = template.Must(template.New("post/all").Parse(allPostStr))

type comment struct {
	Text *string
}

type post struct {
	Id           uint64
	Title        string
	Text         string
	CommentCount int
	Comments     []comment
}

func buildPost(r *sql.Rows) (*post, error) {
	ps, err := buildPosts(r)
	if err != nil {
		return nil, err
	}

	if len(ps) > 1 {
		return nil, fmt.Errorf("expected one post, got %d", len(ps))
	} else if len(ps) < 1 {
		return nil, fmt.Errorf("could not find post")
	}

	return &ps[0], nil
}

func buildPosts(r *sql.Rows) ([]post, error) {
	pm := make(map[uint64]post)
	for r.Next() {
		var p post
		var c comment
		err := r.Scan(&p.Id, &p.Title, &p.Text, &c.Text)
		if err != nil {
			return nil, err
		}

		if c.Text != nil {
			p.Comments = append(pm[p.Id].Comments, c)
			p.CommentCount = len(p.Comments)
		}

		pm[p.Id] = p
	}

	if len(pm) == 0 {
		return nil, nil
	}

	var ps []post
	for _, v := range pm {
		ps = append(ps, v)
	}

	sort.SliceStable(ps, func(i, j int) bool { return ps[i].Id > ps[j].Id })
	return ps, nil
}
