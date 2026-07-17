package answer

import (
	"ai-agent/internal/nlp"
	"ai-agent/internal/search"
	"strings"
)

type Plan struct {
	Intent       nlp.Intent
	Subject      string
	Facts        []string
	Technologies []string
}

func BuildPlan(result search.Result) Plan {
	return BuildPlanFromResults([]search.Result{result})
}

func BuildPlanFromResults(results []search.Result) Plan {
	if len(results) == 0 {
		return Plan{}
	}

	subject := "João Paulo"

	if results[0].HasEntity {
		subject = results[0].Entity.Value
	}

	facts := make([]string, 0, len(results))
	technologies := make([]string, 0)

	knowFacts := make(map[string]struct{})
	knowTechnologies := make(map[string]struct{})

	for _, result := range results {
		fact := strings.TrimSpace(result.Document.Content)

		if fact != "" {
			normalizedFact := strings.ToLower(fact)

			if _, exists := knowFacts[normalizedFact]; !exists {
				knowFacts[normalizedFact] = struct{}{}
				facts = append(facts, fact)
			}
		}

		for _, technology := range ExtractTechnologies(result.Document.Content) {
			normalizedTechnology := strings.ToLower(technology)

			if _, exists := knowTechnologies[normalizedTechnology]; exists {
				continue
			}

			knowTechnologies[normalizedTechnology] = struct{}{}
			technologies = append(technologies, technology)
		}
	}

	return Plan{
		Intent:       results[0].Intent,
		Subject:      subject,
		Facts:        facts,
		Technologies: technologies,
	}
}
