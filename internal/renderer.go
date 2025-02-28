package internal

import (
	"html/template"
)

type RenderArgs struct {
	Name                   string
	Title                  string
	Description            template.HTML
	RecommendedStatusCodes string
}

type Renderer interface {
	Render(args RenderArgs) ([]byte, error)
}
