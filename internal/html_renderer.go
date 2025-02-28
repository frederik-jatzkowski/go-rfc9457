package internal

import (
	"bytes"
	_ "embed"
	"html/template"
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

func (renderer HtmlRenderer) Render(args RenderArgs) ([]byte, error) {
	var buf bytes.Buffer

	err := renderer.tmpl.Execute(&buf, args)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
