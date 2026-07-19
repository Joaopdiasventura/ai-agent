package search

import (
	"ai-agent/internal/domain"
	"ai-agent/internal/nlp"
	"strings"
)

func CalculateIntentBoost(analysis *nlp.QueryAnalysis, document *domain.Document) float64 {
	score := calculateCategoryBoost(analysis, document)
	score += calculateTemporalBoost(analysis, document)
	score += calculateLiteralDocumentBoost(analysis, document)

	if boost := calculateVisitorIntentBoost(analysis.PrimaryIntent, document); boost != 0 {
		return score + boost
	}

	if boost := calculateGeneralTechnologyBoost(analysis.PrimaryIntent, document); boost != 0 {
		return score + boost
	}

	if !analysis.HasEntity {
		return score
	}

	switch analysis.AnswerMode {
	case nlp.AnswerModeAbout:
		return score + calculateProjectAboutBoost(analysis.Entity, document)

	case nlp.AnswerModeTechnology:
		return score + calculateProjectTechnologyBoost(analysis.Entity, document)

	case nlp.AnswerModeComparison:
		return score + calculateProjectComparisonBoost(analysis.Entity, document)

	default:
		return score + calculateProjectDefaultBoost(analysis.Entity, document)
	}

}

func calculateLiteralDocumentBoost(analysis *nlp.QueryAnalysis, document *domain.Document) float64 {
	question := normalizeForBoost(analysis.Question)
	content := normalizeForBoost(document.Content)

	if question == "" || content == "" {
		return 0
	}

	if strings.Contains(question, content) {
		return 100.0
	}

	return 0
}

func calculateCategoryBoost(analysis *nlp.QueryAnalysis, document *domain.Document) float64 {
	switch analysis.CategoryHint {
	case nlp.CategoryHintUnknown:
		return 0
	case nlp.CategoryHintEducation:
		if document.Category == "education" {
			return 0.35
		}
	case nlp.CategoryHintContact:
		if document.Category == "contact" {
			return 0.35
		}
	case nlp.CategoryHintCareer:
		if document.Category == "career" {
			return 0.25
		}
		if document.Category == "impact" && strings.HasPrefix(document.ID, "career-") {
			return 0.15
		}
	case nlp.CategoryHintTechnology:
		if document.Category == "technology" {
			return 0.30
		}
	case nlp.CategoryHintProject:
		if document.Category == "project" {
			return 0.22
		}
		if document.Category == "comparison" {
			return 0.12
		}
	case nlp.CategoryHintImpact:
		if document.Category == "impact" {
			return 0.25
		}
	case nlp.CategoryHintProfile:
		if document.Category == "profile" {
			return 0.25
		}
		if document.Category == "identity" {
			return 0.12
		}
	case nlp.CategoryHintCertificate:
		if document.Category == "certificate" {
			return 0.30
		}
	}

	return -0.08
}

func calculateTemporalBoost(analysis *nlp.QueryAnalysis, document *domain.Document) float64 {
	switch analysis.TemporalContext {
	case nlp.TemporalPresent:
		return presentDocumentBoost(document)
	case nlp.TemporalPast:
		return pastDocumentBoost(document)
	case nlp.TemporalFirst:
		return firstDocumentBoost(document)
	default:
		return 0
	}
}

func presentDocumentBoost(document *domain.Document) float64 {
	switch {
	case documentIDMatches(document, "education-fiap"):
		return 0.45
	case documentIDMatches(document, "education-etec"):
		return -0.30
	case strings.HasPrefix(document.ID, "career-current"):
		return 0.30
	case strings.HasPrefix(document.ID, "career-junior") ||
		strings.HasPrefix(document.ID, "career-intern"):
		return -0.22
	default:
		return 0
	}
}

func pastDocumentBoost(document *domain.Document) float64 {
	switch {
	case documentIDMatches(document, "education-etec"):
		return 0.45
	case documentIDMatches(document, "education-fiap"):
		return -0.30
	case strings.HasPrefix(document.ID, "career-junior") ||
		strings.HasPrefix(document.ID, "career-intern"):
		return 0.20
	case strings.HasPrefix(document.ID, "career-current"):
		return -0.90
	default:
		return 0
	}
}

func firstDocumentBoost(document *domain.Document) float64 {
	if strings.HasPrefix(document.ID, "career-intern") {
		return 0.60
	}

	if strings.HasPrefix(document.ID, "career-current") ||
		strings.HasPrefix(document.ID, "career-junior") {
		return -0.45
	}

	return 0
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
