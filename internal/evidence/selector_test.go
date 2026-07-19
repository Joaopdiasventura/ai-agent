package evidence

import (
	"ai-agent/internal/domain"
	"testing"
)

func TestSelectorUsesSingleEvidenceForFactualQuestion(t *testing.T) {
	selector := DefaultSelector()
	results := []domain.SearchResult{
		evidenceResult("contact-email-pt", "pt", "contact", "", "email", 10),
		evidenceResult("contact-phone-pt", "pt", "contact", "", "phone", 9),
	}

	selected := selector.Select(domain.Query{Text: "Qual é o email dele?", Language: "pt", Category: "contact"}, results)

	if len(selected) != 1 {
		t.Fatalf("Select() returned %d evidences, want 1", len(selected))
	}

	if selected[0].DocumentID != "contact-email-pt" {
		t.Fatalf("Select() evidence = %q, want contact-email-pt", selected[0].DocumentID)
	}
}

func TestSelectorGroupsCompatibleProjectEvidenceForSynthesis(t *testing.T) {
	selector := DefaultSelector()
	results := []domain.SearchResult{
		evidenceResult("project-xtube-leadership-pt", "pt", "impact", "x-tube", "leadership", 10),
		evidenceResult("project-xtube-processing-pt", "pt", "project", "x-tube", "processing", 9),
		evidenceResult("project-auronix-description-pt", "pt", "project", "auronix", "auronix", 8),
		evidenceResult("project-xtube-en", "en", "project", "x-tube", "english", 7),
	}

	selected := selector.Select(domain.Query{Text: "Qual projeto demonstra liderança e impacto?", Language: "pt", Category: "project"}, results)

	if len(selected) != 2 {
		t.Fatalf("Select() returned %d evidences, want 2", len(selected))
	}

	for _, evidence := range selected {
		if evidence.Project != "x-tube" || evidence.Language != "pt" {
			t.Fatalf("Select() returned incompatible evidence: %#v", evidence)
		}
	}
}

func TestSelectorRespectsExplicitProject(t *testing.T) {
	selector := DefaultSelector()
	results := []domain.SearchResult{
		evidenceResult("project-xtube-pt", "pt", "project", "x-tube", "xtube", 10),
		evidenceResult("project-auronix-pt", "pt", "project", "auronix", "auronix", 9),
	}

	selected := selector.Select(domain.Query{Text: "Me fale do Auronix", Language: "pt", Category: "project", Project: "auronix"}, results)

	if len(selected) != 1 || selected[0].DocumentID != "project-auronix-pt" {
		t.Fatalf("Select() returned %#v, want only project-auronix-pt", selected)
	}
}

func evidenceResult(id string, language string, category string, project string, content string, score float64) domain.SearchResult {
	return domain.SearchResult{
		Document: &domain.Document{
			ID:             id,
			Language:       language,
			Category:       category,
			Subject:        id,
			Project:        project,
			TemporalStatus: domain.TemporalTimeless,
			Keywords:       []string{id},
			Content:        content,
		},
		FinalScore: score,
		Sources:    []string{"vector", "lexical"},
	}
}
