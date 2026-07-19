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

func TestSearchProjectRecommendationsSelectExpectedProjects(t *testing.T) {
	tests := []struct {
		name              string
		question          string
		expectedProject   string
		expectedCriterion string
		expectedIDs       []string
		forbiddenIDs      []string
	}{
		{
			name:              "complex problem selects x tube",
			question:          "Qual projeto do João Paulo melhor demonstra capacidade de resolver problemas complexos e qual impacto ele gerou?",
			expectedProject:   "X Tube",
			expectedCriterion: "complex_problem",
			expectedIDs:       []string{"project-xtube-description-pt", "project-xtube-leadership-pt", "project-xtube-processing-pt"},
			forbiddenIDs:      []string{"project-comparison-best-pt"},
		},
		{
			name:              "technical capability selects auronix",
			question:          "Qual projeto do João Paulo melhor demonstra capacidade técnica?",
			expectedProject:   "Auronix",
			expectedCriterion: "technical_capability",
			expectedIDs:       []string{"project-auronix-description-pt", "project-auronix-consistency-pt"},
			forbiddenIDs:      []string{"project-comparison-best-pt", "project-xtube"},
		},
		{
			name:              "recruiter recommendation selects auronix",
			question:          "Que projeto você destacaria para um recrutador?",
			expectedProject:   "Auronix",
			expectedCriterion: "general_recommendation",
			expectedIDs:       []string{"project-auronix-description-pt", "project-auronix-consistency-pt"},
			forbiddenIDs:      []string{"project-comparison-best-pt", "project-xtube"},
		},
		{
			name:              "financial systems selects auronix",
			question:          "Qual projeto melhor demonstra experiência com sistemas financeiros?",
			expectedProject:   "Auronix",
			expectedCriterion: "financial_systems",
			expectedIDs:       []string{"project-auronix-description-pt", "project-auronix-consistency-pt"},
			forbiddenIDs:      []string{"project-comparison-best-pt"},
		},
		{
			name:              "technical leadership selects x tube",
			question:          "Qual projeto melhor demonstra liderança técnica?",
			expectedProject:   "X Tube",
			expectedCriterion: "technical_leadership",
			expectedIDs:       []string{"project-xtube-leadership-pt"},
			forbiddenIDs:      []string{"project-auronix", "project-comparison-best-pt"},
		},
		{
			name:              "go performance selects ggcompress",
			question:          "Qual projeto demonstra melhor desempenho e concorrência?",
			expectedProject:   "GGCompress",
			expectedCriterion: "go_performance",
			expectedIDs:       []string{"project-ggcompress-concurrency-pt", "project-ggcompress-performance-pt"},
			forbiddenIDs:      []string{"project-comparison-best-pt"},
		},
		{
			name:              "auditability selects auditex",
			question:          "Qual projeto melhor demonstra preocupação com auditabilidade?",
			expectedProject:   "Auditex",
			expectedCriterion: "auditability",
			expectedIDs:       []string{"project-auditex-description-pt", "project-auditex-integrity-pt"},
			forbiddenIDs:      []string{"project-ggcompress", "project-comparison-best-pt"},
		},
	}

	engine := newTestEngine()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			searchResult := engine.Search(test.question, 5)

			if !searchResult.Found {
				t.Fatalf("Search(%q) did not find a result", test.question)
			}

			if searchResult.Intent != "project_recommendation" {
				t.Fatalf("Search(%q) intent = %q, want project_recommendation", test.question, searchResult.Intent)
			}

			if searchResult.SelectedProject != test.expectedProject {
				t.Fatalf("Search(%q) selected project = %q, want %q",
					test.question,
					searchResult.SelectedProject,
					test.expectedProject,
				)
			}

			if string(searchResult.ProjectCriterion) != test.expectedCriterion {
				t.Fatalf("Search(%q) criterion = %q, want %q",
					test.question,
					searchResult.ProjectCriterion,
					test.expectedCriterion,
				)
			}

			resultIDs := make([]string, 0, len(searchResult.Results))
			for _, result := range searchResult.Results {
				if result.Document.Language != "pt" {
					t.Fatalf("Search(%q) returned document %q in language %q",
						test.question,
						result.Document.ID,
						result.Document.Language,
					)
				}

				resultIDs = append(resultIDs, result.Document.ID)

				for _, forbiddenID := range test.forbiddenIDs {
					if strings.Contains(result.Document.ID, forbiddenID) {
						t.Fatalf("Search(%q) returned forbidden document %q", test.question, result.Document.ID)
					}
				}
			}

			for _, expectedID := range test.expectedIDs {
				found := false
				for _, resultID := range resultIDs {
					if resultID == expectedID {
						found = true
						break
					}
				}

				if !found {
					t.Fatalf("Search(%q) returned documents %v, want %q", test.question, resultIDs, expectedID)
				}
			}
		})
	}
}
