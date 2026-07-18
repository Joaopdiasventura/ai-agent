package search

import (
	"ai-agent/internal/knowledge"
	"strings"
	"testing"
)

func newTestEngine() *Engine {
	return NewEngine(knowledge.Documents(), 0.1)
}

func TestSearchKeepsResultsInDetectedLanguage(t *testing.T) {
	tests := []struct {
		name     string
		question string
		language string
	}{
		{
			name:     "portuguese question returns portuguese documents",
			question: "Me fale sobre o auronix",
			language: "pt",
		},
		{
			name:     "english question returns english documents",
			question: "Tell me about Auronix",
			language: "en",
		},
	}

	engine := newTestEngine()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			searchResult := engine.Search(test.question, 5)

			if !searchResult.Found {
				t.Fatalf("Search(%q) did not find results", test.question)
			}

			for _, result := range searchResult.Results {
				if result.Document.Language != test.language {
					t.Fatalf("Search(%q) returned document %q with language %q, want %q",
						test.question,
						result.Document.ID,
						result.Document.Language,
						test.language,
					)
				}
			}
		})
	}
}

func TestSearchFindsPortugueseAuronixDocument(t *testing.T) {
	searchResult := newTestEngine().Search("Me fale sobre o auronix", 5)

	if !searchResult.Found {
		t.Fatal("Search() did not find Auronix results")
	}

	for _, result := range searchResult.Results {
		if strings.Contains(result.Document.ID, "project-auronix") &&
			result.Document.Language == "pt" {
			return
		}
	}

	t.Fatalf("Search() did not return a portuguese Auronix document: %#v", searchResult.Results)
}

func TestSearchTechnologiesQuestionReturnsMultipleTechnologyResults(t *testing.T) {
	searchResult := newTestEngine().Search("What technologies does João use?", 5)

	if !searchResult.Found {
		t.Fatal("Search() did not find technology results")
	}

	if len(searchResult.Results) < 2 {
		t.Fatalf("Search() returned %d results, want multiple results", len(searchResult.Results))
	}

	for _, result := range searchResult.Results {
		if result.Document.Language != "en" {
			t.Fatalf("Search() returned non-english document %q", result.Document.ID)
		}
	}
}

func TestSearchInvalidInputsDoNotFindResults(t *testing.T) {
	engine := newTestEngine()

	tests := []struct {
		name     string
		question string
		limit    int
	}{
		{
			name:     "limit less than or equal to zero",
			question: "Tell me about Auronix",
			limit:    0,
		},
		{
			name:     "question without tokens",
			question: "   ???   ",
			limit:    5,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			searchResult := engine.Search(test.question, test.limit)

			if searchResult.Found {
				t.Fatalf("Search(%q, %d) found results, want not found", test.question, test.limit)
			}
		})
	}
}
