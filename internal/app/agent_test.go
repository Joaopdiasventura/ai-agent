package app

import (
	"strings"
	"testing"
)

func TestAgentResponseReturnsPortugueseAnswerForPortugueseQuestion(t *testing.T) {
	response, hasResponse, language := AgentResponse("Me fale sobre o auronix")

	if !hasResponse {
		t.Fatal("AgentResponse() did not find a response")
	}

	if language != "pt" {
		t.Fatalf("AgentResponse() language = %q, want %q", language, "pt")
	}

	if !strings.Contains(response, "Auronix") {
		t.Fatalf("AgentResponse() response %q does not mention Auronix", response)
	}
}

func TestAgentResponseReturnsEnglishAnswerForEnglishQuestion(t *testing.T) {
	response, hasResponse, language := AgentResponse("Tell me about Auronix")

	if !hasResponse {
		t.Fatal("AgentResponse() did not find a response")
	}

	if language != "en" {
		t.Fatalf("AgentResponse() language = %q, want %q", language, "en")
	}

	if !strings.Contains(response, "Auronix") {
		t.Fatalf("AgentResponse() response %q does not mention Auronix", response)
	}
}

func TestAgentResponseReturnsLocalizedFallbackForUnknownQuestion(t *testing.T) {
	tests := []struct {
		name             string
		question         string
		expectedLanguage string
		expectedFallback string
	}{
		{
			name:             "portuguese unknown question",
			question:         "Pergunta sem relação",
			expectedLanguage: "pt",
			expectedFallback: "Não encontrei essa informação específica",
		},
		{
			name:             "english unknown question",
			question:         "Unrelated question",
			expectedLanguage: "en",
			expectedFallback: "I don't have that specific information",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, hasResponse, language := AgentResponse(test.question)

			if hasResponse {
				t.Fatalf("AgentResponse(%q) found a response, want fallback path", test.question)
			}

			if language != test.expectedLanguage {
				t.Fatalf("AgentResponse(%q) language = %q, want %q",
					test.question,
					language,
					test.expectedLanguage,
				)
			}

			fallback := NotFoundMessage(language)
			if !strings.Contains(fallback, test.expectedFallback) {
				t.Fatalf("NotFoundMessage(%q) = %q, want it to contain %q",
					language,
					fallback,
					test.expectedFallback,
				)
			}
		})
	}
}

func TestAgentResponseRegressionExamples(t *testing.T) {
	tests := []struct {
		name           string
		question       string
		language       string
		expectedTerms  []string
		forbiddenTerms []string
	}{
		{
			name:          "current education",
			question:      "onde joão estuda?",
			language:      "pt",
			expectedTerms: []string{"FIAP", "Inteligência Artificial"},
			forbiddenTerms: []string{
				"FIAP segue estudando",
				"Etec",
			},
		},
		{
			name:          "past education",
			question:      "onde joão estudou?",
			language:      "pt",
			expectedTerms: []string{"Etec de Guarulhos", "Desenvolvimento de Sistemas"},
			forbiddenTerms: []string{
				"Engenheiro de Software Full Stack",
				"trabalha como",
			},
		},
		{
			name:          "email pronoun",
			question:      "qual o email dele?",
			language:      "pt",
			expectedTerms: []string{"joaopdias.dev@gmail.com"},
			forbiddenTerms: []string{
				"To contact",
				"phone",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response, hasResponse, language := AgentResponse(test.question)

			if !hasResponse {
				t.Fatalf("AgentResponse(%q) did not find a response", test.question)
			}

			if language != test.language {
				t.Fatalf("AgentResponse(%q) language = %q, want %q", test.question, language, test.language)
			}

			for _, expectedTerm := range test.expectedTerms {
				if !strings.Contains(response, expectedTerm) {
					t.Fatalf("AgentResponse(%q) = %q, want it to contain %q",
						test.question,
						response,
						expectedTerm,
					)
				}
			}

			for _, forbiddenTerm := range test.forbiddenTerms {
				if strings.Contains(response, forbiddenTerm) {
					t.Fatalf("AgentResponse(%q) = %q, want it not to contain %q",
						test.question,
						response,
						forbiddenTerm,
					)
				}
			}
		})
	}
}

func TestAgentResponseProjectRecommendationRegression(t *testing.T) {
	response, hasResponse, language := AgentResponse("Qual projeto do João Paulo melhor demonstra capacidade de resolver problemas complexos e qual impacto ele gerou?")

	if !hasResponse {
		t.Fatal("AgentResponse() did not find a response")
	}

	if language != "pt" {
		t.Fatalf("AgentResponse() language = %q, want %q", language, "pt")
	}

	expectedTerms := []string{
		"X Tube",
		"liderança técnica",
		"upload",
		"processamento",
		"entrega",
		"Amazon S3",
		"Amazon SQS",
		"Kafka",
		"Prometheus",
	}

	for _, term := range expectedTerms {
		if !strings.Contains(response, term) {
			t.Fatalf("AgentResponse() = %q, want it to contain %q", response, term)
		}
	}

	forbiddenTerms := []string{
		"projetos prioritários",
		"Auronix",
		"GGCompress",
		"Auditex",
		"Entre os projetos do portfólio,",
	}

	for _, term := range forbiddenTerms {
		if strings.Contains(response, term) {
			t.Fatalf("AgentResponse() = %q, want it not to contain %q", response, term)
		}
	}
}

