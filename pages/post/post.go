package post

import (
	"database/sql"
	_ "embed"
	"fmt"
	"github.com/eyebrow-fish/stupid-simple-blog/pages"
	"github.com/eyebrow-fish/stupid-simple-blog/pages/user"
	"html/template"
	"sort"
)

var One = pages.NewPage(onePostTemplate, oneHandler)
var All = pages.NewPage(allPostTemplate, allHandler)

//go:embed assets/one.html
var onePostStr string

//go:embed assets/all.html
var allPostStr string

var onePostTemplate = template.Must(template.New("post/one").Parse(onePostStr))
var allPostTemplate = template.Must(template.New("post/all").Parse(allPostStr))

type comment struct {
	Id     uint64
	PostId uint64
	User   user.User
	Text   string
}

type post struct {
	Id           uint64
	Title        string
	Text         string
	User         user.User
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
		var u user.User
		// Necessary to map it manually and to construct it later because we have to assume these could be nil with a
		// left join.
		var (
			cId  *uint64
			cPId *uint64
			cT   *string
			cUId *uint64
			cUE  *string
		)
		err := r.Scan(&p.Id, &p.Title, &p.Text, &u.Id, &u.Email, &cId, &cPId, &cT, &cUId, &cUE)
		if err != nil {
			return nil, err
		}

		p.User = u
		if cId != nil {
			p.Comments = append(
				pm[p.Id].Comments,
				comment{*cId, *cPId, user.User{Id: *cUId, Email: *cUE}, *cT},
			)
			p.CommentCount = len(p.Comments)
		}

		pm[p.Id] = p
	}

	// No posts. We can just 404.
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
