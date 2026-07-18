package search

import (
	"ai-agent/internal/domain"
	"ai-agent/internal/nlp"
)

func FilterDocumentsByIntent(documents []*domain.Document, analysis *nlp.QueryAnalysis) []*domain.Document {
	candidates := make([]*domain.Document, 0)

	for _, document := range documents {
		if document.Language != string(analysis.Language) {
			continue
		}

		if matchesIntent(document, analysis.PrimaryIntent) {
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
		return documentIDMatches(document, "career-current-job")

	case nlp.IntentFirstJob:
		return documentIDMatches(document, "career-first-job")

	case nlp.IntentEducation:
		return document.Category == "education"

	case nlp.IntentProject:
		return document.Category == "project" ||
			document.Category == "comparison"

	case nlp.IntentTechnologies:
		return document.Category == "technology" ||
			document.Category == "project" ||
			document.Category == "career"

	case nlp.IntentContact:
		return document.Category == "contact"

	case nlp.IntentVisitorSummary:
		return document.Category == "identity" ||
			documentIDMatches(document, "profile-focus")

	case nlp.IntentVisitorProjects:
		return documentIDMatches(document, "project-comparison-best") ||
			document.Category == "project"

	case nlp.IntentVisitorServices:
		return document.Category == "service"

	case nlp.IntentHireReason:
		return documentIDMatches(document, "identity-professional-summary") ||
			documentIDMatches(document, "profile-focus") ||
			documentIDMatches(document, "profile-availability") ||
			documentIDMatches(document, "career-current-impact")

	case nlp.IntentAbout:
		return document.Category == "identity"

	default:
		return true
	}
}
