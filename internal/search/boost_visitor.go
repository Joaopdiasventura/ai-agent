package search

import (
	"ai-agent/internal/domain"
	"ai-agent/internal/nlp"
)

func calculateVisitorIntentBoost(intent nlp.Intent, document *domain.Document) float64 {
	switch intent {
	case nlp.IntentVisitorSummary:
		if document.Category == "identity" {
			return 0.20
		}

	case nlp.IntentVisitorProjects:
		if documentIDMatches(document, "project-comparison-best") {
			return 0.30
		}

		if document.Category == "project" {
			return 0.05
		}

	case nlp.IntentVisitorServices:
		if document.Category == "service" {
			return 0.25
		}

	case nlp.IntentHireReason:
		if documentIDMatches(document, "identity-professional-summary") ||
			documentIDMatches(document, "profile-focus") ||
			documentIDMatches(document, "profile-availability") ||
			documentIDMatches(document, "career-current-impact") {
			return 0.20
		}
	}

	return 0
}
