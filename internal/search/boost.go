package search

import (
	"ai-agent/internal/domain"
	"ai-agent/internal/nlp"
	"strings"
)

func CalculateIntentBoost(analysis *nlp.QueryAnalysis, document *domain.Document) float64 {
	if boost := calculateVisitorIntentBoost(analysis.PrimaryIntent, document); boost != 0 {
		return boost
	}

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
