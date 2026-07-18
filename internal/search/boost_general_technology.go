package search

import (
	"ai-agent/internal/domain"
	"ai-agent/internal/nlp"
	"strings"
)

func calculateGeneralTechnologyBoost(intent nlp.Intent, document *domain.Document) float64 {
	if intent != nlp.IntentTechnologies {
		return 0
	}

	score := 0.0

	if document.Category == "technology" {
		score += 0.25
	}

	if strings.Contains(document.ID, "-technologies") {
		score += 0.20
	}

	return score
}
