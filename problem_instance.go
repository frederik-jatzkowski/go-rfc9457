package rfc9457

import "fmt"

type ProblemInstance struct {
	problemType ProblemType
	detail      string
}

func (i ProblemInstance) String() string {
	return fmt.Sprintf(
		`
{
	 "type": "/%s",
}
`,
		i.problemType.name,
	)
}
