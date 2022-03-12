package post

import (
	_ "embed"
	"html/template"
)

//go:embed assets/one.html
var onePostStr string

//go:embed assets/all.html
var allPostStr string

//go:embed assets/create-form.html
var createPostStr string

var oneTemplate = template.Must(template.New("post/one").Parse(onePostStr))
var allTemplate = template.Must(template.New("post/all").Parse(allPostStr))
var createFormTemplate = template.Must(template.New("post/create").Parse(createPostStr))
