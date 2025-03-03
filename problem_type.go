package rfc9457

import "html/template"

type (
	ProblemTypeName        string
	ProblemTypeUrl         string
	ProblemTypeTitle       string
	RecommendedStatusCodes []int
	ProblemTypeDescription template.HTML

	ProblemTypeOption func(*ProblemType)
)

type ProblemType struct {
	Name                   ProblemTypeName        `json:"-"`
	Type                   ProblemTypeUrl         `json:"type"`
	Title                  ProblemTypeTitle       `json:"title"`
	RecommendedStatusCodes RecommendedStatusCodes `json:"-"`
	Description            ProblemTypeDescription `json:"-"`
}

func (t ProblemType) Instantiate(detail string) ProblemDetail {
	return ProblemDetail{
		ProblemType: t,
		Detail:      ProblemInstanceDetail(detail),
	}
}

func WithTitle(title string) ProblemTypeOption {
	return func(t *ProblemType) {
		t.Title = ProblemTypeTitle(title)
	}
}

func WithDescription(description template.HTML) ProblemTypeOption {
	return func(t *ProblemType) {
		t.Description = ProblemTypeDescription(description)
	}
}
