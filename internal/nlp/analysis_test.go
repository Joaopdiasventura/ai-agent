package nlp

import "testing"

func TestAnalyzeQueryDetectsCategoryAndTemporalContext(t *testing.T) {
	tests := []struct {
		name             string
		tokens           []string
		expectedIntent   Intent
		expectedCategory CategoryHint
		expectedTemporal TemporalContext
	}{
		{
			name:             "current education",
			tokens:           []string{"onde", "joão", "estuda"},
			expectedIntent:   IntentEducation,
			expectedCategory: CategoryHintEducation,
			expectedTemporal: TemporalPresent,
		},
		{
			name:             "past education",
			tokens:           []string{"onde", "joão", "estudou"},
			expectedIntent:   IntentEducation,
			expectedCategory: CategoryHintEducation,
			expectedTemporal: TemporalPast,
		},
		{
			name:             "contact email",
			tokens:           []string{"qual", "o", "email", "dele"},
			expectedIntent:   IntentContact,
			expectedCategory: CategoryHintContact,
			expectedTemporal: TemporalUnspecified,
		},
		{
			name:             "first job",
			tokens:           []string{"qual", "foi", "o", "primeiro", "emprego"},
			expectedIntent:   IntentFirstJob,
			expectedCategory: CategoryHintCareer,
			expectedTemporal: TemporalFirst,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			entity, hasEntity := DetectEntity(test.tokens)
			analysis := AnalyzeQuery(test.tokens, entity, hasEntity, LanguagePortuguese)

			if analysis.PrimaryIntent != test.expectedIntent {
				t.Fatalf("PrimaryIntent = %q, want %q", analysis.PrimaryIntent, test.expectedIntent)
			}

			if analysis.CategoryHint != test.expectedCategory {
				t.Fatalf("CategoryHint = %q, want %q", analysis.CategoryHint, test.expectedCategory)
			}

			if analysis.TemporalContext != test.expectedTemporal {
				t.Fatalf("TemporalContext = %q, want %q", analysis.TemporalContext, test.expectedTemporal)
			}
		})
	}
}

func TestDetectEntityResolvesPersonReferences(t *testing.T) {
	tests := [][]string{
		{"joão"},
		{"joao", "paulo"},
		{"ele"},
		{"dele"},
		{"o", "engenheiro"},
		{"o", "profissional"},
	}

	for _, tokens := range tests {
		entity, hasEntity := DetectEntity(tokens)

		if !hasEntity {
			t.Fatalf("DetectEntity(%v) did not find an entity", tokens)
		}

		if entity.Type != EntityPerson || entity.Value != "João Paulo" {
			t.Fatalf("DetectEntity(%v) = %#v, want João Paulo person entity", tokens, entity)
		}
	}
}
