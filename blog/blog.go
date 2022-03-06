package blog

import "html/template"

type Blog struct {
	Title string
}

type Page struct{}

func (p Page) Template() (*template.Template, error) {
	return template.ParseFiles("blog/page.html")
}

func (p Page) Render(m map[string]string) any {
	return Blog{"Foo"}
}
