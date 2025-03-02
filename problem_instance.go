package rfc9457

import (
	"encoding/json"
)

type (
	ProblemInstanceDetail string
)

type ProblemInstance struct {
	ProblemType
	Detail ProblemInstanceDetail `json:"detail"`
}

func (i ProblemInstance) String() string {
	data, err := json.Marshal(i)
	if err != nil {
		return err.Error()
	}

	return string(data)
}
