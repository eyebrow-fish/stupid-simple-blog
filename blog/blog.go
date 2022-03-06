package blog

import (
	_ "embed"
	"html/template"
)

//go:embed page.html
var pageTemplate string

var Template = template.Must(template.New("blog").Parse(pageTemplate))

type Blog struct {
	Title string
}

func Render(m map[string]string) Blog {
	return Blog{"Foo"}
}
