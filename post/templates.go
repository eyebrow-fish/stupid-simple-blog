package post

import (
	_ "embed"
	"html/template"
)

var (
	//go:embed assets/one.html
	onePostStr string
	//go:embed assets/all.html
	allPostStr string
	//go:embed assets/create-form.html
	createPostStr string
)

var (
	oneTemplate        = template.Must(template.New("post/one").Parse(onePostStr))
	allTemplate        = template.Must(template.New("post/all").Parse(allPostStr))
	createFormTemplate = template.Must(template.New("post/create").Parse(createPostStr))
)
