package search

import (
	"ai-agent/internal/domain"
	"ai-agent/internal/nlp"
	"strings"
)

func calculateProjectDefaultBoost(entity nlp.Entity, document domain.Document) float64 {
	content := normalizeForBoost(document.Content)
	entityValue := normalizeForBoost(entity.Value)

	score := 0.0

	if document.Category == "project" {
		score += 0.04
	}

	if strings.Contains(content, entityValue) {
		score += 0.08
	}

	if strings.Contains(document.ID, "-description") {
		score += 0.06
	}

	if strings.Contains(document.ID, "-technologies") {
		score += 0.03
	}

	if strings.Contains(document.ID, "-metrics") ||
		strings.Contains(document.ID, "-performance") ||
		strings.Contains(document.ID, "-architecture") ||
		strings.Contains(document.ID, "-integrity") {
		score += 0.02
	}

	return score
}
