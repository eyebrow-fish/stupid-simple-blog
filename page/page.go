package page

import (
	_ "embed"
	"html/template"
)

var (
	//go:embed assets/main-style.html
	mainStyle string
	//go:embed assets/header.html
	header string
)

type Page[T any] struct {
	MainStyle template.HTML
	Header    template.HTML
	Data      *T
}

func WrapWithPage[T any](t *T) *Page[T] {
	return &Page[T]{template.HTML(mainStyle), template.HTML(header), t}
}

func WrapWithPageAndError[T any](t *T, err error) (*Page[T], error) {
	if err != nil {
		return nil, err
	}

	return &Page[T]{template.HTML(mainStyle), template.HTML(header), t}, nil
}
