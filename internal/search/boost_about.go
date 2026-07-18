package search

import (
	"ai-agent/internal/domain"
	"ai-agent/internal/nlp"
	"strings"
)

func calculateProjectAboutBoost(entity nlp.Entity, document domain.Document) float64 {
	content := normalizeForBoost(document.Content)
	entityValue := normalizeForBoost(entity.Value)

	score := 0.0

	if document.Category != "project" {
		score -= 0.05
	}

	if strings.Contains(document.ID, "-description") {
		score += 0.16
	}

	if strings.HasPrefix(content, entityValue+" é ") {
		score += 0.20
	}

	if strings.Contains(content, entityValue+" é ") {
		score += 0.14
	}

	if strings.Contains(content, entityValue+" foi ") {
		score += 0.08
	}

	if containsAnyPhrase(content, []string{
		entityValue + " é um ",
		entityValue + " é uma ",
	}) {
		score += 0.12
	}

	if containsAnyPhrase(content, []string{
		"tecnologias utilizadas",
		"possui",
		"alcançou",
		"suporta",
	}) {
		score -= 0.04
	}

	if document.Category == "comparison" {
		score -= 0.08
	}

	return score
}
