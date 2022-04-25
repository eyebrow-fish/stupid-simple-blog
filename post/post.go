package post

import (
	"database/sql"
	"fmt"
	"github.com/eyebrow-fish/gosp"
	"github.com/eyebrow-fish/stupid-simple-blog/comment"
	"github.com/eyebrow-fish/stupid-simple-blog/page"
	"github.com/eyebrow-fish/stupid-simple-blog/user"
	"net/http"
	"sort"
)

var One = gosp.NewPageHandler(oneHandler, oneTemplate)
var All = gosp.NewPageHandler(allHandler, allTemplate)
var CreateForm = gosp.NewPageHandler(
	func(request *http.Request) (*page.Page[struct{}], error) {
		return page.WrapWithPage(&struct{}{}), nil
	},
	createFormTemplate,
)

type post struct {
	Id           uint64
	Title        string
	Text         string
	User         user.User
	CommentCount int
	Comments     []comment.Comment
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
				comment.Comment{
					Id:     *cId,
					PostId: *cPId,
					User:   user.User{Id: *cUId, Email: *cUE},
					Text:   *cT,
				},
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
