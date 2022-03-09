package pages

import (
	_ "embed"
	"html/template"
)

//go:embed assets/main-style.html
var mainStyle string

//go:embed assets/header.html
var header string

type pageData[T any] struct {
	MainStyle template.HTML
	Header    template.HTML
	Data      *T
}

func newPageData[T any](t *T) pageData[T] {
	return pageData[T]{template.HTML(mainStyle), template.HTML(header), t}
}
