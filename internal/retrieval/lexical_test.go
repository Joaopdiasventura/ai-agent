package retrieval

import (
	"ai-agent/internal/domain"
	"ai-agent/internal/vectorindex"
	"testing"
)

func TestLexicalSearchFindsExactContactTerms(t *testing.T) {
	index := testLexicalIndex([]domain.Document{
		{ID: "contact-email-pt", Language: "pt", Category: "contact", Subject: "email", TemporalStatus: domain.TemporalTimeless, Keywords: []string{"email"}, Content: "Para contato use joaopdias.dev@gmail.com."},
		{ID: "contact-phone-pt", Language: "pt", Category: "contact", Subject: "phone", TemporalStatus: domain.TemporalTimeless, Keywords: []string{"telefone"}, Content: "Para telefone use +55 (11) 93445-3236."},
	})

	results := LexicalSearch(index, domain.Query{Text: "qual email dele joaopdias.dev@gmail.com"}, 1)

	if len(results) != 1 || results[0].Document.ID != "contact-email-pt" {
		t.Fatalf("LexicalSearch() returned %#v, want contact-email-pt", results)
	}

	results = LexicalSearch(index, domain.Query{Text: "telefone 93445-3236"}, 1)

	if len(results) != 1 || results[0].Document.ID != "contact-phone-pt" {
		t.Fatalf("LexicalSearch() returned %#v, want contact-phone-pt", results)
	}
}

func TestLexicalSearchHandlesAccentsTechnologiesAndProjects(t *testing.T) {
	index := testLexicalIndex([]domain.Document{
		{ID: "project-auronix-pt", Language: "pt", Category: "project", Subject: "auronix", Project: "auronix", TemporalStatus: domain.TemporalTimeless, Keywords: []string{"aws", "eks", "kubernetes"}, Content: "Auronix usa AWS EKS e Kubernetes."},
		{ID: "project-auditex-pt", Language: "pt", Category: "project", Subject: "auditex", Project: "auditex", TemporalStatus: domain.TemporalTimeless, Keywords: []string{"auditabilidade"}, Content: "Auditex demonstra auditabilidade."},
	})

	results := LexicalSearch(index, domain.Query{Text: "auronix aws eks"}, 2)

	if len(results) == 0 || results[0].Document.ID != "project-auronix-pt" {
		t.Fatalf("LexicalSearch() returned %#v, want project-auronix-pt first", results)
	}

	results = LexicalSearch(index, domain.Query{Text: "auditabilidade"}, 1)

	if len(results) != 1 || results[0].Document.ID != "project-auditex-pt" {
		t.Fatalf("LexicalSearch() returned %#v, want project-auditex-pt", results)
	}
}

func TestLexicalSearchHandlesEmptyInputs(t *testing.T) {
	index := testLexicalIndex([]domain.Document{
		{ID: "a", Language: "pt", Category: "profile", Subject: "a", TemporalStatus: domain.TemporalTimeless, Keywords: []string{"a"}, Content: "content"},
	})

	if results := LexicalSearch(index, domain.Query{Text: ""}, 5); len(results) != 0 {
		t.Fatalf("LexicalSearch() returned %d results for empty query", len(results))
	}

	if results := LexicalSearch(index, domain.Query{Text: "content"}, 0); len(results) != 0 {
		t.Fatalf("LexicalSearch() returned %d results for zero limit", len(results))
	}
}

func testLexicalIndex(documents []domain.Document) vectorindex.Index {
	entries := make([]vectorindex.Entry, 0, len(documents))

	for _, document := range documents {
		entries = append(entries, vectorindex.Entry{
			Document:  document,
			Embedding: []float32{1},
		})
	}

	return vectorindex.Index{
		Version:   vectorindex.Version,
		Model:     "test",
		Dimension: 1,
		Entries:   entries,
	}
}
