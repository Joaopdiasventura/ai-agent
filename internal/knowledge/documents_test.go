package knowledge

import (
	"ai-agent/internal/domain"
	"strings"
	"testing"
)

func TestDocumentsStaticBaseInvariants(t *testing.T) {
	documents := Documents()

	if len(documents) != 134 {
		t.Fatalf("Documents() returned %d documents, want 134", len(documents))
	}

	seenIDs := make(map[string]struct{}, len(documents))
	documentsByID := make(map[string]*domain.Document, len(documents))
	languageCounts := map[string]int{
		"pt": 0,
		"en": 0,
	}
	categoryCounts := map[string]int{}
	contents := strings.Builder{}
	validCategories := map[string]struct{}{
		"identity":    {},
		"profile":     {},
		"technology":  {},
		"impact":      {},
		"career":      {},
		"education":   {},
		"certificate": {},
		"contact":     {},
		"service":     {},
		"project":     {},
		"comparison":  {},
	}
	validTemporalStatuses := map[string]struct{}{
		domain.TemporalPast:     {},
		domain.TemporalCurrent:  {},
		domain.TemporalFuture:   {},
		domain.TemporalTimeless: {},
	}
	validProjects := map[string]struct{}{
		"":           {},
		"auronix":    {},
		"x-tube":     {},
		"ggcompress": {},
		"auditex":    {},
	}

	for _, document := range documents {
		if document == nil {
			t.Fatal("Documents() returned a nil document")
		}

		if strings.TrimSpace(document.ID) == "" {
			t.Fatal("document has empty ID")
		}

		if strings.TrimSpace(document.Category) == "" {
			t.Fatalf("document %q has empty category", document.ID)
		}

		if _, valid := validCategories[document.Category]; !valid {
			t.Fatalf("document %q has unsupported category %q", document.ID, document.Category)
		}

		if strings.TrimSpace(document.Language) == "" {
			t.Fatalf("document %q has empty language", document.ID)
		}

		if strings.TrimSpace(document.Content) == "" {
			t.Fatalf("document %q has empty content", document.ID)
		}

		if strings.TrimSpace(document.Subject) == "" {
			t.Fatalf("document %q has empty subject", document.ID)
		}

		if _, valid := validTemporalStatuses[document.TemporalStatus]; !valid {
			t.Fatalf("document %q has unsupported temporal status %q", document.ID, document.TemporalStatus)
		}

		if _, valid := validProjects[document.Project]; !valid {
			t.Fatalf("document %q has unsupported project %q", document.ID, document.Project)
		}

		if len(document.Keywords) == 0 {
			t.Fatalf("document %q has no keywords", document.ID)
		}

		if _, exists := seenIDs[document.ID]; exists {
			t.Fatalf("document ID %q is duplicated", document.ID)
		}

		seenIDs[document.ID] = struct{}{}
		documentsByID[document.ID] = document

		switch document.Language {
		case "pt":
			if !strings.HasSuffix(document.ID, "-pt") {
				t.Fatalf("portuguese document %q does not end with -pt", document.ID)
			}
		case "en":
			if !strings.HasSuffix(document.ID, "-en") {
				t.Fatalf("english document %q does not end with -en", document.ID)
			}
		default:
			t.Fatalf("document %q has unsupported language %q", document.ID, document.Language)
		}

		languageCounts[document.Language]++
		categoryCounts[document.Category]++
		contents.WriteString(document.ID)
		contents.WriteString(" ")
		contents.WriteString(document.Content)
		contents.WriteString(" ")
	}

	for _, document := range documents {
		pairID := pairedDocumentID(document.ID)
		pair := documentsByID[pairID]
		if pair == nil {
			t.Fatalf("document %q has no language pair %q", document.ID, pairID)
		}

		if pair.Category != document.Category {
			t.Fatalf("document %q category = %q, pair %q category = %q", document.ID, document.Category, pair.ID, pair.Category)
		}

		if pair.Project != document.Project {
			t.Fatalf("document %q project = %q, pair %q project = %q", document.ID, document.Project, pair.ID, pair.Project)
		}

		if pair.TemporalStatus != document.TemporalStatus {
			t.Fatalf("document %q temporal status = %q, pair %q temporal status = %q", document.ID, document.TemporalStatus, pair.ID, pair.TemporalStatus)
		}
	}

	expectedTemporalStatuses := map[string]string{
		"education-fiap-pt":     domain.TemporalFuture,
		"education-etec-pt":     domain.TemporalPast,
		"career-current-job-pt": domain.TemporalCurrent,
		"career-junior-job-pt":  domain.TemporalPast,
		"career-intern-job-pt":  domain.TemporalPast,
	}

	for id, expectedTemporalStatus := range expectedTemporalStatuses {
		document := documentsByID[id]
		if document == nil {
			t.Fatalf("expected document %q was not found", id)
		}

		if document.TemporalStatus != expectedTemporalStatus {
			t.Fatalf("document %q temporal status = %q, want %q", id, document.TemporalStatus, expectedTemporalStatus)
		}
	}

	if languageCounts["pt"] != 67 {
		t.Fatalf("portuguese document count = %d, want 67", languageCounts["pt"])
	}

	if languageCounts["en"] != 67 {
		t.Fatalf("english document count = %d, want 67", languageCounts["en"])
	}

	if categoryCounts["impact"] == 0 {
		t.Fatal("Documents() has no impact documents")
	}

	allContent := strings.ToLower(contents.String())

	requiredTerms := []string{
		"+55 (11) 93445-3236",
		"auronix",
		"x tube",
		"ggcompress",
		"auditex",
		"1.23 gb/s",
		"9.77 gb",
	}

	for _, term := range requiredTerms {
		if !strings.Contains(allContent, term) {
			t.Fatalf("Documents() does not contain required term %q", term)
		}
	}

	forbiddenTerms := []string{
		"+55 (11) 986" + "55-3558",
		"project-" + "modularis",
		"project-" + "votrix",
		"project-" + "vox",
		"project-" + "etecfy",
		"capacitor",
		"tauri",
		"bullmq",
		"fastify",
	}

	for _, term := range forbiddenTerms {
		if strings.Contains(allContent, term) {
			t.Fatalf("Documents() contains forbidden term %q", term)
		}
	}
}

func pairedDocumentID(id string) string {
	switch {
	case strings.HasSuffix(id, "-pt"):
		return strings.TrimSuffix(id, "-pt") + "-en"
	case strings.HasSuffix(id, "-en"):
		return strings.TrimSuffix(id, "-en") + "-pt"
	default:
		return ""
	}
}
