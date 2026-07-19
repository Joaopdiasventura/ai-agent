package search

import (
	"ai-agent/internal/domain"
	"ai-agent/internal/knowledge"
	"ai-agent/internal/nlp"
	"fmt"
	"strings"
	"testing"
)

type RetrievalCase struct {
	Name               string
	Question           string
	Language           nlp.Language
	ExpectedDocumentID string
	ExpectedCategory   string
	ExpectedTerms      []string
	ForbiddenTerms     []string
}

func TestRetrievalCatalogCoversEveryDocument(t *testing.T) {
	documents := knowledge.Documents()
	cases := buildRetrievalCatalog(documents)

	casesByDocumentID := make(map[string][]RetrievalCase)
	for _, testCase := range cases {
		casesByDocumentID[testCase.ExpectedDocumentID] = append(casesByDocumentID[testCase.ExpectedDocumentID], testCase)
	}

	for _, document := range documents {
		if len(casesByDocumentID[document.ID]) == 0 {
			t.Fatalf("document %q does not have a retrieval case", document.ID)
		}
	}

	assertLanguagePairs(t, documents)

	engine := newTestEngine()
	recovered := 0
	wrongLanguage := 0
	incoherent := 0

	for _, testCase := range cases {
		t.Run(testCase.Name, func(t *testing.T) {
			searchResult := engine.Search(testCase.Question, 5)

			if !searchResult.Found {
				t.Fatalf("Search(%q) did not find a result", testCase.Question)
			}

			if searchResult.Language != testCase.Language {
				wrongLanguage++
				t.Fatalf("Search(%q) language = %q, want %q",
					testCase.Question,
					searchResult.Language,
					testCase.Language,
				)
			}

			topDocument := searchResult.Results[0].Document
			if topDocument.ID != testCase.ExpectedDocumentID {
				t.Fatalf("Search(%q) top document = %q, want %q",
					testCase.Question,
					topDocument.ID,
					testCase.ExpectedDocumentID,
				)
			}

			if topDocument.Category != testCase.ExpectedCategory {
				t.Fatalf("Search(%q) top category = %q, want %q",
					testCase.Question,
					topDocument.Category,
					testCase.ExpectedCategory,
				)
			}

			if topDocument.Language != string(testCase.Language) {
				wrongLanguage++
				t.Fatalf("Search(%q) top language = %q, want %q",
					testCase.Question,
					topDocument.Language,
					testCase.Language,
				)
			}

			content := strings.ToLower(topDocument.Content)
			for _, expectedTerm := range testCase.ExpectedTerms {
				if !strings.Contains(content, strings.ToLower(expectedTerm)) {
					incoherent++
					t.Fatalf("document %q does not contain expected term %q", topDocument.ID, expectedTerm)
				}
			}

			for _, forbiddenTerm := range testCase.ForbiddenTerms {
				if strings.Contains(content, strings.ToLower(forbiddenTerm)) {
					incoherent++
					t.Fatalf("document %q contains forbidden term %q", topDocument.ID, forbiddenTerm)
				}
			}

			recovered++
		})
	}

	coverage := float64(recovered) / float64(len(documents)) * 100
	t.Logf(
		"retrieval coverage: documents=%d cases=%d recovered=%d wrong_language=%d incoherent=%d coverage=%.2f%%",
		len(documents),
		len(cases),
		recovered,
		wrongLanguage,
		incoherent,
		coverage,
	)

	if recovered != len(documents) {
		t.Fatalf("retrieval coverage = %.2f%%, want 100%%", coverage)
	}
}

func buildRetrievalCatalog(documents []*domain.Document) []RetrievalCase {
	cases := make([]RetrievalCase, 0, len(documents))

	for _, document := range documents {
		language := nlp.Language(document.Language)
		question := portugueseRetrievalQuestion(document)

		if language == nlp.LanguageEnglish {
			question = englishRetrievalQuestion(document)
		}

		forbiddenTerms := []string{}
		if language == nlp.LanguagePortuguese {
			forbiddenTerms = []string{"To contact", "has worked", "can build"}
		} else {
			forbiddenTerms = []string{"Para entrar", "trabalha como", "pode desenvolver"}
		}

		cases = append(cases, RetrievalCase{
			Name:               document.ID,
			Question:           question,
			Language:           language,
			ExpectedDocumentID: document.ID,
			ExpectedCategory:   document.Category,
			ExpectedTerms:      expectedTermsForDocument(document),
			ForbiddenTerms:     forbiddenTerms,
		})
	}

	return cases
}

