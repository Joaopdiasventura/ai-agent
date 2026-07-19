package answer

import (
	"ai-agent/internal/nlp"
	"ai-agent/internal/search"
	"strings"
)

type Plan struct {
	Intent           nlp.Intent
	Language         nlp.Language
	Subject          string
	Facts            []string
	Technologies     []string
	DetailLevel      DetailLevel
	FormattedFacts   string
	SelectedProject  string
	ProjectCriterion nlp.ProjectCriterion
}

func BuildPlan(searchResult *search.SearchResult) Plan {
	if len(searchResult.Results) == 0 {
		return Plan{}
	}

	subject := "João Paulo"

	if searchResult.Results[0].HasEntity &&
		searchResult.Results[0].Entity.Type != nlp.EntityPerson {
		subject = searchResult.Results[0].Entity.Value
	}

	facts := make([]string, 0, len(searchResult.Results))
	technologies := make([]string, 0)

	knowFacts := make(map[string]struct{})
	knowTechnologies := make(map[string]struct{})

	for _, result := range searchResult.Results {
		fact := strings.TrimSpace(result.Document.Content)

		if fact != "" {
			normalizedFact := strings.ToLower(fact)

			if _, exists := knowFacts[normalizedFact]; !exists {
				knowFacts[normalizedFact] = struct{}{}
				facts = append(facts, fact)
			}
		}

		for _, technology := range nlp.ExtractTechnologies(result.Document.Content) {
			normalizedTechnology := strings.ToLower(technology)

			if _, exists := knowTechnologies[normalizedTechnology]; exists {
				continue
			}

			knowTechnologies[normalizedTechnology] = struct{}{}
			technologies = append(technologies, technology)
		}
	}

	detailLevel := SelectDetailLevel(searchResult.Tokens)
	detailLevel = SelectIntentDetailLevel(searchResult.Intent, detailLevel)
	selectedFacts := SelectFactsByDetail(facts, detailLevel)
	formattedFacts := FormatFacts(selectedFacts, searchResult.Language)

	return Plan{
		Intent:           searchResult.Intent,
		Language:         searchResult.Language,
		Subject:          subject,
		Facts:            facts,
		Technologies:     technologies,
		FormattedFacts:   formattedFacts,
		DetailLevel:      detailLevel,
		SelectedProject:  searchResult.SelectedProject,
		ProjectCriterion: searchResult.ProjectCriterion,
	}
}
