package rfc9457

import "html/template"

var aboutBlank = ProblemType{
	name:                   "about:blank",
	title:                  "See HTTP Status Code",
	recommendedStatusCodes: nil,
	description: `
<p>
	The "about:blank" problem type has no additional semantics beyond that of the HTTP status code.
</p>
<p>
	<a href="https://www.rfc-editor.org/rfc/rfc9457.html#name-aboutblank" target="_blank">Further reading.</a>
</p>
`,
}

type ProblemType struct {
	name                   string
	title                  string
	recommendedStatusCodes []int
	description            template.HTML
}

func (t ProblemType) Instantiate(detail string) ProblemInstance {
	return ProblemInstance{
		problemType: t,
		detail:      detail,
	}
}
