package internal

import (
	_ "embed"
	"html/template"
	"io"
)

//go:embed problem-detail.html
var htmlTemplate string

type HtmlRenderer struct {
	tmpl *template.Template
}

func NewHtmlRenderer() (HtmlRenderer, error) {
	tmpl, err := template.New("").Parse(htmlTemplate)

	return HtmlRenderer{tmpl: tmpl}, err
}

type RenderArgs struct {
}

func (renderer HtmlRenderer) RenderTo(w io.Writer) error {
	dot := RenderArgs{}

	return renderer.tmpl.Execute(w, dot)
}
