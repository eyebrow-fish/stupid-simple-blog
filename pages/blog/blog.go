package blog

import (
	"github.com/eyebrow-fish/stupid-simple-blog/pages"
)

var One = pages.NewPage(postTemplate, getPostHandler)
