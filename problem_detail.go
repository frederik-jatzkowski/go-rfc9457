package rfc9457

import (
	"encoding/json"
)

type (
	ProblemInstanceDetail string
)

type ProblemDetail struct {
	ProblemType
	Detail ProblemInstanceDetail `json:"detail"`
}

func (i ProblemDetail) String() string {
	data, err := json.Marshal(i)
	if err != nil {
		return err.Error()
	}

	return string(data)
}