func portugueseRetrievalQuestion(document *domain.Document) string {
	switch document.Category {
	case "contact":
		return "Qual é este contato de João Paulo: " + document.Content
	case "education":
		return "Formação de João Paulo: " + document.Content
	case "career":
		return "Experiência profissional de João Paulo: " + document.Content
	case "technology":
		return "Quais tecnologias aparecem neste fato: " + document.Content
	case "project", "comparison":
		return "Projeto ou comparação do portfólio: " + document.Content
	case "impact":
		return "Qual foi o impacto deste trabalho: " + document.Content
	case "service":
		return "Que serviço este fato descreve: " + document.Content
	case "profile":
		return "Perfil profissional de João Paulo: " + document.Content
	case "certificate":
		return "Certificação de João Paulo: " + document.Content
	default:
		return "Informação de João Paulo: " + document.Content
	}
}

func englishRetrievalQuestion(document *domain.Document) string {
	switch document.Category {
	case "contact":
		return "What contact detail is described here: " + document.Content
	case "education":
		return "Education information about João Paulo: " + document.Content
	case "career":
		return "Professional experience about João Paulo: " + document.Content
	case "technology":
		return "Which technologies appear in this fact: " + document.Content
	case "project", "comparison":
		return "Portfolio project or comparison: " + document.Content
	case "impact":
		return "What impact is described in this work: " + document.Content
	case "service":
		return "What service does this fact describe: " + document.Content
	case "profile":
		return "Professional profile of João Paulo: " + document.Content
	case "certificate":
		return "Certification of João Paulo: " + document.Content
	default:
		return "Information about João Paulo: " + document.Content
	}
}

func expectedTermsForDocument(document *domain.Document) []string {
	content := strings.ToLower(document.Content)

	importantTerms := []string{
		"joão paulo",
		"ufind",
		"representa online",
		"fiap",
		"etec",
		"auronix",
		"x tube",
		"ggcompress",
		"auditex",
		"joaopdias.dev@gmail.com",
		"+55 (11) 93445-3236",
		"linkedin.com/in/joaopdias-dev",
		"github.com/joaopdiasventura",
		"1.23 gb/s",
		"9.77 gb",
	}

	for _, term := range importantTerms {
		if strings.Contains(content, term) {
			return []string{term}
		}
	}

	words := strings.Fields(content)
	for _, word := range words {
		word = strings.Trim(word, ".,;:!?()")
		if len([]rune(word)) >= 6 {
			return []string{word}
		}
	}

	return []string{strings.TrimSpace(document.Content)}
}

func assertLanguagePairs(t *testing.T, documents []*domain.Document) {
	t.Helper()

	byID := make(map[string]*domain.Document, len(documents))
	for _, document := range documents {
		if _, exists := byID[document.ID]; exists {
			t.Fatalf("duplicated document ID %q", document.ID)
		}
		byID[document.ID] = document
	}

	for _, document := range documents {
		baseID, suffix, ok := strings.Cut(document.ID, "-")
		if !ok {
			t.Fatalf("document %q has no language suffix", document.ID)
		}

		lastDash := strings.LastIndex(document.ID, "-")
		baseID = document.ID[:lastDash]
		suffix = document.ID[lastDash+1:]

		if suffix != document.Language {
			t.Fatalf("document %q suffix = %q, language = %q", document.ID, suffix, document.Language)
		}

		counterpartSuffix := "en"
		if suffix == "en" {
			counterpartSuffix = "pt"
		}

		counterpartID := fmt.Sprintf("%s-%s", baseID, counterpartSuffix)
		counterpart, exists := byID[counterpartID]
		if !exists {
			t.Fatalf("document %q does not have counterpart %q", document.ID, counterpartID)
		}

		if counterpart.Category != document.Category {
			t.Fatalf("document %q category = %q, counterpart %q category = %q",
				document.ID,
				document.Category,
				counterpart.ID,
				counterpart.Category,
			)
		}
	}
}
