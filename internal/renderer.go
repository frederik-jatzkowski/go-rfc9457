package internal

import "io"

type Renderer interface {
	RenderTo(w io.Writer)
}
