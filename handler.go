package rfc9457

import (
	"github.com/frederik-jatzkowski/go-rfc9457/internal"
	"net/http"
)

type Handler struct {
	registry Registry
	renderer internal.Renderer
}

func NewHandler(registry Registry) (Handler, error) {
	return Handler{
		registry: registry,
		renderer: internal.HtmlRenderer{},
	}, nil
}

func (h Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	path := request.URL.Path
	if len(path) > 0 && path[0] == '/' {
		path = path[1:]
	}

	problemType, exists := h.registry.types[path]
	if !exists {
		writer.WriteHeader(http.StatusNotFound)

		return
	}

	writer.Header().Set("Content-Type", "text/html")
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(problemType.Error()))
}