func TestAgentResponseProjectRecommendationVariations(t *testing.T) {
	tests := []struct {
		question        string
		expectedProject string
	}{
		{
			question:        "Qual projeto do João Paulo melhor demonstra capacidade técnica?",
			expectedProject: "Auronix",
		},
		{
			question:        "Qual projeto mais demonstra a capacidade técnica do João?",
			expectedProject: "Auronix",
		},
		{
			question:        "Qual projeto dele é mais complexo?",
			expectedProject: "X Tube",
		},
		{
			question:        "Que projeto melhor mostra que ele sabe resolver problemas difíceis?",
			expectedProject: "X Tube",
		},
		{
			question:        "Entre os projetos, qual teve o maior desafio técnico?",
			expectedProject: "X Tube",
		},
		{
			question:        "Qual projeto mostra melhor a experiência dele com sistemas complexos?",
			expectedProject: "X Tube",
		},
		{
			question:        "Que projeto você destacaria para um recrutador?",
			expectedProject: "Auronix",
		},
	}

	for _, test := range tests {
		t.Run(test.question, func(t *testing.T) {
			response, hasResponse, language := AgentResponse(test.question)

			if !hasResponse {
				t.Fatalf("AgentResponse(%q) did not find a response", test.question)
			}

			if language != "pt" {
				t.Fatalf("AgentResponse(%q) language = %q, want %q", test.question, language, "pt")
			}

			if !strings.Contains(response, test.expectedProject) {
				t.Fatalf("AgentResponse(%q) = %q, want it to select %s",
					test.question,
					response,
					test.expectedProject,
				)
			}

			if strings.Contains(response, "projetos prioritários") {
				t.Fatalf("AgentResponse(%q) = %q, contains internal base wording", test.question, response)
			}
		})
	}
}

func TestAgentResponseProjectRecommendationByCriterion(t *testing.T) {
	tests := []struct {
		name           string
		question       string
		expectedTerms  []string
		forbiddenTerms []string
	}{
		{
			name:     "financial systems",
			question: "Qual projeto melhor demonstra experiência com sistemas financeiros?",
			expectedTerms: []string{
				"Auronix",
				"transferências confiáveis",
				"ledger auditável",
			},
			forbiddenTerms: []string{"projetos prioritários", "X Tube demonstra"},
		},
		{
			name:     "technical capability",
			question: "Qual projeto do João Paulo melhor demonstra capacidade técnica?",
			expectedTerms: []string{
				"Auronix",
				"capacidade técnica",
				"plataforma financeira",
				"ledger auditável",
			},
			forbiddenTerms: []string{"projetos prioritários", "X Tube é"},
		},
		{
			name:     "technical leadership",
			question: "Qual projeto melhor demonstra liderança técnica?",
			expectedTerms: []string{
				"X Tube",
				"equipe de três",
				"upload",
				"processamento",
				"entrega",
			},
			forbiddenTerms: []string{"Auronix é", "GGCompress é"},
		},
		{
			name:     "go knowledge",
			question: "Qual projeto melhor demonstra conhecimento em Go?",
			expectedTerms: []string{
				"GGCompress",
				"Go",
				"concorrente",
			},
			forbiddenTerms: []string{"projetos prioritários"},
		},
		{
			name:     "auditability",
			question: "Qual projeto melhor demonstra preocupação com auditabilidade?",
			expectedTerms: []string{
				"Auditex",
				"alterações",
				"validação histórica",
			},
			forbiddenTerms: []string{"descentralizada", "GGCompress"},
		},
		{
			name:     "performance and concurrency",
			question: "Qual projeto demonstra melhor desempenho e concorrência?",
			expectedTerms: []string{
				"GGCompress",
				"concorrente",
				"1.23 GB/s",
				"9.77 GB",
			},
			forbiddenTerms: []string{"Auditex", "blockchain"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response, hasResponse, language := AgentResponse(test.question)

			if !hasResponse {
				t.Fatalf("AgentResponse(%q) did not find a response", test.question)
			}

			if language != "pt" {
				t.Fatalf("AgentResponse(%q) language = %q, want %q", test.question, language, "pt")
			}

			for _, term := range test.expectedTerms {
				if !strings.Contains(response, term) {
					t.Fatalf("AgentResponse(%q) = %q, want it to contain %q",
						test.question,
						response,
						term,
					)
				}
			}

			for _, term := range test.forbiddenTerms {
				if strings.Contains(response, term) {
					t.Fatalf("AgentResponse(%q) = %q, want it not to contain %q",
						test.question,
						response,
						term,
					)
				}
			}
		})
	}
}
