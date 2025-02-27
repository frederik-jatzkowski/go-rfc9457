package rfc9457

type Registry struct {
	types map[string]error
}

func NewRegistry() Registry {
	return Registry{
		types: make(map[string]error),
	}
}

func (r Registry) RegisterError(err error, recommendedStatusCode int) error {
	r.types[err.Error()] = err

	return nil
}
