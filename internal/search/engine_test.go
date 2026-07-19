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

func TestSearchFindsPriorityProjects(t *testing.T) {
	tests := []struct {
		name       string
		question   string
		documentID string
		language   string
	}{
		{
			name:       "english x tube question",
			question:   "Tell me about X Tube",
			documentID: "project-xtube",
			language:   "en",
		},
		{
			name:       "portuguese auditex question",
			question:   "O que é o Auditex?",
			documentID: "project-auditex",
			language:   "pt",
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
				if strings.Contains(result.Document.ID, test.documentID) &&
					result.Document.Language == test.language {
					return
				}
			}

			t.Fatalf("Search(%q) did not return %q document in %q: %#v",
				test.question,
				test.documentID,
				test.language,
				searchResult.Results,
			)
		})
	}
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

func TestSearchRegressionExamplesReturnExpectedTopDocuments(t *testing.T) {
	tests := []struct {
		name       string
		question   string
		documentID string
		language   string
	}{
		{
			name:       "current education in portuguese",
			question:   "onde joão estuda?",
			documentID: "education-fiap-pt",
			language:   "pt",
		},
		{
			name:       "past education in portuguese",
			question:   "onde joão estudou?",
			documentID: "education-etec-pt",
			language:   "pt",
		},
		{
			name:       "email pronoun in portuguese",
			question:   "qual o email dele?",
			documentID: "contact-email-pt",
			language:   "pt",
		},
		{
			name:       "current job in portuguese",
			question:   "onde joão trabalha?",
			documentID: "career-current-job-pt",
			language:   "pt",
		},
		{
			name:       "past job in portuguese",
			question:   "onde joão trabalhou?",
			documentID: "career-junior-job-pt",
			language:   "pt",
		},
		{
			name:       "first job in portuguese",
			question:   "qual foi o primeiro emprego?",
			documentID: "career-intern-job-pt",
			language:   "pt",
		},
		{
			name:       "past job in english",
			question:   "Where did João work before?",
			documentID: "career-junior-job-en",
			language:   "en",
		},
	}

	engine := newTestEngine()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			searchResult := engine.Search(test.question, 5)

			if !searchResult.Found {
				t.Fatalf("Search(%q) did not find a result", test.question)
			}

			if string(searchResult.Language) != test.language {
				t.Fatalf("Search(%q) language = %q, want %q", test.question, searchResult.Language, test.language)
			}

			if len(searchResult.Results) == 0 {
				t.Fatalf("Search(%q) returned no results", test.question)
			}

			topDocument := searchResult.Results[0].Document
			if topDocument.ID != test.documentID {
				t.Fatalf("Search(%q) top document = %q, want %q", test.question, topDocument.ID, test.documentID)
			}

			for _, result := range searchResult.Results {
				if result.Document.Language != test.language {
					t.Fatalf("Search(%q) returned %q with language %q, want %q",
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

func TestSearchShortAmbiguousQuestionsArePredictable(t *testing.T) {
	tests := []struct {
		question   string
		documentID string
		language   string
	}{
		{question: "João", documentID: "identity-basic-pt", language: "pt"},
		{question: "FIAP", documentID: "education-fiap-pt", language: "pt"},
		{question: "email", documentID: "contact-email-pt", language: "pt"},
		{question: "projetos", documentID: "project-comparison-best-pt", language: "pt"},
		{question: "tecnologias", documentID: "technology-messaging-pt", language: "pt"},
		{question: "trabalho", documentID: "career-current-job-pt", language: "pt"},
		{question: "formação", documentID: "education-etec-pt", language: "pt"},
		{question: "Auronix", documentID: "project-auronix-description-pt", language: "pt"},
		{question: "X Tube", documentID: "project-xtube-description-pt", language: "pt"},
		{question: "GGCompress", documentID: "project-ggcompress-description-pt", language: "pt"},
		{question: "Auditex", documentID: "project-auditex-description-pt", language: "pt"},
	}

	engine := newTestEngine()

	for _, test := range tests {
		t.Run(test.question, func(t *testing.T) {
			searchResult := engine.Search(test.question, 5)

			if !searchResult.Found {
				t.Fatalf("Search(%q) did not find a result", test.question)
			}

			topDocument := searchResult.Results[0].Document
			if topDocument.ID != test.documentID {
				t.Fatalf("Search(%q) top document = %q, want %q", test.question, topDocument.ID, test.documentID)
			}

			if topDocument.Language != test.language {
				t.Fatalf("Search(%q) top language = %q, want %q", test.question, topDocument.Language, test.language)
			}
		})
	}
}

func TestSearchNegativeCategoryAndLanguageRegressions(t *testing.T) {
	tests := []struct {
		name             string
		question         string
		forbiddenIDParts []string
		forbiddenTerms   []string
	}{
		{
			name:             "education does not return contact",
			question:         "onde joão estuda?",
			forbiddenIDParts: []string{"contact-"},
			forbiddenTerms:   []string{"+55", "joaopdias.dev@gmail.com"},
		},
		{
			name:             "email does not return projects",
			question:         "qual o email dele?",
			forbiddenIDParts: []string{"project-"},
			forbiddenTerms:   []string{"Auronix", "GGCompress", "X Tube", "Auditex"},
		},
		{
			name:             "current job does not return internship",
			question:         "onde joão trabalha?",
			forbiddenIDParts: []string{"career-intern", "career-junior"},
			forbiddenTerms:   []string{"Estagiário", "Júnior"},
		},
		{
			name:             "auronix does not return ggcompress",
			question:         "Me fale sobre o Auronix",
			forbiddenIDParts: []string{"project-ggcompress"},
			forbiddenTerms:   []string{"GGCompress"},
		},
	}

	engine := newTestEngine()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			searchResult := engine.Search(test.question, 5)

			if !searchResult.Found {
				t.Fatalf("Search(%q) did not find a result", test.question)
			}

			for _, result := range searchResult.Results {
				for _, forbiddenIDPart := range test.forbiddenIDParts {
					if strings.Contains(result.Document.ID, forbiddenIDPart) {
						t.Fatalf("Search(%q) returned forbidden document %q", test.question, result.Document.ID)
					}
				}

				for _, forbiddenTerm := range test.forbiddenTerms {
					if strings.Contains(result.Document.Content, forbiddenTerm) {
						t.Fatalf("Search(%q) returned forbidden term %q in document %q",
							test.question,
							forbiddenTerm,
							result.Document.ID,
						)
					}
				}
			}
		})
	}
}
