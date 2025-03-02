package rfc9457

import (
	"errors"
	"fmt"
	"net/url"
)

var (
	errAboutBlank           = errors.New("about:blank")
	ErrProblemTypeRedefined = errors.New("problem type redefined")
)

type (
	RegistryBasePath string
)

type Registry struct {
	basePath RegistryBasePath
	types    map[ProblemTypeName]ProblemType
}

func NewRegistry(basePath RegistryBasePath) (Registry, error) {
	registry := Registry{
		basePath: basePath,
		types:    make(map[ProblemTypeName]ProblemType),
	}

	err := registry.Define(
		errAboutBlank,
		WithDescription(`
<p>
	The "about:blank" problem type has no additional semantics beyond that of the HTTP status code.
</p>
<p>
	<a href="https://www.rfc-editor.org/rfc/rfc9457.html#name-aboutblank" target="_blank">Further reading.</a>
</p>
`),
	)
	if err != nil {
		return Registry{}, err
	}

	return registry, nil
}

func (r Registry) Define(err error, options ...ProblemTypeOption) error {
	name := ProblemTypeName(err.Error())

	_, found := r.findProblemByName(name)
	if found {
		return fmt.Errorf("defining problem type '%s': %w", name, ErrProblemTypeRedefined)
	}

	t := ProblemType{
		Type:  ProblemTypeUrl(string(r.basePath) + url.PathEscape(string(name))),
		Name:  name,
		Title: ProblemTypeTitle(name),
	}

	for _, option := range options {
		option(&t)
	}

	r.types[name] = t

	return nil
}

func (r Registry) ProblemTypeFor(err error) ProblemType {
	problemType, found := r.findProblemByName(ProblemTypeName(err.Error()))
	if !found {
		return r.types[ProblemTypeName(errAboutBlank.Error())]
	}

	return problemType
}

func (r Registry) findProblemByName(name ProblemTypeName) (ProblemType, bool) {
	problemType, found := r.types[name]
	return problemType, found
}
