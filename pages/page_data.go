package pages

import (
	_ "embed"
	"html/template"
)

//go:embed assets/main.css
var mainStyle string

//go:embed assets/header.html
var header string

type pageData[T any] struct {
	MainStyle template.CSS
	Header    template.HTML
	Data      *T
}

func newPageData[T any](t *T) pageData[T] {
	return pageData[T]{template.CSS(mainStyle), template.HTML(header), t}
}
