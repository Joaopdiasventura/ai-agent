package knowledge

import (
	"strings"
	"testing"
)

func TestDocumentsStaticBaseInvariants(t *testing.T) {
	documents := Documents()

	if len(documents) != 134 {
		t.Fatalf("Documents() returned %d documents, want 134", len(documents))
	}

	seenIDs := make(map[string]struct{}, len(documents))
	languageCounts := map[string]int{
		"pt": 0,
		"en": 0,
	}
	categoryCounts := map[string]int{}
	contents := strings.Builder{}

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

		if strings.TrimSpace(document.Language) == "" {
			t.Fatalf("document %q has empty language", document.ID)
		}

		if strings.TrimSpace(document.Content) == "" {
			t.Fatalf("document %q has empty content", document.ID)
		}

		if _, exists := seenIDs[document.ID]; exists {
			t.Fatalf("document ID %q is duplicated", document.ID)
		}

		seenIDs[document.ID] = struct{}{}

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
