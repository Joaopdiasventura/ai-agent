package knowledge

import (
	"strings"
	"testing"
)

func TestDocumentsStaticBaseInvariants(t *testing.T) {
	documents := Documents()

	if len(documents) != 140 {
		t.Fatalf("Documents() returned %d documents, want 140", len(documents))
	}

	seenIDs := make(map[string]struct{}, len(documents))
	languageCounts := map[string]int{
		"pt": 0,
		"en": 0,
	}
	categoryCounts := map[string]int{}

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
	}

	if languageCounts["pt"] != 70 {
		t.Fatalf("portuguese document count = %d, want 70", languageCounts["pt"])
	}

	if languageCounts["en"] != 70 {
		t.Fatalf("english document count = %d, want 70", languageCounts["en"])
	}

	if categoryCounts["impact"] == 0 {
		t.Fatal("Documents() has no impact documents")
	}
}
