package search

import (
	"ai-agent/internal/domain"
	"ai-agent/internal/nlp"
	"strings"
)

func calculateTechnologyIntentBoost(analysis nlp.QueryAnalysis, document *domain.Document) float64 {
	if !analysis.HasEntity {
		return 0
	}

	entityValue := normalizeForBoost(analysis.Entity.Value)
	content := normalizeForBoost(document.Content)

	score := 0.0

	if document.Category == "technology" {
		score += 0.10
	}

	if document.Category == "project" {
		score += 0.05
	}

	if strings.Contains(document.ID, "-technologies") {
		score += 0.22
	}

	if containsAnyPhrase(content, []string{
		"tecnologias utilizadas",
		"tecnologias usadas",
		"utiliza",
		"utilizam",
		"usa",
		"stack",
	}) {
		score += 0.14
	}

	if strings.Contains(content, entityValue) {
		score += 0.08
	}

	if containsAnyPhrase(content, nlp.KnownTechnologies) {
		score += 0.08
	}

	if strings.Contains(document.ID, "-description") {
		score -= 0.04
	}

	return score
}

func calculateProjectTechnologyBoost(entity nlp.Entity, document *domain.Document) float64 {
	content := normalizeForBoost(document.Content)
	entityValue := normalizeForBoost(entity.Value)

	score := 0.0

	if document.Category != "project" {
		score -= 0.05
	}

	if strings.Contains(document.ID, "-technologies") {
		score += 0.24
	}

	if strings.Contains(content, entityValue) {
		score += 0.08
	}

	if containsAnyPhrase(content, []string{
		"tecnologias utilizadas",
		"tecnologias usadas",
		"stack",
		"utiliza",
		"usa",
	}) {
		score += 0.16
	}

	if strings.Contains(document.ID, "-description") {
		score -= 0.04
	}

	return score
}
