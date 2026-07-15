package answer

import (
	"ai-agent/internal/nlp"
	"ai-agent/internal/search"
)

type Plan struct {
	Intent       nlp.Intent
	Subject      string
	Facts        []string
	Technologies []string
}

func BuildPlan(result search.Result) Plan {
	subject := "João Paulo"

	if result.HasEntity {
		subject = result.Entity.Value
	}

	return Plan{
		Intent:       result.Intent,
		Subject:      subject,
		Facts:        []string{result.Document.Content},
		Technologies: ExtractTechnologies(result.Document.Content),
	}
}
