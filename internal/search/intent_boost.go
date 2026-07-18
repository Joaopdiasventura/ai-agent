package search

import (
	"ai-agent/internal/domain"
	"ai-agent/internal/nlp"
	"strings"
)

func CalculateIntentBoost(analysis *nlp.QueryAnalysis, document domain.Document) float64 {
	if !analysis.HasEntity {
		return 0
	}

	switch analysis.AnswerMode {
	case nlp.AnswerModeAbout:
		return calculateProjectAboutBoost(analysis.Entity, document)

	case nlp.AnswerModeTechnology:
		return calculateProjectTechnologyBoost(analysis.Entity, document)

	case nlp.AnswerModeComparison:
		return calculateProjectComparisonBoost(analysis.Entity, document)

	default:
		return calculateProjectDefaultBoost(analysis.Entity, document)
	}

}

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

func calculateTechnologyIntentBoost(analysis nlp.QueryAnalysis, document domain.Document) float64 {
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

func calculateProjectTechnologyBoost(entity nlp.Entity, document domain.Document) float64 {
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

func normalizeForBoost(value string) string {
	return strings.ToLower(strings.TrimSpace(value))
}

func containsAnyPhrase(content string, phrases []string) bool {
	for _, phrase := range phrases {
		if strings.Contains(content, phrase) {
			return true
		}
	}

	return false
}
