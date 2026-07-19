package confidence

import (
	"ai-agent/internal/domain"
	"testing"
)

func TestPolicyReturnsLowWithoutEvidence(t *testing.T) {
	assessment := DefaultPolicy().Assess(domain.Query{}, nil, nil)

	if assessment.Level != Low {
		t.Fatalf("Assess() level = %q, want %q", assessment.Level, Low)
	}
}

func TestPolicyReturnsHighForCoherentStrongResult(t *testing.T) {
	query := domain.Query{Language: "pt", Category: "contact"}
	result := confidenceResult("contact-email-pt", "pt", "contact", "", 80, []string{"vector", "lexical"})
	evidence := []domain.Evidence{{DocumentID: "contact-email-pt", Language: "pt", Category: "contact"}}

	assessment := DefaultPolicy().Assess(query, []domain.SearchResult{result}, evidence)

	if assessment.Level != High {
		t.Fatalf("Assess() level = %q, want %q", assessment.Level, High)
	}
}

func TestPolicyDowngradesWeakOrPenalizedResult(t *testing.T) {
	query := domain.Query{Language: "pt", Category: "contact"}
	result := confidenceResult("contact-email-en", "en", "contact", "", 30, []string{"vector"})
	result.PenaltyReasons = []string{"language_mismatch"}
	evidence := []domain.Evidence{{DocumentID: "contact-email-en", Language: "en", Category: "contact"}}

	assessment := DefaultPolicy().Assess(query, []domain.SearchResult{result}, evidence)

	if assessment.Level != Low {
		t.Fatalf("Assess() level = %q, want %q", assessment.Level, Low)
	}
}

func confidenceResult(id string, language string, category string, project string, score float64, sources []string) domain.SearchResult {
	return domain.SearchResult{
		Document: &domain.Document{
			ID:             id,
			Language:       language,
			Category:       category,
			Subject:        id,
			Project:        project,
			TemporalStatus: domain.TemporalTimeless,
			Keywords:       []string{id},
			Content:        id,
		},
		FinalScore: score,
		Sources:    sources,
	}
}
