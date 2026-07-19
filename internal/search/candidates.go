package search

import (
	"ai-agent/internal/domain"
	"ai-agent/internal/nlp"
	"strings"
)

func FilterDocumentsByIntent(documents []*domain.Document, analysis *nlp.QueryAnalysis) []*domain.Document {
	languageCandidates := filterDocumentsByLanguage(documents, analysis.Language)
	intentCandidates := make([]*domain.Document, 0)
	compatibleCandidates := make([]*domain.Document, 0)

	for _, document := range languageCandidates {
		if questionContainsDocumentContent(document, analysis) {
			intentCandidates = append(intentCandidates, document)
			compatibleCandidates = append(compatibleCandidates, document)
			continue
		}

		if matchesIntent(document, analysis.PrimaryIntent) {
			intentCandidates = append(intentCandidates, document)

			if documentIsCompatibleWithAnalysis(document, analysis) {
				compatibleCandidates = append(compatibleCandidates, document)
			}
		}
	}

	if len(compatibleCandidates) > 0 {
		return compatibleCandidates
	}

	if len(intentCandidates) == 0 {
		return languageCandidates
	}

	return intentCandidates
}

func filterDocumentsByLanguage(documents []*domain.Document, language nlp.Language) []*domain.Document {
	candidates := make([]*domain.Document, 0, len(documents))

	for _, document := range documents {
		if document.Language == string(language) {
			candidates = append(candidates, document)
		}
	}

	return candidates
}

func questionContainsDocumentContent(document *domain.Document, analysis *nlp.QueryAnalysis) bool {
	question := strings.ToLower(strings.TrimSpace(analysis.Question))
	content := strings.ToLower(strings.TrimSpace(document.Content))

	return question != "" && content != "" && strings.Contains(question, content)
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
		return document.Category == "career" || document.Category == "impact"

	case nlp.IntentFirstJob:
		return documentIDMatches(document, "career-intern-job") ||
			documentIDMatches(document, "career-intern-catalog") ||
			documentIDMatches(document, "career-intern-search") ||
			documentIDMatches(document, "career-intern-auth-performance")

	case nlp.IntentEducation:
		return document.Category == "education"

	case nlp.IntentProject:
		return document.Category == "project" ||
			document.Category == "impact" ||
			document.Category == "comparison"

	case nlp.IntentProjectRecommendation:
		return document.Category == "project" ||
			document.Category == "impact" ||
			document.Category == "technology" ||
			document.Category == "comparison"

	case nlp.IntentTechnologies:
		return document.Category == "technology" ||
			document.Category == "project" ||
			document.Category == "career" ||
			document.Category == "impact" ||
			document.Category == "certificate"

	case nlp.IntentContact:
		return document.Category == "contact"

	case nlp.IntentVisitorSummary:
		return document.Category == "identity" ||
			document.Category == "profile"

	case nlp.IntentVisitorProjects:
		return documentIDMatches(document, "project-comparison-best") ||
			document.Category == "project"

	case nlp.IntentVisitorServices:
		return document.Category == "service"

	case nlp.IntentHireReason:
		return documentIDMatches(document, "identity-professional-summary") ||
			documentIDMatches(document, "profile-focus") ||
			documentIDMatches(document, "profile-availability") ||
			documentIDMatches(document, "career-current-impact") ||
			document.Category == "impact"

	case nlp.IntentAbout:
		return document.Category == "identity"

	default:
		return true
	}
}

func documentIsCompatibleWithAnalysis(document *domain.Document, analysis *nlp.QueryAnalysis) bool {
	if analysis.PrimaryIntent == nlp.IntentProject &&
		analysis.HasEntity &&
		analysis.Entity.Type == nlp.EntityProject &&
		analysis.AnswerMode != nlp.AnswerModeComparison &&
		document.Category == "comparison" {
		return false
	}

	switch analysis.CategoryHint {
	case nlp.CategoryHintEducation:
		return educationDocumentMatchesTemporalContext(document, analysis.TemporalContext)
	case nlp.CategoryHintCareer:
		return careerDocumentMatchesTemporalContext(document, analysis.TemporalContext)
	default:
		return true
	}
}

func educationDocumentMatchesTemporalContext(document *domain.Document, temporalContext nlp.TemporalContext) bool {
	switch temporalContext {
	case nlp.TemporalPresent:
		return !documentIDMatches(document, "education-etec")
	case nlp.TemporalPast:
		return !documentIDMatches(document, "education-fiap")
	default:
		return true
	}
}

func careerDocumentMatchesTemporalContext(document *domain.Document, temporalContext nlp.TemporalContext) bool {
	switch temporalContext {
	case nlp.TemporalPresent:
		return !strings.HasPrefix(document.ID, "career-junior") &&
			!strings.HasPrefix(document.ID, "career-intern")
	case nlp.TemporalPast:
		return !strings.HasPrefix(document.ID, "career-current")
	case nlp.TemporalFirst:
		return strings.HasPrefix(document.ID, "career-intern")
	default:
		return true
	}
}
