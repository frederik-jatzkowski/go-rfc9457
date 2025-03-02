package rfc9457

import (
	"github.com/frederik-jatzkowski/go-rfc9457/internal"
	"html/template"
	"net/http"
)

type Handler struct {
	registry Registry
	renderer internal.Renderer
}

func NewHandler(registry Registry) (Handler, error) {
	renderer, err := internal.NewHtmlRenderer()
	if err != nil {
		return Handler{}, err
	}

	return Handler{
		registry: registry,
		renderer: renderer,
	}, nil
}

func (h Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	path := request.URL.Path
	if len(path) > 0 && path[0] == '/' {
		path = path[1:]
	}

	problemType, exists := h.registry.findProblemByName(ProblemTypeName(path))
	if !exists {
		writer.WriteHeader(http.StatusNotFound)

		return
	}

	data, err := h.renderer.Render(internal.RenderArgs{
		Name:        string(problemType.Name),
		Title:       string(problemType.Title),
		Description: template.HTML(problemType.Description),
	})
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
	}

	writer.Header().Set("Content-Type", "text/html")
	writer.WriteHeader(http.StatusOK)
	writer.Write(data)
}
