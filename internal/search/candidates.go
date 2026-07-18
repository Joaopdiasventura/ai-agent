package search

import (
	"ai-agent/internal/domain"
	"ai-agent/internal/nlp"
)

func FilterDocumentsByIntent(documents []*domain.Document, intent nlp.Intent) []*domain.Document {
	candidates := make([]*domain.Document, 0)

	for _, document := range documents {
		if matchesIntent(document, intent) {
			candidates = append(candidates, document)
		}
	}

	if len(candidates) == 0 {
		return documents
	}

	return candidates
}

func FilterDocumentsByEntity(results []Result, entity nlp.Entity) []Result {
	candidates := make([]Result, 0)

	for _, result := range results {
		if result.Entity.Value == entity.Value {
			candidates = append(candidates, result)
		}
	}

	return candidates
}

func matchesIntent(document *domain.Document, intent nlp.Intent) bool {
	switch intent {
	case nlp.IntentCurrentJob:
		return document.ID == "career-current-job"

	case nlp.IntentFirstJob:
		return document.ID == "career-first-job"

	case nlp.IntentEducation:
		return document.Category == "education"

	case nlp.IntentProject:
		return document.Category == "project" ||
			document.Category == "comparison"

	case nlp.IntentTechnologies:
		return document.Category == "project" ||
			document.Category == "career"

	case nlp.IntentContact:
		return document.Category == "contact"

	case nlp.IntentAbout:
		return document.Category == "identity"

	default:
		return true
	}
}
