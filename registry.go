package rfc9457

import "fmt"

type Registry struct {
	types map[string]ProblemType
}

func NewRegistry() Registry {
	return Registry{
		types: map[string]ProblemType{
			aboutBlank.name: aboutBlank,
		},
	}
}

func (r Registry) RegisterError(err error, recommendedStatusCode int) error {
	name := err.Error()

	_, found := r.findProblemByName(name)
	if found {
		return fmt.Errorf("redefinition of problem type '%s'", name)
	}

	r.types[err.Error()] = ProblemType{
		name: err.Error(),
	}

	return nil
}

func (r Registry) ProblemTypeFor(err error) ProblemType {
	problemType, found := r.findProblemByName(err.Error())
	if !found {
		return aboutBlank
	}

	return problemType
}

func (r Registry) findProblemByName(name string) (ProblemType, bool) {
	problemType, found := r.types[name]
	return problemType, found
}
