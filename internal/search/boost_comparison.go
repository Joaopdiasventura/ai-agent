package search

import (
	"ai-agent/internal/domain"
	"ai-agent/internal/nlp"
	"strings"
)

func calculateProjectComparisonBoost(entity nlp.Entity, document domain.Document) float64 {
	content := normalizeForBoost(document.Content)
	entityValue := normalizeForBoost(entity.Value)

	score := 0.0

	if strings.Contains(content, entityValue) {
		score += 0.08
	}

	if document.Category == "comparison" {
		score += 0.18
	}

	if strings.Contains(document.ID, "comparison") {
		score += 0.14
	}

	if containsAnyPhrase(content, []string{
		"principal exemplo",
		"melhor projeto",
		"comparação",
		"comparacao",
		"comparado",
		"diferença",
		"diferenca",
		"superando",
		"alcançou",
	}) {
		score += 0.12
	}

	if strings.Contains(document.ID, "-description") {
		score -= 0.04
	}

	if strings.Contains(document.ID, "-technologies") {
		score -= 0.04
	}

	return score
}
