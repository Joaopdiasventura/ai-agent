package queryanalysis

import (
	"ai-agent/internal/domain"
	"testing"
)

func TestAnalyzeDetectsLanguageCategoryAndTemporalStatus(t *testing.T) {
	tests := []struct {
		question       string
		language       string
		category       string
		temporalStatus string
	}{
		{
			question:       "Onde João estuda?",
			language:       "pt",
			category:       "education",
			temporalStatus: domain.TemporalFuture,
		},
		{
			question:       "Onde João estudou?",
			language:       "pt",
			category:       "education",
			temporalStatus: domain.TemporalPast,
		},
		{
			question:       "Qual é o email dele?",
			language:       "pt",
			category:       "contact",
			temporalStatus: domain.TemporalTimeless,
		},
		{
			question:       "Where did João work before?",
			language:       "en",
			category:       "career",
			temporalStatus: domain.TemporalPast,
		},
	}

	for _, test := range tests {
		t.Run(test.question, func(t *testing.T) {
			query := Analyze(test.question)

			if query.Language != test.language {
				t.Fatalf("Language = %q, want %q", query.Language, test.language)
			}

			if query.Category != test.category {
				t.Fatalf("Category = %q, want %q", query.Category, test.category)
			}

			if query.TemporalStatus != test.temporalStatus {
				t.Fatalf("TemporalStatus = %q, want %q", query.TemporalStatus, test.temporalStatus)
			}
		})
	}
}

func TestAnalyzeDetectsProjectAndExactTerms(t *testing.T) {
	query := Analyze("Me fale sobre o Auronix e o email joaopdias.dev@gmail.com")

	if query.Project != "auronix" {
		t.Fatalf("Project = %q, want auronix", query.Project)
	}

	if len(query.ExactTerms) == 0 {
		t.Fatal("ExactTerms is empty")
	}
}
