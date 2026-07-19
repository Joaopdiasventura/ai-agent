package app

import (
	"ai-agent/internal/nlp"
	"strings"
	"testing"
)

func TestAgentResponseReturnsPortugueseAnswerForPortugueseQuestion(t *testing.T) {
	response, hasResponse, language := AgentResponse("Me fale sobre o auronix")

	if !hasResponse {
		t.Fatal("AgentResponse() did not find a response")
	}

	if language != nlp.LanguagePortuguese {
		t.Fatalf("AgentResponse() language = %q, want %q", language, nlp.LanguagePortuguese)
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

	if language != nlp.LanguageEnglish {
		t.Fatalf("AgentResponse() language = %q, want %q", language, nlp.LanguageEnglish)
	}

	if !strings.Contains(response, "Auronix") {
		t.Fatalf("AgentResponse() response %q does not mention Auronix", response)
	}
}

func TestAgentResponseReturnsLocalizedFallbackForUnknownQuestion(t *testing.T) {
	tests := []struct {
		name             string
		question         string
		expectedLanguage nlp.Language
		expectedFallback string
	}{
		{
			name:             "portuguese unknown question",
			question:         "Pergunta sem relação",
			expectedLanguage: nlp.LanguagePortuguese,
			expectedFallback: "Não encontrei essa informação específica",
		},
		{
			name:             "english unknown question",
			question:         "Unrelated question",
			expectedLanguage: nlp.LanguageEnglish,
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
		language       nlp.Language
		expectedTerms  []string
		forbiddenTerms []string
	}{
		{
			name:          "current education",
			question:      "onde joão estuda?",
			language:      nlp.LanguagePortuguese,
			expectedTerms: []string{"FIAP", "Inteligência Artificial"},
			forbiddenTerms: []string{
				"FIAP segue estudando",
				"Etec",
			},
		},
		{
			name:          "past education",
			question:      "onde joão estudou?",
			language:      nlp.LanguagePortuguese,
			expectedTerms: []string{"Etec de Guarulhos", "Desenvolvimento de Sistemas"},
			forbiddenTerms: []string{
				"Engenheiro de Software Full Stack",
				"trabalha como",
			},
		},
		{
			name:          "email pronoun",
			question:      "qual o email dele?",
			language:      nlp.LanguagePortuguese,
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
